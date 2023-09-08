package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSGeneralDeviceConfigurationPasswordRequiredType string

const (
	MacOSGeneralDeviceConfigurationPasswordRequiredTypealphanumeric  MacOSGeneralDeviceConfigurationPasswordRequiredType = "Alphanumeric"
	MacOSGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault MacOSGeneralDeviceConfigurationPasswordRequiredType = "DeviceDefault"
	MacOSGeneralDeviceConfigurationPasswordRequiredTypenumeric       MacOSGeneralDeviceConfigurationPasswordRequiredType = "Numeric"
)

func PossibleValuesForMacOSGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(MacOSGeneralDeviceConfigurationPasswordRequiredTypealphanumeric),
		string(MacOSGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault),
		string(MacOSGeneralDeviceConfigurationPasswordRequiredTypenumeric),
	}
}

func parseMacOSGeneralDeviceConfigurationPasswordRequiredType(input string) (*MacOSGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]MacOSGeneralDeviceConfigurationPasswordRequiredType{
		"alphanumeric":  MacOSGeneralDeviceConfigurationPasswordRequiredTypealphanumeric,
		"devicedefault": MacOSGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault,
		"numeric":       MacOSGeneralDeviceConfigurationPasswordRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
