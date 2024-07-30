package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceUserProcessingResultProcessingStatus string

const (
	IdentityGovernanceUserProcessingResultProcessingStatus_Canceled            IdentityGovernanceUserProcessingResultProcessingStatus = "canceled"
	IdentityGovernanceUserProcessingResultProcessingStatus_Completed           IdentityGovernanceUserProcessingResultProcessingStatus = "completed"
	IdentityGovernanceUserProcessingResultProcessingStatus_CompletedWithErrors IdentityGovernanceUserProcessingResultProcessingStatus = "completedWithErrors"
	IdentityGovernanceUserProcessingResultProcessingStatus_Failed              IdentityGovernanceUserProcessingResultProcessingStatus = "failed"
	IdentityGovernanceUserProcessingResultProcessingStatus_InProgress          IdentityGovernanceUserProcessingResultProcessingStatus = "inProgress"
	IdentityGovernanceUserProcessingResultProcessingStatus_Queued              IdentityGovernanceUserProcessingResultProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceUserProcessingResultProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceUserProcessingResultProcessingStatus_Canceled),
		string(IdentityGovernanceUserProcessingResultProcessingStatus_Completed),
		string(IdentityGovernanceUserProcessingResultProcessingStatus_CompletedWithErrors),
		string(IdentityGovernanceUserProcessingResultProcessingStatus_Failed),
		string(IdentityGovernanceUserProcessingResultProcessingStatus_InProgress),
		string(IdentityGovernanceUserProcessingResultProcessingStatus_Queued),
	}
}

func (s *IdentityGovernanceUserProcessingResultProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceUserProcessingResultProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceUserProcessingResultProcessingStatus(input string) (*IdentityGovernanceUserProcessingResultProcessingStatus, error) {
	vals := map[string]IdentityGovernanceUserProcessingResultProcessingStatus{
		"canceled":            IdentityGovernanceUserProcessingResultProcessingStatus_Canceled,
		"completed":           IdentityGovernanceUserProcessingResultProcessingStatus_Completed,
		"completedwitherrors": IdentityGovernanceUserProcessingResultProcessingStatus_CompletedWithErrors,
		"failed":              IdentityGovernanceUserProcessingResultProcessingStatus_Failed,
		"inprogress":          IdentityGovernanceUserProcessingResultProcessingStatus_InProgress,
		"queued":              IdentityGovernanceUserProcessingResultProcessingStatus_Queued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceUserProcessingResultProcessingStatus(input)
	return &out, nil
}
