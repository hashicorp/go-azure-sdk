package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading string

const (
	Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_Allowed       Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading = "allowed"
	Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_Blocked       Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading = "blocked"
	Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_NotConfigured Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationAppsAllowTrustedAppsSideloading() []string {
	return []string{
		string(Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_Allowed),
		string(Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_Blocked),
		string(Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationAppsAllowTrustedAppsSideloading(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationAppsAllowTrustedAppsSideloading(input string) (*Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading, error) {
	vals := map[string]Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading{
		"allowed":       Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_Allowed,
		"blocked":       Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_Blocked,
		"notconfigured": Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationAppsAllowTrustedAppsSideloading(input)
	return &out, nil
}
