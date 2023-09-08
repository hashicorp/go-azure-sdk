package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationPasswordRequiredType string

const (
	Windows81GeneralConfigurationPasswordRequiredTypealphanumeric  Windows81GeneralConfigurationPasswordRequiredType = "Alphanumeric"
	Windows81GeneralConfigurationPasswordRequiredTypedeviceDefault Windows81GeneralConfigurationPasswordRequiredType = "DeviceDefault"
	Windows81GeneralConfigurationPasswordRequiredTypenumeric       Windows81GeneralConfigurationPasswordRequiredType = "Numeric"
)

func PossibleValuesForWindows81GeneralConfigurationPasswordRequiredType() []string {
	return []string{
		string(Windows81GeneralConfigurationPasswordRequiredTypealphanumeric),
		string(Windows81GeneralConfigurationPasswordRequiredTypedeviceDefault),
		string(Windows81GeneralConfigurationPasswordRequiredTypenumeric),
	}
}

func parseWindows81GeneralConfigurationPasswordRequiredType(input string) (*Windows81GeneralConfigurationPasswordRequiredType, error) {
	vals := map[string]Windows81GeneralConfigurationPasswordRequiredType{
		"alphanumeric":  Windows81GeneralConfigurationPasswordRequiredTypealphanumeric,
		"devicedefault": Windows81GeneralConfigurationPasswordRequiredTypedeviceDefault,
		"numeric":       Windows81GeneralConfigurationPasswordRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationPasswordRequiredType(input)
	return &out, nil
}
