package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AlphanumericWithSymbols"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphabetic       AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AtLeastAlphabetic"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphanumeric     AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AtLeastAlphanumeric"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastNumeric          AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "AtLeastNumeric"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault           AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "DeviceDefault"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric    AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "LowSecurityBiometric"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex          AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "NumericComplex"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired                AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType = "Required"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphabetic),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphanumeric),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastNumeric),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired),
	}
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType{
		"alphanumericwithsymbols": AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypealphanumericWithSymbols,
		"atleastalphabetic":       AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphabetic,
		"atleastalphanumeric":     AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastAlphanumeric,
		"atleastnumeric":          AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypeatLeastNumeric,
		"devicedefault":           AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypelowSecurityBiometric,
		"numericcomplex":          AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTypenumericComplex,
		"required":                AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType(input)
	return &out, nil
}
