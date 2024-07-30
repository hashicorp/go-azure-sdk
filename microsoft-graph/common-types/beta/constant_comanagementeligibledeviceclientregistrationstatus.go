package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceClientRegistrationStatus string

const (
	ComanagementEligibleDeviceClientRegistrationStatus_ApprovalPending                ComanagementEligibleDeviceClientRegistrationStatus = "approvalPending"
	ComanagementEligibleDeviceClientRegistrationStatus_CertificateReset               ComanagementEligibleDeviceClientRegistrationStatus = "certificateReset"
	ComanagementEligibleDeviceClientRegistrationStatus_KeyConflict                    ComanagementEligibleDeviceClientRegistrationStatus = "keyConflict"
	ComanagementEligibleDeviceClientRegistrationStatus_NotRegistered                  ComanagementEligibleDeviceClientRegistrationStatus = "notRegistered"
	ComanagementEligibleDeviceClientRegistrationStatus_NotRegisteredPendingEnrollment ComanagementEligibleDeviceClientRegistrationStatus = "notRegisteredPendingEnrollment"
	ComanagementEligibleDeviceClientRegistrationStatus_Registered                     ComanagementEligibleDeviceClientRegistrationStatus = "registered"
	ComanagementEligibleDeviceClientRegistrationStatus_Revoked                        ComanagementEligibleDeviceClientRegistrationStatus = "revoked"
	ComanagementEligibleDeviceClientRegistrationStatus_Unknown                        ComanagementEligibleDeviceClientRegistrationStatus = "unknown"
)

func PossibleValuesForComanagementEligibleDeviceClientRegistrationStatus() []string {
	return []string{
		string(ComanagementEligibleDeviceClientRegistrationStatus_ApprovalPending),
		string(ComanagementEligibleDeviceClientRegistrationStatus_CertificateReset),
		string(ComanagementEligibleDeviceClientRegistrationStatus_KeyConflict),
		string(ComanagementEligibleDeviceClientRegistrationStatus_NotRegistered),
		string(ComanagementEligibleDeviceClientRegistrationStatus_NotRegisteredPendingEnrollment),
		string(ComanagementEligibleDeviceClientRegistrationStatus_Registered),
		string(ComanagementEligibleDeviceClientRegistrationStatus_Revoked),
		string(ComanagementEligibleDeviceClientRegistrationStatus_Unknown),
	}
}

func (s *ComanagementEligibleDeviceClientRegistrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComanagementEligibleDeviceClientRegistrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComanagementEligibleDeviceClientRegistrationStatus(input string) (*ComanagementEligibleDeviceClientRegistrationStatus, error) {
	vals := map[string]ComanagementEligibleDeviceClientRegistrationStatus{
		"approvalpending":                ComanagementEligibleDeviceClientRegistrationStatus_ApprovalPending,
		"certificatereset":               ComanagementEligibleDeviceClientRegistrationStatus_CertificateReset,
		"keyconflict":                    ComanagementEligibleDeviceClientRegistrationStatus_KeyConflict,
		"notregistered":                  ComanagementEligibleDeviceClientRegistrationStatus_NotRegistered,
		"notregisteredpendingenrollment": ComanagementEligibleDeviceClientRegistrationStatus_NotRegisteredPendingEnrollment,
		"registered":                     ComanagementEligibleDeviceClientRegistrationStatus_Registered,
		"revoked":                        ComanagementEligibleDeviceClientRegistrationStatus_Revoked,
		"unknown":                        ComanagementEligibleDeviceClientRegistrationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceClientRegistrationStatus(input)
	return &out, nil
}
