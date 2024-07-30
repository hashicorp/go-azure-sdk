package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderScanType string

const (
	Windows10GeneralConfigurationDefenderScanType_Disabled    Windows10GeneralConfigurationDefenderScanType = "disabled"
	Windows10GeneralConfigurationDefenderScanType_Full        Windows10GeneralConfigurationDefenderScanType = "full"
	Windows10GeneralConfigurationDefenderScanType_Quick       Windows10GeneralConfigurationDefenderScanType = "quick"
	Windows10GeneralConfigurationDefenderScanType_UserDefined Windows10GeneralConfigurationDefenderScanType = "userDefined"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderScanType() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderScanType_Disabled),
		string(Windows10GeneralConfigurationDefenderScanType_Full),
		string(Windows10GeneralConfigurationDefenderScanType_Quick),
		string(Windows10GeneralConfigurationDefenderScanType_UserDefined),
	}
}

func (s *Windows10GeneralConfigurationDefenderScanType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDefenderScanType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDefenderScanType(input string) (*Windows10GeneralConfigurationDefenderScanType, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderScanType{
		"disabled":    Windows10GeneralConfigurationDefenderScanType_Disabled,
		"full":        Windows10GeneralConfigurationDefenderScanType_Full,
		"quick":       Windows10GeneralConfigurationDefenderScanType_Quick,
		"userdefined": Windows10GeneralConfigurationDefenderScanType_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderScanType(input)
	return &out, nil
}
