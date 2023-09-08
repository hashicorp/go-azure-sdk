package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStageDetailStage string

const (
	SubjectRightsRequestStageDetailStagecaseResolved     SubjectRightsRequestStageDetailStage = "CaseResolved"
	SubjectRightsRequestStageDetailStagecontentDeletion  SubjectRightsRequestStageDetailStage = "ContentDeletion"
	SubjectRightsRequestStageDetailStagecontentEstimate  SubjectRightsRequestStageDetailStage = "ContentEstimate"
	SubjectRightsRequestStageDetailStagecontentRetrieval SubjectRightsRequestStageDetailStage = "ContentRetrieval"
	SubjectRightsRequestStageDetailStagecontentReview    SubjectRightsRequestStageDetailStage = "ContentReview"
	SubjectRightsRequestStageDetailStagegenerateReport   SubjectRightsRequestStageDetailStage = "GenerateReport"
)

func PossibleValuesForSubjectRightsRequestStageDetailStage() []string {
	return []string{
		string(SubjectRightsRequestStageDetailStagecaseResolved),
		string(SubjectRightsRequestStageDetailStagecontentDeletion),
		string(SubjectRightsRequestStageDetailStagecontentEstimate),
		string(SubjectRightsRequestStageDetailStagecontentRetrieval),
		string(SubjectRightsRequestStageDetailStagecontentReview),
		string(SubjectRightsRequestStageDetailStagegenerateReport),
	}
}

func parseSubjectRightsRequestStageDetailStage(input string) (*SubjectRightsRequestStageDetailStage, error) {
	vals := map[string]SubjectRightsRequestStageDetailStage{
		"caseresolved":     SubjectRightsRequestStageDetailStagecaseResolved,
		"contentdeletion":  SubjectRightsRequestStageDetailStagecontentDeletion,
		"contentestimate":  SubjectRightsRequestStageDetailStagecontentEstimate,
		"contentretrieval": SubjectRightsRequestStageDetailStagecontentRetrieval,
		"contentreview":    SubjectRightsRequestStageDetailStagecontentReview,
		"generatereport":   SubjectRightsRequestStageDetailStagegenerateReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestStageDetailStage(input)
	return &out, nil
}
