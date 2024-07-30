package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionPolicySetItemStatus string

const (
	ManagedAppProtectionPolicySetItemStatus_Error          ManagedAppProtectionPolicySetItemStatus = "error"
	ManagedAppProtectionPolicySetItemStatus_NotAssigned    ManagedAppProtectionPolicySetItemStatus = "notAssigned"
	ManagedAppProtectionPolicySetItemStatus_PartialSuccess ManagedAppProtectionPolicySetItemStatus = "partialSuccess"
	ManagedAppProtectionPolicySetItemStatus_Success        ManagedAppProtectionPolicySetItemStatus = "success"
	ManagedAppProtectionPolicySetItemStatus_Unknown        ManagedAppProtectionPolicySetItemStatus = "unknown"
	ManagedAppProtectionPolicySetItemStatus_Validating     ManagedAppProtectionPolicySetItemStatus = "validating"
)

func PossibleValuesForManagedAppProtectionPolicySetItemStatus() []string {
	return []string{
		string(ManagedAppProtectionPolicySetItemStatus_Error),
		string(ManagedAppProtectionPolicySetItemStatus_NotAssigned),
		string(ManagedAppProtectionPolicySetItemStatus_PartialSuccess),
		string(ManagedAppProtectionPolicySetItemStatus_Success),
		string(ManagedAppProtectionPolicySetItemStatus_Unknown),
		string(ManagedAppProtectionPolicySetItemStatus_Validating),
	}
}

func (s *ManagedAppProtectionPolicySetItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionPolicySetItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionPolicySetItemStatus(input string) (*ManagedAppProtectionPolicySetItemStatus, error) {
	vals := map[string]ManagedAppProtectionPolicySetItemStatus{
		"error":          ManagedAppProtectionPolicySetItemStatus_Error,
		"notassigned":    ManagedAppProtectionPolicySetItemStatus_NotAssigned,
		"partialsuccess": ManagedAppProtectionPolicySetItemStatus_PartialSuccess,
		"success":        ManagedAppProtectionPolicySetItemStatus_Success,
		"unknown":        ManagedAppProtectionPolicySetItemStatus_Unknown,
		"validating":     ManagedAppProtectionPolicySetItemStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionPolicySetItemStatus(input)
	return &out, nil
}
