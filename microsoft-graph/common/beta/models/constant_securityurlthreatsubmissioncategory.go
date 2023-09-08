package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionCategory string

const (
	SecurityUrlThreatSubmissionCategorymalware  SecurityUrlThreatSubmissionCategory = "Malware"
	SecurityUrlThreatSubmissionCategorynotJunk  SecurityUrlThreatSubmissionCategory = "NotJunk"
	SecurityUrlThreatSubmissionCategoryphishing SecurityUrlThreatSubmissionCategory = "Phishing"
	SecurityUrlThreatSubmissionCategoryspam     SecurityUrlThreatSubmissionCategory = "Spam"
)

func PossibleValuesForSecurityUrlThreatSubmissionCategory() []string {
	return []string{
		string(SecurityUrlThreatSubmissionCategorymalware),
		string(SecurityUrlThreatSubmissionCategorynotJunk),
		string(SecurityUrlThreatSubmissionCategoryphishing),
		string(SecurityUrlThreatSubmissionCategoryspam),
	}
}

func parseSecurityUrlThreatSubmissionCategory(input string) (*SecurityUrlThreatSubmissionCategory, error) {
	vals := map[string]SecurityUrlThreatSubmissionCategory{
		"malware":  SecurityUrlThreatSubmissionCategorymalware,
		"notjunk":  SecurityUrlThreatSubmissionCategorynotJunk,
		"phishing": SecurityUrlThreatSubmissionCategoryphishing,
		"spam":     SecurityUrlThreatSubmissionCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionCategory(input)
	return &out, nil
}
