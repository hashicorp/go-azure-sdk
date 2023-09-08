package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmissionCategory string

const (
	SecurityFileThreatSubmissionCategorymalware  SecurityFileThreatSubmissionCategory = "Malware"
	SecurityFileThreatSubmissionCategorynotJunk  SecurityFileThreatSubmissionCategory = "NotJunk"
	SecurityFileThreatSubmissionCategoryphishing SecurityFileThreatSubmissionCategory = "Phishing"
	SecurityFileThreatSubmissionCategoryspam     SecurityFileThreatSubmissionCategory = "Spam"
)

func PossibleValuesForSecurityFileThreatSubmissionCategory() []string {
	return []string{
		string(SecurityFileThreatSubmissionCategorymalware),
		string(SecurityFileThreatSubmissionCategorynotJunk),
		string(SecurityFileThreatSubmissionCategoryphishing),
		string(SecurityFileThreatSubmissionCategoryspam),
	}
}

func parseSecurityFileThreatSubmissionCategory(input string) (*SecurityFileThreatSubmissionCategory, error) {
	vals := map[string]SecurityFileThreatSubmissionCategory{
		"malware":  SecurityFileThreatSubmissionCategorymalware,
		"notjunk":  SecurityFileThreatSubmissionCategorynotJunk,
		"phishing": SecurityFileThreatSubmissionCategoryphishing,
		"spam":     SecurityFileThreatSubmissionCategoryspam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileThreatSubmissionCategory(input)
	return &out, nil
}
