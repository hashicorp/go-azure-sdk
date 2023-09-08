package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CompliancePolicyPasswordRequiredType string

const (
	WindowsPhone81CompliancePolicyPasswordRequiredTypealphanumeric  WindowsPhone81CompliancePolicyPasswordRequiredType = "Alphanumeric"
	WindowsPhone81CompliancePolicyPasswordRequiredTypedeviceDefault WindowsPhone81CompliancePolicyPasswordRequiredType = "DeviceDefault"
	WindowsPhone81CompliancePolicyPasswordRequiredTypenumeric       WindowsPhone81CompliancePolicyPasswordRequiredType = "Numeric"
)

func PossibleValuesForWindowsPhone81CompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(WindowsPhone81CompliancePolicyPasswordRequiredTypealphanumeric),
		string(WindowsPhone81CompliancePolicyPasswordRequiredTypedeviceDefault),
		string(WindowsPhone81CompliancePolicyPasswordRequiredTypenumeric),
	}
}

func parseWindowsPhone81CompliancePolicyPasswordRequiredType(input string) (*WindowsPhone81CompliancePolicyPasswordRequiredType, error) {
	vals := map[string]WindowsPhone81CompliancePolicyPasswordRequiredType{
		"alphanumeric":  WindowsPhone81CompliancePolicyPasswordRequiredTypealphanumeric,
		"devicedefault": WindowsPhone81CompliancePolicyPasswordRequiredTypedeviceDefault,
		"numeric":       WindowsPhone81CompliancePolicyPasswordRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
