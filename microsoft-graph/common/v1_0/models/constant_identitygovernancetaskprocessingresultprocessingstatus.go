package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskProcessingResultProcessingStatus string

const (
	IdentityGovernanceTaskProcessingResultProcessingStatuscanceled            IdentityGovernanceTaskProcessingResultProcessingStatus = "Canceled"
	IdentityGovernanceTaskProcessingResultProcessingStatuscompleted           IdentityGovernanceTaskProcessingResultProcessingStatus = "Completed"
	IdentityGovernanceTaskProcessingResultProcessingStatuscompletedWithErrors IdentityGovernanceTaskProcessingResultProcessingStatus = "CompletedWithErrors"
	IdentityGovernanceTaskProcessingResultProcessingStatusfailed              IdentityGovernanceTaskProcessingResultProcessingStatus = "Failed"
	IdentityGovernanceTaskProcessingResultProcessingStatusinProgress          IdentityGovernanceTaskProcessingResultProcessingStatus = "InProgress"
	IdentityGovernanceTaskProcessingResultProcessingStatusqueued              IdentityGovernanceTaskProcessingResultProcessingStatus = "Queued"
)

func PossibleValuesForIdentityGovernanceTaskProcessingResultProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceTaskProcessingResultProcessingStatuscanceled),
		string(IdentityGovernanceTaskProcessingResultProcessingStatuscompleted),
		string(IdentityGovernanceTaskProcessingResultProcessingStatuscompletedWithErrors),
		string(IdentityGovernanceTaskProcessingResultProcessingStatusfailed),
		string(IdentityGovernanceTaskProcessingResultProcessingStatusinProgress),
		string(IdentityGovernanceTaskProcessingResultProcessingStatusqueued),
	}
}

func parseIdentityGovernanceTaskProcessingResultProcessingStatus(input string) (*IdentityGovernanceTaskProcessingResultProcessingStatus, error) {
	vals := map[string]IdentityGovernanceTaskProcessingResultProcessingStatus{
		"canceled":            IdentityGovernanceTaskProcessingResultProcessingStatuscanceled,
		"completed":           IdentityGovernanceTaskProcessingResultProcessingStatuscompleted,
		"completedwitherrors": IdentityGovernanceTaskProcessingResultProcessingStatuscompletedWithErrors,
		"failed":              IdentityGovernanceTaskProcessingResultProcessingStatusfailed,
		"inprogress":          IdentityGovernanceTaskProcessingResultProcessingStatusinProgress,
		"queued":              IdentityGovernanceTaskProcessingResultProcessingStatusqueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskProcessingResultProcessingStatus(input)
	return &out, nil
}
