package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAllDeviceCertificateStateCertificateRevokeStatus string

const (
	ManagedAllDeviceCertificateStateCertificateRevokeStatus_Failed  ManagedAllDeviceCertificateStateCertificateRevokeStatus = "failed"
	ManagedAllDeviceCertificateStateCertificateRevokeStatus_Issued  ManagedAllDeviceCertificateStateCertificateRevokeStatus = "issued"
	ManagedAllDeviceCertificateStateCertificateRevokeStatus_None    ManagedAllDeviceCertificateStateCertificateRevokeStatus = "none"
	ManagedAllDeviceCertificateStateCertificateRevokeStatus_Pending ManagedAllDeviceCertificateStateCertificateRevokeStatus = "pending"
	ManagedAllDeviceCertificateStateCertificateRevokeStatus_Revoked ManagedAllDeviceCertificateStateCertificateRevokeStatus = "revoked"
)

func PossibleValuesForManagedAllDeviceCertificateStateCertificateRevokeStatus() []string {
	return []string{
		string(ManagedAllDeviceCertificateStateCertificateRevokeStatus_Failed),
		string(ManagedAllDeviceCertificateStateCertificateRevokeStatus_Issued),
		string(ManagedAllDeviceCertificateStateCertificateRevokeStatus_None),
		string(ManagedAllDeviceCertificateStateCertificateRevokeStatus_Pending),
		string(ManagedAllDeviceCertificateStateCertificateRevokeStatus_Revoked),
	}
}

func (s *ManagedAllDeviceCertificateStateCertificateRevokeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAllDeviceCertificateStateCertificateRevokeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAllDeviceCertificateStateCertificateRevokeStatus(input string) (*ManagedAllDeviceCertificateStateCertificateRevokeStatus, error) {
	vals := map[string]ManagedAllDeviceCertificateStateCertificateRevokeStatus{
		"failed":  ManagedAllDeviceCertificateStateCertificateRevokeStatus_Failed,
		"issued":  ManagedAllDeviceCertificateStateCertificateRevokeStatus_Issued,
		"none":    ManagedAllDeviceCertificateStateCertificateRevokeStatus_None,
		"pending": ManagedAllDeviceCertificateStateCertificateRevokeStatus_Pending,
		"revoked": ManagedAllDeviceCertificateStateCertificateRevokeStatus_Revoked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAllDeviceCertificateStateCertificateRevokeStatus(input)
	return &out, nil
}
