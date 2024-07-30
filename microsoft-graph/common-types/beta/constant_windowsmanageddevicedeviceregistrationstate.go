package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceDeviceRegistrationState string

const (
	WindowsManagedDeviceDeviceRegistrationState_ApprovalPending                WindowsManagedDeviceDeviceRegistrationState = "approvalPending"
	WindowsManagedDeviceDeviceRegistrationState_CertificateReset               WindowsManagedDeviceDeviceRegistrationState = "certificateReset"
	WindowsManagedDeviceDeviceRegistrationState_KeyConflict                    WindowsManagedDeviceDeviceRegistrationState = "keyConflict"
	WindowsManagedDeviceDeviceRegistrationState_NotRegistered                  WindowsManagedDeviceDeviceRegistrationState = "notRegistered"
	WindowsManagedDeviceDeviceRegistrationState_NotRegisteredPendingEnrollment WindowsManagedDeviceDeviceRegistrationState = "notRegisteredPendingEnrollment"
	WindowsManagedDeviceDeviceRegistrationState_Registered                     WindowsManagedDeviceDeviceRegistrationState = "registered"
	WindowsManagedDeviceDeviceRegistrationState_Revoked                        WindowsManagedDeviceDeviceRegistrationState = "revoked"
	WindowsManagedDeviceDeviceRegistrationState_Unknown                        WindowsManagedDeviceDeviceRegistrationState = "unknown"
)

func PossibleValuesForWindowsManagedDeviceDeviceRegistrationState() []string {
	return []string{
		string(WindowsManagedDeviceDeviceRegistrationState_ApprovalPending),
		string(WindowsManagedDeviceDeviceRegistrationState_CertificateReset),
		string(WindowsManagedDeviceDeviceRegistrationState_KeyConflict),
		string(WindowsManagedDeviceDeviceRegistrationState_NotRegistered),
		string(WindowsManagedDeviceDeviceRegistrationState_NotRegisteredPendingEnrollment),
		string(WindowsManagedDeviceDeviceRegistrationState_Registered),
		string(WindowsManagedDeviceDeviceRegistrationState_Revoked),
		string(WindowsManagedDeviceDeviceRegistrationState_Unknown),
	}
}

func (s *WindowsManagedDeviceDeviceRegistrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceDeviceRegistrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceDeviceRegistrationState(input string) (*WindowsManagedDeviceDeviceRegistrationState, error) {
	vals := map[string]WindowsManagedDeviceDeviceRegistrationState{
		"approvalpending":                WindowsManagedDeviceDeviceRegistrationState_ApprovalPending,
		"certificatereset":               WindowsManagedDeviceDeviceRegistrationState_CertificateReset,
		"keyconflict":                    WindowsManagedDeviceDeviceRegistrationState_KeyConflict,
		"notregistered":                  WindowsManagedDeviceDeviceRegistrationState_NotRegistered,
		"notregisteredpendingenrollment": WindowsManagedDeviceDeviceRegistrationState_NotRegisteredPendingEnrollment,
		"registered":                     WindowsManagedDeviceDeviceRegistrationState_Registered,
		"revoked":                        WindowsManagedDeviceDeviceRegistrationState_Revoked,
		"unknown":                        WindowsManagedDeviceDeviceRegistrationState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceDeviceRegistrationState(input)
	return &out, nil
}
