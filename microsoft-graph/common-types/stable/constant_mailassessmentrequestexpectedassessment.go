package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestExpectedAssessment string

const (
	MailAssessmentRequestExpectedAssessment_Block   MailAssessmentRequestExpectedAssessment = "block"
	MailAssessmentRequestExpectedAssessment_Unblock MailAssessmentRequestExpectedAssessment = "unblock"
)

func PossibleValuesForMailAssessmentRequestExpectedAssessment() []string {
	return []string{
		string(MailAssessmentRequestExpectedAssessment_Block),
		string(MailAssessmentRequestExpectedAssessment_Unblock),
	}
}

func (s *MailAssessmentRequestExpectedAssessment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailAssessmentRequestExpectedAssessment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailAssessmentRequestExpectedAssessment(input string) (*MailAssessmentRequestExpectedAssessment, error) {
	vals := map[string]MailAssessmentRequestExpectedAssessment{
		"block":   MailAssessmentRequestExpectedAssessment_Block,
		"unblock": MailAssessmentRequestExpectedAssessment_Unblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestExpectedAssessment(input)
	return &out, nil
}
