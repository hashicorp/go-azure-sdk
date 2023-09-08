package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertStatus string

const (
	ManagedTenantsManagedTenantAlertStatusdismissed  ManagedTenantsManagedTenantAlertStatus = "Dismissed"
	ManagedTenantsManagedTenantAlertStatusinProgress ManagedTenantsManagedTenantAlertStatus = "InProgress"
	ManagedTenantsManagedTenantAlertStatusnewAlert   ManagedTenantsManagedTenantAlertStatus = "NewAlert"
	ManagedTenantsManagedTenantAlertStatusresolved   ManagedTenantsManagedTenantAlertStatus = "Resolved"
	ManagedTenantsManagedTenantAlertStatusunknown    ManagedTenantsManagedTenantAlertStatus = "Unknown"
)

func PossibleValuesForManagedTenantsManagedTenantAlertStatus() []string {
	return []string{
		string(ManagedTenantsManagedTenantAlertStatusdismissed),
		string(ManagedTenantsManagedTenantAlertStatusinProgress),
		string(ManagedTenantsManagedTenantAlertStatusnewAlert),
		string(ManagedTenantsManagedTenantAlertStatusresolved),
		string(ManagedTenantsManagedTenantAlertStatusunknown),
	}
}

func parseManagedTenantsManagedTenantAlertStatus(input string) (*ManagedTenantsManagedTenantAlertStatus, error) {
	vals := map[string]ManagedTenantsManagedTenantAlertStatus{
		"dismissed":  ManagedTenantsManagedTenantAlertStatusdismissed,
		"inprogress": ManagedTenantsManagedTenantAlertStatusinProgress,
		"newalert":   ManagedTenantsManagedTenantAlertStatusnewAlert,
		"resolved":   ManagedTenantsManagedTenantAlertStatusresolved,
		"unknown":    ManagedTenantsManagedTenantAlertStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagedTenantAlertStatus(input)
	return &out, nil
}
