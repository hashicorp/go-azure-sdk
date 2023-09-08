package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionOriginalCategory string

const (
	SecurityEmailUrlThreatSubmissionOriginalCategorymalware  SecurityEmailUrlThreatSubmissionOriginalCategory = "Malware"
	SecurityEmailUrlThreatSubmissionOriginalCategorynotJunk  SecurityEmailUrlThreatSubmissionOriginalCategory = "NotJunk"
	SecurityEmailUrlThreatSubmissionOriginalCategoryphishing SecurityEmailUrlThreatSubmissionOriginalCategory = "Phishing"
	SecurityEmailUrlThreatSubmissionOriginalCategoryspam     SecurityEmailUrlThreatSubmissionOriginalCategory = "Spam"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionOriginalCategory() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionOriginalCategorymalware),
		string(SecurityEmailUrlThreatSubmissionOriginalCategorynotJunk),
		string(SecurityEmailUrlThreatSubmissionOriginalCategoryphishing),
		string(SecurityEmailUrlThreatSubmissionOriginalCategoryspam),
	}
}

func parseSecurityEmailUrlThreatSubmissionOriginalCategory(input string) (*SecurityEmailUrlThreatSubmissionOriginalCategory, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionOriginalCategory{
		"malware":  SecurityEmailUrlThreatSubmissionOriginalCategorymalware,
		"notjunk":  SecurityEmailUrlThreatSubmissionOriginalCategorynotJunk,
		"phishing": SecurityEmailUrlThreatSubmissionOriginalCategoryphishing,
		"spam":     SecurityEmailUrlThreatSubmissionOriginalCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionOriginalCategory(input)
	return &out, nil
}
