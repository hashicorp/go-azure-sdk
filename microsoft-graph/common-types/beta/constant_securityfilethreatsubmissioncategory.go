package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmissionCategory string

const (
	SecurityFileThreatSubmissionCategory_Malware  SecurityFileThreatSubmissionCategory = "malware"
	SecurityFileThreatSubmissionCategory_NotJunk  SecurityFileThreatSubmissionCategory = "notJunk"
	SecurityFileThreatSubmissionCategory_Phishing SecurityFileThreatSubmissionCategory = "phishing"
	SecurityFileThreatSubmissionCategory_Spam     SecurityFileThreatSubmissionCategory = "spam"
)

func PossibleValuesForSecurityFileThreatSubmissionCategory() []string {
	return []string{
		string(SecurityFileThreatSubmissionCategory_Malware),
		string(SecurityFileThreatSubmissionCategory_NotJunk),
		string(SecurityFileThreatSubmissionCategory_Phishing),
		string(SecurityFileThreatSubmissionCategory_Spam),
	}
}

func (s *SecurityFileThreatSubmissionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileThreatSubmissionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileThreatSubmissionCategory(input string) (*SecurityFileThreatSubmissionCategory, error) {
	vals := map[string]SecurityFileThreatSubmissionCategory{
		"malware":  SecurityFileThreatSubmissionCategory_Malware,
		"notjunk":  SecurityFileThreatSubmissionCategory_NotJunk,
		"phishing": SecurityFileThreatSubmissionCategory_Phishing,
		"spam":     SecurityFileThreatSubmissionCategory_Spam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileThreatSubmissionCategory(input)
	return &out, nil
}
