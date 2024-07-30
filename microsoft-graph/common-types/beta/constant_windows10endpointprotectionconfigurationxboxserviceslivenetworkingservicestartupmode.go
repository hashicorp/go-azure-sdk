package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode string

const (
	Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Automatic Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode = "automatic"
	Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Disabled  Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode = "disabled"
	Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Manual    Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode = "manual"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Automatic),
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Disabled),
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Manual),
	}
}

func (s *Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode(input string) (*Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode{
		"automatic": Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Automatic,
		"disabled":  Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Disabled,
		"manual":    Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationXboxServicesLiveNetworkingServiceStartupMode(input)
	return &out, nil
}
