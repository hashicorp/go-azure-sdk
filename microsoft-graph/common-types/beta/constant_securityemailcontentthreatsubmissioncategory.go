package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionCategory string

const (
	SecurityEmailContentThreatSubmissionCategory_Malware  SecurityEmailContentThreatSubmissionCategory = "malware"
	SecurityEmailContentThreatSubmissionCategory_NotJunk  SecurityEmailContentThreatSubmissionCategory = "notJunk"
	SecurityEmailContentThreatSubmissionCategory_Phishing SecurityEmailContentThreatSubmissionCategory = "phishing"
	SecurityEmailContentThreatSubmissionCategory_Spam     SecurityEmailContentThreatSubmissionCategory = "spam"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionCategory() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionCategory_Malware),
		string(SecurityEmailContentThreatSubmissionCategory_NotJunk),
		string(SecurityEmailContentThreatSubmissionCategory_Phishing),
		string(SecurityEmailContentThreatSubmissionCategory_Spam),
	}
}

func (s *SecurityEmailContentThreatSubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailContentThreatSubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailContentThreatSubmissionCategory(input string) (*SecurityEmailContentThreatSubmissionCategory, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionCategory{
		"malware":  SecurityEmailContentThreatSubmissionCategory_Malware,
		"notjunk":  SecurityEmailContentThreatSubmissionCategory_NotJunk,
		"phishing": SecurityEmailContentThreatSubmissionCategory_Phishing,
		"spam":     SecurityEmailContentThreatSubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionCategory(input)
	return &out, nil
}
