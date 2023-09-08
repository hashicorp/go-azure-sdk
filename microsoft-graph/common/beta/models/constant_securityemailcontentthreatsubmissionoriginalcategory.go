package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionOriginalCategory string

const (
	SecurityEmailContentThreatSubmissionOriginalCategorymalware  SecurityEmailContentThreatSubmissionOriginalCategory = "Malware"
	SecurityEmailContentThreatSubmissionOriginalCategorynotJunk  SecurityEmailContentThreatSubmissionOriginalCategory = "NotJunk"
	SecurityEmailContentThreatSubmissionOriginalCategoryphishing SecurityEmailContentThreatSubmissionOriginalCategory = "Phishing"
	SecurityEmailContentThreatSubmissionOriginalCategoryspam     SecurityEmailContentThreatSubmissionOriginalCategory = "Spam"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionOriginalCategory() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionOriginalCategorymalware),
		string(SecurityEmailContentThreatSubmissionOriginalCategorynotJunk),
		string(SecurityEmailContentThreatSubmissionOriginalCategoryphishing),
		string(SecurityEmailContentThreatSubmissionOriginalCategoryspam),
	}
}

func parseSecurityEmailContentThreatSubmissionOriginalCategory(input string) (*SecurityEmailContentThreatSubmissionOriginalCategory, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionOriginalCategory{
		"malware":  SecurityEmailContentThreatSubmissionOriginalCategorymalware,
		"notjunk":  SecurityEmailContentThreatSubmissionOriginalCategorynotJunk,
		"phishing": SecurityEmailContentThreatSubmissionOriginalCategoryphishing,
		"spam":     SecurityEmailContentThreatSubmissionOriginalCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionOriginalCategory(input)
	return &out, nil
}
