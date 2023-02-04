package runasaccounts

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialType string

const (
	CredentialTypeHyperVFabric  CredentialType = "HyperVFabric"
	CredentialTypeLinuxGuest    CredentialType = "LinuxGuest"
	CredentialTypeLinuxServer   CredentialType = "LinuxServer"
	CredentialTypeVMwareFabric  CredentialType = "VMwareFabric"
	CredentialTypeWindowsGuest  CredentialType = "WindowsGuest"
	CredentialTypeWindowsServer CredentialType = "WindowsServer"
)

func PossibleValuesForCredentialType() []string {
	return []string{
		string(CredentialTypeHyperVFabric),
		string(CredentialTypeLinuxGuest),
		string(CredentialTypeLinuxServer),
		string(CredentialTypeVMwareFabric),
		string(CredentialTypeWindowsGuest),
		string(CredentialTypeWindowsServer),
	}
}

func parseCredentialType(input string) (*CredentialType, error) {
	vals := map[string]CredentialType{
		"hypervfabric":  CredentialTypeHyperVFabric,
		"linuxguest":    CredentialTypeLinuxGuest,
		"linuxserver":   CredentialTypeLinuxServer,
		"vmwarefabric":  CredentialTypeVMwareFabric,
		"windowsguest":  CredentialTypeWindowsGuest,
		"windowsserver": CredentialTypeWindowsServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CredentialType(input)
	return &out, nil
}
