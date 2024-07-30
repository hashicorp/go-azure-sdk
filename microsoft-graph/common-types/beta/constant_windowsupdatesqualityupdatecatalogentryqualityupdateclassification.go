package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification string

const (
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_All         WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification = "all"
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_NonSecurity WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification = "nonSecurity"
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_Security    WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification = "security"
)

func PossibleValuesForWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification() []string {
	return []string{
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_All),
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_NonSecurity),
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_Security),
	}
}

func (s *WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification(input string) (*WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification, error) {
	vals := map[string]WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification{
		"all":         WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_All,
		"nonsecurity": WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_NonSecurity,
		"security":    WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification_Security,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateClassification(input)
	return &out, nil
}
