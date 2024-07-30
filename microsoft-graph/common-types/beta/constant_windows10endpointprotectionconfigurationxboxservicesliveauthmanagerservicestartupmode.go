package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode string

const (
	Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Automatic Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode = "automatic"
	Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Disabled  Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode = "disabled"
	Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Manual    Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode = "manual"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Automatic),
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Disabled),
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Manual),
	}
}

func (s *Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode(input string) (*Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode{
		"automatic": Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Automatic,
		"disabled":  Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Disabled,
		"manual":    Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationXboxServicesLiveAuthManagerServiceStartupMode(input)
	return &out, nil
}
