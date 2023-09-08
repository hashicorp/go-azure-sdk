package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceDeviceRegistrationState string

const (
	WindowsManagedDeviceDeviceRegistrationStateapprovalPending                WindowsManagedDeviceDeviceRegistrationState = "ApprovalPending"
	WindowsManagedDeviceDeviceRegistrationStatecertificateReset               WindowsManagedDeviceDeviceRegistrationState = "CertificateReset"
	WindowsManagedDeviceDeviceRegistrationStatekeyConflict                    WindowsManagedDeviceDeviceRegistrationState = "KeyConflict"
	WindowsManagedDeviceDeviceRegistrationStatenotRegistered                  WindowsManagedDeviceDeviceRegistrationState = "NotRegistered"
	WindowsManagedDeviceDeviceRegistrationStatenotRegisteredPendingEnrollment WindowsManagedDeviceDeviceRegistrationState = "NotRegisteredPendingEnrollment"
	WindowsManagedDeviceDeviceRegistrationStateregistered                     WindowsManagedDeviceDeviceRegistrationState = "Registered"
	WindowsManagedDeviceDeviceRegistrationStaterevoked                        WindowsManagedDeviceDeviceRegistrationState = "Revoked"
	WindowsManagedDeviceDeviceRegistrationStateunknown                        WindowsManagedDeviceDeviceRegistrationState = "Unknown"
)

func PossibleValuesForWindowsManagedDeviceDeviceRegistrationState() []string {
	return []string{
		string(WindowsManagedDeviceDeviceRegistrationStateapprovalPending),
		string(WindowsManagedDeviceDeviceRegistrationStatecertificateReset),
		string(WindowsManagedDeviceDeviceRegistrationStatekeyConflict),
		string(WindowsManagedDeviceDeviceRegistrationStatenotRegistered),
		string(WindowsManagedDeviceDeviceRegistrationStatenotRegisteredPendingEnrollment),
		string(WindowsManagedDeviceDeviceRegistrationStateregistered),
		string(WindowsManagedDeviceDeviceRegistrationStaterevoked),
		string(WindowsManagedDeviceDeviceRegistrationStateunknown),
	}
}

func parseWindowsManagedDeviceDeviceRegistrationState(input string) (*WindowsManagedDeviceDeviceRegistrationState, error) {
	vals := map[string]WindowsManagedDeviceDeviceRegistrationState{
		"approvalpending":                WindowsManagedDeviceDeviceRegistrationStateapprovalPending,
		"certificatereset":               WindowsManagedDeviceDeviceRegistrationStatecertificateReset,
		"keyconflict":                    WindowsManagedDeviceDeviceRegistrationStatekeyConflict,
		"notregistered":                  WindowsManagedDeviceDeviceRegistrationStatenotRegistered,
		"notregisteredpendingenrollment": WindowsManagedDeviceDeviceRegistrationStatenotRegisteredPendingEnrollment,
		"registered":                     WindowsManagedDeviceDeviceRegistrationStateregistered,
		"revoked":                        WindowsManagedDeviceDeviceRegistrationStaterevoked,
		"unknown":                        WindowsManagedDeviceDeviceRegistrationStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceDeviceRegistrationState(input)
	return &out, nil
}
