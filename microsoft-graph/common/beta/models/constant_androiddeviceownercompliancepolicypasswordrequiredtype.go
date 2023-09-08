package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCompliancePolicyPasswordRequiredType string

const (
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphabetic              AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "Alphabetic"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumeric            AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "Alphanumeric"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumericWithSymbols AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "AlphanumericWithSymbols"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypecustomPassword          AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "CustomPassword"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypedeviceDefault           AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "DeviceDefault"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypelowSecurityBiometric    AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "LowSecurityBiometric"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypenumeric                 AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "Numeric"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypenumericComplex          AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "NumericComplex"
	AndroidDeviceOwnerCompliancePolicyPasswordRequiredTyperequired                AndroidDeviceOwnerCompliancePolicyPasswordRequiredType = "Required"
)

func PossibleValuesForAndroidDeviceOwnerCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphabetic),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumeric),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumericWithSymbols),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypecustomPassword),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypedeviceDefault),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypelowSecurityBiometric),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypenumeric),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypenumericComplex),
		string(AndroidDeviceOwnerCompliancePolicyPasswordRequiredTyperequired),
	}
}

func parseAndroidDeviceOwnerCompliancePolicyPasswordRequiredType(input string) (*AndroidDeviceOwnerCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AndroidDeviceOwnerCompliancePolicyPasswordRequiredType{
		"alphabetic":              AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphabetic,
		"alphanumeric":            AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumericWithSymbols,
		"custompassword":          AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypecustomPassword,
		"devicedefault":           AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypenumeric,
		"numericcomplex":          AndroidDeviceOwnerCompliancePolicyPasswordRequiredTypenumericComplex,
		"required":                AndroidDeviceOwnerCompliancePolicyPasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
