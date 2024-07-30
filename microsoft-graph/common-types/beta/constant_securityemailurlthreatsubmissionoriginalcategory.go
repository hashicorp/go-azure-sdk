package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionOriginalCategory string

const (
	SecurityEmailUrlThreatSubmissionOriginalCategory_Malware  SecurityEmailUrlThreatSubmissionOriginalCategory = "malware"
	SecurityEmailUrlThreatSubmissionOriginalCategory_NotJunk  SecurityEmailUrlThreatSubmissionOriginalCategory = "notJunk"
	SecurityEmailUrlThreatSubmissionOriginalCategory_Phishing SecurityEmailUrlThreatSubmissionOriginalCategory = "phishing"
	SecurityEmailUrlThreatSubmissionOriginalCategory_Spam     SecurityEmailUrlThreatSubmissionOriginalCategory = "spam"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionOriginalCategory() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionOriginalCategory_Malware),
		string(SecurityEmailUrlThreatSubmissionOriginalCategory_NotJunk),
		string(SecurityEmailUrlThreatSubmissionOriginalCategory_Phishing),
		string(SecurityEmailUrlThreatSubmissionOriginalCategory_Spam),
	}
}

func (s *SecurityEmailUrlThreatSubmissionOriginalCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailUrlThreatSubmissionOriginalCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailUrlThreatSubmissionOriginalCategory(input string) (*SecurityEmailUrlThreatSubmissionOriginalCategory, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionOriginalCategory{
		"malware":  SecurityEmailUrlThreatSubmissionOriginalCategory_Malware,
		"notjunk":  SecurityEmailUrlThreatSubmissionOriginalCategory_NotJunk,
		"phishing": SecurityEmailUrlThreatSubmissionOriginalCategory_Phishing,
		"spam":     SecurityEmailUrlThreatSubmissionOriginalCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionOriginalCategory(input)
	return &out, nil
}
