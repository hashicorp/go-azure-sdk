package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRuleErrorCode string

const (
	DeviceComplianceScriptRuleErrorCode_DatatypeMissing                         DeviceComplianceScriptRuleErrorCode = "datatypeMissing"
	DeviceComplianceScriptRuleErrorCode_DatatypeNotSupported                    DeviceComplianceScriptRuleErrorCode = "datatypeNotSupported"
	DeviceComplianceScriptRuleErrorCode_DescriptionInvalid                      DeviceComplianceScriptRuleErrorCode = "descriptionInvalid"
	DeviceComplianceScriptRuleErrorCode_DescriptionMissing                      DeviceComplianceScriptRuleErrorCode = "descriptionMissing"
	DeviceComplianceScriptRuleErrorCode_DescriptionTooLarge                     DeviceComplianceScriptRuleErrorCode = "descriptionTooLarge"
	DeviceComplianceScriptRuleErrorCode_DuplicateLocales                        DeviceComplianceScriptRuleErrorCode = "duplicateLocales"
	DeviceComplianceScriptRuleErrorCode_DuplicateRules                          DeviceComplianceScriptRuleErrorCode = "duplicateRules"
	DeviceComplianceScriptRuleErrorCode_EnglishLocaleMissing                    DeviceComplianceScriptRuleErrorCode = "englishLocaleMissing"
	DeviceComplianceScriptRuleErrorCode_JsonFileInvalid                         DeviceComplianceScriptRuleErrorCode = "jsonFileInvalid"
	DeviceComplianceScriptRuleErrorCode_JsonFileMissing                         DeviceComplianceScriptRuleErrorCode = "jsonFileMissing"
	DeviceComplianceScriptRuleErrorCode_JsonFileTooLarge                        DeviceComplianceScriptRuleErrorCode = "jsonFileTooLarge"
	DeviceComplianceScriptRuleErrorCode_MoreInfoUriInvalid                      DeviceComplianceScriptRuleErrorCode = "moreInfoUriInvalid"
	DeviceComplianceScriptRuleErrorCode_MoreInfoUriMissing                      DeviceComplianceScriptRuleErrorCode = "moreInfoUriMissing"
	DeviceComplianceScriptRuleErrorCode_MoreInfoUriTooLarge                     DeviceComplianceScriptRuleErrorCode = "moreInfoUriTooLarge"
	DeviceComplianceScriptRuleErrorCode_None                                    DeviceComplianceScriptRuleErrorCode = "none"
	DeviceComplianceScriptRuleErrorCode_OperandInvalid                          DeviceComplianceScriptRuleErrorCode = "operandInvalid"
	DeviceComplianceScriptRuleErrorCode_OperandMissing                          DeviceComplianceScriptRuleErrorCode = "operandMissing"
	DeviceComplianceScriptRuleErrorCode_OperandTooLarge                         DeviceComplianceScriptRuleErrorCode = "operandTooLarge"
	DeviceComplianceScriptRuleErrorCode_OperatorDataTypeCombinationNotSupported DeviceComplianceScriptRuleErrorCode = "operatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptRuleErrorCode_OperatorMissing                         DeviceComplianceScriptRuleErrorCode = "operatorMissing"
	DeviceComplianceScriptRuleErrorCode_OperatorNotSupported                    DeviceComplianceScriptRuleErrorCode = "operatorNotSupported"
	DeviceComplianceScriptRuleErrorCode_RemediationStringsMissing               DeviceComplianceScriptRuleErrorCode = "remediationStringsMissing"
	DeviceComplianceScriptRuleErrorCode_RulesMissing                            DeviceComplianceScriptRuleErrorCode = "rulesMissing"
	DeviceComplianceScriptRuleErrorCode_SettingNameInvalid                      DeviceComplianceScriptRuleErrorCode = "settingNameInvalid"
	DeviceComplianceScriptRuleErrorCode_SettingNameMissing                      DeviceComplianceScriptRuleErrorCode = "settingNameMissing"
	DeviceComplianceScriptRuleErrorCode_SettingNameTooLarge                     DeviceComplianceScriptRuleErrorCode = "settingNameTooLarge"
	DeviceComplianceScriptRuleErrorCode_TitleInvalid                            DeviceComplianceScriptRuleErrorCode = "titleInvalid"
	DeviceComplianceScriptRuleErrorCode_TitleMissing                            DeviceComplianceScriptRuleErrorCode = "titleMissing"
	DeviceComplianceScriptRuleErrorCode_TitleTooLarge                           DeviceComplianceScriptRuleErrorCode = "titleTooLarge"
	DeviceComplianceScriptRuleErrorCode_TooManyRulesSpecified                   DeviceComplianceScriptRuleErrorCode = "tooManyRulesSpecified"
	DeviceComplianceScriptRuleErrorCode_Unknown                                 DeviceComplianceScriptRuleErrorCode = "unknown"
	DeviceComplianceScriptRuleErrorCode_UnrecognizedLocale                      DeviceComplianceScriptRuleErrorCode = "unrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptRuleErrorCode() []string {
	return []string{
		string(DeviceComplianceScriptRuleErrorCode_DatatypeMissing),
		string(DeviceComplianceScriptRuleErrorCode_DatatypeNotSupported),
		string(DeviceComplianceScriptRuleErrorCode_DescriptionInvalid),
		string(DeviceComplianceScriptRuleErrorCode_DescriptionMissing),
		string(DeviceComplianceScriptRuleErrorCode_DescriptionTooLarge),
		string(DeviceComplianceScriptRuleErrorCode_DuplicateLocales),
		string(DeviceComplianceScriptRuleErrorCode_DuplicateRules),
		string(DeviceComplianceScriptRuleErrorCode_EnglishLocaleMissing),
		string(DeviceComplianceScriptRuleErrorCode_JsonFileInvalid),
		string(DeviceComplianceScriptRuleErrorCode_JsonFileMissing),
		string(DeviceComplianceScriptRuleErrorCode_JsonFileTooLarge),
		string(DeviceComplianceScriptRuleErrorCode_MoreInfoUriInvalid),
		string(DeviceComplianceScriptRuleErrorCode_MoreInfoUriMissing),
		string(DeviceComplianceScriptRuleErrorCode_MoreInfoUriTooLarge),
		string(DeviceComplianceScriptRuleErrorCode_None),
		string(DeviceComplianceScriptRuleErrorCode_OperandInvalid),
		string(DeviceComplianceScriptRuleErrorCode_OperandMissing),
		string(DeviceComplianceScriptRuleErrorCode_OperandTooLarge),
		string(DeviceComplianceScriptRuleErrorCode_OperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptRuleErrorCode_OperatorMissing),
		string(DeviceComplianceScriptRuleErrorCode_OperatorNotSupported),
		string(DeviceComplianceScriptRuleErrorCode_RemediationStringsMissing),
		string(DeviceComplianceScriptRuleErrorCode_RulesMissing),
		string(DeviceComplianceScriptRuleErrorCode_SettingNameInvalid),
		string(DeviceComplianceScriptRuleErrorCode_SettingNameMissing),
		string(DeviceComplianceScriptRuleErrorCode_SettingNameTooLarge),
		string(DeviceComplianceScriptRuleErrorCode_TitleInvalid),
		string(DeviceComplianceScriptRuleErrorCode_TitleMissing),
		string(DeviceComplianceScriptRuleErrorCode_TitleTooLarge),
		string(DeviceComplianceScriptRuleErrorCode_TooManyRulesSpecified),
		string(DeviceComplianceScriptRuleErrorCode_Unknown),
		string(DeviceComplianceScriptRuleErrorCode_UnrecognizedLocale),
	}
}

func (s *DeviceComplianceScriptRuleErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRuleErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRuleErrorCode(input string) (*DeviceComplianceScriptRuleErrorCode, error) {
	vals := map[string]DeviceComplianceScriptRuleErrorCode{
		"datatypemissing":                         DeviceComplianceScriptRuleErrorCode_DatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptRuleErrorCode_DatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptRuleErrorCode_DescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptRuleErrorCode_DescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptRuleErrorCode_DescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptRuleErrorCode_DuplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptRuleErrorCode_DuplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptRuleErrorCode_EnglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptRuleErrorCode_JsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptRuleErrorCode_JsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptRuleErrorCode_JsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptRuleErrorCode_MoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptRuleErrorCode_MoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptRuleErrorCode_MoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptRuleErrorCode_None,
		"operandinvalid":                          DeviceComplianceScriptRuleErrorCode_OperandInvalid,
		"operandmissing":                          DeviceComplianceScriptRuleErrorCode_OperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptRuleErrorCode_OperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptRuleErrorCode_OperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptRuleErrorCode_OperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptRuleErrorCode_OperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptRuleErrorCode_RemediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptRuleErrorCode_RulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptRuleErrorCode_SettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptRuleErrorCode_SettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptRuleErrorCode_SettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptRuleErrorCode_TitleInvalid,
		"titlemissing":                            DeviceComplianceScriptRuleErrorCode_TitleMissing,
		"titletoolarge":                           DeviceComplianceScriptRuleErrorCode_TitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptRuleErrorCode_TooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptRuleErrorCode_Unknown,
		"unrecognizedlocale":                      DeviceComplianceScriptRuleErrorCode_UnrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRuleErrorCode(input)
	return &out, nil
}
