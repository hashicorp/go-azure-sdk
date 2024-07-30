package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionOriginalCategory string

const (
	SecurityEmailContentThreatSubmissionOriginalCategory_Malware  SecurityEmailContentThreatSubmissionOriginalCategory = "malware"
	SecurityEmailContentThreatSubmissionOriginalCategory_NotJunk  SecurityEmailContentThreatSubmissionOriginalCategory = "notJunk"
	SecurityEmailContentThreatSubmissionOriginalCategory_Phishing SecurityEmailContentThreatSubmissionOriginalCategory = "phishing"
	SecurityEmailContentThreatSubmissionOriginalCategory_Spam     SecurityEmailContentThreatSubmissionOriginalCategory = "spam"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionOriginalCategory() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionOriginalCategory_Malware),
		string(SecurityEmailContentThreatSubmissionOriginalCategory_NotJunk),
		string(SecurityEmailContentThreatSubmissionOriginalCategory_Phishing),
		string(SecurityEmailContentThreatSubmissionOriginalCategory_Spam),
	}
}

func (s *SecurityEmailContentThreatSubmissionOriginalCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailContentThreatSubmissionOriginalCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailContentThreatSubmissionOriginalCategory(input string) (*SecurityEmailContentThreatSubmissionOriginalCategory, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionOriginalCategory{
		"malware":  SecurityEmailContentThreatSubmissionOriginalCategory_Malware,
		"notjunk":  SecurityEmailContentThreatSubmissionOriginalCategory_NotJunk,
		"phishing": SecurityEmailContentThreatSubmissionOriginalCategory_Phishing,
		"spam":     SecurityEmailContentThreatSubmissionOriginalCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionOriginalCategory(input)
	return &out, nil
}
