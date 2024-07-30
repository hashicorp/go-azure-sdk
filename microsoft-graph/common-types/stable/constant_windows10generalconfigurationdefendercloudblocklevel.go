package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderCloudBlockLevel string

const (
	Windows10GeneralConfigurationDefenderCloudBlockLevel_High          Windows10GeneralConfigurationDefenderCloudBlockLevel = "high"
	Windows10GeneralConfigurationDefenderCloudBlockLevel_HighPlus      Windows10GeneralConfigurationDefenderCloudBlockLevel = "highPlus"
	Windows10GeneralConfigurationDefenderCloudBlockLevel_NotConfigured Windows10GeneralConfigurationDefenderCloudBlockLevel = "notConfigured"
	Windows10GeneralConfigurationDefenderCloudBlockLevel_ZeroTolerance Windows10GeneralConfigurationDefenderCloudBlockLevel = "zeroTolerance"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderCloudBlockLevel() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderCloudBlockLevel_High),
		string(Windows10GeneralConfigurationDefenderCloudBlockLevel_HighPlus),
		string(Windows10GeneralConfigurationDefenderCloudBlockLevel_NotConfigured),
		string(Windows10GeneralConfigurationDefenderCloudBlockLevel_ZeroTolerance),
	}
}

func (s *Windows10GeneralConfigurationDefenderCloudBlockLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDefenderCloudBlockLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDefenderCloudBlockLevel(input string) (*Windows10GeneralConfigurationDefenderCloudBlockLevel, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderCloudBlockLevel{
		"high":          Windows10GeneralConfigurationDefenderCloudBlockLevel_High,
		"highplus":      Windows10GeneralConfigurationDefenderCloudBlockLevel_HighPlus,
		"notconfigured": Windows10GeneralConfigurationDefenderCloudBlockLevel_NotConfigured,
		"zerotolerance": Windows10GeneralConfigurationDefenderCloudBlockLevel_ZeroTolerance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderCloudBlockLevel(input)
	return &out, nil
}
