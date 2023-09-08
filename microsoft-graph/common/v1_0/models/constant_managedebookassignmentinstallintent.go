package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEBookAssignmentInstallIntent string

const (
	ManagedEBookAssignmentInstallIntentavailable                  ManagedEBookAssignmentInstallIntent = "Available"
	ManagedEBookAssignmentInstallIntentavailableWithoutEnrollment ManagedEBookAssignmentInstallIntent = "AvailableWithoutEnrollment"
	ManagedEBookAssignmentInstallIntentrequired                   ManagedEBookAssignmentInstallIntent = "Required"
	ManagedEBookAssignmentInstallIntentuninstall                  ManagedEBookAssignmentInstallIntent = "Uninstall"
)

func PossibleValuesForManagedEBookAssignmentInstallIntent() []string {
	return []string{
		string(ManagedEBookAssignmentInstallIntentavailable),
		string(ManagedEBookAssignmentInstallIntentavailableWithoutEnrollment),
		string(ManagedEBookAssignmentInstallIntentrequired),
		string(ManagedEBookAssignmentInstallIntentuninstall),
	}
}

func parseManagedEBookAssignmentInstallIntent(input string) (*ManagedEBookAssignmentInstallIntent, error) {
	vals := map[string]ManagedEBookAssignmentInstallIntent{
		"available":                  ManagedEBookAssignmentInstallIntentavailable,
		"availablewithoutenrollment": ManagedEBookAssignmentInstallIntentavailableWithoutEnrollment,
		"required":                   ManagedEBookAssignmentInstallIntentrequired,
		"uninstall":                  ManagedEBookAssignmentInstallIntentuninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedEBookAssignmentInstallIntent(input)
	return &out, nil
}
