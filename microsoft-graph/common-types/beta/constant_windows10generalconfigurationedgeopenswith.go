package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeOpensWith string

const (
	Windows10GeneralConfigurationEdgeOpensWith_NewTabPage    Windows10GeneralConfigurationEdgeOpensWith = "newTabPage"
	Windows10GeneralConfigurationEdgeOpensWith_NotConfigured Windows10GeneralConfigurationEdgeOpensWith = "notConfigured"
	Windows10GeneralConfigurationEdgeOpensWith_PreviousPages Windows10GeneralConfigurationEdgeOpensWith = "previousPages"
	Windows10GeneralConfigurationEdgeOpensWith_SpecificPages Windows10GeneralConfigurationEdgeOpensWith = "specificPages"
	Windows10GeneralConfigurationEdgeOpensWith_StartPage     Windows10GeneralConfigurationEdgeOpensWith = "startPage"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeOpensWith() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeOpensWith_NewTabPage),
		string(Windows10GeneralConfigurationEdgeOpensWith_NotConfigured),
		string(Windows10GeneralConfigurationEdgeOpensWith_PreviousPages),
		string(Windows10GeneralConfigurationEdgeOpensWith_SpecificPages),
		string(Windows10GeneralConfigurationEdgeOpensWith_StartPage),
	}
}

func (s *Windows10GeneralConfigurationEdgeOpensWith) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationEdgeOpensWith(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationEdgeOpensWith(input string) (*Windows10GeneralConfigurationEdgeOpensWith, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeOpensWith{
		"newtabpage":    Windows10GeneralConfigurationEdgeOpensWith_NewTabPage,
		"notconfigured": Windows10GeneralConfigurationEdgeOpensWith_NotConfigured,
		"previouspages": Windows10GeneralConfigurationEdgeOpensWith_PreviousPages,
		"specificpages": Windows10GeneralConfigurationEdgeOpensWith_SpecificPages,
		"startpage":     Windows10GeneralConfigurationEdgeOpensWith_StartPage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeOpensWith(input)
	return &out, nil
}
