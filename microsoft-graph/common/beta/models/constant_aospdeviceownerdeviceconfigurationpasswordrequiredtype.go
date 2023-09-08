package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerDeviceConfigurationPasswordRequiredType string

const (
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphabetic              AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "Alphabetic"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphanumeric            AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "Alphanumeric"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "AlphanumericWithSymbols"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTypecustomPassword          AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "CustomPassword"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTypedeviceDefault           AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "DeviceDefault"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTypelowSecurityBiometric    AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "LowSecurityBiometric"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTypenumeric                 AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "Numeric"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTypenumericComplex          AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "NumericComplex"
	AospDeviceOwnerDeviceConfigurationPasswordRequiredTyperequired                AospDeviceOwnerDeviceConfigurationPasswordRequiredType = "Required"
)

func PossibleValuesForAospDeviceOwnerDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphabetic),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphanumeric),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTypecustomPassword),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTypedeviceDefault),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTypelowSecurityBiometric),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTypenumeric),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTypenumericComplex),
		string(AospDeviceOwnerDeviceConfigurationPasswordRequiredTyperequired),
	}
}

func parseAospDeviceOwnerDeviceConfigurationPasswordRequiredType(input string) (*AospDeviceOwnerDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AospDeviceOwnerDeviceConfigurationPasswordRequiredType{
		"alphabetic":              AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphabetic,
		"alphanumeric":            AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AospDeviceOwnerDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols,
		"custompassword":          AospDeviceOwnerDeviceConfigurationPasswordRequiredTypecustomPassword,
		"devicedefault":           AospDeviceOwnerDeviceConfigurationPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AospDeviceOwnerDeviceConfigurationPasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AospDeviceOwnerDeviceConfigurationPasswordRequiredTypenumeric,
		"numericcomplex":          AospDeviceOwnerDeviceConfigurationPasswordRequiredTypenumericComplex,
		"required":                AospDeviceOwnerDeviceConfigurationPasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
