package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyOperationOperationStatus string

const (
	GroupPolicyOperationOperationStatus_Failed     GroupPolicyOperationOperationStatus = "failed"
	GroupPolicyOperationOperationStatus_InProgress GroupPolicyOperationOperationStatus = "inProgress"
	GroupPolicyOperationOperationStatus_Success    GroupPolicyOperationOperationStatus = "success"
	GroupPolicyOperationOperationStatus_Unknown    GroupPolicyOperationOperationStatus = "unknown"
)

func PossibleValuesForGroupPolicyOperationOperationStatus() []string {
	return []string{
		string(GroupPolicyOperationOperationStatus_Failed),
		string(GroupPolicyOperationOperationStatus_InProgress),
		string(GroupPolicyOperationOperationStatus_Success),
		string(GroupPolicyOperationOperationStatus_Unknown),
	}
}

func (s *GroupPolicyOperationOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyOperationOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyOperationOperationStatus(input string) (*GroupPolicyOperationOperationStatus, error) {
	vals := map[string]GroupPolicyOperationOperationStatus{
		"failed":     GroupPolicyOperationOperationStatus_Failed,
		"inprogress": GroupPolicyOperationOperationStatus_InProgress,
		"success":    GroupPolicyOperationOperationStatus_Success,
		"unknown":    GroupPolicyOperationOperationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyOperationOperationStatus(input)
	return &out, nil
}
