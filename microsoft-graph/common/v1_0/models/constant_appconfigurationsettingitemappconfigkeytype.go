package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConfigurationSettingItemAppConfigKeyType string

const (
	AppConfigurationSettingItemAppConfigKeyTypebooleanType AppConfigurationSettingItemAppConfigKeyType = "BooleanType"
	AppConfigurationSettingItemAppConfigKeyTypeintegerType AppConfigurationSettingItemAppConfigKeyType = "IntegerType"
	AppConfigurationSettingItemAppConfigKeyTyperealType    AppConfigurationSettingItemAppConfigKeyType = "RealType"
	AppConfigurationSettingItemAppConfigKeyTypestringType  AppConfigurationSettingItemAppConfigKeyType = "StringType"
	AppConfigurationSettingItemAppConfigKeyTypetokenType   AppConfigurationSettingItemAppConfigKeyType = "TokenType"
)

func PossibleValuesForAppConfigurationSettingItemAppConfigKeyType() []string {
	return []string{
		string(AppConfigurationSettingItemAppConfigKeyTypebooleanType),
		string(AppConfigurationSettingItemAppConfigKeyTypeintegerType),
		string(AppConfigurationSettingItemAppConfigKeyTyperealType),
		string(AppConfigurationSettingItemAppConfigKeyTypestringType),
		string(AppConfigurationSettingItemAppConfigKeyTypetokenType),
	}
}

func parseAppConfigurationSettingItemAppConfigKeyType(input string) (*AppConfigurationSettingItemAppConfigKeyType, error) {
	vals := map[string]AppConfigurationSettingItemAppConfigKeyType{
		"booleantype": AppConfigurationSettingItemAppConfigKeyTypebooleanType,
		"integertype": AppConfigurationSettingItemAppConfigKeyTypeintegerType,
		"realtype":    AppConfigurationSettingItemAppConfigKeyTyperealType,
		"stringtype":  AppConfigurationSettingItemAppConfigKeyTypestringType,
		"tokentype":   AppConfigurationSettingItemAppConfigKeyTypetokenType,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppConfigurationSettingItemAppConfigKeyType(input)
	return &out, nil
}
