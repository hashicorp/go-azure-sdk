package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CompliancePolicyPasswordRequiredType string

const (
	Windows81CompliancePolicyPasswordRequiredTypealphanumeric  Windows81CompliancePolicyPasswordRequiredType = "Alphanumeric"
	Windows81CompliancePolicyPasswordRequiredTypedeviceDefault Windows81CompliancePolicyPasswordRequiredType = "DeviceDefault"
	Windows81CompliancePolicyPasswordRequiredTypenumeric       Windows81CompliancePolicyPasswordRequiredType = "Numeric"
)

func PossibleValuesForWindows81CompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(Windows81CompliancePolicyPasswordRequiredTypealphanumeric),
		string(Windows81CompliancePolicyPasswordRequiredTypedeviceDefault),
		string(Windows81CompliancePolicyPasswordRequiredTypenumeric),
	}
}

func parseWindows81CompliancePolicyPasswordRequiredType(input string) (*Windows81CompliancePolicyPasswordRequiredType, error) {
	vals := map[string]Windows81CompliancePolicyPasswordRequiredType{
		"alphanumeric":  Windows81CompliancePolicyPasswordRequiredTypealphanumeric,
		"devicedefault": Windows81CompliancePolicyPasswordRequiredTypedeviceDefault,
		"numeric":       Windows81CompliancePolicyPasswordRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
