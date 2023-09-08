package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityResultEnrollmentState string

const (
	ImportedAppleDeviceIdentityResultEnrollmentStateblocked      ImportedAppleDeviceIdentityResultEnrollmentState = "Blocked"
	ImportedAppleDeviceIdentityResultEnrollmentStateenrolled     ImportedAppleDeviceIdentityResultEnrollmentState = "Enrolled"
	ImportedAppleDeviceIdentityResultEnrollmentStatefailed       ImportedAppleDeviceIdentityResultEnrollmentState = "Failed"
	ImportedAppleDeviceIdentityResultEnrollmentStatenotContacted ImportedAppleDeviceIdentityResultEnrollmentState = "NotContacted"
	ImportedAppleDeviceIdentityResultEnrollmentStatependingReset ImportedAppleDeviceIdentityResultEnrollmentState = "PendingReset"
	ImportedAppleDeviceIdentityResultEnrollmentStateunknown      ImportedAppleDeviceIdentityResultEnrollmentState = "Unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityResultEnrollmentState() []string {
	return []string{
		string(ImportedAppleDeviceIdentityResultEnrollmentStateblocked),
		string(ImportedAppleDeviceIdentityResultEnrollmentStateenrolled),
		string(ImportedAppleDeviceIdentityResultEnrollmentStatefailed),
		string(ImportedAppleDeviceIdentityResultEnrollmentStatenotContacted),
		string(ImportedAppleDeviceIdentityResultEnrollmentStatependingReset),
		string(ImportedAppleDeviceIdentityResultEnrollmentStateunknown),
	}
}

func parseImportedAppleDeviceIdentityResultEnrollmentState(input string) (*ImportedAppleDeviceIdentityResultEnrollmentState, error) {
	vals := map[string]ImportedAppleDeviceIdentityResultEnrollmentState{
		"blocked":      ImportedAppleDeviceIdentityResultEnrollmentStateblocked,
		"enrolled":     ImportedAppleDeviceIdentityResultEnrollmentStateenrolled,
		"failed":       ImportedAppleDeviceIdentityResultEnrollmentStatefailed,
		"notcontacted": ImportedAppleDeviceIdentityResultEnrollmentStatenotContacted,
		"pendingreset": ImportedAppleDeviceIdentityResultEnrollmentStatependingReset,
		"unknown":      ImportedAppleDeviceIdentityResultEnrollmentStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityResultEnrollmentState(input)
	return &out, nil
}
