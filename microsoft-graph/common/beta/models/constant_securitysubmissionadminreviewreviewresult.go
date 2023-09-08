package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionAdminReviewReviewResult string

const (
	SecuritySubmissionAdminReviewReviewResultallowedByPolicy   SecuritySubmissionAdminReviewReviewResult = "AllowedByPolicy"
	SecuritySubmissionAdminReviewReviewResultblockedByPolicy   SecuritySubmissionAdminReviewReviewResult = "BlockedByPolicy"
	SecuritySubmissionAdminReviewReviewResultmalware           SecuritySubmissionAdminReviewReviewResult = "Malware"
	SecuritySubmissionAdminReviewReviewResultnoResultAvailable SecuritySubmissionAdminReviewReviewResult = "NoResultAvailable"
	SecuritySubmissionAdminReviewReviewResultnotJunk           SecuritySubmissionAdminReviewReviewResult = "NotJunk"
	SecuritySubmissionAdminReviewReviewResultphishing          SecuritySubmissionAdminReviewReviewResult = "Phishing"
	SecuritySubmissionAdminReviewReviewResultspam              SecuritySubmissionAdminReviewReviewResult = "Spam"
	SecuritySubmissionAdminReviewReviewResultspoof             SecuritySubmissionAdminReviewReviewResult = "Spoof"
	SecuritySubmissionAdminReviewReviewResultunknown           SecuritySubmissionAdminReviewReviewResult = "Unknown"
)

func PossibleValuesForSecuritySubmissionAdminReviewReviewResult() []string {
	return []string{
		string(SecuritySubmissionAdminReviewReviewResultallowedByPolicy),
		string(SecuritySubmissionAdminReviewReviewResultblockedByPolicy),
		string(SecuritySubmissionAdminReviewReviewResultmalware),
		string(SecuritySubmissionAdminReviewReviewResultnoResultAvailable),
		string(SecuritySubmissionAdminReviewReviewResultnotJunk),
		string(SecuritySubmissionAdminReviewReviewResultphishing),
		string(SecuritySubmissionAdminReviewReviewResultspam),
		string(SecuritySubmissionAdminReviewReviewResultspoof),
		string(SecuritySubmissionAdminReviewReviewResultunknown),
	}
}

func parseSecuritySubmissionAdminReviewReviewResult(input string) (*SecuritySubmissionAdminReviewReviewResult, error) {
	vals := map[string]SecuritySubmissionAdminReviewReviewResult{
		"allowedbypolicy":   SecuritySubmissionAdminReviewReviewResultallowedByPolicy,
		"blockedbypolicy":   SecuritySubmissionAdminReviewReviewResultblockedByPolicy,
		"malware":           SecuritySubmissionAdminReviewReviewResultmalware,
		"noresultavailable": SecuritySubmissionAdminReviewReviewResultnoResultAvailable,
		"notjunk":           SecuritySubmissionAdminReviewReviewResultnotJunk,
		"phishing":          SecuritySubmissionAdminReviewReviewResultphishing,
		"spam":              SecuritySubmissionAdminReviewReviewResultspam,
		"spoof":             SecuritySubmissionAdminReviewReviewResultspoof,
		"unknown":           SecuritySubmissionAdminReviewReviewResultunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionAdminReviewReviewResult(input)
	return &out, nil
}
