package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceManagementFeatures string

const (
	WindowsManagedDeviceManagementFeaturesmicrosoftManagedDesktop WindowsManagedDeviceManagementFeatures = "MicrosoftManagedDesktop"
	WindowsManagedDeviceManagementFeaturesnone                    WindowsManagedDeviceManagementFeatures = "None"
)

func PossibleValuesForWindowsManagedDeviceManagementFeatures() []string {
	return []string{
		string(WindowsManagedDeviceManagementFeaturesmicrosoftManagedDesktop),
		string(WindowsManagedDeviceManagementFeaturesnone),
	}
}

func parseWindowsManagedDeviceManagementFeatures(input string) (*WindowsManagedDeviceManagementFeatures, error) {
	vals := map[string]WindowsManagedDeviceManagementFeatures{
		"microsoftmanageddesktop": WindowsManagedDeviceManagementFeaturesmicrosoftManagedDesktop,
		"none":                    WindowsManagedDeviceManagementFeaturesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceManagementFeatures(input)
	return &out, nil
}
