package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionPolicySetItemErrorCode string

const (
	ManagedAppProtectionPolicySetItemErrorCode_Deleted      ManagedAppProtectionPolicySetItemErrorCode = "deleted"
	ManagedAppProtectionPolicySetItemErrorCode_NoError      ManagedAppProtectionPolicySetItemErrorCode = "noError"
	ManagedAppProtectionPolicySetItemErrorCode_NotFound     ManagedAppProtectionPolicySetItemErrorCode = "notFound"
	ManagedAppProtectionPolicySetItemErrorCode_Unauthorized ManagedAppProtectionPolicySetItemErrorCode = "unauthorized"
)

func PossibleValuesForManagedAppProtectionPolicySetItemErrorCode() []string {
	return []string{
		string(ManagedAppProtectionPolicySetItemErrorCode_Deleted),
		string(ManagedAppProtectionPolicySetItemErrorCode_NoError),
		string(ManagedAppProtectionPolicySetItemErrorCode_NotFound),
		string(ManagedAppProtectionPolicySetItemErrorCode_Unauthorized),
	}
}

func (s *ManagedAppProtectionPolicySetItemErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionPolicySetItemErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionPolicySetItemErrorCode(input string) (*ManagedAppProtectionPolicySetItemErrorCode, error) {
	vals := map[string]ManagedAppProtectionPolicySetItemErrorCode{
		"deleted":      ManagedAppProtectionPolicySetItemErrorCode_Deleted,
		"noerror":      ManagedAppProtectionPolicySetItemErrorCode_NoError,
		"notfound":     ManagedAppProtectionPolicySetItemErrorCode_NotFound,
		"unauthorized": ManagedAppProtectionPolicySetItemErrorCode_Unauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionPolicySetItemErrorCode(input)
	return &out, nil
}
