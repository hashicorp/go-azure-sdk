package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestStatus string

const (
	UrlAssessmentRequestStatus_Completed UrlAssessmentRequestStatus = "completed"
	UrlAssessmentRequestStatus_Pending   UrlAssessmentRequestStatus = "pending"
)

func PossibleValuesForUrlAssessmentRequestStatus() []string {
	return []string{
		string(UrlAssessmentRequestStatus_Completed),
		string(UrlAssessmentRequestStatus_Pending),
	}
}

func (s *UrlAssessmentRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUrlAssessmentRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUrlAssessmentRequestStatus(input string) (*UrlAssessmentRequestStatus, error) {
	vals := map[string]UrlAssessmentRequestStatus{
		"completed": UrlAssessmentRequestStatus_Completed,
		"pending":   UrlAssessmentRequestStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestStatus(input)
	return &out, nil
}
