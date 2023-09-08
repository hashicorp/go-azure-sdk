package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfigurationPasswordRequiredType string

const (
	AndroidGeneralDeviceConfigurationPasswordRequiredTypealphabetic              AndroidGeneralDeviceConfigurationPasswordRequiredType = "Alphabetic"
	AndroidGeneralDeviceConfigurationPasswordRequiredTypealphanumeric            AndroidGeneralDeviceConfigurationPasswordRequiredType = "Alphanumeric"
	AndroidGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols AndroidGeneralDeviceConfigurationPasswordRequiredType = "AlphanumericWithSymbols"
	AndroidGeneralDeviceConfigurationPasswordRequiredTypeany                     AndroidGeneralDeviceConfigurationPasswordRequiredType = "Any"
	AndroidGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault           AndroidGeneralDeviceConfigurationPasswordRequiredType = "DeviceDefault"
	AndroidGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric    AndroidGeneralDeviceConfigurationPasswordRequiredType = "LowSecurityBiometric"
	AndroidGeneralDeviceConfigurationPasswordRequiredTypenumeric                 AndroidGeneralDeviceConfigurationPasswordRequiredType = "Numeric"
	AndroidGeneralDeviceConfigurationPasswordRequiredTypenumericComplex          AndroidGeneralDeviceConfigurationPasswordRequiredType = "NumericComplex"
)

func PossibleValuesForAndroidGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AndroidGeneralDeviceConfigurationPasswordRequiredTypealphabetic),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredTypealphanumeric),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredTypeany),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredTypenumeric),
		string(AndroidGeneralDeviceConfigurationPasswordRequiredTypenumericComplex),
	}
}

func parseAndroidGeneralDeviceConfigurationPasswordRequiredType(input string) (*AndroidGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AndroidGeneralDeviceConfigurationPasswordRequiredType{
		"alphabetic":              AndroidGeneralDeviceConfigurationPasswordRequiredTypealphabetic,
		"alphanumeric":            AndroidGeneralDeviceConfigurationPasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AndroidGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols,
		"any":                     AndroidGeneralDeviceConfigurationPasswordRequiredTypeany,
		"devicedefault":           AndroidGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AndroidGeneralDeviceConfigurationPasswordRequiredTypenumeric,
		"numericcomplex":          AndroidGeneralDeviceConfigurationPasswordRequiredTypenumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
