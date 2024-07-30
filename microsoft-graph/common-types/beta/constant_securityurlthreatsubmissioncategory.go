package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionCategory string

const (
	SecurityUrlThreatSubmissionCategory_Malware  SecurityUrlThreatSubmissionCategory = "malware"
	SecurityUrlThreatSubmissionCategory_NotJunk  SecurityUrlThreatSubmissionCategory = "notJunk"
	SecurityUrlThreatSubmissionCategory_Phishing SecurityUrlThreatSubmissionCategory = "phishing"
	SecurityUrlThreatSubmissionCategory_Spam     SecurityUrlThreatSubmissionCategory = "spam"
)

func PossibleValuesForSecurityUrlThreatSubmissionCategory() []string {
	return []string{
		string(SecurityUrlThreatSubmissionCategory_Malware),
		string(SecurityUrlThreatSubmissionCategory_NotJunk),
		string(SecurityUrlThreatSubmissionCategory_Phishing),
		string(SecurityUrlThreatSubmissionCategory_Spam),
	}
}

func (s *SecurityUrlThreatSubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUrlThreatSubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUrlThreatSubmissionCategory(input string) (*SecurityUrlThreatSubmissionCategory, error) {
	vals := map[string]SecurityUrlThreatSubmissionCategory{
		"malware":  SecurityUrlThreatSubmissionCategory_Malware,
		"notjunk":  SecurityUrlThreatSubmissionCategory_NotJunk,
		"phishing": SecurityUrlThreatSubmissionCategory_Phishing,
		"spam":     SecurityUrlThreatSubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionCategory(input)
	return &out, nil
}
