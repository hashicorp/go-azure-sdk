package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteMergeType string

const (
	BrowserSiteMergeTypedefault BrowserSiteMergeType = "Default"
	BrowserSiteMergeTypenoMerge BrowserSiteMergeType = "NoMerge"
)

func PossibleValuesForBrowserSiteMergeType() []string {
	return []string{
		string(BrowserSiteMergeTypedefault),
		string(BrowserSiteMergeTypenoMerge),
	}
}

func parseBrowserSiteMergeType(input string) (*BrowserSiteMergeType, error) {
	vals := map[string]BrowserSiteMergeType{
		"default": BrowserSiteMergeTypedefault,
		"nomerge": BrowserSiteMergeTypenoMerge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteMergeType(input)
	return &out, nil
}
