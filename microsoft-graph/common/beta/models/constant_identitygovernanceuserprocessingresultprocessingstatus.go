package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceUserProcessingResultProcessingStatus string

const (
	IdentityGovernanceUserProcessingResultProcessingStatuscanceled            IdentityGovernanceUserProcessingResultProcessingStatus = "Canceled"
	IdentityGovernanceUserProcessingResultProcessingStatuscompleted           IdentityGovernanceUserProcessingResultProcessingStatus = "Completed"
	IdentityGovernanceUserProcessingResultProcessingStatuscompletedWithErrors IdentityGovernanceUserProcessingResultProcessingStatus = "CompletedWithErrors"
	IdentityGovernanceUserProcessingResultProcessingStatusfailed              IdentityGovernanceUserProcessingResultProcessingStatus = "Failed"
	IdentityGovernanceUserProcessingResultProcessingStatusinProgress          IdentityGovernanceUserProcessingResultProcessingStatus = "InProgress"
	IdentityGovernanceUserProcessingResultProcessingStatusqueued              IdentityGovernanceUserProcessingResultProcessingStatus = "Queued"
)

func PossibleValuesForIdentityGovernanceUserProcessingResultProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceUserProcessingResultProcessingStatuscanceled),
		string(IdentityGovernanceUserProcessingResultProcessingStatuscompleted),
		string(IdentityGovernanceUserProcessingResultProcessingStatuscompletedWithErrors),
		string(IdentityGovernanceUserProcessingResultProcessingStatusfailed),
		string(IdentityGovernanceUserProcessingResultProcessingStatusinProgress),
		string(IdentityGovernanceUserProcessingResultProcessingStatusqueued),
	}
}

func parseIdentityGovernanceUserProcessingResultProcessingStatus(input string) (*IdentityGovernanceUserProcessingResultProcessingStatus, error) {
	vals := map[string]IdentityGovernanceUserProcessingResultProcessingStatus{
		"canceled":            IdentityGovernanceUserProcessingResultProcessingStatuscanceled,
		"completed":           IdentityGovernanceUserProcessingResultProcessingStatuscompleted,
		"completedwitherrors": IdentityGovernanceUserProcessingResultProcessingStatuscompletedWithErrors,
		"failed":              IdentityGovernanceUserProcessingResultProcessingStatusfailed,
		"inprogress":          IdentityGovernanceUserProcessingResultProcessingStatusinProgress,
		"queued":              IdentityGovernanceUserProcessingResultProcessingStatusqueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceUserProcessingResultProcessingStatus(input)
	return &out, nil
}
