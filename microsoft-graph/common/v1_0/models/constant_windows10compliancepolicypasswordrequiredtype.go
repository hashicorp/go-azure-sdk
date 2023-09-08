package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CompliancePolicyPasswordRequiredType string

const (
	Windows10CompliancePolicyPasswordRequiredTypealphanumeric  Windows10CompliancePolicyPasswordRequiredType = "Alphanumeric"
	Windows10CompliancePolicyPasswordRequiredTypedeviceDefault Windows10CompliancePolicyPasswordRequiredType = "DeviceDefault"
	Windows10CompliancePolicyPasswordRequiredTypenumeric       Windows10CompliancePolicyPasswordRequiredType = "Numeric"
)

func PossibleValuesForWindows10CompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(Windows10CompliancePolicyPasswordRequiredTypealphanumeric),
		string(Windows10CompliancePolicyPasswordRequiredTypedeviceDefault),
		string(Windows10CompliancePolicyPasswordRequiredTypenumeric),
	}
}

func parseWindows10CompliancePolicyPasswordRequiredType(input string) (*Windows10CompliancePolicyPasswordRequiredType, error) {
	vals := map[string]Windows10CompliancePolicyPasswordRequiredType{
		"alphanumeric":  Windows10CompliancePolicyPasswordRequiredTypealphanumeric,
		"devicedefault": Windows10CompliancePolicyPasswordRequiredTypedeviceDefault,
		"numeric":       Windows10CompliancePolicyPasswordRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
