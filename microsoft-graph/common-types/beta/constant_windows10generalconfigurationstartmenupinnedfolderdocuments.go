package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderDocuments string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_Hide          Windows10GeneralConfigurationStartMenuPinnedFolderDocuments = "hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_NotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderDocuments = "notConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_Show          Windows10GeneralConfigurationStartMenuPinnedFolderDocuments = "show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderDocuments() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_Hide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_NotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_Show),
	}
}

func (s *Windows10GeneralConfigurationStartMenuPinnedFolderDocuments) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuPinnedFolderDocuments(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderDocuments(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderDocuments, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderDocuments{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_Hide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_NotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderDocuments_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderDocuments(input)
	return &out, nil
}
