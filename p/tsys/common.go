package tsys

//1080/1081 FIELD TYPE ENUMERATIONS
type AccountDataSourceType rune

const (
	SpaceOrEmptyAccountDataSource                           AccountDataSourceType = ' '
	BarCodeRead                                             AccountDataSourceType = 'A'
	MicrDataWasAcquiredByOcrReader                          AccountDataSourceType = 'B'
	CheckImagingDevicePosCheckService                       AccountDataSourceType = 'C'
	FullMagneticStripeReadAndTransmitTrack2                 AccountDataSourceType = 'D'
	ChipCardReadDataCCPS                                    AccountDataSourceType = 'G'
	FullMagneticStripeReadAndTransmitTrack1                 AccountDataSourceType = 'H'
	ManuallyKeyedChipCardReadCapableTerminal                AccountDataSourceType = 'P'
	ProximityPaymentUsingTrackDataRules                     AccountDataSourceType = 'Q'
	ProximityPaymentUsingEmvRules                           AccountDataSourceType = 'R'
	FullMagneticStripeReadChipCardCapableTerminal           AccountDataSourceType = 'S'
	ManuallyKeyedTrack2Capable                              AccountDataSourceType = 'T'
	ManuallyKeyedTrack1Capable                              AccountDataSourceType = 'X'
	ManuallyKeyedTerminalHasNoCardReadingCapability         AccountDataSourceType = '@'
	ChipCardTransProcessedAsMagStripe_NoEmvAppOnTerminal    AccountDataSourceType = 'W'
	ChipCardTransProcessedAsMagStripe_CardOrTerminalFailure AccountDataSourceType = 'Z'
)

func (x AccountDataSourceType) String() string {
	switch x {
	case SpaceOrEmptyAccountDataSource:
		return "SpaceOrEmptyAccountDataSource"
	case BarCodeRead:
		return "BarCodeRead"
	case MicrDataWasAcquiredByOcrReader:
		return "MicrDataWasAcquiredByOcrReader"
	case CheckImagingDevicePosCheckService:
		return "CheckImagingDevicePosCheckService"
	case FullMagneticStripeReadAndTransmitTrack2:
		return "FullMagneticStripeReadAndTransmitTrack2"
	case ChipCardReadDataCCPS:
		return "ChipCardReadDataCCPS"
	case FullMagneticStripeReadAndTransmitTrack1:
		return "FullMagneticStripeReadAndTransmitTrack1"
	case ManuallyKeyedChipCardReadCapableTerminal:
		return "ManuallyKeyedChipCardReadCapableTerminal"
	case ProximityPaymentUsingTrackDataRules:
		return "ProximityPaymentUsingTrackDataRules"
	case ProximityPaymentUsingEmvRules:
		return "ProximityPaymentUsingEmvRules"
	case FullMagneticStripeReadChipCardCapableTerminal:
		return "FullMagneticStripeReadChipCardCapableTerminal"
	case ManuallyKeyedTrack2Capable:
		return "ManuallyKeyedTrack2Capable"
	case ManuallyKeyedTrack1Capable:
		return "ManuallyKeyedTrack1Capable"
	case ManuallyKeyedTerminalHasNoCardReadingCapability:
		return "ManuallyKeyedTerminalHasNoCardReadingCapability"
	case ChipCardTransProcessedAsMagStripe_NoEmvAppOnTerminal:
		return "ChipCardTransProcessedAsMagStripe_NoEmvAppOnTerminal"
	case ChipCardTransProcessedAsMagStripe_CardOrTerminalFailure:
		return "ChipCardTransProcessedAsMagStripe_CardOrTerminalFailure"
	}
	return ""
}

type AdditionalAmountSignType rune

const (
	PositiveBalance AdditionalAmountSignType = 'C'
	NegativeBalance AdditionalAmountSignType = 'D'
)

func (x AdditionalAmountSignType) String() string {
	switch x {
	case PositiveBalance:
		return "PositiveBalance"
	case NegativeBalance:
		return "NegativeBalance"
	}
	return ""
}

type AdditionalAmountAccountType uint

const (
	AccountNotSpecified AdditionalAmountAccountType = 0
	SavingsAccount      AdditionalAmountAccountType = 10
	CheckingAccount     AdditionalAmountAccountType = 20
	CreditCardAccount   AdditionalAmountAccountType = 30
	UniversalAccount    AdditionalAmountAccountType = 40
	StoredValueAccount  AdditionalAmountAccountType = 60
	CashBenefitsAccount AdditionalAmountAccountType = 96 // (for use by Electronic Benefits Transfer transactions only)
	FoodStampsAccount   AdditionalAmountAccountType = 98 // (for use by Electronic Benefits Transfer transactions only)
)

func (x AdditionalAmountAccountType) String() string {
	switch x {
	case AccountNotSpecified:
		return "AccountNotSpecified"
	case SavingsAccount:
		return "SavingsAccount"
	case CheckingAccount:
		return "CheckingAccount"
	case CreditCardAccount:
		return "CreditCardAccount"
	case UniversalAccount:
		return "UniversalAccount"
	case StoredValueAccount:
		return "StoredValueAccount"
	case CashBenefitsAccount:
		return "CashBenefitsAccount"
	case FoodStampsAccount:
		return "FoodStampsAccount"
	}
	return ""
}

type AdditionalAmountType string

const (
	DepositAccountLedgerBalance    AdditionalAmountType = "01" // Credit card account: Open to buy
	DepositAccountAvailableBalance AdditionalAmountType = "02" // Credit card account: Credit limit
	AmexPrepaidCard                AdditionalAmountType = "05" // available amount remaining (Amex only)
	AmountHealthcare               AdditionalAmountType = "4S" // (Visa authorization request only)
	AmountTransit                  AdditionalAmountType = "4T" // (Visa authorization request only)
	AmountCoPayment                AdditionalAmountType = "3S"
	OriginalAmount                 AdditionalAmountType = "57"
	PartialAuthorizedAmount        AdditionalAmountType = "10"
	AmountPrescriptionRx           AdditionalAmountType = "4U" // (Visa and MC authorization requests)
	AmountVisionOptical            AdditionalAmountType = "4V" // (Visa authorization request only)
	AmountClinic                   AdditionalAmountType = "4W" // other qualified medical (Visa authorization request only)
	AmountDental                   AdditionalAmountType = "4X" // (Visa authorization request only)
	AmountCashOver                 AdditionalAmountType = "80" // Discover only
	OriginalAmountCashOver         AdditionalAmountType = "81" // Discover Only
)

func (x AdditionalAmountType) String() string {
	switch x {
	case DepositAccountLedgerBalance:
		return "DepositAccountLedgerBalance"
	case DepositAccountAvailableBalance:
		return "DepositAccountAvailableBalance"
	case AmexPrepaidCard:
		return "AmexPrepaidCard"
	case AmountHealthcare:
		return "AmountHealthcare"
	case AmountTransit:
		return "AmountTransit"
	case AmountCoPayment:
		return "AmountCoPayment"
	case OriginalAmount:
		return "OriginalAmount"
	case PartialAuthorizedAmount:
		return "PartialAuthorizedAmount"
	case AmountPrescriptionRx:
		return "AmountPrescriptionRx"
	case AmountVisionOptical:
		return "AmountVisionOptical"
	case AmountClinic:
		return "AmountClinic"
	case AmountDental:
		return "AmountDental"
	case AmountCashOver:
		return "AmountCashOver"
	case OriginalAmountCashOver:
		return "OriginalAmountCashOver"
	}
	return ""
}

type AdllType rune

const (
	SpaceOrEmptyAdll AdllType = ' '
	PartialDownload  AdllType = 'P'
	FullDownload     AdllType = 'F'
)

func (x AdllType) String() string {
	switch x {
	case SpaceOrEmptyAdll:
		return "SpaceOrEmptyAdll"
	case PartialDownload:
		return "PartialDownload"
	case FullDownload:
		return "FullDownload"
	}
	return ""
}

type AddressVerificationResultCodeType rune

const (
	ApprovedAddressVerificationNotRequested_0           AddressVerificationResultCodeType = '0'
	AddressMatchAddressMatchOnly_A                      AddressVerificationResultCodeType = 'A'
	AddressMatchStreetAddressMatchInternational_B       AddressVerificationResultCodeType = 'B'
	ServUnavailableStreetAndPostalCodeNotVerified_C     AddressVerificationResultCodeType = 'C'
	ExactMatchStreetAddressMatch_D                      AddressVerificationResultCodeType = 'D'
	ExactMatchStreetAddressAndPostalCodeMatch_F         AddressVerificationResultCodeType = 'F'
	VerUnavailableIssuerDoesNotParticipate_G            AddressVerificationResultCodeType = 'G'
	VerUnavailable_I                                    AddressVerificationResultCodeType = 'I'
	ExactMatchStreetAddressMatchInternational_M         AddressVerificationResultCodeType = 'M'
	NoMatchNoAddressOrZIPMatch_N                        AddressVerificationResultCodeType = 'N'
	ZipMatchPostalCodesMatchInternational_P             AddressVerificationResultCodeType = 'P'
	RetryIssuerSystemUnavailable_R                      AddressVerificationResultCodeType = 'R'
	ServUnavailableServiceNotSupported_S                AddressVerificationResultCodeType = 'S'
	VerUnavailableAddressUnavailable_U                  AddressVerificationResultCodeType = 'U'
	ZipMatchNineCharacterNumericZIPMatchOnly_W          AddressVerificationResultCodeType = 'W'
	ExactMatchExactMatchNineCharacterNumericZIP_X       AddressVerificationResultCodeType = 'X'
	ExactMatchExactMatchFiveCharacterNumericZIP_Y       AddressVerificationResultCodeType = 'Y'
	ZipMatchFiveCharacterNumericZIPMatchOnly_Z          AddressVerificationResultCodeType = 'Z'
	CardholderNameAndZIPMatchAMEXOnly_1                 AddressVerificationResultCodeType = '1'
	CardholderNameAddressAndZIPMatchAMEXOnly_2          AddressVerificationResultCodeType = '2'
	CardholderNameAndAddressMatchAMEXOnly_3             AddressVerificationResultCodeType = '3'
	CardholderNameMatchAMEXOnly_4                       AddressVerificationResultCodeType = '4'
	CardholderNameIncorrectZIPMatchAMEXOnly_5           AddressVerificationResultCodeType = '5'
	CardholderNameIncorrectAddressAndZIPMatchAMEXOnly_6 AddressVerificationResultCodeType = '6'
	CardholderNameIncorrectAddressMatchAMEXOnly_7       AddressVerificationResultCodeType = '7'
	CardholderAllDoNotMatchAMEXOnly_8                   AddressVerificationResultCodeType = '8'
)

func (x AddressVerificationResultCodeType) String() string {
	switch x {
	case ApprovedAddressVerificationNotRequested_0:
		return "ApprovedAddressVerificationNotRequested_0"
	case AddressMatchAddressMatchOnly_A:
		return "AddressMatchAddressMatchOnly_A"
	case AddressMatchStreetAddressMatchInternational_B:
		return "AddressMatchStreetAddressMatchInternational_B"
	case ServUnavailableStreetAndPostalCodeNotVerified_C:
		return "ServUnavailableStreetAndPostalCodeNotVerified_C"
	case ExactMatchStreetAddressMatch_D:
		return "ExactMatchStreetAddressMatch_D"
	case ExactMatchStreetAddressAndPostalCodeMatch_F:
		return "ExactMatchStreetAddressAndPostalCodeMatch_F"
	case VerUnavailableIssuerDoesNotParticipate_G:
		return "VerUnavailableIssuerDoesNotParticipate_G"
	case VerUnavailable_I:
		return "VerUnavailable_I"
	case ExactMatchStreetAddressMatchInternational_M:
		return "ExactMatchStreetAddressMatchInternational_M"
	case NoMatchNoAddressOrZIPMatch_N:
		return "NoMatchNoAddressOrZIPMatch_N"
	case ZipMatchPostalCodesMatchInternational_P:
		return "ZipMatchPostalCodesMatchInternational_P"
	case RetryIssuerSystemUnavailable_R:
		return "RetryIssuerSystemUnavailable_R"
	case ServUnavailableServiceNotSupported_S:
		return "ServUnavailableServiceNotSupported_S"
	case VerUnavailableAddressUnavailable_U:
		return "VerUnavailableAddressUnavailable_U"
	case ZipMatchNineCharacterNumericZIPMatchOnly_W:
		return "ZipMatchNineCharacterNumericZIPMatchOnly_W"
	case ExactMatchExactMatchNineCharacterNumericZIP_X:
		return "ExactMatchExactMatchNineCharacterNumericZIP_X"
	case ExactMatchExactMatchFiveCharacterNumericZIP_Y:
		return "ExactMatchExactMatchFiveCharacterNumericZIP_Y"
	case ZipMatchFiveCharacterNumericZIPMatchOnly_Z:
		return "ZipMatchFiveCharacterNumericZIPMatchOnly_Z"
	case CardholderNameAndZIPMatchAMEXOnly_1:
		return "CardholderNameAndZIPMatchAMEXOnly_1"
	case CardholderNameAddressAndZIPMatchAMEXOnly_2:
		return "CardholderNameAddressAndZIPMatchAMEXOnly_2"
	case CardholderNameAndAddressMatchAMEXOnly_3:
		return "CardholderNameAndAddressMatchAMEXOnly_3"
	case CardholderNameMatchAMEXOnly_4:
		return "CardholderNameMatchAMEXOnly_4"
	case CardholderNameIncorrectZIPMatchAMEXOnly_5:
		return "CardholderNameIncorrectZIPMatchAMEXOnly_5"
	case CardholderNameIncorrectAddressAndZIPMatchAMEXOnly_6:
		return "CardholderNameIncorrectAddressAndZIPMatchAMEXOnly_6"
	case CardholderNameIncorrectAddressMatchAMEXOnly_7:
		return "CardholderNameIncorrectAddressMatchAMEXOnly_7"
	case CardholderAllDoNotMatchAMEXOnly_8:
		return "CardholderAllDoNotMatchAMEXOnly_8"
	}
	return ""
}

type AirportCodeType string

const (
	Unassigned                            AirportCodeType = " "
	AbiAbileneTx                          AirportCodeType = "ABI"
	AbqAlbuquerqueNm                      AirportCodeType = "ABQ"
	AcaAcapulcoMexico                     AirportCodeType = "ACA"
	ActWacoTx                             AirportCodeType = "ACT"
	AexAlexandriaLa                       AirportCodeType = "AEX"
	AklAucklandNewZealand                 AirportCodeType = "AKL"
	AlbAlbanyNy                           AirportCodeType = "ALB"
	AloWaterlooIa                         AirportCodeType = "ALO"
	AmaAmarilloTx                         AirportCodeType = "AMA"
	AmsAmsterdamNetherlands               AirportCodeType = "AMS"
	AnuAntigua                            AirportCodeType = "ANU"
	ApfNaplesFl                           AirportCodeType = "APF"
	ArnStockholmSweden                    AirportCodeType = "ARN"
	AsuAsuncionParaguay                   AirportCodeType = "ASU"
	AtlAtlantaGa                          AirportCodeType = "ATL"
	AuaAruba                              AirportCodeType = "AUA"
	AuhAbuDhabiUae                        AirportCodeType = "AUH"
	AusAustinTx                           AirportCodeType = "AUS"
	AvlAshevilleNc                        AirportCodeType = "AVL"
	AxaAnguilla                           AirportCodeType = "AXA"
	AzoKalamazooMi                        AirportCodeType = "AZO"
	BahBahrainBahrain                     AirportCodeType = "BAH"
	BaqBarranquillaColombia               AirportCodeType = "BAQ"
	BdaBermuda                            AirportCodeType = "BDA"
	BdlHartfordCtSpringfieldMa            AirportCodeType = "BDL"
	BflBakersfieldCa                      AirportCodeType = "BFL"
	BfsBelfastIrelandUk                   AirportCodeType = "BFS"
	BgiBarbados                           AirportCodeType = "BGI"
	BhmBirminghamAl                       AirportCodeType = "BHM"
	BhxBirminghamUk                       AirportCodeType = "BHX"
	BjxLeonGuanajuatoMexico               AirportCodeType = "BJX"
	BmiBloomingtonIl                      AirportCodeType = "BMI"
	BnaNashvilleTn                        AirportCodeType = "BNA"
	BogBogotaColombia                     AirportCodeType = "BOG"
	BosBostonMa                           AirportCodeType = "BOS"
	BptBeaumontPortArthurTx               AirportCodeType = "BPT"
	BqnAguadillaPuertoRico                AirportCodeType = "BQN"
	BruBrusselsBelgium                    AirportCodeType = "BRU"
	BtrBatonRougeLa                       AirportCodeType = "BTR"
	BufBuffaloNy                          AirportCodeType = "BUF"
	BurBurbankCa                          AirportCodeType = "BUR"
	BwiBaltimoreMd                        AirportCodeType = "BWI"
	BzeBelizeCityBelize                   AirportCodeType = "BZE"
	CaeColumbiaSc                         AirportCodeType = "CAE"
	CcsCaracasVenezuela                   AirportCodeType = "CCS"
	ChaChattanoogaTn                      AirportCodeType = "CHA"
	CidCedarRapidsIa                      AirportCodeType = "CID"
	CldCarlsbadCa                         AirportCodeType = "CLD"
	CleClevelandOh                        AirportCodeType = "CLE"
	CllCollegeStationTx                   AirportCodeType = "CLL"
	CloCaliColombia                       AirportCodeType = "CLO"
	CltCharlotteNc                        AirportCodeType = "CLT"
	CmhColumbusOh                         AirportCodeType = "CMH"
	CmiChampaignUrbanaIl                  AirportCodeType = "CMI"
	CnfBeloHorizonteBrazil                AirportCodeType = "CNF"
	CosColoradoSpringsCo                  AirportCodeType = "COS"
	CrpCorpusChristiTx                    AirportCodeType = "CRP"
	CunCancunMexico                       AirportCodeType = "CUN"
	CurCuracao                            AirportCodeType = "CUR"
	CvgCincinnatiOh                       AirportCodeType = "CVG"
	CwaWausauStevensPointWi               AirportCodeType = "CWA"
	DayDaytonOh                           AirportCodeType = "DAY"
	DbqDubuqueIa                          AirportCodeType = "DBQ"
	DcaWashingtonDcNational               AirportCodeType = "DCA"
	DenDenverCo                           AirportCodeType = "DEN"
	DfwDallasFortWorthTx                  AirportCodeType = "DFW"
	DohDohaQatar                          AirportCodeType = "DOH"
	DomDominicaDominica                   AirportCodeType = "DOM"
	DroDurangoCo                          AirportCodeType = "DRO"
	DsmDesMoinesIa                        AirportCodeType = "DSM"
	DtwDetroitMi                          AirportCodeType = "DTW"
	DusDuesseldorfGermany                 AirportCodeType = "DUS"
	EdiEdinburghScotland                  AirportCodeType = "EDI"
	EgeVailCo                             AirportCodeType = "EGE"
	EisTortolaVirginGordaBrvii            AirportCodeType = "EIS"
	ElpElPasoTexas                        AirportCodeType = "ELP"
	EvvEvansvilleIn                       AirportCodeType = "EVV"
	EwrNewYorknewark                      AirportCodeType = "EWR"
	EywKeyWestFl                          AirportCodeType = "EYW"
	EzeBuenosAiresArgentina               AirportCodeType = "EZE"
	FatFresnoCa                           AirportCodeType = "FAT"
	FdfFortDeFranceMartinique             AirportCodeType = "FDF"
	FllFortLauderdaleFl                   AirportCodeType = "FLL"
	FpoFreeportBahamas                    AirportCodeType = "FPO"
	FraFrankfurtGermany                   AirportCodeType = "FRA"
	FsmFortSmithAr                        AirportCodeType = "FSM"
	FwaFortWayneIn                        AirportCodeType = "FWA"
	FyvFayettevilleAr                     AirportCodeType = "FYV"
	GcmGrandCayman                        AirportCodeType = "GCM"
	GdlGuadalajaraMexico                  AirportCodeType = "GDL"
	GggLongviewTx                         AirportCodeType = "GGG"
	GgtGeorgeTownBahamas                  AirportCodeType = "GGT"
	GhbGovernorsHarbourBahamas            AirportCodeType = "GHB"
	GigRioDeJaneiroBrazil                 AirportCodeType = "GIG"
	GlaGlasgowScotland                    AirportCodeType = "GLA"
	GndGrenada                            AirportCodeType = "GND"
	GrbGreenBayWi                         AirportCodeType = "GRB"
	GrrGrandRapidsMi                      AirportCodeType = "GRR"
	GruSaoPauloBrazil                     AirportCodeType = "GRU"
	GsoGreensboroNc                       AirportCodeType = "GSO"
	GuaGuatemalaCityGuatemala             AirportCodeType = "GUA"
	GucGunnisonCrestedButteCo             AirportCodeType = "GUC"
	GyeGuayaquilEcuador                   AirportCodeType = "GYE"
	HdnSteamboatSpringsCo                 AirportCodeType = "HDN"
	HnlHonoluluOahuHi                     AirportCodeType = "HNL"
	HouHoustonTxhobby                     AirportCodeType = "HOU"
	HpnWhitePlainsNy                      AirportCodeType = "HPN"
	HrlHarlingenTx                        AirportCodeType = "HRL"
	HsvHuntsvilleDecaturAl                AirportCodeType = "HSV"
	IadWashingtonDcdulles                 AirportCodeType = "IAD"
	IahHoustonTxintercontinental          AirportCodeType = "IAH"
	IctWichitaKs                          AirportCodeType = "ICT"
	IleKilleenTx                          AirportCodeType = "ILE"
	IndIndianapolisIn                     AirportCodeType = "IND"
	IspLongIslandNyMacarthur              AirportCodeType = "ISP"
	IykInyokernCa                         AirportCodeType = "IYK"
	JacJacksonHoleWy                      AirportCodeType = "JAC"
	JanJacksonMs                          AirportCodeType = "JAN"
	JaxJacksonvilleFl                     AirportCodeType = "JAX"
	JfkNewYorkkennedy                     AirportCodeType = "JFK"
	JnbJohannesburgSouthAfrica            AirportCodeType = "JNB"
	KinKingstonJamaica                    AirportCodeType = "KIN"
	LanLansingMi                          AirportCodeType = "LAN"
	LasLasVegasNv                         AirportCodeType = "LAS"
	LawLawtonOk                           AirportCodeType = "LAW"
	LaxLosAngelesCa                       AirportCodeType = "LAX"
	LbaLeedsBradfordUk                    AirportCodeType = "LBA"
	LbbLubbockTx                          AirportCodeType = "LBB"
	LchLakeCharlesLa                      AirportCodeType = "LCH"
	LexLexingtonKy                        AirportCodeType = "LEX"
	LftLafayetteLa                        AirportCodeType = "LFT"
	LgaNewYorklaguardia                   AirportCodeType = "LGA"
	LgwLondonUkgatwick                    AirportCodeType = "LGW"
	LhrLondonUkheathrow                   AirportCodeType = "LHR"
	LimLimaPeru                           AirportCodeType = "LIM"
	LitLittleRockAr                       AirportCodeType = "LIT"
	LonLondonUk                           AirportCodeType = "LON"
	LpbLaPazBolivia                       AirportCodeType = "LPB"
	LrdLaredoTx                           AirportCodeType = "LRD"
	LrmCasaDeCampoLaRomanaDomrep          AirportCodeType = "LRM"
	LseLaCrosseWi                         AirportCodeType = "LSE"
	MadMadridSpain                        AirportCodeType = "MAD"
	MafMidlandOdessaTx                    AirportCodeType = "MAF"
	ManManchesterUk                       AirportCodeType = "MAN"
	MarMaracaiboVenezuela                 AirportCodeType = "MAR"
	MazMayaguezPuertoRico                 AirportCodeType = "MAZ"
	MbjMontegoBayJamaica                  AirportCodeType = "MBJ"
	MciKansasCityMo                       AirportCodeType = "MCI"
	McoOrlandoFl                          AirportCodeType = "MCO"
	MctMuscatOman                         AirportCodeType = "MCT"
	MdtHarrisburgPa                       AirportCodeType = "MDT"
	MelMelbourneAustralia                 AirportCodeType = "MEL"
	MemMemphisTn                          AirportCodeType = "MEM"
	MexMexicoCityMexico                   AirportCodeType = "MEX"
	MfeMcallenTx                          AirportCodeType = "MFE"
	MgaManaguaNicaragua                   AirportCodeType = "MGA"
	MhhMarshHarbourBahamas                AirportCodeType = "MHH"
	MiaMiamiFl                            AirportCodeType = "MIA"
	MkeMilwaukeeWi                        AirportCodeType = "MKE"
	MkgMuskegonMi                         AirportCodeType = "MKG"
	MlbMelbourneFl                        AirportCodeType = "MLB"
	MliMolineQuadCitiesIl                 AirportCodeType = "MLI"
	MluMonroeLa                           AirportCodeType = "MLU"
	MmeTeessideUk                         AirportCodeType = "MME"
	MobMobileAl                           AirportCodeType = "MOB"
	MqtMarquetteMi                        AirportCodeType = "MQT"
	MryMontereyCa                         AirportCodeType = "MRY"
	MsnMadisonWi                          AirportCodeType = "MSN"
	MspMinneapolisStPaulMn                AirportCodeType = "MSP"
	MsyNewOrleansLa                       AirportCodeType = "MSY"
	MthMarathonFl                         AirportCodeType = "MTH"
	MtyMonterreyMexico                    AirportCodeType = "MTY"
	MxpMilanItaly                         AirportCodeType = "MXP"
	NasNassauBahamas                      AirportCodeType = "NAS"
	NrtTokyoJapan                         AirportCodeType = "NRT"
	NycNewYork                            AirportCodeType = "NYC"
	OakOaklandCa                          AirportCodeType = "OAK"
	OggKahuluiMauiHi                      AirportCodeType = "OGG"
	OkcOklahomaCityOk                     AirportCodeType = "OKC"
	OmaOmahaNe                            AirportCodeType = "OMA"
	OntOntarioCa                          AirportCodeType = "ONT"
	OrdChicagoIl                          AirportCodeType = "ORD"
	OrfNorfolkVirginiaBeachVa             AirportCodeType = "ORF"
	OryParisFrance                        AirportCodeType = "ORY"
	OwbOwensboroKy                        AirportCodeType = "OWB"
	OxrOxnardCa                           AirportCodeType = "OXR"
	PahPaducahKy                          AirportCodeType = "PAH"
	PapPortAuPrinceHaiti                  AirportCodeType = "PAP"
	PbiWestPalmBeachFl                    AirportCodeType = "PBI"
	PdxPortlandOr                         AirportCodeType = "PDX"
	PhlPhiladelphiaPaWilmingtonDe         AirportCodeType = "PHL"
	PhxPhoenixAz                          AirportCodeType = "PHX"
	PiaPeoriaIl                           AirportCodeType = "PIA"
	PitPittsburghPa                       AirportCodeType = "PIT"
	PlsProvidencialesTurksCaicosIs        AirportCodeType = "PLS"
	PopPuertoPlataDomRep                  AirportCodeType = "POP"
	PosPortOfSpainTrinidadTobago          AirportCodeType = "POS"
	PsePoncePuertoRico                    AirportCodeType = "PSE"
	PspPalmSpringsCa                      AirportCodeType = "PSP"
	PtpPointeAPitreGuadeloupe             AirportCodeType = "PTP"
	PtyPanamaCityPanama                   AirportCodeType = "PTY"
	PujPuntaCanaDomRep                    AirportCodeType = "PUJ"
	PvdProvidenceNewportRi                AirportCodeType = "PVD"
	PvrPuertoVallartaMexico               AirportCodeType = "PVR"
	QhoHoustonTx                          AirportCodeType = "QHO"
	QslStLucia                            AirportCodeType = "QSL"
	RduRaleighDurhamNc                    AirportCodeType = "RDU"
	RfdRockfordIl                         AirportCodeType = "RFD"
	RicRichmondVa                         AirportCodeType = "RIC"
	RnoRenoNv                             AirportCodeType = "RNO"
	RocRochesterNy                        AirportCodeType = "ROC"
	RstRochesterMn                        AirportCodeType = "RST"
	RswFortMyersFl                        AirportCodeType = "RSW"
	SalSanSalvadorElSalvador              AirportCodeType = "SAL"
	SanSanDiegoCa                         AirportCodeType = "SAN"
	SapSanPedroSulaHonduras               AirportCodeType = "SAP"
	SatSanAntonioTx                       AirportCodeType = "SAT"
	SbaSantaBarbaraCa                     AirportCodeType = "SBA"
	SbnSouthBendIn                        AirportCodeType = "SBN"
	SbpSanLuisObispoCa                    AirportCodeType = "SBP"
	SclSantiagoChile                      AirportCodeType = "SCL"
	SdfLouisvilleKy                       AirportCodeType = "SDF"
	SdqSantoDomingoDomRep                 AirportCodeType = "SDQ"
	SeaSeattleTacomaWa                    AirportCodeType = "SEA"
	SfoSanFranciscoCa                     AirportCodeType = "SFO"
	SgfSpringfieldMo                      AirportCodeType = "SGF"
	ShvShreveportLa                       AirportCodeType = "SHV"
	SjcSanJoseCa                          AirportCodeType = "SJC"
	SjoSanJoseCostaRica                   AirportCodeType = "SJO"
	SjtSanAngeloTx                        AirportCodeType = "SJT"
	SjuSanJuanPuertoRico                  AirportCodeType = "SJU"
	SkbStKittsNevis                       AirportCodeType = "SKB"
	SlcSaltLakeCityUt                     AirportCodeType = "SLC"
	SluStLuciavigieField                  AirportCodeType = "SLU"
	SmfSacramentoCa                       AirportCodeType = "SMF"
	SmxSantaMariaCa                       AirportCodeType = "SMX"
	SnaOrangeCountyCa                     AirportCodeType = "SNA"
	SpiSpringfieldIl                      AirportCodeType = "SPI"
	SpsWichitaFallsTx                     AirportCodeType = "SPS"
	SrqSarasotaFl                         AirportCodeType = "SRQ"
	StlStLouisMo                          AirportCodeType = "STL"
	SttStThomasVi                         AirportCodeType = "STT"
	StxStCroixVi                          AirportCodeType = "STX"
	SwfNewburghNyStewartField             AirportCodeType = "SWF"
	SxmStMaarten                          AirportCodeType = "SXM"
	SydSydneyAustralia                    AirportCodeType = "SYD"
	SyrSyracuseNy                         AirportCodeType = "SYR"
	TcbTreasureCayBahamas                 AirportCodeType = "TCB"
	TclTuscaloosaAl                       AirportCodeType = "TCL"
	TguTegucigalpaHonduras                AirportCodeType = "TGU"
	TlhTallahasseeFl                      AirportCodeType = "TLH"
	TolToledoOh                           AirportCodeType = "TOL"
	TpaTampaFl                            AirportCodeType = "TPA"
	TplTempleTx                           AirportCodeType = "TPL"
	TriTriCityTn                          AirportCodeType = "TRI"
	TulTulsaOk                            AirportCodeType = "TUL"
	TupTupeloMs                           AirportCodeType = "TUP"
	TusTucsonAz                           AirportCodeType = "TUS"
	TvcTraverseCityMi                     AirportCodeType = "TVC"
	TxkTexarkanaAr                        AirportCodeType = "TXK"
	TyrTylerTx                            AirportCodeType = "TYR"
	TysKnoxvilleTn                        AirportCodeType = "TYS"
	UioQuitoEcuador                       AirportCodeType = "UIO"
	UvfStLuciaheWanorra                   AirportCodeType = "UVF"
	VrbVeroBeachFl                        AirportCodeType = "VRB"
	VviSantaCruzBolivia                   AirportCodeType = "VVI"
	WasWashingtonDc                       AirportCodeType = "WAS"
	YedEdmontonAlbertaCanada              AirportCodeType = "YED"
	YegEdmontonAlbertaCanadaInternational AirportCodeType = "YEG"
	YhzHalifaxNsCanada                    AirportCodeType = "YHZ"
	YowOttawaOnCanada                     AirportCodeType = "YOW"
	YqbQuebecCityPqCanada                 AirportCodeType = "YQB"
	YulMontrealPqCanada                   AirportCodeType = "YUL"
	YvrVancouverBcCanada                  AirportCodeType = "YVR"
	YwgWinnipegMbCanada                   AirportCodeType = "YWG"
	YxdEdmontonAlbertaCanadaMunicipal     AirportCodeType = "YXD"
	YycCalgaryAbCanada                    AirportCodeType = "YYC"
	YyzTorontoOnCanada                    AirportCodeType = "YYZ"
	ZrhZurichSwitzerland                  AirportCodeType = "ZRH"
)

func (x AirportCodeType) String() string {
	switch x {
	case Unassigned:
		return "Unassigned"
	case AbiAbileneTx:
		return "AbiAbileneTx"
	case AbqAlbuquerqueNm:
		return "AbqAlbuquerqueNm"
	case AcaAcapulcoMexico:
		return "AcaAcapulcoMexico"
	case ActWacoTx:
		return "ActWacoTx"
	case AexAlexandriaLa:
		return "AexAlexandriaLa"
	case AklAucklandNewZealand:
		return "AklAucklandNewZealand"
	case AlbAlbanyNy:
		return "AlbAlbanyNy"
	case AloWaterlooIa:
		return "AloWaterlooIa"
	case AmaAmarilloTx:
		return "AmaAmarilloTx"
	case AmsAmsterdamNetherlands:
		return "AmsAmsterdamNetherlands"
	case AnuAntigua:
		return "AnuAntigua"
	case ApfNaplesFl:
		return "ApfNaplesFl"
	case ArnStockholmSweden:
		return "ArnStockholmSweden"
	case AsuAsuncionParaguay:
		return "AsuAsuncionParaguay"
	case AtlAtlantaGa:
		return "AtlAtlantaGa"
	case AuaAruba:
		return "AuaAruba"
	case AuhAbuDhabiUae:
		return "AuhAbuDhabiUae"
	case AusAustinTx:
		return "AusAustinTx"
	case AvlAshevilleNc:
		return "AvlAshevilleNc"
	case AxaAnguilla:
		return "AxaAnguilla"
	case AzoKalamazooMi:
		return "AzoKalamazooMi"
	case BahBahrainBahrain:
		return "BahBahrainBahrain"
	case BaqBarranquillaColombia:
		return "BaqBarranquillaColombia"
	case BdaBermuda:
		return "BdaBermuda"
	case BdlHartfordCtSpringfieldMa:
		return "BdlHartfordCtSpringfieldMa"
	case BflBakersfieldCa:
		return "BflBakersfieldCa"
	case BfsBelfastIrelandUk:
		return "BfsBelfastIrelandUk"
	case BgiBarbados:
		return "BgiBarbados"
	case BhmBirminghamAl:
		return "BhmBirminghamAl"
	case BhxBirminghamUk:
		return "BhxBirminghamUk"
	case BjxLeonGuanajuatoMexico:
		return "BjxLeonGuanajuatoMexico"
	case BmiBloomingtonIl:
		return "BmiBloomingtonIl"
	case BnaNashvilleTn:
		return "BnaNashvilleTn"
	case BogBogotaColombia:
		return "BogBogotaColombia"
	case BosBostonMa:
		return "BosBostonMa"
	case BptBeaumontPortArthurTx:
		return "BptBeaumontPortArthurTx"
	case BqnAguadillaPuertoRico:
		return "BqnAguadillaPuertoRico"
	case BruBrusselsBelgium:
		return "BruBrusselsBelgium"
	case BtrBatonRougeLa:
		return "BtrBatonRougeLa"
	case BufBuffaloNy:
		return "BufBuffaloNy"
	case BurBurbankCa:
		return "BurBurbankCa"
	case BwiBaltimoreMd:
		return "BwiBaltimoreMd"
	case BzeBelizeCityBelize:
		return "BzeBelizeCityBelize"
	case CaeColumbiaSc:
		return "CaeColumbiaSc"
	case CcsCaracasVenezuela:
		return "CcsCaracasVenezuela"
	case ChaChattanoogaTn:
		return "ChaChattanoogaTn"
	case CidCedarRapidsIa:
		return "CidCedarRapidsIa"
	case CldCarlsbadCa:
		return "CldCarlsbadCa"
	case CleClevelandOh:
		return "CleClevelandOh"
	case CllCollegeStationTx:
		return "CllCollegeStationTx"
	case CloCaliColombia:
		return "CloCaliColombia"
	case CltCharlotteNc:
		return "CltCharlotteNc"
	case CmhColumbusOh:
		return "CmhColumbusOh"
	case CmiChampaignUrbanaIl:
		return "CmiChampaignUrbanaIl"
	case CnfBeloHorizonteBrazil:
		return "CnfBeloHorizonteBrazil"
	case CosColoradoSpringsCo:
		return "CosColoradoSpringsCo"
	case CrpCorpusChristiTx:
		return "CrpCorpusChristiTx"
	case CunCancunMexico:
		return "CunCancunMexico"
	case CurCuracao:
		return "CurCuracao"
	case CvgCincinnatiOh:
		return "CvgCincinnatiOh"
	case CwaWausauStevensPointWi:
		return "CwaWausauStevensPointWi"
	case DayDaytonOh:
		return "DayDaytonOh"
	case DbqDubuqueIa:
		return "DbqDubuqueIa"
	case DcaWashingtonDcNational:
		return "DcaWashingtonDcNational"
	case DenDenverCo:
		return "DenDenverCo"
	case DfwDallasFortWorthTx:
		return "DfwDallasFortWorthTx"
	case DohDohaQatar:
		return "DohDohaQatar"
	case DomDominicaDominica:
		return "DomDominicaDominica"
	case DroDurangoCo:
		return "DroDurangoCo"
	case DsmDesMoinesIa:
		return "DsmDesMoinesIa"
	case DtwDetroitMi:
		return "DtwDetroitMi"
	case DusDuesseldorfGermany:
		return "DusDuesseldorfGermany"
	case EdiEdinburghScotland:
		return "EdiEdinburghScotland"
	case EgeVailCo:
		return "EgeVailCo"
	case EisTortolaVirginGordaBrvii:
		return "EisTortolaVirginGordaBrvii"
	case ElpElPasoTexas:
		return "ElpElPasoTexas"
	case EvvEvansvilleIn:
		return "EvvEvansvilleIn"
	case EwrNewYorknewark:
		return "EwrNewYorknewark"
	case EywKeyWestFl:
		return "EywKeyWestFl"
	case EzeBuenosAiresArgentina:
		return "EzeBuenosAiresArgentina"
	case FatFresnoCa:
		return "FatFresnoCa"
	case FdfFortDeFranceMartinique:
		return "FdfFortDeFranceMartinique"
	case FllFortLauderdaleFl:
		return "FllFortLauderdaleFl"
	case FpoFreeportBahamas:
		return "FpoFreeportBahamas"
	case FraFrankfurtGermany:
		return "FraFrankfurtGermany"
	case FsmFortSmithAr:
		return "FsmFortSmithAr"
	case FwaFortWayneIn:
		return "FwaFortWayneIn"
	case FyvFayettevilleAr:
		return "FyvFayettevilleAr"
	case GcmGrandCayman:
		return "GcmGrandCayman"
	case GdlGuadalajaraMexico:
		return "GdlGuadalajaraMexico"
	case GggLongviewTx:
		return "GggLongviewTx"
	case GgtGeorgeTownBahamas:
		return "GgtGeorgeTownBahamas"
	case GhbGovernorsHarbourBahamas:
		return "GhbGovernorsHarbourBahamas"
	case GigRioDeJaneiroBrazil:
		return "GigRioDeJaneiroBrazil"
	case GlaGlasgowScotland:
		return "GlaGlasgowScotland"
	case GndGrenada:
		return "GndGrenada"
	case GrbGreenBayWi:
		return "GrbGreenBayWi"
	case GrrGrandRapidsMi:
		return "GrrGrandRapidsMi"
	case GruSaoPauloBrazil:
		return "GruSaoPauloBrazil"
	case GsoGreensboroNc:
		return "GsoGreensboroNc"
	case GuaGuatemalaCityGuatemala:
		return "GuaGuatemalaCityGuatemala"
	case GucGunnisonCrestedButteCo:
		return "GucGunnisonCrestedButteCo"
	case GyeGuayaquilEcuador:
		return "GyeGuayaquilEcuador"
	case HdnSteamboatSpringsCo:
		return "HdnSteamboatSpringsCo"
	case HnlHonoluluOahuHi:
		return "HnlHonoluluOahuHi"
	case HouHoustonTxhobby:
		return "HouHoustonTxhobby"
	case HpnWhitePlainsNy:
		return "HpnWhitePlainsNy"
	case HrlHarlingenTx:
		return "HrlHarlingenTx"
	case HsvHuntsvilleDecaturAl:
		return "HsvHuntsvilleDecaturAl"
	case IadWashingtonDcdulles:
		return "IadWashingtonDcdulles"
	case IahHoustonTxintercontinental:
		return "IahHoustonTxintercontinental"
	case IctWichitaKs:
		return "IctWichitaKs"
	case IleKilleenTx:
		return "IleKilleenTx"
	case IndIndianapolisIn:
		return "IndIndianapolisIn"
	case IspLongIslandNyMacarthur:
		return "IspLongIslandNyMacarthur"
	case IykInyokernCa:
		return "IykInyokernCa"
	case JacJacksonHoleWy:
		return "JacJacksonHoleWy"
	case JanJacksonMs:
		return "JanJacksonMs"
	case JaxJacksonvilleFl:
		return "JaxJacksonvilleFl"
	case JfkNewYorkkennedy:
		return "JfkNewYorkkennedy"
	case JnbJohannesburgSouthAfrica:
		return "JnbJohannesburgSouthAfrica"
	case KinKingstonJamaica:
		return "KinKingstonJamaica"
	case LanLansingMi:
		return "LanLansingMi"
	case LasLasVegasNv:
		return "LasLasVegasNv"
	case LawLawtonOk:
		return "LawLawtonOk"
	case LaxLosAngelesCa:
		return "LaxLosAngelesCa"
	case LbaLeedsBradfordUk:
		return "LbaLeedsBradfordUk"
	case LbbLubbockTx:
		return "LbbLubbockTx"
	case LchLakeCharlesLa:
		return "LchLakeCharlesLa"
	case LexLexingtonKy:
		return "LexLexingtonKy"
	case LftLafayetteLa:
		return "LftLafayetteLa"
	case LgaNewYorklaguardia:
		return "LgaNewYorklaguardia"
	case LgwLondonUkgatwick:
		return "LgwLondonUkgatwick"
	case LhrLondonUkheathrow:
		return "LhrLondonUkheathrow"
	case LimLimaPeru:
		return "LimLimaPeru"
	case LitLittleRockAr:
		return "LitLittleRockAr"
	case LonLondonUk:
		return "LonLondonUk"
	case LpbLaPazBolivia:
		return "LpbLaPazBolivia"
	case LrdLaredoTx:
		return "LrdLaredoTx"
	case LrmCasaDeCampoLaRomanaDomrep:
		return "LrmCasaDeCampoLaRomanaDomrep"
	case LseLaCrosseWi:
		return "LseLaCrosseWi"
	case MadMadridSpain:
		return "MadMadridSpain"
	case MafMidlandOdessaTx:
		return "MafMidlandOdessaTx"
	case ManManchesterUk:
		return "ManManchesterUk"
	case MarMaracaiboVenezuela:
		return "MarMaracaiboVenezuela"
	case MazMayaguezPuertoRico:
		return "MazMayaguezPuertoRico"
	case MbjMontegoBayJamaica:
		return "MbjMontegoBayJamaica"
	case MciKansasCityMo:
		return "MciKansasCityMo"
	case McoOrlandoFl:
		return "McoOrlandoFl"
	case MctMuscatOman:
		return "MctMuscatOman"
	case MdtHarrisburgPa:
		return "MdtHarrisburgPa"
	case MelMelbourneAustralia:
		return "MelMelbourneAustralia"
	case MemMemphisTn:
		return "MemMemphisTn"
	case MexMexicoCityMexico:
		return "MexMexicoCityMexico"
	case MfeMcallenTx:
		return "MfeMcallenTx"
	case MgaManaguaNicaragua:
		return "MgaManaguaNicaragua"
	case MhhMarshHarbourBahamas:
		return "MhhMarshHarbourBahamas"
	case MiaMiamiFl:
		return "MiaMiamiFl"
	case MkeMilwaukeeWi:
		return "MkeMilwaukeeWi"
	case MkgMuskegonMi:
		return "MkgMuskegonMi"
	case MlbMelbourneFl:
		return "MlbMelbourneFl"
	case MliMolineQuadCitiesIl:
		return "MliMolineQuadCitiesIl"
	case MluMonroeLa:
		return "MluMonroeLa"
	case MmeTeessideUk:
		return "MmeTeessideUk"
	case MobMobileAl:
		return "MobMobileAl"
	case MqtMarquetteMi:
		return "MqtMarquetteMi"
	case MryMontereyCa:
		return "MryMontereyCa"
	case MsnMadisonWi:
		return "MsnMadisonWi"
	case MspMinneapolisStPaulMn:
		return "MspMinneapolisStPaulMn"
	case MsyNewOrleansLa:
		return "MsyNewOrleansLa"
	case MthMarathonFl:
		return "MthMarathonFl"
	case MtyMonterreyMexico:
		return "MtyMonterreyMexico"
	case MxpMilanItaly:
		return "MxpMilanItaly"
	case NasNassauBahamas:
		return "NasNassauBahamas"
	case NrtTokyoJapan:
		return "NrtTokyoJapan"
	case NycNewYork:
		return "NycNewYork"
	case OakOaklandCa:
		return "OakOaklandCa"
	case OggKahuluiMauiHi:
		return "OggKahuluiMauiHi"
	case OkcOklahomaCityOk:
		return "OkcOklahomaCityOk"
	case OmaOmahaNe:
		return "OmaOmahaNe"
	case OntOntarioCa:
		return "OntOntarioCa"
	case OrdChicagoIl:
		return "OrdChicagoIl"
	case OrfNorfolkVirginiaBeachVa:
		return "OrfNorfolkVirginiaBeachVa"
	case OryParisFrance:
		return "OryParisFrance"
	case OwbOwensboroKy:
		return "OwbOwensboroKy"
	case OxrOxnardCa:
		return "OxrOxnardCa"
	case PahPaducahKy:
		return "PahPaducahKy"
	case PapPortAuPrinceHaiti:
		return "PapPortAuPrinceHaiti"
	case PbiWestPalmBeachFl:
		return "PbiWestPalmBeachFl"
	case PdxPortlandOr:
		return "PdxPortlandOr"
	case PhlPhiladelphiaPaWilmingtonDe:
		return "PhlPhiladelphiaPaWilmingtonDe"
	case PhxPhoenixAz:
		return "PhxPhoenixAz"
	case PiaPeoriaIl:
		return "PiaPeoriaIl"
	case PitPittsburghPa:
		return "PitPittsburghPa"
	case PlsProvidencialesTurksCaicosIs:
		return "PlsProvidencialesTurksCaicosIs"
	case PopPuertoPlataDomRep:
		return "PopPuertoPlataDomRep"
	case PosPortOfSpainTrinidadTobago:
		return "PosPortOfSpainTrinidadTobago"
	case PsePoncePuertoRico:
		return "PsePoncePuertoRico"
	case PspPalmSpringsCa:
		return "PspPalmSpringsCa"
	case PtpPointeAPitreGuadeloupe:
		return "PtpPointeAPitreGuadeloupe"
	case PtyPanamaCityPanama:
		return "PtyPanamaCityPanama"
	case PujPuntaCanaDomRep:
		return "PujPuntaCanaDomRep"
	case PvdProvidenceNewportRi:
		return "PvdProvidenceNewportRi"
	case PvrPuertoVallartaMexico:
		return "PvrPuertoVallartaMexico"
	case QhoHoustonTx:
		return "QhoHoustonTx"
	case QslStLucia:
		return "QslStLucia"
	case RduRaleighDurhamNc:
		return "RduRaleighDurhamNc"
	case RfdRockfordIl:
		return "RfdRockfordIl"
	case RicRichmondVa:
		return "RicRichmondVa"
	case RnoRenoNv:
		return "RnoRenoNv"
	case RocRochesterNy:
		return "RocRochesterNy"
	case RstRochesterMn:
		return "RstRochesterMn"
	case RswFortMyersFl:
		return "RswFortMyersFl"
	case SalSanSalvadorElSalvador:
		return "SalSanSalvadorElSalvador"
	case SanSanDiegoCa:
		return "SanSanDiegoCa"
	case SapSanPedroSulaHonduras:
		return "SapSanPedroSulaHonduras"
	case SatSanAntonioTx:
		return "SatSanAntonioTx"
	case SbaSantaBarbaraCa:
		return "SbaSantaBarbaraCa"
	case SbnSouthBendIn:
		return "SbnSouthBendIn"
	case SbpSanLuisObispoCa:
		return "SbpSanLuisObispoCa"
	case SclSantiagoChile:
		return "SclSantiagoChile"
	case SdfLouisvilleKy:
		return "SdfLouisvilleKy"
	case SdqSantoDomingoDomRep:
		return "SdqSantoDomingoDomRep"
	case SeaSeattleTacomaWa:
		return "SeaSeattleTacomaWa"
	case SfoSanFranciscoCa:
		return "SfoSanFranciscoCa"
	case SgfSpringfieldMo:
		return "SgfSpringfieldMo"
	case ShvShreveportLa:
		return "ShvShreveportLa"
	case SjcSanJoseCa:
		return "SjcSanJoseCa"
	case SjoSanJoseCostaRica:
		return "SjoSanJoseCostaRica"
	case SjtSanAngeloTx:
		return "SjtSanAngeloTx"
	case SjuSanJuanPuertoRico:
		return "SjuSanJuanPuertoRico"
	case SkbStKittsNevis:
		return "SkbStKittsNevis"
	case SlcSaltLakeCityUt:
		return "SlcSaltLakeCityUt"
	case SluStLuciavigieField:
		return "SluStLuciavigieField"
	case SmfSacramentoCa:
		return "SmfSacramentoCa"
	case SmxSantaMariaCa:
		return "SmxSantaMariaCa"
	case SnaOrangeCountyCa:
		return "SnaOrangeCountyCa"
	case SpiSpringfieldIl:
		return "SpiSpringfieldIl"
	case SpsWichitaFallsTx:
		return "SpsWichitaFallsTx"
	case SrqSarasotaFl:
		return "SrqSarasotaFl"
	case StlStLouisMo:
		return "StlStLouisMo"
	case SttStThomasVi:
		return "SttStThomasVi"
	case StxStCroixVi:
		return "StxStCroixVi"
	case SwfNewburghNyStewartField:
		return "SwfNewburghNyStewartField"
	case SxmStMaarten:
		return "SxmStMaarten"
	case SydSydneyAustralia:
		return "SydSydneyAustralia"
	case SyrSyracuseNy:
		return "SyrSyracuseNy"
	case TcbTreasureCayBahamas:
		return "TcbTreasureCayBahamas"
	case TclTuscaloosaAl:
		return "TclTuscaloosaAl"
	case TguTegucigalpaHonduras:
		return "TguTegucigalpaHonduras"
	case TlhTallahasseeFl:
		return "TlhTallahasseeFl"
	case TolToledoOh:
		return "TolToledoOh"
	case TpaTampaFl:
		return "TpaTampaFl"
	case TplTempleTx:
		return "TplTempleTx"
	case TriTriCityTn:
		return "TriTriCityTn"
	case TulTulsaOk:
		return "TulTulsaOk"
	case TupTupeloMs:
		return "TupTupeloMs"
	case TusTucsonAz:
		return "TusTucsonAz"
	case TvcTraverseCityMi:
		return "TvcTraverseCityMi"
	case TxkTexarkanaAr:
		return "TxkTexarkanaAr"
	case TyrTylerTx:
		return "TyrTylerTx"
	case TysKnoxvilleTn:
		return "TysKnoxvilleTn"
	case UioQuitoEcuador:
		return "UioQuitoEcuador"
	case UvfStLuciaheWanorra:
		return "UvfStLuciaheWanorra"
	case VrbVeroBeachFl:
		return "VrbVeroBeachFl"
	case VviSantaCruzBolivia:
		return "VviSantaCruzBolivia"
	case WasWashingtonDc:
		return "WasWashingtonDc"
	case YedEdmontonAlbertaCanada:
		return "YedEdmontonAlbertaCanada"
	case YegEdmontonAlbertaCanadaInternational:
		return "YegEdmontonAlbertaCanadaInternational"
	case YhzHalifaxNsCanada:
		return "YhzHalifaxNsCanada"
	case YowOttawaOnCanada:
		return "YowOttawaOnCanada"
	case YqbQuebecCityPqCanada:
		return "YqbQuebecCityPqCanada"
	case YulMontrealPqCanada:
		return "YulMontrealPqCanada"
	case YvrVancouverBcCanada:
		return "YvrVancouverBcCanada"
	case YwgWinnipegMbCanada:
		return "YwgWinnipegMbCanada"
	case YxdEdmontonAlbertaCanadaMunicipal:
		return "YxdEdmontonAlbertaCanadaMunicipal"
	case YycCalgaryAbCanada:
		return "YycCalgaryAbCanada"
	case YyzTorontoOnCanada:
		return "YyzTorontoOnCanada"
	case ZrhZurichSwitzerland:
		return "ZrhZurichSwitzerland"
	}
	return ""
}

type AmexSpecialProgramType rune

const (
	AmexPurchase             AmexSpecialProgramType = '1'
	AssuredReservationNoShow AmexSpecialProgramType = '2'
	CardDeposit              AmexSpecialProgramType = '3'
	DelayedCharge            AmexSpecialProgramType = '4'
	ExpressService           AmexSpecialProgramType = '5'
	AssuredReservation       AmexSpecialProgramType = '6'
	Other                    AmexSpecialProgramType = ' '
)

func (x AmexSpecialProgramType) String() string {
	switch x {
	case AmexPurchase:
		return "AmexPurchase"
	case AssuredReservationNoShow:
		return "AssuredReservationNoShow"
	case CardDeposit:
		return "CardDeposit"
	case DelayedCharge:
		return "DelayedCharge"
	case ExpressService:
		return "ExpressService"
	case AssuredReservation:
		return "AssuredReservation"
	case Other:
		return "Other"
	}
	return ""
}

type ApplicationIndicatorType uint

const (
	SingleAuthorizationPerConnection                         ApplicationIndicatorType = 0
	SingleBatchPerConnection                                 ApplicationIndicatorType = 1
	MutlipleAuthorizationsPerConnectionSingleThreaded        ApplicationIndicatorType = 2
	MultipleBatchPerConnection                               ApplicationIndicatorType = 3
	MultipleAuthorizationsPerConnectionFullDuplexInterleaved ApplicationIndicatorType = 4
	VitalCentralDataCapture05                                ApplicationIndicatorType = 5
	VitalCentralDataCapture06                                ApplicationIndicatorType = 6
	Reserved                                                 ApplicationIndicatorType = 9
)

func (x ApplicationIndicatorType) String() string {
	switch x {
	case SingleAuthorizationPerConnection:
		return "SingleAuthorizationPerConnection"
	case SingleBatchPerConnection:
		return "SingleBatchPerConnection"
	case MutlipleAuthorizationsPerConnectionSingleThreaded:
		return "MutlipleAuthorizationsPerConnectionSingleThreaded"
	case MultipleBatchPerConnection:
		return "MultipleBatchPerConnection"
	case MultipleAuthorizationsPerConnectionFullDuplexInterleaved:
		return "MultipleAuthorizationsPerConnectionFullDuplexInterleaved"
	case VitalCentralDataCapture05:
		return "VitalCentralDataCapture05"
	case VitalCentralDataCapture06:
		return "VitalCentralDataCapture06"
	case Reserved:
		return "Reserved"
	}
	return ""
}

type AuthorizationSourceType rune

const (
	SpaceOrEmptyAuthorizationSource                       AuthorizationSourceType = ' '
	StipStandinProcessingTimeoutResponse                  AuthorizationSourceType = '1'
	StipAmountBelowIssuerLimit                            AuthorizationSourceType = '2'
	StipIssuerInSuppressInquiryMode                       AuthorizationSourceType = '3'
	DirectConnectIssuerGeneratedResponse                  AuthorizationSourceType = '4'
	IssuerGeneratedResponse                               AuthorizationSourceType = '5'
	OffLineApprovalPosDeviceGenerated                     AuthorizationSourceType = '6'
	AcquirerApprovalBase1Unavailable                      AuthorizationSourceType = '7'
	AcquirerApprovalOfReferral                            AuthorizationSourceType = '8'
	CreditRefundOrNonauthorizedTransactions               AuthorizationSourceType = '9'
	ThirdPartyAuthorizingAgentPosCheckService             AuthorizationSourceType = 'A'
	ReferralAuthorizationCodeManuallyKeyed                AuthorizationSourceType = 'D'
	OffLineApprovalAuthorizationCodeManuallyKeyed         AuthorizationSourceType = 'E'
	CafisInterfaceOffLinePostAuthJapanAcquirerServicesJas AuthorizationSourceType = 'F'
	IssuerApprovalPostAuth                                AuthorizationSourceType = 'G'
)

func (x AuthorizationSourceType) String() string {
	switch x {
	case SpaceOrEmptyAuthorizationSource:
		return "SpaceOrEmptyAuthorizationSource"
	case StipStandinProcessingTimeoutResponse:
		return "StipStandinProcessingTimeoutResponse"
	case StipAmountBelowIssuerLimit:
		return "StipAmountBelowIssuerLimit"
	case StipIssuerInSuppressInquiryMode:
		return "StipIssuerInSuppressInquiryMode"
	case DirectConnectIssuerGeneratedResponse:
		return "DirectConnectIssuerGeneratedResponse"
	case IssuerGeneratedResponse:
		return "IssuerGeneratedResponse"
	case OffLineApprovalPosDeviceGenerated:
		return "OffLineApprovalPosDeviceGenerated"
	case AcquirerApprovalBase1Unavailable:
		return "AcquirerApprovalBase1Unavailable"
	case AcquirerApprovalOfReferral:
		return "AcquirerApprovalOfReferral"
	case CreditRefundOrNonauthorizedTransactions:
		return "CreditRefundOrNonauthorizedTransactions"
	case ThirdPartyAuthorizingAgentPosCheckService:
		return "ThirdPartyAuthorizingAgentPosCheckService"
	case ReferralAuthorizationCodeManuallyKeyed:
		return "ReferralAuthorizationCodeManuallyKeyed"
	case OffLineApprovalAuthorizationCodeManuallyKeyed:
		return "OffLineApprovalAuthorizationCodeManuallyKeyed"
	case CafisInterfaceOffLinePostAuthJapanAcquirerServicesJas:
		return "CafisInterfaceOffLinePostAuthJapanAcquirerServicesJas"
	case IssuerApprovalPostAuth:
		return "IssuerApprovalPostAuth"
	}
	return ""
}

type AvsResultType rune

const (
	SpaceOrEmptyAvsResult                    AvsResultType = ' '
	CreditOrOffline                          AvsResultType = '0'
	AddressMatchZipDoesNot                   AvsResultType = 'A'
	IneligibleTransactionNotMoto             AvsResultType = 'E'
	NonUsIssuerDoesNotParticipate            AvsResultType = 'G'
	NeitherAddressNorZipMatches              AvsResultType = 'N'
	RetrySystemUnavailableOrTimeout          AvsResultType = 'R'
	CardTypeNotSupported                     AvsResultType = 'S'
	AddressInformationUnavailable            AvsResultType = 'U'
	NineDigitZipMatchAddressDoesNot          AvsResultType = 'W'
	ExactMatchNineDigitZipAndAddress         AvsResultType = 'X'
	AddressAndFiveDigitZipMatch              AvsResultType = 'Y'
	FiveDigitZipMatchesAddressDoesNot        AvsResultType = 'Z'
	IntlStreetMatchOnlyPostalCodeNotVerified AvsResultType = 'B'
	IntlNoMatch                              AvsResultType = 'C'
	IntlStreetMatchOnly                      AvsResultType = 'D'
	IntlStreetNotVerified                    AvsResultType = 'I'
	IntlStreetAddrMatch                      AvsResultType = 'M'
	IntlPostalCodeMatchStreetNotVerified     AvsResultType = 'P'
)

func (x AvsResultType) String() string {
	switch x {
	case SpaceOrEmptyAvsResult:
		return "SpaceOrEmptyAvsResult"
	case CreditOrOffline:
		return "CreditOrOffline"
	case AddressMatchZipDoesNot:
		return "AddressMatchZipDoesNot"
	case IneligibleTransactionNotMoto:
		return "IneligibleTransactionNotMoto"
	case NonUsIssuerDoesNotParticipate:
		return "NonUsIssuerDoesNotParticipate"
	case NeitherAddressNorZipMatches:
		return "NeitherAddressNorZipMatches"
	case RetrySystemUnavailableOrTimeout:
		return "RetrySystemUnavailableOrTimeout"
	case CardTypeNotSupported:
		return "CardTypeNotSupported"
	case AddressInformationUnavailable:
		return "AddressInformationUnavailable"
	case NineDigitZipMatchAddressDoesNot:
		return "NineDigitZipMatchAddressDoesNot"
	case ExactMatchNineDigitZipAndAddress:
		return "ExactMatchNineDigitZipAndAddress"
	case AddressAndFiveDigitZipMatch:
		return "AddressAndFiveDigitZipMatch"
	case FiveDigitZipMatchesAddressDoesNot:
		return "FiveDigitZipMatchesAddressDoesNot"
	case IntlStreetMatchOnlyPostalCodeNotVerified:
		return "IntlStreetMatchOnlyPostalCodeNotVerified"
	case IntlNoMatch:
		return "IntlNoMatch"
	case IntlStreetMatchOnly:
		return "IntlStreetMatchOnly"
	case IntlStreetNotVerified:
		return "IntlStreetNotVerified"
	case IntlStreetAddrMatch:
		return "IntlStreetAddrMatch"
	case IntlPostalCodeMatchStreetNotVerified:
		return "IntlPostalCodeMatchStreetNotVerified"
	}
	return ""
}

//type OptionalDataGroupRecordType uint

// "D@@@@", "Base Group"
// "DP@@@", "Encrypted Detail Record"
// "D@@@B", "(Group 2) Detail Record (restaurant)"
// "D@@‘D", "(Groups 3 and 12) Detail Record (direct marketing)"
// "D@@@P",  "(Group 5) Detail Record (Hotel)"
// "D@@AP", "(Groups 5 and 7) Detail Record (Hotel/American Express)"
// "D@@@H", "(Group 4) Detail Record (auto rental)"
// "D@A@H", "(Groups 4 and 13) Detail Record (auto rental/MasterCard)"
// "D@@D@", "(Group 9) Detail Record (Passenger Transport)"
// "D@@H@", "(Group 10) Detail Record (Direct Debit)"
// "D@@H@", "(Group 10) Detail Record (EBT)"
// "D@@’D", "Detail Record (e-Commerce)"

// const (
// 	SpaceOrEmptyBatchRecord               OptionalDataGroupRecordType = 0
// 	BatchDetailRecordAutoRental           OptionalDataGroupRecordType = 1  //D@@@H (group 4) auto rental
// 	BatchDetailRecordAutoRentalMastercard OptionalDataGroupRecordType = 2  //D@A@H (groups 4 and 13) auto rental/master card
// 	BatchDetailRecordBaseGroup            OptionalDataGroupRecordType = 3  //D@@@@ base group
// 	BatchDetailRecordCcsPrivateLabel      OptionalDataGroupRecordType = 4  //DA@@@ (group 19) ccs private label
// 	BatchDetailRecordDirectDebit          OptionalDataGroupRecordType = 5  //D@@H@ (group 10) direct debit
// 	BatchDetailRecordDirectMarket         OptionalDataGroupRecordType = 6  //D@@`D (groups 3 and 12) direct marketing
// 	BatchDetailRecordEbt                  OptionalDataGroupRecordType = 7  //D@@H@ (group 10) ebt
// 	BatchDetailRecordEcom                 OptionalDataGroupRecordType = 8  //D@@'D ecom
// 	BatchDetailRecordEncrypted            OptionalDataGroupRecordType = 9  //DP@@@
// 	BatchDetailRecordHotel                OptionalDataGroupRecordType = 10 //D@@@P (group 5) hotel
// 	BatchDetailRecordHotelAmex            OptionalDataGroupRecordType = 11 //D@@AP (groups 5 and 7) hotel amex
// 	BatchDetailRecordPassengerTransport   OptionalDataGroupRecordType = 12 //D@@D@ (group 9) passenger transport
// 	BatchDetailRecordPosCheck             OptionalDataGroupRecordType = 13 //D@DH@ pos check service without cash back (optional groups 10 and 15)
// 	BatchDetailRecordPosCheckWCashback    OptionalDataGroupRecordType = 14 //D@DHA pos check service with cash back (optional groups 1 10 and 15)
// 	BatchDetailRecordQuasiCash            OptionalDataGroupRecordType = 15 //D@@@@ quasi cash
// 	BatchDetailRecordRestaurant           OptionalDataGroupRecordType = 16 //D@@@B (group 2) restaurant
// 	BatchDetailRecordWithToken            OptionalDataGroupRecordType = 17 //DP@@@
// 	BatchHeaderRecordBaseGroup            OptionalDataGroupRecordType = 18 //H@@@@ base group
// 	BatchHeaderRecordHotelAutoRental      OptionalDataGroupRecordType = 19 //H@@@B (group 2) hotel auto rental
// 	BatchHeaderRecordPassengerTransport   OptionalDataGroupRecordType = 20 //H@@@A (group 1) passenger transport
// 	BatchParameterRecordBaseGroup         OptionalDataGroupRecordType = 21 //P@@@@ base group
// 	BatchParameterRecordDirectMarket      OptionalDataGroupRecordType = 22 //P@@@@ direct marketing
// 	BatchResponseRecord                   OptionalDataGroupRecordType = 23 //R@@@@
// 	BatchTrailerRecord                    OptionalDataGroupRecordType = 24 //T@@@@
// )

// func (x OptionalDataGroupRecordType) String() string {
// 	switch x {
// 	case SpaceOrEmptyBatchRecord:
// 		return "SpaceOrEmptyBatchRecord"
// 	case BatchHeaderRecordBaseGroup:
// 		return "BatchHeaderRecordBaseGroup"
// 	case BatchHeaderRecordPassengerTransport:
// 		return "BatchHeaderRecordPassengerTransport"
// 	case BatchHeaderRecordHotelAutoRental:
// 		return "BatchHeaderRecordHotelAutoRental"
// 	case BatchParameterRecordBaseGroup:
// 		return "BatchParameterRecordBaseGroup"
// 	case BatchParameterRecordDirectMarket:
// 		return "BatchParameterRecordDirectMarket"
// 	case BatchDetailRecordBaseGroup:
// 		return "BatchDetailRecordBaseGroup"
// 	case BatchDetailRecordRestaurant:
// 		return "BatchDetailRecordRestaurant"
// 	case BatchDetailRecordDirectMarket:
// 		return "BatchDetailRecordDirectMarket"
// 	case BatchDetailRecordHotel:
// 		return "BatchDetailRecordHotel"
// 	case BatchDetailRecordHotelAmex:
// 		return "BatchDetailRecordHotelAmex"
// 	case BatchDetailRecordAutoRental:
// 		return "BatchDetailRecordAutoRental"
// 	case BatchDetailRecordAutoRentalMastercard:
// 		return "BatchDetailRecordAutoRentalMastercard"
// 	case BatchDetailRecordPassengerTransport:
// 		return "BatchDetailRecordPassengerTransport"
// 	case BatchDetailRecordDirectDebit:
// 		return "BatchDetailRecordDirectDebit"
// 	case BatchDetailRecordEbt:
// 		return "BatchDetailRecordEbt"
// 	case BatchDetailRecordEcom:
// 		return "BatchDetailRecordEcom"
// 	case BatchDetailRecordQuasiCash:
// 		return "BatchDetailRecordQuasiCash"
// 	case BatchDetailRecordCcsPrivateLabel:
// 		return "BatchDetailRecordCcsPrivateLabel"
// 	case BatchDetailRecordPosCheck:
// 		return "BatchDetailRecordPosCheck"
// 	case BatchDetailRecordPosCheckWCashback:
// 		return "BatchDetailRecordPosCheckWCashback"
// 	case BatchTrailerRecord:
// 		return "BatchTrailerRecord"
// 	case BatchResponseRecord:
// 		return "BatchResponseRecord"
// 	}
// 	return ""
// }

type BatchResponseCodeType string

const (
	GoodBatch      BatchResponseCodeType = "GB"
	DuplicateBatch BatchResponseCodeType = "QD"
	RejectedBatch  BatchResponseCodeType = "RB"
)

func (x BatchResponseCodeType) String() string {
	switch x {
	case GoodBatch:
		return "GoodBatch"
	case DuplicateBatch:
		return "DuplicateBatch"
	case RejectedBatch:
		return "RejectedBatch"
	}
	return ""
}

type BlockingIndicatorType rune

const (
	NotBlocked BlockingIndicatorType = '0'
	Blocked    BlockingIndicatorType = '2'
)

func (x BlockingIndicatorType) String() string {
	switch x {
	case NotBlocked:
		return "NotBlocked"
	case Blocked:
		return "Blocked"
	}
	return ""
}

type CardholderIdCodeType rune

const (
	SpaceOrEmptyCardholderId                                   CardholderIdCodeType = ' '
	PersonalIdentificationNumber32CharStaticKeyNonUsa          CardholderIdCodeType = 'A'
	PinAtAutomatedDispensingMachine32CharStaticKey             CardholderIdCodeType = 'B'
	SelfServiceLimitedAmountTerminalNoIdMethodAvailable        CardholderIdCodeType = 'C'
	SelfServiceTerminalNoIdMethodAvailable                     CardholderIdCodeType = 'D'
	CustomerActivatedAutomatedFuelDispenserNoIdMethodAvailable CardholderIdCodeType = 'E'
	PinAuthenticationByIccChipCard                             CardholderIdCodeType = 'F'
	ClearTextPin                                               CardholderIdCodeType = 'G'
	PinAtAutomatedDispensingMachine32CharDukpt                 CardholderIdCodeType = 'J'
	PersonalIdentificationNumber32CharDukpt                    CardholderIdCodeType = 'K'
	CardPresentUnableToReadMagStripeSendingAvsData             CardholderIdCodeType = 'M'
	CardNotPresentIncludesEcomAndFullAvsData                   CardholderIdCodeType = 'N'
	NoCvmRequiredEMVContactOrContactless                       CardholderIdCodeType = 'P'
	PersonalIdentificationNumber32CharStaticKey                CardholderIdCodeType = 'S'
	CardholderSignatureTerminalHasPinPad                       CardholderIdCodeType = 'Z'
	CardholderSignatureNoPinPadAvailable                       CardholderIdCodeType = '@'
)

func (x CardholderIdCodeType) String() string {
	switch x {
	case SpaceOrEmptyCardholderId:
		return "SpaceOrEmptyCardholderId"
	case PersonalIdentificationNumber32CharStaticKeyNonUsa:
		return "PersonalIdentificationNumber32CharStaticKeyNonUsa"
	case PinAtAutomatedDispensingMachine32CharStaticKey:
		return "PinAtAutomatedDispensingMachine32CharStaticKey"
	case SelfServiceLimitedAmountTerminalNoIdMethodAvailable:
		return "SelfServiceLimitedAmountTerminalNoIdMethodAvailable"
	case SelfServiceTerminalNoIdMethodAvailable:
		return "SelfServiceTerminalNoIdMethodAvailable"
	case CustomerActivatedAutomatedFuelDispenserNoIdMethodAvailable:
		return "CustomerActivatedAutomatedFuelDispenserNoIdMethodAvailable"
	case PinAuthenticationByIccChipCard:
		return "PinAuthenticationByIccChipCard"
	case ClearTextPin:
		return "ClearTextPin"
	case PinAtAutomatedDispensingMachine32CharDukpt:
		return "PinAtAutomatedDispensingMachine32CharDukpt"
	case PersonalIdentificationNumber32CharDukpt:
		return "PersonalIdentificationNumber32CharDukpt"
	case CardPresentUnableToReadMagStripeSendingAvsData:
		return "CardPresentUnableToReadMagStripeSendingAvsData"
	case CardNotPresentIncludesEcomAndFullAvsData:
		return "CardNotPresentIncludesEcomAndFullAvsData"
	case NoCvmRequiredEMVContactOrContactless:
		return "NoCvmRequiredEMVContactOrContactless"
	case PersonalIdentificationNumber32CharStaticKey:
		return "PersonalIdentificationNumber32CharStaticKey"
	case CardholderSignatureTerminalHasPinPad:
		return "CardholderSignatureTerminalHasPinPad"
	case CardholderSignatureNoPinPadAvailable:
		return "CardholderSignatureNoPinPadAvailable"
	}
	return ""
}

type CavvResultCodeType rune

const (
	BlankOrNotPresentCavvNotValidated                                     CavvResultCodeType = ' '
	CavvNotValidatedErroneousDataSubmitted                                CavvResultCodeType = '0'
	CavvFailedValidation                                                  CavvResultCodeType = '1'
	CavvPassedValidation                                                  CavvResultCodeType = '2'
	CavvValidationCouldNotBePerformedIssuerAttemptIncomplete              CavvResultCodeType = '3'
	CavvValidationCouldNotBePerformedIssuerSystemError                    CavvResultCodeType = '4'
	ReservedForFutureUse05                                                CavvResultCodeType = '5'
	ReservedForFutureUse06                                                CavvResultCodeType = '6'
	CavvAttemptFailedValidationIssuerAvailableUsIssuedCardNonUsAcquirer   CavvResultCodeType = '7'
	CavvAttemptPassedValidationIssuerAvailableUsIssuedCardNonUsAcquirer   CavvResultCodeType = '8'
	CavvAttemptFailedValidationIssuerUnavailableUsIssuedCardNonUsAcquirer CavvResultCodeType = '9'
	CavvAttemptPassedValidationIssuerUnavailableUsIssuedCardNonUsAcquirer CavvResultCodeType = 'A'
	CavvPassedValidationInformationOnlyNoLiabilityShift                   CavvResultCodeType = 'B'
)

func (x CavvResultCodeType) String() string {
	switch x {
	case BlankOrNotPresentCavvNotValidated:
		return "BlankOrNotPresentCavvNotValidated"
	case CavvNotValidatedErroneousDataSubmitted:
		return "CavvNotValidatedErroneousDataSubmitted"
	case CavvFailedValidation:
		return "CavvFailedValidation"
	case CavvPassedValidation:
		return "CavvPassedValidation"
	case CavvValidationCouldNotBePerformedIssuerAttemptIncomplete:
		return "CavvValidationCouldNotBePerformedIssuerAttemptIncomplete"
	case CavvValidationCouldNotBePerformedIssuerSystemError:
		return "CavvValidationCouldNotBePerformedIssuerSystemError"
	case ReservedForFutureUse05:
		return "ReservedForFutureUse05"
	case ReservedForFutureUse06:
		return "ReservedForFutureUse06"
	case CavvAttemptFailedValidationIssuerAvailableUsIssuedCardNonUsAcquirer:
		return "CavvAttemptFailedValidationIssuerAvailableUsIssuedCardNonUsAcquirer"
	case CavvAttemptPassedValidationIssuerAvailableUsIssuedCardNonUsAcquirer:
		return "CavvAttemptPassedValidationIssuerAvailableUsIssuedCardNonUsAcquirer"
	case CavvAttemptFailedValidationIssuerUnavailableUsIssuedCardNonUsAcquirer:
		return "CavvAttemptFailedValidationIssuerUnavailableUsIssuedCardNonUsAcquirer"
	case CavvAttemptPassedValidationIssuerUnavailableUsIssuedCardNonUsAcquirer:
		return "CavvAttemptPassedValidationIssuerUnavailableUsIssuedCardNonUsAcquirer"
	case CavvPassedValidationInformationOnlyNoLiabilityShift:
		return "CavvPassedValidationInformationOnlyNoLiabilityShift"
	}
	return ""
}

type CheckSettlementCodeType rune

const (
	VitalCaptureSystem CheckSettlementCodeType = '1'
	AchCapture         CheckSettlementCodeType = '2'
)

func (x CheckSettlementCodeType) String() string {
	switch x {
	case VitalCaptureSystem:
		return "VitalCaptureSystem"
	case AchCapture:
		return "AchCapture"
	}
	return ""
}

type CheckValidationStateCodeEquifaxType uint

const (
	AL_Alabama         CheckValidationStateCodeEquifaxType = 41
	MT_Montana         CheckValidationStateCodeEquifaxType = 68
	AK_Alaska          CheckValidationStateCodeEquifaxType = 42
	NE_Nebraska        CheckValidationStateCodeEquifaxType = 69
	AZ_Arizona         CheckValidationStateCodeEquifaxType = 43
	NV_Nevada          CheckValidationStateCodeEquifaxType = 70
	AR_Arkansas        CheckValidationStateCodeEquifaxType = 45
	NH_NewHampshire    CheckValidationStateCodeEquifaxType = 71
	CA_California      CheckValidationStateCodeEquifaxType = 46
	NJ_NewJersey       CheckValidationStateCodeEquifaxType = 66
	CO_Colorado        CheckValidationStateCodeEquifaxType = 47
	NM_NewMexico       CheckValidationStateCodeEquifaxType = 72
	CT_Connecticut     CheckValidationStateCodeEquifaxType = 48
	NY_NewYork         CheckValidationStateCodeEquifaxType = 55
	DE_Delaware        CheckValidationStateCodeEquifaxType = 77
	NC_North_Carolina  CheckValidationStateCodeEquifaxType = 73
	DC_DistofColumbia  CheckValidationStateCodeEquifaxType = 91
	ND_NorthDakota     CheckValidationStateCodeEquifaxType = 74
	FL_Florida         CheckValidationStateCodeEquifaxType = 49
	OH_Ohio            CheckValidationStateCodeEquifaxType = 75
	GA_Georgia         CheckValidationStateCodeEquifaxType = 50
	OK_Oklahoma        CheckValidationStateCodeEquifaxType = 76
	HI_Hawaii          CheckValidationStateCodeEquifaxType = 51
	OR_Oregon          CheckValidationStateCodeEquifaxType = 78
	ID_Idaho           CheckValidationStateCodeEquifaxType = 52
	PA_Pennsylvania    CheckValidationStateCodeEquifaxType = 44
	IL_Illinois        CheckValidationStateCodeEquifaxType = 53
	RI_RhodeIsland     CheckValidationStateCodeEquifaxType = 79
	IN_Indiana         CheckValidationStateCodeEquifaxType = 54
	SC_SouthCarolina   CheckValidationStateCodeEquifaxType = 80
	IA_Iowa            CheckValidationStateCodeEquifaxType = 56
	SD_SouthDakota     CheckValidationStateCodeEquifaxType = 81
	KS_Kansas          CheckValidationStateCodeEquifaxType = 57
	TN_Tennessee       CheckValidationStateCodeEquifaxType = 82
	KY_Kentucky        CheckValidationStateCodeEquifaxType = 58
	TX_Texas           CheckValidationStateCodeEquifaxType = 83
	LA_Louisiana       CheckValidationStateCodeEquifaxType = 59
	UT_Utah            CheckValidationStateCodeEquifaxType = 84
	ME_Maine           CheckValidationStateCodeEquifaxType = 60
	VT_Vermont         CheckValidationStateCodeEquifaxType = 85
	MD_Maryland        CheckValidationStateCodeEquifaxType = 61
	VA_Virginia        CheckValidationStateCodeEquifaxType = 86
	MA_Massachusetts   CheckValidationStateCodeEquifaxType = 62
	WA_Washington      CheckValidationStateCodeEquifaxType = 87
	MI_Michigan        CheckValidationStateCodeEquifaxType = 63
	WV_WestVirginia    CheckValidationStateCodeEquifaxType = 88
	MN_Minnesota       CheckValidationStateCodeEquifaxType = 64
	WI_Wisconsin       CheckValidationStateCodeEquifaxType = 89
	MS_Mississippi     CheckValidationStateCodeEquifaxType = 65
	WY_Wyoming         CheckValidationStateCodeEquifaxType = 90
	MO_Missouri        CheckValidationStateCodeEquifaxType = 67
	AB_Alberta         CheckValidationStateCodeEquifaxType = 28
	NS_NovaScotia      CheckValidationStateCodeEquifaxType = 34
	BC_BritishColumbia CheckValidationStateCodeEquifaxType = 29
	ON_Ontario         CheckValidationStateCodeEquifaxType = 35
	MB_Manitoba        CheckValidationStateCodeEquifaxType = 30
	PE_PrinceEdwardIs  CheckValidationStateCodeEquifaxType = 36
	NB_NewBrunswick    CheckValidationStateCodeEquifaxType = 31
	PQ_Quebec          CheckValidationStateCodeEquifaxType = 37
	NF_Newfoundland    CheckValidationStateCodeEquifaxType = 32
	SK_Saskatchewan    CheckValidationStateCodeEquifaxType = 38
	NT_NWTerritories   CheckValidationStateCodeEquifaxType = 33
	YT_YukonTerritory  CheckValidationStateCodeEquifaxType = 39
)

func (x CheckValidationStateCodeEquifaxType) String() string {
	switch x {
	case AL_Alabama:
		return "AL_Alabama"
	case MT_Montana:
		return "MT_Montana"
	case AK_Alaska:
		return "AK_Alaska"
	case NE_Nebraska:
		return "NE_Nebraska"
	case AZ_Arizona:
		return "AZ_Arizona"
	case NV_Nevada:
		return "NV_Nevada"
	case AR_Arkansas:
		return "AR_Arkansas"
	case NH_NewHampshire:
		return "NH_NewHampshire"
	case CA_California:
		return "CA_California"
	case NJ_NewJersey:
		return "NJ_NewJersey"
	case CO_Colorado:
		return "CO_Colorado"
	case NM_NewMexico:
		return "NM_NewMexico"
	case CT_Connecticut:
		return "CT_Connecticut"
	case NY_NewYork:
		return "NY_NewYork"
	case DE_Delaware:
		return "DE_Delaware"
	case NC_North_Carolina:
		return "NC_North_Carolina"
	case DC_DistofColumbia:
		return "DC_DistofColumbia"
	case ND_NorthDakota:
		return "ND_NorthDakota"
	case FL_Florida:
		return "FL_Florida"
	case OH_Ohio:
		return "OH_Ohio"
	case GA_Georgia:
		return "GA_Georgia"
	case OK_Oklahoma:
		return "OK_Oklahoma"
	case HI_Hawaii:
		return "HI_Hawaii"
	case OR_Oregon:
		return "OR_Oregon"
	case ID_Idaho:
		return "ID_Idaho"
	case PA_Pennsylvania:
		return "PA_Pennsylvania"
	case IL_Illinois:
		return "IL_Illinois"
	case RI_RhodeIsland:
		return "RI_RhodeIsland"
	case IN_Indiana:
		return "IN_Indiana"
	case SC_SouthCarolina:
		return "SC_SouthCarolina"
	case IA_Iowa:
		return "IA_Iowa"
	case SD_SouthDakota:
		return "SD_SouthDakota"
	case KS_Kansas:
		return "KS_Kansas"
	case TN_Tennessee:
		return "TN_Tennessee"
	case KY_Kentucky:
		return "KY_Kentucky"
	case TX_Texas:
		return "TX_Texas"
	case LA_Louisiana:
		return "LA_Louisiana"
	case UT_Utah:
		return "UT_Utah"
	case ME_Maine:
		return "ME_Maine"
	case VT_Vermont:
		return "VT_Vermont"
	case MD_Maryland:
		return "MD_Maryland"
	case VA_Virginia:
		return "VA_Virginia"
	case MA_Massachusetts:
		return "MA_Massachusetts"
	case WA_Washington:
		return "WA_Washington"
	case MI_Michigan:
		return "MI_Michigan"
	case WV_WestVirginia:
		return "WV_WestVirginia"
	case MN_Minnesota:
		return "MN_Minnesota"
	case WI_Wisconsin:
		return "WI_Wisconsin"
	case MS_Mississippi:
		return "MS_Mississippi"
	case WY_Wyoming:
		return "WY_Wyoming"
	case MO_Missouri:
		return "MO_Missouri"
	case AB_Alberta:
		return "AB_Alberta"
	case NS_NovaScotia:
		return "NS_NovaScotia"
	case BC_BritishColumbia:
		return "BC_BritishColumbia"
	case ON_Ontario:
		return "ON_Ontario"
	case MB_Manitoba:
		return "MB_Manitoba"
	case PE_PrinceEdwardIs:
		return "PE_PrinceEdwardIs"
	case NB_NewBrunswick:
		return "NB_NewBrunswick"
	case PQ_Quebec:
		return "PQ_Quebec"
	case NF_Newfoundland:
		return "NF_Newfoundland"
	case SK_Saskatchewan:
		return "SK_Saskatchewan"
	case NT_NWTerritories:
		return "NT_NWTerritories"
	case YT_YukonTerritory:
		return "YT_YukonTerritory"
	}
	return ""
}

type CheckValidationStateCodeTelecheckType uint

const (
	AL_AlabamaTelecheck        CheckValidationStateCodeTelecheckType = 25
	MT_MontanaTelecheck        CheckValidationStateCodeTelecheckType = 68
	LK_AlaskaTelecheck         CheckValidationStateCodeTelecheckType = 55
	NE_NebraskaTelecheck       CheckValidationStateCodeTelecheckType = 63
	AZ_ArizonaTelecheck        CheckValidationStateCodeTelecheckType = 20
	EV_NevadaTelecheck         CheckValidationStateCodeTelecheckType = 38
	AR_ArkansasTelecheck       CheckValidationStateCodeTelecheckType = 27
	HP_NewHampshireTelecheck   CheckValidationStateCodeTelecheckType = 47
	CF_CaliforniaTelecheck     CheckValidationStateCodeTelecheckType = 23
	JE_NewJerseyTelecheck      CheckValidationStateCodeTelecheckType = 53
	CO_ColoradoTelecheck       CheckValidationStateCodeTelecheckType = 26
	EX_NewMexicoTelecheck      CheckValidationStateCodeTelecheckType = 39
	CT_ConnecticutTelecheck    CheckValidationStateCodeTelecheckType = 28
	NY_CTNewYorkTelecheck      CheckValidationStateCodeTelecheckType = 69
	DE_DelawareTelecheck       CheckValidationStateCodeTelecheckType = 33
	RL_NorthCarolinaTelecheck  CheckValidationStateCodeTelecheckType = 75
	WD_DistofColumbiaTelecheck CheckValidationStateCodeTelecheckType = 93
	DN_NorthDakotaTelecheck    CheckValidationStateCodeTelecheckType = 36
	FL_FloridaTelecheck        CheckValidationStateCodeTelecheckType = 35
	OZ_OhioTelecheck           CheckValidationStateCodeTelecheckType = 60
	GA_GeorgiaTelecheck        CheckValidationStateCodeTelecheckType = 42
	OK_OklahomaTelecheck       CheckValidationStateCodeTelecheckType = 65
	HI_HawaiiTelecheck         CheckValidationStateCodeTelecheckType = 44
	OR_OregonTelecheck         CheckValidationStateCodeTelecheckType = 67
	ID_IdahoTelecheck          CheckValidationStateCodeTelecheckType = 43
	PV_PennsylvaniaTelecheck   CheckValidationStateCodeTelecheckType = 78
	IL_IllinoisTelecheck       CheckValidationStateCodeTelecheckType = 45
	//TelecheckRhodeIsland79RIRI
	IN_IndianaTelecheck         CheckValidationStateCodeTelecheckType = 46
	SC_SouthCarolinaTelecheck   CheckValidationStateCodeTelecheckType = 75
	IW_IowaTelecheck            CheckValidationStateCodeTelecheckType = 49
	SD_SouthDakotaTelecheck     CheckValidationStateCodeTelecheckType = 73
	KS_KansasTelecheck          CheckValidationStateCodeTelecheckType = 57
	TN_TennesseeTelecheck       CheckValidationStateCodeTelecheckType = 86
	KY_KentuckyTelecheck        CheckValidationStateCodeTelecheckType = 59
	TX_TexasTelecheck           CheckValidationStateCodeTelecheckType = 89
	LA_LouisianaTelecheck       CheckValidationStateCodeTelecheckType = 52
	UT_UtahTelecheck            CheckValidationStateCodeTelecheckType = 88
	LO_MaineTelecheck           CheckValidationStateCodeTelecheckType = 56
	VE_VermontTelecheck         CheckValidationStateCodeTelecheckType = 83
	RY_MarylandTelecheck        CheckValidationStateCodeTelecheckType = 79
	VA_VirginiaTelecheck        CheckValidationStateCodeTelecheckType = 82
	RY_MassachusettsTelecheck   CheckValidationStateCodeTelecheckType = 87
	WA_WashingtonTelecheck      CheckValidationStateCodeTelecheckType = 92
	GZ_MichiganTelecheck        CheckValidationStateCodeTelecheckType = 40
	WV_WestVirginiaTelecheck    CheckValidationStateCodeTelecheckType = 98
	MI_MinnesotaTelecheck       CheckValidationStateCodeTelecheckType = 64
	WI_WisconsinTelecheck       CheckValidationStateCodeTelecheckType = 94
	SS_MississippiTelecheck     CheckValidationStateCodeTelecheckType = 77
	WY_WyomingTelecheck         CheckValidationStateCodeTelecheckType = 99
	MO_MissouriTelecheck        CheckValidationStateCodeTelecheckType = 66
	AB_AlbertaTelecheck         CheckValidationStateCodeTelecheckType = 21
	NS_NovaScotiaTelecheck      CheckValidationStateCodeTelecheckType = 41
	BC_BritishColumbiaTelecheck CheckValidationStateCodeTelecheckType = 54
	ON_OntarioTelecheck         CheckValidationStateCodeTelecheckType = 51
	MB_ManitobaTelecheck        CheckValidationStateCodeTelecheckType = 61
	PE_PrinceEdwardIsTelecheck  CheckValidationStateCodeTelecheckType = 81
	NB_NewBrunswickTelecheck    CheckValidationStateCodeTelecheckType = 61
	PQ_QuebecTelecheck          CheckValidationStateCodeTelecheckType = 71
	NF_NewfoundlandTelecheck    CheckValidationStateCodeTelecheckType = 31
	SK_SaskatchewanTelecheck    CheckValidationStateCodeTelecheckType = 58
	NW_NWTerritoriesTelecheck   CheckValidationStateCodeTelecheckType = 37
	YU_YukonTerritoryTelecheck  CheckValidationStateCodeTelecheckType = 91
)

func (x CheckValidationStateCodeTelecheckType) String() string {
	switch x {
	case AL_AlabamaTelecheck:
		return "AL_AlabamaTelecheck"
	case MT_MontanaTelecheck:
		return "MT_MontanaTelecheck"
	case LK_AlaskaTelecheck:
		return "LK_AlaskaTelecheck"
	case NE_NebraskaTelecheck:
		return "NE_NebraskaTelecheck"
	case AZ_ArizonaTelecheck:
		return "AZ_ArizonaTelecheck"
	case EV_NevadaTelecheck:
		return "EV_NevadaTelecheck"
	case AR_ArkansasTelecheck:
		return "AR_ArkansasTelecheck"
	case HP_NewHampshireTelecheck:
		return "HP_NewHampshireTelecheck"
	case CF_CaliforniaTelecheck:
		return "CF_CaliforniaTelecheck"
	case JE_NewJerseyTelecheck:
		return "JE_NewJerseyTelecheck"
	case CO_ColoradoTelecheck:
		return "CO_ColoradoTelecheck"
	case EX_NewMexicoTelecheck:
		return "EX_NewMexicoTelecheck"
	case CT_ConnecticutTelecheck:
		return "CT_ConnecticutTelecheck"
	case NY_CTNewYorkTelecheck:
		return "NY_CTNewYorkTelecheck"
	case DE_DelawareTelecheck:
		return "DE_DelawareTelecheck"
	case RL_NorthCarolinaTelecheck:
		return "RL_NorthCarolinaTelecheck or SC_SouthCarolinaTelecheck"
	case WD_DistofColumbiaTelecheck:
		return "WD_DistofColumbiaTelecheck"
	case DN_NorthDakotaTelecheck:
		return "DN_NorthDakotaTelecheck"
	case FL_FloridaTelecheck:
		return "FL_FloridaTelecheck"
	case OZ_OhioTelecheck:
		return "OZ_OhioTelecheck"
	case GA_GeorgiaTelecheck:
		return "GA_GeorgiaTelecheck"
	case OK_OklahomaTelecheck:
		return "OK_OklahomaTelecheck"
	case HI_HawaiiTelecheck:
		return "HI_HawaiiTelecheck"
	case OR_OregonTelecheck:
		return "OR_OregonTelecheck"
	case ID_IdahoTelecheck:
		return "ID_IdahoTelecheck"
	case PV_PennsylvaniaTelecheck:
		return "PV_PennsylvaniaTelecheck"
	case IL_IllinoisTelecheck:
		return "IL_IllinoisTelecheck"
	case IN_IndianaTelecheck:
		return "IN_IndianaTelecheck"
	// case SC_SouthCarolinaTelecheck:
	// 	return "SC_SouthCarolinaTelecheck"
	case IW_IowaTelecheck:
		return "IW_IowaTelecheck"
	case SD_SouthDakotaTelecheck:
		return "SD_SouthDakotaTelecheck"
	case KS_KansasTelecheck:
		return "KS_KansasTelecheck"
	case TN_TennesseeTelecheck:
		return "TN_TennesseeTelecheck"
	case KY_KentuckyTelecheck:
		return "KY_KentuckyTelecheck"
	case TX_TexasTelecheck:
		return "TX_TexasTelecheck"
	case LA_LouisianaTelecheck:
		return "LA_LouisianaTelecheck"
	case UT_UtahTelecheck:
		return "UT_UtahTelecheck"
	case LO_MaineTelecheck:
		return "LO_MaineTelecheck"
	case VE_VermontTelecheck:
		return "VE_VermontTelecheck"
	case RY_MarylandTelecheck:
		return "RY_MarylandTelecheck"
	case VA_VirginiaTelecheck:
		return "VA_VirginiaTelecheck"
	case RY_MassachusettsTelecheck:
		return "RY_MassachusettsTelecheck"
	case WA_WashingtonTelecheck:
		return "WA_WashingtonTelecheck"
	case GZ_MichiganTelecheck:
		return "GZ_MichiganTelecheck"
	case WV_WestVirginiaTelecheck:
		return "WV_WestVirginiaTelecheck"
	case MI_MinnesotaTelecheck:
		return "MI_MinnesotaTelecheck"
	case WI_WisconsinTelecheck:
		return "WI_WisconsinTelecheck"
	case SS_MississippiTelecheck:
		return "SS_MississippiTelecheck"
	case WY_WyomingTelecheck:
		return "WY_WyomingTelecheck"
	case MO_MissouriTelecheck:
		return "MO_MissouriTelecheck"
	case AB_AlbertaTelecheck:
		return "AB_AlbertaTelecheck"
	case NS_NovaScotiaTelecheck:
		return "NS_NovaScotiaTelecheck"
	case BC_BritishColumbiaTelecheck:
		return "BC_BritishColumbiaTelecheck"
	case ON_OntarioTelecheck:
		return "ON_OntarioTelecheck"
	case MB_ManitobaTelecheck:
		return "MB_ManitobaTelecheck or NB_NewBrunswickTelecheck"
	case PE_PrinceEdwardIsTelecheck:
		return "PE_PrinceEdwardIsTelecheck"
	// case NB_NewBrunswickTelecheck:
	// 	return "NB_NewBrunswickTelecheck"
	case PQ_QuebecTelecheck:
		return "PQ_QuebecTelecheck"
	case NF_NewfoundlandTelecheck:
		return "NF_NewfoundlandTelecheck"
	case SK_SaskatchewanTelecheck:
		return "SK_SaskatchewanTelecheck"
	case NW_NWTerritoriesTelecheck:
		return "NW_NWTerritoriesTelecheck"
	case YU_YukonTerritoryTelecheck:
		return "YU_YukonTerritoryTelecheck"
	}
	return ""
}

type ChipConditionCodeType rune

const (
	SpaceOrEmptyChipConditionCode           ChipConditionCodeType = ' '
	ServiceCodeDoesNotBeginWith2Or6         ChipConditionCodeType = '0'
	LastCcpsReadSuccessOrNotChipTransaction ChipConditionCodeType = '1'
	UnsuccessfulChipRead                    ChipConditionCodeType = '2'
)

func (x ChipConditionCodeType) String() string {
	switch x {
	case SpaceOrEmptyChipConditionCode:
		return "SpaceOrEmptyChipConditionCode"
	case ServiceCodeDoesNotBeginWith2Or6:
		return "ServiceCodeDoesNotBeginWith2Or6"
	case LastCcpsReadSuccessOrNotChipTransaction:
		return "LastCcpsReadSuccessOrNotChipTransaction"
	case UnsuccessfulChipRead:
		return "UnsuccessfulChipRead"
	}
	return ""
}

type ChipCardTransactionType string

const (
	ChipCardGoodsOrServicePurchase        ChipCardTransactionType = "00"
	ChipCardWithdrawalCashAdvance         ChipCardTransactionType = "01"
	ChipCardAdjustmentDebit               ChipCardTransactionType = "02"
	ChipCardCheckGuaranteeFundsGuaranteed ChipCardTransactionType = "03"
	ChipCardQuasiCashTransaction          ChipCardTransactionType = "11"
	ChipCardScript                        ChipCardTransactionType = "17"
	ChipCardFeeCollection                 ChipCardTransactionType = "19"
	ChipCardReturnOfGoods                 ChipCardTransactionType = "20"
	ChipCardAdjustmentCredit              ChipCardTransactionType = "22"
	ChipCardFundsDisbursement             ChipCardTransactionType = "29"
	ChipCardAvailableFundsInquiry         ChipCardTransactionType = "30"
	ChipCardCardholderAccountTransfer     ChipCardTransactionType = "40"
	ChipCardCOPACGoodsOrServicePurchaser  ChipCardTransactionType = "A0"
	ChipCardCOPACWithdrawalCashAdvance    ChipCardTransactionType = "A1"
)

func (x ChipCardTransactionType) String() string {
	switch x {
	case ChipCardGoodsOrServicePurchase:
		return "ChipCardGoodsOrServicePurchase"
	case ChipCardWithdrawalCashAdvance:
		return "ChipCardWithdrawalCashAdvance"
	case ChipCardAdjustmentDebit:
		return "ChipCardAdjustmentDebit"
	case ChipCardCheckGuaranteeFundsGuaranteed:
		return "ChipCardCheckGuaranteeFundsGuaranteed"
	case ChipCardQuasiCashTransaction:
		return "ChipCardQuasiCashTransaction"
	case ChipCardScript:
		return "ChipCardScript"
	case ChipCardFeeCollection:
		return "ChipCardFeeCollection"
	case ChipCardReturnOfGoods:
		return "ChipCardReturnOfGoods"
	case ChipCardAdjustmentCredit:
		return "ChipCardAdjustmentCredit"
	case ChipCardFundsDisbursement:
		return "ChipCardFundsDisbursement"
	case ChipCardAvailableFundsInquiry:
		return "ChipCardAvailableFundsInquiry"
	case ChipCardCardholderAccountTransfer:
		return "ChipCardCardholderAccountTransfer"
	case ChipCardCOPACGoodsOrServicePurchaser:
		return "ChipCardCOPACGoodsOrServicePurchaser"
	case ChipCardCOPACWithdrawalCashAdvance:
		return "ChipCardCOPACWithdrawalCashAdvance"
	}
	return ""
}

type ComputerReservationSystemType string

const (
	EmptyComputerReservation ComputerReservationSystemType = "    "
	Start                    ComputerReservationSystemType = "STRT"
	TWA                      ComputerReservationSystemType = "PARS"
	DeltaSystem              ComputerReservationSystemType = "DATS"
	Sabre                    ComputerReservationSystemType = "SABR"
	CoviaApollo              ComputerReservationSystemType = "DALA"
	DrBlank                  ComputerReservationSystemType = "BLAN"
	DER                      ComputerReservationSystemType = "DERD"
	TUI                      ComputerReservationSystemType = "TUID"
)

func (x ComputerReservationSystemType) String() string {
	switch x {
	case EmptyComputerReservation:
		return "EmptyComputerReservation"
	case Start:
		return "Start"
	case TWA:
		return "TWA"
	case DeltaSystem:
		return "DeltaSystem"
	case Sabre:
		return "Sabre"
	case CoviaApollo:
		return "CoviaApollo"
	case DrBlank:
		return "DrBlank"
	case DER:
		return "DER"
	case TUI:
		return "TUI"
	}
	return ""
}

type CountryCodeType uint32

const (
	Antarctica CountryCodeType = 0
	Albania    CountryCodeType = 8
	Algeria    CountryCodeType = 12
	Azerbaijan CountryCodeType = 31
	Argentina  CountryCodeType = 32
	Australia  CountryCodeType = 36
	//ChristmasIsland CountryCodeType = 36
	//CocosKeelingIslands CountryCodeType = 36
	//HeardIslandAndMcdonaldIslands CountryCodeType = 36
	//Kiribati CountryCodeType = 36
	//Nauru CountryCodeType = 36
	//NorfolkIsland CountryCodeType = 36
	//Tuvalu CountryCodeType = 36
	Bahamas          CountryCodeType = 44
	Bahrain          CountryCodeType = 48
	Bangladesh       CountryCodeType = 50
	Armenia          CountryCodeType = 51
	Barbados         CountryCodeType = 52
	Bermuda          CountryCodeType = 60
	Bhutan           CountryCodeType = 64
	Bolivia          CountryCodeType = 68
	Botswana         CountryCodeType = 72
	Belize           CountryCodeType = 84
	SolomonIslands   CountryCodeType = 90
	BruneiDarussalam CountryCodeType = 96
	Myanmar          CountryCodeType = 104
	Burundi          CountryCodeType = 108
	Cambodia         CountryCodeType = 116
	Canada           CountryCodeType = 124
	CapeVerde        CountryCodeType = 132
	CaymanIslands    CountryCodeType = 136
	SriLanka         CountryCodeType = 144
	Chile            CountryCodeType = 152
	China            CountryCodeType = 156
	Colombia         CountryCodeType = 170
	Comoros          CountryCodeType = 174
	CostaRica        CountryCodeType = 188
	Croatia          CountryCodeType = 191
	Cuba             CountryCodeType = 192
	Cyprus           CountryCodeType = 196
	CzechRepublic    CountryCodeType = 203
	Denmark          CountryCodeType = 208
	//FaroeIslands CountryCodeType = 208
	//Greenland CountryCodeType = 208
	DominicanRepublic       CountryCodeType = 214
	ElSalvador              CountryCodeType = 222
	Ethiopia                CountryCodeType = 230
	Eritrea                 CountryCodeType = 232
	Estonia                 CountryCodeType = 233
	FalklandIslandsMalvinas CountryCodeType = 238
	Fiji                    CountryCodeType = 242
	Djibouti                CountryCodeType = 262
	Gambia                  CountryCodeType = 270
	Ghana                   CountryCodeType = 288
	Gibraltar               CountryCodeType = 292
	Guatemala               CountryCodeType = 320
	Guinea                  CountryCodeType = 324
	Guyana                  CountryCodeType = 328
	Haiti                   CountryCodeType = 332
	Honduras                CountryCodeType = 340
	HongKong                CountryCodeType = 344
	Hungary                 CountryCodeType = 348
	Iceland                 CountryCodeType = 352
	//Bhutan CountryCodeType = 356
	India                            CountryCodeType = 356
	Indonesia                        CountryCodeType = 360
	IranIslamicRepublicOf            CountryCodeType = 364
	Iraq                             CountryCodeType = 368
	Israel                           CountryCodeType = 376
	Jamaica                          CountryCodeType = 388
	Japan                            CountryCodeType = 392
	Kazakhstan                       CountryCodeType = 398
	Jordan                           CountryCodeType = 400
	Kenya                            CountryCodeType = 404
	KoreaDemocraticPeoplesRepublicOf CountryCodeType = 408
	KoreaRepublicOf                  CountryCodeType = 410
	Kuwait                           CountryCodeType = 414
	Kyrgyzstan                       CountryCodeType = 417
	LaoPeoplesDemocraticRepublic     CountryCodeType = 418
	Lebanon                          CountryCodeType = 422
	Lesotho                          CountryCodeType = 426
	Latvia                           CountryCodeType = 428
	Liberia                          CountryCodeType = 430
	LibyanArabJamahiriya             CountryCodeType = 434
	Lithuania                        CountryCodeType = 440
	Macao                            CountryCodeType = 446
	Madagascar                       CountryCodeType = 450
	Malawi                           CountryCodeType = 454
	Malaysia                         CountryCodeType = 458
	Maldives                         CountryCodeType = 462
	Malta                            CountryCodeType = 470
	Mauritania                       CountryCodeType = 478
	Mauritius                        CountryCodeType = 480
	Mexico                           CountryCodeType = 484
	Mongolia                         CountryCodeType = 496
	MoldovaRepublicOf                CountryCodeType = 498
	Morocco                          CountryCodeType = 504
	WesternSahara                    CountryCodeType = 504
	Mozambique                       CountryCodeType = 508
	Oman                             CountryCodeType = 512
	Namibia                          CountryCodeType = 516
	Nepal                            CountryCodeType = 524
	NetherlandsAntilles              CountryCodeType = 532
	Aruba                            CountryCodeType = 533
	Vanuatu                          CountryCodeType = 548
	//CookIslands CountryCodeType = 554
	NewZealand CountryCodeType = 554
	//Niue CountryCodeType = 554
	//Pitcairn CountryCodeType = 554
	//Tokelau CountryCodeType = 554
	Nicaragua CountryCodeType = 558
	Nigeria   CountryCodeType = 566
	//BouvetIsland CountryCodeType = 578
	Norway CountryCodeType = 578
	//SvalbardAndJanMayen CountryCodeType = 578
	Pakistan           CountryCodeType = 586
	Panama             CountryCodeType = 590
	PapuaNewGuinea     CountryCodeType = 598
	Paraguay           CountryCodeType = 600
	Peru               CountryCodeType = 604
	Philippines        CountryCodeType = 608
	GuineaBissau       CountryCodeType = 624
	Qatar              CountryCodeType = 634
	Romania            CountryCodeType = 642
	RussianFederation  CountryCodeType = 643
	Rwanda             CountryCodeType = 646
	SaintHelena        CountryCodeType = 654
	SaoTomeAndPrincipe CountryCodeType = 678
	SaudiArabia        CountryCodeType = 682
	Seychelles         CountryCodeType = 690
	SierraLeone        CountryCodeType = 694
	Singapore          CountryCodeType = 702
	Slovakia           CountryCodeType = 703
	VietNam            CountryCodeType = 704
	Slovenia           CountryCodeType = 705
	Somalia            CountryCodeType = 706
	//Lesotho CountryCodeType = 710
	//Namibia CountryCodeType = 710
	SouthAfrica CountryCodeType = 710
	Zimbabwe    CountryCodeType = 716
	Sudan       CountryCodeType = 736
	Swaziland   CountryCodeType = 748
	Sweden      CountryCodeType = 752
	//Liechtenstein CountryCodeType = 756
	Switzerland                          CountryCodeType = 756
	SyrianArabRepublic                   CountryCodeType = 760
	Thailand                             CountryCodeType = 764
	Tonga                                CountryCodeType = 776
	TrinidadAndTobago                    CountryCodeType = 780
	UnitedArabEmirates                   CountryCodeType = 784
	Tunisia                              CountryCodeType = 788
	Turkey                               CountryCodeType = 792
	Turkmenistan                         CountryCodeType = 795
	Uganda                               CountryCodeType = 800
	MacedoniaTheFormerYugoslavRepublicOf CountryCodeType = 807
	//RussianFederation CountryCodeType = 810
	Egypt                    CountryCodeType = 818
	UnitedKingdom            CountryCodeType = 826
	TanzaniaUnitedRepublicOf CountryCodeType = 834
	//AmericanSamoa CountryCodeType = 840
	//BritishIndianOceanTerritory CountryCodeType = 840
	//Ecuador CountryCodeType = 840
	//ElSalvador CountryCodeType = 840
	//Guam CountryCodeType = 840
	//Haiti840 CountryCodeType = 840
	//MarshallIslands CountryCodeType = 840
	//MicronesiaFederatedStatesOf CountryCodeType = 840
	//NorthernMarianaIslands CountryCodeType = 840
	//Palau CountryCodeType = 840
	//Panama CountryCodeType = 840
	//PuertoRico CountryCodeType = 840
	//TimorLeste CountryCodeType = 840
	//TurksAndCaicosIslands CountryCodeType = 840
	UnitedStates CountryCodeType = 840
	//UnitedStatesMinorOutlyingIslands CountryCodeType = 840
	//VirginIslandsBritish CountryCodeType = 840
	//VirginIslandsUs CountryCodeType = 840
	Uruguay               CountryCodeType = 858
	Uzbekistan            CountryCodeType = 860
	Venezuela             CountryCodeType = 862
	Samoa                 CountryCodeType = 882
	Yemen                 CountryCodeType = 886
	SerbiaMontenegro      CountryCodeType = 891
	Zambia                CountryCodeType = 894
	TaiwanProvinceOfChina CountryCodeType = 901
	CentralAfrica         CountryCodeType = 950
	//Cameroon CountryCodeType = 950
	//CentralAfricanRepublic CountryCodeType = 950
	//Chad CountryCodeType = 950
	//Congo CountryCodeType = 950
	//EquatorialGuinea CountryCodeType = 950
	//Gabon CountryCodeType = 950
	CaribbeanIslands CountryCodeType = 951
	//Anguilla CountryCodeType = 951
	//AntiguaAndBarbuda CountryCodeType = 951
	//Dominica CountryCodeType = 951
	//Grenada CountryCodeType = 951
	//Montserrat CountryCodeType = 951
	//SaintKittsAndNevis CountryCodeType = 951
	//SaintLucia CountryCodeType = 951
	//SaintVincentAndTheGrenadines CountryCodeType = 951
	//Benin CountryCodeType = 952
	//BurkinaFaso CountryCodeType = 952
	//CoteDivoire CountryCodeType = 952
	//GuineaBissau CountryCodeType = 952
	//Mali CountryCodeType = 952
	//Niger CountryCodeType = 952
	//Senegal CountryCodeType = 952
	//Togo CountryCodeType = 952
	FrenchPolynesia CountryCodeType = 953
	//NewCaledonia CountryCodeType = 953
	//WallisAndFutuna CountryCodeType = 953
	Suriname CountryCodeType = 968
	//Madagascar CountryCodeType = 969
	//Colombia CountryCodeType = 970
	Afghanistan                  CountryCodeType = 971
	Tajikistan                   CountryCodeType = 972
	Angola                       CountryCodeType = 973
	Belarus                      CountryCodeType = 974
	Bulgaria                     CountryCodeType = 975
	CongoTheDemocraticRepublicOf CountryCodeType = 976
	BosniaAndHerzegovina         CountryCodeType = 977
	//Andorra CountryCodeType = 978
	//Austria CountryCodeType = 978
	//Belgium CountryCodeType = 978
	//Finland CountryCodeType = 978
	//France CountryCodeType = 978
	//FrenchGuiana CountryCodeType = 978
	//FrenchSouthernTerritories CountryCodeType = 978
	//Germany CountryCodeType = 978
	//Greece CountryCodeType = 978
	//Guadeloupe CountryCodeType = 978
	//HolySeeVaticanCityState CountryCodeType = 978
	//Ireland CountryCodeType = 978
	//Italy CountryCodeType = 978
	//Luxembourg CountryCodeType = 978
	//Martinique CountryCodeType = 978
	//Mayotte CountryCodeType = 978
	//Monaco CountryCodeType = 978
	//Netherlands CountryCodeType = 978
	//Portugal CountryCodeType = 978
	//Reunion CountryCodeType = 978
	//SaintPierreAndMiquelon CountryCodeType = 978
	//SanMarino CountryCodeType = 978
	//SerbiaMontenegro CountryCodeType = 978
	//Spain CountryCodeType = 978
	WesternEurope CountryCodeType = 978
	//Mexico CountryCodeType = 979
	Ukraine CountryCodeType = 980
	Georgia CountryCodeType = 981
	//Bolivia CountryCodeType = 984
	Poland CountryCodeType = 985
	Brazil CountryCodeType = 986
	//Chile CountryCodeType = 990
	UnitedStatesNd CountryCodeType = 997
	UnitedStatesSd CountryCodeType = 998
)

func (x CountryCodeType) String() string {
	switch x {
	case Antarctica:
		return "Antarctica"
	case Albania:
		return "Albania"
	case Algeria:
		return "Algeria"
	case Azerbaijan:
		return "Azerbaijan"
	case Argentina:
		return "Argentina"
	case Australia:
		return "Australia"
	//ChristmasIsland CountryCodeType = 36
	//CocosKeelingIslands CountryCodeType = 36
	//HeardIslandAndMcdonaldIslands CountryCodeType = 36
	//Kiribati CountryCodeType = 36
	//Nauru CountryCodeType = 36
	//NorfolkIsland CountryCodeType = 36
	//Tuvalu CountryCodeType = 36
	case Bahamas:
		return "Bahamas"
	case Bahrain:
		return "Bahrain"
	case Bangladesh:
		return "Bangladesh"
	case Armenia:
		return "Armenia"
	case Barbados:
		return "Barbados"
	case Bermuda:
		return "Bermuda"
	case Bhutan:
		return "Bhutan"
	case Bolivia:
		return "Bolivia"
	case Botswana:
		return "Botswana"
	case Belize:
		return "Belize"
	case SolomonIslands:
		return "SolomonIslands"
	case BruneiDarussalam:
		return "BruneiDarussalam"
	case Myanmar:
		return "Myanmar"
	case Burundi:
		return "Burundi"
	case Cambodia:
		return "Cambodia"
	case Canada:
		return "Canada"
	case CapeVerde:
		return "CapeVerde"
	case CaymanIslands:
		return "CaymanIslands"
	case SriLanka:
		return "SriLanka"
	case Chile:
		return "Chile"
	case China:
		return "China"
	case Colombia:
		return "Colombia"
	case Comoros:
		return "Comoros"
	case CostaRica:
		return "CostaRica"
	case Croatia:
		return "Croatia"
	case Cuba:
		return "Cuba"
	case Cyprus:
		return "Cyprus"
	case CzechRepublic:
		return "CzechRepublic"
	case Denmark:
		return "Denmark"
	//FaroeIslands CountryCodeType = 208
	//Greenland CountryCodeType = 208
	case DominicanRepublic:
		return "DominicanRepublic"
	case ElSalvador:
		return "ElSalvador"
	case Ethiopia:
		return "Ethiopia"
	case Eritrea:
		return "Eritrea"
	case Estonia:
		return "Estonia"
	case FalklandIslandsMalvinas:
		return "FalklandIslandsMalvinas"
	case Fiji:
		return "Fiji"
	case Djibouti:
		return "Djibouti"
	case Gambia:
		return "Gambia"
	case Ghana:
		return "Ghana"
	case Gibraltar:
		return "Gibraltar"
	case Guatemala:
		return "Guatemala"
	case Guinea:
		return "Guinea"
	case Guyana:
		return "Guyana"
	case Haiti:
		return "Haiti"
	case Honduras:
		return "Honduras"
	case HongKong:
		return "HongKong"
	case Hungary:
		return "Hungary"
	case Iceland:
		return "Iceland"
	//Bhutan CountryCodeType = 356
	case India:
		return "India"
	case Indonesia:
		return "Indonesia"
	case IranIslamicRepublicOf:
		return "IranIslamicRepublicOf"
	case Iraq:
		return "Iraq"
	case Israel:
		return "Israel"
	case Jamaica:
		return "Jamaica"
	case Japan:
		return "Japan"
	case Kazakhstan:
		return "Kazakhstan"
	case Jordan:
		return "Jordan"
	case Kenya:
		return "Kenya"
	case KoreaDemocraticPeoplesRepublicOf:
		return "KoreaDemocraticPeoplesRepublicOf"
	case KoreaRepublicOf:
		return "KoreaRepublicOf"
	case Kuwait:
		return "Kuwait"
	case Kyrgyzstan:
		return "Kyrgyzstan"
	case LaoPeoplesDemocraticRepublic:
		return "LaoPeoplesDemocraticRepublic"
	case Lebanon:
		return "Lebanon"
	case Lesotho:
		return "Lesotho"
	case Latvia:
		return "Latvia"
	case Liberia:
		return "Liberia"
	case LibyanArabJamahiriya:
		return "LibyanArabJamahiriya"
	case Lithuania:
		return "Lithuania"
	case Macao:
		return "Macao"
	case Madagascar:
		return "Madagascar"
	case Malawi:
		return "Malawi"
	case Malaysia:
		return "Malaysia"
	case Maldives:
		return "Maldives"
	case Malta:
		return "Malta"
	case Mauritania:
		return "Mauritania"
	case Mauritius:
		return "Mauritius"
	case Mexico:
		return "Mexico"
	case Mongolia:
		return "Mongolia"
	case MoldovaRepublicOf:
		return "MoldovaRepublicOf"
	case Morocco:
		return "Morocco"
	// case WesternSahara:
	// 	return "WesternSahara"
	case Mozambique:
		return "Mozambique"
	case Oman:
		return "Oman"
	case Namibia:
		return "Namibia"
	case Nepal:
		return "Nepal"
	case NetherlandsAntilles:
		return "NetherlandsAntilles"
	case Aruba:
		return "Aruba"
	case Vanuatu:
		return "Vanuatu"
	//CookIslands CountryCodeType = 554
	case NewZealand:
		return "NewZealand"
	//Niue CountryCodeType = 554
	//Pitcairn CountryCodeType = 554
	//Tokelau CountryCodeType = 554
	case Nicaragua:
		return "Nicaragua"
	case Nigeria:
		return "Nigeria"
	//BouvetIsland CountryCodeType = 578
	case Norway:
		return "Norway"
	//SvalbardAndJanMayen CountryCodeType = 578
	case Pakistan:
		return "Pakistan"
	case Panama:
		return "Panama"
	case PapuaNewGuinea:
		return "PapuaNewGuinea"
	case Paraguay:
		return "Paraguay"
	case Peru:
		return "Peru"
	case Philippines:
		return "Philippines"
	case GuineaBissau:
		return "GuineaBissau"
	case Qatar:
		return "Qatar"
	case Romania:
		return "Romania"
	case RussianFederation:
		return "RussianFederation"
	case Rwanda:
		return "Rwanda"
	case SaintHelena:
		return "SaintHelena"
	case SaoTomeAndPrincipe:
		return "SaoTomeAndPrincipe"
	case SaudiArabia:
		return "SaudiArabia"
	case Seychelles:
		return "Seychelles"
	case SierraLeone:
		return "SierraLeone"
	case Singapore:
		return "Singapore"
	case Slovakia:
		return "Slovakia"
	case VietNam:
		return "VietNam"
	case Slovenia:
		return "Slovenia"
	case Somalia:
		return "Somalia"
	//Lesotho CountryCodeType = 710
	//Namibia CountryCodeType = 710
	case SouthAfrica:
		return "SouthAfrica"
	case Zimbabwe:
		return "Zimbabwe"
	case Sudan:
		return "Sudan"
	case Swaziland:
		return "Swaziland"
	case Sweden:
		return "Sweden"
	//Liechtenstein CountryCodeType = 756
	case Switzerland:
		return "Switzerland"
	case SyrianArabRepublic:
		return "SyrianArabRepublic"
	case Thailand:
		return "Thailand"
	case Tonga:
		return "Tonga"
	case TrinidadAndTobago:
		return "TrinidadAndTobago"
	case UnitedArabEmirates:
		return "UnitedArabEmirates"
	case Tunisia:
		return "Tunisia"
	case Turkey:
		return "Turkey"
	case Turkmenistan:
		return "Turkmenistan"
	case Uganda:
		return "Uganda"
	case MacedoniaTheFormerYugoslavRepublicOf:
		return "MacedoniaTheFormerYugoslavRepublicOf"
	//RussianFederation CountryCodeType = 810
	case Egypt:
		return "Egypt"
	case UnitedKingdom:
		return "UnitedKingdom"
	case TanzaniaUnitedRepublicOf:
		return "TanzaniaUnitedRepublicOf"
	//AmericanSamoa CountryCodeType = 840
	//BritishIndianOceanTerritory CountryCodeType = 840
	//Ecuador CountryCodeType = 840
	//ElSalvador CountryCodeType = 840
	//Guam CountryCodeType = 840
	//Haiti840 CountryCodeType = 840
	//MarshallIslands CountryCodeType = 840
	//MicronesiaFederatedStatesOf CountryCodeType = 840
	//NorthernMarianaIslands CountryCodeType = 840
	//Palau CountryCodeType = 840
	//Panama CountryCodeType = 840
	//PuertoRico CountryCodeType = 840
	//TimorLeste CountryCodeType = 840
	//TurksAndCaicosIslands CountryCodeType = 840
	case UnitedStates:
		return "UnitedStates"
	//UnitedStatesMinorOutlyingIslands CountryCodeType = 840
	//VirginIslandsBritish CountryCodeType = 840
	//VirginIslandsUs CountryCodeType = 840
	case Uruguay:
		return "Uruguay"
	case Uzbekistan:
		return "Uzbekistan"
	case Venezuela:
		return "Venezuela"
	case Samoa:
		return "Samoa"
	case Yemen:
		return "Yemen"
	case SerbiaMontenegro:
		return "SerbiaMontenegro"
	case Zambia:
		return "Zambia"
	case TaiwanProvinceOfChina:
		return "TaiwanProvinceOfChina"
	case CentralAfrica:
		return "CentralAfrica"
	//Cameroon CountryCodeType = 950
	//CentralAfricanRepublic CountryCodeType = 950
	//Chad CountryCodeType = 950
	//Congo CountryCodeType = 950
	//EquatorialGuinea CountryCodeType = 950
	//Gabon CountryCodeType = 950
	case CaribbeanIslands:
		return "CaribbeanIslands"
	//Anguilla CountryCodeType = 951
	//AntiguaAndBarbuda CountryCodeType = 951
	//Dominica CountryCodeType = 951
	//Grenada CountryCodeType = 951
	//Montserrat CountryCodeType = 951
	//SaintKittsAndNevis CountryCodeType = 951
	//SaintLucia CountryCodeType = 951
	//SaintVincentAndTheGrenadines CountryCodeType = 951
	//Benin CountryCodeType = 952
	//BurkinaFaso CountryCodeType = 952
	//CoteDivoire CountryCodeType = 952
	//GuineaBissau CountryCodeType = 952
	//Mali CountryCodeType = 952
	//Niger CountryCodeType = 952
	//Senegal CountryCodeType = 952
	//Togo CountryCodeType = 952
	case FrenchPolynesia:
		return "FrenchPolynesia"
	//NewCaledonia CountryCodeType = 953
	//WallisAndFutuna CountryCodeType = 953
	case Suriname:
		return "Suriname"
	//Madagascar CountryCodeType = 969
	//Colombia CountryCodeType = 970
	case Afghanistan:
		return "Afghanistan"
	case Tajikistan:
		return "Tajikistan"
	case Angola:
		return "Angola"
	case Belarus:
		return "Belarus"
	case Bulgaria:
		return "Bulgaria"
	case CongoTheDemocraticRepublicOf:
		return "CongoTheDemocraticRepublicOf"
	case BosniaAndHerzegovina:
		return "BosniaAndHerzegovina"
	//Andorra CountryCodeType = 978
	//Austria CountryCodeType = 978
	//Belgium CountryCodeType = 978
	//Finland CountryCodeType = 978
	//France CountryCodeType = 978
	//FrenchGuiana CountryCodeType = 978
	//FrenchSouthernTerritories CountryCodeType = 978
	//Germany CountryCodeType = 978
	//Greece CountryCodeType = 978
	//Guadeloupe CountryCodeType = 978
	//HolySeeVaticanCityState CountryCodeType = 978
	//Ireland CountryCodeType = 978
	//Italy CountryCodeType = 978
	//Luxembourg CountryCodeType = 978
	//Martinique CountryCodeType = 978
	//Mayotte CountryCodeType = 978
	//Monaco CountryCodeType = 978
	//Netherlands CountryCodeType = 978
	//Portugal CountryCodeType = 978
	//Reunion CountryCodeType = 978
	//SaintPierreAndMiquelon CountryCodeType = 978
	//SanMarino CountryCodeType = 978
	//SerbiaMontenegro CountryCodeType = 978
	//Spain CountryCodeType = 978
	case WesternEurope:
		return "WesternEurope"
	//Mexico CountryCodeType = 979
	case Ukraine:
		return "Ukraine"
	case Georgia:
		return "Georgia"
	//Bolivia CountryCodeType = 984
	case Poland:
		return "Poland"
	case Brazil:
		return "Brazil"
	//Chile CountryCodeType = 990
	case UnitedStatesNd:
		return "UnitedStatesNd"
	case UnitedStatesSd:
		return "UnitedStatesSd"
	}
	return ""
}

type CurrencyCodeType uint

const (
	AfghanistanAfghani                             CurrencyCodeType = 971
	AlbaniaLek                                     CurrencyCodeType = 8
	AlgeriaAlgerianDinar                           CurrencyCodeType = 12
	AngolaKwanza                                   CurrencyCodeType = 973
	ArgentinaArgentinePeso                         CurrencyCodeType = 32
	ArmeniaArmenianDram                            CurrencyCodeType = 51
	ArubaArubanGuilder                             CurrencyCodeType = 533
	AustraliaAustralianDollar                      CurrencyCodeType = 36
	AzerbaijanAzerbaijanianManat                   CurrencyCodeType = 31
	BahamasBahamianDollar                          CurrencyCodeType = 44
	BahrainBahrainiDinar                           CurrencyCodeType = 48
	BangladeshTaka                                 CurrencyCodeType = 50
	BarbadosBarbadosDollar                         CurrencyCodeType = 52
	BelarusBelarussianRuble                        CurrencyCodeType = 974
	BelizeBelizeDollar                             CurrencyCodeType = 84
	BermudaBermudianDollar                         CurrencyCodeType = 60
	BhutanNgultrum                                 CurrencyCodeType = 64
	BoliviaBoliviano                               CurrencyCodeType = 68
	BoliviaMvdol                                   CurrencyCodeType = 984
	BosniaAndHerzegovinaConvertibleMarks           CurrencyCodeType = 977
	BotswanaPula                                   CurrencyCodeType = 72
	BrazilBrazilianReal                            CurrencyCodeType = 986
	BruneiDarussalamBruneiDollar                   CurrencyCodeType = 96
	BulgariaBulgarianLev                           CurrencyCodeType = 975
	BurundiBurundiFranc                            CurrencyCodeType = 108
	CambodiaRiel                                   CurrencyCodeType = 116
	CanadaCanadianDollar                           CurrencyCodeType = 124
	CapeVerdeCapeVerdeEscudo                       CurrencyCodeType = 132
	CaymanIslandsCaymanIslandsDollar               CurrencyCodeType = 136
	CentralAfricanRepublicCfaFrancBeac             CurrencyCodeType = 950
	CfaFrancBceao                                  CurrencyCodeType = 952
	ChileChileanPeso                               CurrencyCodeType = 152
	ChileUnidadesDeFomento                         CurrencyCodeType = 990
	ChinaYuanRenminbi                              CurrencyCodeType = 156
	ColombiaColombianPeso                          CurrencyCodeType = 170
	ColombiaUnidadDeValorReal                      CurrencyCodeType = 970
	ComorosComoroFranc                             CurrencyCodeType = 174
	CongoTheDemocraticRepublicOfFrancCongolais     CurrencyCodeType = 976
	CostaRicaCostaRicanColon                       CurrencyCodeType = 188
	CroatiaCroatianKuna                            CurrencyCodeType = 191
	CubaCubanPeso                                  CurrencyCodeType = 192
	CyprusCyprusPound                              CurrencyCodeType = 196
	CzechRepublicCzechKoruna                       CurrencyCodeType = 203
	DenmarkDanishKrone                             CurrencyCodeType = 208
	DjiboutiDjiboutiFranc                          CurrencyCodeType = 262
	DominicanRepublicDominicanPeso                 CurrencyCodeType = 214
	EastCarribbeanDollar                           CurrencyCodeType = 951
	EgyptEgyptianPound                             CurrencyCodeType = 818
	ElSalvadorElSalvadorColon                      CurrencyCodeType = 222
	EritreaNakfa                                   CurrencyCodeType = 232
	EstoniaKroon                                   CurrencyCodeType = 233
	EthiopiaEthiopianBirr                          CurrencyCodeType = 230
	Euro                                           CurrencyCodeType = 978
	FalklandIslandsMalvinasFalklandIslandsPound    CurrencyCodeType = 238
	FijiFijiDollar                                 CurrencyCodeType = 242
	FrenchPolynesiaCfpFranc                        CurrencyCodeType = 953
	GambiaDalasi                                   CurrencyCodeType = 270
	GeorgiaLari                                    CurrencyCodeType = 981
	GhanaCedi                                      CurrencyCodeType = 288
	GibraltarGibraltarPound                        CurrencyCodeType = 292
	GuatemalaQuetzal                               CurrencyCodeType = 320
	GuineaBissauGuineaBissauPeso                   CurrencyCodeType = 624
	GuineaGuineaFranc                              CurrencyCodeType = 324
	GuyanaGuyanaDollar                             CurrencyCodeType = 328
	HaitiGourde                                    CurrencyCodeType = 332
	HondurasLempira                                CurrencyCodeType = 340
	HongKongHongKongDollar                         CurrencyCodeType = 344
	HungaryForint                                  CurrencyCodeType = 348
	IcelandIcelandKrona                            CurrencyCodeType = 352
	IndiaIndianRupee                               CurrencyCodeType = 356
	IndonesiaRupiah                                CurrencyCodeType = 360
	IranIslamicRepublicOfIranianRial               CurrencyCodeType = 364
	IraqIraqiDinar                                 CurrencyCodeType = 368
	IsraelNewIsraeliSheqel                         CurrencyCodeType = 376
	JamaicaJamaicanDollar                          CurrencyCodeType = 388
	JapanYen                                       CurrencyCodeType = 392
	JordanJordanianDinar                           CurrencyCodeType = 400
	KazakhstanTenge                                CurrencyCodeType = 398
	KenyaKenyanShilling                            CurrencyCodeType = 404
	KoreaDemocraticPeoplesRepublicOfNorthKoreanWon CurrencyCodeType = 408
	KoreaRepublicOfWon                             CurrencyCodeType = 410
	KuwaitKuwaitiDinar                             CurrencyCodeType = 414
	KyrgyzstanSom                                  CurrencyCodeType = 417
	LaoPeoplesDemocraticRepublicKip                CurrencyCodeType = 418
	LatviaLatvianLats                              CurrencyCodeType = 428
	LebanonLebanesePound                           CurrencyCodeType = 422
	LesothoLoti                                    CurrencyCodeType = 426
	LiberiaLiberianDollar                          CurrencyCodeType = 430
	LibyanArabJamahiriyaLybianDinar                CurrencyCodeType = 434
	LithuaniaLithuanianLitas                       CurrencyCodeType = 440
	MacaoPataca                                    CurrencyCodeType = 446
	MacedoniaTheFormerYugoslavRepublicOfDenar      CurrencyCodeType = 807
	MadagascarAriary                               CurrencyCodeType = 969
	MadagascarMalagasyFranc                        CurrencyCodeType = 450
	MalawiKwacha                                   CurrencyCodeType = 454
	MalaysiaMalaysianRinggit                       CurrencyCodeType = 458
	MaldivesRufiyaa                                CurrencyCodeType = 462
	MaltaMalteseLira                               CurrencyCodeType = 470
	MauritaniaOuguiya                              CurrencyCodeType = 478
	MauritiusMauritiusRupee                        CurrencyCodeType = 480
	MexicoMexicanPeso                              CurrencyCodeType = 484
	MexicoMexicanUnidadDeInversionUdi              CurrencyCodeType = 979
	MoldovaRepublicOfMoldovanLeu                   CurrencyCodeType = 498
	MongoliaTugrik                                 CurrencyCodeType = 496
	MoroccoMoroccanDirham                          CurrencyCodeType = 504
	MozambiqueMetical                              CurrencyCodeType = 508
	MyanmarKyat                                    CurrencyCodeType = 104
	NamibiaNamibiaDollar                           CurrencyCodeType = 516
	NepalNepaleseRupee                             CurrencyCodeType = 524
	NetherlandsAntillesNetherlandsAntillanGuilder  CurrencyCodeType = 532
	NewZealandNewZealandDollar                     CurrencyCodeType = 554
	NicaraguaCordobaOro                            CurrencyCodeType = 558
	NigeriaNaira                                   CurrencyCodeType = 566
	NorwayNorwegianKrone                           CurrencyCodeType = 578
	OmanRialOmani                                  CurrencyCodeType = 512
	PakistanPakistanRupee                          CurrencyCodeType = 586
	PanamaBalboa                                   CurrencyCodeType = 590
	PapuaNewGuineaKina                             CurrencyCodeType = 598
	ParaguayGuarani                                CurrencyCodeType = 600
	PeruNuevoSol                                   CurrencyCodeType = 604
	PhilippinesPhilippinePeso                      CurrencyCodeType = 608
	PolandZloty                                    CurrencyCodeType = 985
	QatarQatariRial                                CurrencyCodeType = 634
	RomaniaLeu                                     CurrencyCodeType = 642
	RussianFederationRussianRuble                  CurrencyCodeType = 643
	RwandaRwandaFranc                              CurrencyCodeType = 646
	SaintHelenaSaintHelenaPound                    CurrencyCodeType = 654
	SamoaTala                                      CurrencyCodeType = 882
	SaoTomeAndPrincipeDobra                        CurrencyCodeType = 678
	SaudiArabiaSaudiRiyal                          CurrencyCodeType = 682
	SerbiaMontenegroSerbianDinar                   CurrencyCodeType = 891
	SeychellesSeychellesRupee                      CurrencyCodeType = 690
	SierraLeoneLeone                               CurrencyCodeType = 694
	SingaporeSingaporeDollar                       CurrencyCodeType = 702
	SlovakiaSlovakKoruna                           CurrencyCodeType = 703
	SloveniaTolar                                  CurrencyCodeType = 705
	SolomonIslandsSolomonIslandsDollar             CurrencyCodeType = 90
	SomaliaSomaliShilling                          CurrencyCodeType = 706
	SouthAfricaRand                                CurrencyCodeType = 710
	SriLankaSriLankaRupee                          CurrencyCodeType = 144
	SudanSudaneseDinar                             CurrencyCodeType = 736
	SurinameSurinameDollar                         CurrencyCodeType = 968
	SwazilandLilangeni                             CurrencyCodeType = 748
	SwedenSwedishKrona                             CurrencyCodeType = 752
	SwitzerlandSwissFranc                          CurrencyCodeType = 756
	SyrianArabRepublicSyrianPound                  CurrencyCodeType = 760
	TaiwanProvinceOfChinaNewTaiwanDollar           CurrencyCodeType = 901
	TajikistanSomoni                               CurrencyCodeType = 972
	TanzaniaUnitedRepublicOfTanzanianShilling      CurrencyCodeType = 834
	ThailandBaht                                   CurrencyCodeType = 764
	TongaPaanga                                    CurrencyCodeType = 776
	TrinidadAndTobagoTrinidadAndTobagoDollar       CurrencyCodeType = 780
	TunisiaTunisianDinar                           CurrencyCodeType = 788
	TurkeyTurkishLira                              CurrencyCodeType = 792
	TurkmenistanManat                              CurrencyCodeType = 795
	UgandaUgandaShilling                           CurrencyCodeType = 800
	UkraineHryvnia                                 CurrencyCodeType = 980
	UnitedArabEmiratesUaeDirham                    CurrencyCodeType = 784
	UnitedKingdomPoundSterling                     CurrencyCodeType = 826
	UnitedStatesUsDollar                           CurrencyCodeType = 840
	UnitedStatesUsDollarNextDay                    CurrencyCodeType = 997
	UnitedStatesUsDollarSameDay                    CurrencyCodeType = 998
	UruguayPesoUruguayo                            CurrencyCodeType = 858
	UzbekistanUzbekistanSum                        CurrencyCodeType = 860
	VanuatuVatu                                    CurrencyCodeType = 548
	VenezuelaBolivar                               CurrencyCodeType = 862
	VietNamDong                                    CurrencyCodeType = 704
	YemenYemeniRial                                CurrencyCodeType = 886
	ZambiaKwacha                                   CurrencyCodeType = 894
	ZimbabweZimbabweDollar                         CurrencyCodeType = 716
)

func (x CurrencyCodeType) String() string {
	switch x {
	case UnitedArabEmiratesUaeDirham:
		return "UnitedArabEmiratesUaeDirham"
	case AfghanistanAfghani:
		return "AfghanistanAfghani"
	case AlbaniaLek:
		return "AlbaniaLek"
	case ArmeniaArmenianDram:
		return "ArmeniaArmenianDram"
	case NetherlandsAntillesNetherlandsAntillanGuilder:
		return "NetherlandsAntillesNetherlandsAntillanGuilder"
	case AngolaKwanza:
		return "AngolaKwanza"
	case ArgentinaArgentinePeso:
		return "ArgentinaArgentinePeso"
	case AustraliaAustralianDollar:
		return "AustraliaAustralianDollar"
	case ArubaArubanGuilder:
		return "ArubaArubanGuilder"
	case AzerbaijanAzerbaijanianManat:
		return "AzerbaijanAzerbaijanianManat"
	case BosniaAndHerzegovinaConvertibleMarks:
		return "BosniaAndHerzegovinaConvertibleMarks"
	case BarbadosBarbadosDollar:
		return "BarbadosBarbadosDollar"
	case BangladeshTaka:
		return "BangladeshTaka"
	case BulgariaBulgarianLev:
		return "BulgariaBulgarianLev"
	case BahrainBahrainiDinar:
		return "BahrainBahrainiDinar"
	case BurundiBurundiFranc:
		return "BurundiBurundiFranc"
	case BermudaBermudianDollar:
		return "BermudaBermudianDollar"
	case BruneiDarussalamBruneiDollar:
		return "BruneiDarussalamBruneiDollar"
	case BoliviaBoliviano:
		return "BoliviaBoliviano"
	case BoliviaMvdol:
		return "BoliviaMvdol"
	case BrazilBrazilianReal:
		return "BrazilBrazilianReal"
	case BahamasBahamianDollar:
		return "BahamasBahamianDollar"
	case BhutanNgultrum:
		return "BhutanNgultrum"
	case BotswanaPula:
		return "BotswanaPula"
	case BelarusBelarussianRuble:
		return "BelarusBelarussianRuble"
	case BelizeBelizeDollar:
		return "BelizeBelizeDollar"
	case CanadaCanadianDollar:
		return "CanadaCanadianDollar"
	case CongoTheDemocraticRepublicOfFrancCongolais:
		return "CongoTheDemocraticRepublicOfFrancCongolais"
	case SwitzerlandSwissFranc:
		return "SwitzerlandSwissFranc"
	case ChileUnidadesDeFomento:
		return "ChileUnidadesDeFomento"
	case ChileChileanPeso:
		return "ChileChileanPeso"
	case ChinaYuanRenminbi:
		return "ChinaYuanRenminbi"
	case ColombiaColombianPeso:
		return "ColombiaColombianPeso"
	case ColombiaUnidadDeValorReal:
		return "ColombiaUnidadDeValorReal"
	case CostaRicaCostaRicanColon:
		return "CostaRicaCostaRicanColon"
	case SerbiaMontenegroSerbianDinar:
		return "SerbiaMontenegroSerbianDinar"
	case CubaCubanPeso:
		return "CubaCubanPeso"
	case CapeVerdeCapeVerdeEscudo:
		return "CapeVerdeCapeVerdeEscudo"
	case CyprusCyprusPound:
		return "CyprusCyprusPound"
	case CzechRepublicCzechKoruna:
		return "CzechRepublicCzechKoruna"
	case DjiboutiDjiboutiFranc:
		return "DjiboutiDjiboutiFranc"
	case DenmarkDanishKrone:
		return "DenmarkDanishKrone"
	case DominicanRepublicDominicanPeso:
		return "DominicanRepublicDominicanPeso"
	case AlgeriaAlgerianDinar:
		return "AlgeriaAlgerianDinar"
	case EstoniaKroon:
		return "EstoniaKroon"
	case EgyptEgyptianPound:
		return "EgyptEgyptianPound"
	case EritreaNakfa:
		return "EritreaNakfa"
	case EthiopiaEthiopianBirr:
		return "EthiopiaEthiopianBirr"
	case Euro:
		return "Euro"
	case FijiFijiDollar:
		return "FijiFijiDollar"
	case FalklandIslandsMalvinasFalklandIslandsPound:
		return "FalklandIslandsMalvinasFalklandIslandsPound"
	case UnitedKingdomPoundSterling:
		return "UnitedKingdomPoundSterling"
	case GeorgiaLari:
		return "GeorgiaLari"
	case GhanaCedi:
		return "GhanaCedi"
	case GibraltarGibraltarPound:
		return "GibraltarGibraltarPound"
	case GambiaDalasi:
		return "GambiaDalasi"
	case GuineaGuineaFranc:
		return "GuineaGuineaFranc"
	case GuatemalaQuetzal:
		return "GuatemalaQuetzal"
	case GuineaBissauGuineaBissauPeso:
		return "GuineaBissauGuineaBissauPeso"
	case GuyanaGuyanaDollar:
		return "GuyanaGuyanaDollar"
	case HongKongHongKongDollar:
		return "HongKongHongKongDollar"
	case HondurasLempira:
		return "HondurasLempira"
	case CroatiaCroatianKuna:
		return "CroatiaCroatianKuna"
	case HaitiGourde:
		return "HaitiGourde"
	case HungaryForint:
		return "HungaryForint"
	case IndonesiaRupiah:
		return "IndonesiaRupiah"
	case IsraelNewIsraeliSheqel:
		return "IsraelNewIsraeliSheqel"
	case IndiaIndianRupee:
		return "IndiaIndianRupee"
	case IraqIraqiDinar:
		return "IraqIraqiDinar"
	case IranIslamicRepublicOfIranianRial:
		return "IranIslamicRepublicOfIranianRial"
	case IcelandIcelandKrona:
		return "IcelandIcelandKrona"
	case JamaicaJamaicanDollar:
		return "JamaicaJamaicanDollar"
	case JordanJordanianDinar:
		return "JordanJordanianDinar"
	case JapanYen:
		return "JapanYen"
	case KenyaKenyanShilling:
		return "KenyaKenyanShilling"
	case KyrgyzstanSom:
		return "KyrgyzstanSom"
	case CambodiaRiel:
		return "CambodiaRiel"
	case ComorosComoroFranc:
		return "ComorosComoroFranc"
	case KoreaDemocraticPeoplesRepublicOfNorthKoreanWon:
		return "KoreaDemocraticPeoplesRepublicOfNorthKoreanWon"
	case KoreaRepublicOfWon:
		return "KoreaRepublicOfWon"
	case KuwaitKuwaitiDinar:
		return "KuwaitKuwaitiDinar"
	case CaymanIslandsCaymanIslandsDollar:
		return "CaymanIslandsCaymanIslandsDollar"
	case KazakhstanTenge:
		return "KazakhstanTenge"
	case LaoPeoplesDemocraticRepublicKip:
		return "LaoPeoplesDemocraticRepublicKip"
	case LebanonLebanesePound:
		return "LebanonLebanesePound"
	case SriLankaSriLankaRupee:
		return "SriLankaSriLankaRupee"
	case LiberiaLiberianDollar:
		return "LiberiaLiberianDollar"
	case LesothoLoti:
		return "LesothoLoti"
	case LithuaniaLithuanianLitas:
		return "LithuaniaLithuanianLitas"
	case LatviaLatvianLats:
		return "LatviaLatvianLats"
	case LibyanArabJamahiriyaLybianDinar:
		return "LibyanArabJamahiriyaLybianDinar"
	case MoroccoMoroccanDirham:
		return "MoroccoMoroccanDirham"
	case MoldovaRepublicOfMoldovanLeu:
		return "MoldovaRepublicOfMoldovanLeu"
	case MadagascarAriary:
		return "MadagascarAriary"
	case MadagascarMalagasyFranc:
		return "MadagascarMalagasyFranc"
	case MacedoniaTheFormerYugoslavRepublicOfDenar:
		return "MacedoniaTheFormerYugoslavRepublicOfDenar"
	case MyanmarKyat:
		return "MyanmarKyat"
	case MongoliaTugrik:
		return "MongoliaTugrik"
	case MacaoPataca:
		return "MacaoPataca"
	case MauritaniaOuguiya:
		return "MauritaniaOuguiya"
	case MaltaMalteseLira:
		return "MaltaMalteseLira"
	case MauritiusMauritiusRupee:
		return "MauritiusMauritiusRupee"
	case MaldivesRufiyaa:
		return "MaldivesRufiyaa"
	case MalawiKwacha:
		return "MalawiKwacha"
	case MexicoMexicanPeso:
		return "MexicoMexicanPeso"
	case MexicoMexicanUnidadDeInversionUdi:
		return "MexicoMexicanUnidadDeInversionUdi"
	case MalaysiaMalaysianRinggit:
		return "MalaysiaMalaysianRinggit"
	case MozambiqueMetical:
		return "MozambiqueMetical"
	case NamibiaNamibiaDollar:
		return "NamibiaNamibiaDollar"
	case NigeriaNaira:
		return "NigeriaNaira"
	case NicaraguaCordobaOro:
		return "NicaraguaCordobaOro"
	case NorwayNorwegianKrone:
		return "NorwayNorwegianKrone"
	case NepalNepaleseRupee:
		return "NepalNepaleseRupee"
	case NewZealandNewZealandDollar:
		return "NewZealandNewZealandDollar"
	case OmanRialOmani:
		return "OmanRialOmani"
	case PanamaBalboa:
		return "PanamaBalboa"
	case PeruNuevoSol:
		return "PeruNuevoSol"
	case PapuaNewGuineaKina:
		return "PapuaNewGuineaKina"
	case PhilippinesPhilippinePeso:
		return "PhilippinesPhilippinePeso"
	case PakistanPakistanRupee:
		return "PakistanPakistanRupee"
	case PolandZloty:
		return "PolandZloty"
	case ParaguayGuarani:
		return "ParaguayGuarani"
	case QatarQatariRial:
		return "QatarQatariRial"
	case RomaniaLeu:
		return "RomaniaLeu"
	case RussianFederationRussianRuble:
		return "RussianFederationRussianRuble"
	case RwandaRwandaFranc:
		return "RwandaRwandaFranc"
	case SaudiArabiaSaudiRiyal:
		return "SaudiArabiaSaudiRiyal"
	case SolomonIslandsSolomonIslandsDollar:
		return "SolomonIslandsSolomonIslandsDollar"
	case SeychellesSeychellesRupee:
		return "SeychellesSeychellesRupee"
	case SudanSudaneseDinar:
		return "SudanSudaneseDinar"
	case SwedenSwedishKrona:
		return "SwedenSwedishKrona"
	case SingaporeSingaporeDollar:
		return "SingaporeSingaporeDollar"
	case SaintHelenaSaintHelenaPound:
		return "SaintHelenaSaintHelenaPound"
	case SloveniaTolar:
		return "SloveniaTolar"
	case SlovakiaSlovakKoruna:
		return "SlovakiaSlovakKoruna"
	case SierraLeoneLeone:
		return "SierraLeoneLeone"
	case SomaliaSomaliShilling:
		return "SomaliaSomaliShilling"
	case SurinameSurinameDollar:
		return "SurinameSurinameDollar"
	case SaoTomeAndPrincipeDobra:
		return "SaoTomeAndPrincipeDobra"
	case ElSalvadorElSalvadorColon:
		return "ElSalvadorElSalvadorColon"
	case SyrianArabRepublicSyrianPound:
		return "SyrianArabRepublicSyrianPound"
	case SwazilandLilangeni:
		return "SwazilandLilangeni"
	case ThailandBaht:
		return "ThailandBaht"
	case TajikistanSomoni:
		return "TajikistanSomoni"
	case TurkmenistanManat:
		return "TurkmenistanManat"
	case TunisiaTunisianDinar:
		return "TunisiaTunisianDinar"
	case TongaPaanga:
		return "TongaPaanga"
	case TurkeyTurkishLira:
		return "TurkeyTurkishLira"
	case TrinidadAndTobagoTrinidadAndTobagoDollar:
		return "TrinidadAndTobagoTrinidadAndTobagoDollar"
	case TaiwanProvinceOfChinaNewTaiwanDollar:
		return "TaiwanProvinceOfChinaNewTaiwanDollar"
	case TanzaniaUnitedRepublicOfTanzanianShilling:
		return "TanzaniaUnitedRepublicOfTanzanianShilling"
	case UkraineHryvnia:
		return "UkraineHryvnia"
	case UgandaUgandaShilling:
		return "UgandaUgandaShilling"
	case UnitedStatesUsDollar:
		return "UnitedStatesUsDollar"
	case UnitedStatesUsDollarNextDay:
		return "UnitedStatesUsDollarNextDay"
	case UnitedStatesUsDollarSameDay:
		return "UnitedStatesUsDollarSameDay"
	case UruguayPesoUruguayo:
		return "UruguayPesoUruguayo"
	case UzbekistanUzbekistanSum:
		return "UzbekistanUzbekistanSum"
	case VenezuelaBolivar:
		return "VenezuelaBolivar"
	case VietNamDong:
		return "VietNamDong"
	case VanuatuVatu:
		return "VanuatuVatu"
	case SamoaTala:
		return "SamoaTala"
	case CentralAfricanRepublicCfaFrancBeac:
		return "CentralAfricanRepublicCfaFrancBeac"
	case EastCarribbeanDollar:
		return "EastCarribbeanDollar"
	case CfaFrancBceao:
		return "CfaFrancBceao"
	case FrenchPolynesiaCfpFranc:
		return "FrenchPolynesiaCfpFranc"
	case YemenYemeniRial:
		return "YemenYemeniRial"
	case SouthAfricaRand:
		return "SouthAfricaRand"
	case ZambiaKwacha:
		return "ZambiaKwacha"
	case ZimbabweZimbabweDollar:
		return "ZimbabweZimbabweDollar"
	}
	return ""
}

func (x CurrencyCodeType) IsoCode() string {
	switch x {
	case UnitedArabEmiratesUaeDirham:
		return "AED"
	case AfghanistanAfghani:
		return "AFN"
	case AlbaniaLek:
		return "ALL"
	case ArmeniaArmenianDram:
		return "AMD"
	case NetherlandsAntillesNetherlandsAntillanGuilder:
		return "ANG"
	case AngolaKwanza:
		return "AOA"
	case ArgentinaArgentinePeso:
		return "ARS"
	case AustraliaAustralianDollar:
		return "AUD"
	case ArubaArubanGuilder:
		return "AWG"
	case AzerbaijanAzerbaijanianManat:
		return "AZM"
	case BosniaAndHerzegovinaConvertibleMarks:
		return "BAM"
	case BarbadosBarbadosDollar:
		return "BBD"
	case BangladeshTaka:
		return "BDT"
	case BulgariaBulgarianLev:
		return "BGN"
	case BahrainBahrainiDinar:
		return "BHD"
	case BurundiBurundiFranc:
		return "BIF"
	case BermudaBermudianDollar:
		return "BMD"
	case BruneiDarussalamBruneiDollar:
		return "BND"
	case BoliviaBoliviano:
		return "BOB"
	case BoliviaMvdol:
		return "BOV"
	case BrazilBrazilianReal:
		return "BRL"
	case BahamasBahamianDollar:
		return "BSD"
	case BhutanNgultrum:
		return "BTN"
	case BotswanaPula:
		return "BWP"
	case BelarusBelarussianRuble:
		return "BYR"
	case BelizeBelizeDollar:
		return "BZD"
	case CanadaCanadianDollar:
		return "CAD"
	case CongoTheDemocraticRepublicOfFrancCongolais:
		return "CDF"
	case SwitzerlandSwissFranc:
		return "CHF"
	case ChileUnidadesDeFomento:
		return "CLF"
	case ChileChileanPeso:
		return "CLP"
	case ChinaYuanRenminbi:
		return "CNY"
	case ColombiaColombianPeso:
		return "COP"
	case ColombiaUnidadDeValorReal:
		return "COU"
	case CostaRicaCostaRicanColon:
		return "CRC"
	case SerbiaMontenegroSerbianDinar:
		return "CSD"
	case CubaCubanPeso:
		return "CUP"
	case CapeVerdeCapeVerdeEscudo:
		return "CVE"
	case CyprusCyprusPound:
		return "CYP"
	case CzechRepublicCzechKoruna:
		return "CZK"
	case DjiboutiDjiboutiFranc:
		return "DJF"
	case DenmarkDanishKrone:
		return "DKK"
	case DominicanRepublicDominicanPeso:
		return "DOP"
	case AlgeriaAlgerianDinar:
		return "DZD"
	case EstoniaKroon:
		return "EEK"
	case EgyptEgyptianPound:
		return "EGP"
	case EritreaNakfa:
		return "ERN"
	case EthiopiaEthiopianBirr:
		return "ETB"
	case Euro:
		return "EUR"
	case FijiFijiDollar:
		return "FJD"
	case FalklandIslandsMalvinasFalklandIslandsPound:
		return "FKP"
	case UnitedKingdomPoundSterling:
		return "GBP"
	case GeorgiaLari:
		return "GEL"
	case GhanaCedi:
		return "GHC"
	case GibraltarGibraltarPound:
		return "GIP"
	case GambiaDalasi:
		return "GMD"
	case GuineaGuineaFranc:
		return "GNF"
	case GuatemalaQuetzal:
		return "GTQ"
	case GuineaBissauGuineaBissauPeso:
		return "GWP"
	case GuyanaGuyanaDollar:
		return "GYD"
	case HongKongHongKongDollar:
		return "HKD"
	case HondurasLempira:
		return "HNL"
	case CroatiaCroatianKuna:
		return "HRK"
	case HaitiGourde:
		return "HTG"
	case HungaryForint:
		return "HUF"
	case IndonesiaRupiah:
		return "IDR"
	case IsraelNewIsraeliSheqel:
		return "ILS"
	case IndiaIndianRupee:
		return "INR"
	case IraqIraqiDinar:
		return "IQD"
	case IranIslamicRepublicOfIranianRial:
		return "IRR"
	case IcelandIcelandKrona:
		return "ISK"
	case JamaicaJamaicanDollar:
		return "JMD"
	case JordanJordanianDinar:
		return "JOD"
	case JapanYen:
		return "JPY"
	case KenyaKenyanShilling:
		return "KES"
	case KyrgyzstanSom:
		return "KGS"
	case CambodiaRiel:
		return "KHR"
	case ComorosComoroFranc:
		return "KMF"
	case KoreaDemocraticPeoplesRepublicOfNorthKoreanWon:
		return "KPW"
	case KoreaRepublicOfWon:
		return "KRW"
	case KuwaitKuwaitiDinar:
		return "KWD"
	case CaymanIslandsCaymanIslandsDollar:
		return "KYD"
	case KazakhstanTenge:
		return "KZT"
	case LaoPeoplesDemocraticRepublicKip:
		return "LAK"
	case LebanonLebanesePound:
		return "LBP"
	case SriLankaSriLankaRupee:
		return "LKR"
	case LiberiaLiberianDollar:
		return "LRD"
	case LesothoLoti:
		return "LSL"
	case LithuaniaLithuanianLitas:
		return "LTL"
	case LatviaLatvianLats:
		return "LVL"
	case LibyanArabJamahiriyaLybianDinar:
		return "LYD"
	case MoroccoMoroccanDirham:
		return "MAD"
	case MoldovaRepublicOfMoldovanLeu:
		return "MDL"
	case MadagascarAriary:
		return "MGA"
	case MadagascarMalagasyFranc:
		return "MGF"
	case MacedoniaTheFormerYugoslavRepublicOfDenar:
		return "MKD"
	case MyanmarKyat:
		return "MMK"
	case MongoliaTugrik:
		return "MNT"
	case MacaoPataca:
		return "MOP"
	case MauritaniaOuguiya:
		return "MRO"
	case MaltaMalteseLira:
		return "MTL"
	case MauritiusMauritiusRupee:
		return "MUR"
	case MaldivesRufiyaa:
		return "MVR"
	case MalawiKwacha:
		return "MWK"
	case MexicoMexicanPeso:
		return "MXN"
	case MexicoMexicanUnidadDeInversionUdi:
		return "MXV"
	case MalaysiaMalaysianRinggit:
		return "MYR"
	case MozambiqueMetical:
		return "MZM"
	case NamibiaNamibiaDollar:
		return "NAD"
	case NigeriaNaira:
		return "NGN"
	case NicaraguaCordobaOro:
		return "NIO"
	case NorwayNorwegianKrone:
		return "NOK"
	case NepalNepaleseRupee:
		return "NPR"
	case NewZealandNewZealandDollar:
		return "NZD"
	case OmanRialOmani:
		return "OMR"
	case PanamaBalboa:
		return "PAB"
	case PeruNuevoSol:
		return "PEN"
	case PapuaNewGuineaKina:
		return "PGK"
	case PhilippinesPhilippinePeso:
		return "PHP"
	case PakistanPakistanRupee:
		return "PKR"
	case PolandZloty:
		return "PLN"
	case ParaguayGuarani:
		return "PYG"
	case QatarQatariRial:
		return "QAR"
	case RomaniaLeu:
		return "ROL"
	case RussianFederationRussianRuble:
		return "RUB"
	case RwandaRwandaFranc:
		return "RWF"
	case SaudiArabiaSaudiRiyal:
		return "SAR"
	case SolomonIslandsSolomonIslandsDollar:
		return "SBD"
	case SeychellesSeychellesRupee:
		return "SCR"
	case SudanSudaneseDinar:
		return "SDD"
	case SwedenSwedishKrona:
		return "SEK"
	case SingaporeSingaporeDollar:
		return "SGD"
	case SaintHelenaSaintHelenaPound:
		return "SHP"
	case SloveniaTolar:
		return "SIT"
	case SlovakiaSlovakKoruna:
		return "SKK"
	case SierraLeoneLeone:
		return "SLL"
	case SomaliaSomaliShilling:
		return "SOS"
	case SurinameSurinameDollar:
		return "SRD"
	case SaoTomeAndPrincipeDobra:
		return "STD"
	case ElSalvadorElSalvadorColon:
		return "SVC"
	case SyrianArabRepublicSyrianPound:
		return "SYP"
	case SwazilandLilangeni:
		return "SZL"
	case ThailandBaht:
		return "THB"
	case TajikistanSomoni:
		return "TJS"
	case TurkmenistanManat:
		return "TMM"
	case TunisiaTunisianDinar:
		return "TND"
	case TongaPaanga:
		return "TOP"
	case TurkeyTurkishLira:
		return "TRL"
	case TrinidadAndTobagoTrinidadAndTobagoDollar:
		return "TTD"
	case TaiwanProvinceOfChinaNewTaiwanDollar:
		return "TWD"
	case TanzaniaUnitedRepublicOfTanzanianShilling:
		return "TZS"
	case UkraineHryvnia:
		return "UAH"
	case UgandaUgandaShilling:
		return "UGX"
	case UnitedStatesUsDollar:
		return "USD"
	case UnitedStatesUsDollarNextDay:
		return "USN"
	case UnitedStatesUsDollarSameDay:
		return "USS"
	case UruguayPesoUruguayo:
		return "UYU"
	case UzbekistanUzbekistanSum:
		return "UZS"
	case VenezuelaBolivar:
		return "VEB"
	case VietNamDong:
		return "VND"
	case VanuatuVatu:
		return "VUV"
	case SamoaTala:
		return "WST"
	case CentralAfricanRepublicCfaFrancBeac:
		return "XAF"
	case EastCarribbeanDollar:
		return "XCD"
	case CfaFrancBceao:
		return "XOF"
	case FrenchPolynesiaCfpFranc:
		return "XPF"
	case YemenYemeniRial:
		return "YER"
	case SouthAfricaRand:
		return "ZAR"
	case ZambiaKwacha:
		return "ZMK"
	case ZimbabweZimbabweDollar:
		return "ZWD"
	}
	return ""
}

type CurrencyExponentIndicatorType uint

const (
	Aed CurrencyExponentIndicatorType = 2
	Afn CurrencyExponentIndicatorType = 2
	All CurrencyExponentIndicatorType = 2
	Amd CurrencyExponentIndicatorType = 2
	Ang CurrencyExponentIndicatorType = 2
	Aoa CurrencyExponentIndicatorType = 2
	Ars CurrencyExponentIndicatorType = 2
	Aud CurrencyExponentIndicatorType = 2
	Awg CurrencyExponentIndicatorType = 2
	Azm CurrencyExponentIndicatorType = 2
	Bam CurrencyExponentIndicatorType = 2
	Bbd CurrencyExponentIndicatorType = 2
	Bdt CurrencyExponentIndicatorType = 2
	Bgn CurrencyExponentIndicatorType = 2
	Bhd CurrencyExponentIndicatorType = 3
	Bif CurrencyExponentIndicatorType = 0
	Bmd CurrencyExponentIndicatorType = 2
	Bnd CurrencyExponentIndicatorType = 2
	Bob CurrencyExponentIndicatorType = 2
	Bov CurrencyExponentIndicatorType = 2
	Brl CurrencyExponentIndicatorType = 2
	Bsd CurrencyExponentIndicatorType = 2
	Btn CurrencyExponentIndicatorType = 2
	Bwp CurrencyExponentIndicatorType = 2
	Byr CurrencyExponentIndicatorType = 0
	Bzd CurrencyExponentIndicatorType = 2
	Cad CurrencyExponentIndicatorType = 2
	Cdf CurrencyExponentIndicatorType = 2
	Chf CurrencyExponentIndicatorType = 2
	Clf CurrencyExponentIndicatorType = 0
	Clp CurrencyExponentIndicatorType = 0
	Cny CurrencyExponentIndicatorType = 2
	Cop CurrencyExponentIndicatorType = 2
	Cou CurrencyExponentIndicatorType = 2
	Crc CurrencyExponentIndicatorType = 2
	Csd CurrencyExponentIndicatorType = 2
	Cup CurrencyExponentIndicatorType = 2
	Cve CurrencyExponentIndicatorType = 2
	Cyp CurrencyExponentIndicatorType = 2
	Czk CurrencyExponentIndicatorType = 2
	Djf CurrencyExponentIndicatorType = 0
	Dkk CurrencyExponentIndicatorType = 2
	Dop CurrencyExponentIndicatorType = 2
	Dzd CurrencyExponentIndicatorType = 2
	Eek CurrencyExponentIndicatorType = 2
	Egp CurrencyExponentIndicatorType = 2
	Ern CurrencyExponentIndicatorType = 2
	Etb CurrencyExponentIndicatorType = 2
	Eur CurrencyExponentIndicatorType = 2
	Fjd CurrencyExponentIndicatorType = 2
	Fkp CurrencyExponentIndicatorType = 2
	Gbp CurrencyExponentIndicatorType = 2
	Gel CurrencyExponentIndicatorType = 2
	Ghc CurrencyExponentIndicatorType = 2
	Gip CurrencyExponentIndicatorType = 2
	Gmd CurrencyExponentIndicatorType = 2
	Gnf CurrencyExponentIndicatorType = 0
	Gtq CurrencyExponentIndicatorType = 2
	Gwp CurrencyExponentIndicatorType = 2
	Gyd CurrencyExponentIndicatorType = 2
	Hkd CurrencyExponentIndicatorType = 2
	Hnl CurrencyExponentIndicatorType = 2
	Hrk CurrencyExponentIndicatorType = 2
	Htg CurrencyExponentIndicatorType = 2
	Huf CurrencyExponentIndicatorType = 2
	Idr CurrencyExponentIndicatorType = 2
	Ils CurrencyExponentIndicatorType = 2
	Inr CurrencyExponentIndicatorType = 2
	Iqd CurrencyExponentIndicatorType = 3
	Irr CurrencyExponentIndicatorType = 2
	Isk CurrencyExponentIndicatorType = 2
	Jmd CurrencyExponentIndicatorType = 2
	Jod CurrencyExponentIndicatorType = 3
	Jpy CurrencyExponentIndicatorType = 0
	Kes CurrencyExponentIndicatorType = 2
	Kgs CurrencyExponentIndicatorType = 2
	Khr CurrencyExponentIndicatorType = 2
	Kmf CurrencyExponentIndicatorType = 0
	Kpw CurrencyExponentIndicatorType = 2
	Krw CurrencyExponentIndicatorType = 0
	Kwd CurrencyExponentIndicatorType = 3
	Kyd CurrencyExponentIndicatorType = 2
	Kzt CurrencyExponentIndicatorType = 2
	Lak CurrencyExponentIndicatorType = 2
	Lbp CurrencyExponentIndicatorType = 2
	Lkr CurrencyExponentIndicatorType = 2
	Lrd CurrencyExponentIndicatorType = 2
	Lsl CurrencyExponentIndicatorType = 2
	Ltl CurrencyExponentIndicatorType = 2
	Lvl CurrencyExponentIndicatorType = 2
	Lyd CurrencyExponentIndicatorType = 3
	Mad CurrencyExponentIndicatorType = 2
	Mdl CurrencyExponentIndicatorType = 2
	Mga CurrencyExponentIndicatorType = 0
	Mgf CurrencyExponentIndicatorType = 0
	Mkd CurrencyExponentIndicatorType = 2
	Mmk CurrencyExponentIndicatorType = 2
	Mnt CurrencyExponentIndicatorType = 2
	Mop CurrencyExponentIndicatorType = 2
	Mro CurrencyExponentIndicatorType = 2
	Mtl CurrencyExponentIndicatorType = 2
	Mur CurrencyExponentIndicatorType = 2
	Mvr CurrencyExponentIndicatorType = 2
	Mwk CurrencyExponentIndicatorType = 2
	Mxn CurrencyExponentIndicatorType = 2
	Mxv CurrencyExponentIndicatorType = 2
	Myr CurrencyExponentIndicatorType = 2
	Mzm CurrencyExponentIndicatorType = 2
	Nad CurrencyExponentIndicatorType = 2
	Ngn CurrencyExponentIndicatorType = 2
	Nio CurrencyExponentIndicatorType = 2
	Nok CurrencyExponentIndicatorType = 2
	Npr CurrencyExponentIndicatorType = 2
	Nzd CurrencyExponentIndicatorType = 2
	Omr CurrencyExponentIndicatorType = 3
	Pab CurrencyExponentIndicatorType = 2
	Pen CurrencyExponentIndicatorType = 2
	Pgk CurrencyExponentIndicatorType = 2
	Php CurrencyExponentIndicatorType = 2
	Pkr CurrencyExponentIndicatorType = 2
	Pln CurrencyExponentIndicatorType = 2
	Pyg CurrencyExponentIndicatorType = 0
	Qar CurrencyExponentIndicatorType = 2
	Rol CurrencyExponentIndicatorType = 2
	Rub CurrencyExponentIndicatorType = 2
	Rur CurrencyExponentIndicatorType = 2
	Rwf CurrencyExponentIndicatorType = 0
	Sar CurrencyExponentIndicatorType = 2
	Sbd CurrencyExponentIndicatorType = 2
	Scr CurrencyExponentIndicatorType = 2
	Sdd CurrencyExponentIndicatorType = 2
	Sek CurrencyExponentIndicatorType = 2
	Sgd CurrencyExponentIndicatorType = 2
	Shp CurrencyExponentIndicatorType = 2
	Sit CurrencyExponentIndicatorType = 2
	Skk CurrencyExponentIndicatorType = 2
	Sll CurrencyExponentIndicatorType = 2
	Sos CurrencyExponentIndicatorType = 2
	Srd CurrencyExponentIndicatorType = 2
	Std CurrencyExponentIndicatorType = 2
	Svc CurrencyExponentIndicatorType = 2
	Syp CurrencyExponentIndicatorType = 2
	Szl CurrencyExponentIndicatorType = 2
	Thb CurrencyExponentIndicatorType = 2
	Tjs CurrencyExponentIndicatorType = 2
	Tmm CurrencyExponentIndicatorType = 2
	Tnd CurrencyExponentIndicatorType = 3
	Top CurrencyExponentIndicatorType = 2
	Trl CurrencyExponentIndicatorType = 0
	Ttd CurrencyExponentIndicatorType = 2
	Twd CurrencyExponentIndicatorType = 2
	Tzs CurrencyExponentIndicatorType = 2
	Uah CurrencyExponentIndicatorType = 2
	Ugx CurrencyExponentIndicatorType = 2
	Usd CurrencyExponentIndicatorType = 2
	Usn CurrencyExponentIndicatorType = 2
	Uss CurrencyExponentIndicatorType = 2
	Uyu CurrencyExponentIndicatorType = 2
	Uzs CurrencyExponentIndicatorType = 2
	Veb CurrencyExponentIndicatorType = 2
	Vnd CurrencyExponentIndicatorType = 2
	Vuv CurrencyExponentIndicatorType = 0
	Wst CurrencyExponentIndicatorType = 2
	Xaf CurrencyExponentIndicatorType = 0
	Xcd CurrencyExponentIndicatorType = 2
	Xof CurrencyExponentIndicatorType = 0
	Xpf CurrencyExponentIndicatorType = 0
	Yer CurrencyExponentIndicatorType = 2
	Zar CurrencyExponentIndicatorType = 2
	Zmk CurrencyExponentIndicatorType = 2
	Zwd CurrencyExponentIndicatorType = 2
)

type Cvv2ResultCodeType rune

const (
	SpaceOrEmptyCvv2Result     Cvv2ResultCodeType = ' '
	Cvv2Match                  Cvv2ResultCodeType = 'M'
	Cvv2NoMatch                Cvv2ResultCodeType = 'N'
	NotProcessed               Cvv2ResultCodeType = 'P'
	Cvv2NotPresentOnCard       Cvv2ResultCodeType = 'S'
	IssuerNotVisaCvv2Certified Cvv2ResultCodeType = 'U'
)

func (x Cvv2ResultCodeType) String() string {
	switch x {
	case SpaceOrEmptyCvv2Result:
		return "SpaceOrEmptyCvv2Result"
	case Cvv2Match:
		return "Cvv2Match"
	case Cvv2NoMatch:
		return "Cvv2NoMatch"
	case NotProcessed:
		return "NotProcessed"
	case Cvv2NotPresentOnCard:
		return "Cvv2NotPresentOnCard"
	case IssuerNotVisaCvv2Certified:
		return "IssuerNotVisaCvv2Certified"
	}
	return ""
}

type DeviceCodeType rune

const (
	UnknownOrUnsure                   DeviceCodeType = 'O'
	Pc                                DeviceCodeType = 'C'
	DialTerminal                      DeviceCodeType = 'D'
	ElectronicCashRegister            DeviceCodeType = 'E'
	InStorePromotion                  DeviceCodeType = 'I'
	MainFrame                         DeviceCodeType = 'M'
	PosPort                           DeviceCodeType = 'P'
	ReservedForThirdPartyDevelopers   DeviceCodeType = 'Q'
	PosPortR                          DeviceCodeType = 'R'
	PosPartner                        DeviceCodeType = 'S'
	SuppressPs2000MeritResponseFields DeviceCodeType = 'Z'
)

func (x DeviceCodeType) String() string {
	switch x {
	case UnknownOrUnsure:
		return "UnknownOrUnsure"
	case Pc:
		return "Pc"
	case DialTerminal:
		return "DialTerminal"
	case ElectronicCashRegister:
		return "ElectronicCashRegister"
	case InStorePromotion:
		return "InStorePromotion"
	case MainFrame:
		return "MainFrame"
	case PosPort:
		return "PosPort"
	case ReservedForThirdPartyDevelopers:
		return "ReservedForThirdPartyDevelopers"
	case PosPortR:
		return "PosPortR"
	case PosPartner:
		return "PosPartner"
	case SuppressPs2000MeritResponseFields:
		return "SuppressPs2000MeritResponseFields"
	}
	return ""
}

type EcommerceGoodsIndicatorType rune

const (
	DigitalGoods  EcommerceGoodsIndicatorType = 'D'
	PhysicalGoods EcommerceGoodsIndicatorType = 'P'
)

func (x EcommerceGoodsIndicatorType) String() string {
	switch x {
	case DigitalGoods:
		return "DigitalGoods"
	case PhysicalGoods:
		return "PhysicalGoods"
	}
	return ""
}

type ErrorCodeType rune

const (
	SpaceOrEmptyErrorCode          ErrorCodeType = ' '
	BlockedTerminal                ErrorCodeType = 'B'
	CardTypeError                  ErrorCodeType = 'C'
	DeviceError                    ErrorCodeType = 'D'
	ErrorInBatch                   ErrorCodeType = 'E'
	VitalResidencyRequirementError ErrorCodeType = 'P'
	SequenceError                  ErrorCodeType = 'S'
	TransmissionError              ErrorCodeType = 'T'
	UnknownError                   ErrorCodeType = 'U'
	RoutingError                   ErrorCodeType = 'V'
)

func (x ErrorCodeType) String() string {
	switch x {
	case SpaceOrEmptyErrorCode:
		return "SpaceOrEmptyErrorCode"
	case BlockedTerminal:
		return "BlockedTerminal"
	case CardTypeError:
		return "CardTypeError"
	case DeviceError:
		return "DeviceError"
	case ErrorInBatch:
		return "ErrorInBatch"
	case VitalResidencyRequirementError:
		return "VitalResidencyRequirementError"
	case SequenceError:
		return "SequenceError"
	case TransmissionError:
		return "TransmissionError"
	case UnknownError:
		return "UnknownError"
	case RoutingError:
		return "RoutingError"
	}
	return ""
}

type ErrorRecordType rune

const (
	SpaceOrEmptyErrorRecord                ErrorRecordType = ' '
	HeaderRecordError                      ErrorRecordType = 'H'
	ParameterRecordError                   ErrorRecordType = 'P'
	DetailRecordError                      ErrorRecordType = 'D'
	LineItemDetailRecordError              ErrorRecordType = 'L'
	TrailerRecordError                     ErrorRecordType = 'T'
	HierarchyError                         ErrorRecordType = 'K'
	TripLegDetailRecordCommercialCardError ErrorRecordType = 'M'
	TrailerResponseRecordError             ErrorRecordType = 'R'
)

func (x ErrorRecordType) String() string {
	switch x {
	case SpaceOrEmptyErrorRecord:
		return "SpaceOrEmptyErrorRecord"
	case HeaderRecordError:
		return "HeaderRecordError"
	case ParameterRecordError:
		return "ParameterRecordError"
	case DetailRecordError:
		return "DetailRecordError"
	case LineItemDetailRecordError:
		return "LineItemDetailRecordError"
	case TrailerRecordError:
		return "TrailerRecordError"
	case HierarchyError:
		return "HierarchyError"
	case TripLegDetailRecordCommercialCardError:
		return "TripLegDetailRecordCommercialCardError"
	case TrailerResponseRecordError:
		return "TrailerResponseRecordError"
	}
	return ""
}

type ExistingDebtIndicatorType rune

const (
	SpaceOrEmptyExistingDebt ExistingDebtIndicatorType = ' '
	ExistingDebtTransaction  ExistingDebtIndicatorType = '9'
)

func (x ExistingDebtIndicatorType) String() string {
	switch x {
	case SpaceOrEmptyExistingDebt:
		return "SpaceOrEmptyExistingDebt"
	case ExistingDebtTransaction:
		return "ExistingDebtTransaction"
	}
	return ""
}

type AutoRentalExtraChargeType rune

const (
	NoExtraCharge            AutoRentalExtraChargeType = '0'
	NoExtraChargeSpace       AutoRentalExtraChargeType = ' '
	Gasoline                 AutoRentalExtraChargeType = '1'
	ExtraMileage             AutoRentalExtraChargeType = '2'
	LateReturn               AutoRentalExtraChargeType = '3'
	OneWayServiceFee         AutoRentalExtraChargeType = '4'
	ParkingOrMovingViolation AutoRentalExtraChargeType = '5'
)

func (x AutoRentalExtraChargeType) String() string {
	switch x {
	case NoExtraCharge:
		return "NoExtraCharge"
	case NoExtraChargeSpace:
		return "NoExtraChargeSpace"
	case Gasoline:
		return "Gasoline"
	case ExtraMileage:
		return "ExtraMileage"
	case LateReturn:
		return "LateReturn"
	case OneWayServiceFee:
		return "OneWayServiceFee"
	case ParkingOrMovingViolation:
		return "ParkingOrMovingViolation"
	}
	return ""
}

type GroupMapBitType uint

const (
	//Table 3.2 Transaction header record - optional data groups

	HeaderRecordBit_PassengerTravel                  GroupMapBitType = 1
	HeaderRecordBit_HotelAutoRentalAndCardNotPresent GroupMapBitType = 2
	HeaderRecordBit_Gen2AuthenticationGenkey         GroupMapBitType = 4 //OnlyforGen2Authenticationparticipants
	HeaderRecordBit_DeveloperID                      GroupMapBitType = 5
	HeaderRecordBit_GatewayID                        GroupMapBitType = 6

	//Table 3.6 Transaction detail record - optional data groups
	DetailRecordBit_CashBack                                          GroupMapBitType = 1
	DetailRecordBit_Restaurant                                        GroupMapBitType = 2
	DetailRecordBit_DirectMarketing                                   GroupMapBitType = 3
	DetailRecordBit_AutoRental                                        GroupMapBitType = 4
	DetailRecordBit_Hotel                                             GroupMapBitType = 5
	DetailRecordBit_HotelNotFrontDeskNonLodgingPurchase               GroupMapBitType = 6
	DetailRecordBit_HotelAmericanExpress                              GroupMapBitType = 7
	DetailRecordBit_NonTECommercialCardLevelIIVisaMasterCardOnly      GroupMapBitType = 8
	DetailRecordBit_PassengerTransportVisaOnly                        GroupMapBitType = 9
	DetailRecordBit_DirectDebit                                       GroupMapBitType = 10
	DetailRecordBit_DirectMarketingInstallmentOrRecurringPayments     GroupMapBitType = 12
	DetailRecordBit_AutoRentalReturn                                  GroupMapBitType = 13
	DetailRecordBit_ServiceDevelopmentIndicator                       GroupMapBitType = 14
	DetailRecordBit_ExistingDebtIndicatorVisaOnly                     GroupMapBitType = 16
	DetailRecordBit_UniversalCardholderAuthentication                 GroupMapBitType = 17
	DetailRecordBit_ElectronicCommerceGoodsIndicator                  GroupMapBitType = 18
	DetailRecordBit_CCSPrivateLabel                                   GroupMapBitType = 19 //RestrictedforCCSUseOnly MandatoryforPrivateLabeltransactions
	DetailRecordBit_MerchantVerificationValue                         GroupMapBitType = 20
	DetailRecordBit_AmericanExpressCorporatePurchasingCard            GroupMapBitType = 21
	DetailRecordBit_USOnlyNonTECommercialCardLevelIIVisaAndMasterCard GroupMapBitType = 22
	DetailRecordBit_GroupMapExtension                                 GroupMapBitType = 23 //Groups 25-48
	DetailRecordBit_ReservedPOSPartnerUseOnly                         GroupMapBitType = 24

	//EXTENDEDBITMAPSECTION

	DetailRecordBit_CommercialCardLevelIINonTEFleetNonFuel       GroupMapBitType = 25
	DetailRecordBit_VisaCardEnhancedDataNonTEFleetNonFuel        GroupMapBitType = 26
	DetailRecordBit_MasterCardEnhancedDataNonTEFleetNonFuel      GroupMapBitType = 27
	DetailRecordBit_VisaCardEnhancedDataFleet                    GroupMapBitType = 28
	DetailRecordBit_MasterCardEnhancedDataFleetFuelOnly          GroupMapBitType = 29
	DetailRecordBit_AmericanExpressCPCChargeDescriptors          GroupMapBitType = 30
	DetailRecordBit_MarketSpecificDataIndicator                  GroupMapBitType = 31
	DetailRecordBit_POSDataCode                                  GroupMapBitType = 32
	DetailRecordBit_RetailTransactionAdviceAddenda               GroupMapBitType = 33
	DetailRecordBit_LodgingTransactionAdviceAddenda              GroupMapBitType = 34
	DetailRecordBit_TransactionAdviceLineItemCounts              GroupMapBitType = 35
	DetailRecordBit_LocationDetailTransactionAdviceAddenda       GroupMapBitType = 36
	DetailRecordBit_MasterCardMiscellaneousFields                GroupMapBitType = 37
	DetailRecordBit_VisaMiscellaneousFields                      GroupMapBitType = 38
	DetailRecordBit_AmexCAPNCorporatePurchasingSolutionExtension GroupMapBitType = 39
	DetailRecordBit_DiscoverPayPalMiscellaneousFields            GroupMapBitType = 40
	DetailRecordBit_DetailExtension                              GroupMapBitType = 41
	DetailRecordBit_AmexAutoRental                               GroupMapBitType = 42
	DetailRecordBit_AmexInsurance                                GroupMapBitType = 43
	DetailRecordBit_ACHTSYSInternalUseOnly                       GroupMapBitType = 44
	DetailRecordBit_TransactionSecurityIndicator                 GroupMapBitType = 45
	DetailRecordBit_MerchantSoftDescriptor                       GroupMapBitType = 46
)

func (x GroupMapBitType) String() string {
	switch x {
	case DetailRecordBit_CashBack:
		return "DetailRecordBit_CashBack"
	case DetailRecordBit_Restaurant:
		return "DetailRecordBit_Restaurant"
	case DetailRecordBit_DirectMarketing:
		return "DetailRecordBit_DirectMarketing"
	case DetailRecordBit_AutoRental:
		return "DetailRecordBit_AutoRental"
	case DetailRecordBit_Hotel:
		return "DetailRecordBit_Hotel"
	case DetailRecordBit_HotelNotFrontDeskNonLodgingPurchase:
		return "DetailRecordBit_HotelNotFrontDeskNonLodgingPurchase"
	case DetailRecordBit_HotelAmericanExpress:
		return "DetailRecordBit_HotelAmericanExpress"
	case DetailRecordBit_NonTECommercialCardLevelIIVisaMasterCardOnly:
		return "DetailRecordBit_NonTECommercialCardLevelIIVisaMasterCardOnly"
	case DetailRecordBit_PassengerTransportVisaOnly:
		return "DetailRecordBit_PassengerTransportVisaOnly"
	case DetailRecordBit_DirectDebit:
		return "DetailRecordBit_DirectDebit"
	case DetailRecordBit_DirectMarketingInstallmentOrRecurringPayments:
		return "DetailRecordBit_DirectMarketingInstallmentOrRecurringPayments"
	case DetailRecordBit_AutoRentalReturn:
		return "DetailRecordBit_AutoRentalReturn"
	case DetailRecordBit_ServiceDevelopmentIndicator:
		return "DetailRecordBit_ServiceDevelopmentIndicator"
	case DetailRecordBit_ExistingDebtIndicatorVisaOnly:
		return "DetailRecordBit_ExistingDebtIndicatorVisaOnly"
	case DetailRecordBit_UniversalCardholderAuthentication:
		return "DetailRecordBit_UniversalCardholderAuthentication"
	case DetailRecordBit_ElectronicCommerceGoodsIndicator:
		return "DetailRecordBit_ElectronicCommerceGoodsIndicator"
	case DetailRecordBit_CCSPrivateLabel:
		return "DetailRecordBit_CCSPrivateLabel"
	case DetailRecordBit_MerchantVerificationValue:
		return "DetailRecordBit_MerchantVerificationValue"
	case DetailRecordBit_AmericanExpressCorporatePurchasingCard:
		return "DetailRecordBit_AmericanExpressCorporatePurchasingCard"
	case DetailRecordBit_USOnlyNonTECommercialCardLevelIIVisaAndMasterCard:
		return "DetailRecordBit_USOnlyNonTECommercialCardLevelIIVisaAndMasterCard"
	case DetailRecordBit_GroupMapExtension:
		return "DetailRecordBit_GroupMapExtension"
	case DetailRecordBit_ReservedPOSPartnerUseOnly:
		return "DetailRecordBit_ReservedPOSPartnerUseOnly"
	case DetailRecordBit_CommercialCardLevelIINonTEFleetNonFuel:
		return "DetailRecordBit_CommercialCardLevelIINonTEFleetNonFuel"
	case DetailRecordBit_VisaCardEnhancedDataNonTEFleetNonFuel:
		return "DetailRecordBit_VisaCardEnhancedDataNonTEFleetNonFuel"
	case DetailRecordBit_MasterCardEnhancedDataNonTEFleetNonFuel:
		return "DetailRecordBit_MasterCardEnhancedDataNonTEFleetNonFuel"
	case DetailRecordBit_VisaCardEnhancedDataFleet:
		return "DetailRecordBit_VisaCardEnhancedDataFleet"
	case DetailRecordBit_MasterCardEnhancedDataFleetFuelOnly:
		return "DetailRecordBit_MasterCardEnhancedDataFleetFuelOnly"
	case DetailRecordBit_AmericanExpressCPCChargeDescriptors:
		return "DetailRecordBit_AmericanExpressCPCChargeDescriptors"
	case DetailRecordBit_MarketSpecificDataIndicator:
		return "DetailRecordBit_MarketSpecificDataIndicator"
	case DetailRecordBit_POSDataCode:
		return "DetailRecordBit_POSDataCode"
	case DetailRecordBit_RetailTransactionAdviceAddenda:
		return "DetailRecordBit_RetailTransactionAdviceAddenda"
	case DetailRecordBit_LodgingTransactionAdviceAddenda:
		return "DetailRecordBit_LodgingTransactionAdviceAddenda"
	case DetailRecordBit_TransactionAdviceLineItemCounts:
		return "DetailRecordBit_TransactionAdviceLineItemCounts"
	case DetailRecordBit_LocationDetailTransactionAdviceAddenda:
		return "DetailRecordBit_LocationDetailTransactionAdviceAddenda"
	case DetailRecordBit_MasterCardMiscellaneousFields:
		return "DetailRecordBit_MasterCardMiscellaneousFields"
	case DetailRecordBit_VisaMiscellaneousFields:
		return "DetailRecordBit_VisaMiscellaneousFields"
	case DetailRecordBit_AmexCAPNCorporatePurchasingSolutionExtension:
		return "DetailRecordBit_AmexCAPNCorporatePurchasingSolutionExtension"
	case DetailRecordBit_DiscoverPayPalMiscellaneousFields:
		return "DetailRecordBit_DiscoverPayPalMiscellaneousFields"
	case DetailRecordBit_DetailExtension:
		return "DetailRecordBit_DetailExtension"
	case DetailRecordBit_AmexAutoRental:
		return "DetailRecordBit_AmexAutoRental"
	case DetailRecordBit_AmexInsurance:
		return "DetailRecordBit_AmexInsurance"
	case DetailRecordBit_ACHTSYSInternalUseOnly:
		return "DetailRecordBit_ACHTSYSInternalUseOnly"
	case DetailRecordBit_TransactionSecurityIndicator:
		return "DetailRecordBit_TransactionSecurityIndicator"
	case DetailRecordBit_MerchantSoftDescriptor:
		return "DetailRecordBit_MerchantSoftDescriptor"
	}
	return ""
}

type HotelExtraChargeType rune

const (
	PositionNotUsed          HotelExtraChargeType = '0'
	PositionNotUsedSpace     HotelExtraChargeType = ' '
	HotelExtraChargeReserved HotelExtraChargeType = '1'
	Restaurant               HotelExtraChargeType = '2'
	GiftShop                 HotelExtraChargeType = '3'
	MiniBar                  HotelExtraChargeType = '4'
	Telephone                HotelExtraChargeType = '5'
	HotelExtraChargeOther    HotelExtraChargeType = '6'
	Laundry                  HotelExtraChargeType = '7'
)

func (x HotelExtraChargeType) String() string {
	switch x {
	case PositionNotUsed:
		return "PositionNotUsed"
	case PositionNotUsedSpace:
		return "PositionNotUsedSpace"
	case HotelExtraChargeReserved:
		return "HotelExtraChargeReserved"
	case Restaurant:
		return "Restaurant"
	case GiftShop:
		return "GiftShop"
	case MiniBar:
		return "MiniBar"
	case Telephone:
		return "Telephone"
	case HotelExtraChargeOther:
		return "HotelExtraChargeOther"
	case Laundry:
		return "Laundry"
	}
	return ""
}

type FuelPurchaseType rune

const (
	FuelPurchase           FuelPurchaseType = '1'
	NonFuelPurchase        FuelPurchaseType = '2'
	FuelAndNonFuelPurchase FuelPurchaseType = '3'
)

func (x FuelPurchaseType) String() string {
	switch x {
	case FuelPurchase:
		return "FuelPurchase"
	case NonFuelPurchase:
		return "NonFuelPurchase"
	case FuelAndNonFuelPurchase:
		return "FuelAndNonFuelPurchase"
	}
	return ""
}

type FuelType uint

const (
	FuelOther                               FuelType = 0
	UnleadedRegular86                       FuelType = 1
	UnleadedRegular87                       FuelType = 2
	UnleadedMidGrade88                      FuelType = 3
	UnleadedMidGrade89                      FuelType = 4
	UnleadedPremium90                       FuelType = 5
	UnleadedPremium91                       FuelType = 6
	UnleadedSuper92                         FuelType = 7
	UnleadedSuper93                         FuelType = 8
	UnleadedSuper94                         FuelType = 9
	RegularLeaded                           FuelType = 11
	Diesel                                  FuelType = 12
	DieselPremium                           FuelType = 13
	Kerosene                                FuelType = 14
	Lpg                                     FuelType = 15
	Gasohol                                 FuelType = 16
	Cng                                     FuelType = 17
	Methanol85                              FuelType = 18
	Methanol10                              FuelType = 19
	Methanol7                               FuelType = 20
	Methanol5                               FuelType = 21
	Ethanol85                               FuelType = 22
	Ethanol10                               FuelType = 23
	Ethanol7                                FuelType = 24
	Ethanol5                                FuelType = 25
	JetFuel                                 FuelType = 26
	AviationFuel                            FuelType = 27
	OffRoadDiesel                           FuelType = 28
	Marine                                  FuelType = 29
	MotorOil                                FuelType = 30
	OilChange                               FuelType = 31
	EngineService                           FuelType = 32
	TransmissionService                     FuelType = 33
	BrakeService                            FuelType = 34
	UnassignedRepairValuesA                 FuelType = 35
	UnassignedRepairValuesB                 FuelType = 36
	UnassignedRepairValuesC                 FuelType = 37
	UnassignedRepairValuesD                 FuelType = 38
	MiscellaneousRepairs                    FuelType = 39
	TiresBatteriesAccessories               FuelType = 40
	Tires                                   FuelType = 41
	Batteries                               FuelType = 42
	AutomotiveAccessories                   FuelType = 43
	AutomotiveGlass                         FuelType = 44
	CarWash                                 FuelType = 45
	UnassignedAutomotiveProductsOrServicesA FuelType = 46
	UnassignedAutomotiveProductsOrServicesB FuelType = 47
	UnassignedAutomotiveProductsOrServicesC FuelType = 48
	UnassignedAutomotiveProductsOrServicesD FuelType = 49
	UnassignedAutomotiveProductsOrServicesE FuelType = 50
	UnassignedAutomotiveProductsOrServicesF FuelType = 51
	UnassignedAutomotiveProductsOrServicesG FuelType = 52
	UnassignedAutomotiveProductsOrServicesH FuelType = 53
	UnassignedAutomotiveProductsOrServicesI FuelType = 54
	UnassignedAutomotiveProductsOrServicesJ FuelType = 55
	UnassignedAutomotiveProductsOrServicesK FuelType = 56
	UnassignedAutomotiveProductsOrServicesL FuelType = 57
	UnassignedAutomotiveProductsOrServicesM FuelType = 58
	UnassignedAutomotiveProductsOrServicesN FuelType = 59
	UnassignedAutomotiveProductsOrServicesO FuelType = 60
	UnassignedAutomotiveProductsOrServicesP FuelType = 61
	UnassignedAutomotiveProductsOrServicesQ FuelType = 62
	UnassignedAutomotiveProductsOrServicesR FuelType = 63
	UnassignedAutomotiveProductsOrServicesS FuelType = 64
	UnassignedAutomotiveProductsOrServicesT FuelType = 65
	UnassignedAutomotiveProductsOrServicesU FuelType = 66
	UnassignedAutomotiveProductsOrServicesV FuelType = 67
	UnassignedAutomotiveProductsOrServicesX FuelType = 68
	UnassignedAutomotiveProductsOrServicesY FuelType = 69
	CigarettesOrTobacco                     FuelType = 70
	UnassignedFoodOrGroceryItemsA           FuelType = 71
	UnassignedFoodOrGroceryItemsB           FuelType = 72
	UnassignedFoodOrGroceryItemsC           FuelType = 73
	UnassignedFoodOrGroceryItemsD           FuelType = 74
	UnassignedFoodOrGroceryItemsE           FuelType = 75
	UnassignedFoodOrGroceryItemsF           FuelType = 76
	UnassignedFoodOrGroceryItemsG           FuelType = 77
	HealthAndBeautyAid                      FuelType = 78
	MiscellaneousGrocery                    FuelType = 79
	Soda                                    FuelType = 80
	BeerAndWine                             FuelType = 81
	MilkAndJuice                            FuelType = 82
	UnassignedBeverageItemsA                FuelType = 83
	UnassignedBeverageItemsB                FuelType = 84
	UnassignedBeverageItemsC                FuelType = 85
	UnassignedBeverageItemsD                FuelType = 86
	UnassignedBeverageItemsE                FuelType = 87
	UnassignedBeverageItemsF                FuelType = 88
	UnassignedBeverageItemsG                FuelType = 89
	Miscellaneous                           FuelType = 90
)

func (x FuelType) String() string {
	switch x {
	case FuelOther:
		return "FuelOther"
	case UnleadedRegular86:
		return "UnleadedRegular86"
	case UnleadedRegular87:
		return "UnleadedRegular87"
	case UnleadedMidGrade88:
		return "UnleadedMidGrade88"
	case UnleadedMidGrade89:
		return "UnleadedMidGrade89"
	case UnleadedPremium90:
		return "UnleadedPremium90"
	case UnleadedPremium91:
		return "UnleadedPremium91"
	case UnleadedSuper92:
		return "UnleadedSuper92"
	case UnleadedSuper93:
		return "UnleadedSuper93"
	case UnleadedSuper94:
		return "UnleadedSuper94"
	case RegularLeaded:
		return "RegularLeaded"
	case Diesel:
		return "Diesel"
	case DieselPremium:
		return "DieselPremium"
	case Kerosene:
		return "Kerosene"
	case Lpg:
		return "Lpg"
	case Gasohol:
		return "Gasohol"
	case Cng:
		return "Cng"
	case Methanol85:
		return "Methanol85"
	case Methanol10:
		return "Methanol10"
	case Methanol7:
		return "Methanol7"
	case Methanol5:
		return "Methanol5"
	case Ethanol85:
		return "Ethanol85"
	case Ethanol10:
		return "Ethanol10"
	case Ethanol7:
		return "Ethanol7"
	case Ethanol5:
		return "Ethanol5"
	case JetFuel:
		return "JetFuel"
	case AviationFuel:
		return "AviationFuel"
	case OffRoadDiesel:
		return "OffRoadDiesel"
	case Marine:
		return "Marine"
	case MotorOil:
		return "MotorOil"
	case OilChange:
		return "OilChange"
	case EngineService:
		return "EngineService"
	case TransmissionService:
		return "TransmissionService"
	case BrakeService:
		return "BrakeService"
	case UnassignedRepairValuesA:
		return "UnassignedRepairValuesA"
	case UnassignedRepairValuesB:
		return "UnassignedRepairValuesB"
	case UnassignedRepairValuesC:
		return "UnassignedRepairValuesC"
	case UnassignedRepairValuesD:
		return "UnassignedRepairValuesD"
	case MiscellaneousRepairs:
		return "MiscellaneousRepairs"
	case TiresBatteriesAccessories:
		return "TiresBatteriesAccessories"
	case Tires:
		return "Tires"
	case Batteries:
		return "Batteries"
	case AutomotiveAccessories:
		return "AutomotiveAccessories"
	case AutomotiveGlass:
		return "AutomotiveGlass"
	case CarWash:
		return "CarWash"
	case UnassignedAutomotiveProductsOrServicesA:
		return "UnassignedAutomotiveProductsOrServicesA"
	case UnassignedAutomotiveProductsOrServicesB:
		return "UnassignedAutomotiveProductsOrServicesB"
	case UnassignedAutomotiveProductsOrServicesC:
		return "UnassignedAutomotiveProductsOrServicesC"
	case UnassignedAutomotiveProductsOrServicesD:
		return "UnassignedAutomotiveProductsOrServicesD"
	case UnassignedAutomotiveProductsOrServicesE:
		return "UnassignedAutomotiveProductsOrServicesE"
	case UnassignedAutomotiveProductsOrServicesF:
		return "UnassignedAutomotiveProductsOrServicesF"
	case UnassignedAutomotiveProductsOrServicesG:
		return "UnassignedAutomotiveProductsOrServicesG"
	case UnassignedAutomotiveProductsOrServicesH:
		return "UnassignedAutomotiveProductsOrServicesH"
	case UnassignedAutomotiveProductsOrServicesI:
		return "UnassignedAutomotiveProductsOrServicesI"
	case UnassignedAutomotiveProductsOrServicesJ:
		return "UnassignedAutomotiveProductsOrServicesJ"
	case UnassignedAutomotiveProductsOrServicesK:
		return "UnassignedAutomotiveProductsOrServicesK"
	case UnassignedAutomotiveProductsOrServicesL:
		return "UnassignedAutomotiveProductsOrServicesL"
	case UnassignedAutomotiveProductsOrServicesM:
		return "UnassignedAutomotiveProductsOrServicesM"
	case UnassignedAutomotiveProductsOrServicesN:
		return "UnassignedAutomotiveProductsOrServicesN"
	case UnassignedAutomotiveProductsOrServicesO:
		return "UnassignedAutomotiveProductsOrServicesO"
	case UnassignedAutomotiveProductsOrServicesP:
		return "UnassignedAutomotiveProductsOrServicesP"
	case UnassignedAutomotiveProductsOrServicesQ:
		return "UnassignedAutomotiveProductsOrServicesQ"
	case UnassignedAutomotiveProductsOrServicesR:
		return "UnassignedAutomotiveProductsOrServicesR"
	case UnassignedAutomotiveProductsOrServicesS:
		return "UnassignedAutomotiveProductsOrServicesS"
	case UnassignedAutomotiveProductsOrServicesT:
		return "UnassignedAutomotiveProductsOrServicesT"
	case UnassignedAutomotiveProductsOrServicesU:
		return "UnassignedAutomotiveProductsOrServicesU"
	case UnassignedAutomotiveProductsOrServicesV:
		return "UnassignedAutomotiveProductsOrServicesV"
	case UnassignedAutomotiveProductsOrServicesX:
		return "UnassignedAutomotiveProductsOrServicesX"
	case UnassignedAutomotiveProductsOrServicesY:
		return "UnassignedAutomotiveProductsOrServicesY"
	case CigarettesOrTobacco:
		return "CigarettesOrTobacco"
	case UnassignedFoodOrGroceryItemsA:
		return "UnassignedFoodOrGroceryItemsA"
	case UnassignedFoodOrGroceryItemsB:
		return "UnassignedFoodOrGroceryItemsB"
	case UnassignedFoodOrGroceryItemsC:
		return "UnassignedFoodOrGroceryItemsC"
	case UnassignedFoodOrGroceryItemsD:
		return "UnassignedFoodOrGroceryItemsD"
	case UnassignedFoodOrGroceryItemsE:
		return "UnassignedFoodOrGroceryItemsE"
	case UnassignedFoodOrGroceryItemsF:
		return "UnassignedFoodOrGroceryItemsF"
	case UnassignedFoodOrGroceryItemsG:
		return "UnassignedFoodOrGroceryItemsG"
	case HealthAndBeautyAid:
		return "HealthAndBeautyAid"
	case MiscellaneousGrocery:
		return "MiscellaneousGrocery"
	case Soda:
		return "Soda"
	case BeerAndWine:
		return "BeerAndWine"
	case MilkAndJuice:
		return "MilkAndJuice"
	case UnassignedBeverageItemsA:
		return "UnassignedBeverageItemsA"
	case UnassignedBeverageItemsB:
		return "UnassignedBeverageItemsB"
	case UnassignedBeverageItemsC:
		return "UnassignedBeverageItemsC"
	case UnassignedBeverageItemsD:
		return "UnassignedBeverageItemsD"
	case UnassignedBeverageItemsE:
		return "UnassignedBeverageItemsE"
	case UnassignedBeverageItemsF:
		return "UnassignedBeverageItemsF"
	case UnassignedBeverageItemsG:
		return "UnassignedBeverageItemsG"
	case Miscellaneous:
		return "Miscellaneous"
	}
	return ""
}

type G3VersionType uint

const (
	NoAddendumData                                    G3VersionType = 0
	G3v001CommercialCard                              G3VersionType = 1
	ReservedA                                         G3VersionType = 2
	ReservedB                                         G3VersionType = 3
	ReservedC                                         G3VersionType = 4
	ReservedD                                         G3VersionType = 5
	ReservedE                                         G3VersionType = 6
	G3v007CardVerificationCode                        G3VersionType = 7
	G3v008FleetFuelingCard                            G3VersionType = 8
	G3v009SeteCommerce                                G3VersionType = 9
	G3v010Ccps                                        G3VersionType = 10
	G3v011ChipCondition                               G3VersionType = 11
	G3v012CommercialCardLargeTicket                   G3VersionType = 12
	G3v013ElectronicBenefitsTransfer                  G3VersionType = 13
	G3v014MotoEcommerce                               G3VersionType = 14
	G3v015ServiceDevelopmentIndicator                 G3VersionType = 15
	G3v016PosCheckServiceAuthorization                G3VersionType = 16
	G3v017Visa3dSecurEcomVerification                 G3VersionType = 17
	G3v018ExistingDebtIndicator                       G3VersionType = 18
	G3v019MastercardUniversalCardholderAuthentication G3VersionType = 19
	G3v020DeveloperInformation                        G3VersionType = 20
	G3v021MerchantVerificationValue                   G3VersionType = 21
	G3v022AdditionalAmounts                           G3VersionType = 22
	G3v023VisaMastercardHealthcare                    G3VersionType = 23
	G3v024MerchantAdviceCode                          G3VersionType = 24
	G3v025TransactionFeeAmount                        G3VersionType = 25
	G3v026ProductParticipationGroup                   G3VersionType = 26
	G3v027PosData                                     G3VersionType = 27
	G3v028AmericanExpressAdditionalData               G3VersionType = 28
	G3v029ExtendedAVSdata                             G3VersionType = 29
	G3v030AmericanExpressMerchantNameLocation         G3VersionType = 30
	G3v031AgentIdentificationService                  G3VersionType = 31
	G3v032CurrencyConversionData                      G3VersionType = 32
	G3v033ReversalRequestCode                         G3VersionType = 33
	G3v034CardProductCode                             G3VersionType = 34
	G3v035PromotionalCode                             G3VersionType = 35
	G3v036PaymentTransactionIdentifier                G3VersionType = 36
	G3v037RealTimeSubstantiation                      G3VersionType = 37
	G3v038ElectroMagneticSignature                    G3VersionType = 38
	G3v039CardholderVerificationMethod                G3VersionType = 39
	G3v040VisaISAChargeIndicator                      G3VersionType = 40
	G3v041NtiaUpcSkuData                              G3VersionType = 41
	G3v042VisaContactless                             G3VersionType = 42
	G3v043NetworkID                                   G3VersionType = 43
	G3v044AutomatedTellerMachine                      G3VersionType = 44
	G3v045IntegratedChipCard                          G3VersionType = 45
	G3v046CardTypeGroup                               G3VersionType = 46
	G3v047TSYSInternalUseOnly                         G3VersionType = 47
	G3v048AmexCardholderVerificationResults           G3VersionType = 48
	G3v049Gen2TerminalAuthentication                  G3VersionType = 49
	G3v050AssociationTimestamp                        G3VersionType = 50
	G3v051MasterCardEventMonitoringRealTimeScoring    G3VersionType = 51
	G3v052VoltageEncryptionTransmissionBlock          G3VersionType = 52
	G3v053Token                                       G3VersionType = 53
	G3v054TransitProgram                              G3VersionType = 54
	G3v055IntegratedChipCardEmvTlv                    G3VersionType = 55
	G3v056MessageReasonCode                           G3VersionType = 56
	G3v057DiscoverOrPayPalAdditionalData              G3VersionType = 57
	G3v058AlternateAccountID                          G3VersionType = 58
	G3v059MasterCardPayPassMappingService             G3VersionType = 59
	G3v060PayPassMobile                               G3VersionType = 60
	G3v061SpendQualifiedIndicator                     G3VersionType = 61
	G3v200GiftCard                                    G3VersionType = 200
)

func (x G3VersionType) String() string {
	switch x {
	case NoAddendumData:
		return "NoAddendumData"
	case G3v001CommercialCard:
		return "G3v001CommercialCard"
	case ReservedA:
		return "ReservedA"
	case ReservedB:
		return "ReservedB"
	case ReservedC:
		return "ReservedC"
	case ReservedD:
		return "ReservedD"
	case ReservedE:
		return "ReservedE"
	case G3v007CardVerificationCode:
		return "G3v007CardVerificationCode"
	case G3v008FleetFuelingCard:
		return "G3v008FleetFuelingCard"
	case G3v009SeteCommerce:
		return "G3v009SeteCommerce"
	case G3v010Ccps:
		return "G3v010Ccps"
	case G3v011ChipCondition:
		return "G3v011ChipCondition"
	case G3v012CommercialCardLargeTicket:
		return "G3v012CommercialCardLargeTicket"
	case G3v013ElectronicBenefitsTransfer:
		return "G3v013ElectronicBenefitsTransfer"
	case G3v014MotoEcommerce:
		return "G3v014MotoEcommerce"
	case G3v015ServiceDevelopmentIndicator:
		return "G3v015ServiceDevelopmentIndicator"
	case G3v016PosCheckServiceAuthorization:
		return "G3v016PosCheckServiceAuthorization"
	case G3v017Visa3dSecurEcomVerification:
		return "G3v017Visa3dSecurEcomVerification"
	case G3v018ExistingDebtIndicator:
		return "G3v018ExistingDebtIndicator"
	case G3v019MastercardUniversalCardholderAuthentication:
		return "G3v019MastercardUniversalCardholderAuthentication"
	case G3v020DeveloperInformation:
		return "G3v020DeveloperInformation"
	case G3v021MerchantVerificationValue:
		return "G3v021MerchantVerificationValue"
	case G3v022AdditionalAmounts:
		return "G3v022AdditionalAmounts"
	case G3v023VisaMastercardHealthcare:
		return "G3v023VisaMastercardHealthcare"
	case G3v024MerchantAdviceCode:
		return "G3v024MerchantAdviceCode"
	case G3v025TransactionFeeAmount:
		return "G3v025TransactionFeeAmount"
	case G3v026ProductParticipationGroup:
		return "G3v026ProductParticipationGroup"
	case G3v027PosData:
		return "G3v027PosData"
	case G3v028AmericanExpressAdditionalData:
		return "G3v028AmericanExpressAdditionalData"
	case G3v029ExtendedAVSdata:
		return "G3v029ExtendedAVSdata"
	case G3v030AmericanExpressMerchantNameLocation:
		return "G3v030AmericanExpressMerchantNameLocation"
	case G3v031AgentIdentificationService:
		return "G3v031AgentIdentificationService"
	case G3v032CurrencyConversionData:
		return "G3v032CurrencyConversionData"
	case G3v033ReversalRequestCode:
		return "G3v033ReversalRequestCode"
	case G3v034CardProductCode:
		return "G3v034CardProductCode"
	case G3v035PromotionalCode:
		return "G3v035PromotionalCode"
	case G3v036PaymentTransactionIdentifier:
		return "G3v036PaymentTransactionIdentifier"
	case G3v037RealTimeSubstantiation:
		return "G3v037RealTimeSubstantiation"
	case G3v038ElectroMagneticSignature:
		return "G3v038ElectroMagneticSignature"
	case G3v039CardholderVerificationMethod:
		return "G3v039CardholderVerificationMethod"
	case G3v040VisaISAChargeIndicator:
		return "G3v040VisaISAChargeIndicator"
	case G3v041NtiaUpcSkuData:
		return "G3v041NtiaUpcSkuData"
	case G3v042VisaContactless:
		return "G3v042VisaContactless"
	case G3v043NetworkID:
		return "G3v043NetworkID"
	case G3v044AutomatedTellerMachine:
		return "G3v044AutomatedTellerMachine"
	case G3v045IntegratedChipCard:
		return "G3v045IntegratedChipCard"
	case G3v046CardTypeGroup:
		return "G3v046CardTypeGroup"
	case G3v047TSYSInternalUseOnly:
		return "G3v047TSYSInternalUseOnly"
	case G3v048AmexCardholderVerificationResults:
		return "G3v048AmexCardholderVerificationResults"
	case G3v049Gen2TerminalAuthentication:
		return "G3v049Gen2TerminalAuthentication"
	case G3v050AssociationTimestamp:
		return "G3v050AssociationTimestamp"
	case G3v051MasterCardEventMonitoringRealTimeScoring:
		return "G3v051MasterCardEventMonitoringRealTimeScoring"
	case G3v052VoltageEncryptionTransmissionBlock:
		return "G3v052VoltageEncryptionTransmissionBlock"
	case G3v053Token:
		return "G3v053Token"
	case G3v054TransitProgram:
		return "G3v054TransitProgram"
	case G3v055IntegratedChipCardEmvTlv:
		return "G3v055IntegratedChipCardEmvTlv"
	case G3v056MessageReasonCode:
		return "G3v056MessageReasonCode"
	case G3v057DiscoverOrPayPalAdditionalData:
		return "G3v057DiscoverOrPayPalAdditionalData"
	case G3v058AlternateAccountID:
		return "G3v058AlternateAccountID"
	case G3v059MasterCardPayPassMappingService:
		return "G3v059MasterCardPayPassMappingService"
	case G3v060PayPassMobile:
		return "G3v060PayPassMobile"
	case G3v061SpendQualifiedIndicator:
		return "G3v061SpendQualifiedIndicator"
	case G3v200GiftCard:
		return "G3v200GiftCard"
	}
	return ""
}

type HotelChargeType uint

const (
	SpaceOrEmptyHotelCharge HotelChargeType = 0
	HotelChargeHotel        HotelChargeType = 1
	HotelChargeRestaurant   HotelChargeType = 2
	HotelChargeGiftShop     HotelChargeType = 3
)

func (x HotelChargeType) String() string {
	switch x {
	case SpaceOrEmptyHotelCharge:
		return "SpaceOrEmptyHotelCharge"
	case HotelChargeHotel:
		return "HotelChargeHotel"
	case HotelChargeRestaurant:
		return "HotelChargeRestaurant"
	case HotelChargeGiftShop:
		return "HotelChargeGiftShop"
	}
	return ""
}

type IdNumberPromptType uint

const (
	IdNumberAndOdometerReading IdNumberPromptType = 1
	VehicleNumber              IdNumberPromptType = 2
	DriverIdAndOdometerReading IdNumberPromptType = 3
	OdometerReading            IdNumberPromptType = 4
	NoPrompting                IdNumberPromptType = 5
	IdNumber                   IdNumberPromptType = 6
)

func (x IdNumberPromptType) String() string {
	switch x {
	case IdNumberAndOdometerReading:
		return "IdNumberAndOdometerReading"
	case VehicleNumber:
		return "VehicleNumber"
	case DriverIdAndOdometerReading:
		return "DriverIdAndOdometerReading"
	case OdometerReading:
		return "OdometerReading"
	case NoPrompting:
		return "NoPrompting"
	case IdNumber:
		return "IdNumber"
	}
	return ""
}

type IndustryCodeType rune

const (
	AllTypesExceptHotelAndAuto       IndustryCodeType = ' '
	IndustryCodeUnknownOrUnsure      IndustryCodeType = '0'
	AutoRental                       IndustryCodeType = 'A'
	BankFinancialInstitution         IndustryCodeType = 'B'
	DirectMarketing                  IndustryCodeType = 'D'
	FoodRestaurant                   IndustryCodeType = 'F'
	GroceryStoreSupermarket          IndustryCodeType = 'G'
	HotelAndLodging                  IndustryCodeType = 'H'
	LimitedAmountTerminal            IndustryCodeType = 'L'
	OilCompanyAutomatedFuelingSystem IndustryCodeType = 'O'
	PassengerTransport               IndustryCodeType = 'P'
	Retail                           IndustryCodeType = 'R'
	Internet                         IndustryCodeType = 'I'
)

func (x IndustryCodeType) String() string {
	switch x {
	case AllTypesExceptHotelAndAuto:
		return "AllTypesExceptHotelAndAuto"
	case IndustryCodeUnknownOrUnsure:
		return "IndustryCodeUnknownOrUnsure"
	case AutoRental:
		return "AutoRental"
	case BankFinancialInstitution:
		return "BankFinancialInstitution"
	case DirectMarketing:
		return "DirectMarketing"
	case FoodRestaurant:
		return "FoodRestaurant"
	case GroceryStoreSupermarket:
		return "GroceryStoreSupermarket"
	case HotelAndLodging:
		return "HotelAndLodging"
	case LimitedAmountTerminal:
		return "LimitedAmountTerminal"
	case OilCompanyAutomatedFuelingSystem:
		return "OilCompanyAutomatedFuelingSystem"
	case PassengerTransport:
		return "PassengerTransport"
	case Retail:
		return "Retail"
	case Internet:
		return "Internet"
	}
	return ""
}

type InternetIndicatorType rune

const (
	InternetIndicatorNo  InternetIndicatorType = 'N'
	InternetIndicatorYes InternetIndicatorType = 'Y'
)

func (x InternetIndicatorType) String() string {
	switch x {
	case InternetIndicatorNo:
		return "InternetIndicatorNo"
	case InternetIndicatorYes:
		return "InternetIndicatorYes"
	}
	return ""
}

type LanguageIndicatorType uint

const (
	English    LanguageIndicatorType = 0
	Spanish    LanguageIndicatorType = 1
	Portuguese LanguageIndicatorType = 2
	Irish      LanguageIndicatorType = 3
	French     LanguageIndicatorType = 4
	German     LanguageIndicatorType = 5
	Italian    LanguageIndicatorType = 6
	Dutch      LanguageIndicatorType = 7
)

func (x LanguageIndicatorType) String() string {
	switch x {
	case English:
		return "English"
	case Spanish:
		return "Spanish"
	case Portuguese:
		return "Portuguese"
	case Irish:
		return "Irish"
	case French:
		return "French"
	case German:
		return "German"
	case Italian:
		return "Italian"
	case Dutch:
		return "Dutch"
	}
	return ""
}

type Level3ServiceType rune

const (
	NonFuelTransaction Level3ServiceType = ' '
	FullService        Level3ServiceType = 'F'
	SelfService        Level3ServiceType = 'S'
)

func (x Level3ServiceType) String() string {
	switch x {
	case NonFuelTransaction:
		return "NonFuelTransaction"
	case FullService:
		return "FullService"
	case SelfService:
		return "SelfService"
	}
	return ""
}

type LocalTaxIndicatorType uint

const (
	SalesTaxNotProvided   LocalTaxIndicatorType = 0
	LocalOrSalesTaxAmount LocalTaxIndicatorType = 1
	TaxExempt             LocalTaxIndicatorType = 2
)

func (x LocalTaxIndicatorType) String() string {
	switch x {
	case SalesTaxNotProvided:
		return "SalesTaxNotProvided"
	case LocalOrSalesTaxAmount:
		return "LocalOrSalesTaxAmount"
	case TaxExempt:
		return "TaxExempt"
	}
	return ""
}

type MarketSpecificDataType rune

const (
	MarketSpecificAutoRental                      MarketSpecificDataType = 'A'
	MarketSpecificBillPayment                     MarketSpecificDataType = 'B'
	MarketSpecificEcommerceTransactionAggregation MarketSpecificDataType = 'E'
	HotelLodging                                  MarketSpecificDataType = 'H'
	B2BInvoicePayments                            MarketSpecificDataType = 'J'
	MarketSpecificHealthCare                      MarketSpecificDataType = 'M'
	MarketSpecificFailedMarket                    MarketSpecificDataType = 'N'
	MarketSpecificTransit                         MarketSpecificDataType = 'T'
	OtherIndustries                               MarketSpecificDataType = ' '
)

func (x MarketSpecificDataType) String() string {
	switch x {
	case MarketSpecificAutoRental:
		return "MarketSpecificAutoRental"
	case HotelLodging:
		return "HotelLodging"
	case OtherIndustries:
		return "OtherIndustries"
	}
	return ""
}

type MerchantCategoryType uint

const (
	SpaceOrEmptyMerchantCategory                                                                        MerchantCategoryType = 0
	GreenbriahResorts                                                                                   MerchantCategoryType = 3753
	AmeliaIslandPlanation                                                                               MerchantCategoryType = 3754
	TheHomestead                                                                                        MerchantCategoryType = 3755
	SouthSeasResorts                                                                                    MerchantCategoryType = 3756
	PassengerTransportationRailroadsFeriesLocal                                                         MerchantCategoryType = 4111
	PassengerRailways                                                                                   MerchantCategoryType = 4112
	AmbulanceServices                                                                                   MerchantCategoryType = 4119
	TaxicabsAndLimousines                                                                               MerchantCategoryType = 4121
	BusLinesIncludingChartersTourBuses                                                                  MerchantCategoryType = 4131
	TourBuses                                                                                           MerchantCategoryType = 4131
	DeliveryServicesLocal                                                                               MerchantCategoryType = 4214
	FreightCarriers                                                                                     MerchantCategoryType = 4214
	MotorFreightCarriers                                                                                MerchantCategoryType = 4214
	MovingAndStorageCompanies                                                                           MerchantCategoryType = 4214
	TruckingLocallongDistance                                                                           MerchantCategoryType = 4214
	CourierServicesAirOrGround                                                                          MerchantCategoryType = 4215
	FrightForwarders                                                                                    MerchantCategoryType = 4215
	Storage                                                                                             MerchantCategoryType = 4225
	WarehousingPublic                                                                                   MerchantCategoryType = 4225
	CruiseLines                                                                                         MerchantCategoryType = 4411
	SteamshipLines                                                                                      MerchantCategoryType = 4411
	BoatRentalsAndLeases                                                                                MerchantCategoryType = 4457
	MarinasMarineServiceAndSupplies                                                                     MerchantCategoryType = 4468
	AirlinesAirCarriersNotListedElsewhere                                                               MerchantCategoryType = 4511
	AirportsAirportTerminals                                                                            MerchantCategoryType = 4582
	FlyingFields                                                                                        MerchantCategoryType = 4582
	TravelAgenciesAndTourOperations                                                                     MerchantCategoryType = 4722
	PackageTourOperatorsForUseInGermanyOnly                                                             MerchantCategoryType = 4723
	TollAndBridgeFees                                                                                   MerchantCategoryType = 4784
	TransportationServicesNotElsewhereClassified                                                        MerchantCategoryType = 4789
	TelecommunicationsEquipmentIncludingTelephoneSales                                                  MerchantCategoryType = 4812
	FaxServices                                                                                         MerchantCategoryType = 4814
	TelecommunicationService                                                                            MerchantCategoryType = 4814
	Visaphone                                                                                           MerchantCategoryType = 4815
	TelegraphServices                                                                                   MerchantCategoryType = 4821
	MoneyOrdersWireTransfer                                                                             MerchantCategoryType = 4829
	CableAndOtherPayTelevisionPreviouslyCableServices                                                   MerchantCategoryType = 4899
	ElectricGasSanitaryAndWaterUtilities                                                                MerchantCategoryType = 4900
	MotorVehicleSuppliesAndNewParts                                                                     MerchantCategoryType = 5013
	OfficeAndCommercialFurniture                                                                        MerchantCategoryType = 5021
	ConstructionMaterialsNotElsewhereClassified                                                         MerchantCategoryType = 5039
	OfficePhotographicPhotocopyAndMicrofilmEquipment                                                    MerchantCategoryType = 5044
	ComputersComputerPeripheralEquipmentSoftware                                                        MerchantCategoryType = 5045
	CommercialEquipmentNotElsewhereClassified                                                           MerchantCategoryType = 5046
	MedicalDentalOphthalmicHospitalEquipmentAndSupplies                                                 MerchantCategoryType = 5047
	MetalServiceCentersAndOffices                                                                       MerchantCategoryType = 5051
	ElectricalPartsAndEquipment                                                                         MerchantCategoryType = 5065
	HardwareEquipmentAndSupplies                                                                        MerchantCategoryType = 5072
	PlumbingAndHeatingEquipmentAndSupplies                                                              MerchantCategoryType = 5074
	IndustrialSuppliesNotElsewhereClassified                                                            MerchantCategoryType = 5085
	PreciousStonesAndMetalsWatchesAndJewelry                                                            MerchantCategoryType = 5094
	DurableGoodsNotElsewhereClassified                                                                  MerchantCategoryType = 5099
	StationeryOfficeSuppliesPrintingAndWritingPaper                                                     MerchantCategoryType = 5111
	DrugsDrugProprietorsAndDruggistsSundries                                                            MerchantCategoryType = 5122
	PieceGoodsNotionsAndOtherDryGoods                                                                   MerchantCategoryType = 5131
	MensWomensAndChildrensUniformsAndCommercialClothing                                                 MerchantCategoryType = 5137
	CommercialFootwear                                                                                  MerchantCategoryType = 5139
	ChemicalsAndAlliedProductsNotElsewhereClassified                                                    MerchantCategoryType = 5169
	PetroleumAndPetroleumProducts                                                                       MerchantCategoryType = 5172
	BooksPeriodicalsAndNewspapers                                                                       MerchantCategoryType = 5192
	FloristsSuppliesNurseryStockAndFlowers                                                              MerchantCategoryType = 5193
	PaintsVarnishesAndSupplies                                                                          MerchantCategoryType = 5198
	NonDurableGoodsNotElsewhereClassified                                                               MerchantCategoryType = 5199
	HomeSupplyWarehouseStores                                                                           MerchantCategoryType = 5200
	LumberAndBuildingMaterialsStores                                                                    MerchantCategoryType = 5211
	GlassStores                                                                                         MerchantCategoryType = 5231
	PaintAndWallpaperStores                                                                             MerchantCategoryType = 5231
	WallpaperStores                                                                                     MerchantCategoryType = 5231
	HardwareStores                                                                                      MerchantCategoryType = 5251
	NurseriesLawnAndGardenSupplyStore                                                                   MerchantCategoryType = 5261
	MobileHomeDealers                                                                                   MerchantCategoryType = 5271
	WholesaleClubs                                                                                      MerchantCategoryType = 5300
	DutyFreeStore                                                                                       MerchantCategoryType = 5309
	DiscountStores                                                                                      MerchantCategoryType = 5310
	DepartmentStores                                                                                    MerchantCategoryType = 5311
	VarietyStores                                                                                       MerchantCategoryType = 5331
	MiscGeneralMerchandise                                                                              MerchantCategoryType = 5399
	GroceryStores                                                                                       MerchantCategoryType = 5411
	Supermarkets                                                                                        MerchantCategoryType = 5411
	FreezerAndLockerMeatProvisioners                                                                    MerchantCategoryType = 5422
	MeatProvisionersFreezerAndLocker                                                                    MerchantCategoryType = 5422
	CandyStores                                                                                         MerchantCategoryType = 5441
	ConfectioneryStores                                                                                 MerchantCategoryType = 5441
	NutStores                                                                                           MerchantCategoryType = 5441
	DairyProductsStores                                                                                 MerchantCategoryType = 5451
	Bakeries                                                                                            MerchantCategoryType = 5462
	MiscFoodStoresConvenienceStoresAndSpecialtyMarkets                                                  MerchantCategoryType = 5499
	CarAndTruckDealersNewAndUsedSalesServiceRepairsPartsAndLeasing                                      MerchantCategoryType = 5511
	AutomobileAndTruckDealersUsedOnly                                                                   MerchantCategoryType = 5521
	AutomobileSupplyStores                                                                              MerchantCategoryType = 5531
	AutomotiveTireStores                                                                                MerchantCategoryType = 5532
	AutomotivePartsAccessoriesStores                                                                    MerchantCategoryType = 5533
	ServiceStationsWithOrWithoutAncillaryServices                                                       MerchantCategoryType = 5541
	AutomatedFuelDispensers                                                                             MerchantCategoryType = 5542
	BoatDealers                                                                                         MerchantCategoryType = 5551
	RecreationalAndUtilityTrailersCampDealers                                                           MerchantCategoryType = 5561
	MotorcycleDealers                                                                                   MerchantCategoryType = 5571
	MotorHomeDealers                                                                                    MerchantCategoryType = 5592
	SnowmobileDealers                                                                                   MerchantCategoryType = 5598
	MensAndBoysClothingAndAccessoriesStores                                                             MerchantCategoryType = 5611
	WomensReadyToWearStores                                                                             MerchantCategoryType = 5621
	WomensAccessoryAndSpecialtyShops                                                                    MerchantCategoryType = 5631
	ChildrensAndInfantsWearStores                                                                       MerchantCategoryType = 5641
	FamilyClothingStores                                                                                MerchantCategoryType = 5651
	SportsApparelRidingApparelStores                                                                    MerchantCategoryType = 5655
	ShoeStores                                                                                          MerchantCategoryType = 5661
	FurriersAndFurShops                                                                                 MerchantCategoryType = 5681
	MensAndWomensClothingStores                                                                         MerchantCategoryType = 5691
	TailorsSeamstressMendingAndAlterations                                                              MerchantCategoryType = 5697
	WigAndToupeeStores                                                                                  MerchantCategoryType = 5698
	MiscellaneousApparelAndAccessoryShops                                                               MerchantCategoryType = 5699
	FurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances                                          MerchantCategoryType = 5712
	FloorCoveringStores                                                                                 MerchantCategoryType = 5713
	DraperyWindowCoveringAndUpholsteryStores                                                            MerchantCategoryType = 5714
	FireplaceScreensAndAccessoriesStores                                                                MerchantCategoryType = 5718
	MiscellaneousHomeFurnishingSpecialtyStores                                                          MerchantCategoryType = 5719
	HouseholdApplianceStores                                                                            MerchantCategoryType = 5722
	ElectronicSales                                                                                     MerchantCategoryType = 5732
	MusicStoresMusicalInstrumentsPianoSheetMusic                                                        MerchantCategoryType = 5733
	ComputerSoftwareStores                                                                              MerchantCategoryType = 5734
	RecordShops                                                                                         MerchantCategoryType = 5735
	Caterers                                                                                            MerchantCategoryType = 5811
	EatingPlacesAndRestaurants                                                                          MerchantCategoryType = 5812
	AlcoholicBeveragesBarsTavernsCocktailLoungesNightClubsAndDiscotheques                               MerchantCategoryType = 5813
	FastFoodRestaurants                                                                                 MerchantCategoryType = 5814
	DrugStoresAndPharmacies                                                                             MerchantCategoryType = 5912
	PackageStoresBeerWineAndLiquor                                                                      MerchantCategoryType = 5921
	UsedMerchandiseAndSecondhandStores                                                                  MerchantCategoryType = 5931
	AntiqueShopsSalesRepairsAndRestorationServices                                                      MerchantCategoryType = 5832
	PawnShopsAndSalvageYards                                                                            MerchantCategoryType = 5933
	WreckingAndSalvageYards                                                                             MerchantCategoryType = 5935
	AntiqueReproductions                                                                                MerchantCategoryType = 5937
	BicycleShopsSalesAndService                                                                         MerchantCategoryType = 5940
	SportingGoodsStores                                                                                 MerchantCategoryType = 5941
	BookStores                                                                                          MerchantCategoryType = 5942
	StationeryStoresOfficeAndSchoolSupplyStores                                                         MerchantCategoryType = 5943
	WatchClockJewelryAndSilverwareStores                                                                MerchantCategoryType = 5944
	HobbyToyAndGameShops                                                                                MerchantCategoryType = 5945
	CameraAndPhotographicSupplyStores                                                                   MerchantCategoryType = 5946
	CardShopsGiftNoveltyAndSouvenirShops                                                                MerchantCategoryType = 5947
	LeatherFoodsStores                                                                                  MerchantCategoryType = 5948
	SewingNeedleFabricAndPriceGoodsStores                                                               MerchantCategoryType = 5949
	GlasswarecrystalStores                                                                              MerchantCategoryType = 5950
	DirectMarketingInsuranceService                                                                     MerchantCategoryType = 5960
	MailOrderHousesIncludingCatalogOrderStoresBookrecordClubsNoLongerPermittedForUsOriginalPresentments MerchantCategoryType = 5961
	DirectMarketingTravelRelatedArrangementsServices                                                    MerchantCategoryType = 5962
	DoorToDoorSales                                                                                     MerchantCategoryType = 5963
	DirectMarketingCatalogMerchant                                                                      MerchantCategoryType = 5964
	DirectMarketingCatalogAndCatalogAndRetailMerchant                                                   MerchantCategoryType = 5965
	DirectMarketingOutboundTelemarketingMerchant                                                        MerchantCategoryType = 5966
	DirectMarketingInboundTeleservicesMerchant                                                          MerchantCategoryType = 5967
	DirectMarketingContinuitysubscriptionMerchant                                                       MerchantCategoryType = 5968
	DirectMarketingNotElsewhereClassified                                                               MerchantCategoryType = 5969
	ArtistsSupplyAndCraftShops                                                                          MerchantCategoryType = 5970
	ArtDealersAndGalleries                                                                              MerchantCategoryType = 5971
	StampAndCoinStoresPhilatelicAndNumismaticSupplies                                                   MerchantCategoryType = 5972
	ReligiousGoodsStores                                                                                MerchantCategoryType = 5973
	HearingAidsSalesServiceAndSupplyStores                                                              MerchantCategoryType = 5975
	OrthopedicGoodsProstheticDevices                                                                    MerchantCategoryType = 5976
	CosmeticStores                                                                                      MerchantCategoryType = 5977
	TypewriterStoresSalesRentalService                                                                  MerchantCategoryType = 5978
	FuelFuelOilWoodCoalLiquefiedPetroleum                                                               MerchantCategoryType = 5983
	Florists                                                                                            MerchantCategoryType = 5992
	CigarStoresAndStands                                                                                MerchantCategoryType = 5993
	NewsDealersAndNewsstands                                                                            MerchantCategoryType = 5994
	PetShopsPetFoodsAndSuppliesStores                                                                   MerchantCategoryType = 5995
	SwimmingPoolsSalesServiceAndSupplies                                                                MerchantCategoryType = 5996
	ElectricRazorStoresSalesAndService                                                                  MerchantCategoryType = 5997
	TentAndAwningShops                                                                                  MerchantCategoryType = 5998
	MiscellaneousAndSpecialtyRetailStores                                                               MerchantCategoryType = 5999
	FinancialInstitutionsManualCashDisbursements                                                        MerchantCategoryType = 6010
	FinancialInstitutionsAutomatedCash                                                                  MerchantCategoryType = 6011
	FinancialInstitutionsMerchandiseAndServices                                                         MerchantCategoryType = 6012
	NonFinancialInstitutionsForeignCurrencyMoneyOrdersNotWireTransferAndTravelersCheques                MerchantCategoryType = 6051
	SecurityBrokersdealers                                                                              MerchantCategoryType = 6211
	InsuranceSalesUnderwritingAndPremiums                                                               MerchantCategoryType = 6300
	InsurancePremiumsNoLongerValidForFirstPresentmentWork                                               MerchantCategoryType = 6381
	InsuranceNotElsewhereClassifiedNoLongerValidForFirstPresentmentWork                                 MerchantCategoryType = 6399
	LodgingHotelsMotelsResortsCentralReservationServicesNotElsewhereClassified                          MerchantCategoryType = 7011
	Timeshares                                                                                          MerchantCategoryType = 7012
	SportingAndRecreationalCamps                                                                        MerchantCategoryType = 7032
	TrailerParksAndCampGrounds                                                                          MerchantCategoryType = 7033
	LaundryCleaningAndGarmentServices                                                                   MerchantCategoryType = 7210
	LaundryFamilyAndCommercial                                                                          MerchantCategoryType = 7211
	DryCleaners                                                                                         MerchantCategoryType = 7216
	CarpetAndUpholsteryCleaning                                                                         MerchantCategoryType = 7217
	PhotographicStudios                                                                                 MerchantCategoryType = 7221
	BarberAndBeautyShops                                                                                MerchantCategoryType = 7230
	ShopRepairShopsAndShoeShineParlorsAndHatCleaningShops                                               MerchantCategoryType = 7251
	FuneralServiceAndCrematories                                                                        MerchantCategoryType = 7261
	DatingAndEscortServices                                                                             MerchantCategoryType = 7273
	TaxPreparationService                                                                               MerchantCategoryType = 7276
	CounselingServiceDebtMarriagePersonal                                                               MerchantCategoryType = 7277
	BuyingshoppingServicesClubs                                                                         MerchantCategoryType = 7278
	ClothingRentalCostumesFormalWearUniforms                                                            MerchantCategoryType = 7296
	MassageParlors                                                                                      MerchantCategoryType = 7297
	HealthAndBeautyShops                                                                                MerchantCategoryType = 7298
	MiscellaneousPersonalServicesNotElsewhereClassifies                                                 MerchantCategoryType = 7299
	AdvertisingServices                                                                                 MerchantCategoryType = 7311
	ConsumerCreditReportingAgencies                                                                     MerchantCategoryType = 7321
	BlueprintingAndPhotocopyingServices                                                                 MerchantCategoryType = 7332
	CommercialPhotographyArtAndGraphics                                                                 MerchantCategoryType = 7333
	QuickCopyReproductionAndBlueprintingServices                                                        MerchantCategoryType = 7338
	StenographicAndSecretarialSupportServices                                                           MerchantCategoryType = 7339
	DisinfectingServices                                                                                MerchantCategoryType = 7342
	ExterminatingAndDisinfectingServices                                                                MerchantCategoryType = 7342
	CleaningAndMaintenanceJanitorialServices                                                            MerchantCategoryType = 7349
	EmploymentAgenciesTemporaryHelpServices                                                             MerchantCategoryType = 7361
	ComputerProgrammingIntegratedSystemsDesignAndDataProcessingServices                                 MerchantCategoryType = 7372
	InformationRetrievalServices                                                                        MerchantCategoryType = 7375
	ComputerMaintenanceAndRepairServicesNotElsewhereClassified                                          MerchantCategoryType = 7379
	ManagementConsultingAndPublicRelationsServices                                                      MerchantCategoryType = 7392
	ProtectiveAndSecurityServicesIncludingArmoredCarsAndGuardDogs                                       MerchantCategoryType = 7393
	EquipmentRentalAndLeasingServicesToolRentalFurnitureRentalAndApplianceRental                        MerchantCategoryType = 7394
	PhotofinishingLaboratoriesPhotoDeveloping                                                           MerchantCategoryType = 7395
	BusinessServicesNotElsewhereClassified                                                              MerchantCategoryType = 7399
	CarRentalCompaniesNotListedBelow                                                                    MerchantCategoryType = 7512
	TruckAndUtilityTrailerRentals                                                                       MerchantCategoryType = 7513
	MotorHomeAndRecreationalVehicleRentals                                                              MerchantCategoryType = 7519
	AutomobileParkingLotsAndGarages                                                                     MerchantCategoryType = 7523
	AutomotiveBodyRepairShops                                                                           MerchantCategoryType = 7531
	TireReTreadingAndRepairShops                                                                        MerchantCategoryType = 7534
	PaintShopsAutomotive                                                                                MerchantCategoryType = 7535
	AutomotiveServiceShops                                                                              MerchantCategoryType = 7538
	CarWashes                                                                                           MerchantCategoryType = 7542
	TowingServices                                                                                      MerchantCategoryType = 7549
	RadioRepairShops                                                                                    MerchantCategoryType = 7622
	AirConditioningAndRefrigerationRepairShops                                                          MerchantCategoryType = 7623
	ElectricalAndSmallApplianceRepairShops                                                              MerchantCategoryType = 7629
	WatchClockAndJewelryRepair                                                                          MerchantCategoryType = 7631
	FurnitureFurnitureRepairAndFurnitureRefinishing                                                     MerchantCategoryType = 7641
	WeldingRepair                                                                                       MerchantCategoryType = 7692
	RepairShopsAndRelatedServicesMiscellaneous                                                          MerchantCategoryType = 7699
	MotionPicturesAndVideoTapeProductionAndDistribution                                                 MerchantCategoryType = 7829
	MotionPictureTheaters                                                                               MerchantCategoryType = 7832
	VideoTapeRentalStores                                                                               MerchantCategoryType = 7841
	DanceHallsStudiosAndSchools                                                                         MerchantCategoryType = 7911
	TheatricalProducersExceptMotionPicturesTicketAgencies                                               MerchantCategoryType = 7922
	BandsOrchestrasAndMiscellaneousEntertainersNotElsewhereClassified                                   MerchantCategoryType = 7929
	BilliardAndPoolEstablishments                                                                       MerchantCategoryType = 7932
	BowlingAlleys                                                                                       MerchantCategoryType = 7933
	CommercialSportsAthleticFieldsProfessionalSportClubsAndSportPromoters                               MerchantCategoryType = 7941
	TouristAttractionsAndExhibits                                                                       MerchantCategoryType = 7991
	GolfCoursesPublic                                                                                   MerchantCategoryType = 7992
	VideoAmusementGameSupplies                                                                          MerchantCategoryType = 7993
	VideoGameArcadesestablishments                                                                      MerchantCategoryType = 7994
	BettingIncludingLotteryTicketsCasinoGamingChipsOffTrackBettingAndWagers                             MerchantCategoryType = 7995
	AmusementParksCarnivalsCircusesFortuneTellers                                                       MerchantCategoryType = 7996
	MembershipClubsSportsRecreationAthleticCountryClubsAndPrivateGolfCourses                            MerchantCategoryType = 7997
	AquariumsSeaAquariumsDolphinariums                                                                  MerchantCategoryType = 7998
	RecreationServicesNotElsewhereClassified                                                            MerchantCategoryType = 7999
	DoctorsAndPhysiciansNotElsewhereClassified                                                          MerchantCategoryType = 8011
	DentistsAndOrthodontists                                                                            MerchantCategoryType = 8021
	Osteopaths                                                                                          MerchantCategoryType = 8031
	Chiropractors                                                                                       MerchantCategoryType = 8041
	OptometristsAndOphthalmologists                                                                     MerchantCategoryType = 8042
	OpticiansOpticiansGoodsAndEyeglasses                                                                MerchantCategoryType = 8043
	OpticiansOpticalGoodsAndEyeglassesNoLongerValidForFirstPresentments                                 MerchantCategoryType = 8044
	PodiatristsAndChiropodists                                                                          MerchantCategoryType = 8049
	NursingAndPersonalCareFacilities                                                                    MerchantCategoryType = 8050
	Hospitals                                                                                           MerchantCategoryType = 8062
	MedicalAndDentalLaboratories                                                                        MerchantCategoryType = 8071
	MedicalServicesAndHealthPractitionersNotElsewhereClassified                                         MerchantCategoryType = 8099
	LegalServicesAndAttorneys                                                                           MerchantCategoryType = 8111
	ElementaryAndSecondarySchools                                                                       MerchantCategoryType = 8211
	CollegesJuniorCollegesUniversitiesAndProfessionalSchools                                            MerchantCategoryType = 8220
	CorrespondenceSchools                                                                               MerchantCategoryType = 8241
	BusinessAndSecretarialSchools                                                                       MerchantCategoryType = 8244
	VocationalSchoolsAndTradeSchools                                                                    MerchantCategoryType = 8249
	SchoolsAndEducationalServicesNotElsewhereClassified                                                 MerchantCategoryType = 8299
	ChildCareServices                                                                                   MerchantCategoryType = 8351
	CharitableAndSocialServiceOrganizations                                                             MerchantCategoryType = 8398
	CivicFraternalAndSocialAssociations                                                                 MerchantCategoryType = 8641
	PoliticalOrganizations                                                                              MerchantCategoryType = 8651
	ReligiousOrganizations                                                                              MerchantCategoryType = 8661
	AutomobileAssociations                                                                              MerchantCategoryType = 8675
	MembershipOrganizationsNotElsewhereClassified                                                       MerchantCategoryType = 8699
	TestingLaboratoriesNonMedical                                                                       MerchantCategoryType = 8734
	ArchitecturalEngineeringAndSurveyingServices                                                        MerchantCategoryType = 8911
	AccountingAuditingAndBookkeepingServices                                                            MerchantCategoryType = 8931
	ProfessionalServicesNotElsewhereDefined                                                             MerchantCategoryType = 8999
	CourtCostsIncludingAlimonyAndChildSupport                                                           MerchantCategoryType = 9211
	Fines                                                                                               MerchantCategoryType = 9222
	BailAndBondPayments                                                                                 MerchantCategoryType = 9223
	TaxPayments                                                                                         MerchantCategoryType = 9311
	GovernmentServicesNotElsewhereClassified                                                            MerchantCategoryType = 9399
	PostalServicesGovernmentOnly                                                                        MerchantCategoryType = 9402
	IntraGovernmentTransactions                                                                         MerchantCategoryType = 9405
	AutomatedReferralServiceForVisaOnly                                                                 MerchantCategoryType = 9700
	VisaCredentialServiceForVisaOnly                                                                    MerchantCategoryType = 9701
	GcasEmergencyServicesForVisaOnly                                                                    MerchantCategoryType = 9702
	IntraCompanyPurchasesForVisaOnly                                                                    MerchantCategoryType = 9950
)

func (x MerchantCategoryType) String() string {
	switch x {
	case SpaceOrEmptyMerchantCategory:
		return "SpaceOrEmptyMerchantCategory"
	case GreenbriahResorts:
		return "GreenbriahResorts"
	case AmeliaIslandPlanation:
		return "AmeliaIslandPlanation"
	case TheHomestead:
		return "TheHomestead"
	case SouthSeasResorts:
		return "SouthSeasResorts"
	case PassengerTransportationRailroadsFeriesLocal:
		return "PassengerTransportationRailroadsFeriesLocal"
	case PassengerRailways:
		return "PassengerRailways"
	case AmbulanceServices:
		return "AmbulanceServices"
	case TaxicabsAndLimousines:
		return "TaxicabsAndLimousines"
	case BusLinesIncludingChartersTourBuses:
		return "BusLinesIncludingChartersTourBuses"
	// case TourBuses:
	// 	return "TourBuses"
	case DeliveryServicesLocal:
		return "DeliveryServicesLocal"
	// case FreightCarriers:
	// 	return "FreightCarriers"
	// case MotorFreightCarriers:
	// 	return "MotorFreightCarriers"
	// case MovingAndStorageCompanies:
	// 	return "MovingAndStorageCompanies"
	// case TruckingLocallongDistance:
	// 	return "TruckingLocallongDistance"
	// case CourierServicesAirOrGround:
	// 	return "CourierServicesAirOrGround"
	// case FrightForwarders:
	// 	return "FrightForwarders"
	// case Storage:
	// 	return "Storage"
	// case WarehousingPublic:
	// 	return "WarehousingPublic"
	// case CruiseLines:
	// 	return "CruiseLines"
	// case SteamshipLines:
	// 	return "SteamshipLines"
	// case BoatRentalsAndLeases:
	// 	return "BoatRentalsAndLeases"
	// case MarinasMarineServiceAndSupplies:
	// 	return "MarinasMarineServiceAndSupplies"
	// case AirlinesAirCarriersNotListedElsewhere:
	// 	return "AirlinesAirCarriersNotListedElsewhere"
	// case AirportsAirportTerminals:
	// 	return "AirportsAirportTerminals"
	// case FlyingFields:
	// 	return "FlyingFields"
	// case TravelAgenciesAndTourOperations:
	// 	return "TravelAgenciesAndTourOperations"
	// case PackageTourOperatorsForUseInGermanyOnly:
	// 	return "PackageTourOperatorsForUseInGermanyOnly"
	// case TollAndBridgeFees:
	// 	return "TollAndBridgeFees"
	// case TransportationServicesNotElsewhereClassified:
	// 	return "TransportationServicesNotElsewhereClassified"
	// case TelecommunicationsEquipmentIncludingTelephoneSales:
	// 	return "TelecommunicationsEquipmentIncludingTelephoneSales"
	// case FaxServices:
	// 	return "FaxServices"
	case TelecommunicationService:
		return "TelecommunicationService"
	case Visaphone:
		return "Visaphone"
	case TelegraphServices:
		return "TelegraphServices"
	case MoneyOrdersWireTransfer:
		return "MoneyOrdersWireTransfer"
	case CableAndOtherPayTelevisionPreviouslyCableServices:
		return "CableAndOtherPayTelevisionPreviouslyCableServices"
	case ElectricGasSanitaryAndWaterUtilities:
		return "ElectricGasSanitaryAndWaterUtilities"
	case MotorVehicleSuppliesAndNewParts:
		return "MotorVehicleSuppliesAndNewParts"
	case OfficeAndCommercialFurniture:
		return "OfficeAndCommercialFurniture"
	case ConstructionMaterialsNotElsewhereClassified:
		return "ConstructionMaterialsNotElsewhereClassified"
	case OfficePhotographicPhotocopyAndMicrofilmEquipment:
		return "OfficePhotographicPhotocopyAndMicrofilmEquipment"
	case ComputersComputerPeripheralEquipmentSoftware:
		return "ComputersComputerPeripheralEquipmentSoftware"
	case CommercialEquipmentNotElsewhereClassified:
		return "CommercialEquipmentNotElsewhereClassified"
	case MedicalDentalOphthalmicHospitalEquipmentAndSupplies:
		return "MedicalDentalOphthalmicHospitalEquipmentAndSupplies"
	case MetalServiceCentersAndOffices:
		return "MetalServiceCentersAndOffices"
	case ElectricalPartsAndEquipment:
		return "ElectricalPartsAndEquipment"
	case HardwareEquipmentAndSupplies:
		return "HardwareEquipmentAndSupplies"
	case PlumbingAndHeatingEquipmentAndSupplies:
		return "PlumbingAndHeatingEquipmentAndSupplies"
	case IndustrialSuppliesNotElsewhereClassified:
		return "IndustrialSuppliesNotElsewhereClassified"
	case PreciousStonesAndMetalsWatchesAndJewelry:
		return "PreciousStonesAndMetalsWatchesAndJewelry"
	case DurableGoodsNotElsewhereClassified:
		return "DurableGoodsNotElsewhereClassified"
	case StationeryOfficeSuppliesPrintingAndWritingPaper:
		return "StationeryOfficeSuppliesPrintingAndWritingPaper"
	case DrugsDrugProprietorsAndDruggistsSundries:
		return "DrugsDrugProprietorsAndDruggistsSundries"
	case PieceGoodsNotionsAndOtherDryGoods:
		return "PieceGoodsNotionsAndOtherDryGoods"
	case MensWomensAndChildrensUniformsAndCommercialClothing:
		return "MensWomensAndChildrensUniformsAndCommercialClothing"
	case CommercialFootwear:
		return "CommercialFootwear"
	case ChemicalsAndAlliedProductsNotElsewhereClassified:
		return "ChemicalsAndAlliedProductsNotElsewhereClassified"
	case PetroleumAndPetroleumProducts:
		return "PetroleumAndPetroleumProducts"
	case BooksPeriodicalsAndNewspapers:
		return "BooksPeriodicalsAndNewspapers"
	case FloristsSuppliesNurseryStockAndFlowers:
		return "FloristsSuppliesNurseryStockAndFlowers"
	case PaintsVarnishesAndSupplies:
		return "PaintsVarnishesAndSupplies"
	case NonDurableGoodsNotElsewhereClassified:
		return "NonDurableGoodsNotElsewhereClassified"
	case HomeSupplyWarehouseStores:
		return "HomeSupplyWarehouseStores"
	case LumberAndBuildingMaterialsStores:
		return "LumberAndBuildingMaterialsStores"
	case GlassStores:
		return "GlassStores"
	// case PaintAndWallpaperStores:
	// 	return "PaintAndWallpaperStores"
	// case WallpaperStores:
	// 	return "WallpaperStores"
	case HardwareStores:
		return "HardwareStores"
	case NurseriesLawnAndGardenSupplyStore:
		return "NurseriesLawnAndGardenSupplyStore"
	case MobileHomeDealers:
		return "MobileHomeDealers"
	case WholesaleClubs:
		return "WholesaleClubs"
	case DutyFreeStore:
		return "DutyFreeStore"
	case DiscountStores:
		return "DiscountStores"
	case DepartmentStores:
		return "DepartmentStores"
	case VarietyStores:
		return "VarietyStores"
	case MiscGeneralMerchandise:
		return "MiscGeneralMerchandise"
	case GroceryStores:
		return "GroceryStores"
	// case Supermarkets:
	// 	return "Supermarkets"
	// case FreezerAndLockerMeatProvisioners:
	// 	return "FreezerAndLockerMeatProvisioners"
	// case MeatProvisionersFreezerAndLocker:
	// 	return "MeatProvisionersFreezerAndLocker"
	// case CandyStores:
	// 	return "CandyStores"
	// case ConfectioneryStores:
	// 	return "ConfectioneryStores"
	// case NutStores:
	// 	return "NutStores"
	case DairyProductsStores:
		return "DairyProductsStores"
	case Bakeries:
		return "Bakeries"
	case MiscFoodStoresConvenienceStoresAndSpecialtyMarkets:
		return "MiscFoodStoresConvenienceStoresAndSpecialtyMarkets"
	case CarAndTruckDealersNewAndUsedSalesServiceRepairsPartsAndLeasing:
		return "CarAndTruckDealersNewAndUsedSalesServiceRepairsPartsAndLeasing"
	case AutomobileAndTruckDealersUsedOnly:
		return "AutomobileAndTruckDealersUsedOnly"
	case AutomobileSupplyStores:
		return "AutomobileSupplyStores"
	case AutomotiveTireStores:
		return "AutomotiveTireStores"
	case AutomotivePartsAccessoriesStores:
		return "AutomotivePartsAccessoriesStores"
	case ServiceStationsWithOrWithoutAncillaryServices:
		return "ServiceStationsWithOrWithoutAncillaryServices"
	case AutomatedFuelDispensers:
		return "AutomatedFuelDispensers"
	case BoatDealers:
		return "BoatDealers"
	case RecreationalAndUtilityTrailersCampDealers:
		return "RecreationalAndUtilityTrailersCampDealers"
	case MotorcycleDealers:
		return "MotorcycleDealers"
	case MotorHomeDealers:
		return "MotorHomeDealers"
	case SnowmobileDealers:
		return "SnowmobileDealers"
	case MensAndBoysClothingAndAccessoriesStores:
		return "MensAndBoysClothingAndAccessoriesStores"
	case WomensReadyToWearStores:
		return "WomensReadyToWearStores"
	case WomensAccessoryAndSpecialtyShops:
		return "WomensAccessoryAndSpecialtyShops"
	case ChildrensAndInfantsWearStores:
		return "ChildrensAndInfantsWearStores"
	case FamilyClothingStores:
		return "FamilyClothingStores"
	case SportsApparelRidingApparelStores:
		return "SportsApparelRidingApparelStores"
	case ShoeStores:
		return "ShoeStores"
	case FurriersAndFurShops:
		return "FurriersAndFurShops"
	case MensAndWomensClothingStores:
		return "MensAndWomensClothingStores"
	case TailorsSeamstressMendingAndAlterations:
		return "TailorsSeamstressMendingAndAlterations"
	case WigAndToupeeStores:
		return "WigAndToupeeStores"
	case MiscellaneousApparelAndAccessoryShops:
		return "MiscellaneousApparelAndAccessoryShops"
	case FurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances:
		return "FurnitureHomeFurnishingsAndEquipmentStoresExceptAppliances"
	case FloorCoveringStores:
		return "FloorCoveringStores"
	case DraperyWindowCoveringAndUpholsteryStores:
		return "DraperyWindowCoveringAndUpholsteryStores"
	case FireplaceScreensAndAccessoriesStores:
		return "FireplaceScreensAndAccessoriesStores"
	case MiscellaneousHomeFurnishingSpecialtyStores:
		return "MiscellaneousHomeFurnishingSpecialtyStores"
	case HouseholdApplianceStores:
		return "HouseholdApplianceStores"
	case ElectronicSales:
		return "ElectronicSales"
	case MusicStoresMusicalInstrumentsPianoSheetMusic:
		return "MusicStoresMusicalInstrumentsPianoSheetMusic"
	case ComputerSoftwareStores:
		return "ComputerSoftwareStores"
	case RecordShops:
		return "RecordShops"
	case Caterers:
		return "Caterers"
	case EatingPlacesAndRestaurants:
		return "EatingPlacesAndRestaurants"
	case AlcoholicBeveragesBarsTavernsCocktailLoungesNightClubsAndDiscotheques:
		return "AlcoholicBeveragesBarsTavernsCocktailLoungesNightClubsAndDiscotheques"
	case FastFoodRestaurants:
		return "FastFoodRestaurants"
	case DrugStoresAndPharmacies:
		return "DrugStoresAndPharmacies"
	case PackageStoresBeerWineAndLiquor:
		return "PackageStoresBeerWineAndLiquor"
	case UsedMerchandiseAndSecondhandStores:
		return "UsedMerchandiseAndSecondhandStores"
	case AntiqueShopsSalesRepairsAndRestorationServices:
		return "AntiqueShopsSalesRepairsAndRestorationServices"
	case PawnShopsAndSalvageYards:
		return "PawnShopsAndSalvageYards"
	case WreckingAndSalvageYards:
		return "WreckingAndSalvageYards"
	case AntiqueReproductions:
		return "AntiqueReproductions"
	case BicycleShopsSalesAndService:
		return "BicycleShopsSalesAndService"
	case SportingGoodsStores:
		return "SportingGoodsStores"
	case BookStores:
		return "BookStores"
	case StationeryStoresOfficeAndSchoolSupplyStores:
		return "StationeryStoresOfficeAndSchoolSupplyStores"
	case WatchClockJewelryAndSilverwareStores:
		return "WatchClockJewelryAndSilverwareStores"
	case HobbyToyAndGameShops:
		return "HobbyToyAndGameShops"
	case CameraAndPhotographicSupplyStores:
		return "CameraAndPhotographicSupplyStores"
	case CardShopsGiftNoveltyAndSouvenirShops:
		return "CardShopsGiftNoveltyAndSouvenirShops"
	case LeatherFoodsStores:
		return "LeatherFoodsStores"
	case SewingNeedleFabricAndPriceGoodsStores:
		return "SewingNeedleFabricAndPriceGoodsStores"
	case GlasswarecrystalStores:
		return "GlasswarecrystalStores"
	case DirectMarketingInsuranceService:
		return "DirectMarketingInsuranceService"
	case MailOrderHousesIncludingCatalogOrderStoresBookrecordClubsNoLongerPermittedForUsOriginalPresentments:
		return "MailOrderHousesIncludingCatalogOrderStoresBookrecordClubsNoLongerPermittedForUsOriginalPresentments"
	case DirectMarketingTravelRelatedArrangementsServices:
		return "DirectMarketingTravelRelatedArrangementsServices"
	case DoorToDoorSales:
		return "DoorToDoorSales"
	case DirectMarketingCatalogMerchant:
		return "DirectMarketingCatalogMerchant"
	case DirectMarketingCatalogAndCatalogAndRetailMerchant:
		return "DirectMarketingCatalogAndCatalogAndRetailMerchant"
	case DirectMarketingOutboundTelemarketingMerchant:
		return "DirectMarketingOutboundTelemarketingMerchant"
	case DirectMarketingInboundTeleservicesMerchant:
		return "DirectMarketingInboundTeleservicesMerchant"
	case DirectMarketingContinuitysubscriptionMerchant:
		return "DirectMarketingContinuitysubscriptionMerchant"
	case DirectMarketingNotElsewhereClassified:
		return "DirectMarketingNotElsewhereClassified"
	case ArtistsSupplyAndCraftShops:
		return "ArtistsSupplyAndCraftShops"
	case ArtDealersAndGalleries:
		return "ArtDealersAndGalleries"
	case StampAndCoinStoresPhilatelicAndNumismaticSupplies:
		return "StampAndCoinStoresPhilatelicAndNumismaticSupplies"
	case ReligiousGoodsStores:
		return "ReligiousGoodsStores"
	case HearingAidsSalesServiceAndSupplyStores:
		return "HearingAidsSalesServiceAndSupplyStores"
	case OrthopedicGoodsProstheticDevices:
		return "OrthopedicGoodsProstheticDevices"
	case CosmeticStores:
		return "CosmeticStores"
	case TypewriterStoresSalesRentalService:
		return "TypewriterStoresSalesRentalService"
	case FuelFuelOilWoodCoalLiquefiedPetroleum:
		return "FuelFuelOilWoodCoalLiquefiedPetroleum"
	case Florists:
		return "Florists"
	case CigarStoresAndStands:
		return "CigarStoresAndStands"
	case NewsDealersAndNewsstands:
		return "NewsDealersAndNewsstands"
	case PetShopsPetFoodsAndSuppliesStores:
		return "PetShopsPetFoodsAndSuppliesStores"
	case SwimmingPoolsSalesServiceAndSupplies:
		return "SwimmingPoolsSalesServiceAndSupplies"
	case ElectricRazorStoresSalesAndService:
		return "ElectricRazorStoresSalesAndService"
	case TentAndAwningShops:
		return "TentAndAwningShops"
	case MiscellaneousAndSpecialtyRetailStores:
		return "MiscellaneousAndSpecialtyRetailStores"
	case FinancialInstitutionsManualCashDisbursements:
		return "FinancialInstitutionsManualCashDisbursements"
	case FinancialInstitutionsAutomatedCash:
		return "FinancialInstitutionsAutomatedCash"
	case FinancialInstitutionsMerchandiseAndServices:
		return "FinancialInstitutionsMerchandiseAndServices"
	case NonFinancialInstitutionsForeignCurrencyMoneyOrdersNotWireTransferAndTravelersCheques:
		return "NonFinancialInstitutionsForeignCurrencyMoneyOrdersNotWireTransferAndTravelersCheques"
	case SecurityBrokersdealers:
		return "SecurityBrokersdealers"
	case InsuranceSalesUnderwritingAndPremiums:
		return "InsuranceSalesUnderwritingAndPremiums"
	case InsurancePremiumsNoLongerValidForFirstPresentmentWork:
		return "InsurancePremiumsNoLongerValidForFirstPresentmentWork"
	case InsuranceNotElsewhereClassifiedNoLongerValidForFirstPresentmentWork:
		return "InsuranceNotElsewhereClassifiedNoLongerValidForFirstPresentmentWork"
	case LodgingHotelsMotelsResortsCentralReservationServicesNotElsewhereClassified:
		return "LodgingHotelsMotelsResortsCentralReservationServicesNotElsewhereClassified"
	case Timeshares:
		return "Timeshares"
	case SportingAndRecreationalCamps:
		return "SportingAndRecreationalCamps"
	case TrailerParksAndCampGrounds:
		return "TrailerParksAndCampGrounds"
	case LaundryCleaningAndGarmentServices:
		return "LaundryCleaningAndGarmentServices"
	case LaundryFamilyAndCommercial:
		return "LaundryFamilyAndCommercial"
	case DryCleaners:
		return "DryCleaners"
	case CarpetAndUpholsteryCleaning:
		return "CarpetAndUpholsteryCleaning"
	case PhotographicStudios:
		return "PhotographicStudios"
	case BarberAndBeautyShops:
		return "BarberAndBeautyShops"
	case ShopRepairShopsAndShoeShineParlorsAndHatCleaningShops:
		return "ShopRepairShopsAndShoeShineParlorsAndHatCleaningShops"
	case FuneralServiceAndCrematories:
		return "FuneralServiceAndCrematories"
	case DatingAndEscortServices:
		return "DatingAndEscortServices"
	case TaxPreparationService:
		return "TaxPreparationService"
	case CounselingServiceDebtMarriagePersonal:
		return "CounselingServiceDebtMarriagePersonal"
	case BuyingshoppingServicesClubs:
		return "BuyingshoppingServicesClubs"
	case ClothingRentalCostumesFormalWearUniforms:
		return "ClothingRentalCostumesFormalWearUniforms"
	case MassageParlors:
		return "MassageParlors"
	case HealthAndBeautyShops:
		return "HealthAndBeautyShops"
	case MiscellaneousPersonalServicesNotElsewhereClassifies:
		return "MiscellaneousPersonalServicesNotElsewhereClassifies"
	case AdvertisingServices:
		return "AdvertisingServices"
	case ConsumerCreditReportingAgencies:
		return "ConsumerCreditReportingAgencies"
	case BlueprintingAndPhotocopyingServices:
		return "BlueprintingAndPhotocopyingServices"
	case CommercialPhotographyArtAndGraphics:
		return "CommercialPhotographyArtAndGraphics"
	case QuickCopyReproductionAndBlueprintingServices:
		return "QuickCopyReproductionAndBlueprintingServices"
	case StenographicAndSecretarialSupportServices:
		return "StenographicAndSecretarialSupportServices"
	// case DisinfectingServices:
	// 	return "DisinfectingServices"
	// case ExterminatingAndDisinfectingServices:
	// 	return "ExterminatingAndDisinfectingServices"
	case CleaningAndMaintenanceJanitorialServices:
		return "CleaningAndMaintenanceJanitorialServices"
	case EmploymentAgenciesTemporaryHelpServices:
		return "EmploymentAgenciesTemporaryHelpServices"
	case ComputerProgrammingIntegratedSystemsDesignAndDataProcessingServices:
		return "ComputerProgrammingIntegratedSystemsDesignAndDataProcessingServices"
	case InformationRetrievalServices:
		return "InformationRetrievalServices"
	case ComputerMaintenanceAndRepairServicesNotElsewhereClassified:
		return "ComputerMaintenanceAndRepairServicesNotElsewhereClassified"
	case ManagementConsultingAndPublicRelationsServices:
		return "ManagementConsultingAndPublicRelationsServices"
	case ProtectiveAndSecurityServicesIncludingArmoredCarsAndGuardDogs:
		return "ProtectiveAndSecurityServicesIncludingArmoredCarsAndGuardDogs"
	case EquipmentRentalAndLeasingServicesToolRentalFurnitureRentalAndApplianceRental:
		return "EquipmentRentalAndLeasingServicesToolRentalFurnitureRentalAndApplianceRental"
	case PhotofinishingLaboratoriesPhotoDeveloping:
		return "PhotofinishingLaboratoriesPhotoDeveloping"
	case BusinessServicesNotElsewhereClassified:
		return "BusinessServicesNotElsewhereClassified"
	case CarRentalCompaniesNotListedBelow:
		return "CarRentalCompaniesNotListedBelow"
	case TruckAndUtilityTrailerRentals:
		return "TruckAndUtilityTrailerRentals"
	case MotorHomeAndRecreationalVehicleRentals:
		return "MotorHomeAndRecreationalVehicleRentals"
	case AutomobileParkingLotsAndGarages:
		return "AutomobileParkingLotsAndGarages"
	case AutomotiveBodyRepairShops:
		return "AutomotiveBodyRepairShops"
	case TireReTreadingAndRepairShops:
		return "TireReTreadingAndRepairShops"
	case PaintShopsAutomotive:
		return "PaintShopsAutomotive"
	case AutomotiveServiceShops:
		return "AutomotiveServiceShops"
	case CarWashes:
		return "CarWashes"
	case TowingServices:
		return "TowingServices"
	case RadioRepairShops:
		return "RadioRepairShops"
	case AirConditioningAndRefrigerationRepairShops:
		return "AirConditioningAndRefrigerationRepairShops"
	case ElectricalAndSmallApplianceRepairShops:
		return "ElectricalAndSmallApplianceRepairShops"
	case WatchClockAndJewelryRepair:
		return "WatchClockAndJewelryRepair"
	case FurnitureFurnitureRepairAndFurnitureRefinishing:
		return "FurnitureFurnitureRepairAndFurnitureRefinishing"
	case WeldingRepair:
		return "WeldingRepair"
	case RepairShopsAndRelatedServicesMiscellaneous:
		return "RepairShopsAndRelatedServicesMiscellaneous"
	case MotionPicturesAndVideoTapeProductionAndDistribution:
		return "MotionPicturesAndVideoTapeProductionAndDistribution"
	case MotionPictureTheaters:
		return "MotionPictureTheaters"
	case VideoTapeRentalStores:
		return "VideoTapeRentalStores"
	case DanceHallsStudiosAndSchools:
		return "DanceHallsStudiosAndSchools"
	case TheatricalProducersExceptMotionPicturesTicketAgencies:
		return "TheatricalProducersExceptMotionPicturesTicketAgencies"
	case BandsOrchestrasAndMiscellaneousEntertainersNotElsewhereClassified:
		return "BandsOrchestrasAndMiscellaneousEntertainersNotElsewhereClassified"
	case BilliardAndPoolEstablishments:
		return "BilliardAndPoolEstablishments"
	case BowlingAlleys:
		return "BowlingAlleys"
	case CommercialSportsAthleticFieldsProfessionalSportClubsAndSportPromoters:
		return "CommercialSportsAthleticFieldsProfessionalSportClubsAndSportPromoters"
	case TouristAttractionsAndExhibits:
		return "TouristAttractionsAndExhibits"
	case GolfCoursesPublic:
		return "GolfCoursesPublic"
	case VideoAmusementGameSupplies:
		return "VideoAmusementGameSupplies"
	case VideoGameArcadesestablishments:
		return "VideoGameArcadesestablishments"
	case BettingIncludingLotteryTicketsCasinoGamingChipsOffTrackBettingAndWagers:
		return "BettingIncludingLotteryTicketsCasinoGamingChipsOffTrackBettingAndWagers"
	case AmusementParksCarnivalsCircusesFortuneTellers:
		return "AmusementParksCarnivalsCircusesFortuneTellers"
	case MembershipClubsSportsRecreationAthleticCountryClubsAndPrivateGolfCourses:
		return "MembershipClubsSportsRecreationAthleticCountryClubsAndPrivateGolfCourses"
	case AquariumsSeaAquariumsDolphinariums:
		return "AquariumsSeaAquariumsDolphinariums"
	case RecreationServicesNotElsewhereClassified:
		return "RecreationServicesNotElsewhereClassified"
	case DoctorsAndPhysiciansNotElsewhereClassified:
		return "DoctorsAndPhysiciansNotElsewhereClassified"
	case DentistsAndOrthodontists:
		return "DentistsAndOrthodontists"
	case Osteopaths:
		return "Osteopaths"
	case Chiropractors:
		return "Chiropractors"
	case OptometristsAndOphthalmologists:
		return "OptometristsAndOphthalmologists"
	case OpticiansOpticiansGoodsAndEyeglasses:
		return "OpticiansOpticiansGoodsAndEyeglasses"
	case OpticiansOpticalGoodsAndEyeglassesNoLongerValidForFirstPresentments:
		return "OpticiansOpticalGoodsAndEyeglassesNoLongerValidForFirstPresentments"
	case PodiatristsAndChiropodists:
		return "PodiatristsAndChiropodists"
	case NursingAndPersonalCareFacilities:
		return "NursingAndPersonalCareFacilities"
	case Hospitals:
		return "Hospitals"
	case MedicalAndDentalLaboratories:
		return "MedicalAndDentalLaboratories"
	case MedicalServicesAndHealthPractitionersNotElsewhereClassified:
		return "MedicalServicesAndHealthPractitionersNotElsewhereClassified"
	case LegalServicesAndAttorneys:
		return "LegalServicesAndAttorneys"
	case ElementaryAndSecondarySchools:
		return "ElementaryAndSecondarySchools"
	case CollegesJuniorCollegesUniversitiesAndProfessionalSchools:
		return "CollegesJuniorCollegesUniversitiesAndProfessionalSchools"
	case CorrespondenceSchools:
		return "CorrespondenceSchools"
	case BusinessAndSecretarialSchools:
		return "BusinessAndSecretarialSchools"
	case VocationalSchoolsAndTradeSchools:
		return "VocationalSchoolsAndTradeSchools"
	case SchoolsAndEducationalServicesNotElsewhereClassified:
		return "SchoolsAndEducationalServicesNotElsewhereClassified"
	case ChildCareServices:
		return "ChildCareServices"
	case CharitableAndSocialServiceOrganizations:
		return "CharitableAndSocialServiceOrganizations"
	case CivicFraternalAndSocialAssociations:
		return "CivicFraternalAndSocialAssociations"
	case PoliticalOrganizations:
		return "PoliticalOrganizations"
	case ReligiousOrganizations:
		return "ReligiousOrganizations"
	case AutomobileAssociations:
		return "AutomobileAssociations"
	case MembershipOrganizationsNotElsewhereClassified:
		return "MembershipOrganizationsNotElsewhereClassified"
	case TestingLaboratoriesNonMedical:
		return "TestingLaboratoriesNonMedical"
	case ArchitecturalEngineeringAndSurveyingServices:
		return "ArchitecturalEngineeringAndSurveyingServices"
	case AccountingAuditingAndBookkeepingServices:
		return "AccountingAuditingAndBookkeepingServices"
	case ProfessionalServicesNotElsewhereDefined:
		return "ProfessionalServicesNotElsewhereDefined"
	case CourtCostsIncludingAlimonyAndChildSupport:
		return "CourtCostsIncludingAlimonyAndChildSupport"
	case Fines:
		return "Fines"
	case BailAndBondPayments:
		return "BailAndBondPayments"
	case TaxPayments:
		return "TaxPayments"
	case GovernmentServicesNotElsewhereClassified:
		return "GovernmentServicesNotElsewhereClassified"
	case PostalServicesGovernmentOnly:
		return "PostalServicesGovernmentOnly"
	case IntraGovernmentTransactions:
		return "IntraGovernmentTransactions"
	case AutomatedReferralServiceForVisaOnly:
		return "AutomatedReferralServiceForVisaOnly"
	case VisaCredentialServiceForVisaOnly:
		return "VisaCredentialServiceForVisaOnly"
	case GcasEmergencyServicesForVisaOnly:
		return "GcasEmergencyServicesForVisaOnly"
	case IntraCompanyPurchasesForVisaOnly:
		return "IntraCompanyPurchasesForVisaOnly"
	}
	return ""
}

type MotoEcomIndicatorType rune

const (
	NonMailOrderTransaction                           MotoEcomIndicatorType = ' '
	SpaceOrEmptyMoto                                  MotoEcomIndicatorType = '0'
	SingleTransaction                                 MotoEcomIndicatorType = '1'
	RecurringTransaction                              MotoEcomIndicatorType = '2'
	InstallmentPayment                                MotoEcomIndicatorType = '3'
	UnknownClassification                             MotoEcomIndicatorType = '4'
	FullyAuthenticated3dSecureTransaction             MotoEcomIndicatorType = '5'
	No3dsecureAuthenticationFromCardholder            MotoEcomIndicatorType = '6'
	TransmittedWithSslOrChannelEncryption             MotoEcomIndicatorType = '7'
	NonSecureTransactionInTheClear                    MotoEcomIndicatorType = '8'
	NoSetAuthenticationFromCardholder                 MotoEcomIndicatorType = '9'
	FirstRecurringSecureCodePhoneOrderTransaction     MotoEcomIndicatorType = 'R'
	SingleNonRecurringSecureCodePhoneOrderTransaction MotoEcomIndicatorType = 'T'
)

func (x MotoEcomIndicatorType) String() string {
	switch x {
	case NonMailOrderTransaction:
		return "NonMailOrderTransaction"
	case SpaceOrEmptyMoto:
		return "SpaceOrEmptyMoto"
	case SingleTransaction:
		return "SingleTransaction"
	case RecurringTransaction:
		return "RecurringTransaction"
	case InstallmentPayment:
		return "InstallmentPayment"
	case UnknownClassification:
		return "UnknownClassification"
	case FullyAuthenticated3dSecureTransaction:
		return "FullyAuthenticated3dSecureTransaction"
	case No3dsecureAuthenticationFromCardholder:
		return "No3dsecureAuthenticationFromCardholder"
	case TransmittedWithSslOrChannelEncryption:
		return "TransmittedWithSslOrChannelEncryption"
	case NonSecureTransactionInTheClear:
		return "NonSecureTransactionInTheClear"
	case NoSetAuthenticationFromCardholder:
		return "NoSetAuthenticationFromCardholder"
	case FirstRecurringSecureCodePhoneOrderTransaction:
		return "FirstRecurringSecureCodePhoneOrderTransaction"
	case SingleNonRecurringSecureCodePhoneOrderTransaction:
		return "SingleNonRecurringSecureCodePhoneOrderTransaction"
	}
	return ""
}

type NationalTaxFlagType rune

const (
	TaxNotIncluded NationalTaxFlagType = 0
	TaxIncluded    NationalTaxFlagType = 1
)

func (x NationalTaxFlagType) String() string {
	switch x {
	case TaxNotIncluded:
		return "TaxNotIncluded"
	case TaxIncluded:
		return "TaxIncluded"
	}
	return ""
}

type NetworkIdentificationType rune

const (
	SpaceOrEmptyNetworkId NetworkIdentificationType = ' '
	EftIllinois           NetworkIdentificationType = '1'
	AlaskaOption          NetworkIdentificationType = '3'
	VisaCheckCardII       NetworkIdentificationType = '5'
	ShazamITS             NetworkIdentificationType = '7'
	MaestroNetwork        NetworkIdentificationType = '8'
	Transfund             NetworkIdentificationType = 'A'
	MellonSpecial         NetworkIdentificationType = 'B'
	Alert                 NetworkIdentificationType = 'C'
	Accel                 NetworkIdentificationType = 'E'
	MagicLine             NetworkIdentificationType = 'F'
	Interlink             NetworkIdentificationType = 'G'
	Tyme                  NetworkIdentificationType = 'H'
	Most                  NetworkIdentificationType = 'I'
	Gulfnet               NetworkIdentificationType = 'J'
	EbtNetwork            NetworkIdentificationType = 'K'
	Pulse                 NetworkIdentificationType = 'L'
	Bankmate              NetworkIdentificationType = 'M'
	CashStation           NetworkIdentificationType = 'N'
	StarWestExplore       NetworkIdentificationType = 'Q'
	MoneyStation          NetworkIdentificationType = 'S'
	Affn                  NetworkIdentificationType = 'U'
	VisaNetwork           NetworkIdentificationType = 'V'
	StarSoutheastHonor    NetworkIdentificationType = 'W'
	Infyank24nyce         NetworkIdentificationType = 'Y'
	StarNortheastMac      NetworkIdentificationType = 'Z'
)

func (x NetworkIdentificationType) String() string {
	switch x {
	case SpaceOrEmptyNetworkId:
		return "SpaceOrEmptyNetworkId"
	case EftIllinois:
		return "EftIllinois"
	case AlaskaOption:
		return "AlaskaOption"
	case VisaCheckCardII:
		return "VisaCheckCardII"
	case ShazamITS:
		return "ShazamITS"
	case MaestroNetwork:
		return "MaestroNetwork"
	case Transfund:
		return "Transfund"
	case MellonSpecial:
		return "MellonSpecial"
	case Alert:
		return "Alert"
	case Accel:
		return "Accel"
	case MagicLine:
		return "MagicLine"
	case Interlink:
		return "Interlink"
	case Tyme:
		return "Tyme"
	case Most:
		return "Most"
	case Gulfnet:
		return "Gulfnet"
	case EbtNetwork:
		return "EbtNetwork"
	case Pulse:
		return "Pulse"
	case Bankmate:
		return "Bankmate"
	case CashStation:
		return "CashStation"
	case StarWestExplore:
		return "StarWestExplore"
	case MoneyStation:
		return "MoneyStation"
	case Affn:
		return "Affn"
	case VisaNetwork:
		return "VisaNetwork"
	case StarSoutheastHonor:
		return "StarSoutheastHonor"
	case Infyank24nyce:
		return "Infyank24nyce"
	case StarNortheastMac:
		return "StarNortheastMac"
	}
	return ""
}

type NoShowIndicatorType uint

const (
	NotApplicable NoShowIndicatorType = 0
	NoShow        NoShowIndicatorType = 1
)

func (x NoShowIndicatorType) String() string {
	switch x {
	case NotApplicable:
		return "NotApplicable"
	case NoShow:
		return "NoShow"
	}
	return ""
}

type PrestigiousPropertyType rune

const (
	AutoRentalOrNonParticipatingProperty PrestigiousPropertyType = ' '
	PropertyWith500Limit                 PrestigiousPropertyType = 'D'
	PropertyWith1000Limit                PrestigiousPropertyType = 'B'
	PropertyWith1500Limit                PrestigiousPropertyType = 'S'
)

func (x PrestigiousPropertyType) String() string {
	switch x {
	case AutoRentalOrNonParticipatingProperty:
		return "AutoRentalOrNonParticipatingProperty"
	case PropertyWith500Limit:
		return "PropertyWith500Limit"
	case PropertyWith1000Limit:
		return "PropertyWith1000Limit"
	case PropertyWith1500Limit:
		return "PropertyWith1500Limit"
	}
	return ""
}

type ProtocolCompletionStatusType uint

const (
	HostAcknowledgementReceived    ProtocolCompletionStatusType = 0
	HostAcknowledgementNotReceived ProtocolCompletionStatusType = 1
)

func (x ProtocolCompletionStatusType) String() string {
	switch x {
	case HostAcknowledgementReceived:
		return "HostAcknowledgementReceived"
	case HostAcknowledgementNotReceived:
		return "HostAcknowledgementNotReceived"
	}
	return ""
}

type PurchaseIdFormatType rune

const (
	NotUsed                    PurchaseIdFormatType = 'Z'
	PurchaseIdFormatReserved   PurchaseIdFormatType = '0'
	ReservedSpace              PurchaseIdFormatType = ' '
	DirectMarketingOrderNumber PurchaseIdFormatType = '1'
	Reserved2                  PurchaseIdFormatType = '2'
	AutoRentalAgreementNumber  PurchaseIdFormatType = '3'
	HotelFolioNumber           PurchaseIdFormatType = '4'
)

func (x PurchaseIdFormatType) String() string {
	switch x {
	case NotUsed:
		return "NotUsed"
	case PurchaseIdFormatReserved:
		return "PurchaseIdFormatReserved"
	case ReservedSpace:
		return "ReservedSpace"
	case DirectMarketingOrderNumber:
		return "DirectMarketingOrderNumber"
	case Reserved2:
		return "Reserved2"
	case AutoRentalAgreementNumber:
		return "AutoRentalAgreementNumber"
	case HotelFolioNumber:
		return "HotelFolioNumber"
	}
	return ""
}

type RecordFormatType rune

const (
	CreditCardAuthorizationRequest  RecordFormatType = 'D'
	CreditCardAuthorizationResponse RecordFormatType = 'E'
	DebitEbtRequest                 RecordFormatType = 'T'
	DebitEbtResponse                RecordFormatType = 'U'
	DebitEbtConfirmation            RecordFormatType = 'V'
	Settlement                      RecordFormatType = 'K'
	PekExchangeRequest              RecordFormatType = 'P'
	KeepAlive                       RecordFormatType = 'A'
	Iso8583                         RecordFormatType = 'I'
)

func (x RecordFormatType) String() string {
	switch x {
	case CreditCardAuthorizationRequest:
		return "CreditCardAuthorizationRequest"
	case CreditCardAuthorizationResponse:
		return "CreditCardAuthorizationResponse"
	case DebitEbtRequest:
		return "DebitEbtRequest"
	case DebitEbtResponse:
		return "DebitEbtResponse"
	case DebitEbtConfirmation:
		return "DebitEbtConfirmation"
	case Settlement:
		return "Settlement"
	case PekExchangeRequest:
		return "PekExchangeRequest"
	case KeepAlive:
		return "KeepAlive"
	case Iso8583:
		return "Iso8583"
	}
	return ""
}

type RecordIndicatorType rune

const (
	HeaderRecord                        RecordIndicatorType = 'H'
	ParameterRecord                     RecordIndicatorType = 'P'
	DetailRecord                        RecordIndicatorType = 'D'
	AmericanExpressLineItemDetailRecord RecordIndicatorType = 'A'
	CommercialCardLineItemDetailRecord  RecordIndicatorType = 'L'
	CommercialCardTripLegDetailRecord   RecordIndicatorType = 'M'
	FleetLineItemDetailRecord           RecordIndicatorType = 'F'
	EncryptedTransmissionBlockRecord    RecordIndicatorType = 'E'
	ChipCardAddendumRecord              RecordIndicatorType = 'C'
	TrailerRecord                       RecordIndicatorType = 'T'
	TrailerResponseRecord               RecordIndicatorType = 'R'
	GroupMapExtensionRecord             RecordIndicatorType = 'X'
)

func (x RecordIndicatorType) String() string {
	switch x {
	case HeaderRecord:
		return "HeaderRecord"
	case ParameterRecord:
		return "ParameterRecord"
	case DetailRecord:
		return "DetailRecord"
	case CommercialCardLineItemDetailRecord:
		return "CommercialCardLineItemDetailRecord"
	case CommercialCardTripLegDetailRecord:
		return "CommercialCardTripLegDetailRecord"
	case TrailerRecord:
		return "TrailerRecord"
	case TrailerResponseRecord:
		return "TrailerResponseRecord"
	}
	return ""
}

type ReimbursementAttributeType rune

const (
	EbtNonDebitOrNonInterlinkDebit                   ReimbursementAttributeType = '0'
	PreexistingQualifiedInterlinkSuperMarketMerchant ReimbursementAttributeType = 'W'
	PreexistingQualifiedInterlinkRetailMerchant      ReimbursementAttributeType = 'X'
	QualifiedInterlinkSuperMarketMerchant            ReimbursementAttributeType = 'Y'
	StandardInterlinkRetailMerchant                  ReimbursementAttributeType = 'Z'
)

func (x ReimbursementAttributeType) String() string {
	switch x {
	case EbtNonDebitOrNonInterlinkDebit:
		return "EbtNonDebitOrNonInterlinkDebit"
	case PreexistingQualifiedInterlinkSuperMarketMerchant:
		return "PreexistingQualifiedInterlinkSuperMarketMerchant"
	case PreexistingQualifiedInterlinkRetailMerchant:
		return "PreexistingQualifiedInterlinkRetailMerchant"
	case QualifiedInterlinkSuperMarketMerchant:
		return "QualifiedInterlinkSuperMarketMerchant"
	case StandardInterlinkRetailMerchant:
		return "StandardInterlinkRetailMerchant"
	}
	return ""
}

type RequestAciType rune

const (
	DeviceIsNotCpsMeritCapableOrPosCheck      RequestAciType = 'N'
	DeviceIsCpsMeritCapableOrCreditOrOffline  RequestAciType = 'Y'
	CpsCapableManuallyKeyedOrMcTips           RequestAciType = 'P'
	CpsCapableIncrementalAuthorizationRequest RequestAciType = 'I'
	RecurringOrInstallmentPayments            RequestAciType = 'R'
)

func (x RequestAciType) String() string {
	switch x {
	case DeviceIsNotCpsMeritCapableOrPosCheck:
		return "DeviceIsNotCpsMeritCapableOrPosCheck"
	case DeviceIsCpsMeritCapableOrCreditOrOffline:
		return "DeviceIsCpsMeritCapableOrCreditOrOffline"
	case CpsCapableManuallyKeyedOrMcTips:
		return "CpsCapableManuallyKeyedOrMcTips"
	case CpsCapableIncrementalAuthorizationRequest:
		return "CpsCapableIncrementalAuthorizationRequest"
	case RecurringOrInstallmentPayments:
		return "RecurringOrInstallmentPayments"
	}
	return ""
}

type ResponseCodeType string

const (
	CreditRefundOrOfflineVoiceApprovedTransaction ResponseCodeType = "  "
	OfflineChipCardDeclinedUnableToGoOnlineZ3     ResponseCodeType = "Z3"
	OfflineChipCardApprovalY1                     ResponseCodeType = "Y1"
	OfflineChipCardDeclinedZ1                     ResponseCodeType = "Z1"
	OnlineApprovedTransactionT0                   ResponseCodeType = "T0"
	StopPaymentOrderR0                            ResponseCodeType = "R0"
	RevocationOfAuthorizationOrderR1              ResponseCodeType = "R1"
	ApprovalApprovedAndCompleted00                ResponseCodeType = "00"
	CallReferToIssuer01                           ResponseCodeType = "01"
	CallReferToIssueSpecialCondition02            ResponseCodeType = "02"
	TermIDErrorInvalidMerchantID03                ResponseCodeType = "03"
	HoldCallOrPickUpCardNoFraud04                 ResponseCodeType = "04"
	DeclineDoNotHonor05                           ResponseCodeType = "05"
	GeneralError06                                ResponseCodeType = "06"
	//*CheckServiceCustomTextErrorResponseTextFromCheckService06 ResponseCodeType = "06"
	HoldCallOrPickUpCardPickUpCardSpecialConditionFraudAccount07        ResponseCodeType = "07"
	ApprovalHonorMastercardWithId08                                     ResponseCodeType = "08"
	PartialApprovalForTheAuthorizedAmountReturnedInGroupIiiVersion02210 ResponseCodeType = "10"
	ApprovalVipApproval11                                               ResponseCodeType = "11"
	InvalidTransInvalidTransaction12                                    ResponseCodeType = "12"
	AmountErrorInvalidAmount13                                          ResponseCodeType = "13"
	CardNoErrorInvalidCardNumber14                                      ResponseCodeType = "14"
	NoSuchIssuer15                                                      ResponseCodeType = "15"
	ReEnterTransaction19                                                ResponseCodeType = "19"
	NoActionTakenUnableToBackOutTransaction21                           ResponseCodeType = "21"
	NoReplyFileIsTemporarilyUnavailable28                               ResponseCodeType = "28"
	//TransactionCancelledMastercardUseOnlyResponseCodeType = "34" //TransactionCancelledFraudConcernUsedInReversalRequestsOnly34
	NoCreditAcctNoCreditAccount39                                                  ResponseCodeType = "39"
	HoldCallOrPickUpCardLostCardFraudAccount41                                     ResponseCodeType = "41"
	HoldCallOrPickUpCardStolenCardFraudAccount43                                   ResponseCodeType = "43"
	DeclineInsufficientFunds51                                                     ResponseCodeType = "51"
	NoCheckAccountNoCheckingAccount52                                              ResponseCodeType = "52"
	NoSaveAccountNoSavingsAccount53                                                ResponseCodeType = "53"
	ExpiredCardExpiredCard54                                                       ResponseCodeType = "54"
	WrongPinIncorrectPin55                                                         ResponseCodeType = "55"
	ServNotAllowedTransactionNotPermittedCard57                                    ResponseCodeType = "57"
	ServNotAllowedTransactionNotPermittedTerminal58                                ResponseCodeType = "58"
	ServNotAllowedTransactionNotPermittedMerchant59                                ResponseCodeType = "59"
	DeclinedExceedsWithdrawalLimit61                                               ResponseCodeType = "61"
	DeclinedInvalidServiceCodeRestricted62                                         ResponseCodeType = "62"
	SecViolationSecurityViolation63                                                ResponseCodeType = "63"
	DeclinedActivityLimitExceeded65                                                ResponseCodeType = "65"
	PinExceededPinTriedExceeded75                                                  ResponseCodeType = "75"
	UnsolicatedReversalUnableToLocateNoMatch76                                     ResponseCodeType = "76"
	NoActionTakenInconsistantDataReversedOrRepeat77                                ResponseCodeType = "77"
	NoAccountNoAccount78                                                           ResponseCodeType = "78"
	AlreadyReversedAlreadyReversedAtSwitch79                                       ResponseCodeType = "79"
	NoImpactNoFinancialImpactUsedInReversalResponsesToDeclinedOriginals80          ResponseCodeType = "80"
	EncryptionErrorCryptographicError81                                            ResponseCodeType = "81"
	IncorrectCvvCvvDataIsNotCorrect82                                              ResponseCodeType = "82"
	CannotVerifyPinCannotVerifyPin83                                               ResponseCodeType = "83"
	CardOkNoReasonToDecline85                                                      ResponseCodeType = "85"
	CannotVerifyPinCannotVerifyPin86                                               ResponseCodeType = "86"
	NoReplyIssuerOrSwitchIsUnavailable91                                           ResponseCodeType = "91"
	InvalidRoutingDestinationNotFound92                                            ResponseCodeType = "92"
	DeclineViolationCannotComplete93                                               ResponseCodeType = "93"
	DuplicateTransUnableToLocateNoMatch94                                          ResponseCodeType = "94"
	SystemErrorSystemMalfunction96                                                 ResponseCodeType = "96"
	ActivatedPosDeviceAuthenticationSuccessfulA1                                   ResponseCodeType = "A1"
	NotActivatedPosDeviceAuthenticationNotSuccessfulA2                             ResponseCodeType = "A2"
	DeactivatedPosDeviceDeactivationSuccessfulA3                                   ResponseCodeType = "A3"
	SrchgNotAllowedSurchargeAmountNotPermittedOnDebitCardsOrEbtFoodStampsB1        ResponseCodeType = "B1"
	SrchgNotAllowedSurchargeAmountNotSupportedByDebitNetworkIssuerB2               ResponseCodeType = "B2"
	FailureCvCardTypeVerificationErrorCV                                           ResponseCodeType = "CV"
	EncrNotConfigdEncryptionIsNotConfiguredE1                                      ResponseCodeType = "E1"
	TermNotAuthentTerminalIsNotAuthenticatedE2                                     ResponseCodeType = "E2"
	DecryptFailureDataCouldNotBeDecryptedE3                                        ResponseCodeType = "E3"
	AcctLengthErrVerificationErrorEA                                               ResponseCodeType = "EA"
	CheckDigitErrVerificationErrorEB                                               ResponseCodeType = "EB"
	CidFormatErrorVerificationErrorEC                                              ResponseCodeType = "EC"
	FailureHvHierarchyVerificationErrorHV                                          ResponseCodeType = "HV"
	TokenResponseTokenRequestWasProcessedK0                                        ResponseCodeType = "K0"
	TokenNotConfigTokenizationIsNotConfiguredK1                                    ResponseCodeType = "K1"
	TermNotAuthentTerminalIsNotAuthenticatedK2                                     ResponseCodeType = "K2"
	TokenFailureDataCouldNotBeDeTokenizedK3                                        ResponseCodeType = "K3"
	CashbackNotAvlCashBackServiceNotAvailableN3                                    ResponseCodeType = "N3"
	DeclineExceedsIssuerWithdrawalLimitN4                                          ResponseCodeType = "N4"
	Ccv2MismatchCvv2ValueSuppliedIsInvalidN7                                       ResponseCodeType = "N7"
	StopRecurringCustomerRequestedStopOfSpecificRecurringPaymentR0                 ResponseCodeType = "R0"
	StopRecurringCustomerRequestedStopOfAllRecurringPaymentsFromSpecificMerchantR1 ResponseCodeType = "R1"
	ApprovalFirstCheckIsOkAndHasBeenConvertedT0                                    ResponseCodeType = "T0"
	CannotConvertCheckIsOkButCannotBeConvertedThisIsADeclinedTransactionT1         ResponseCodeType = "T1"
	InvalidAbaInvalidAbaNumberNotAnAchParticipantT2                                ResponseCodeType = "T2"
	AmountErrorAmountGreaterThanTheLimitT3                                         ResponseCodeType = "T3"
	UnpaidItemsUnpaidItemsFailedNegativeFileCheckT4                                ResponseCodeType = "T4"
	DuplicateNumberDuplicateCheckNumberT5                                          ResponseCodeType = "T5"
	MicrErrorT6                                                                    ResponseCodeType = "T6"
	TooManyChecksTooManyChecksOverMerchantOrBankLimitT7                            ResponseCodeType = "T7"
	FailureVmDailyThresholdExceededV1                                              ResponseCodeType = "V1"

	// MIP spec
	ApprovedHonorWithId08        ResponseCodeType = "08"
	PartialApproval10            ResponseCodeType = "10"
	FormatError30                ResponseCodeType = "30"
	ContactCardIssuer70          ResponseCodeType = "70"
	PinNotChanged71              ResponseCodeType = "71"
	InvalidToAccount76           ResponseCodeType = "76"
	InvalidFromAccount77         ResponseCodeType = "77"
	InvalidAccount78             ResponseCodeType = "78"
	InvalidAuthLifeCycle84       ResponseCodeType = "84"
	ApprovedPurchaseNoCashback87 ResponseCodeType = "87"
	CryptographicFailure88       ResponseCodeType = "88"
	UnacceptablePin89            ResponseCodeType = "89"
	DuplicateTransmission94      ResponseCodeType = "94"
)

func (x ResponseCodeType) String() string {
	switch x {
	case CreditRefundOrOfflineVoiceApprovedTransaction:
		return "CreditRefundOrOfflineVoiceApprovedTransaction"
	case OfflineChipCardDeclinedUnableToGoOnlineZ3:
		return "OfflineChipCardDeclinedUnableToGoOnlineZ3"
	case OfflineChipCardApprovalY1:
		return "OfflineChipCardApprovalY1"
	case OfflineChipCardDeclinedZ1:
		return "OfflineChipCardDeclinedZ1"
	case OnlineApprovedTransactionT0:
		return "OnlineApprovedTransactionT0"
	case StopPaymentOrderR0:
		return "StopPaymentOrderR0"
	case RevocationOfAuthorizationOrderR1:
		return "RevocationOfAuthorizationOrderR1"
	case ApprovalApprovedAndCompleted00:
		return "ApprovalApprovedAndCompleted00"
	case CallReferToIssuer01:
		return "CallReferToIssuer01"
	case CallReferToIssueSpecialCondition02:
		return "CallReferToIssueSpecialCondition02"
	case TermIDErrorInvalidMerchantID03:
		return "TermIDErrorInvalidMerchantID03"
	case HoldCallOrPickUpCardNoFraud04:
		return "HoldCallOrPickUpCardNoFraud04"
	case DeclineDoNotHonor05:
		return "DeclineDoNotHonor05"
	case GeneralError06:
		return "GeneralError06"
	//*case CheckServiceCustomTextErrorResponseTextFromCheckService06: return "CheckServiceCustomTextErrorResponseTextFromCheckService06"
	case HoldCallOrPickUpCardPickUpCardSpecialConditionFraudAccount07:
		return "HoldCallOrPickUpCardPickUpCardSpecialConditionFraudAccount07"
	case ApprovalHonorMastercardWithId08:
		return "ApprovalHonorMastercardWithId08"
	case PartialApprovalForTheAuthorizedAmountReturnedInGroupIiiVersion02210:
		return "PartialApprovalForTheAuthorizedAmountReturnedInGroupIiiVersion02210"
	case ApprovalVipApproval11:
		return "ApprovalVipApproval11"
	case InvalidTransInvalidTransaction12:
		return "InvalidTransInvalidTransaction12"
	case AmountErrorInvalidAmount13:
		return "AmountErrorInvalidAmount13"
	case CardNoErrorInvalidCardNumber14:
		return "CardNoErrorInvalidCardNumber14"
	case NoSuchIssuer15:
		return "NoSuchIssuer15"
	case ReEnterTransaction19:
		return "ReEnterTransaction19"
	case NoActionTakenUnableToBackOutTransaction21:
		return "NoActionTakenUnableToBackOutTransaction21"
	case NoReplyFileIsTemporarilyUnavailable28:
		return "NoReplyFileIsTemporarilyUnavailable28"
	//case TransactionCancelledMastercardUseOnlyResponseCodeType: return "TransactionCancelledMastercardUseOnlyResponseCodeType"
	case NoCreditAcctNoCreditAccount39:
		return "NoCreditAcctNoCreditAccount39"
	case HoldCallOrPickUpCardLostCardFraudAccount41:
		return "HoldCallOrPickUpCardLostCardFraudAccount41"
	case HoldCallOrPickUpCardStolenCardFraudAccount43:
		return "HoldCallOrPickUpCardStolenCardFraudAccount43"
	case DeclineInsufficientFunds51:
		return "DeclineInsufficientFunds51"
	case NoCheckAccountNoCheckingAccount52:
		return "NoCheckAccountNoCheckingAccount52"
	case NoSaveAccountNoSavingsAccount53:
		return "NoSaveAccountNoSavingsAccount53"
	case ExpiredCardExpiredCard54:
		return "ExpiredCardExpiredCard54"
	case WrongPinIncorrectPin55:
		return "WrongPinIncorrectPin55"
	case ServNotAllowedTransactionNotPermittedCard57:
		return "ServNotAllowedTransactionNotPermittedCard57"
	case ServNotAllowedTransactionNotPermittedTerminal58:
		return "ServNotAllowedTransactionNotPermittedTerminal58"
	case ServNotAllowedTransactionNotPermittedMerchant59:
		return "ServNotAllowedTransactionNotPermittedMerchant59"
	case DeclinedExceedsWithdrawalLimit61:
		return "DeclinedExceedsWithdrawalLimit61"
	case DeclinedInvalidServiceCodeRestricted62:
		return "DeclinedInvalidServiceCodeRestricted62"
	case SecViolationSecurityViolation63:
		return "SecViolationSecurityViolation63"
	case DeclinedActivityLimitExceeded65:
		return "DeclinedActivityLimitExceeded65"
	case PinExceededPinTriedExceeded75:
		return "PinExceededPinTriedExceeded75"
	case UnsolicatedReversalUnableToLocateNoMatch76:
		return "UnsolicatedReversalUnableToLocateNoMatch76"
	case NoActionTakenInconsistantDataReversedOrRepeat77:
		return "NoActionTakenInconsistantDataReversedOrRepeat77"
	case NoAccountNoAccount78:
		return "NoAccountNoAccount78"
	case AlreadyReversedAlreadyReversedAtSwitch79:
		return "AlreadyReversedAlreadyReversedAtSwitch79"
	case NoImpactNoFinancialImpactUsedInReversalResponsesToDeclinedOriginals80:
		return "NoImpactNoFinancialImpactUsedInReversalResponsesToDeclinedOriginals80"
	case EncryptionErrorCryptographicError81:
		return "EncryptionErrorCryptographicError81"
	case IncorrectCvvCvvDataIsNotCorrect82:
		return "IncorrectCvvCvvDataIsNotCorrect82"
	case CannotVerifyPinCannotVerifyPin83:
		return "CannotVerifyPinCannotVerifyPin83"
	case CardOkNoReasonToDecline85:
		return "CardOkNoReasonToDecline85"
	case CannotVerifyPinCannotVerifyPin86:
		return "CannotVerifyPinCannotVerifyPin86"
	case NoReplyIssuerOrSwitchIsUnavailable91:
		return "NoReplyIssuerOrSwitchIsUnavailable91"
	case InvalidRoutingDestinationNotFound92:
		return "InvalidRoutingDestinationNotFound92"
	case DeclineViolationCannotComplete93:
		return "DeclineViolationCannotComplete93"
	case DuplicateTransUnableToLocateNoMatch94:
		return "DuplicateTransUnableToLocateNoMatch94"
	case SystemErrorSystemMalfunction96:
		return "SystemErrorSystemMalfunction96"
	case ActivatedPosDeviceAuthenticationSuccessfulA1:
		return "ActivatedPosDeviceAuthenticationSuccessfulA1"
	case NotActivatedPosDeviceAuthenticationNotSuccessfulA2:
		return "NotActivatedPosDeviceAuthenticationNotSuccessfulA2"
	case DeactivatedPosDeviceDeactivationSuccessfulA3:
		return "DeactivatedPosDeviceDeactivationSuccessfulA3"
	case SrchgNotAllowedSurchargeAmountNotPermittedOnDebitCardsOrEbtFoodStampsB1:
		return "SrchgNotAllowedSurchargeAmountNotPermittedOnDebitCardsOrEbtFoodStampsB1"
	case SrchgNotAllowedSurchargeAmountNotSupportedByDebitNetworkIssuerB2:
		return "SrchgNotAllowedSurchargeAmountNotSupportedByDebitNetworkIssuerB2"
	case FailureCvCardTypeVerificationErrorCV:
		return "FailureCvCardTypeVerificationErrorCV"
	case EncrNotConfigdEncryptionIsNotConfiguredE1:
		return "EncrNotConfigdEncryptionIsNotConfiguredE1"
	case TermNotAuthentTerminalIsNotAuthenticatedE2:
		return "TermNotAuthentTerminalIsNotAuthenticatedE2"
	case DecryptFailureDataCouldNotBeDecryptedE3:
		return "DecryptFailureDataCouldNotBeDecryptedE3"
	case AcctLengthErrVerificationErrorEA:
		return "AcctLengthErrVerificationErrorEA"
	case CheckDigitErrVerificationErrorEB:
		return "CheckDigitErrVerificationErrorEB"
	case CidFormatErrorVerificationErrorEC:
		return "CidFormatErrorVerificationErrorEC"
	case FailureHvHierarchyVerificationErrorHV:
		return "FailureHvHierarchyVerificationErrorHV"
	case TokenResponseTokenRequestWasProcessedK0:
		return "TokenResponseTokenRequestWasProcessedK0"
	case TokenNotConfigTokenizationIsNotConfiguredK1:
		return "TokenNotConfigTokenizationIsNotConfiguredK1"
	case TermNotAuthentTerminalIsNotAuthenticatedK2:
		return "TermNotAuthentTerminalIsNotAuthenticatedK2"
	case TokenFailureDataCouldNotBeDeTokenizedK3:
		return "TokenFailureDataCouldNotBeDeTokenizedK3"
	case CashbackNotAvlCashBackServiceNotAvailableN3:
		return "CashbackNotAvlCashBackServiceNotAvailableN3"
	case DeclineExceedsIssuerWithdrawalLimitN4:
		return "DeclineExceedsIssuerWithdrawalLimitN4"
	case Ccv2MismatchCvv2ValueSuppliedIsInvalidN7:
		return "Ccv2MismatchCvv2ValueSuppliedIsInvalidN7"
	// case StopRecurringCustomerRequestedStopOfSpecificRecurringPaymentR0:
	// 	return "StopRecurringCustomerRequestedStopOfSpecificRecurringPaymentR0"
	// case StopRecurringCustomerRequestedStopOfAllRecurringPaymentsFromSpecificMerchantR1:
	// 	return "StopRecurringCustomerRequestedStopOfAllRecurringPaymentsFromSpecificMerchantR1"
	// case ApprovalFirstCheckIsOkAndHasBeenConvertedT0:
	// 	return "ApprovalFirstCheckIsOkAndHasBeenConvertedT0"
	case CannotConvertCheckIsOkButCannotBeConvertedThisIsADeclinedTransactionT1:
		return "CannotConvertCheckIsOkButCannotBeConvertedThisIsADeclinedTransactionT1"
	case InvalidAbaInvalidAbaNumberNotAnAchParticipantT2:
		return "InvalidAbaInvalidAbaNumberNotAnAchParticipantT2"
	case AmountErrorAmountGreaterThanTheLimitT3:
		return "AmountErrorAmountGreaterThanTheLimitT3"
	case UnpaidItemsUnpaidItemsFailedNegativeFileCheckT4:
		return "UnpaidItemsUnpaidItemsFailedNegativeFileCheckT4"
	case DuplicateNumberDuplicateCheckNumberT5:
		return "DuplicateNumberDuplicateCheckNumberT5"
	case MicrErrorT6:
		return "MicrErrorT6"
	case TooManyChecksTooManyChecksOverMerchantOrBankLimitT7:
		return "TooManyChecksTooManyChecksOverMerchantOrBankLimitT7"
	case FailureVmDailyThresholdExceededV1:
		return "FailureVmDailyThresholdExceededV1"

	// MIP spec
	// case ApprovedHonorWithId08:
	// 	return "ApprovedHonorWithId08"
	// case PartialApproval10:
	// 	return "PartialApproval10"
	case FormatError30:
		return "FormatError30"
	case ContactCardIssuer70:
		return "ContactCardIssuer70"
	case PinNotChanged71:
		return "PinNotChanged71"
	// case InvalidToAccount76:
	// 	return "InvalidToAccount76"
	// case InvalidFromAccount77:
	// 	return "InvalidFromAccount77"
	// case InvalidAccount78:
	// 	return "InvalidAccount78"
	// case InvalidAuthLifeCycle84:
	// 	return "InvalidAuthLifeCycle84"
	case ApprovedPurchaseNoCashback87:
		return "ApprovedPurchaseNoCashback87"
	case CryptographicFailure88:
		return "CryptographicFailure88"
	case UnacceptablePin89:
		return "UnacceptablePin89"
		// case DuplicateTransmission94:
		// 	return "DuplicateTransmission94"
	}
	return ""
}

type RestrictedTicketIndicatorType rune

const (
	NoRestrictionSpace            RestrictedTicketIndicatorType = ' '
	NoRestriction                 RestrictedTicketIndicatorType = '0'
	RestrictedTicketNonRefundable RestrictedTicketIndicatorType = '1'
)

func (x RestrictedTicketIndicatorType) String() string {
	switch x {
	case NoRestrictionSpace:
		return "NoRestrictionSpace"
	case NoRestriction:
		return "NoRestriction"
	case RestrictedTicketNonRefundable:
		return "RestrictedTicketNonRefundable"
	}
	return ""
}

type ReturnAciType rune

const (
	SpaceOrEmptyAci                                     ReturnAciType = '0'
	CpsQualified                                        ReturnAciType = 'A'
	CpsQualifiedCardAcceptorDataWasSupplied             ReturnAciType = 'E'
	CpsQualifiedForVisaAccountFundingTxns               ReturnAciType = 'F'
	CpsQualifiedIncrementalAuthRequest                  ReturnAciType = 'I'
	CpsQualifiedForASelfServiceAutomatedFuelDispenser   ReturnAciType = 'C'
	CpsQualifiedIncludedAvsRequestUnableToReadMagstripe ReturnAciType = 'K'
	ReservedGermanDirectMarketing                       ReturnAciType = 'M'
	NotCpsQualified                                     ReturnAciType = 'N'
	CpsQualifiedAcceptedForPreferredCustomer            ReturnAciType = 'P'
	ReturnAciRecurringOrInstallmentPayments             ReturnAciType = 'R'
	CpsAttemptedForPreferredEcom3dSecure                ReturnAciType = 'S'
	CpsQualifiedForPreferredEcom3dSecure                ReturnAciType = 'U'
	CpsQualifiedIncludedAvsRequestInAuthRequest         ReturnAciType = 'V'
	CpsQualifiedForBasicEConnectionsNon3dSecure         ReturnAciType = 'W'
	NotEligibleToParticipateInCPS                       ReturnAciType = 'T'
	VasDowngradeOrCreditOrOffline                       ReturnAciType = ' '
)

func (x ReturnAciType) String() string {
	switch x {
	case SpaceOrEmptyAci:
		return "SpaceOrEmptyAci"
	case CpsQualified:
		return "CpsQualified"
	case CpsQualifiedCardAcceptorDataWasSupplied:
		return "CpsQualifiedCardAcceptorDataWasSupplied"
	case CpsQualifiedForVisaAccountFundingTxns:
		return "CpsQualifiedForVisaAccountFundingTxns"
	case CpsQualifiedIncrementalAuthRequest:
		return "CpsQualifiedIncrementalAuthRequest"
	case CpsQualifiedForASelfServiceAutomatedFuelDispenser:
		return "CpsQualifiedForASelfServiceAutomatedFuelDispenser"
	case CpsQualifiedIncludedAvsRequestUnableToReadMagstripe:
		return "CpsQualifiedIncludedAvsRequestUnableToReadMagstripe"
	case ReservedGermanDirectMarketing:
		return "ReservedGermanDirectMarketing"
	case NotCpsQualified:
		return "NotCpsQualified"
	case CpsQualifiedAcceptedForPreferredCustomer:
		return "CpsQualifiedAcceptedForPreferredCustomer"
	case ReturnAciRecurringOrInstallmentPayments:
		return "ReturnAciRecurringOrInstallmentPayments"
	case CpsAttemptedForPreferredEcom3dSecure:
		return "CpsAttemptedForPreferredEcom3dSecure"
	case CpsQualifiedForPreferredEcom3dSecure:
		return "CpsQualifiedForPreferredEcom3dSecure"
	case CpsQualifiedIncludedAvsRequestInAuthRequest:
		return "CpsQualifiedIncludedAvsRequestInAuthRequest"
	case CpsQualifiedForBasicEConnectionsNon3dSecure:
		return "CpsQualifiedForBasicEConnectionsNon3dSecure"
	case NotEligibleToParticipateInCPS:
		return "NotEligibleToParticipateInCPS"
	case VasDowngradeOrCreditOrOffline:
		return "VasDowngradeOrCreditOrOffline"
	}
	return ""
}

type ReversalTransmissionStatusType int

const (
	AuthorizationReversalNotTransmitted ReversalTransmissionStatusType = 0
	AuthorizationReversalTransmitted    ReversalTransmissionStatusType = 1
)

func (x ReversalTransmissionStatusType) String() string {
	switch x {
	case AuthorizationReversalNotTransmitted:
		return "AuthorizationReversalNotTransmitted"
	case AuthorizationReversalTransmitted:
		return "AuthorizationReversalTransmitted"
	}
	return ""
}

type ReceivingInstitutionIdType uint

const ( //Receiving Institution Identification
	SpaceOrEmptyReceivingInst  ReceivingInstitutionIdType = 0
	Bankserv                   ReceivingInstitutionIdType = 763057
	CertegyEquifaxEast         ReceivingInstitutionIdType = 894400
	CertegyEquifaxWest         ReceivingInstitutionIdType = 894300
	CiticorpPrivateLabel       ReceivingInstitutionIdType = 911111
	Efunds                     ReceivingInstitutionIdType = 762135
	IcsCbsNpc                  ReceivingInstitutionIdType = 810000
	RockyMountainRetailSystems ReceivingInstitutionIdType = 763060
	Scan                       ReceivingInstitutionIdType = 813500
	Telecheck                  ReceivingInstitutionIdType = 861400
)

func (x ReceivingInstitutionIdType) String() string {
	switch x {
	case SpaceOrEmptyReceivingInst:
		return "SpaceOrEmptyReceivingInst"
	case Bankserv:
		return "Bankserv"
	case CertegyEquifaxEast:
		return "CertegyEquifaxEast"
	case CertegyEquifaxWest:
		return "CertegyEquifaxWest"
	case CiticorpPrivateLabel:
		return "CiticorpPrivateLabel"
	case Efunds:
		return "Efunds"
	case IcsCbsNpc:
		return "IcsCbsNpc"
	case RockyMountainRetailSystems:
		return "RockyMountainRetailSystems"
	case Scan:
		return "Scan"
	case Telecheck:
		return "Telecheck"
	}
	return ""
}

type ServiceDevelopmentType rune

const (
	SpaceOrEmptyServiceDev  ServiceDevelopmentType = ' '
	Transponder             ServiceDevelopmentType = '5'
	RelationshipParticipant ServiceDevelopmentType = '6'
	DeferredBilling         ServiceDevelopmentType = '7'
)

func (x ServiceDevelopmentType) String() string {
	switch x {
	case SpaceOrEmptyServiceDev:
		return "SpaceOrEmptyServiceDev"
	case Transponder:
		return "Transponder"
	case RelationshipParticipant:
		return "RelationshipParticipant"
	case DeferredBilling:
		return "DeferredBilling"
	}
	return ""
}

type StateIdType uint

const (
	SpaceOrEmptyStateId                       StateIdType = 0
	Alabama                                   StateIdType = 1
	Alaska                                    StateIdType = 2
	Arizona                                   StateIdType = 4
	Arkansas                                  StateIdType = 5
	California                                StateIdType = 6
	Colorado                                  StateIdType = 8
	Connecticut                               StateIdType = 9
	CourtesyCard                              StateIdType = 90
	Delaware                                  StateIdType = 10
	DistrictOfColumbia                        StateIdType = 11
	Florida                                   StateIdType = 12
	GeorgiaState                              StateIdType = 13
	Hawaii                                    StateIdType = 15
	Idaho                                     StateIdType = 16
	Illinois                                  StateIdType = 17
	Indiana                                   StateIdType = 18
	Iowa                                      StateIdType = 19
	Kansas                                    StateIdType = 20
	Kentucky                                  StateIdType = 21
	Louisiana                                 StateIdType = 22
	Maine                                     StateIdType = 23
	Maryland                                  StateIdType = 24
	Massachusetts                             StateIdType = 25
	Michigan                                  StateIdType = 26
	MilitaryId                                StateIdType = 91
	Minnesota                                 StateIdType = 27
	Mississippi                               StateIdType = 28
	Missouri                                  StateIdType = 29
	Montana                                   StateIdType = 30
	Nevada                                    StateIdType = 32
	NewHampshire                              StateIdType = 33
	NewJersey                                 StateIdType = 34
	NewMexico                                 StateIdType = 35
	NewYork                                   StateIdType = 36
	NorthCarolina                             StateIdType = 37
	NorthDakota                               StateIdType = 38
	Ohio                                      StateIdType = 39
	Oklahoma                                  StateIdType = 40
	Oregon                                    StateIdType = 41
	Pennsylvania                              StateIdType = 42
	ProprietaryCard                           StateIdType = 93
	RhodeIsland                               StateIdType = 44
	SocialSecurityNumber                      StateIdType = 92
	SouthCarolina                             StateIdType = 45
	SouthDakota                               StateIdType = 46
	Tennessee                                 StateIdType = 47
	Texas                                     StateIdType = 48
	UsMilitaryBaseEmbassiesTravelingMerchants StateIdType = 99
	Utah                                      StateIdType = 49
	Vermont                                   StateIdType = 50
	Virginia                                  StateIdType = 51
	Washington                                StateIdType = 53
	WestVirginia                              StateIdType = 54
	Wisconsin                                 StateIdType = 55
	Wyoming                                   StateIdType = 56
)

func (x StateIdType) String() string {
	switch x {
	case SpaceOrEmptyStateId:
		return "SpaceOrEmptyStateId"
	case Alabama:
		return "Alabama"
	case Alaska:
		return "Alaska"
	case Arizona:
		return "Arizona"
	case Arkansas:
		return "Arkansas"
	case California:
		return "California"
	case Colorado:
		return "Colorado"
	case Connecticut:
		return "Connecticut"
	case CourtesyCard:
		return "CourtesyCard"
	case Delaware:
		return "Delaware"
	case DistrictOfColumbia:
		return "DistrictOfColumbia"
	case Florida:
		return "Florida"
	case GeorgiaState:
		return "GeorgiaState"
	case Hawaii:
		return "Hawaii"
	case Idaho:
		return "Idaho"
	case Illinois:
		return "Illinois"
	case Indiana:
		return "Indiana"
	case Iowa:
		return "Iowa"
	case Kansas:
		return "Kansas"
	case Kentucky:
		return "Kentucky"
	case Louisiana:
		return "Louisiana"
	case Maine:
		return "Maine"
	case Maryland:
		return "Maryland"
	case Massachusetts:
		return "Massachusetts"
	case Michigan:
		return "Michigan"
	case MilitaryId:
		return "MilitaryId"
	case Minnesota:
		return "Minnesota"
	case Mississippi:
		return "Mississippi"
	case Missouri:
		return "Missouri"
	case Montana:
		return "Montana"
	case Nevada:
		return "Nevada"
	case NewHampshire:
		return "NewHampshire"
	case NewJersey:
		return "NewJersey"
	case NewMexico:
		return "NewMexico"
	case NewYork:
		return "NewYork"
	case NorthCarolina:
		return "NorthCarolina"
	case NorthDakota:
		return "NorthDakota"
	case Ohio:
		return "Ohio"
	case Oklahoma:
		return "Oklahoma"
	case Oregon:
		return "Oregon"
	case Pennsylvania:
		return "Pennsylvania"
	case ProprietaryCard:
		return "ProprietaryCard"
	case RhodeIsland:
		return "RhodeIsland"
	case SocialSecurityNumber:
		return "SocialSecurityNumber"
	case SouthCarolina:
		return "SouthCarolina"
	case SouthDakota:
		return "SouthDakota"
	case Tennessee:
		return "Tennessee"
	case Texas:
		return "Texas"
	case UsMilitaryBaseEmbassiesTravelingMerchants:
		return "UsMilitaryBaseEmbassiesTravelingMerchants"
	case Utah:
		return "Utah"
	case Vermont:
		return "Vermont"
	case Virginia:
		return "Virginia"
	case Washington:
		return "Washington"
	case WestVirginia:
		return "WestVirginia"
	case Wisconsin:
		return "Wisconsin"
	case Wyoming:
		return "Wyoming"
	}
	return ""
}

type StopOverCodeType rune

const (
	StopoverCodeNotApplicable StopOverCodeType = ' '
	StopOverAllowed           StopOverCodeType = 'O'
	StopOverNotAllowed        StopOverCodeType = 'X'
)

func (x StopOverCodeType) String() string {
	switch x {
	case StopoverCodeNotApplicable:
		return "StopoverCodeNotApplicable"
	case StopOverAllowed:
		return "StopOverAllowed"
	case StopOverNotAllowed:
		return "StopOverNotAllowed"
	}
	return ""
}

type TransactionCodeType string

const (
	Purchase                                          TransactionCodeType = "54"
	PurchaseRepeat                                    TransactionCodeType = "64"
	CashAdvance                                       TransactionCodeType = "55" // Use in Banks and Financial Institutions Only (Visa, MasterCard, Discover only).
	CashAdvanceRepeat                                 TransactionCodeType = "65"
	PurchaseCardNotPresent                            TransactionCodeType = "56" // Used primarily in Direct Marketing, Mail, and Telephone Order Environments
	PurchaseCardNotPresentRepeat                      TransactionCodeType = "66"
	QuasiCash                                         TransactionCodeType = "57" // Used for purchase of semi-cash items such as casino chips, travelers checks, stamps. (Visa and MasterCard)
	QuasiCashRepeat                                   TransactionCodeType = "67"
	CardAuthentication                                TransactionCodeType = "58" // To determine if an account is open for use or to verify cardholder information. Transaction amounts must be zero filled. AVS and CVV2 may be attempted. (Visa, MasterCard, American Express, Discover, PayPal only)
	CardAuthenticationRepeat                          TransactionCodeType = "68"
	OnlineAuthorizationReversal                       TransactionCodeType = "59" // Used to reverse a completed credit authorization prior to batch settlement.
	OnlineAuthorizationReversalRepeat                 TransactionCodeType = "69"
	StoreAndForwardAuthorizationReversal              TransactionCodeType = "5A" // Same as code “59,” except the reversal is being submitted after batch settlement occurred
	StoreAndForwardAuthorizationReversalRepeat        TransactionCodeType = "6A"
	BillPayTransaction                                TransactionCodeType = "5B" // By Credit Card only
	BillPayTransactionRepeat                          TransactionCodeType = "6B"
	CreditAdvice                                      TransactionCodeType = "5C" // Advice message for an AFD final sale amount (Visa and MasterCard only)
	CreditAccountFundingOrPayment                     TransactionCodeType = "5G" // Credit card used for payment of account to account transactions
	CreditAccountFundingOrPaymentRepeat               TransactionCodeType = "6G" // Credit Account Funding Transaction (Visa), Credit Payment Transaction (MC)
	CardNotPresentCreditAccountFundingOrPayment       TransactionCodeType = "5H" // Card-not-present credit payment of account to account transactions
	CardNotPresentCreditAccountFundingOrPaymentRepeat TransactionCodeType = "6H"
	CardPresentCreditCardholderFundsTransfer          TransactionCodeType = "5J" // Card-present credit payment of cardholder funds transfer
	CardPresentCreditCardholderFundsTransferRepeat    TransactionCodeType = "6J"
	CardholderFundsTransferCardNotPresent             TransactionCodeType = "5K"
	CardholderFundsTransferCardNotPresentRepeat       TransactionCodeType = "6K"
	CreditStoredValueBalanceInquiry                   TransactionCodeType = "5L" // Balance Inquiry
	CreditStoredValueBalanceInquiryRepeat             TransactionCodeType = "6L"
	HealthcareEligibilityInquiry                      TransactionCodeType = "5M" // Visa only
	HealthcareEligibilityInquiryRepeat                TransactionCodeType = "6M"
	BalanceInquiryReversal                            TransactionCodeType = "5N" // Used to reverse a balance inquiry transaction (MasterCard only)
	BalanceInquiryReversalRepeat                      TransactionCodeType = "6N"
	ProductEligibilityInquiry                         TransactionCodeType = "5P" // Visa and Mastercard only
	TokenRequest                                      TransactionCodeType = "5T" // This indicates a request for Token only. No card verification will be performed. (Group III version 053 is needed)
	CheckGuarantee                                    TransactionCodeType = "70" // Must specify check service in RIID Field
	PosCheckConversionWithGuarantee                   TransactionCodeType = "76"
	PosCheckConversionWithVerification                TransactionCodeType = "77"
	PosCheckConversionOnly                            TransactionCodeType = "78"
	PosCheckReversalConversionWithGuarantee           TransactionCodeType = "79"
	PosCheckReversalConversionWithVerification        TransactionCodeType = "7A"
	PosCheckReversalConversionOnly                    TransactionCodeType = "7B"
	PrivateLabelPurchase                              TransactionCodeType = "84"
	PrivateLabelCashAdvance                           TransactionCodeType = "85"
	PrivateLabelCardNotPresent                        TransactionCodeType = "86"
	PrivateLabelQuasiCash                             TransactionCodeType = "87"
	PrivateLabelCardAuthentication                    TransactionCodeType = "88"
	FoodStampsReturnEbt                               TransactionCodeType = "92"
	DirectDebitPurchase                               TransactionCodeType = "93"
	DirectDebitPurchaseReturn                         TransactionCodeType = "94"
	CashBenefitsCashWithdrawalEbt                     TransactionCodeType = "96"
	FoodStampPurchaseEbt                              TransactionCodeType = "98"
	DirectDebitBalanceInquiry                         TransactionCodeType = "9A"

	DebitBillPaymentTransaction        TransactionCodeType = "9B"
	PinlessDebitBillPaymentTransaction TransactionCodeType = "9C" // Debit transactions without a PIN. The Network ID of 0000 can not be used.

	FoodStampsElectronicVoucherEbt                     TransactionCodeType = "9E"
	EbtCashBenefitsPurchaseOrPurchaseWithCashBack      TransactionCodeType = "9F"
	DebitAccountFundingPurchase                        TransactionCodeType = "9G"
	DebitAccountFundingReturn                          TransactionCodeType = "9H"
	DebitCardholderFundsTransfer                       TransactionCodeType = "9J"
	DebitFundsTransferReturn                           TransactionCodeType = "9K"
	EbtFoodStampBalanceInquiry                         TransactionCodeType = "9L"
	EbtCashBenefitsBalanceInquiry                      TransactionCodeType = "9M"
	ChipCardTransactionAdviceRecordLimitedAvailability TransactionCodeType = "A0"
	AutomaticReversalDirectDebitPurchase               TransactionCodeType = "A3"
	AutomaticReversalDirectDebitPurchaseReturn         TransactionCodeType = "A4"
	AutomaticReversalInterlinkDirectDebitCancel        TransactionCodeType = "A5"

	ATMCashDisbursement                  TransactionCodeType = "B1"
	ATMBalanceInquiry                    TransactionCodeType = "B2"
	ATMDeposit                           TransactionCodeType = "B3"
	ATMCardholderAccountTransfer         TransactionCodeType = "B4"
	ATMCashDisbursementReversal          TransactionCodeType = "C1"
	ATMDepositReversal                   TransactionCodeType = "C3"
	ATMCardholderAccountTransferReversal TransactionCodeType = "C4"
	ATMAdjustmentUpCredit                TransactionCodeType = "C8"
	ATMAdjustmentDownDebit               TransactionCodeType = "C9"
	GiftCardCloseCard                    TransactionCodeType = "G2"
	GiftCardBalanceInquiry               TransactionCodeType = "G3"
	GiftCardPurchaseRedemption           TransactionCodeType = "G4"
	GiftCardReturnRefund                 TransactionCodeType = "G5"
	GiftCardAddValueLoadCard             TransactionCodeType = "G6"
	GiftCardDecreaseValueUnloadCard      TransactionCodeType = "G7"
	GiftCardStandAloneTip                TransactionCodeType = "GB"
	GiftCardIssueGiftCard                TransactionCodeType = "GC"
	GiftCardIssueVirtualGiftCard         TransactionCodeType = "GD"
	GiftCardMerchantInitiatedCancel      TransactionCodeType = "GE"
	GiftCardMerchantInitiatedReversal    TransactionCodeType = "GF"
	GiftCardCashBack                     TransactionCodeType = "GG"
	Q1PrepaidCardActivation              TransactionCodeType = "P1"
	Q2PrepaidCardActivationReversal      TransactionCodeType = "P2"
	Q3PrepaidCardLoad                    TransactionCodeType = "P3"
	Q4PrepaidCardLoadReversal            TransactionCodeType = "P4"
	PrepaidCardActivation                TransactionCodeType = "R1"
	PrepaidCardActivationReversal        TransactionCodeType = "R2"
	PrepaidCardLoad                      TransactionCodeType = "R3"
	PrepaidCardLoadReversal              TransactionCodeType = "R4"
	TerminalAuthentication               TransactionCodeType = "TA" // Used to authenticate a device using Group III version 49.
	TerminalDeactivation                 TransactionCodeType = "TD" // Deactivates a device previously authenticated; authenticate again after deactivation

	CreditReturn                        TransactionCodeType = "CR"
	AccountFundingCreditReturn          TransactionCodeType = "FR"
	CardholderFundsTransferCreditReturn TransactionCodeType = "TR"
)

func (x TransactionCodeType) String() string {
	switch x {
	case Purchase:
		return "Purchase"
	case PurchaseRepeat:
		return "PurchaseRepeat"
	case CashAdvance:
		return "CashAdvance"
	case CashAdvanceRepeat:
		return "CashAdvanceRepeat"
	case PurchaseCardNotPresent:
		return "PurchaseCardNotPresent"
	case PurchaseCardNotPresentRepeat:
		return "PurchaseCardNotPresentRepeat"
	case QuasiCash:
		return "QuasiCash"
	case QuasiCashRepeat:
		return "QuasiCashRepeat"
	case CardAuthentication:
		return "CardAuthentication"
	case CardAuthenticationRepeat:
		return "CardAuthenticationRepeat"
	case OnlineAuthorizationReversal:
		return "OnlineAuthorizationReversal"
	case OnlineAuthorizationReversalRepeat:
		return "OnlineAuthorizationReversalRepeat"
	case StoreAndForwardAuthorizationReversal:
		return "StoreAndForwardAuthorizationReversal"
	case StoreAndForwardAuthorizationReversalRepeat:
		return "StoreAndForwardAuthorizationReversalRepeat"
	case BillPayTransaction:
		return "BillPayTransaction"
	case BillPayTransactionRepeat:
		return "BillPayTransactionRepeat"
	case CreditAdvice:
		return "CreditAdvice"
	case CreditAccountFundingOrPayment:
		return "CreditAccountFundingOrPayment"
	case CreditAccountFundingOrPaymentRepeat:
		return "CreditAccountFundingOrPaymentRepeat"
	case CardNotPresentCreditAccountFundingOrPayment:
		return "CardNotPresentCreditAccountFundingOrPayment"
	case CardNotPresentCreditAccountFundingOrPaymentRepeat:
		return "CardNotPresentCreditAccountFundingOrPaymentRepeat"
	case CardPresentCreditCardholderFundsTransfer:
		return "CardPresentCreditCardholderFundsTransfer"
	case CardPresentCreditCardholderFundsTransferRepeat:
		return "CardPresentCreditCardholderFundsTransferRepeat"
	case CardholderFundsTransferCardNotPresent:
		return "CardholderFundsTransferCardNotPresent"
	case CardholderFundsTransferCardNotPresentRepeat:
		return "CardholderFundsTransferCardNotPresentRepeat"
	case CreditStoredValueBalanceInquiry:
		return "CreditStoredValueBalanceInquiry"
	case CreditStoredValueBalanceInquiryRepeat:
		return "CreditStoredValueBalanceInquiryRepeat"
	case HealthcareEligibilityInquiry:
		return "HealthcareEligibilityInquiry"
	case HealthcareEligibilityInquiryRepeat:
		return "HealthcareEligibilityInquiryRepeat"
	case BalanceInquiryReversal:
		return "BalanceInquiryReversal"
	case BalanceInquiryReversalRepeat:
		return "BalanceInquiryReversalRepeat"
	case ProductEligibilityInquiry:
		return "ProductEligibilityInquiry"
	case TokenRequest:
		return "TokenRequest"
	case CheckGuarantee:
		return "CheckGuarantee"
	case PosCheckConversionWithGuarantee:
		return "PosCheckConversionWithGuarantee"
	case PosCheckConversionWithVerification:
		return "PosCheckConversionWithVerification"
	case PosCheckConversionOnly:
		return "PosCheckConversionOnly"
	case PosCheckReversalConversionWithGuarantee:
		return "PosCheckReversalConversionWithGuarantee"
	case PosCheckReversalConversionWithVerification:
		return "PosCheckReversalConversionWithVerification"
	case PosCheckReversalConversionOnly:
		return "PosCheckReversalConversionOnly"
	case PrivateLabelPurchase:
		return "PrivateLabelPurchase"
	case PrivateLabelCashAdvance:
		return "PrivateLabelCashAdvance"
	case PrivateLabelCardNotPresent:
		return "PrivateLabelCardNotPresent"
	case PrivateLabelQuasiCash:
		return "PrivateLabelQuasiCash"
	case PrivateLabelCardAuthentication:
		return "PrivateLabelCardAuthentication"
	case FoodStampsReturnEbt:
		return "FoodStampsReturnEbt"
	case DirectDebitPurchase:
		return "DirectDebitPurchase"
	case DirectDebitPurchaseReturn:
		return "DirectDebitPurchaseReturn"
	case CashBenefitsCashWithdrawalEbt:
		return "CashBenefitsCashWithdrawalEbt"
	case FoodStampPurchaseEbt:
		return "FoodStampPurchaseEbt"
	case DirectDebitBalanceInquiry:
		return "DirectDebitBalanceInquiry"

	case DebitBillPaymentTransaction:
		return "DebitBillPaymentTransaction"
	case PinlessDebitBillPaymentTransaction:
		return "PinlessDebitBillPaymentTransaction"

	case FoodStampsElectronicVoucherEbt:
		return "FoodStampsElectronicVoucherEbt"
	case EbtCashBenefitsPurchaseOrPurchaseWithCashBack:
		return "EbtCashBenefitsPurchaseOrPurchaseWithCashBack"
	case DebitAccountFundingPurchase:
		return "DebitAccountFundingPurchase"
	case DebitAccountFundingReturn:
		return "DebitAccountFundingReturn"
	case DebitCardholderFundsTransfer:
		return "DebitCardholderFundsTransfer"
	case DebitFundsTransferReturn:
		return "DebitFundsTransferReturn"
	case EbtFoodStampBalanceInquiry:
		return "EbtFoodStampBalanceInquiry"
	case EbtCashBenefitsBalanceInquiry:
		return "EbtCashBenefitsBalanceInquiry"
	case ChipCardTransactionAdviceRecordLimitedAvailability:
		return "ChipCardTransactionAdviceRecordLimitedAvailability"
	case AutomaticReversalDirectDebitPurchase:
		return "AutomaticReversalDirectDebitPurchase"
	case AutomaticReversalDirectDebitPurchaseReturn:
		return "AutomaticReversalDirectDebitPurchaseReturn"
	case AutomaticReversalInterlinkDirectDebitCancel:
		return "AutomaticReversalInterlinkDirectDebitCancel"

	case ATMCashDisbursement:
		return "ATMCashDisbursement"
	case ATMBalanceInquiry:
		return "ATMBalanceInquiry"
	case ATMDeposit:
		return "ATMDeposit"
	case ATMCardholderAccountTransfer:
		return "ATMCardholderAccountTransfer"
	case ATMCashDisbursementReversal:
		return "ATMCashDisbursementReversal"
	case ATMDepositReversal:
		return "ATMDepositReversal"
	case ATMCardholderAccountTransferReversal:
		return "ATMCardholderAccountTransferReversal"
	case ATMAdjustmentUpCredit:
		return "ATMAdjustmentUpCredit"
	case ATMAdjustmentDownDebit:
		return "ATMAdjustmentDownDebit"
	case GiftCardCloseCard:
		return "GiftCardCloseCard"
	case GiftCardBalanceInquiry:
		return "GiftCardBalanceInquiry"
	case GiftCardPurchaseRedemption:
		return "GiftCardPurchaseRedemption"
	case GiftCardReturnRefund:
		return "GiftCardReturnRefund"
	case GiftCardAddValueLoadCard:
		return "GiftCardAddValueLoadCard"
	case GiftCardDecreaseValueUnloadCard:
		return "GiftCardDecreaseValueUnloadCard"
	case GiftCardStandAloneTip:
		return "GiftCardStandAloneTip"
	case GiftCardIssueGiftCard:
		return "GiftCardIssueGiftCard"
	case GiftCardIssueVirtualGiftCard:
		return "GiftCardIssueVirtualGiftCard"
	case GiftCardMerchantInitiatedCancel:
		return "GiftCardMerchantInitiatedCancel"
	case GiftCardMerchantInitiatedReversal:
		return "GiftCardMerchantInitiatedReversal"
	case GiftCardCashBack:
		return "GiftCardCashBack"
	case Q1PrepaidCardActivation:
		return "Q1PrepaidCardActivation"
	case Q2PrepaidCardActivationReversal:
		return "Q2PrepaidCardActivationReversal"
	case Q3PrepaidCardLoad:
		return "Q3PrepaidCardLoad"
	case Q4PrepaidCardLoadReversal:
		return "Q4PrepaidCardLoadReversal"
	case PrepaidCardActivation:
		return "PrepaidCardActivation"
	case PrepaidCardActivationReversal:
		return "PrepaidCardActivationReversal"
	case PrepaidCardLoad:
		return "PrepaidCardLoad"
	case PrepaidCardLoadReversal:
		return "PrepaidCardLoadReversal"
	case TerminalAuthentication:
		return "TerminalAuthentication"
	case TerminalDeactivation:
		return "TerminalDeactivation"

	case CreditReturn:
		return "CreditReturn"
	case AccountFundingCreditReturn:
		return "AccountFundingCreditReturn"
	case CardholderFundsTransferCreditReturn:
		return "CardholderFundsTransferCreditReturn"
	}
	return ""
}

type TransactionSecurityIndicatorType rune

const (
	CardAccountNumberIsEncrypted TransactionSecurityIndicatorType = 'E'
	CardAccountNumberIsToken     TransactionSecurityIndicatorType = 'T'
)

func (x TransactionSecurityIndicatorType) String() string {
	switch x {
	case CardAccountNumberIsEncrypted:
		return "CardAccountNumberIsEncrypted"
	case CardAccountNumberIsToken:
		return "CardAccountNumberIsToken"
	}
	return ""
}

type UcafCollectionType rune

const (
	UcafNotSupportedByMerchantWebsite UcafCollectionType = '0'
	UcafSupportedButDataNotPopulated  UcafCollectionType = '1'
	UcafDataWasPopulated              UcafCollectionType = '2'
	VisaCavvDataPresent               UcafCollectionType = '4'
)

func (x UcafCollectionType) String() string {
	switch x {
	case UcafNotSupportedByMerchantWebsite:
		return "UcafNotSupportedByMerchantWebsite"
	case UcafSupportedButDataNotPopulated:
		return "UcafSupportedButDataNotPopulated"
	case UcafDataWasPopulated:
		return "UcafDataWasPopulated"
	case VisaCavvDataPresent:
		return "VisaCavvDataPresent"
	}
	return ""
}

type UnitMeasureType rune

const (
	Litre          UnitMeasureType = 'L'
	UsGallon       UnitMeasureType = 'G'
	ImperialGallon UnitMeasureType = 'I'
	Kilo           UnitMeasureType = 'K'
	Pound          UnitMeasureType = 'P'
)

func (x UnitMeasureType) String() string {
	switch x {
	case Litre:
		return "Litre"
	case UsGallon:
		return "UsGallon"
	case ImperialGallon:
		return "ImperialGallon"
	case Kilo:
		return "Kilo"
	case Pound:
		return "Pound"
	}
	return ""
}

type VoidIndicatorType rune

const (
	Voided    VoidIndicatorType = 'V'
	NotVoided VoidIndicatorType = ' '
)

func (x VoidIndicatorType) String() string {
	switch x {
	case Voided:
		return "Voided"
	case NotVoided:
		return "NotVoided"
	}
	return ""
}

type X25RoutingType rune

const (
	Primary X25RoutingType = 'Z'
)

func (x X25RoutingType) String() string {
	switch x {
	case Primary:
		return "Primary"
	}
	return ""
}
