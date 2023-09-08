package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyPasswordRequiredType string

const (
	AndroidCompliancePolicyPasswordRequiredTypealphabetic              AndroidCompliancePolicyPasswordRequiredType = "Alphabetic"
	AndroidCompliancePolicyPasswordRequiredTypealphanumeric            AndroidCompliancePolicyPasswordRequiredType = "Alphanumeric"
	AndroidCompliancePolicyPasswordRequiredTypealphanumericWithSymbols AndroidCompliancePolicyPasswordRequiredType = "AlphanumericWithSymbols"
	AndroidCompliancePolicyPasswordRequiredTypeany                     AndroidCompliancePolicyPasswordRequiredType = "Any"
	AndroidCompliancePolicyPasswordRequiredTypedeviceDefault           AndroidCompliancePolicyPasswordRequiredType = "DeviceDefault"
	AndroidCompliancePolicyPasswordRequiredTypelowSecurityBiometric    AndroidCompliancePolicyPasswordRequiredType = "LowSecurityBiometric"
	AndroidCompliancePolicyPasswordRequiredTypenumeric                 AndroidCompliancePolicyPasswordRequiredType = "Numeric"
	AndroidCompliancePolicyPasswordRequiredTypenumericComplex          AndroidCompliancePolicyPasswordRequiredType = "NumericComplex"
)

func PossibleValuesForAndroidCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AndroidCompliancePolicyPasswordRequiredTypealphabetic),
		string(AndroidCompliancePolicyPasswordRequiredTypealphanumeric),
		string(AndroidCompliancePolicyPasswordRequiredTypealphanumericWithSymbols),
		string(AndroidCompliancePolicyPasswordRequiredTypeany),
		string(AndroidCompliancePolicyPasswordRequiredTypedeviceDefault),
		string(AndroidCompliancePolicyPasswordRequiredTypelowSecurityBiometric),
		string(AndroidCompliancePolicyPasswordRequiredTypenumeric),
		string(AndroidCompliancePolicyPasswordRequiredTypenumericComplex),
	}
}

func parseAndroidCompliancePolicyPasswordRequiredType(input string) (*AndroidCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AndroidCompliancePolicyPasswordRequiredType{
		"alphabetic":              AndroidCompliancePolicyPasswordRequiredTypealphabetic,
		"alphanumeric":            AndroidCompliancePolicyPasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AndroidCompliancePolicyPasswordRequiredTypealphanumericWithSymbols,
		"any":                     AndroidCompliancePolicyPasswordRequiredTypeany,
		"devicedefault":           AndroidCompliancePolicyPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidCompliancePolicyPasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AndroidCompliancePolicyPasswordRequiredTypenumeric,
		"numericcomplex":          AndroidCompliancePolicyPasswordRequiredTypenumericComplex,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
