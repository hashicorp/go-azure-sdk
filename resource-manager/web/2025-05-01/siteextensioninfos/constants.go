package siteextensioninfos

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteExtensionType string

const (
	SiteExtensionTypeGallery SiteExtensionType = "Gallery"
	SiteExtensionTypeWebRoot SiteExtensionType = "WebRoot"
)

func PossibleValuesForSiteExtensionType() []string {
	return []string{
		string(SiteExtensionTypeGallery),
		string(SiteExtensionTypeWebRoot),
	}
}

func (s *SiteExtensionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSiteExtensionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSiteExtensionType(input string) (*SiteExtensionType, error) {
	vals := map[string]SiteExtensionType{
		"gallery": SiteExtensionTypeGallery,
		"webroot": SiteExtensionTypeWebRoot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SiteExtensionType(input)
	return &out, nil
}
