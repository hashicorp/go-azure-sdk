package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphabetic              AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "Alphabetic"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumeric            AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "Alphanumeric"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AlphanumericWithSymbols"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypecustomPassword          AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "CustomPassword"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault           AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "DeviceDefault"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric    AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "LowSecurityBiometric"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumeric                 AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "Numeric"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex          AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "NumericComplex"
	AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired                AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "Required"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphabetic),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumeric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypecustomPassword),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumeric),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType{
		"alphabetic":              AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphabetic,
		"alphanumeric":            AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols,
		"custompassword":          AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypecustomPassword,
		"devicedefault":           AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumeric,
		"numericcomplex":          AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex,
		"required":                AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input)
	return &out, nil
}
