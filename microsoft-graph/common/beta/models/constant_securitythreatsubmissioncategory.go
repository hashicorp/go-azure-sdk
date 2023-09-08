package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionCategory string

const (
	SecurityThreatSubmissionCategorymalware  SecurityThreatSubmissionCategory = "Malware"
	SecurityThreatSubmissionCategorynotJunk  SecurityThreatSubmissionCategory = "NotJunk"
	SecurityThreatSubmissionCategoryphishing SecurityThreatSubmissionCategory = "Phishing"
	SecurityThreatSubmissionCategoryspam     SecurityThreatSubmissionCategory = "Spam"
)

func PossibleValuesForSecurityThreatSubmissionCategory() []string {
	return []string{
		string(SecurityThreatSubmissionCategorymalware),
		string(SecurityThreatSubmissionCategorynotJunk),
		string(SecurityThreatSubmissionCategoryphishing),
		string(SecurityThreatSubmissionCategoryspam),
	}
}

func parseSecurityThreatSubmissionCategory(input string) (*SecurityThreatSubmissionCategory, error) {
	vals := map[string]SecurityThreatSubmissionCategory{
		"malware":  SecurityThreatSubmissionCategorymalware,
		"notjunk":  SecurityThreatSubmissionCategorynotJunk,
		"phishing": SecurityThreatSubmissionCategoryphishing,
		"spam":     SecurityThreatSubmissionCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionCategory(input)
	return &out, nil
}
