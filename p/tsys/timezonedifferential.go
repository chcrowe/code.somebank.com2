package tsys

import (
	"fmt"

	"code.somebank.com/p/bytes"
)

type TimeZoneDirectionType uint

const (
	HoursAheadOfGMT                           TimeZoneDirectionType = 0
	QuarterHoursAheadOfGMT                    TimeZoneDirectionType = 2
	QuarterHoursAheadOfGMTWithDaylightSavings TimeZoneDirectionType = 4
	HoursAheadOfGMTWithDaylightSavings        TimeZoneDirectionType = 6
	HoursBehindGMT                            TimeZoneDirectionType = 1
	QuarterHoursBehindGMT                     TimeZoneDirectionType = 3
	QuarterHoursBehindGMTWithDaylightSavings  TimeZoneDirectionType = 5
	HoursBehindGMTWithDaylightSavings         TimeZoneDirectionType = 7
)

func (tz TimeZoneDirectionType) String() string {
	switch tz {

	case HoursAheadOfGMT:
		return "HoursAheadOfGMT"
	case QuarterHoursAheadOfGMT:
		return "QuarterHoursAheadOfGMT"
	case QuarterHoursAheadOfGMTWithDaylightSavings:
		return "QuarterHoursAheadOfGMTWithDaylightSavings"
	case HoursAheadOfGMTWithDaylightSavings:
		return "HoursAheadOfGMTWithDaylightSavings"
	case HoursBehindGMT:
		return "HoursBehindGMT"
	case QuarterHoursBehindGMT:
		return "QuarterHoursBehindGMT"
	case QuarterHoursBehindGMTWithDaylightSavings:
		return "QuarterHoursBehindGMTWithDaylightSavings"
	case HoursBehindGMTWithDaylightSavings:
		return "HoursBehindGMTWithDaylightSavings"

	}
	return ""
}

type TimeZoneDifferential struct {
	TimeZoneDirection  TimeZoneDirectionType
	TimeZoneDifference uint
}

func NewTimeZoneDifferential(direction TimeZoneDirectionType, zonediff uint) *TimeZoneDifferential {
	return &TimeZoneDifferential{direction, zonediff}
}

func NewTimeZoneDifferentialFromUint(tmzonediff uint) *TimeZoneDifferential {

	tz := TimeZoneDifferential{HoursBehindGMTWithDaylightSavings, 0}

	fmt.Sscanf(fmt.Sprintf("%03d", tmzonediff), "%1d%02d", &tz.TimeZoneDirection, &tz.TimeZoneDifference)

	return &tz
}

func (t *TimeZoneDifferential) String() string {
	return fmt.Sprintf("%d%02d", t.TimeZoneDirection, t.TimeZoneDifference)
}

func (t *TimeZoneDifferential) VerboseString() string {
	buffer := bytes.NewSafeBuffer(512)
	buffer.AppendFormat("\n%8s{\n", " ")

	var hrs float32

	switch t.TimeZoneDirection {
	case HoursAheadOfGMT, HoursAheadOfGMTWithDaylightSavings:
		hrs = float32(t.TimeZoneDifference)
	case HoursBehindGMT, HoursBehindGMTWithDaylightSavings:
		hrs = (-1.0) * float32(t.TimeZoneDifference)
	case QuarterHoursAheadOfGMT, QuarterHoursAheadOfGMTWithDaylightSavings:
		hrs = (0.25) * float32(t.TimeZoneDifference)
	case QuarterHoursBehindGMT, QuarterHoursBehindGMTWithDaylightSavings:
		hrs = (-0.25) * float32(t.TimeZoneDifference)
	}

	buffer.AppendFormat("%8s%-32s%d (%[3]s)\n", " ", "TimeZoneDirection", t.TimeZoneDirection)
	buffer.AppendFormat("%8s%-32s%02d (GMT %+02.2f HRS)\n", " ", "TimeZoneDifference", t.TimeZoneDifference, hrs)

	buffer.AppendFormat("%8s}", " ")
	return buffer.String()
}

func ParseTimeZoneDifferential(s string) *TimeZoneDifferential {
	tz := TimeZoneDifferential{}

	if 3 > len(s) {
		return &tz
	}

	fmt.Sscanf(s, "%1d%02d", &tz.TimeZoneDirection, &tz.TimeZoneDifference)

	return &tz
}
