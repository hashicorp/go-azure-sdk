package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuAppListVisibility string

const (
	Windows10GeneralConfigurationStartMenuAppListVisibility_Collapse           Windows10GeneralConfigurationStartMenuAppListVisibility = "collapse"
	Windows10GeneralConfigurationStartMenuAppListVisibility_DisableSettingsApp Windows10GeneralConfigurationStartMenuAppListVisibility = "disableSettingsApp"
	Windows10GeneralConfigurationStartMenuAppListVisibility_Remove             Windows10GeneralConfigurationStartMenuAppListVisibility = "remove"
	Windows10GeneralConfigurationStartMenuAppListVisibility_UserDefined        Windows10GeneralConfigurationStartMenuAppListVisibility = "userDefined"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuAppListVisibility() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuAppListVisibility_Collapse),
		string(Windows10GeneralConfigurationStartMenuAppListVisibility_DisableSettingsApp),
		string(Windows10GeneralConfigurationStartMenuAppListVisibility_Remove),
		string(Windows10GeneralConfigurationStartMenuAppListVisibility_UserDefined),
	}
}

func (s *Windows10GeneralConfigurationStartMenuAppListVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationStartMenuAppListVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationStartMenuAppListVisibility(input string) (*Windows10GeneralConfigurationStartMenuAppListVisibility, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuAppListVisibility{
		"collapse":           Windows10GeneralConfigurationStartMenuAppListVisibility_Collapse,
		"disablesettingsapp": Windows10GeneralConfigurationStartMenuAppListVisibility_DisableSettingsApp,
		"remove":             Windows10GeneralConfigurationStartMenuAppListVisibility_Remove,
		"userdefined":        Windows10GeneralConfigurationStartMenuAppListVisibility_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuAppListVisibility(input)
	return &out, nil
}
