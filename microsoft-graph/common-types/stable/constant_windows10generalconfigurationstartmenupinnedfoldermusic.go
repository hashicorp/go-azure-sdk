package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderMusic string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderMusic_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderMusic = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderMusic_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderMusic = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderMusic_Show          Windows10GeneralConfigurationStartMenuPinnedFolderMusic = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderMusic() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderMusic_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderMusic_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderMusic_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderMusic) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderMusic(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderMusic(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderMusic, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderMusic{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderMusic_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderMusic_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderMusic_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderMusic(input)
	return &out, nil
}
