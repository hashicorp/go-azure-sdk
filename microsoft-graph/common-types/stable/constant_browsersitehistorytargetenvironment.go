package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistoryTargetEnvironment string

const (
	BrowserSiteHistoryTargetEnvironment_Configurable         BrowserSiteHistoryTargetEnvironment = "configurable"
	BrowserSiteHistoryTargetEnvironment_InternetExplorer11   BrowserSiteHistoryTargetEnvironment = "internetExplorer11"
	BrowserSiteHistoryTargetEnvironment_InternetExplorerMode BrowserSiteHistoryTargetEnvironment = "internetExplorerMode"
	BrowserSiteHistoryTargetEnvironment_MicrosoftEdge        BrowserSiteHistoryTargetEnvironment = "microsoftEdge"
	BrowserSiteHistoryTargetEnvironment_None                 BrowserSiteHistoryTargetEnvironment = "none"
)

func PossibleValuesForBrowserSiteHistoryTargetEnvironment() []string {
	return []string{
		string(BrowserSiteHistoryTargetEnvironment_Configurable),
		string(BrowserSiteHistoryTargetEnvironment_InternetExplorer11),
		string(BrowserSiteHistoryTargetEnvironment_InternetExplorerMode),
		string(BrowserSiteHistoryTargetEnvironment_MicrosoftEdge),
		string(BrowserSiteHistoryTargetEnvironment_None),
	}
}

func (s *BrowserSiteHistoryTargetEnvironment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSiteHistoryTargetEnvironment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSiteHistoryTargetEnvironment(input string) (*BrowserSiteHistoryTargetEnvironment, error) {
	vals := map[string]BrowserSiteHistoryTargetEnvironment{
		"configurable":         BrowserSiteHistoryTargetEnvironment_Configurable,
		"internetexplorer11":   BrowserSiteHistoryTargetEnvironment_InternetExplorer11,
		"internetexplorermode": BrowserSiteHistoryTargetEnvironment_InternetExplorerMode,
		"microsoftedge":        BrowserSiteHistoryTargetEnvironment_MicrosoftEdge,
		"none":                 BrowserSiteHistoryTargetEnvironment_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteHistoryTargetEnvironment(input)
	return &out, nil
}
