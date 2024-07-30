package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRunProcessingStatus string

const (
	IdentityGovernanceRunProcessingStatus_Canceled            IdentityGovernanceRunProcessingStatus = "canceled"
	IdentityGovernanceRunProcessingStatus_Completed           IdentityGovernanceRunProcessingStatus = "completed"
	IdentityGovernanceRunProcessingStatus_CompletedWithErrors IdentityGovernanceRunProcessingStatus = "completedWithErrors"
	IdentityGovernanceRunProcessingStatus_Failed              IdentityGovernanceRunProcessingStatus = "failed"
	IdentityGovernanceRunProcessingStatus_InProgress          IdentityGovernanceRunProcessingStatus = "inProgress"
	IdentityGovernanceRunProcessingStatus_Queued              IdentityGovernanceRunProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceRunProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceRunProcessingStatus_Canceled),
		string(IdentityGovernanceRunProcessingStatus_Completed),
		string(IdentityGovernanceRunProcessingStatus_CompletedWithErrors),
		string(IdentityGovernanceRunProcessingStatus_Failed),
		string(IdentityGovernanceRunProcessingStatus_InProgress),
		string(IdentityGovernanceRunProcessingStatus_Queued),
	}
}

func (s *IdentityGovernanceRunProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceRunProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceRunProcessingStatus(input string) (*IdentityGovernanceRunProcessingStatus, error) {
	vals := map[string]IdentityGovernanceRunProcessingStatus{
		"canceled":            IdentityGovernanceRunProcessingStatus_Canceled,
		"completed":           IdentityGovernanceRunProcessingStatus_Completed,
		"completedwitherrors": IdentityGovernanceRunProcessingStatus_CompletedWithErrors,
		"failed":              IdentityGovernanceRunProcessingStatus_Failed,
		"inprogress":          IdentityGovernanceRunProcessingStatus_InProgress,
		"queued":              IdentityGovernanceRunProcessingStatus_Queued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceRunProcessingStatus(input)
	return &out, nil
}
