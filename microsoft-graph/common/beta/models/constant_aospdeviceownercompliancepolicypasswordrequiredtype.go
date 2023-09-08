package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerCompliancePolicyPasswordRequiredType string

const (
	AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphabetic              AospDeviceOwnerCompliancePolicyPasswordRequiredType = "Alphabetic"
	AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumeric            AospDeviceOwnerCompliancePolicyPasswordRequiredType = "Alphanumeric"
	AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumericWithSymbols AospDeviceOwnerCompliancePolicyPasswordRequiredType = "AlphanumericWithSymbols"
	AospDeviceOwnerCompliancePolicyPasswordRequiredTypecustomPassword          AospDeviceOwnerCompliancePolicyPasswordRequiredType = "CustomPassword"
	AospDeviceOwnerCompliancePolicyPasswordRequiredTypedeviceDefault           AospDeviceOwnerCompliancePolicyPasswordRequiredType = "DeviceDefault"
	AospDeviceOwnerCompliancePolicyPasswordRequiredTypelowSecurityBiometric    AospDeviceOwnerCompliancePolicyPasswordRequiredType = "LowSecurityBiometric"
	AospDeviceOwnerCompliancePolicyPasswordRequiredTypenumeric                 AospDeviceOwnerCompliancePolicyPasswordRequiredType = "Numeric"
	AospDeviceOwnerCompliancePolicyPasswordRequiredTypenumericComplex          AospDeviceOwnerCompliancePolicyPasswordRequiredType = "NumericComplex"
	AospDeviceOwnerCompliancePolicyPasswordRequiredTyperequired                AospDeviceOwnerCompliancePolicyPasswordRequiredType = "Required"
)

func PossibleValuesForAospDeviceOwnerCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphabetic),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumeric),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumericWithSymbols),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTypecustomPassword),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTypedeviceDefault),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTypelowSecurityBiometric),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTypenumeric),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTypenumericComplex),
		string(AospDeviceOwnerCompliancePolicyPasswordRequiredTyperequired),
	}
}

func parseAospDeviceOwnerCompliancePolicyPasswordRequiredType(input string) (*AospDeviceOwnerCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]AospDeviceOwnerCompliancePolicyPasswordRequiredType{
		"alphabetic":              AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphabetic,
		"alphanumeric":            AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumeric,
		"alphanumericwithsymbols": AospDeviceOwnerCompliancePolicyPasswordRequiredTypealphanumericWithSymbols,
		"custompassword":          AospDeviceOwnerCompliancePolicyPasswordRequiredTypecustomPassword,
		"devicedefault":           AospDeviceOwnerCompliancePolicyPasswordRequiredTypedeviceDefault,
		"lowsecuritybiometric":    AospDeviceOwnerCompliancePolicyPasswordRequiredTypelowSecurityBiometric,
		"numeric":                 AospDeviceOwnerCompliancePolicyPasswordRequiredTypenumeric,
		"numericcomplex":          AospDeviceOwnerCompliancePolicyPasswordRequiredTypenumericComplex,
		"required":                AospDeviceOwnerCompliancePolicyPasswordRequiredTyperequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
