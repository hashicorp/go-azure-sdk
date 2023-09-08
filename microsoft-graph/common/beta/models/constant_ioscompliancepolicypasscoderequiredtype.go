package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCompliancePolicyPasscodeRequiredType string

const (
	IosCompliancePolicyPasscodeRequiredTypealphanumeric  IosCompliancePolicyPasscodeRequiredType = "Alphanumeric"
	IosCompliancePolicyPasscodeRequiredTypedeviceDefault IosCompliancePolicyPasscodeRequiredType = "DeviceDefault"
	IosCompliancePolicyPasscodeRequiredTypenumeric       IosCompliancePolicyPasscodeRequiredType = "Numeric"
)

func PossibleValuesForIosCompliancePolicyPasscodeRequiredType() []string {
	return []string{
		string(IosCompliancePolicyPasscodeRequiredTypealphanumeric),
		string(IosCompliancePolicyPasscodeRequiredTypedeviceDefault),
		string(IosCompliancePolicyPasscodeRequiredTypenumeric),
	}
}

func parseIosCompliancePolicyPasscodeRequiredType(input string) (*IosCompliancePolicyPasscodeRequiredType, error) {
	vals := map[string]IosCompliancePolicyPasscodeRequiredType{
		"alphanumeric":  IosCompliancePolicyPasscodeRequiredTypealphanumeric,
		"devicedefault": IosCompliancePolicyPasscodeRequiredTypedeviceDefault,
		"numeric":       IosCompliancePolicyPasscodeRequiredTypenumeric,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCompliancePolicyPasscodeRequiredType(input)
	return &out, nil
}
