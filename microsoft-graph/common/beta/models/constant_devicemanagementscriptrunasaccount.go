package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptRunAsAccount string

const (
	DeviceManagementScriptRunAsAccountsystem DeviceManagementScriptRunAsAccount = "System"
	DeviceManagementScriptRunAsAccountuser   DeviceManagementScriptRunAsAccount = "User"
)

func PossibleValuesForDeviceManagementScriptRunAsAccount() []string {
	return []string{
		string(DeviceManagementScriptRunAsAccountsystem),
		string(DeviceManagementScriptRunAsAccountuser),
	}
}

func parseDeviceManagementScriptRunAsAccount(input string) (*DeviceManagementScriptRunAsAccount, error) {
	vals := map[string]DeviceManagementScriptRunAsAccount{
		"system": DeviceManagementScriptRunAsAccountsystem,
		"user":   DeviceManagementScriptRunAsAccountuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptRunAsAccount(input)
	return &out, nil
}
