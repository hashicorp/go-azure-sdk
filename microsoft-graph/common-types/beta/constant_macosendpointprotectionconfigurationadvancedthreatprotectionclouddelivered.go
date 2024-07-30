package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered string

const (
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_Disabled      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered = "disabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_Enabled       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered = "enabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_NotConfigured MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered = "notConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_Disabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_Enabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_NotConfigured),
	}
}

func (s *MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered(input string) (*MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered{
		"disabled":      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_Disabled,
		"enabled":       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_Enabled,
		"notconfigured": MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationAdvancedThreatProtectionCloudDelivered(input)
	return &out, nil
}
