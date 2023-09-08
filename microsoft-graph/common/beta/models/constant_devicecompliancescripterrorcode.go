package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptErrorCode string

const (
	DeviceComplianceScriptErrorCodedatatypeMissing                         DeviceComplianceScriptErrorCode = "DatatypeMissing"
	DeviceComplianceScriptErrorCodedatatypeNotSupported                    DeviceComplianceScriptErrorCode = "DatatypeNotSupported"
	DeviceComplianceScriptErrorCodedescriptionInvalid                      DeviceComplianceScriptErrorCode = "DescriptionInvalid"
	DeviceComplianceScriptErrorCodedescriptionMissing                      DeviceComplianceScriptErrorCode = "DescriptionMissing"
	DeviceComplianceScriptErrorCodedescriptionTooLarge                     DeviceComplianceScriptErrorCode = "DescriptionTooLarge"
	DeviceComplianceScriptErrorCodeduplicateLocales                        DeviceComplianceScriptErrorCode = "DuplicateLocales"
	DeviceComplianceScriptErrorCodeduplicateRules                          DeviceComplianceScriptErrorCode = "DuplicateRules"
	DeviceComplianceScriptErrorCodeenglishLocaleMissing                    DeviceComplianceScriptErrorCode = "EnglishLocaleMissing"
	DeviceComplianceScriptErrorCodejsonFileInvalid                         DeviceComplianceScriptErrorCode = "JsonFileInvalid"
	DeviceComplianceScriptErrorCodejsonFileMissing                         DeviceComplianceScriptErrorCode = "JsonFileMissing"
	DeviceComplianceScriptErrorCodejsonFileTooLarge                        DeviceComplianceScriptErrorCode = "JsonFileTooLarge"
	DeviceComplianceScriptErrorCodemoreInfoUriInvalid                      DeviceComplianceScriptErrorCode = "MoreInfoUriInvalid"
	DeviceComplianceScriptErrorCodemoreInfoUriMissing                      DeviceComplianceScriptErrorCode = "MoreInfoUriMissing"
	DeviceComplianceScriptErrorCodemoreInfoUriTooLarge                     DeviceComplianceScriptErrorCode = "MoreInfoUriTooLarge"
	DeviceComplianceScriptErrorCodenone                                    DeviceComplianceScriptErrorCode = "None"
	DeviceComplianceScriptErrorCodeoperandInvalid                          DeviceComplianceScriptErrorCode = "OperandInvalid"
	DeviceComplianceScriptErrorCodeoperandMissing                          DeviceComplianceScriptErrorCode = "OperandMissing"
	DeviceComplianceScriptErrorCodeoperandTooLarge                         DeviceComplianceScriptErrorCode = "OperandTooLarge"
	DeviceComplianceScriptErrorCodeoperatorDataTypeCombinationNotSupported DeviceComplianceScriptErrorCode = "OperatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptErrorCodeoperatorMissing                         DeviceComplianceScriptErrorCode = "OperatorMissing"
	DeviceComplianceScriptErrorCodeoperatorNotSupported                    DeviceComplianceScriptErrorCode = "OperatorNotSupported"
	DeviceComplianceScriptErrorCoderemediationStringsMissing               DeviceComplianceScriptErrorCode = "RemediationStringsMissing"
	DeviceComplianceScriptErrorCoderulesMissing                            DeviceComplianceScriptErrorCode = "RulesMissing"
	DeviceComplianceScriptErrorCodesettingNameInvalid                      DeviceComplianceScriptErrorCode = "SettingNameInvalid"
	DeviceComplianceScriptErrorCodesettingNameMissing                      DeviceComplianceScriptErrorCode = "SettingNameMissing"
	DeviceComplianceScriptErrorCodesettingNameTooLarge                     DeviceComplianceScriptErrorCode = "SettingNameTooLarge"
	DeviceComplianceScriptErrorCodetitleInvalid                            DeviceComplianceScriptErrorCode = "TitleInvalid"
	DeviceComplianceScriptErrorCodetitleMissing                            DeviceComplianceScriptErrorCode = "TitleMissing"
	DeviceComplianceScriptErrorCodetitleTooLarge                           DeviceComplianceScriptErrorCode = "TitleTooLarge"
	DeviceComplianceScriptErrorCodetooManyRulesSpecified                   DeviceComplianceScriptErrorCode = "TooManyRulesSpecified"
	DeviceComplianceScriptErrorCodeunknown                                 DeviceComplianceScriptErrorCode = "Unknown"
	DeviceComplianceScriptErrorCodeunrecognizedLocale                      DeviceComplianceScriptErrorCode = "UnrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptErrorCode() []string {
	return []string{
		string(DeviceComplianceScriptErrorCodedatatypeMissing),
		string(DeviceComplianceScriptErrorCodedatatypeNotSupported),
		string(DeviceComplianceScriptErrorCodedescriptionInvalid),
		string(DeviceComplianceScriptErrorCodedescriptionMissing),
		string(DeviceComplianceScriptErrorCodedescriptionTooLarge),
		string(DeviceComplianceScriptErrorCodeduplicateLocales),
		string(DeviceComplianceScriptErrorCodeduplicateRules),
		string(DeviceComplianceScriptErrorCodeenglishLocaleMissing),
		string(DeviceComplianceScriptErrorCodejsonFileInvalid),
		string(DeviceComplianceScriptErrorCodejsonFileMissing),
		string(DeviceComplianceScriptErrorCodejsonFileTooLarge),
		string(DeviceComplianceScriptErrorCodemoreInfoUriInvalid),
		string(DeviceComplianceScriptErrorCodemoreInfoUriMissing),
		string(DeviceComplianceScriptErrorCodemoreInfoUriTooLarge),
		string(DeviceComplianceScriptErrorCodenone),
		string(DeviceComplianceScriptErrorCodeoperandInvalid),
		string(DeviceComplianceScriptErrorCodeoperandMissing),
		string(DeviceComplianceScriptErrorCodeoperandTooLarge),
		string(DeviceComplianceScriptErrorCodeoperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptErrorCodeoperatorMissing),
		string(DeviceComplianceScriptErrorCodeoperatorNotSupported),
		string(DeviceComplianceScriptErrorCoderemediationStringsMissing),
		string(DeviceComplianceScriptErrorCoderulesMissing),
		string(DeviceComplianceScriptErrorCodesettingNameInvalid),
		string(DeviceComplianceScriptErrorCodesettingNameMissing),
		string(DeviceComplianceScriptErrorCodesettingNameTooLarge),
		string(DeviceComplianceScriptErrorCodetitleInvalid),
		string(DeviceComplianceScriptErrorCodetitleMissing),
		string(DeviceComplianceScriptErrorCodetitleTooLarge),
		string(DeviceComplianceScriptErrorCodetooManyRulesSpecified),
		string(DeviceComplianceScriptErrorCodeunknown),
		string(DeviceComplianceScriptErrorCodeunrecognizedLocale),
	}
}

func parseDeviceComplianceScriptErrorCode(input string) (*DeviceComplianceScriptErrorCode, error) {
	vals := map[string]DeviceComplianceScriptErrorCode{
		"datatypemissing":                         DeviceComplianceScriptErrorCodedatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptErrorCodedatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptErrorCodedescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptErrorCodedescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptErrorCodedescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptErrorCodeduplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptErrorCodeduplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptErrorCodeenglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptErrorCodejsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptErrorCodejsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptErrorCodejsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptErrorCodemoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptErrorCodemoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptErrorCodemoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptErrorCodenone,
		"operandinvalid":                          DeviceComplianceScriptErrorCodeoperandInvalid,
		"operandmissing":                          DeviceComplianceScriptErrorCodeoperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptErrorCodeoperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptErrorCodeoperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptErrorCodeoperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptErrorCodeoperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptErrorCoderemediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptErrorCoderulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptErrorCodesettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptErrorCodesettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptErrorCodesettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptErrorCodetitleInvalid,
		"titlemissing":                            DeviceComplianceScriptErrorCodetitleMissing,
		"titletoolarge":                           DeviceComplianceScriptErrorCodetitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptErrorCodetooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptErrorCodeunknown,
		"unrecognizedlocale":                      DeviceComplianceScriptErrorCodeunrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptErrorCode(input)
	return &out, nil
}
