package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10MobileCompliancePolicyPasswordRequiredType string

const (
	Windows10MobileCompliancePolicyPasswordRequiredTypealphanumeric  Windows10MobileCompliancePolicyPasswordRequiredType = "Alphanumeric"
	Windows10MobileCompliancePolicyPasswordRequiredTypedeviceDefault Windows10MobileCompliancePolicyPasswordRequiredType = "DeviceDefault"
	Windows10MobileCompliancePolicyPasswordRequiredTypenumeric       Windows10MobileCompliancePolicyPasswordRequiredType = "Numeric"
)

func PossibleValuesForWindows10MobileCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(Windows10MobileCompliancePolicyPasswordRequiredTypealphanumeric),
		string(Windows10MobileCompliancePolicyPasswordRequiredTypedeviceDefault),
		string(Windows10MobileCompliancePolicyPasswordRequiredTypenumeric),
	}
}

func parseWindows10MobileCompliancePolicyPasswordRequiredType(input string) (*Windows10MobileCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]Windows10MobileCompliancePolicyPasswordRequiredType{
		"alphanumeric":  Windows10MobileCompliancePolicyPasswordRequiredTypealphanumeric,
		"devicedefault": Windows10MobileCompliancePolicyPasswordRequiredTypedeviceDefault,
		"numeric":       Windows10MobileCompliancePolicyPasswordRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10MobileCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
