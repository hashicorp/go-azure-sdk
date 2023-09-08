package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError string

const (
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordatatypeMissing                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "DatatypeMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordatatypeNotSupported                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "DatatypeNotSupported"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionInvalid                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "DescriptionInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionMissing                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "DescriptionMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionTooLarge                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "DescriptionTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorduplicateLocales                        DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "DuplicateLocales"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorduplicateRules                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "DuplicateRules"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorenglishLocaleMissing                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "EnglishLocaleMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileInvalid                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "JsonFileInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileMissing                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "JsonFileMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileTooLarge                        DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "JsonFileTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriInvalid                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "MoreInfoUriInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriMissing                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "MoreInfoUriMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriTooLarge                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "MoreInfoUriTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrornone                                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "None"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandInvalid                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "OperandInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandMissing                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "OperandMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandTooLarge                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "OperandTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorDataTypeCombinationNotSupported DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "OperatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorMissing                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "OperatorMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorNotSupported                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "OperatorNotSupported"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorremediationStringsMissing               DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "RemediationStringsMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorrulesMissing                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "RulesMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameInvalid                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "SettingNameInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameMissing                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "SettingNameMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameTooLarge                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "SettingNameTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleInvalid                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "TitleInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleMissing                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "TitleMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleTooLarge                           DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "TitleTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortooManyRulesSpecified                   DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "TooManyRulesSpecified"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorunknown                                 DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "Unknown"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorunrecognizedLocale                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "UnrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError() []string {
	return []string{
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordatatypeMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordatatypeNotSupported),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorduplicateLocales),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorduplicateRules),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorenglishLocaleMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrornone),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorNotSupported),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorremediationStringsMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorrulesMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortooManyRulesSpecified),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorunknown),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorunrecognizedLocale),
	}
}

func parseDeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError(input string) (*DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError, error) {
	vals := map[string]DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError{
		"datatypemissing":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrordescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorduplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorduplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorenglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorjsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrormoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrornone,
		"operandinvalid":                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandInvalid,
		"operandmissing":                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErroroperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorremediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorrulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorsettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleInvalid,
		"titlemissing":                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleMissing,
		"titletoolarge":                           DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrortooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorunknown,
		"unrecognizedlocale":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationErrorunrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError(input)
	return &out, nil
}
