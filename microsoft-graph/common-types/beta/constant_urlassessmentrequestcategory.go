package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UrlAssessmentRequestCategory string

const (
	UrlAssessmentRequestCategory_Malware   UrlAssessmentRequestCategory = "malware"
	UrlAssessmentRequestCategory_Phishing  UrlAssessmentRequestCategory = "phishing"
	UrlAssessmentRequestCategory_Spam      UrlAssessmentRequestCategory = "spam"
	UrlAssessmentRequestCategory_Undefined UrlAssessmentRequestCategory = "undefined"
)

func PossibleValuesForUrlAssessmentRequestCategory() []string {
	return []string{
		string(UrlAssessmentRequestCategory_Malware),
		string(UrlAssessmentRequestCategory_Phishing),
		string(UrlAssessmentRequestCategory_Spam),
		string(UrlAssessmentRequestCategory_Undefined),
	}
}

func (s *UrlAssessmentRequestCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUrlAssessmentRequestCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUrlAssessmentRequestCategory(input string) (*UrlAssessmentRequestCategory, error) {
	vals := map[string]UrlAssessmentRequestCategory{
		"malware":   UrlAssessmentRequestCategory_Malware,
		"phishing":  UrlAssessmentRequestCategory_Phishing,
		"spam":      UrlAssessmentRequestCategory_Spam,
		"undefined": UrlAssessmentRequestCategory_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UrlAssessmentRequestCategory(input)
	return &out, nil
}
