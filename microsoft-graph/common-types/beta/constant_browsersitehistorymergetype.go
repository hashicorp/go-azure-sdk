package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistoryMergeType string

const (
	BrowserSiteHistoryMergeType_Default BrowserSiteHistoryMergeType = "default"
	BrowserSiteHistoryMergeType_NoMerge BrowserSiteHistoryMergeType = "noMerge"
)

func PossibleValuesForBrowserSiteHistoryMergeType() []string {
	return []string{
		string(BrowserSiteHistoryMergeType_Default),
		string(BrowserSiteHistoryMergeType_NoMerge),
	}
}

func (s *BrowserSiteHistoryMergeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBrowserSiteHistoryMergeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBrowserSiteHistoryMergeType(input string) (*BrowserSiteHistoryMergeType, error) {
	vals := map[string]BrowserSiteHistoryMergeType{
		"default": BrowserSiteHistoryMergeType_Default,
		"nomerge": BrowserSiteHistoryMergeType_NoMerge,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BrowserSiteHistoryMergeType(input)
	return &out, nil
}
