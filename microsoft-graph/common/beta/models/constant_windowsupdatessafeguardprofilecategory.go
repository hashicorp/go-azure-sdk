package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSafeguardProfileCategory string

const (
	WindowsUpdatesSafeguardProfileCategorylikelyIssues WindowsUpdatesSafeguardProfileCategory = "LikelyIssues"
)

func PossibleValuesForWindowsUpdatesSafeguardProfileCategory() []string {
	return []string{
		string(WindowsUpdatesSafeguardProfileCategorylikelyIssues),
	}
}

func parseWindowsUpdatesSafeguardProfileCategory(input string) (*WindowsUpdatesSafeguardProfileCategory, error) {
	vals := map[string]WindowsUpdatesSafeguardProfileCategory{
		"likelyissues": WindowsUpdatesSafeguardProfileCategorylikelyIssues,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesSafeguardProfileCategory(input)
	return &out, nil
}
