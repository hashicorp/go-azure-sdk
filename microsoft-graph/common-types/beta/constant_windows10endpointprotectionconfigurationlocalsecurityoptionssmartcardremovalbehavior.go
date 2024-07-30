package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_DisconnectRemoteDesktopSession Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior = "disconnectRemoteDesktopSession"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_ForceLogoff                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior = "forceLogoff"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_LockWorkstation                Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior = "lockWorkstation"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_NoAction                       Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior = "noAction"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_DisconnectRemoteDesktopSession),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_ForceLogoff),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_LockWorkstation),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_NoAction),
	}
}

func (s *Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior{
		"disconnectremotedesktopsession": Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_DisconnectRemoteDesktopSession,
		"forcelogoff":                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_ForceLogoff,
		"lockworkstation":                Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_LockWorkstation,
		"noaction":                       Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior_NoAction,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsSmartCardRemovalBehavior(input)
	return &out, nil
}
