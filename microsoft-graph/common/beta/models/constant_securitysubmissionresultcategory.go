package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionResultCategory string

const (
	SecuritySubmissionResultCategoryallowedByPolicy   SecuritySubmissionResultCategory = "AllowedByPolicy"
	SecuritySubmissionResultCategoryblockedByPolicy   SecuritySubmissionResultCategory = "BlockedByPolicy"
	SecuritySubmissionResultCategorymalware           SecuritySubmissionResultCategory = "Malware"
	SecuritySubmissionResultCategorynoResultAvailable SecuritySubmissionResultCategory = "NoResultAvailable"
	SecuritySubmissionResultCategorynotJunk           SecuritySubmissionResultCategory = "NotJunk"
	SecuritySubmissionResultCategoryphishing          SecuritySubmissionResultCategory = "Phishing"
	SecuritySubmissionResultCategoryspam              SecuritySubmissionResultCategory = "Spam"
	SecuritySubmissionResultCategoryspoof             SecuritySubmissionResultCategory = "Spoof"
	SecuritySubmissionResultCategoryunknown           SecuritySubmissionResultCategory = "Unknown"
)

func PossibleValuesForSecuritySubmissionResultCategory() []string {
	return []string{
		string(SecuritySubmissionResultCategoryallowedByPolicy),
		string(SecuritySubmissionResultCategoryblockedByPolicy),
		string(SecuritySubmissionResultCategorymalware),
		string(SecuritySubmissionResultCategorynoResultAvailable),
		string(SecuritySubmissionResultCategorynotJunk),
		string(SecuritySubmissionResultCategoryphishing),
		string(SecuritySubmissionResultCategoryspam),
		string(SecuritySubmissionResultCategoryspoof),
		string(SecuritySubmissionResultCategoryunknown),
	}
}

func parseSecuritySubmissionResultCategory(input string) (*SecuritySubmissionResultCategory, error) {
	vals := map[string]SecuritySubmissionResultCategory{
		"allowedbypolicy":   SecuritySubmissionResultCategoryallowedByPolicy,
		"blockedbypolicy":   SecuritySubmissionResultCategoryblockedByPolicy,
		"malware":           SecuritySubmissionResultCategorymalware,
		"noresultavailable": SecuritySubmissionResultCategorynoResultAvailable,
		"notjunk":           SecuritySubmissionResultCategorynotJunk,
		"phishing":          SecuritySubmissionResultCategoryphishing,
		"spam":              SecuritySubmissionResultCategoryspam,
		"spoof":             SecuritySubmissionResultCategoryspoof,
		"unknown":           SecuritySubmissionResultCategoryunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionResultCategory(input)
	return &out, nil
}
