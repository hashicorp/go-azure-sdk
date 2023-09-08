package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode string

const (
	AndroidManagedStoreAppAssignmentSettingsAutoUpdateModedefault   AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode = "Default"
	AndroidManagedStoreAppAssignmentSettingsAutoUpdateModepostponed AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode = "Postponed"
	AndroidManagedStoreAppAssignmentSettingsAutoUpdateModepriority  AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode = "Priority"
)

func PossibleValuesForAndroidManagedStoreAppAssignmentSettingsAutoUpdateMode() []string {
	return []string{
		string(AndroidManagedStoreAppAssignmentSettingsAutoUpdateModedefault),
		string(AndroidManagedStoreAppAssignmentSettingsAutoUpdateModepostponed),
		string(AndroidManagedStoreAppAssignmentSettingsAutoUpdateModepriority),
	}
}

func parseAndroidManagedStoreAppAssignmentSettingsAutoUpdateMode(input string) (*AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode, error) {
	vals := map[string]AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode{
		"default":   AndroidManagedStoreAppAssignmentSettingsAutoUpdateModedefault,
		"postponed": AndroidManagedStoreAppAssignmentSettingsAutoUpdateModepostponed,
		"priority":  AndroidManagedStoreAppAssignmentSettingsAutoUpdateModepriority,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAppAssignmentSettingsAutoUpdateMode(input)
	return &out, nil
}
