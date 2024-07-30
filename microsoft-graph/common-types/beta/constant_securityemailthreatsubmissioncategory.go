package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionCategory string

const (
	SecurityEmailThreatSubmissionCategory_Malware  SecurityEmailThreatSubmissionCategory = "malware"
	SecurityEmailThreatSubmissionCategory_NotJunk  SecurityEmailThreatSubmissionCategory = "notJunk"
	SecurityEmailThreatSubmissionCategory_Phishing SecurityEmailThreatSubmissionCategory = "phishing"
	SecurityEmailThreatSubmissionCategory_Spam     SecurityEmailThreatSubmissionCategory = "spam"
)

func PossibleValuesForSecurityEmailThreatSubmissionCategory() []string {
	return []string{
		string(SecurityEmailThreatSubmissionCategory_Malware),
		string(SecurityEmailThreatSubmissionCategory_NotJunk),
		string(SecurityEmailThreatSubmissionCategory_Phishing),
		string(SecurityEmailThreatSubmissionCategory_Spam),
	}
}

func (s *SecurityEmailThreatSubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailThreatSubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailThreatSubmissionCategory(input string) (*SecurityEmailThreatSubmissionCategory, error) {
	vals := map[string]SecurityEmailThreatSubmissionCategory{
		"malware":  SecurityEmailThreatSubmissionCategory_Malware,
		"notjunk":  SecurityEmailThreatSubmissionCategory_NotJunk,
		"phishing": SecurityEmailThreatSubmissionCategory_Phishing,
		"spam":     SecurityEmailThreatSubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionCategory(input)
	return &out, nil
}
