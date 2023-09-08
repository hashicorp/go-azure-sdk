package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceShellScriptRunAsAccount string

const (
	DeviceShellScriptRunAsAccountsystem DeviceShellScriptRunAsAccount = "System"
	DeviceShellScriptRunAsAccountuser   DeviceShellScriptRunAsAccount = "User"
)

func PossibleValuesForDeviceShellScriptRunAsAccount() []string {
	return []string{
		string(DeviceShellScriptRunAsAccountsystem),
		string(DeviceShellScriptRunAsAccountuser),
	}
}

func parseDeviceShellScriptRunAsAccount(input string) (*DeviceShellScriptRunAsAccount, error) {
	vals := map[string]DeviceShellScriptRunAsAccount{
		"system": DeviceShellScriptRunAsAccountsystem,
		"user":   DeviceShellScriptRunAsAccountuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceShellScriptRunAsAccount(input)
	return &out, nil
}
