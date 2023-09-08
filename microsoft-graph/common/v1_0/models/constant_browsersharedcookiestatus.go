package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieStatus string

const (
	BrowserSharedCookieStatuspendingAdd    BrowserSharedCookieStatus = "PendingAdd"
	BrowserSharedCookieStatuspendingDelete BrowserSharedCookieStatus = "PendingDelete"
	BrowserSharedCookieStatuspendingEdit   BrowserSharedCookieStatus = "PendingEdit"
	BrowserSharedCookieStatuspublished     BrowserSharedCookieStatus = "Published"
)

func PossibleValuesForBrowserSharedCookieStatus() []string {
	return []string{
		string(BrowserSharedCookieStatuspendingAdd),
		string(BrowserSharedCookieStatuspendingDelete),
		string(BrowserSharedCookieStatuspendingEdit),
		string(BrowserSharedCookieStatuspublished),
	}
}

func parseBrowserSharedCookieStatus(input string) (*BrowserSharedCookieStatus, error) {
	vals := map[string]BrowserSharedCookieStatus{
		"pendingadd":    BrowserSharedCookieStatuspendingAdd,
		"pendingdelete": BrowserSharedCookieStatuspendingDelete,
		"pendingedit":   BrowserSharedCookieStatuspendingEdit,
		"published":     BrowserSharedCookieStatuspublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSharedCookieStatus(input)
	return &out, nil
}
