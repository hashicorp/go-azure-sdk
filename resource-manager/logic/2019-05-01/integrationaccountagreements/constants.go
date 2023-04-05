package integrationaccountagreements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgreementType string

const (
	AgreementTypeASTwo        AgreementType = "AS2"
	AgreementTypeEdifact      AgreementType = "Edifact"
	AgreementTypeNotSpecified AgreementType = "NotSpecified"
	AgreementTypeXOneTwo      AgreementType = "X12"
)

func PossibleValuesForAgreementType() []string {
	return []string{
		string(AgreementTypeASTwo),
		string(AgreementTypeEdifact),
		string(AgreementTypeNotSpecified),
		string(AgreementTypeXOneTwo),
	}
}

type EdifactCharacterSet string

const (
	EdifactCharacterSetKECA         EdifactCharacterSet = "KECA"
	EdifactCharacterSetNotSpecified EdifactCharacterSet = "NotSpecified"
	EdifactCharacterSetUNOA         EdifactCharacterSet = "UNOA"
	EdifactCharacterSetUNOB         EdifactCharacterSet = "UNOB"
	EdifactCharacterSetUNOC         EdifactCharacterSet = "UNOC"
	EdifactCharacterSetUNOD         EdifactCharacterSet = "UNOD"
	EdifactCharacterSetUNOE         EdifactCharacterSet = "UNOE"
	EdifactCharacterSetUNOF         EdifactCharacterSet = "UNOF"
	EdifactCharacterSetUNOG         EdifactCharacterSet = "UNOG"
	EdifactCharacterSetUNOH         EdifactCharacterSet = "UNOH"
	EdifactCharacterSetUNOI         EdifactCharacterSet = "UNOI"
	EdifactCharacterSetUNOJ         EdifactCharacterSet = "UNOJ"
	EdifactCharacterSetUNOK         EdifactCharacterSet = "UNOK"
	EdifactCharacterSetUNOX         EdifactCharacterSet = "UNOX"
	EdifactCharacterSetUNOY         EdifactCharacterSet = "UNOY"
)

func PossibleValuesForEdifactCharacterSet() []string {
	return []string{
		string(EdifactCharacterSetKECA),
		string(EdifactCharacterSetNotSpecified),
		string(EdifactCharacterSetUNOA),
		string(EdifactCharacterSetUNOB),
		string(EdifactCharacterSetUNOC),
		string(EdifactCharacterSetUNOD),
		string(EdifactCharacterSetUNOE),
		string(EdifactCharacterSetUNOF),
		string(EdifactCharacterSetUNOG),
		string(EdifactCharacterSetUNOH),
		string(EdifactCharacterSetUNOI),
		string(EdifactCharacterSetUNOJ),
		string(EdifactCharacterSetUNOK),
		string(EdifactCharacterSetUNOX),
		string(EdifactCharacterSetUNOY),
	}
}

type EdifactDecimalIndicator string

const (
	EdifactDecimalIndicatorComma        EdifactDecimalIndicator = "Comma"
	EdifactDecimalIndicatorDecimal      EdifactDecimalIndicator = "Decimal"
	EdifactDecimalIndicatorNotSpecified EdifactDecimalIndicator = "NotSpecified"
)

func PossibleValuesForEdifactDecimalIndicator() []string {
	return []string{
		string(EdifactDecimalIndicatorComma),
		string(EdifactDecimalIndicatorDecimal),
		string(EdifactDecimalIndicatorNotSpecified),
	}
}

type EncryptionAlgorithm string

const (
	EncryptionAlgorithmAESOneNineTwo  EncryptionAlgorithm = "AES192"
	EncryptionAlgorithmAESOneTwoEight EncryptionAlgorithm = "AES128"
	EncryptionAlgorithmAESTwoFiveSix  EncryptionAlgorithm = "AES256"
	EncryptionAlgorithmDESThree       EncryptionAlgorithm = "DES3"
	EncryptionAlgorithmNone           EncryptionAlgorithm = "None"
	EncryptionAlgorithmNotSpecified   EncryptionAlgorithm = "NotSpecified"
	EncryptionAlgorithmRCTwo          EncryptionAlgorithm = "RC2"
)

func PossibleValuesForEncryptionAlgorithm() []string {
	return []string{
		string(EncryptionAlgorithmAESOneNineTwo),
		string(EncryptionAlgorithmAESOneTwoEight),
		string(EncryptionAlgorithmAESTwoFiveSix),
		string(EncryptionAlgorithmDESThree),
		string(EncryptionAlgorithmNone),
		string(EncryptionAlgorithmNotSpecified),
		string(EncryptionAlgorithmRCTwo),
	}
}

type HashingAlgorithm string

const (
	HashingAlgorithmMDFive               HashingAlgorithm = "MD5"
	HashingAlgorithmNone                 HashingAlgorithm = "None"
	HashingAlgorithmNotSpecified         HashingAlgorithm = "NotSpecified"
	HashingAlgorithmSHAOne               HashingAlgorithm = "SHA1"
	HashingAlgorithmSHATwoFiveOneTwo     HashingAlgorithm = "SHA2512"
	HashingAlgorithmSHATwoThreeEightFour HashingAlgorithm = "SHA2384"
	HashingAlgorithmSHATwoTwoFiveSix     HashingAlgorithm = "SHA2256"
)

func PossibleValuesForHashingAlgorithm() []string {
	return []string{
		string(HashingAlgorithmMDFive),
		string(HashingAlgorithmNone),
		string(HashingAlgorithmNotSpecified),
		string(HashingAlgorithmSHAOne),
		string(HashingAlgorithmSHATwoFiveOneTwo),
		string(HashingAlgorithmSHATwoThreeEightFour),
		string(HashingAlgorithmSHATwoTwoFiveSix),
	}
}

type KeyType string

const (
	KeyTypeNotSpecified KeyType = "NotSpecified"
	KeyTypePrimary      KeyType = "Primary"
	KeyTypeSecondary    KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypeNotSpecified),
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

type MessageFilterType string

const (
	MessageFilterTypeExclude      MessageFilterType = "Exclude"
	MessageFilterTypeInclude      MessageFilterType = "Include"
	MessageFilterTypeNotSpecified MessageFilterType = "NotSpecified"
)

func PossibleValuesForMessageFilterType() []string {
	return []string{
		string(MessageFilterTypeExclude),
		string(MessageFilterTypeInclude),
		string(MessageFilterTypeNotSpecified),
	}
}

type SegmentTerminatorSuffix string

const (
	SegmentTerminatorSuffixCR           SegmentTerminatorSuffix = "CR"
	SegmentTerminatorSuffixCRLF         SegmentTerminatorSuffix = "CRLF"
	SegmentTerminatorSuffixLF           SegmentTerminatorSuffix = "LF"
	SegmentTerminatorSuffixNone         SegmentTerminatorSuffix = "None"
	SegmentTerminatorSuffixNotSpecified SegmentTerminatorSuffix = "NotSpecified"
)

func PossibleValuesForSegmentTerminatorSuffix() []string {
	return []string{
		string(SegmentTerminatorSuffixCR),
		string(SegmentTerminatorSuffixCRLF),
		string(SegmentTerminatorSuffixLF),
		string(SegmentTerminatorSuffixNone),
		string(SegmentTerminatorSuffixNotSpecified),
	}
}

type SigningAlgorithm string

const (
	SigningAlgorithmDefault              SigningAlgorithm = "Default"
	SigningAlgorithmNotSpecified         SigningAlgorithm = "NotSpecified"
	SigningAlgorithmSHAOne               SigningAlgorithm = "SHA1"
	SigningAlgorithmSHATwoFiveOneTwo     SigningAlgorithm = "SHA2512"
	SigningAlgorithmSHATwoThreeEightFour SigningAlgorithm = "SHA2384"
	SigningAlgorithmSHATwoTwoFiveSix     SigningAlgorithm = "SHA2256"
)

func PossibleValuesForSigningAlgorithm() []string {
	return []string{
		string(SigningAlgorithmDefault),
		string(SigningAlgorithmNotSpecified),
		string(SigningAlgorithmSHAOne),
		string(SigningAlgorithmSHATwoFiveOneTwo),
		string(SigningAlgorithmSHATwoThreeEightFour),
		string(SigningAlgorithmSHATwoTwoFiveSix),
	}
}

type TrailingSeparatorPolicy string

const (
	TrailingSeparatorPolicyMandatory    TrailingSeparatorPolicy = "Mandatory"
	TrailingSeparatorPolicyNotAllowed   TrailingSeparatorPolicy = "NotAllowed"
	TrailingSeparatorPolicyNotSpecified TrailingSeparatorPolicy = "NotSpecified"
	TrailingSeparatorPolicyOptional     TrailingSeparatorPolicy = "Optional"
)

func PossibleValuesForTrailingSeparatorPolicy() []string {
	return []string{
		string(TrailingSeparatorPolicyMandatory),
		string(TrailingSeparatorPolicyNotAllowed),
		string(TrailingSeparatorPolicyNotSpecified),
		string(TrailingSeparatorPolicyOptional),
	}
}

type UsageIndicator string

const (
	UsageIndicatorInformation  UsageIndicator = "Information"
	UsageIndicatorNotSpecified UsageIndicator = "NotSpecified"
	UsageIndicatorProduction   UsageIndicator = "Production"
	UsageIndicatorTest         UsageIndicator = "Test"
)

func PossibleValuesForUsageIndicator() []string {
	return []string{
		string(UsageIndicatorInformation),
		string(UsageIndicatorNotSpecified),
		string(UsageIndicatorProduction),
		string(UsageIndicatorTest),
	}
}

type X12CharacterSet string

const (
	X12CharacterSetBasic        X12CharacterSet = "Basic"
	X12CharacterSetExtended     X12CharacterSet = "Extended"
	X12CharacterSetNotSpecified X12CharacterSet = "NotSpecified"
	X12CharacterSetUTFEight     X12CharacterSet = "UTF8"
)

func PossibleValuesForX12CharacterSet() []string {
	return []string{
		string(X12CharacterSetBasic),
		string(X12CharacterSetExtended),
		string(X12CharacterSetNotSpecified),
		string(X12CharacterSetUTFEight),
	}
}

type X12DateFormat string

const (
	X12DateFormatCCYYMMDD     X12DateFormat = "CCYYMMDD"
	X12DateFormatNotSpecified X12DateFormat = "NotSpecified"
	X12DateFormatYYMMDD       X12DateFormat = "YYMMDD"
)

func PossibleValuesForX12DateFormat() []string {
	return []string{
		string(X12DateFormatCCYYMMDD),
		string(X12DateFormatNotSpecified),
		string(X12DateFormatYYMMDD),
	}
}

type X12TimeFormat string

const (
	X12TimeFormatHHMM         X12TimeFormat = "HHMM"
	X12TimeFormatHHMMSS       X12TimeFormat = "HHMMSS"
	X12TimeFormatHHMMSSd      X12TimeFormat = "HHMMSSd"
	X12TimeFormatHHMMSSdd     X12TimeFormat = "HHMMSSdd"
	X12TimeFormatNotSpecified X12TimeFormat = "NotSpecified"
)

func PossibleValuesForX12TimeFormat() []string {
	return []string{
		string(X12TimeFormatHHMM),
		string(X12TimeFormatHHMMSS),
		string(X12TimeFormatHHMMSSd),
		string(X12TimeFormatHHMMSSdd),
		string(X12TimeFormatNotSpecified),
	}
}
