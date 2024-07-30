package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError string

const (
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DatatypeMissing                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "datatypeMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DatatypeNotSupported                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "datatypeNotSupported"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionInvalid                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "descriptionInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionMissing                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "descriptionMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionTooLarge                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "descriptionTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DuplicateLocales                        DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "duplicateLocales"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DuplicateRules                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "duplicateRules"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_EnglishLocaleMissing                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "englishLocaleMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileInvalid                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "jsonFileInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileMissing                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "jsonFileMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileTooLarge                        DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "jsonFileTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "moreInfoUriInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriMissing                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "moreInfoUriMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "moreInfoUriTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_None                                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "none"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandInvalid                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "operandInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandMissing                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "operandMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandTooLarge                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "operandTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "operatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorMissing                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "operatorMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorNotSupported                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "operatorNotSupported"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_RemediationStringsMissing               DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "remediationStringsMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_RulesMissing                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "rulesMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameInvalid                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "settingNameInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameMissing                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "settingNameMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameTooLarge                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "settingNameTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleInvalid                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "titleInvalid"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleMissing                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "titleMissing"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleTooLarge                           DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "titleTooLarge"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TooManyRulesSpecified                   DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "tooManyRulesSpecified"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_Unknown                                 DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "unknown"
	DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_UnrecognizedLocale                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError = "unrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError() []string {
	return []string{
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DatatypeMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DatatypeNotSupported),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DuplicateLocales),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DuplicateRules),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_EnglishLocaleMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_None),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorNotSupported),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_RemediationStringsMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_RulesMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleInvalid),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleMissing),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleTooLarge),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TooManyRulesSpecified),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_Unknown),
		string(DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_UnrecognizedLocale),
	}
}

func (s *DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError(input string) (*DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError, error) {
	vals := map[string]DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError{
		"datatypemissing":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DuplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_DuplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_EnglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_JsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_MoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_None,
		"operandinvalid":                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandInvalid,
		"operandmissing":                          DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_OperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_RemediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_RulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_SettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleInvalid,
		"titlemissing":                            DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleMissing,
		"titletoolarge":                           DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_TooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_Unknown,
		"unrecognizedlocale":                      DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError_UnrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptErrorDeviceComplianceScriptRulesValidationError(input)
	return &out, nil
}
