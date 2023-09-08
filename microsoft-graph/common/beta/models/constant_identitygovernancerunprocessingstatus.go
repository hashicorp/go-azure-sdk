package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRunProcessingStatus string

const (
	IdentityGovernanceRunProcessingStatuscanceled            IdentityGovernanceRunProcessingStatus = "Canceled"
	IdentityGovernanceRunProcessingStatuscompleted           IdentityGovernanceRunProcessingStatus = "Completed"
	IdentityGovernanceRunProcessingStatuscompletedWithErrors IdentityGovernanceRunProcessingStatus = "CompletedWithErrors"
	IdentityGovernanceRunProcessingStatusfailed              IdentityGovernanceRunProcessingStatus = "Failed"
	IdentityGovernanceRunProcessingStatusinProgress          IdentityGovernanceRunProcessingStatus = "InProgress"
	IdentityGovernanceRunProcessingStatusqueued              IdentityGovernanceRunProcessingStatus = "Queued"
)

func PossibleValuesForIdentityGovernanceRunProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceRunProcessingStatuscanceled),
		string(IdentityGovernanceRunProcessingStatuscompleted),
		string(IdentityGovernanceRunProcessingStatuscompletedWithErrors),
		string(IdentityGovernanceRunProcessingStatusfailed),
		string(IdentityGovernanceRunProcessingStatusinProgress),
		string(IdentityGovernanceRunProcessingStatusqueued),
	}
}

func parseIdentityGovernanceRunProcessingStatus(input string) (*IdentityGovernanceRunProcessingStatus, error) {
	vals := map[string]IdentityGovernanceRunProcessingStatus{
		"canceled":            IdentityGovernanceRunProcessingStatuscanceled,
		"completed":           IdentityGovernanceRunProcessingStatuscompleted,
		"completedwitherrors": IdentityGovernanceRunProcessingStatuscompletedWithErrors,
		"failed":              IdentityGovernanceRunProcessingStatusfailed,
		"inprogress":          IdentityGovernanceRunProcessingStatusinProgress,
		"queued":              IdentityGovernanceRunProcessingStatusqueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceRunProcessingStatus(input)
	return &out, nil
}
