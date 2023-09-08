package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleErrorCode string

const (
	DeviceComplianceScriptRuleErrorCodedatatypeMissing                         DeviceComplianceScriptRuleErrorCode = "DatatypeMissing"
	DeviceComplianceScriptRuleErrorCodedatatypeNotSupported                    DeviceComplianceScriptRuleErrorCode = "DatatypeNotSupported"
	DeviceComplianceScriptRuleErrorCodedescriptionInvalid                      DeviceComplianceScriptRuleErrorCode = "DescriptionInvalid"
	DeviceComplianceScriptRuleErrorCodedescriptionMissing                      DeviceComplianceScriptRuleErrorCode = "DescriptionMissing"
	DeviceComplianceScriptRuleErrorCodedescriptionTooLarge                     DeviceComplianceScriptRuleErrorCode = "DescriptionTooLarge"
	DeviceComplianceScriptRuleErrorCodeduplicateLocales                        DeviceComplianceScriptRuleErrorCode = "DuplicateLocales"
	DeviceComplianceScriptRuleErrorCodeduplicateRules                          DeviceComplianceScriptRuleErrorCode = "DuplicateRules"
	DeviceComplianceScriptRuleErrorCodeenglishLocaleMissing                    DeviceComplianceScriptRuleErrorCode = "EnglishLocaleMissing"
	DeviceComplianceScriptRuleErrorCodejsonFileInvalid                         DeviceComplianceScriptRuleErrorCode = "JsonFileInvalid"
	DeviceComplianceScriptRuleErrorCodejsonFileMissing                         DeviceComplianceScriptRuleErrorCode = "JsonFileMissing"
	DeviceComplianceScriptRuleErrorCodejsonFileTooLarge                        DeviceComplianceScriptRuleErrorCode = "JsonFileTooLarge"
	DeviceComplianceScriptRuleErrorCodemoreInfoUriInvalid                      DeviceComplianceScriptRuleErrorCode = "MoreInfoUriInvalid"
	DeviceComplianceScriptRuleErrorCodemoreInfoUriMissing                      DeviceComplianceScriptRuleErrorCode = "MoreInfoUriMissing"
	DeviceComplianceScriptRuleErrorCodemoreInfoUriTooLarge                     DeviceComplianceScriptRuleErrorCode = "MoreInfoUriTooLarge"
	DeviceComplianceScriptRuleErrorCodenone                                    DeviceComplianceScriptRuleErrorCode = "None"
	DeviceComplianceScriptRuleErrorCodeoperandInvalid                          DeviceComplianceScriptRuleErrorCode = "OperandInvalid"
	DeviceComplianceScriptRuleErrorCodeoperandMissing                          DeviceComplianceScriptRuleErrorCode = "OperandMissing"
	DeviceComplianceScriptRuleErrorCodeoperandTooLarge                         DeviceComplianceScriptRuleErrorCode = "OperandTooLarge"
	DeviceComplianceScriptRuleErrorCodeoperatorDataTypeCombinationNotSupported DeviceComplianceScriptRuleErrorCode = "OperatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptRuleErrorCodeoperatorMissing                         DeviceComplianceScriptRuleErrorCode = "OperatorMissing"
	DeviceComplianceScriptRuleErrorCodeoperatorNotSupported                    DeviceComplianceScriptRuleErrorCode = "OperatorNotSupported"
	DeviceComplianceScriptRuleErrorCoderemediationStringsMissing               DeviceComplianceScriptRuleErrorCode = "RemediationStringsMissing"
	DeviceComplianceScriptRuleErrorCoderulesMissing                            DeviceComplianceScriptRuleErrorCode = "RulesMissing"
	DeviceComplianceScriptRuleErrorCodesettingNameInvalid                      DeviceComplianceScriptRuleErrorCode = "SettingNameInvalid"
	DeviceComplianceScriptRuleErrorCodesettingNameMissing                      DeviceComplianceScriptRuleErrorCode = "SettingNameMissing"
	DeviceComplianceScriptRuleErrorCodesettingNameTooLarge                     DeviceComplianceScriptRuleErrorCode = "SettingNameTooLarge"
	DeviceComplianceScriptRuleErrorCodetitleInvalid                            DeviceComplianceScriptRuleErrorCode = "TitleInvalid"
	DeviceComplianceScriptRuleErrorCodetitleMissing                            DeviceComplianceScriptRuleErrorCode = "TitleMissing"
	DeviceComplianceScriptRuleErrorCodetitleTooLarge                           DeviceComplianceScriptRuleErrorCode = "TitleTooLarge"
	DeviceComplianceScriptRuleErrorCodetooManyRulesSpecified                   DeviceComplianceScriptRuleErrorCode = "TooManyRulesSpecified"
	DeviceComplianceScriptRuleErrorCodeunknown                                 DeviceComplianceScriptRuleErrorCode = "Unknown"
	DeviceComplianceScriptRuleErrorCodeunrecognizedLocale                      DeviceComplianceScriptRuleErrorCode = "UnrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptRuleErrorCode() []string {
	return []string{
		string(DeviceComplianceScriptRuleErrorCodedatatypeMissing),
		string(DeviceComplianceScriptRuleErrorCodedatatypeNotSupported),
		string(DeviceComplianceScriptRuleErrorCodedescriptionInvalid),
		string(DeviceComplianceScriptRuleErrorCodedescriptionMissing),
		string(DeviceComplianceScriptRuleErrorCodedescriptionTooLarge),
		string(DeviceComplianceScriptRuleErrorCodeduplicateLocales),
		string(DeviceComplianceScriptRuleErrorCodeduplicateRules),
		string(DeviceComplianceScriptRuleErrorCodeenglishLocaleMissing),
		string(DeviceComplianceScriptRuleErrorCodejsonFileInvalid),
		string(DeviceComplianceScriptRuleErrorCodejsonFileMissing),
		string(DeviceComplianceScriptRuleErrorCodejsonFileTooLarge),
		string(DeviceComplianceScriptRuleErrorCodemoreInfoUriInvalid),
		string(DeviceComplianceScriptRuleErrorCodemoreInfoUriMissing),
		string(DeviceComplianceScriptRuleErrorCodemoreInfoUriTooLarge),
		string(DeviceComplianceScriptRuleErrorCodenone),
		string(DeviceComplianceScriptRuleErrorCodeoperandInvalid),
		string(DeviceComplianceScriptRuleErrorCodeoperandMissing),
		string(DeviceComplianceScriptRuleErrorCodeoperandTooLarge),
		string(DeviceComplianceScriptRuleErrorCodeoperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptRuleErrorCodeoperatorMissing),
		string(DeviceComplianceScriptRuleErrorCodeoperatorNotSupported),
		string(DeviceComplianceScriptRuleErrorCoderemediationStringsMissing),
		string(DeviceComplianceScriptRuleErrorCoderulesMissing),
		string(DeviceComplianceScriptRuleErrorCodesettingNameInvalid),
		string(DeviceComplianceScriptRuleErrorCodesettingNameMissing),
		string(DeviceComplianceScriptRuleErrorCodesettingNameTooLarge),
		string(DeviceComplianceScriptRuleErrorCodetitleInvalid),
		string(DeviceComplianceScriptRuleErrorCodetitleMissing),
		string(DeviceComplianceScriptRuleErrorCodetitleTooLarge),
		string(DeviceComplianceScriptRuleErrorCodetooManyRulesSpecified),
		string(DeviceComplianceScriptRuleErrorCodeunknown),
		string(DeviceComplianceScriptRuleErrorCodeunrecognizedLocale),
	}
}

func parseDeviceComplianceScriptRuleErrorCode(input string) (*DeviceComplianceScriptRuleErrorCode, error) {
	vals := map[string]DeviceComplianceScriptRuleErrorCode{
		"datatypemissing":                         DeviceComplianceScriptRuleErrorCodedatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptRuleErrorCodedatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptRuleErrorCodedescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptRuleErrorCodedescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptRuleErrorCodedescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptRuleErrorCodeduplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptRuleErrorCodeduplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptRuleErrorCodeenglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptRuleErrorCodejsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptRuleErrorCodejsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptRuleErrorCodejsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptRuleErrorCodemoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptRuleErrorCodemoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptRuleErrorCodemoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptRuleErrorCodenone,
		"operandinvalid":                          DeviceComplianceScriptRuleErrorCodeoperandInvalid,
		"operandmissing":                          DeviceComplianceScriptRuleErrorCodeoperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptRuleErrorCodeoperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptRuleErrorCodeoperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptRuleErrorCodeoperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptRuleErrorCodeoperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptRuleErrorCoderemediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptRuleErrorCoderulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptRuleErrorCodesettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptRuleErrorCodesettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptRuleErrorCodesettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptRuleErrorCodetitleInvalid,
		"titlemissing":                            DeviceComplianceScriptRuleErrorCodetitleMissing,
		"titletoolarge":                           DeviceComplianceScriptRuleErrorCodetitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptRuleErrorCodetooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptRuleErrorCodeunknown,
		"unrecognizedlocale":                      DeviceComplianceScriptRuleErrorCodeunrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleErrorCode(input)
	return &out, nil
}
