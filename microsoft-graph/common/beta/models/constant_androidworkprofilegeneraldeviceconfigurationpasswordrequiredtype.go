package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType string

const (
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "AlphanumericWithSymbols"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphabetic       AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "AtLeastAlphabetic"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphanumeric     AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "AtLeastAlphanumeric"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastNumeric          AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "AtLeastNumeric"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault           AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "DeviceDefault"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric    AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "LowSecurityBiometric"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypenumericComplex          AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "NumericComplex"
	AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTyperequired                AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType = "Required"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphabetic),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphanumeric),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastNumeric),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypenumericComplex),
		string(AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTyperequired),
	}
}

func parseAndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType(input string) (*AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType{
		"alphanumericwithsymbols": AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols,
		"atleastalphabetic":       AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphabetic,
		"atleastalphanumeric":     AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphanumeric,
		"atleastnumeric":          AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypeatLeastNumeric,
		"devicedefault":           AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric,
		"numericcomplex":          AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTypenumericComplex,
		"required":                AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
