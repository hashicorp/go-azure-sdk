package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionCategory string

const (
	SecurityEmailUrlThreatSubmissionCategorymalware  SecurityEmailUrlThreatSubmissionCategory = "Malware"
	SecurityEmailUrlThreatSubmissionCategorynotJunk  SecurityEmailUrlThreatSubmissionCategory = "NotJunk"
	SecurityEmailUrlThreatSubmissionCategoryphishing SecurityEmailUrlThreatSubmissionCategory = "Phishing"
	SecurityEmailUrlThreatSubmissionCategoryspam     SecurityEmailUrlThreatSubmissionCategory = "Spam"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionCategory() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionCategorymalware),
		string(SecurityEmailUrlThreatSubmissionCategorynotJunk),
		string(SecurityEmailUrlThreatSubmissionCategoryphishing),
		string(SecurityEmailUrlThreatSubmissionCategoryspam),
	}
}

func parseSecurityEmailUrlThreatSubmissionCategory(input string) (*SecurityEmailUrlThreatSubmissionCategory, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionCategory{
		"malware":  SecurityEmailUrlThreatSubmissionCategorymalware,
		"notjunk":  SecurityEmailUrlThreatSubmissionCategorynotJunk,
		"phishing": SecurityEmailUrlThreatSubmissionCategoryphishing,
		"spam":     SecurityEmailUrlThreatSubmissionCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionCategory(input)
	return &out, nil
}
