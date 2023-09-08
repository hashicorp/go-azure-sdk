package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionCategory string

const (
	SecurityEmailThreatSubmissionCategorymalware  SecurityEmailThreatSubmissionCategory = "Malware"
	SecurityEmailThreatSubmissionCategorynotJunk  SecurityEmailThreatSubmissionCategory = "NotJunk"
	SecurityEmailThreatSubmissionCategoryphishing SecurityEmailThreatSubmissionCategory = "Phishing"
	SecurityEmailThreatSubmissionCategoryspam     SecurityEmailThreatSubmissionCategory = "Spam"
)

func PossibleValuesForSecurityEmailThreatSubmissionCategory() []string {
	return []string{
		string(SecurityEmailThreatSubmissionCategorymalware),
		string(SecurityEmailThreatSubmissionCategorynotJunk),
		string(SecurityEmailThreatSubmissionCategoryphishing),
		string(SecurityEmailThreatSubmissionCategoryspam),
	}
}

func parseSecurityEmailThreatSubmissionCategory(input string) (*SecurityEmailThreatSubmissionCategory, error) {
	vals := map[string]SecurityEmailThreatSubmissionCategory{
		"malware":  SecurityEmailThreatSubmissionCategorymalware,
		"notjunk":  SecurityEmailThreatSubmissionCategorynotJunk,
		"phishing": SecurityEmailThreatSubmissionCategoryphishing,
		"spam":     SecurityEmailThreatSubmissionCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionCategory(input)
	return &out, nil
}
