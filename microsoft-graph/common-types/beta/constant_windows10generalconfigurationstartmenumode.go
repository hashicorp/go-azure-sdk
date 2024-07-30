package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuMode string

const (
	Windows10GeneralConfigurationStartMenuMode_FullScreen    Windows10GeneralConfigurationStartMenuMode = "fullScreen"
	Windows10GeneralConfigurationStartMenuMode_NonFullScreen Windows10GeneralConfigurationStartMenuMode = "nonFullScreen"
	Windows10GeneralConfigurationStartMenuMode_UserDefined   Windows10GeneralConfigurationStartMenuMode = "userDefined"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuMode() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuMode_FullScreen),
		string(Windows10GeneralConfigurationStartMenuMode_NonFullScreen),
		string(Windows10GeneralConfigurationStartMenuMode_UserDefined),
	}
}

func (s *Windows10GeneralConfigurationStartMenuMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuMode(input string) (*Windows10GeneralConfigurationStartMenuMode, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuMode{
		"fullscreen":    Windows10GeneralConfigurationStartMenuMode_FullScreen,
		"nonfullscreen": Windows10GeneralConfigurationStartMenuMode_NonFullScreen,
		"userdefined":   Windows10GeneralConfigurationStartMenuMode_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuMode(input)
	return &out, nil
}
