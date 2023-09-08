package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskReportProcessingStatus string

const (
	IdentityGovernanceTaskReportProcessingStatuscanceled            IdentityGovernanceTaskReportProcessingStatus = "Canceled"
	IdentityGovernanceTaskReportProcessingStatuscompleted           IdentityGovernanceTaskReportProcessingStatus = "Completed"
	IdentityGovernanceTaskReportProcessingStatuscompletedWithErrors IdentityGovernanceTaskReportProcessingStatus = "CompletedWithErrors"
	IdentityGovernanceTaskReportProcessingStatusfailed              IdentityGovernanceTaskReportProcessingStatus = "Failed"
	IdentityGovernanceTaskReportProcessingStatusinProgress          IdentityGovernanceTaskReportProcessingStatus = "InProgress"
	IdentityGovernanceTaskReportProcessingStatusqueued              IdentityGovernanceTaskReportProcessingStatus = "Queued"
)

func PossibleValuesForIdentityGovernanceTaskReportProcessingStatus() []string {
	return []string{
		string(IdentityGovernanceTaskReportProcessingStatuscanceled),
		string(IdentityGovernanceTaskReportProcessingStatuscompleted),
		string(IdentityGovernanceTaskReportProcessingStatuscompletedWithErrors),
		string(IdentityGovernanceTaskReportProcessingStatusfailed),
		string(IdentityGovernanceTaskReportProcessingStatusinProgress),
		string(IdentityGovernanceTaskReportProcessingStatusqueued),
	}
}

func parseIdentityGovernanceTaskReportProcessingStatus(input string) (*IdentityGovernanceTaskReportProcessingStatus, error) {
	vals := map[string]IdentityGovernanceTaskReportProcessingStatus{
		"canceled":            IdentityGovernanceTaskReportProcessingStatuscanceled,
		"completed":           IdentityGovernanceTaskReportProcessingStatuscompleted,
		"completedwitherrors": IdentityGovernanceTaskReportProcessingStatuscompletedWithErrors,
		"failed":              IdentityGovernanceTaskReportProcessingStatusfailed,
		"inprogress":          IdentityGovernanceTaskReportProcessingStatusinProgress,
		"queued":              IdentityGovernanceTaskReportProcessingStatusqueued,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceTaskReportProcessingStatus(input)
	return &out, nil
}
