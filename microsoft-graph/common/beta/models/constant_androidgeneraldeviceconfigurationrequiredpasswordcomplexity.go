package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfigurationRequiredPasswordComplexity string

const (
	AndroidGeneralDeviceConfigurationRequiredPasswordComplexityhigh   AndroidGeneralDeviceConfigurationRequiredPasswordComplexity = "High"
	AndroidGeneralDeviceConfigurationRequiredPasswordComplexitylow    AndroidGeneralDeviceConfigurationRequiredPasswordComplexity = "Low"
	AndroidGeneralDeviceConfigurationRequiredPasswordComplexitymedium AndroidGeneralDeviceConfigurationRequiredPasswordComplexity = "Medium"
	AndroidGeneralDeviceConfigurationRequiredPasswordComplexitynone   AndroidGeneralDeviceConfigurationRequiredPasswordComplexity = "None"
)

func PossibleValuesForAndroidGeneralDeviceConfigurationRequiredPasswordComplexity() []string {
	return []string{
		string(AndroidGeneralDeviceConfigurationRequiredPasswordComplexityhigh),
		string(AndroidGeneralDeviceConfigurationRequiredPasswordComplexitylow),
		string(AndroidGeneralDeviceConfigurationRequiredPasswordComplexitymedium),
		string(AndroidGeneralDeviceConfigurationRequiredPasswordComplexitynone),
	}
}

func parseAndroidGeneralDeviceConfigurationRequiredPasswordComplexity(input string) (*AndroidGeneralDeviceConfigurationRequiredPasswordComplexity, error) {
	vals := map[string]AndroidGeneralDeviceConfigurationRequiredPasswordComplexity{
		"high":   AndroidGeneralDeviceConfigurationRequiredPasswordComplexityhigh,
		"low":    AndroidGeneralDeviceConfigurationRequiredPasswordComplexitylow,
		"medium": AndroidGeneralDeviceConfigurationRequiredPasswordComplexitymedium,
		"none":   AndroidGeneralDeviceConfigurationRequiredPasswordComplexitynone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidGeneralDeviceConfigurationRequiredPasswordComplexity(input)
	return &out, nil
}
