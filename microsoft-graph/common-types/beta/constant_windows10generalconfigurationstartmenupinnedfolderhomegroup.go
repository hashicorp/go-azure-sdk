package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_Show          Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderHomeGroup() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderHomeGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderHomeGroup(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderHomeGroup(input)
	return &out, nil
}
