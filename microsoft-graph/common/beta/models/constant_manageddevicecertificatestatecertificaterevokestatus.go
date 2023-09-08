package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateRevokeStatus string

const (
	ManagedDeviceCertificateStateCertificateRevokeStatusfailed  ManagedDeviceCertificateStateCertificateRevokeStatus = "Failed"
	ManagedDeviceCertificateStateCertificateRevokeStatusissued  ManagedDeviceCertificateStateCertificateRevokeStatus = "Issued"
	ManagedDeviceCertificateStateCertificateRevokeStatusnone    ManagedDeviceCertificateStateCertificateRevokeStatus = "None"
	ManagedDeviceCertificateStateCertificateRevokeStatuspending ManagedDeviceCertificateStateCertificateRevokeStatus = "Pending"
	ManagedDeviceCertificateStateCertificateRevokeStatusrevoked ManagedDeviceCertificateStateCertificateRevokeStatus = "Revoked"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateRevokeStatus() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateRevokeStatusfailed),
		string(ManagedDeviceCertificateStateCertificateRevokeStatusissued),
		string(ManagedDeviceCertificateStateCertificateRevokeStatusnone),
		string(ManagedDeviceCertificateStateCertificateRevokeStatuspending),
		string(ManagedDeviceCertificateStateCertificateRevokeStatusrevoked),
	}
}

func parseManagedDeviceCertificateStateCertificateRevokeStatus(input string) (*ManagedDeviceCertificateStateCertificateRevokeStatus, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateRevokeStatus{
		"failed":  ManagedDeviceCertificateStateCertificateRevokeStatusfailed,
		"issued":  ManagedDeviceCertificateStateCertificateRevokeStatusissued,
		"none":    ManagedDeviceCertificateStateCertificateRevokeStatusnone,
		"pending": ManagedDeviceCertificateStateCertificateRevokeStatuspending,
		"revoked": ManagedDeviceCertificateStateCertificateRevokeStatusrevoked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateRevokeStatus(input)
	return &out, nil
}
