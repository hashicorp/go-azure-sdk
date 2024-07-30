package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskReportProcessingStatus string

const (
	IdentityGovernanceTaskReportProcessingStatus_Canceled            IdentityGovernanceTaskReportProcessingStatus = "canceled"
	IdentityGovernanceTaskReportProcessingStatus_Completed           IdentityGovernanceTaskReportProcessingStatus = "completed"
	IdentityGovernanceTaskReportProcessingStatus_CompletedWithErrors IdentityGovernanceTaskReportProcessingStatus = "completedWithErrors"
	IdentityGovernanceTaskReportProcessingStatus_Failed              IdentityGovernanceTaskReportProcessingStatus = "failed"
	IdentityGovernanceTaskReportProcessingStatus_InProgress          IdentityGovernanceTaskReportProcessingStatus = "inProgress"
	IdentityGovernanceTaskReportProcessingStatus_Queued              IdentityGovernanceTaskReportProcessingStatus = "queued"
)

func PossibleValuesForIdentityGovernanceTaskReportProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceTaskReportProcessingStatus_Canceled),
		string(IdentityGovernanceTaskReportProcessingStatus_Completed),
		string(IdentityGovernanceTaskReportProcessingStatus_CompletedWithErrors),
		string(IdentityGovernanceTaskReportProcessingStatus_Failed),
		string(IdentityGovernanceTaskReportProcessingStatus_InProgress),
		string(IdentityGovernanceTaskReportProcessingStatus_Queued),
	}
}

func (s *IdentityGovernanceTaskReportProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceTaskReportProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceTaskReportProcessingStatus(input string) (*IdentityGovernanceTaskReportProcessingStatus, error) {
	vals := map[string]IdentityGovernanceTaskReportProcessingStatus{
		"canceled":            IdentityGovernanceTaskReportProcessingStatus_Canceled,
		"completed":           IdentityGovernanceTaskReportProcessingStatus_Completed,
		"completedwitherrors": IdentityGovernanceTaskReportProcessingStatus_CompletedWithErrors,
		"failed":              IdentityGovernanceTaskReportProcessingStatus_Failed,
		"inprogress":          IdentityGovernanceTaskReportProcessingStatus_InProgress,
		"queued":              IdentityGovernanceTaskReportProcessingStatus_Queued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskReportProcessingStatus(input)
	return &out, nil
}
