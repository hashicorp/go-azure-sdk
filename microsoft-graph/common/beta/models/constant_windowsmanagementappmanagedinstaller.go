package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementAppManagedInstaller string

const (
	WindowsManagementAppManagedInstallerdisabled WindowsManagementAppManagedInstaller = "Disabled"
	WindowsManagementAppManagedInstallerenabled  WindowsManagementAppManagedInstaller = "Enabled"
)

func PossibleValuesForWindowsManagementAppManagedInstaller() []string {
	return []string{
		string(WindowsManagementAppManagedInstallerdisabled),
		string(WindowsManagementAppManagedInstallerenabled),
	}
}

func parseWindowsManagementAppManagedInstaller(input string) (*WindowsManagementAppManagedInstaller, error) {
	vals := map[string]WindowsManagementAppManagedInstaller{
		"disabled": WindowsManagementAppManagedInstallerdisabled,
		"enabled":  WindowsManagementAppManagedInstallerenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagementAppManagedInstaller(input)
	return &out, nil
}
