package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityEnrollmentState string

const (
	ImportedDeviceIdentityEnrollmentStateblocked      ImportedDeviceIdentityEnrollmentState = "Blocked"
	ImportedDeviceIdentityEnrollmentStateenrolled     ImportedDeviceIdentityEnrollmentState = "Enrolled"
	ImportedDeviceIdentityEnrollmentStatefailed       ImportedDeviceIdentityEnrollmentState = "Failed"
	ImportedDeviceIdentityEnrollmentStatenotContacted ImportedDeviceIdentityEnrollmentState = "NotContacted"
	ImportedDeviceIdentityEnrollmentStatependingReset ImportedDeviceIdentityEnrollmentState = "PendingReset"
	ImportedDeviceIdentityEnrollmentStateunknown      ImportedDeviceIdentityEnrollmentState = "Unknown"
)

func PossibleValuesForImportedDeviceIdentityEnrollmentState() []string {
	return []string{
		string(ImportedDeviceIdentityEnrollmentStateblocked),
		string(ImportedDeviceIdentityEnrollmentStateenrolled),
		string(ImportedDeviceIdentityEnrollmentStatefailed),
		string(ImportedDeviceIdentityEnrollmentStatenotContacted),
		string(ImportedDeviceIdentityEnrollmentStatependingReset),
		string(ImportedDeviceIdentityEnrollmentStateunknown),
	}
}

func parseImportedDeviceIdentityEnrollmentState(input string) (*ImportedDeviceIdentityEnrollmentState, error) {
	vals := map[string]ImportedDeviceIdentityEnrollmentState{
		"blocked":      ImportedDeviceIdentityEnrollmentStateblocked,
		"enrolled":     ImportedDeviceIdentityEnrollmentStateenrolled,
		"failed":       ImportedDeviceIdentityEnrollmentStatefailed,
		"notcontacted": ImportedDeviceIdentityEnrollmentStatenotContacted,
		"pendingreset": ImportedDeviceIdentityEnrollmentStatependingReset,
		"unknown":      ImportedDeviceIdentityEnrollmentStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityEnrollmentState(input)
	return &out, nil
}
