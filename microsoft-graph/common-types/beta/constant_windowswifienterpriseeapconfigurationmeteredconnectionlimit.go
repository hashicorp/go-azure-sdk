package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit string

const (
	WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Fixed        WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit = "fixed"
	WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Unrestricted WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit = "unrestricted"
	WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Variable     WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit = "variable"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Fixed),
		string(WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Unrestricted),
		string(WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Variable),
	}
}

func (s *WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit(input string) (*WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit{
		"fixed":        WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Fixed,
		"unrestricted": WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Unrestricted,
		"variable":     WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit_Variable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationMeteredConnectionLimit(input)
	return &out, nil
}
