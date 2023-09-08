package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError string

const (
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordatatypeMissing                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "DatatypeMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordatatypeNotSupported                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "DatatypeNotSupported"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionInvalid                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "DescriptionInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionMissing                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "DescriptionMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionTooLarge                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "DescriptionTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorduplicateLocales                        DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "DuplicateLocales"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorduplicateRules                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "DuplicateRules"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorenglishLocaleMissing                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "EnglishLocaleMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileInvalid                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "JsonFileInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileMissing                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "JsonFileMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileTooLarge                        DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "JsonFileTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriInvalid                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "MoreInfoUriInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriMissing                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "MoreInfoUriMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriTooLarge                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "MoreInfoUriTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrornone                                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "None"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandInvalid                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "OperandInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandMissing                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "OperandMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandTooLarge                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "OperandTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorDataTypeCombinationNotSupported DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "OperatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorMissing                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "OperatorMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorNotSupported                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "OperatorNotSupported"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorremediationStringsMissing               DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "RemediationStringsMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorrulesMissing                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "RulesMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameInvalid                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "SettingNameInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameMissing                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "SettingNameMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameTooLarge                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "SettingNameTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleInvalid                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "TitleInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleMissing                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "TitleMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleTooLarge                           DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "TitleTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortooManyRulesSpecified                   DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "TooManyRulesSpecified"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorunknown                                 DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "Unknown"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorunrecognizedLocale                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "UnrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError() []string {
	return []string{
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordatatypeMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordatatypeNotSupported),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorduplicateLocales),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorduplicateRules),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorenglishLocaleMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrornone),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorNotSupported),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorremediationStringsMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorrulesMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortooManyRulesSpecified),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorunknown),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorunrecognizedLocale),
	}
}

func parseDeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError(input string) (*DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError, error) {
	vals := map[string]DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError{
		"datatypemissing":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrordescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorduplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorduplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorenglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorjsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrornone,
		"operandinvalid":                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandInvalid,
		"operandmissing":                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErroroperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorremediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorrulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorsettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleInvalid,
		"titlemissing":                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleMissing,
		"titletoolarge":                           DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrortooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorunknown,
		"unrecognizedlocale":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationErrorunrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError(input)
	return &out, nil
}
