package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderVideos string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderVideos_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderVideos = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderVideos_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderVideos = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderVideos_Show          Windows10GeneralConfigurationStartMenuPinnedFolderVideos = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderVideos() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderVideos_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderVideos_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderVideos_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderVideos) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderVideos(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderVideos(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderVideos, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderVideos{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderVideos_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderVideos_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderVideos_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderVideos(input)
	return &out, nil
}
