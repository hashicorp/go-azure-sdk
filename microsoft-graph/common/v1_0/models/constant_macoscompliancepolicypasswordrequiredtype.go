package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCompliancePolicyPasswordRequiredType string

const (
	MacOSCompliancePolicyPasswordRequiredTypealphanumeric  MacOSCompliancePolicyPasswordRequiredType = "Alphanumeric"
	MacOSCompliancePolicyPasswordRequiredTypedeviceDefault MacOSCompliancePolicyPasswordRequiredType = "DeviceDefault"
	MacOSCompliancePolicyPasswordRequiredTypenumeric       MacOSCompliancePolicyPasswordRequiredType = "Numeric"
)

func PossibleValuesForMacOSCompliancePolicyPasswordRequiredType() []string {
	return []string{
		string(MacOSCompliancePolicyPasswordRequiredTypealphanumeric),
		string(MacOSCompliancePolicyPasswordRequiredTypedeviceDefault),
		string(MacOSCompliancePolicyPasswordRequiredTypenumeric),
	}
}

func parseMacOSCompliancePolicyPasswordRequiredType(input string) (*MacOSCompliancePolicyPasswordRequiredType, error) {
	vals := map[string]MacOSCompliancePolicyPasswordRequiredType{
		"alphanumeric":  MacOSCompliancePolicyPasswordRequiredTypealphanumeric,
		"devicedefault": MacOSCompliancePolicyPasswordRequiredTypedeviceDefault,
		"numeric":       MacOSCompliancePolicyPasswordRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCompliancePolicyPasswordRequiredType(input)
	return &out, nil
}
