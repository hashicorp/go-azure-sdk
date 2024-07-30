package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_Show          Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderPersonalFolder(input)
	return &out, nil
}
