package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionCategory string

const (
	SecurityEmailContentThreatSubmissionCategorymalware  SecurityEmailContentThreatSubmissionCategory = "Malware"
	SecurityEmailContentThreatSubmissionCategorynotJunk  SecurityEmailContentThreatSubmissionCategory = "NotJunk"
	SecurityEmailContentThreatSubmissionCategoryphishing SecurityEmailContentThreatSubmissionCategory = "Phishing"
	SecurityEmailContentThreatSubmissionCategoryspam     SecurityEmailContentThreatSubmissionCategory = "Spam"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionCategory() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionCategorymalware),
		string(SecurityEmailContentThreatSubmissionCategorynotJunk),
		string(SecurityEmailContentThreatSubmissionCategoryphishing),
		string(SecurityEmailContentThreatSubmissionCategoryspam),
	}
}

func parseSecurityEmailContentThreatSubmissionCategory(input string) (*SecurityEmailContentThreatSubmissionCategory, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionCategory{
		"malware":  SecurityEmailContentThreatSubmissionCategorymalware,
		"notjunk":  SecurityEmailContentThreatSubmissionCategorynotJunk,
		"phishing": SecurityEmailContentThreatSubmissionCategoryphishing,
		"spam":     SecurityEmailContentThreatSubmissionCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionCategory(input)
	return &out, nil
}
