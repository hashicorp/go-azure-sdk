package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderDownloads string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderDownloads = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderDownloads = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_Show          Windows10GeneralConfigurationStartMenuPinnedFolderDownloads = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderDownloads() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderDownloads) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderDownloads(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderDownloads(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderDownloads, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderDownloads{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderDownloads_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderDownloads(input)
	return &out, nil
}
