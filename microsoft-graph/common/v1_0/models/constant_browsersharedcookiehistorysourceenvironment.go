package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieHistorySourceEnvironment string

const (
	BrowserSharedCookieHistorySourceEnvironmentboth               BrowserSharedCookieHistorySourceEnvironment = "Both"
	BrowserSharedCookieHistorySourceEnvironmentinternetExplorer11 BrowserSharedCookieHistorySourceEnvironment = "InternetExplorer11"
	BrowserSharedCookieHistorySourceEnvironmentmicrosoftEdge      BrowserSharedCookieHistorySourceEnvironment = "MicrosoftEdge"
)

func PossibleValuesForBrowserSharedCookieHistorySourceEnvironment() []string {
	return []string{
		string(BrowserSharedCookieHistorySourceEnvironmentboth),
		string(BrowserSharedCookieHistorySourceEnvironmentinternetExplorer11),
		string(BrowserSharedCookieHistorySourceEnvironmentmicrosoftEdge),
	}
}

func parseBrowserSharedCookieHistorySourceEnvironment(input string) (*BrowserSharedCookieHistorySourceEnvironment, error) {
	vals := map[string]BrowserSharedCookieHistorySourceEnvironment{
		"both":               BrowserSharedCookieHistorySourceEnvironmentboth,
		"internetexplorer11": BrowserSharedCookieHistorySourceEnvironmentinternetExplorer11,
		"microsoftedge":      BrowserSharedCookieHistorySourceEnvironmentmicrosoftEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSharedCookieHistorySourceEnvironment(input)
	return &out, nil
}
