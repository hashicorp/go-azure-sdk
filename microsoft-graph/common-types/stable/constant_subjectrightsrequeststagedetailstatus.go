package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStageDetailStatus string

const (
	SubjectRightsRequestStageDetailStatus_Completed  SubjectRightsRequestStageDetailStatus = "completed"
	SubjectRightsRequestStageDetailStatus_Current    SubjectRightsRequestStageDetailStatus = "current"
	SubjectRightsRequestStageDetailStatus_Failed     SubjectRightsRequestStageDetailStatus = "failed"
	SubjectRightsRequestStageDetailStatus_NotStarted SubjectRightsRequestStageDetailStatus = "notStarted"
)

func PossibleValuesForSubjectRightsRequestStageDetailStatus() []string {
	return []string{
		string(SubjectRightsRequestStageDetailStatus_Completed),
		string(SubjectRightsRequestStageDetailStatus_Current),
		string(SubjectRightsRequestStageDetailStatus_Failed),
		string(SubjectRightsRequestStageDetailStatus_NotStarted),
	}
}

func (s *SubjectRightsRequestStageDetailStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestStageDetailStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestStageDetailStatus(input string) (*SubjectRightsRequestStageDetailStatus, error) {
	vals := map[string]SubjectRightsRequestStageDetailStatus{
		"completed":  SubjectRightsRequestStageDetailStatus_Completed,
		"current":    SubjectRightsRequestStageDetailStatus_Current,
		"failed":     SubjectRightsRequestStageDetailStatus_Failed,
		"notstarted": SubjectRightsRequestStageDetailStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestStageDetailStatus(input)
	return &out, nil
}
