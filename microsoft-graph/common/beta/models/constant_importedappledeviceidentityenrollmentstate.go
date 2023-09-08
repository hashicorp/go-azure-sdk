package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityEnrollmentState string

const (
	ImportedAppleDeviceIdentityEnrollmentStateblocked      ImportedAppleDeviceIdentityEnrollmentState = "Blocked"
	ImportedAppleDeviceIdentityEnrollmentStateenrolled     ImportedAppleDeviceIdentityEnrollmentState = "Enrolled"
	ImportedAppleDeviceIdentityEnrollmentStatefailed       ImportedAppleDeviceIdentityEnrollmentState = "Failed"
	ImportedAppleDeviceIdentityEnrollmentStatenotContacted ImportedAppleDeviceIdentityEnrollmentState = "NotContacted"
	ImportedAppleDeviceIdentityEnrollmentStatependingReset ImportedAppleDeviceIdentityEnrollmentState = "PendingReset"
	ImportedAppleDeviceIdentityEnrollmentStateunknown      ImportedAppleDeviceIdentityEnrollmentState = "Unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityEnrollmentState() []string {
	return []string{
		string(ImportedAppleDeviceIdentityEnrollmentStateblocked),
		string(ImportedAppleDeviceIdentityEnrollmentStateenrolled),
		string(ImportedAppleDeviceIdentityEnrollmentStatefailed),
		string(ImportedAppleDeviceIdentityEnrollmentStatenotContacted),
		string(ImportedAppleDeviceIdentityEnrollmentStatependingReset),
		string(ImportedAppleDeviceIdentityEnrollmentStateunknown),
	}
}

func parseImportedAppleDeviceIdentityEnrollmentState(input string) (*ImportedAppleDeviceIdentityEnrollmentState, error) {
	vals := map[string]ImportedAppleDeviceIdentityEnrollmentState{
		"blocked":      ImportedAppleDeviceIdentityEnrollmentStateblocked,
		"enrolled":     ImportedAppleDeviceIdentityEnrollmentStateenrolled,
		"failed":       ImportedAppleDeviceIdentityEnrollmentStatefailed,
		"notcontacted": ImportedAppleDeviceIdentityEnrollmentStatenotContacted,
		"pendingreset": ImportedAppleDeviceIdentityEnrollmentStatependingReset,
		"unknown":      ImportedAppleDeviceIdentityEnrollmentStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityEnrollmentState(input)
	return &out, nil
}
