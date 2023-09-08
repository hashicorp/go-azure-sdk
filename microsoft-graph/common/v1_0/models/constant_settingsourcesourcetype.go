package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingSourceSourceType string

const (
	SettingSourceSourceTypedeviceConfiguration SettingSourceSourceType = "DeviceConfiguration"
	SettingSourceSourceTypedeviceIntent        SettingSourceSourceType = "DeviceIntent"
)

func PossibleValuesForSettingSourceSourceType() []string {
	return []string{
		string(SettingSourceSourceTypedeviceConfiguration),
		string(SettingSourceSourceTypedeviceIntent),
	}
}

func parseSettingSourceSourceType(input string) (*SettingSourceSourceType, error) {
	vals := map[string]SettingSourceSourceType{
		"deviceconfiguration": SettingSourceSourceTypedeviceConfiguration,
		"deviceintent":        SettingSourceSourceTypedeviceIntent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingSourceSourceType(input)
	return &out, nil
}
