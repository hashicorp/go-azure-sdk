package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStageDetailStatus string

const (
	SubjectRightsRequestStageDetailStatuscompleted  SubjectRightsRequestStageDetailStatus = "Completed"
	SubjectRightsRequestStageDetailStatuscurrent    SubjectRightsRequestStageDetailStatus = "Current"
	SubjectRightsRequestStageDetailStatusfailed     SubjectRightsRequestStageDetailStatus = "Failed"
	SubjectRightsRequestStageDetailStatusnotStarted SubjectRightsRequestStageDetailStatus = "NotStarted"
)

func PossibleValuesForSubjectRightsRequestStageDetailStatus() []string {
	return []string{
		string(SubjectRightsRequestStageDetailStatuscompleted),
		string(SubjectRightsRequestStageDetailStatuscurrent),
		string(SubjectRightsRequestStageDetailStatusfailed),
		string(SubjectRightsRequestStageDetailStatusnotStarted),
	}
}

func parseSubjectRightsRequestStageDetailStatus(input string) (*SubjectRightsRequestStageDetailStatus, error) {
	vals := map[string]SubjectRightsRequestStageDetailStatus{
		"completed":  SubjectRightsRequestStageDetailStatuscompleted,
		"current":    SubjectRightsRequestStageDetailStatuscurrent,
		"failed":     SubjectRightsRequestStageDetailStatusfailed,
		"notstarted": SubjectRightsRequestStageDetailStatusnotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestStageDetailStatus(input)
	return &out, nil
}
