package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateIssuanceState string

const (
	ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeIssueFailed         ManagedDeviceCertificateStateCertificateIssuanceState = "challengeIssueFailed"
	ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeIssued              ManagedDeviceCertificateStateCertificateIssuanceState = "challengeIssued"
	ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeValidationFailed    ManagedDeviceCertificateStateCertificateIssuanceState = "challengeValidationFailed"
	ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeValidationSucceeded ManagedDeviceCertificateStateCertificateIssuanceState = "challengeValidationSucceeded"
	ManagedDeviceCertificateStateCertificateIssuanceState_DeleteFailed                 ManagedDeviceCertificateStateCertificateIssuanceState = "deleteFailed"
	ManagedDeviceCertificateStateCertificateIssuanceState_Deleted                      ManagedDeviceCertificateStateCertificateIssuanceState = "deleted"
	ManagedDeviceCertificateStateCertificateIssuanceState_EnrollmentNotNeeded          ManagedDeviceCertificateStateCertificateIssuanceState = "enrollmentNotNeeded"
	ManagedDeviceCertificateStateCertificateIssuanceState_EnrollmentSucceeded          ManagedDeviceCertificateStateCertificateIssuanceState = "enrollmentSucceeded"
	ManagedDeviceCertificateStateCertificateIssuanceState_InstallFailed                ManagedDeviceCertificateStateCertificateIssuanceState = "installFailed"
	ManagedDeviceCertificateStateCertificateIssuanceState_Installed                    ManagedDeviceCertificateStateCertificateIssuanceState = "installed"
	ManagedDeviceCertificateStateCertificateIssuanceState_IssueFailed                  ManagedDeviceCertificateStateCertificateIssuanceState = "issueFailed"
	ManagedDeviceCertificateStateCertificateIssuanceState_IssuePending                 ManagedDeviceCertificateStateCertificateIssuanceState = "issuePending"
	ManagedDeviceCertificateStateCertificateIssuanceState_Issued                       ManagedDeviceCertificateStateCertificateIssuanceState = "issued"
	ManagedDeviceCertificateStateCertificateIssuanceState_RemovedFromCollection        ManagedDeviceCertificateStateCertificateIssuanceState = "removedFromCollection"
	ManagedDeviceCertificateStateCertificateIssuanceState_RenewVerified                ManagedDeviceCertificateStateCertificateIssuanceState = "renewVerified"
	ManagedDeviceCertificateStateCertificateIssuanceState_RenewalRequested             ManagedDeviceCertificateStateCertificateIssuanceState = "renewalRequested"
	ManagedDeviceCertificateStateCertificateIssuanceState_RequestCreationFailed        ManagedDeviceCertificateStateCertificateIssuanceState = "requestCreationFailed"
	ManagedDeviceCertificateStateCertificateIssuanceState_RequestSubmitFailed          ManagedDeviceCertificateStateCertificateIssuanceState = "requestSubmitFailed"
	ManagedDeviceCertificateStateCertificateIssuanceState_Requested                    ManagedDeviceCertificateStateCertificateIssuanceState = "requested"
	ManagedDeviceCertificateStateCertificateIssuanceState_ResponsePending              ManagedDeviceCertificateStateCertificateIssuanceState = "responsePending"
	ManagedDeviceCertificateStateCertificateIssuanceState_ResponseProcessingFailed     ManagedDeviceCertificateStateCertificateIssuanceState = "responseProcessingFailed"
	ManagedDeviceCertificateStateCertificateIssuanceState_Revoked                      ManagedDeviceCertificateStateCertificateIssuanceState = "revoked"
	ManagedDeviceCertificateStateCertificateIssuanceState_Unknown                      ManagedDeviceCertificateStateCertificateIssuanceState = "unknown"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateIssuanceState() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeIssueFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeIssued),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeValidationFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeValidationSucceeded),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_DeleteFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_Deleted),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_EnrollmentNotNeeded),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_EnrollmentSucceeded),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_InstallFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_Installed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_IssueFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_IssuePending),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_Issued),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_RemovedFromCollection),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_RenewVerified),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_RenewalRequested),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_RequestCreationFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_RequestSubmitFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_Requested),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_ResponsePending),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_ResponseProcessingFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_Revoked),
		string(ManagedDeviceCertificateStateCertificateIssuanceState_Unknown),
	}
}

func (s *ManagedDeviceCertificateStateCertificateIssuanceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceCertificateStateCertificateIssuanceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceCertificateStateCertificateIssuanceState(input string) (*ManagedDeviceCertificateStateCertificateIssuanceState, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateIssuanceState{
		"challengeissuefailed":         ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeIssueFailed,
		"challengeissued":              ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeIssued,
		"challengevalidationfailed":    ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeValidationFailed,
		"challengevalidationsucceeded": ManagedDeviceCertificateStateCertificateIssuanceState_ChallengeValidationSucceeded,
		"deletefailed":                 ManagedDeviceCertificateStateCertificateIssuanceState_DeleteFailed,
		"deleted":                      ManagedDeviceCertificateStateCertificateIssuanceState_Deleted,
		"enrollmentnotneeded":          ManagedDeviceCertificateStateCertificateIssuanceState_EnrollmentNotNeeded,
		"enrollmentsucceeded":          ManagedDeviceCertificateStateCertificateIssuanceState_EnrollmentSucceeded,
		"installfailed":                ManagedDeviceCertificateStateCertificateIssuanceState_InstallFailed,
		"installed":                    ManagedDeviceCertificateStateCertificateIssuanceState_Installed,
		"issuefailed":                  ManagedDeviceCertificateStateCertificateIssuanceState_IssueFailed,
		"issuepending":                 ManagedDeviceCertificateStateCertificateIssuanceState_IssuePending,
		"issued":                       ManagedDeviceCertificateStateCertificateIssuanceState_Issued,
		"removedfromcollection":        ManagedDeviceCertificateStateCertificateIssuanceState_RemovedFromCollection,
		"renewverified":                ManagedDeviceCertificateStateCertificateIssuanceState_RenewVerified,
		"renewalrequested":             ManagedDeviceCertificateStateCertificateIssuanceState_RenewalRequested,
		"requestcreationfailed":        ManagedDeviceCertificateStateCertificateIssuanceState_RequestCreationFailed,
		"requestsubmitfailed":          ManagedDeviceCertificateStateCertificateIssuanceState_RequestSubmitFailed,
		"requested":                    ManagedDeviceCertificateStateCertificateIssuanceState_Requested,
		"responsepending":              ManagedDeviceCertificateStateCertificateIssuanceState_ResponsePending,
		"responseprocessingfailed":     ManagedDeviceCertificateStateCertificateIssuanceState_ResponseProcessingFailed,
		"revoked":                      ManagedDeviceCertificateStateCertificateIssuanceState_Revoked,
		"unknown":                      ManagedDeviceCertificateStateCertificateIssuanceState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateIssuanceState(input)
	return &out, nil
}
