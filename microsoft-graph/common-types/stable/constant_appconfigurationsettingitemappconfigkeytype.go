package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConfigurationSettingItemAppConfigKeyType string

const (
	AppConfigurationSettingItemAppConfigKeyType_BooleanType AppConfigurationSettingItemAppConfigKeyType = "booleanType"
	AppConfigurationSettingItemAppConfigKeyType_IntegerType AppConfigurationSettingItemAppConfigKeyType = "integerType"
	AppConfigurationSettingItemAppConfigKeyType_RealType    AppConfigurationSettingItemAppConfigKeyType = "realType"
	AppConfigurationSettingItemAppConfigKeyType_StringType  AppConfigurationSettingItemAppConfigKeyType = "stringType"
	AppConfigurationSettingItemAppConfigKeyType_TokenType   AppConfigurationSettingItemAppConfigKeyType = "tokenType"
)

func PossibleValuesForAppConfigurationSettingItemAppConfigKeyType() []string {
	return []string{
		string(AppConfigurationSettingItemAppConfigKeyType_BooleanType),
		string(AppConfigurationSettingItemAppConfigKeyType_IntegerType),
		string(AppConfigurationSettingItemAppConfigKeyType_RealType),
		string(AppConfigurationSettingItemAppConfigKeyType_StringType),
		string(AppConfigurationSettingItemAppConfigKeyType_TokenType),
	}
}

func (s *AppConfigurationSettingItemAppConfigKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppConfigurationSettingItemAppConfigKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppConfigurationSettingItemAppConfigKeyType(input string) (*AppConfigurationSettingItemAppConfigKeyType, error) {
	vals := map[string]AppConfigurationSettingItemAppConfigKeyType{
		"booleantype": AppConfigurationSettingItemAppConfigKeyType_BooleanType,
		"integertype": AppConfigurationSettingItemAppConfigKeyType_IntegerType,
		"realtype":    AppConfigurationSettingItemAppConfigKeyType_RealType,
		"stringtype":  AppConfigurationSettingItemAppConfigKeyType_StringType,
		"tokentype":   AppConfigurationSettingItemAppConfigKeyType_TokenType,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppConfigurationSettingItemAppConfigKeyType(input)
	return &out, nil
}
