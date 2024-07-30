package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieHistorySourceEnvironment string

const (
	BrowserSharedCookieHistorySourceEnvironment_Both               BrowserSharedCookieHistorySourceEnvironment = "both"
	BrowserSharedCookieHistorySourceEnvironment_InternetExplorer11 BrowserSharedCookieHistorySourceEnvironment = "internetExplorer11"
	BrowserSharedCookieHistorySourceEnvironment_MicrosoftEdge      BrowserSharedCookieHistorySourceEnvironment = "microsoftEdge"
)

func PossibleValuesForBrowserSharedCookieHistorySourceEnvironment() []string {
	return []string{
		string(BrowserSharedCookieHistorySourceEnvironment_Both),
		string(BrowserSharedCookieHistorySourceEnvironment_InternetExplorer11),
		string(BrowserSharedCookieHistorySourceEnvironment_MicrosoftEdge),
	}
}

func (s *BrowserSharedCookieHistorySourceEnvironment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSharedCookieHistorySourceEnvironment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSharedCookieHistorySourceEnvironment(input string) (*BrowserSharedCookieHistorySourceEnvironment, error) {
	vals := map[string]BrowserSharedCookieHistorySourceEnvironment{
		"both":               BrowserSharedCookieHistorySourceEnvironment_Both,
		"internetexplorer11": BrowserSharedCookieHistorySourceEnvironment_InternetExplorer11,
		"microsoftedge":      BrowserSharedCookieHistorySourceEnvironment_MicrosoftEdge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSharedCookieHistorySourceEnvironment(input)
	return &out, nil
}
