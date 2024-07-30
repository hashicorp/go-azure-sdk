package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceRegistrationState string

const (
	ManagedDeviceDeviceRegistrationState_ApprovalPending                ManagedDeviceDeviceRegistrationState = "approvalPending"
	ManagedDeviceDeviceRegistrationState_CertificateReset               ManagedDeviceDeviceRegistrationState = "certificateReset"
	ManagedDeviceDeviceRegistrationState_KeyConflict                    ManagedDeviceDeviceRegistrationState = "keyConflict"
	ManagedDeviceDeviceRegistrationState_NotRegistered                  ManagedDeviceDeviceRegistrationState = "notRegistered"
	ManagedDeviceDeviceRegistrationState_NotRegisteredPendingEnrollment ManagedDeviceDeviceRegistrationState = "notRegisteredPendingEnrollment"
	ManagedDeviceDeviceRegistrationState_Registered                     ManagedDeviceDeviceRegistrationState = "registered"
	ManagedDeviceDeviceRegistrationState_Revoked                        ManagedDeviceDeviceRegistrationState = "revoked"
	ManagedDeviceDeviceRegistrationState_Unknown                        ManagedDeviceDeviceRegistrationState = "unknown"
)

func PossibleValuesForManagedDeviceDeviceRegistrationState() []string {
	return []string{
		string(ManagedDeviceDeviceRegistrationState_ApprovalPending),
		string(ManagedDeviceDeviceRegistrationState_CertificateReset),
		string(ManagedDeviceDeviceRegistrationState_KeyConflict),
		string(ManagedDeviceDeviceRegistrationState_NotRegistered),
		string(ManagedDeviceDeviceRegistrationState_NotRegisteredPendingEnrollment),
		string(ManagedDeviceDeviceRegistrationState_Registered),
		string(ManagedDeviceDeviceRegistrationState_Revoked),
		string(ManagedDeviceDeviceRegistrationState_Unknown),
	}
}

func (s *ManagedDeviceDeviceRegistrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceDeviceRegistrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceDeviceRegistrationState(input string) (*ManagedDeviceDeviceRegistrationState, error) {
	vals := map[string]ManagedDeviceDeviceRegistrationState{
		"approvalpending":                ManagedDeviceDeviceRegistrationState_ApprovalPending,
		"certificatereset":               ManagedDeviceDeviceRegistrationState_CertificateReset,
		"keyconflict":                    ManagedDeviceDeviceRegistrationState_KeyConflict,
		"notregistered":                  ManagedDeviceDeviceRegistrationState_NotRegistered,
		"notregisteredpendingenrollment": ManagedDeviceDeviceRegistrationState_NotRegisteredPendingEnrollment,
		"registered":                     ManagedDeviceDeviceRegistrationState_Registered,
		"revoked":                        ManagedDeviceDeviceRegistrationState_Revoked,
		"unknown":                        ManagedDeviceDeviceRegistrationState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceRegistrationState(input)
	return &out, nil
}
