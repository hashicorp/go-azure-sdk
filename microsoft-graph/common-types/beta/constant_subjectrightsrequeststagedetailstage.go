package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStageDetailStage string

const (
	SubjectRightsRequestStageDetailStage_Approval         SubjectRightsRequestStageDetailStage = "approval"
	SubjectRightsRequestStageDetailStage_CaseResolved     SubjectRightsRequestStageDetailStage = "caseResolved"
	SubjectRightsRequestStageDetailStage_ContentDeletion  SubjectRightsRequestStageDetailStage = "contentDeletion"
	SubjectRightsRequestStageDetailStage_ContentEstimate  SubjectRightsRequestStageDetailStage = "contentEstimate"
	SubjectRightsRequestStageDetailStage_ContentRetrieval SubjectRightsRequestStageDetailStage = "contentRetrieval"
	SubjectRightsRequestStageDetailStage_ContentReview    SubjectRightsRequestStageDetailStage = "contentReview"
	SubjectRightsRequestStageDetailStage_GenerateReport   SubjectRightsRequestStageDetailStage = "generateReport"
)

func PossibleValuesForSubjectRightsRequestStageDetailStage() []string {
	return []string{
		string(SubjectRightsRequestStageDetailStage_Approval),
		string(SubjectRightsRequestStageDetailStage_CaseResolved),
		string(SubjectRightsRequestStageDetailStage_ContentDeletion),
		string(SubjectRightsRequestStageDetailStage_ContentEstimate),
		string(SubjectRightsRequestStageDetailStage_ContentRetrieval),
		string(SubjectRightsRequestStageDetailStage_ContentReview),
		string(SubjectRightsRequestStageDetailStage_GenerateReport),
	}
}

func (s *SubjectRightsRequestStageDetailStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestStageDetailStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestStageDetailStage(input string) (*SubjectRightsRequestStageDetailStage, error) {
	vals := map[string]SubjectRightsRequestStageDetailStage{
		"approval":         SubjectRightsRequestStageDetailStage_Approval,
		"caseresolved":     SubjectRightsRequestStageDetailStage_CaseResolved,
		"contentdeletion":  SubjectRightsRequestStageDetailStage_ContentDeletion,
		"contentestimate":  SubjectRightsRequestStageDetailStage_ContentEstimate,
		"contentretrieval": SubjectRightsRequestStageDetailStage_ContentRetrieval,
		"contentreview":    SubjectRightsRequestStageDetailStage_ContentReview,
		"generatereport":   SubjectRightsRequestStageDetailStage_GenerateReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestStageDetailStage(input)
	return &out, nil
}
