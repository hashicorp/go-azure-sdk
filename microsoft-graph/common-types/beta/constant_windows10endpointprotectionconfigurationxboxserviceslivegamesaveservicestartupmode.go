package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode string

const (
	Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Automatic Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode = "automatic"
	Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Disabled  Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode = "disabled"
	Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Manual    Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode = "manual"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Automatic),
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Disabled),
		string(Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Manual),
	}
}

func (s *Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode(input string) (*Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode{
		"automatic": Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Automatic,
		"disabled":  Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Disabled,
		"manual":    Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationXboxServicesLiveGameSaveServiceStartupMode(input)
	return &out, nil
}
