package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestHistoryStageStatus string

const (
	SubjectRightsRequestHistoryStageStatus_Completed  SubjectRightsRequestHistoryStageStatus = "completed"
	SubjectRightsRequestHistoryStageStatus_Current    SubjectRightsRequestHistoryStageStatus = "current"
	SubjectRightsRequestHistoryStageStatus_Failed     SubjectRightsRequestHistoryStageStatus = "failed"
	SubjectRightsRequestHistoryStageStatus_NotStarted SubjectRightsRequestHistoryStageStatus = "notStarted"
)

func PossibleValuesForSubjectRightsRequestHistoryStageStatus() []string {
	return []string{
		string(SubjectRightsRequestHistoryStageStatus_Completed),
		string(SubjectRightsRequestHistoryStageStatus_Current),
		string(SubjectRightsRequestHistoryStageStatus_Failed),
		string(SubjectRightsRequestHistoryStageStatus_NotStarted),
	}
}

func (s *SubjectRightsRequestHistoryStageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestHistoryStageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestHistoryStageStatus(input string) (*SubjectRightsRequestHistoryStageStatus, error) {
	vals := map[string]SubjectRightsRequestHistoryStageStatus{
		"completed":  SubjectRightsRequestHistoryStageStatus_Completed,
		"current":    SubjectRightsRequestHistoryStageStatus_Current,
		"failed":     SubjectRightsRequestHistoryStageStatus_Failed,
		"notstarted": SubjectRightsRequestHistoryStageStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestHistoryStageStatus(input)
	return &out, nil
}
