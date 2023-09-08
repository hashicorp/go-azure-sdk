package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AlphanumericWithSymbols"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphabetic       AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AtLeastAlphabetic"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphanumeric     AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AtLeastAlphanumeric"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastNumeric          AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AtLeastNumeric"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault           AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "DeviceDefault"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric    AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "LowSecurityBiometric"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex          AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "NumericComplex"
	AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired                AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "Required"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphabetic),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphanumeric),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastNumeric),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired),
	}
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType{
		"alphanumericwithsymbols": AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols,
		"atleastalphabetic":       AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphabetic,
		"atleastalphanumeric":     AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphanumeric,
		"atleastnumeric":          AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastNumeric,
		"devicedefault":           AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric,
		"numericcomplex":          AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex,
		"required":                AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input)
	return &out, nil
}
