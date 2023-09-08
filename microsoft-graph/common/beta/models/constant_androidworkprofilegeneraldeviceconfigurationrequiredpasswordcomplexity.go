package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity string

const (
	AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexityhigh   AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity = "High"
	AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitylow    AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity = "Low"
	AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitymedium AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity = "Medium"
	AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitynone   AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity = "None"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexityhigh),
		string(AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitylow),
		string(AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitymedium),
		string(AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitynone),
	}
}

func parseAndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity(input string) (*AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity{
		"high":   AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexityhigh,
		"low":    AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitylow,
		"medium": AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitymedium,
		"none":   AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexitynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationRequiredPasswordComplexity(input)
	return &out, nil
}
