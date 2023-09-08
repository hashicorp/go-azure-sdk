package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestHistoryStageStatus string

const (
	SubjectRightsRequestHistoryStageStatuscompleted  SubjectRightsRequestHistoryStageStatus = "Completed"
	SubjectRightsRequestHistoryStageStatuscurrent    SubjectRightsRequestHistoryStageStatus = "Current"
	SubjectRightsRequestHistoryStageStatusfailed     SubjectRightsRequestHistoryStageStatus = "Failed"
	SubjectRightsRequestHistoryStageStatusnotStarted SubjectRightsRequestHistoryStageStatus = "NotStarted"
)

func PossibleValuesForSubjectRightsRequestHistoryStageStatus() []string {
	return []string{
		string(SubjectRightsRequestHistoryStageStatuscompleted),
		string(SubjectRightsRequestHistoryStageStatuscurrent),
		string(SubjectRightsRequestHistoryStageStatusfailed),
		string(SubjectRightsRequestHistoryStageStatusnotStarted),
	}
}

func parseSubjectRightsRequestHistoryStageStatus(input string) (*SubjectRightsRequestHistoryStageStatus, error) {
	vals := map[string]SubjectRightsRequestHistoryStageStatus{
		"completed":  SubjectRightsRequestHistoryStageStatuscompleted,
		"current":    SubjectRightsRequestHistoryStageStatuscurrent,
		"failed":     SubjectRightsRequestHistoryStageStatusfailed,
		"notstarted": SubjectRightsRequestHistoryStageStatusnotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestHistoryStageStatus(input)
	return &out, nil
}
