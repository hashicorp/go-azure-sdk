package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceClientRegistrationStatus string

const (
	ComanagementEligibleDeviceClientRegistrationStatusapprovalPending                ComanagementEligibleDeviceClientRegistrationStatus = "ApprovalPending"
	ComanagementEligibleDeviceClientRegistrationStatuscertificateReset               ComanagementEligibleDeviceClientRegistrationStatus = "CertificateReset"
	ComanagementEligibleDeviceClientRegistrationStatuskeyConflict                    ComanagementEligibleDeviceClientRegistrationStatus = "KeyConflict"
	ComanagementEligibleDeviceClientRegistrationStatusnotRegistered                  ComanagementEligibleDeviceClientRegistrationStatus = "NotRegistered"
	ComanagementEligibleDeviceClientRegistrationStatusnotRegisteredPendingEnrollment ComanagementEligibleDeviceClientRegistrationStatus = "NotRegisteredPendingEnrollment"
	ComanagementEligibleDeviceClientRegistrationStatusregistered                     ComanagementEligibleDeviceClientRegistrationStatus = "Registered"
	ComanagementEligibleDeviceClientRegistrationStatusrevoked                        ComanagementEligibleDeviceClientRegistrationStatus = "Revoked"
	ComanagementEligibleDeviceClientRegistrationStatusunknown                        ComanagementEligibleDeviceClientRegistrationStatus = "Unknown"
)

func PossibleValuesForComanagementEligibleDeviceClientRegistrationStatus() []string {
	return []string{
		string(ComanagementEligibleDeviceClientRegistrationStatusapprovalPending),
		string(ComanagementEligibleDeviceClientRegistrationStatuscertificateReset),
		string(ComanagementEligibleDeviceClientRegistrationStatuskeyConflict),
		string(ComanagementEligibleDeviceClientRegistrationStatusnotRegistered),
		string(ComanagementEligibleDeviceClientRegistrationStatusnotRegisteredPendingEnrollment),
		string(ComanagementEligibleDeviceClientRegistrationStatusregistered),
		string(ComanagementEligibleDeviceClientRegistrationStatusrevoked),
		string(ComanagementEligibleDeviceClientRegistrationStatusunknown),
	}
}

func parseComanagementEligibleDeviceClientRegistrationStatus(input string) (*ComanagementEligibleDeviceClientRegistrationStatus, error) {
	vals := map[string]ComanagementEligibleDeviceClientRegistrationStatus{
		"approvalpending":                ComanagementEligibleDeviceClientRegistrationStatusapprovalPending,
		"certificatereset":               ComanagementEligibleDeviceClientRegistrationStatuscertificateReset,
		"keyconflict":                    ComanagementEligibleDeviceClientRegistrationStatuskeyConflict,
		"notregistered":                  ComanagementEligibleDeviceClientRegistrationStatusnotRegistered,
		"notregisteredpendingenrollment": ComanagementEligibleDeviceClientRegistrationStatusnotRegisteredPendingEnrollment,
		"registered":                     ComanagementEligibleDeviceClientRegistrationStatusregistered,
		"revoked":                        ComanagementEligibleDeviceClientRegistrationStatusrevoked,
		"unknown":                        ComanagementEligibleDeviceClientRegistrationStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceClientRegistrationStatus(input)
	return &out, nil
}
