package tsys

import (
	"code.google.com/p/go-uuid/uuid"
	"code.somebank.com/p/bytes"

	//	"fmt"
	"strings"
	"time"
)

type BatchMessageRequest struct {
	BatchUuid     string
	BatchDateTime time.Time
	Header        *BatchHeaderRecord
	Parameter     *BatchParameterRecord
	Details       []BatchDetailRecord
	Trailer       *BatchTrailerRecord
}

func NewBatchMessageRequest(id string,
	batchdatetimestamp time.Time,
	head *BatchHeaderRecord,
	param *BatchParameterRecord,
	detailrecords []BatchDetailRecord) *BatchMessageRequest {

	b := BatchMessageRequest{}
	b.BatchDateTime = batchdatetimestamp //time.Now()

	if 0 < len(id) {
		b.BatchUuid = id
	} else {
		b.BatchUuid = uuid.New()
	}

	return &b
}

func (b *BatchMessageRequest) Bytes() []byte {

	buf := bytes.NewSafeBuffer(1024)

	buf.AppendBytes(bytes.FrameAndEncodePacket(128, b.Header.Bytes(), byte(bytes.EndOfTransmissionBlock)))
	buf.AppendBytes(bytes.FrameAndEncodePacket(128, b.Parameter.Bytes(), byte(bytes.EndOfTransmissionBlock)))

	for _, detail := range b.Details {
		//append an encrypted transaction block record?
		buf.AppendBytes(bytes.FrameAndEncodePacket(512, detail.Bytes(), byte(bytes.EndOfTransmissionBlock)))
		//append a chip card addendum?
	}

	buf.AppendBytes(bytes.FrameAndEncodePacket(128, b.Trailer.Bytes(), byte(bytes.EndOfText)))

	return buf.Bytes()
}

func ParseBatchMessageRequest(s string) *BatchMessageRequest {

	b := BatchMessageRequest{}
	records := strings.Split(s, string(bytes.StartOfText))

	for _, rec := range records {

		if 4 < len(rec) {

			record := bytes.RemovePacketFraming(rec)

			switch RecordIndicatorType(record[4]) {
			case HeaderRecord: // 'H'
				b.Header, _ = ParseBatchHeaderRecord(record)
			case ParameterRecord: // 'P'
				b.Parameter, _ = ParseBatchParameterRecord(record)
			case DetailRecord: // 'D'
				detail, _ := ParseBatchDetailRecord(record)
				b.Details = append(b.Details, *detail)
			// case AmericanExpressLineItemDetailRecord: // 'A'
			// case CommercialCardLineItemDetailRecord: // 'L'
			// case CommercialCardTripLegDetailRecord: // 'M'
			// case FleetLineItemDetailRecord: // 'F'
			case EncryptedTransmissionBlockRecord: // 'E'
			case ChipCardAddendumRecord: // 'C'
			case TrailerRecord: // 'T'
				b.Trailer, _ = ParseBatchTrailerRecord(record)
				// case TrailerResponseRecord: // 'R'
				// case GroupMapExtensionRecord: // 'X'
			}
		}
	}

	return &b
}
