package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptRunAsAccount string

const (
	DeviceHealthScriptRunAsAccountsystem DeviceHealthScriptRunAsAccount = "System"
	DeviceHealthScriptRunAsAccountuser   DeviceHealthScriptRunAsAccount = "User"
)

func PossibleValuesForDeviceHealthScriptRunAsAccount() []string {
	return []string{
		string(DeviceHealthScriptRunAsAccountsystem),
		string(DeviceHealthScriptRunAsAccountuser),
	}
}

func parseDeviceHealthScriptRunAsAccount(input string) (*DeviceHealthScriptRunAsAccount, error) {
	vals := map[string]DeviceHealthScriptRunAsAccount{
		"system": DeviceHealthScriptRunAsAccountsystem,
		"user":   DeviceHealthScriptRunAsAccountuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptRunAsAccount(input)
	return &out, nil
}
