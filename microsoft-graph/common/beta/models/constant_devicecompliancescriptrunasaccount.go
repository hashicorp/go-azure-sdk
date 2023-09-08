package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRunAsAccount string

const (
	DeviceComplianceScriptRunAsAccountsystem DeviceComplianceScriptRunAsAccount = "System"
	DeviceComplianceScriptRunAsAccountuser   DeviceComplianceScriptRunAsAccount = "User"
)

func PossibleValuesForDeviceComplianceScriptRunAsAccount() []string {
	return []string{
		string(DeviceComplianceScriptRunAsAccountsystem),
		string(DeviceComplianceScriptRunAsAccountuser),
	}
}

func parseDeviceComplianceScriptRunAsAccount(input string) (*DeviceComplianceScriptRunAsAccount, error) {
	vals := map[string]DeviceComplianceScriptRunAsAccount{
		"system": DeviceComplianceScriptRunAsAccountsystem,
		"user":   DeviceComplianceScriptRunAsAccountuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRunAsAccount(input)
	return &out, nil
}
