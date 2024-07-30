package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistoryCompatibilityMode string

const (
	BrowserSiteHistoryCompatibilityMode_Default                     BrowserSiteHistoryCompatibilityMode = "default"
	BrowserSiteHistoryCompatibilityMode_InternetExplorer10          BrowserSiteHistoryCompatibilityMode = "internetExplorer10"
	BrowserSiteHistoryCompatibilityMode_InternetExplorer11          BrowserSiteHistoryCompatibilityMode = "internetExplorer11"
	BrowserSiteHistoryCompatibilityMode_InternetExplorer5           BrowserSiteHistoryCompatibilityMode = "internetExplorer5"
	BrowserSiteHistoryCompatibilityMode_InternetExplorer7           BrowserSiteHistoryCompatibilityMode = "internetExplorer7"
	BrowserSiteHistoryCompatibilityMode_InternetExplorer7Enterprise BrowserSiteHistoryCompatibilityMode = "internetExplorer7Enterprise"
	BrowserSiteHistoryCompatibilityMode_InternetExplorer8           BrowserSiteHistoryCompatibilityMode = "internetExplorer8"
	BrowserSiteHistoryCompatibilityMode_InternetExplorer8Enterprise BrowserSiteHistoryCompatibilityMode = "internetExplorer8Enterprise"
	BrowserSiteHistoryCompatibilityMode_InternetExplorer9           BrowserSiteHistoryCompatibilityMode = "internetExplorer9"
)

func PossibleValuesForBrowserSiteHistoryCompatibilityMode() []string {
	return []string{
		string(BrowserSiteHistoryCompatibilityMode_Default),
		string(BrowserSiteHistoryCompatibilityMode_InternetExplorer10),
		string(BrowserSiteHistoryCompatibilityMode_InternetExplorer11),
		string(BrowserSiteHistoryCompatibilityMode_InternetExplorer5),
		string(BrowserSiteHistoryCompatibilityMode_InternetExplorer7),
		string(BrowserSiteHistoryCompatibilityMode_InternetExplorer7Enterprise),
		string(BrowserSiteHistoryCompatibilityMode_InternetExplorer8),
		string(BrowserSiteHistoryCompatibilityMode_InternetExplorer8Enterprise),
		string(BrowserSiteHistoryCompatibilityMode_InternetExplorer9),
	}
}

func (s *BrowserSiteHistoryCompatibilityMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSiteHistoryCompatibilityMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSiteHistoryCompatibilityMode(input string) (*BrowserSiteHistoryCompatibilityMode, error) {
	vals := map[string]BrowserSiteHistoryCompatibilityMode{
		"default":                     BrowserSiteHistoryCompatibilityMode_Default,
		"internetexplorer10":          BrowserSiteHistoryCompatibilityMode_InternetExplorer10,
		"internetexplorer11":          BrowserSiteHistoryCompatibilityMode_InternetExplorer11,
		"internetexplorer5":           BrowserSiteHistoryCompatibilityMode_InternetExplorer5,
		"internetexplorer7":           BrowserSiteHistoryCompatibilityMode_InternetExplorer7,
		"internetexplorer7enterprise": BrowserSiteHistoryCompatibilityMode_InternetExplorer7Enterprise,
		"internetexplorer8":           BrowserSiteHistoryCompatibilityMode_InternetExplorer8,
		"internetexplorer8enterprise": BrowserSiteHistoryCompatibilityMode_InternetExplorer8Enterprise,
		"internetexplorer9":           BrowserSiteHistoryCompatibilityMode_InternetExplorer9,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteHistoryCompatibilityMode(input)
	return &out, nil
}
