package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptErrorCode string

const (
	DeviceComplianceScriptErrorCode_DatatypeMissing                         DeviceComplianceScriptErrorCode = "datatypeMissing"
	DeviceComplianceScriptErrorCode_DatatypeNotSupported                    DeviceComplianceScriptErrorCode = "datatypeNotSupported"
	DeviceComplianceScriptErrorCode_DescriptionInvalid                      DeviceComplianceScriptErrorCode = "descriptionInvalid"
	DeviceComplianceScriptErrorCode_DescriptionMissing                      DeviceComplianceScriptErrorCode = "descriptionMissing"
	DeviceComplianceScriptErrorCode_DescriptionTooLarge                     DeviceComplianceScriptErrorCode = "descriptionTooLarge"
	DeviceComplianceScriptErrorCode_DuplicateLocales                        DeviceComplianceScriptErrorCode = "duplicateLocales"
	DeviceComplianceScriptErrorCode_DuplicateRules                          DeviceComplianceScriptErrorCode = "duplicateRules"
	DeviceComplianceScriptErrorCode_EnglishLocaleMissing                    DeviceComplianceScriptErrorCode = "englishLocaleMissing"
	DeviceComplianceScriptErrorCode_JsonFileInvalid                         DeviceComplianceScriptErrorCode = "jsonFileInvalid"
	DeviceComplianceScriptErrorCode_JsonFileMissing                         DeviceComplianceScriptErrorCode = "jsonFileMissing"
	DeviceComplianceScriptErrorCode_JsonFileTooLarge                        DeviceComplianceScriptErrorCode = "jsonFileTooLarge"
	DeviceComplianceScriptErrorCode_MoreInfoUriInvalid                      DeviceComplianceScriptErrorCode = "moreInfoUriInvalid"
	DeviceComplianceScriptErrorCode_MoreInfoUriMissing                      DeviceComplianceScriptErrorCode = "moreInfoUriMissing"
	DeviceComplianceScriptErrorCode_MoreInfoUriTooLarge                     DeviceComplianceScriptErrorCode = "moreInfoUriTooLarge"
	DeviceComplianceScriptErrorCode_None                                    DeviceComplianceScriptErrorCode = "none"
	DeviceComplianceScriptErrorCode_OperandInvalid                          DeviceComplianceScriptErrorCode = "operandInvalid"
	DeviceComplianceScriptErrorCode_OperandMissing                          DeviceComplianceScriptErrorCode = "operandMissing"
	DeviceComplianceScriptErrorCode_OperandTooLarge                         DeviceComplianceScriptErrorCode = "operandTooLarge"
	DeviceComplianceScriptErrorCode_OperatorDataTypeCombinationNotSupported DeviceComplianceScriptErrorCode = "operatorDataTypeCombinationNotSupported"
	DeviceComplianceScriptErrorCode_OperatorMissing                         DeviceComplianceScriptErrorCode = "operatorMissing"
	DeviceComplianceScriptErrorCode_OperatorNotSupported                    DeviceComplianceScriptErrorCode = "operatorNotSupported"
	DeviceComplianceScriptErrorCode_RemediationStringsMissing               DeviceComplianceScriptErrorCode = "remediationStringsMissing"
	DeviceComplianceScriptErrorCode_RulesMissing                            DeviceComplianceScriptErrorCode = "rulesMissing"
	DeviceComplianceScriptErrorCode_SettingNameInvalid                      DeviceComplianceScriptErrorCode = "settingNameInvalid"
	DeviceComplianceScriptErrorCode_SettingNameMissing                      DeviceComplianceScriptErrorCode = "settingNameMissing"
	DeviceComplianceScriptErrorCode_SettingNameTooLarge                     DeviceComplianceScriptErrorCode = "settingNameTooLarge"
	DeviceComplianceScriptErrorCode_TitleInvalid                            DeviceComplianceScriptErrorCode = "titleInvalid"
	DeviceComplianceScriptErrorCode_TitleMissing                            DeviceComplianceScriptErrorCode = "titleMissing"
	DeviceComplianceScriptErrorCode_TitleTooLarge                           DeviceComplianceScriptErrorCode = "titleTooLarge"
	DeviceComplianceScriptErrorCode_TooManyRulesSpecified                   DeviceComplianceScriptErrorCode = "tooManyRulesSpecified"
	DeviceComplianceScriptErrorCode_Unknown                                 DeviceComplianceScriptErrorCode = "unknown"
	DeviceComplianceScriptErrorCode_UnrecognizedLocale                      DeviceComplianceScriptErrorCode = "unrecognizedLocale"
)

func PossibleValuesForDeviceComplianceScriptErrorCode() []string {
	return []string{
		string(DeviceComplianceScriptErrorCode_DatatypeMissing),
		string(DeviceComplianceScriptErrorCode_DatatypeNotSupported),
		string(DeviceComplianceScriptErrorCode_DescriptionInvalid),
		string(DeviceComplianceScriptErrorCode_DescriptionMissing),
		string(DeviceComplianceScriptErrorCode_DescriptionTooLarge),
		string(DeviceComplianceScriptErrorCode_DuplicateLocales),
		string(DeviceComplianceScriptErrorCode_DuplicateRules),
		string(DeviceComplianceScriptErrorCode_EnglishLocaleMissing),
		string(DeviceComplianceScriptErrorCode_JsonFileInvalid),
		string(DeviceComplianceScriptErrorCode_JsonFileMissing),
		string(DeviceComplianceScriptErrorCode_JsonFileTooLarge),
		string(DeviceComplianceScriptErrorCode_MoreInfoUriInvalid),
		string(DeviceComplianceScriptErrorCode_MoreInfoUriMissing),
		string(DeviceComplianceScriptErrorCode_MoreInfoUriTooLarge),
		string(DeviceComplianceScriptErrorCode_None),
		string(DeviceComplianceScriptErrorCode_OperandInvalid),
		string(DeviceComplianceScriptErrorCode_OperandMissing),
		string(DeviceComplianceScriptErrorCode_OperandTooLarge),
		string(DeviceComplianceScriptErrorCode_OperatorDataTypeCombinationNotSupported),
		string(DeviceComplianceScriptErrorCode_OperatorMissing),
		string(DeviceComplianceScriptErrorCode_OperatorNotSupported),
		string(DeviceComplianceScriptErrorCode_RemediationStringsMissing),
		string(DeviceComplianceScriptErrorCode_RulesMissing),
		string(DeviceComplianceScriptErrorCode_SettingNameInvalid),
		string(DeviceComplianceScriptErrorCode_SettingNameMissing),
		string(DeviceComplianceScriptErrorCode_SettingNameTooLarge),
		string(DeviceComplianceScriptErrorCode_TitleInvalid),
		string(DeviceComplianceScriptErrorCode_TitleMissing),
		string(DeviceComplianceScriptErrorCode_TitleTooLarge),
		string(DeviceComplianceScriptErrorCode_TooManyRulesSpecified),
		string(DeviceComplianceScriptErrorCode_Unknown),
		string(DeviceComplianceScriptErrorCode_UnrecognizedLocale),
	}
}

func (s *DeviceComplianceScriptErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptErrorCode(input string) (*DeviceComplianceScriptErrorCode, error) {
	vals := map[string]DeviceComplianceScriptErrorCode{
		"datatypemissing":                         DeviceComplianceScriptErrorCode_DatatypeMissing,
		"datatypenotsupported":                    DeviceComplianceScriptErrorCode_DatatypeNotSupported,
		"descriptioninvalid":                      DeviceComplianceScriptErrorCode_DescriptionInvalid,
		"descriptionmissing":                      DeviceComplianceScriptErrorCode_DescriptionMissing,
		"descriptiontoolarge":                     DeviceComplianceScriptErrorCode_DescriptionTooLarge,
		"duplicatelocales":                        DeviceComplianceScriptErrorCode_DuplicateLocales,
		"duplicaterules":                          DeviceComplianceScriptErrorCode_DuplicateRules,
		"englishlocalemissing":                    DeviceComplianceScriptErrorCode_EnglishLocaleMissing,
		"jsonfileinvalid":                         DeviceComplianceScriptErrorCode_JsonFileInvalid,
		"jsonfilemissing":                         DeviceComplianceScriptErrorCode_JsonFileMissing,
		"jsonfiletoolarge":                        DeviceComplianceScriptErrorCode_JsonFileTooLarge,
		"moreinfouriinvalid":                      DeviceComplianceScriptErrorCode_MoreInfoUriInvalid,
		"moreinfourimissing":                      DeviceComplianceScriptErrorCode_MoreInfoUriMissing,
		"moreinfouritoolarge":                     DeviceComplianceScriptErrorCode_MoreInfoUriTooLarge,
		"none":                                    DeviceComplianceScriptErrorCode_None,
		"operandinvalid":                          DeviceComplianceScriptErrorCode_OperandInvalid,
		"operandmissing":                          DeviceComplianceScriptErrorCode_OperandMissing,
		"operandtoolarge":                         DeviceComplianceScriptErrorCode_OperandTooLarge,
		"operatordatatypecombinationnotsupported": DeviceComplianceScriptErrorCode_OperatorDataTypeCombinationNotSupported,
		"operatormissing":                         DeviceComplianceScriptErrorCode_OperatorMissing,
		"operatornotsupported":                    DeviceComplianceScriptErrorCode_OperatorNotSupported,
		"remediationstringsmissing":               DeviceComplianceScriptErrorCode_RemediationStringsMissing,
		"rulesmissing":                            DeviceComplianceScriptErrorCode_RulesMissing,
		"settingnameinvalid":                      DeviceComplianceScriptErrorCode_SettingNameInvalid,
		"settingnamemissing":                      DeviceComplianceScriptErrorCode_SettingNameMissing,
		"settingnametoolarge":                     DeviceComplianceScriptErrorCode_SettingNameTooLarge,
		"titleinvalid":                            DeviceComplianceScriptErrorCode_TitleInvalid,
		"titlemissing":                            DeviceComplianceScriptErrorCode_TitleMissing,
		"titletoolarge":                           DeviceComplianceScriptErrorCode_TitleTooLarge,
		"toomanyrulesspecified":                   DeviceComplianceScriptErrorCode_TooManyRulesSpecified,
		"unknown":                                 DeviceComplianceScriptErrorCode_Unknown,
		"unrecognizedlocale":                      DeviceComplianceScriptErrorCode_UnrecognizedLocale,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptErrorCode(input)
	return &out, nil
}
