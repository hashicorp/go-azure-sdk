package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestStatus string

const (
	ThreatAssessmentRequestStatus_Completed ThreatAssessmentRequestStatus = "completed"
	ThreatAssessmentRequestStatus_Pending   ThreatAssessmentRequestStatus = "pending"
)

func PossibleValuesForThreatAssessmentRequestStatus() []string {
	return []string{
		string(ThreatAssessmentRequestStatus_Completed),
		string(ThreatAssessmentRequestStatus_Pending),
	}
}

func (s *ThreatAssessmentRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestStatus(input string) (*ThreatAssessmentRequestStatus, error) {
	vals := map[string]ThreatAssessmentRequestStatus{
		"completed": ThreatAssessmentRequestStatus_Completed,
		"pending":   ThreatAssessmentRequestStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestStatus(input)
	return &out, nil
}
