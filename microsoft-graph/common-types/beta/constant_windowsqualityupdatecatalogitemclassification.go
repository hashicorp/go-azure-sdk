package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateCatalogItemClassification string

const (
	WindowsQualityUpdateCatalogItemClassification_All         WindowsQualityUpdateCatalogItemClassification = "all"
	WindowsQualityUpdateCatalogItemClassification_NonSecurity WindowsQualityUpdateCatalogItemClassification = "nonSecurity"
	WindowsQualityUpdateCatalogItemClassification_Security    WindowsQualityUpdateCatalogItemClassification = "security"
)

func PossibleValuesForWindowsQualityUpdateCatalogItemClassification() []string {
	return []string{
		string(WindowsQualityUpdateCatalogItemClassification_All),
		string(WindowsQualityUpdateCatalogItemClassification_NonSecurity),
		string(WindowsQualityUpdateCatalogItemClassification_Security),
	}
}

func (s *WindowsQualityUpdateCatalogItemClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsQualityUpdateCatalogItemClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsQualityUpdateCatalogItemClassification(input string) (*WindowsQualityUpdateCatalogItemClassification, error) {
	vals := map[string]WindowsQualityUpdateCatalogItemClassification{
		"all":         WindowsQualityUpdateCatalogItemClassification_All,
		"nonsecurity": WindowsQualityUpdateCatalogItemClassification_NonSecurity,
		"security":    WindowsQualityUpdateCatalogItemClassification_Security,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsQualityUpdateCatalogItemClassification(input)
	return &out, nil
}
