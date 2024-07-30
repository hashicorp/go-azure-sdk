package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestExpectedAssessment string

const (
	UrlAssessmentRequestExpectedAssessment_Block   UrlAssessmentRequestExpectedAssessment = "block"
	UrlAssessmentRequestExpectedAssessment_Unblock UrlAssessmentRequestExpectedAssessment = "unblock"
)

func PossibleValuesForUrlAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(UrlAssessmentRequestExpectedAssessment_Block),
		string(UrlAssessmentRequestExpectedAssessment_Unblock),
	}
}

func (s *UrlAssessmentRequestExpectedAssessment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUrlAssessmentRequestExpectedAssessment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUrlAssessmentRequestExpectedAssessment(input string) (*UrlAssessmentRequestExpectedAssessment, error) {
	vals := map[string]UrlAssessmentRequestExpectedAssessment{
		"block":   UrlAssessmentRequestExpectedAssessment_Block,
		"unblock": UrlAssessmentRequestExpectedAssessment_Unblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
