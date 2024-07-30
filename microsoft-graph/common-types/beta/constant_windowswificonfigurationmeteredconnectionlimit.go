package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiConfigurationMeteredConnectionLimit string

const (
	WindowsWifiConfigurationMeteredConnectionLimit_Fixed        WindowsWifiConfigurationMeteredConnectionLimit = "fixed"
	WindowsWifiConfigurationMeteredConnectionLimit_Unrestricted WindowsWifiConfigurationMeteredConnectionLimit = "unrestricted"
	WindowsWifiConfigurationMeteredConnectionLimit_Variable     WindowsWifiConfigurationMeteredConnectionLimit = "variable"
)

func PossibleValuesForWindowsWifiConfigurationMeteredConnectionLimit() []string {
	return []string{
		string(WindowsWifiConfigurationMeteredConnectionLimit_Fixed),
		string(WindowsWifiConfigurationMeteredConnectionLimit_Unrestricted),
		string(WindowsWifiConfigurationMeteredConnectionLimit_Variable),
	}
}

func (s *WindowsWifiConfigurationMeteredConnectionLimit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiConfigurationMeteredConnectionLimit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiConfigurationMeteredConnectionLimit(input string) (*WindowsWifiConfigurationMeteredConnectionLimit, error) {
	vals := map[string]WindowsWifiConfigurationMeteredConnectionLimit{
		"fixed":        WindowsWifiConfigurationMeteredConnectionLimit_Fixed,
		"unrestricted": WindowsWifiConfigurationMeteredConnectionLimit_Unrestricted,
		"variable":     WindowsWifiConfigurationMeteredConnectionLimit_Variable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiConfigurationMeteredConnectionLimit(input)
	return &out, nil
}
