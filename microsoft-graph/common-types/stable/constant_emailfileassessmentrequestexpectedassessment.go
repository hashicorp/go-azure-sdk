package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestExpectedAssessment string

const (
	EmailFileAssessmentRequestExpectedAssessment_Block   EmailFileAssessmentRequestExpectedAssessment = "block"
	EmailFileAssessmentRequestExpectedAssessment_Unblock EmailFileAssessmentRequestExpectedAssessment = "unblock"
)

func PossibleValuesForEmailFileAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(EmailFileAssessmentRequestExpectedAssessment_Block),
		string(EmailFileAssessmentRequestExpectedAssessment_Unblock),
	}
}

func (s *EmailFileAssessmentRequestExpectedAssessment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailFileAssessmentRequestExpectedAssessment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailFileAssessmentRequestExpectedAssessment(input string) (*EmailFileAssessmentRequestExpectedAssessment, error) {
	vals := map[string]EmailFileAssessmentRequestExpectedAssessment{
		"block":   EmailFileAssessmentRequestExpectedAssessment_Block,
		"unblock": EmailFileAssessmentRequestExpectedAssessment_Unblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
