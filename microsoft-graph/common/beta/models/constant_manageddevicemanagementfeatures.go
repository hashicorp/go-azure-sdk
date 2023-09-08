package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceManagementFeatures string

const (
	ManagedDeviceManagementFeaturesmicrosoftManagedDesktop ManagedDeviceManagementFeatures = "MicrosoftManagedDesktop"
	ManagedDeviceManagementFeaturesnone                    ManagedDeviceManagementFeatures = "None"
)

func PossibleValuesForManagedDeviceManagementFeatures() []string {
	return []string{
		string(ManagedDeviceManagementFeaturesmicrosoftManagedDesktop),
		string(ManagedDeviceManagementFeaturesnone),
	}
}

func parseManagedDeviceManagementFeatures(input string) (*ManagedDeviceManagementFeatures, error) {
	vals := map[string]ManagedDeviceManagementFeatures{
		"microsoftmanageddesktop": ManagedDeviceManagementFeaturesmicrosoftManagedDesktop,
		"none":                    ManagedDeviceManagementFeaturesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementFeatures(input)
	return &out, nil
}
