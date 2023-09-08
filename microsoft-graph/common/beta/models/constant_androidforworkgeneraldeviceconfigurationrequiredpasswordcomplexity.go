package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity string

const (
	AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexityhigh   AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity = "High"
	AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitylow    AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity = "Low"
	AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitymedium AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity = "Medium"
	AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitynone   AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity = "None"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexityhigh),
		string(AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitylow),
		string(AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitymedium),
		string(AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitynone),
	}
}

func parseAndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity(input string) (*AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity{
		"high":   AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexityhigh,
		"low":    AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitylow,
		"medium": AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitymedium,
		"none":   AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexitynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity(input)
	return &out, nil
}
