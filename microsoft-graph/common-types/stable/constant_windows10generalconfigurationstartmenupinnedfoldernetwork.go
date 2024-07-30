package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderNetwork string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderNetwork = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderNetwork = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_Show          Windows10GeneralConfigurationStartMenuPinnedFolderNetwork = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderNetwork() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderNetwork) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderNetwork(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderNetwork(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderNetwork, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderNetwork{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderNetwork_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderNetwork(input)
	return &out, nil
}
