package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestHistoryStage string

const (
	SubjectRightsRequestHistoryStageapproval         SubjectRightsRequestHistoryStage = "Approval"
	SubjectRightsRequestHistoryStagecaseResolved     SubjectRightsRequestHistoryStage = "CaseResolved"
	SubjectRightsRequestHistoryStagecontentDeletion  SubjectRightsRequestHistoryStage = "ContentDeletion"
	SubjectRightsRequestHistoryStagecontentEstimate  SubjectRightsRequestHistoryStage = "ContentEstimate"
	SubjectRightsRequestHistoryStagecontentRetrieval SubjectRightsRequestHistoryStage = "ContentRetrieval"
	SubjectRightsRequestHistoryStagecontentReview    SubjectRightsRequestHistoryStage = "ContentReview"
	SubjectRightsRequestHistoryStagegenerateReport   SubjectRightsRequestHistoryStage = "GenerateReport"
)

func PossibleValuesForSubjectRightsRequestHistoryStage() []string {
	return []string{
		string(SubjectRightsRequestHistoryStageapproval),
		string(SubjectRightsRequestHistoryStagecaseResolved),
		string(SubjectRightsRequestHistoryStagecontentDeletion),
		string(SubjectRightsRequestHistoryStagecontentEstimate),
		string(SubjectRightsRequestHistoryStagecontentRetrieval),
		string(SubjectRightsRequestHistoryStagecontentReview),
		string(SubjectRightsRequestHistoryStagegenerateReport),
	}
}

func parseSubjectRightsRequestHistoryStage(input string) (*SubjectRightsRequestHistoryStage, error) {
	vals := map[string]SubjectRightsRequestHistoryStage{
		"approval":         SubjectRightsRequestHistoryStageapproval,
		"caseresolved":     SubjectRightsRequestHistoryStagecaseResolved,
		"contentdeletion":  SubjectRightsRequestHistoryStagecontentDeletion,
		"contentestimate":  SubjectRightsRequestHistoryStagecontentEstimate,
		"contentretrieval": SubjectRightsRequestHistoryStagecontentRetrieval,
		"contentreview":    SubjectRightsRequestHistoryStagecontentReview,
		"generatereport":   SubjectRightsRequestHistoryStagegenerateReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestHistoryStage(input)
	return &out, nil
}
