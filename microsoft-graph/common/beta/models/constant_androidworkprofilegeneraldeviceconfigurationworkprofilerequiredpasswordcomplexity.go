package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexityhigh   AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "High"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitylow    AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "Low"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitymedium AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "Medium"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitynone   AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity = "None"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexityhigh),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitylow),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitymedium),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitynone),
	}
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity{
		"high":   AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexityhigh,
		"low":    AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitylow,
		"medium": AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitymedium,
		"none":   AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexitynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity(input)
	return &out, nil
}
