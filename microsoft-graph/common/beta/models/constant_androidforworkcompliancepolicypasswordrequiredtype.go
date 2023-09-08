package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCompliancePolicyPasswordRequiredType string

const (
	AndroidForWorkCompliancePolicyPasswordRequiredTypealphabetic              AndroidForWorkCompliancePolicyPasswordRequiredType = "Alphabetic"
	AndroidForWorkCompliancePolicyPasswordRequiredTypealphanumeric            AndroidForWorkCompliancePolicyPasswordRequiredType = "Alphanumeric"
	AndroidForWorkCompliancePolicyPasswordRequiredTypealphanumericWithSymbols AndroidForWorkCompliancePolicyPasswordRequiredType = "AlphanumericWithSymbols"
	AndroidForWorkCompliancePolicyPasswordRequiredTypeany                     AndroidForWorkCompliancePolicyPasswordRequiredType = "Any"
	AndroidForWorkCompliancePolicyPasswordRequiredTypedeviceDefault           AndroidForWorkCompliancePolicyPasswordRequiredType = "DeviceDefault"
	AndroidForWorkCompliancePolicyPasswordRequiredTypelowSecurityBiometric    AndroidForWorkCompliancePolicyPasswordRequiredType = "LowSecurityBiometric"
	AndroidForWorkCompliancePolicyPasswordRequiredTypenumeric                 AndroidForWorkCompliancePolicyPasswordRequiredType = "Numeric"
	AndroidForWorkCompliancePolicyPasswordRequiredTypenumericComplex          AndroidForWorkCompliancePolicyPasswordRequiredType = "NumericComplex"
)

func PossibleValuesForAndroidForWorkCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AndroidForWorkCompliancePolicyPasswordRequiredTypealphabetic),
		string(AndroidForWorkCompliancePolicyPasswordRequiredTypealphanumeric),
		string(AndroidForWorkCompliancePolicyPasswordRequiredTypealphanumericWithSymbols),
		string(AndroidForWorkCompliancePolicyPasswordRequiredTypeany),
		string(AndroidForWorkCompliancePolicyPasswordRequiredTypedeviceDefault),
		string(AndroidForWorkCompliancePolicyPasswordRequiredTypelowSecurityBiometric),
		string(AndroidForWorkCompliancePolicyPasswordRequiredTypenumeric),
		string(AndroidForWorkCompliancePolicyPasswordRequiredTypenumericComplex),
	}
}

func parseAndroidForWorkCompliancePolicyPasswordRequiredType(input string) (*AndroidForWorkCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AndroidForWorkCompliancePolicyPasswordRequiredType{
		"alphabetic":              AndroidForWorkCompliancePolicyPasswordRequiredTypealphabetic,
		"alphanumeric":            AndroidForWorkCompliancePolicyPasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AndroidForWorkCompliancePolicyPasswordRequiredTypealphanumericWithSymbols,
		"any":                     AndroidForWorkCompliancePolicyPasswordRequiredTypeany,
		"devicedefault":           AndroidForWorkCompliancePolicyPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidForWorkCompliancePolicyPasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AndroidForWorkCompliancePolicyPasswordRequiredTypenumeric,
		"numericcomplex":          AndroidForWorkCompliancePolicyPasswordRequiredTypenumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
