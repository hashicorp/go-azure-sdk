package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType string

const (
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "AlphanumericWithSymbols"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphabetic       AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "AtLeastAlphabetic"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphanumeric     AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "AtLeastAlphanumeric"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastNumeric          AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "AtLeastNumeric"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault           AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "DeviceDefault"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric    AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "LowSecurityBiometric"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypenumericComplex          AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "NumericComplex"
	AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTyperequired                AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType = "Required"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationPasswordRequiredType() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphabetic),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphanumeric),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastNumeric),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypenumericComplex),
		string(AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTyperequired),
	}
}

func parseAndroidForWorkGeneralDeviceConfigurationPasswordRequiredType(input string) (*AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType{
		"alphanumericwithsymbols": AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypealphanumericWithSymbols,
		"atleastalphabetic":       AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphabetic,
		"atleastalphanumeric":     AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastAlphanumeric,
		"atleastnumeric":          AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypeatLeastNumeric,
		"devicedefault":           AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypelowSecurityBiometric,
		"numericcomplex":          AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTypenumericComplex,
		"required":                AndroidForWorkGeneralDeviceConfigurationPasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType(input)
	return &out, nil
}
