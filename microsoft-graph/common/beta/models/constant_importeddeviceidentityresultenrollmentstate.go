package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityResultEnrollmentState string

const (
	ImportedDeviceIdentityResultEnrollmentStateblocked      ImportedDeviceIdentityResultEnrollmentState = "Blocked"
	ImportedDeviceIdentityResultEnrollmentStateenrolled     ImportedDeviceIdentityResultEnrollmentState = "Enrolled"
	ImportedDeviceIdentityResultEnrollmentStatefailed       ImportedDeviceIdentityResultEnrollmentState = "Failed"
	ImportedDeviceIdentityResultEnrollmentStatenotContacted ImportedDeviceIdentityResultEnrollmentState = "NotContacted"
	ImportedDeviceIdentityResultEnrollmentStatependingReset ImportedDeviceIdentityResultEnrollmentState = "PendingReset"
	ImportedDeviceIdentityResultEnrollmentStateunknown      ImportedDeviceIdentityResultEnrollmentState = "Unknown"
)

func PossibleValuesForImportedDeviceIdentityResultEnrollmentState() []string {
	return []string{
		string(ImportedDeviceIdentityResultEnrollmentStateblocked),
		string(ImportedDeviceIdentityResultEnrollmentStateenrolled),
		string(ImportedDeviceIdentityResultEnrollmentStatefailed),
		string(ImportedDeviceIdentityResultEnrollmentStatenotContacted),
		string(ImportedDeviceIdentityResultEnrollmentStatependingReset),
		string(ImportedDeviceIdentityResultEnrollmentStateunknown),
	}
}

func parseImportedDeviceIdentityResultEnrollmentState(input string) (*ImportedDeviceIdentityResultEnrollmentState, error) {
	vals := map[string]ImportedDeviceIdentityResultEnrollmentState{
		"blocked":      ImportedDeviceIdentityResultEnrollmentStateblocked,
		"enrolled":     ImportedDeviceIdentityResultEnrollmentStateenrolled,
		"failed":       ImportedDeviceIdentityResultEnrollmentStatefailed,
		"notcontacted": ImportedDeviceIdentityResultEnrollmentStatenotContacted,
		"pendingreset": ImportedDeviceIdentityResultEnrollmentStatependingReset,
		"unknown":      ImportedDeviceIdentityResultEnrollmentStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityResultEnrollmentState(input)
	return &out, nil
}
