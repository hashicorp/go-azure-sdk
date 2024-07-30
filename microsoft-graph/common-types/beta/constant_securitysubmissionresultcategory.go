package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionResultCategory string

const (
	SecuritySubmissionResultCategory_AllowedByPolicy   SecuritySubmissionResultCategory = "allowedByPolicy"
	SecuritySubmissionResultCategory_BlockedByPolicy   SecuritySubmissionResultCategory = "blockedByPolicy"
	SecuritySubmissionResultCategory_Malware           SecuritySubmissionResultCategory = "malware"
	SecuritySubmissionResultCategory_NoResultAvailable SecuritySubmissionResultCategory = "noResultAvailable"
	SecuritySubmissionResultCategory_NotJunk           SecuritySubmissionResultCategory = "notJunk"
	SecuritySubmissionResultCategory_Phishing          SecuritySubmissionResultCategory = "phishing"
	SecuritySubmissionResultCategory_Spam              SecuritySubmissionResultCategory = "spam"
	SecuritySubmissionResultCategory_Spoof             SecuritySubmissionResultCategory = "spoof"
	SecuritySubmissionResultCategory_Unknown           SecuritySubmissionResultCategory = "unknown"
)

func PossibleValuesForSecuritySubmissionResultCategory() []string {
	return []string{
		string(SecuritySubmissionResultCategory_AllowedByPolicy),
		string(SecuritySubmissionResultCategory_BlockedByPolicy),
		string(SecuritySubmissionResultCategory_Malware),
		string(SecuritySubmissionResultCategory_NoResultAvailable),
		string(SecuritySubmissionResultCategory_NotJunk),
		string(SecuritySubmissionResultCategory_Phishing),
		string(SecuritySubmissionResultCategory_Spam),
		string(SecuritySubmissionResultCategory_Spoof),
		string(SecuritySubmissionResultCategory_Unknown),
	}
}

func (s *SecuritySubmissionResultCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionResultCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionResultCategory(input string) (*SecuritySubmissionResultCategory, error) {
	vals := map[string]SecuritySubmissionResultCategory{
		"allowedbypolicy":   SecuritySubmissionResultCategory_AllowedByPolicy,
		"blockedbypolicy":   SecuritySubmissionResultCategory_BlockedByPolicy,
		"malware":           SecuritySubmissionResultCategory_Malware,
		"noresultavailable": SecuritySubmissionResultCategory_NoResultAvailable,
		"notjunk":           SecuritySubmissionResultCategory_NotJunk,
		"phishing":          SecuritySubmissionResultCategory_Phishing,
		"spam":              SecuritySubmissionResultCategory_Spam,
		"spoof":             SecuritySubmissionResultCategory_Spoof,
		"unknown":           SecuritySubmissionResultCategory_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionResultCategory(input)
	return &out, nil
}
