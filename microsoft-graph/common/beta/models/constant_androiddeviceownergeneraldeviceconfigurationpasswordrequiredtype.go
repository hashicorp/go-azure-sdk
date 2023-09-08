package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphabetic              AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "Alphabetic"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphanumeric            AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "Alphanumeric"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "AlphanumericWithSymbols"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypecustomPassword          AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "CustomPassword"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "DeviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "LowSecurityBiometric"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypenumeric                 AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "Numeric"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypenumericComplex          AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "NumericComplex"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTyperequired                AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType = "Required"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphabetic),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphanumeric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypecustomPassword),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypenumeric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypenumericComplex),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTyperequired),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType{
		"alphabetic":              AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphabetic,
		"alphanumeric":            AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols,
		"custompassword":          AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypecustomPassword,
		"devicedefault":           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypenumeric,
		"numericcomplex":          AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTypenumericComplex,
		"required":                AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
