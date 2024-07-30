package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime string

const (
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_Disabled      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime = "disabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_Enabled       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime = "enabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_NotConfigured MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime = "notConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_Disabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_Enabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_NotConfigured),
	}
}

func (s *MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime(input string) (*MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime{
		"disabled":      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_Disabled,
		"enabled":       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_Enabled,
		"notconfigured": MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationAdvancedThreatProtectionRealTime(input)
	return &out, nil
}
