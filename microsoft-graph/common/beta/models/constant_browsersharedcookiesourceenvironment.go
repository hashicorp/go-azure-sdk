package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieSourceEnvironment string

const (
	BrowserSharedCookieSourceEnvironmentboth               BrowserSharedCookieSourceEnvironment = "Both"
	BrowserSharedCookieSourceEnvironmentinternetExplorer11 BrowserSharedCookieSourceEnvironment = "InternetExplorer11"
	BrowserSharedCookieSourceEnvironmentmicrosoftEdge      BrowserSharedCookieSourceEnvironment = "MicrosoftEdge"
)

func PossibleValuesForBrowserSharedCookieSourceEnvironment() []string {
	return []string{
		string(BrowserSharedCookieSourceEnvironmentboth),
		string(BrowserSharedCookieSourceEnvironmentinternetExplorer11),
		string(BrowserSharedCookieSourceEnvironmentmicrosoftEdge),
	}
}

func parseBrowserSharedCookieSourceEnvironment(input string) (*BrowserSharedCookieSourceEnvironment, error) {
	vals := map[string]BrowserSharedCookieSourceEnvironment{
		"both":               BrowserSharedCookieSourceEnvironmentboth,
		"internetexplorer11": BrowserSharedCookieSourceEnvironmentinternetExplorer11,
		"microsoftedge":      BrowserSharedCookieSourceEnvironmentmicrosoftEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSharedCookieSourceEnvironment(input)
	return &out, nil
}
