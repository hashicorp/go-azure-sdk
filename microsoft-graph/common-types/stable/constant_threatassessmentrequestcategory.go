package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestCategory string

const (
	ThreatAssessmentRequestCategory_Malware   ThreatAssessmentRequestCategory = "malware"
	ThreatAssessmentRequestCategory_Phishing  ThreatAssessmentRequestCategory = "phishing"
	ThreatAssessmentRequestCategory_Spam      ThreatAssessmentRequestCategory = "spam"
	ThreatAssessmentRequestCategory_Undefined ThreatAssessmentRequestCategory = "undefined"
)

func PossibleValuesForThreatAssessmentRequestCategory() []string {
	return []string{
		string(ThreatAssessmentRequestCategory_Malware),
		string(ThreatAssessmentRequestCategory_Phishing),
		string(ThreatAssessmentRequestCategory_Spam),
		string(ThreatAssessmentRequestCategory_Undefined),
	}
}

func (s *ThreatAssessmentRequestCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestCategory(input string) (*ThreatAssessmentRequestCategory, error) {
	vals := map[string]ThreatAssessmentRequestCategory{
		"malware":   ThreatAssessmentRequestCategory_Malware,
		"phishing":  ThreatAssessmentRequestCategory_Phishing,
		"spam":      ThreatAssessmentRequestCategory_Spam,
		"undefined": ThreatAssessmentRequestCategory_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestCategory(input)
	return &out, nil
}
