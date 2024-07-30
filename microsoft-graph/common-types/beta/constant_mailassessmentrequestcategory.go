package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailAssessmentRequestCategory string

const (
	MailAssessmentRequestCategory_Malware   MailAssessmentRequestCategory = "malware"
	MailAssessmentRequestCategory_Phishing  MailAssessmentRequestCategory = "phishing"
	MailAssessmentRequestCategory_Spam      MailAssessmentRequestCategory = "spam"
	MailAssessmentRequestCategory_Undefined MailAssessmentRequestCategory = "undefined"
)

func PossibleValuesForMailAssessmentRequestCategory() []string {
	return []string{
		string(MailAssessmentRequestCategory_Malware),
		string(MailAssessmentRequestCategory_Phishing),
		string(MailAssessmentRequestCategory_Spam),
		string(MailAssessmentRequestCategory_Undefined),
	}
}

func (s *MailAssessmentRequestCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailAssessmentRequestCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailAssessmentRequestCategory(input string) (*MailAssessmentRequestCategory, error) {
	vals := map[string]MailAssessmentRequestCategory{
		"malware":   MailAssessmentRequestCategory_Malware,
		"phishing":  MailAssessmentRequestCategory_Phishing,
		"spam":      MailAssessmentRequestCategory_Spam,
		"undefined": MailAssessmentRequestCategory_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailAssessmentRequestCategory(input)
	return &out, nil
}
