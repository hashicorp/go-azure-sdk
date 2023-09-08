package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceRegistrationState string

const (
	ManagedDeviceDeviceRegistrationStateapprovalPending                ManagedDeviceDeviceRegistrationState = "ApprovalPending"
	ManagedDeviceDeviceRegistrationStatecertificateReset               ManagedDeviceDeviceRegistrationState = "CertificateReset"
	ManagedDeviceDeviceRegistrationStatekeyConflict                    ManagedDeviceDeviceRegistrationState = "KeyConflict"
	ManagedDeviceDeviceRegistrationStatenotRegistered                  ManagedDeviceDeviceRegistrationState = "NotRegistered"
	ManagedDeviceDeviceRegistrationStatenotRegisteredPendingEnrollment ManagedDeviceDeviceRegistrationState = "NotRegisteredPendingEnrollment"
	ManagedDeviceDeviceRegistrationStateregistered                     ManagedDeviceDeviceRegistrationState = "Registered"
	ManagedDeviceDeviceRegistrationStaterevoked                        ManagedDeviceDeviceRegistrationState = "Revoked"
	ManagedDeviceDeviceRegistrationStateunknown                        ManagedDeviceDeviceRegistrationState = "Unknown"
)

func PossibleValuesForManagedDeviceDeviceRegistrationState() []string {
	return []string{
		string(ManagedDeviceDeviceRegistrationStateapprovalPending),
		string(ManagedDeviceDeviceRegistrationStatecertificateReset),
		string(ManagedDeviceDeviceRegistrationStatekeyConflict),
		string(ManagedDeviceDeviceRegistrationStatenotRegistered),
		string(ManagedDeviceDeviceRegistrationStatenotRegisteredPendingEnrollment),
		string(ManagedDeviceDeviceRegistrationStateregistered),
		string(ManagedDeviceDeviceRegistrationStaterevoked),
		string(ManagedDeviceDeviceRegistrationStateunknown),
	}
}

func parseManagedDeviceDeviceRegistrationState(input string) (*ManagedDeviceDeviceRegistrationState, error) {
	vals := map[string]ManagedDeviceDeviceRegistrationState{
		"approvalpending":                ManagedDeviceDeviceRegistrationStateapprovalPending,
		"certificatereset":               ManagedDeviceDeviceRegistrationStatecertificateReset,
		"keyconflict":                    ManagedDeviceDeviceRegistrationStatekeyConflict,
		"notregistered":                  ManagedDeviceDeviceRegistrationStatenotRegistered,
		"notregisteredpendingenrollment": ManagedDeviceDeviceRegistrationStatenotRegisteredPendingEnrollment,
		"registered":                     ManagedDeviceDeviceRegistrationStateregistered,
		"revoked":                        ManagedDeviceDeviceRegistrationStaterevoked,
		"unknown":                        ManagedDeviceDeviceRegistrationStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceRegistrationState(input)
	return &out, nil
}
