package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestExpectedAssessment string

const (
	ThreatAssessmentRequestExpectedAssessment_Block   ThreatAssessmentRequestExpectedAssessment = "block"
	ThreatAssessmentRequestExpectedAssessment_Unblock ThreatAssessmentRequestExpectedAssessment = "unblock"
)

func PossibleValuesForThreatAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(ThreatAssessmentRequestExpectedAssessment_Block),
		string(ThreatAssessmentRequestExpectedAssessment_Unblock),
	}
}

func (s *ThreatAssessmentRequestExpectedAssessment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestExpectedAssessment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestExpectedAssessment(input string) (*ThreatAssessmentRequestExpectedAssessment, error) {
	vals := map[string]ThreatAssessmentRequestExpectedAssessment{
		"block":   ThreatAssessmentRequestExpectedAssessment_Block,
		"unblock": ThreatAssessmentRequestExpectedAssessment_Unblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
