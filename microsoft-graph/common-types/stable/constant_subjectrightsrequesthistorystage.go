package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestHistoryStage string

const (
	SubjectRightsRequestHistoryStage_Approval         SubjectRightsRequestHistoryStage = "approval"
	SubjectRightsRequestHistoryStage_CaseResolved     SubjectRightsRequestHistoryStage = "caseResolved"
	SubjectRightsRequestHistoryStage_ContentDeletion  SubjectRightsRequestHistoryStage = "contentDeletion"
	SubjectRightsRequestHistoryStage_ContentEstimate  SubjectRightsRequestHistoryStage = "contentEstimate"
	SubjectRightsRequestHistoryStage_ContentRetrieval SubjectRightsRequestHistoryStage = "contentRetrieval"
	SubjectRightsRequestHistoryStage_ContentReview    SubjectRightsRequestHistoryStage = "contentReview"
	SubjectRightsRequestHistoryStage_GenerateReport   SubjectRightsRequestHistoryStage = "generateReport"
)

func PossibleValuesForSubjectRightsRequestHistoryStage() []string {
	return []string{
		string(SubjectRightsRequestHistoryStage_Approval),
		string(SubjectRightsRequestHistoryStage_CaseResolved),
		string(SubjectRightsRequestHistoryStage_ContentDeletion),
		string(SubjectRightsRequestHistoryStage_ContentEstimate),
		string(SubjectRightsRequestHistoryStage_ContentRetrieval),
		string(SubjectRightsRequestHistoryStage_ContentReview),
		string(SubjectRightsRequestHistoryStage_GenerateReport),
	}
}

func (s *SubjectRightsRequestHistoryStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestHistoryStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestHistoryStage(input string) (*SubjectRightsRequestHistoryStage, error) {
	vals := map[string]SubjectRightsRequestHistoryStage{
		"approval":         SubjectRightsRequestHistoryStage_Approval,
		"caseresolved":     SubjectRightsRequestHistoryStage_CaseResolved,
		"contentdeletion":  SubjectRightsRequestHistoryStage_ContentDeletion,
		"contentestimate":  SubjectRightsRequestHistoryStage_ContentEstimate,
		"contentretrieval": SubjectRightsRequestHistoryStage_ContentRetrieval,
		"contentreview":    SubjectRightsRequestHistoryStage_ContentReview,
		"generatereport":   SubjectRightsRequestHistoryStage_GenerateReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestHistoryStage(input)
	return &out, nil
}
