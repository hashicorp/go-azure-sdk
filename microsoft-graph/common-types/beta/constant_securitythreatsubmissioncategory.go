package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionCategory string

const (
	SecurityThreatSubmissionCategory_Malware  SecurityThreatSubmissionCategory = "malware"
	SecurityThreatSubmissionCategory_NotJunk  SecurityThreatSubmissionCategory = "notJunk"
	SecurityThreatSubmissionCategory_Phishing SecurityThreatSubmissionCategory = "phishing"
	SecurityThreatSubmissionCategory_Spam     SecurityThreatSubmissionCategory = "spam"
)

func PossibleValuesForSecurityThreatSubmissionCategory() []string {
	return []string{
		string(SecurityThreatSubmissionCategory_Malware),
		string(SecurityThreatSubmissionCategory_NotJunk),
		string(SecurityThreatSubmissionCategory_Phishing),
		string(SecurityThreatSubmissionCategory_Spam),
	}
}

func (s *SecurityThreatSubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityThreatSubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityThreatSubmissionCategory(input string) (*SecurityThreatSubmissionCategory, error) {
	vals := map[string]SecurityThreatSubmissionCategory{
		"malware":  SecurityThreatSubmissionCategory_Malware,
		"notjunk":  SecurityThreatSubmissionCategory_NotJunk,
		"phishing": SecurityThreatSubmissionCategory_Phishing,
		"spam":     SecurityThreatSubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionCategory(input)
	return &out, nil
}
