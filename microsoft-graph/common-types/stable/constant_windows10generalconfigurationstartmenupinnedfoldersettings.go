package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderSettings string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderSettings_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderSettings = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderSettings_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderSettings = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderSettings_Show          Windows10GeneralConfigurationStartMenuPinnedFolderSettings = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderSettings() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderSettings_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderSettings_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderSettings_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderSettings(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderSettings, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderSettings{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderSettings_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderSettings_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderSettings_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderSettings(input)
	return &out, nil
}
