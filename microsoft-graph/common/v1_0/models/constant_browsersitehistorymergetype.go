package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistoryMergeType string

const (
	BrowserSiteHistoryMergeTypedefault BrowserSiteHistoryMergeType = "Default"
	BrowserSiteHistoryMergeTypenoMerge BrowserSiteHistoryMergeType = "NoMerge"
)

func PossibleValuesForBrowserSiteHistoryMergeType() []string {
	return []string{
		string(BrowserSiteHistoryMergeTypedefault),
		string(BrowserSiteHistoryMergeTypenoMerge),
	}
}

func parseBrowserSiteHistoryMergeType(input string) (*BrowserSiteHistoryMergeType, error) {
	vals := map[string]BrowserSiteHistoryMergeType{
		"default": BrowserSiteHistoryMergeTypedefault,
		"nomerge": BrowserSiteHistoryMergeTypenoMerge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteHistoryMergeType(input)
	return &out, nil
}
