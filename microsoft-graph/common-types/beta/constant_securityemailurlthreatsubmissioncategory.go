package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionCategory string

const (
	SecurityEmailUrlThreatSubmissionCategory_Malware  SecurityEmailUrlThreatSubmissionCategory = "malware"
	SecurityEmailUrlThreatSubmissionCategory_NotJunk  SecurityEmailUrlThreatSubmissionCategory = "notJunk"
	SecurityEmailUrlThreatSubmissionCategory_Phishing SecurityEmailUrlThreatSubmissionCategory = "phishing"
	SecurityEmailUrlThreatSubmissionCategory_Spam     SecurityEmailUrlThreatSubmissionCategory = "spam"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionCategory() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionCategory_Malware),
		string(SecurityEmailUrlThreatSubmissionCategory_NotJunk),
		string(SecurityEmailUrlThreatSubmissionCategory_Phishing),
		string(SecurityEmailUrlThreatSubmissionCategory_Spam),
	}
}

func (s *SecurityEmailUrlThreatSubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailUrlThreatSubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailUrlThreatSubmissionCategory(input string) (*SecurityEmailUrlThreatSubmissionCategory, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionCategory{
		"malware":  SecurityEmailUrlThreatSubmissionCategory_Malware,
		"notjunk":  SecurityEmailUrlThreatSubmissionCategory_NotJunk,
		"phishing": SecurityEmailUrlThreatSubmissionCategory_Phishing,
		"spam":     SecurityEmailUrlThreatSubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionCategory(input)
	return &out, nil
}
