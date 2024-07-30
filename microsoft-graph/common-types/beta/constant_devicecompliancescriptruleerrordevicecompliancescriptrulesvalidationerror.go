package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError string

const (
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DatatypeMissing                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "datatypeMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DatatypeNotSupported                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "datatypeNotSupported"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionInvalid                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "descriptionInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionMissing                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "descriptionMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionTooLarge                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "descriptionTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DuplicateLocales                        DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "duplicateLocales"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DuplicateRules                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "duplicateRules"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_EnglishLocaleMissing                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "englishLocaleMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileInvalid                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "jsonFileInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileMissing                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "jsonFileMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileTooLarge                        DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "jsonFileTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "moreInfoUriInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriMissing                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "moreInfoUriMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "moreInfoUriTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_None                                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "none"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandInvalid                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "operandInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandMissing                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "operandMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandTooLarge                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "operandTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "operatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorMissing                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "operatorMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorNotSupported                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "operatorNotSupported"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_RemediationStringsMissing               DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "remediationStringsMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_RulesMissing                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "rulesMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameInvalid                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "settingNameInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameMissing                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "settingNameMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameTooLarge                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "settingNameTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleInvalid                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "titleInvalid"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleMissing                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "titleMissing"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleTooLarge                           DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "titleTooLarge"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TooManyRulesSpecified                   DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "tooManyRulesSpecified"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_Unknown                                 DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "unknown"
	DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_UnrecognizedLocale                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError = "unrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError() []string {
	return []string{
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DatatypeMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DatatypeNotSupported),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DuplicateLocales),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DuplicateRules),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_EnglishLocaleMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_None),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorNotSupported),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_RemediationStringsMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_RulesMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleInvalid),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleMissing),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleTooLarge),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TooManyRulesSpecified),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_Unknown),
		string(DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_UnrecognizedLocale),
	}
}

func (s *DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError(input string) (*DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError, error) {
	vals := map[string]DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError{
		"datatypemissing":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DuplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_DuplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_EnglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_JsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_None,
		"operandinvalid":                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandInvalid,
		"operandmissing":                          DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_OperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_RemediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_RulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_SettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleInvalid,
		"titlemissing":                            DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleMissing,
		"titletoolarge":                           DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_TooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_Unknown,
		"unrecognizedlocale":                      DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError_UnrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleErrorDeviceComplianceScriptRulesValidationError(input)
	return &out, nil
}
