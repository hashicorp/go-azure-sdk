package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteStatus string

const (
	BrowserSiteStatuspendingAdd    BrowserSiteStatus = "PendingAdd"
	BrowserSiteStatuspendingDelete BrowserSiteStatus = "PendingDelete"
	BrowserSiteStatuspendingEdit   BrowserSiteStatus = "PendingEdit"
	BrowserSiteStatuspublished     BrowserSiteStatus = "Published"
)

func PossibleValuesForBrowserSiteStatus() []string {
	return []string{
		string(BrowserSiteStatuspendingAdd),
		string(BrowserSiteStatuspendingDelete),
		string(BrowserSiteStatuspendingEdit),
		string(BrowserSiteStatuspublished),
	}
}

func parseBrowserSiteStatus(input string) (*BrowserSiteStatus, error) {
	vals := map[string]BrowserSiteStatus{
		"pendingadd":    BrowserSiteStatuspendingAdd,
		"pendingdelete": BrowserSiteStatuspendingDelete,
		"pendingedit":   BrowserSiteStatuspendingEdit,
		"published":     BrowserSiteStatuspublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteStatus(input)
	return &out, nil
}
