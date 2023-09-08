package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionCategory string

const (
	SecurityFileContentThreatSubmissionCategorymalware  SecurityFileContentThreatSubmissionCategory = "Malware"
	SecurityFileContentThreatSubmissionCategorynotJunk  SecurityFileContentThreatSubmissionCategory = "NotJunk"
	SecurityFileContentThreatSubmissionCategoryphishing SecurityFileContentThreatSubmissionCategory = "Phishing"
	SecurityFileContentThreatSubmissionCategoryspam     SecurityFileContentThreatSubmissionCategory = "Spam"
)

func PossibleValuesForSecurityFileContentThreatSubmissionCategory() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionCategorymalware),
		string(SecurityFileContentThreatSubmissionCategorynotJunk),
		string(SecurityFileContentThreatSubmissionCategoryphishing),
		string(SecurityFileContentThreatSubmissionCategoryspam),
	}
}

func parseSecurityFileContentThreatSubmissionCategory(input string) (*SecurityFileContentThreatSubmissionCategory, error) {
	vals := map[string]SecurityFileContentThreatSubmissionCategory{
		"malware":  SecurityFileContentThreatSubmissionCategorymalware,
		"notjunk":  SecurityFileContentThreatSubmissionCategorynotJunk,
		"phishing": SecurityFileContentThreatSubmissionCategoryphishing,
		"spam":     SecurityFileContentThreatSubmissionCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionCategory(input)
	return &out, nil
}
