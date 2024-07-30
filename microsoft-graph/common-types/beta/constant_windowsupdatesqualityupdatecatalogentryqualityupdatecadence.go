package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence string

const (
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence_Monthly   WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence = "monthly"
	WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence_OutOfBand WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence = "outOfBand"
)

func PossibleValuesForWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence() []string {
	return []string{
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence_Monthly),
		string(WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence_OutOfBand),
	}
}

func (s *WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence(input string) (*WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence, error) {
	vals := map[string]WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence{
		"monthly":   WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence_Monthly,
		"outofband": WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence_OutOfBand,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesQualityUpdateCatalogEntryQualityUpdateCadence(input)
	return &out, nil
}
