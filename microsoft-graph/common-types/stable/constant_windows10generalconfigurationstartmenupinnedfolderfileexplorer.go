package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_Show          Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderFileExplorer() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderFileExplorer(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderFileExplorer(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderFileExplorer(input)
	return &out, nil
}
