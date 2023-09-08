package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81GeneralConfigurationPasswordRequiredType string

const (
	WindowsPhone81GeneralConfigurationPasswordRequiredTypealphanumeric  WindowsPhone81GeneralConfigurationPasswordRequiredType = "Alphanumeric"
	WindowsPhone81GeneralConfigurationPasswordRequiredTypedeviceDefault WindowsPhone81GeneralConfigurationPasswordRequiredType = "DeviceDefault"
	WindowsPhone81GeneralConfigurationPasswordRequiredTypenumeric       WindowsPhone81GeneralConfigurationPasswordRequiredType = "Numeric"
)

func PossibleValuesForWindowsPhone81GeneralConfigurationPasswordRequiredType() []string {
	return []string{
		string(WindowsPhone81GeneralConfigurationPasswordRequiredTypealphanumeric),
		string(WindowsPhone81GeneralConfigurationPasswordRequiredTypedeviceDefault),
		string(WindowsPhone81GeneralConfigurationPasswordRequiredTypenumeric),
	}
}

func parseWindowsPhone81GeneralConfigurationPasswordRequiredType(input string) (*WindowsPhone81GeneralConfigurationPasswordRequiredType, error) {
	vals := map[string]WindowsPhone81GeneralConfigurationPasswordRequiredType{
		"alphanumeric":  WindowsPhone81GeneralConfigurationPasswordRequiredTypealphanumeric,
		"devicedefault": WindowsPhone81GeneralConfigurationPasswordRequiredTypedeviceDefault,
		"numeric":       WindowsPhone81GeneralConfigurationPasswordRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81GeneralConfigurationPasswordRequiredType(input)
	return &out, nil
}
