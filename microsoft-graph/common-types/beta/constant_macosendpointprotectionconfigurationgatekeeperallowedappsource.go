package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource string

const (
	MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_Anywhere                           MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource = "anywhere"
	MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_MacAppStore                        MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource = "macAppStore"
	MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_MacAppStoreAndIdentifiedDevelopers MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource = "macAppStoreAndIdentifiedDevelopers"
	MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_NotConfigured                      MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource = "notConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_Anywhere),
		string(MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_MacAppStore),
		string(MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_MacAppStoreAndIdentifiedDevelopers),
		string(MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_NotConfigured),
	}
}

func (s *MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource(input string) (*MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource{
		"anywhere":                           MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_Anywhere,
		"macappstore":                        MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_MacAppStore,
		"macappstoreandidentifieddevelopers": MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_MacAppStoreAndIdentifiedDevelopers,
		"notconfigured":                      MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationGatekeeperAllowedAppSource(input)
	return &out, nil
}
