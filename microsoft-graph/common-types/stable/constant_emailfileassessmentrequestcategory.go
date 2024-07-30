package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestCategory string

const (
	EmailFileAssessmentRequestCategory_Malware   EmailFileAssessmentRequestCategory = "malware"
	EmailFileAssessmentRequestCategory_Phishing  EmailFileAssessmentRequestCategory = "phishing"
	EmailFileAssessmentRequestCategory_Spam      EmailFileAssessmentRequestCategory = "spam"
	EmailFileAssessmentRequestCategory_Undefined EmailFileAssessmentRequestCategory = "undefined"
)

func PossibleValuesForEmailFileAssessmentRequestCategory() []string {
	return []string{
		string(EmailFileAssessmentRequestCategory_Malware),
		string(EmailFileAssessmentRequestCategory_Phishing),
		string(EmailFileAssessmentRequestCategory_Spam),
		string(EmailFileAssessmentRequestCategory_Undefined),
	}
}

func (s *EmailFileAssessmentRequestCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEmailFileAssessmentRequestCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEmailFileAssessmentRequestCategory(input string) (*EmailFileAssessmentRequestCategory, error) {
	vals := map[string]EmailFileAssessmentRequestCategory{
		"malware":   EmailFileAssessmentRequestCategory_Malware,
		"phishing":  EmailFileAssessmentRequestCategory_Phishing,
		"spam":      EmailFileAssessmentRequestCategory_Spam,
		"undefined": EmailFileAssessmentRequestCategory_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EmailFileAssessmentRequestCategory(input)
	return &out, nil
}
