package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode string

const (
	Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Automatic Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode = "automatic"
	Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Disabled  Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode = "disabled"
	Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Manual    Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode = "manual"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Automatic),
		string(Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Disabled),
		string(Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Manual),
	}
}

func (s *Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode(input string) (*Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode{
		"automatic": Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Automatic,
		"disabled":  Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Disabled,
		"manual":    Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationXboxServicesAccessoryManagementServiceStartupMode(input)
	return &out, nil
}
