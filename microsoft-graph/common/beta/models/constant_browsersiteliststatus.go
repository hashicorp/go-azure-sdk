package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteListStatus string

const (
	BrowserSiteListStatusdraft     BrowserSiteListStatus = "Draft"
	BrowserSiteListStatuspending   BrowserSiteListStatus = "Pending"
	BrowserSiteListStatuspublished BrowserSiteListStatus = "Published"
)

func PossibleValuesForBrowserSiteListStatus() []string {
	return []string{
		string(BrowserSiteListStatusdraft),
		string(BrowserSiteListStatuspending),
		string(BrowserSiteListStatuspublished),
	}
}

func parseBrowserSiteListStatus(input string) (*BrowserSiteListStatus, error) {
	vals := map[string]BrowserSiteListStatus{
		"draft":     BrowserSiteListStatusdraft,
		"pending":   BrowserSiteListStatuspending,
		"published": BrowserSiteListStatuspublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteListStatus(input)
	return &out, nil
}
