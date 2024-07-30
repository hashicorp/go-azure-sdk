package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileUrlThreatSubmissionCategory string

const (
	SecurityFileUrlThreatSubmissionCategory_Malware  SecurityFileUrlThreatSubmissionCategory = "malware"
	SecurityFileUrlThreatSubmissionCategory_NotJunk  SecurityFileUrlThreatSubmissionCategory = "notJunk"
	SecurityFileUrlThreatSubmissionCategory_Phishing SecurityFileUrlThreatSubmissionCategory = "phishing"
	SecurityFileUrlThreatSubmissionCategory_Spam     SecurityFileUrlThreatSubmissionCategory = "spam"
)

func PossibleValuesForSecurityFileUrlThreatSubmissionCategory() []string {
	return []string{
		string(SecurityFileUrlThreatSubmissionCategory_Malware),
		string(SecurityFileUrlThreatSubmissionCategory_NotJunk),
		string(SecurityFileUrlThreatSubmissionCategory_Phishing),
		string(SecurityFileUrlThreatSubmissionCategory_Spam),
	}
}

func (s *SecurityFileUrlThreatSubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileUrlThreatSubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileUrlThreatSubmissionCategory(input string) (*SecurityFileUrlThreatSubmissionCategory, error) {
	vals := map[string]SecurityFileUrlThreatSubmissionCategory{
		"malware":  SecurityFileUrlThreatSubmissionCategory_Malware,
		"notjunk":  SecurityFileUrlThreatSubmissionCategory_NotJunk,
		"phishing": SecurityFileUrlThreatSubmissionCategory_Phishing,
		"spam":     SecurityFileUrlThreatSubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileUrlThreatSubmissionCategory(input)
	return &out, nil
}
