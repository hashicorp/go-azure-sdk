package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateIssuanceState string

const (
	ManagedDeviceCertificateStateCertificateIssuanceStatechallengeIssueFailed         ManagedDeviceCertificateStateCertificateIssuanceState = "ChallengeIssueFailed"
	ManagedDeviceCertificateStateCertificateIssuanceStatechallengeIssued              ManagedDeviceCertificateStateCertificateIssuanceState = "ChallengeIssued"
	ManagedDeviceCertificateStateCertificateIssuanceStatechallengeValidationFailed    ManagedDeviceCertificateStateCertificateIssuanceState = "ChallengeValidationFailed"
	ManagedDeviceCertificateStateCertificateIssuanceStatechallengeValidationSucceeded ManagedDeviceCertificateStateCertificateIssuanceState = "ChallengeValidationSucceeded"
	ManagedDeviceCertificateStateCertificateIssuanceStatedeleteFailed                 ManagedDeviceCertificateStateCertificateIssuanceState = "DeleteFailed"
	ManagedDeviceCertificateStateCertificateIssuanceStatedeleted                      ManagedDeviceCertificateStateCertificateIssuanceState = "Deleted"
	ManagedDeviceCertificateStateCertificateIssuanceStateenrollmentNotNeeded          ManagedDeviceCertificateStateCertificateIssuanceState = "EnrollmentNotNeeded"
	ManagedDeviceCertificateStateCertificateIssuanceStateenrollmentSucceeded          ManagedDeviceCertificateStateCertificateIssuanceState = "EnrollmentSucceeded"
	ManagedDeviceCertificateStateCertificateIssuanceStateinstallFailed                ManagedDeviceCertificateStateCertificateIssuanceState = "InstallFailed"
	ManagedDeviceCertificateStateCertificateIssuanceStateinstalled                    ManagedDeviceCertificateStateCertificateIssuanceState = "Installed"
	ManagedDeviceCertificateStateCertificateIssuanceStateissueFailed                  ManagedDeviceCertificateStateCertificateIssuanceState = "IssueFailed"
	ManagedDeviceCertificateStateCertificateIssuanceStateissuePending                 ManagedDeviceCertificateStateCertificateIssuanceState = "IssuePending"
	ManagedDeviceCertificateStateCertificateIssuanceStateissued                       ManagedDeviceCertificateStateCertificateIssuanceState = "Issued"
	ManagedDeviceCertificateStateCertificateIssuanceStateremovedFromCollection        ManagedDeviceCertificateStateCertificateIssuanceState = "RemovedFromCollection"
	ManagedDeviceCertificateStateCertificateIssuanceStaterenewVerified                ManagedDeviceCertificateStateCertificateIssuanceState = "RenewVerified"
	ManagedDeviceCertificateStateCertificateIssuanceStaterenewalRequested             ManagedDeviceCertificateStateCertificateIssuanceState = "RenewalRequested"
	ManagedDeviceCertificateStateCertificateIssuanceStaterequestCreationFailed        ManagedDeviceCertificateStateCertificateIssuanceState = "RequestCreationFailed"
	ManagedDeviceCertificateStateCertificateIssuanceStaterequestSubmitFailed          ManagedDeviceCertificateStateCertificateIssuanceState = "RequestSubmitFailed"
	ManagedDeviceCertificateStateCertificateIssuanceStaterequested                    ManagedDeviceCertificateStateCertificateIssuanceState = "Requested"
	ManagedDeviceCertificateStateCertificateIssuanceStateresponsePending              ManagedDeviceCertificateStateCertificateIssuanceState = "ResponsePending"
	ManagedDeviceCertificateStateCertificateIssuanceStateresponseProcessingFailed     ManagedDeviceCertificateStateCertificateIssuanceState = "ResponseProcessingFailed"
	ManagedDeviceCertificateStateCertificateIssuanceStaterevoked                      ManagedDeviceCertificateStateCertificateIssuanceState = "Revoked"
	ManagedDeviceCertificateStateCertificateIssuanceStateunknown                      ManagedDeviceCertificateStateCertificateIssuanceState = "Unknown"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateIssuanceState() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateIssuanceStatechallengeIssueFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceStatechallengeIssued),
		string(ManagedDeviceCertificateStateCertificateIssuanceStatechallengeValidationFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceStatechallengeValidationSucceeded),
		string(ManagedDeviceCertificateStateCertificateIssuanceStatedeleteFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceStatedeleted),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateenrollmentNotNeeded),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateenrollmentSucceeded),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateinstallFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateinstalled),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateissueFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateissuePending),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateissued),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateremovedFromCollection),
		string(ManagedDeviceCertificateStateCertificateIssuanceStaterenewVerified),
		string(ManagedDeviceCertificateStateCertificateIssuanceStaterenewalRequested),
		string(ManagedDeviceCertificateStateCertificateIssuanceStaterequestCreationFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceStaterequestSubmitFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceStaterequested),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateresponsePending),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateresponseProcessingFailed),
		string(ManagedDeviceCertificateStateCertificateIssuanceStaterevoked),
		string(ManagedDeviceCertificateStateCertificateIssuanceStateunknown),
	}
}

func parseManagedDeviceCertificateStateCertificateIssuanceState(input string) (*ManagedDeviceCertificateStateCertificateIssuanceState, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateIssuanceState{
		"challengeissuefailed":         ManagedDeviceCertificateStateCertificateIssuanceStatechallengeIssueFailed,
		"challengeissued":              ManagedDeviceCertificateStateCertificateIssuanceStatechallengeIssued,
		"challengevalidationfailed":    ManagedDeviceCertificateStateCertificateIssuanceStatechallengeValidationFailed,
		"challengevalidationsucceeded": ManagedDeviceCertificateStateCertificateIssuanceStatechallengeValidationSucceeded,
		"deletefailed":                 ManagedDeviceCertificateStateCertificateIssuanceStatedeleteFailed,
		"deleted":                      ManagedDeviceCertificateStateCertificateIssuanceStatedeleted,
		"enrollmentnotneeded":          ManagedDeviceCertificateStateCertificateIssuanceStateenrollmentNotNeeded,
		"enrollmentsucceeded":          ManagedDeviceCertificateStateCertificateIssuanceStateenrollmentSucceeded,
		"installfailed":                ManagedDeviceCertificateStateCertificateIssuanceStateinstallFailed,
		"installed":                    ManagedDeviceCertificateStateCertificateIssuanceStateinstalled,
		"issuefailed":                  ManagedDeviceCertificateStateCertificateIssuanceStateissueFailed,
		"issuepending":                 ManagedDeviceCertificateStateCertificateIssuanceStateissuePending,
		"issued":                       ManagedDeviceCertificateStateCertificateIssuanceStateissued,
		"removedfromcollection":        ManagedDeviceCertificateStateCertificateIssuanceStateremovedFromCollection,
		"renewverified":                ManagedDeviceCertificateStateCertificateIssuanceStaterenewVerified,
		"renewalrequested":             ManagedDeviceCertificateStateCertificateIssuanceStaterenewalRequested,
		"requestcreationfailed":        ManagedDeviceCertificateStateCertificateIssuanceStaterequestCreationFailed,
		"requestsubmitfailed":          ManagedDeviceCertificateStateCertificateIssuanceStaterequestSubmitFailed,
		"requested":                    ManagedDeviceCertificateStateCertificateIssuanceStaterequested,
		"responsepending":              ManagedDeviceCertificateStateCertificateIssuanceStateresponsePending,
		"responseprocessingfailed":     ManagedDeviceCertificateStateCertificateIssuanceStateresponseProcessingFailed,
		"revoked":                      ManagedDeviceCertificateStateCertificateIssuanceStaterevoked,
		"unknown":                      ManagedDeviceCertificateStateCertificateIssuanceStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateIssuanceState(input)
	return &out, nil
}
