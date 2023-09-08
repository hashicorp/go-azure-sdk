package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexityhigh   AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "High"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitylow    AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "Low"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitymedium AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "Medium"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitynone   AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "None"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexityhigh),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitylow),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitymedium),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitynone),
	}
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity{
		"high":   AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexityhigh,
		"low":    AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitylow,
		"medium": AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitymedium,
		"none":   AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity(input)
	return &out, nil
}
