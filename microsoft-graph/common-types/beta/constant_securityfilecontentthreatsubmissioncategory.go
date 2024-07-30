package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionCategory string

const (
	SecurityFileContentThreatSubmissionCategory_Malware  SecurityFileContentThreatSubmissionCategory = "malware"
	SecurityFileContentThreatSubmissionCategory_NotJunk  SecurityFileContentThreatSubmissionCategory = "notJunk"
	SecurityFileContentThreatSubmissionCategory_Phishing SecurityFileContentThreatSubmissionCategory = "phishing"
	SecurityFileContentThreatSubmissionCategory_Spam     SecurityFileContentThreatSubmissionCategory = "spam"
)

func PossibleValuesForSecurityFileContentThreatSubmissionCategory() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionCategory_Malware),
		string(SecurityFileContentThreatSubmissionCategory_NotJunk),
		string(SecurityFileContentThreatSubmissionCategory_Phishing),
		string(SecurityFileContentThreatSubmissionCategory_Spam),
	}
}

func (s *SecurityFileContentThreatSubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileContentThreatSubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileContentThreatSubmissionCategory(input string) (*SecurityFileContentThreatSubmissionCategory, error) {
	vals := map[string]SecurityFileContentThreatSubmissionCategory{
		"malware":  SecurityFileContentThreatSubmissionCategory_Malware,
		"notjunk":  SecurityFileContentThreatSubmissionCategory_NotJunk,
		"phishing": SecurityFileContentThreatSubmissionCategory_Phishing,
		"spam":     SecurityFileContentThreatSubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionCategory(input)
	return &out, nil
}
