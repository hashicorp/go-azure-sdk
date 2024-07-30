package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionAdminReviewReviewResult string

const (
	SecuritySubmissionAdminReviewReviewResult_AllowedByPolicy   SecuritySubmissionAdminReviewReviewResult = "allowedByPolicy"
	SecuritySubmissionAdminReviewReviewResult_BlockedByPolicy   SecuritySubmissionAdminReviewReviewResult = "blockedByPolicy"
	SecuritySubmissionAdminReviewReviewResult_Malware           SecuritySubmissionAdminReviewReviewResult = "malware"
	SecuritySubmissionAdminReviewReviewResult_NoResultAvailable SecuritySubmissionAdminReviewReviewResult = "noResultAvailable"
	SecuritySubmissionAdminReviewReviewResult_NotJunk           SecuritySubmissionAdminReviewReviewResult = "notJunk"
	SecuritySubmissionAdminReviewReviewResult_Phishing          SecuritySubmissionAdminReviewReviewResult = "phishing"
	SecuritySubmissionAdminReviewReviewResult_Spam              SecuritySubmissionAdminReviewReviewResult = "spam"
	SecuritySubmissionAdminReviewReviewResult_Spoof             SecuritySubmissionAdminReviewReviewResult = "spoof"
	SecuritySubmissionAdminReviewReviewResult_Unknown           SecuritySubmissionAdminReviewReviewResult = "unknown"
)

func PossibleValuesForSecuritySubmissionAdminReviewReviewResult() []string {
	return []string{
		string(SecuritySubmissionAdminReviewReviewResult_AllowedByPolicy),
		string(SecuritySubmissionAdminReviewReviewResult_BlockedByPolicy),
		string(SecuritySubmissionAdminReviewReviewResult_Malware),
		string(SecuritySubmissionAdminReviewReviewResult_NoResultAvailable),
		string(SecuritySubmissionAdminReviewReviewResult_NotJunk),
		string(SecuritySubmissionAdminReviewReviewResult_Phishing),
		string(SecuritySubmissionAdminReviewReviewResult_Spam),
		string(SecuritySubmissionAdminReviewReviewResult_Spoof),
		string(SecuritySubmissionAdminReviewReviewResult_Unknown),
	}
}

func (s *SecuritySubmissionAdminReviewReviewResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionAdminReviewReviewResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionAdminReviewReviewResult(input string) (*SecuritySubmissionAdminReviewReviewResult, error) {
	vals := map[string]SecuritySubmissionAdminReviewReviewResult{
		"allowedbypolicy":   SecuritySubmissionAdminReviewReviewResult_AllowedByPolicy,
		"blockedbypolicy":   SecuritySubmissionAdminReviewReviewResult_BlockedByPolicy,
		"malware":           SecuritySubmissionAdminReviewReviewResult_Malware,
		"noresultavailable": SecuritySubmissionAdminReviewReviewResult_NoResultAvailable,
		"notjunk":           SecuritySubmissionAdminReviewReviewResult_NotJunk,
		"phishing":          SecuritySubmissionAdminReviewReviewResult_Phishing,
		"spam":              SecuritySubmissionAdminReviewReviewResult_Spam,
		"spoof":             SecuritySubmissionAdminReviewReviewResult_Spoof,
		"unknown":           SecuritySubmissionAdminReviewReviewResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionAdminReviewReviewResult(input)
	return &out, nil
}
