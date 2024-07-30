package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskProcessingResultProcessingStatus string

const (
	IdentityGovernanceTaskProcessingResultProcessingStatus_Canceled            IdentityGovernanceTaskProcessingResultProcessingStatus = "canceled"
	IdentityGovernanceTaskProcessingResultProcessingStatus_Completed           IdentityGovernanceTaskProcessingResultProcessingStatus = "completed"
	IdentityGovernanceTaskProcessingResultProcessingStatus_CompletedWithErrors IdentityGovernanceTaskProcessingResultProcessingStatus = "completedWithErrors"
	IdentityGovernanceTaskProcessingResultProcessingStatus_Failed              IdentityGovernanceTaskProcessingResultProcessingStatus = "failed"
	IdentityGovernanceTaskProcessingResultProcessingStatus_InProgress          IdentityGovernanceTaskProcessingResultProcessingStatus = "inProgress"
	IdentityGovernanceTaskProcessingResultProcessingStatus_Queued              IdentityGovernanceTaskProcessingResultProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceTaskProcessingResultProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceTaskProcessingResultProcessingStatus_Canceled),
		string(IdentityGovernanceTaskProcessingResultProcessingStatus_Completed),
		string(IdentityGovernanceTaskProcessingResultProcessingStatus_CompletedWithErrors),
		string(IdentityGovernanceTaskProcessingResultProcessingStatus_Failed),
		string(IdentityGovernanceTaskProcessingResultProcessingStatus_InProgress),
		string(IdentityGovernanceTaskProcessingResultProcessingStatus_Queued),
	}
}

func (s *IdentityGovernanceTaskProcessingResultProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTaskProcessingResultProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTaskProcessingResultProcessingStatus(input string) (*IdentityGovernanceTaskProcessingResultProcessingStatus, error) {
	vals := map[string]IdentityGovernanceTaskProcessingResultProcessingStatus{
		"canceled":            IdentityGovernanceTaskProcessingResultProcessingStatus_Canceled,
		"completed":           IdentityGovernanceTaskProcessingResultProcessingStatus_Completed,
		"completedwitherrors": IdentityGovernanceTaskProcessingResultProcessingStatus_CompletedWithErrors,
		"failed":              IdentityGovernanceTaskProcessingResultProcessingStatus_Failed,
		"inprogress":          IdentityGovernanceTaskProcessingResultProcessingStatus_InProgress,
		"queued":              IdentityGovernanceTaskProcessingResultProcessingStatus_Queued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskProcessingResultProcessingStatus(input)
	return &out, nil
}
