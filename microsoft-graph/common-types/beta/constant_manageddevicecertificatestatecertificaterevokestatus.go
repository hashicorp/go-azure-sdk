package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateRevokeStatus string

const (
	ManagedDeviceCertificateStateCertificateRevokeStatus_Failed  ManagedDeviceCertificateStateCertificateRevokeStatus = "failed"
	ManagedDeviceCertificateStateCertificateRevokeStatus_Issued  ManagedDeviceCertificateStateCertificateRevokeStatus = "issued"
	ManagedDeviceCertificateStateCertificateRevokeStatus_None    ManagedDeviceCertificateStateCertificateRevokeStatus = "none"
	ManagedDeviceCertificateStateCertificateRevokeStatus_Pending ManagedDeviceCertificateStateCertificateRevokeStatus = "pending"
	ManagedDeviceCertificateStateCertificateRevokeStatus_Revoked ManagedDeviceCertificateStateCertificateRevokeStatus = "revoked"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateRevokeStatus() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateRevokeStatus_Failed),
		string(ManagedDeviceCertificateStateCertificateRevokeStatus_Issued),
		string(ManagedDeviceCertificateStateCertificateRevokeStatus_None),
		string(ManagedDeviceCertificateStateCertificateRevokeStatus_Pending),
		string(ManagedDeviceCertificateStateCertificateRevokeStatus_Revoked),
	}
}

func (s *ManagedDeviceCertificateStateCertificateRevokeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceCertificateStateCertificateRevokeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceCertificateStateCertificateRevokeStatus(input string) (*ManagedDeviceCertificateStateCertificateRevokeStatus, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateRevokeStatus{
		"failed":  ManagedDeviceCertificateStateCertificateRevokeStatus_Failed,
		"issued":  ManagedDeviceCertificateStateCertificateRevokeStatus_Issued,
		"none":    ManagedDeviceCertificateStateCertificateRevokeStatus_None,
		"pending": ManagedDeviceCertificateStateCertificateRevokeStatus_Pending,
		"revoked": ManagedDeviceCertificateStateCertificateRevokeStatus_Revoked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateRevokeStatus(input)
	return &out, nil
}
