package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionOriginalCategory string

const (
	SecurityEmailThreatSubmissionOriginalCategory_Malware  SecurityEmailThreatSubmissionOriginalCategory = "malware"
	SecurityEmailThreatSubmissionOriginalCategory_NotJunk  SecurityEmailThreatSubmissionOriginalCategory = "notJunk"
	SecurityEmailThreatSubmissionOriginalCategory_Phishing SecurityEmailThreatSubmissionOriginalCategory = "phishing"
	SecurityEmailThreatSubmissionOriginalCategory_Spam     SecurityEmailThreatSubmissionOriginalCategory = "spam"
)

func PossibleValuesForSecurityEmailThreatSubmissionOriginalCategory() []string {
	return []string{
		string(SecurityEmailThreatSubmissionOriginalCategory_Malware),
		string(SecurityEmailThreatSubmissionOriginalCategory_NotJunk),
		string(SecurityEmailThreatSubmissionOriginalCategory_Phishing),
		string(SecurityEmailThreatSubmissionOriginalCategory_Spam),
	}
}

func (s *SecurityEmailThreatSubmissionOriginalCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailThreatSubmissionOriginalCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailThreatSubmissionOriginalCategory(input string) (*SecurityEmailThreatSubmissionOriginalCategory, error) {
	vals := map[string]SecurityEmailThreatSubmissionOriginalCategory{
		"malware":  SecurityEmailThreatSubmissionOriginalCategory_Malware,
		"notjunk":  SecurityEmailThreatSubmissionOriginalCategory_NotJunk,
		"phishing": SecurityEmailThreatSubmissionOriginalCategory_Phishing,
		"spam":     SecurityEmailThreatSubmissionOriginalCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionOriginalCategory(input)
	return &out, nil
}
