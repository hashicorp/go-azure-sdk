package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicyPasswordRequiredType string

const (
	AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphabetic              AndroidWorkProfileCompliancePolicyPasswordRequiredType = "Alphabetic"
	AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphanumeric            AndroidWorkProfileCompliancePolicyPasswordRequiredType = "Alphanumeric"
	AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphanumericWithSymbols AndroidWorkProfileCompliancePolicyPasswordRequiredType = "AlphanumericWithSymbols"
	AndroidWorkProfileCompliancePolicyPasswordRequiredTypeany                     AndroidWorkProfileCompliancePolicyPasswordRequiredType = "Any"
	AndroidWorkProfileCompliancePolicyPasswordRequiredTypedeviceDefault           AndroidWorkProfileCompliancePolicyPasswordRequiredType = "DeviceDefault"
	AndroidWorkProfileCompliancePolicyPasswordRequiredTypelowSecurityBiometric    AndroidWorkProfileCompliancePolicyPasswordRequiredType = "LowSecurityBiometric"
	AndroidWorkProfileCompliancePolicyPasswordRequiredTypenumeric                 AndroidWorkProfileCompliancePolicyPasswordRequiredType = "Numeric"
	AndroidWorkProfileCompliancePolicyPasswordRequiredTypenumericComplex          AndroidWorkProfileCompliancePolicyPasswordRequiredType = "NumericComplex"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphabetic),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphanumeric),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphanumericWithSymbols),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredTypeany),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredTypedeviceDefault),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredTypelowSecurityBiometric),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredTypenumeric),
		string(AndroidWorkProfileCompliancePolicyPasswordRequiredTypenumericComplex),
	}
}

func parseAndroidWorkProfileCompliancePolicyPasswordRequiredType(input string) (*AndroidWorkProfileCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicyPasswordRequiredType{
		"alphabetic":              AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphabetic,
		"alphanumeric":            AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AndroidWorkProfileCompliancePolicyPasswordRequiredTypealphanumericWithSymbols,
		"any":                     AndroidWorkProfileCompliancePolicyPasswordRequiredTypeany,
		"devicedefault":           AndroidWorkProfileCompliancePolicyPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidWorkProfileCompliancePolicyPasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AndroidWorkProfileCompliancePolicyPasswordRequiredTypenumeric,
		"numericcomplex":          AndroidWorkProfileCompliancePolicyPasswordRequiredTypenumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
