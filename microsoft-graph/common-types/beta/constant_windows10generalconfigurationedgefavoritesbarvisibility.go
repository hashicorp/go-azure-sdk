package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeFavoritesBarVisibility string

const (
	Windows10GeneralConfigurationEdgeFavoritesBarVisibility_Hide          Windows10GeneralConfigurationEdgeFavoritesBarVisibility = "hide"
	Windows10GeneralConfigurationEdgeFavoritesBarVisibility_NotConfigured Windows10GeneralConfigurationEdgeFavoritesBarVisibility = "notConfigured"
	Windows10GeneralConfigurationEdgeFavoritesBarVisibility_Show          Windows10GeneralConfigurationEdgeFavoritesBarVisibility = "show"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeFavoritesBarVisibility() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeFavoritesBarVisibility_Hide),
		string(Windows10GeneralConfigurationEdgeFavoritesBarVisibility_NotConfigured),
		string(Windows10GeneralConfigurationEdgeFavoritesBarVisibility_Show),
	}
}

func (s *Windows10GeneralConfigurationEdgeFavoritesBarVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationEdgeFavoritesBarVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationEdgeFavoritesBarVisibility(input string) (*Windows10GeneralConfigurationEdgeFavoritesBarVisibility, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeFavoritesBarVisibility{
		"hide":          Windows10GeneralConfigurationEdgeFavoritesBarVisibility_Hide,
		"notconfigured": Windows10GeneralConfigurationEdgeFavoritesBarVisibility_NotConfigured,
		"show":          Windows10GeneralConfigurationEdgeFavoritesBarVisibility_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeFavoritesBarVisibility(input)
	return &out, nil
}
