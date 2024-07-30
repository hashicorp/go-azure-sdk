package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderPictures string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderPictures_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderPictures = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderPictures_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderPictures = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderPictures_Show          Windows10GeneralConfigurationStartMenuPinnedFolderPictures = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderPictures() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPictures_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPictures_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderPictures_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderPictures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderPictures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderPictures(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderPictures, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderPictures{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderPictures_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderPictures_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderPictures_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderPictures(input)
	return &out, nil
}
