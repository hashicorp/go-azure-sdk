package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeShellScriptRunAsAccount string

const (
	DeviceCustomAttributeShellScriptRunAsAccountsystem DeviceCustomAttributeShellScriptRunAsAccount = "System"
	DeviceCustomAttributeShellScriptRunAsAccountuser   DeviceCustomAttributeShellScriptRunAsAccount = "User"
)

func PossibleValuesForDeviceCustomAttributeShellScriptRunAsAccount() []string {
	return []string{
		string(DeviceCustomAttributeShellScriptRunAsAccountsystem),
		string(DeviceCustomAttributeShellScriptRunAsAccountuser),
	}
}

func parseDeviceCustomAttributeShellScriptRunAsAccount(input string) (*DeviceCustomAttributeShellScriptRunAsAccount, error) {
	vals := map[string]DeviceCustomAttributeShellScriptRunAsAccount{
		"system": DeviceCustomAttributeShellScriptRunAsAccountsystem,
		"user":   DeviceCustomAttributeShellScriptRunAsAccountuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCustomAttributeShellScriptRunAsAccount(input)
	return &out, nil
}
