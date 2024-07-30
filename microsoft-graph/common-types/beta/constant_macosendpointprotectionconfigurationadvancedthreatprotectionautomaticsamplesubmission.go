package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission string

const (
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_Disabled      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission = "disabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_Enabled       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission = "enabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_NotConfigured MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission = "notConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_Disabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_Enabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_NotConfigured),
	}
}

func (s *MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission(input string) (*MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission{
		"disabled":      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_Disabled,
		"enabled":       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_Enabled,
		"notconfigured": MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationAdvancedThreatProtectionAutomaticSampleSubmission(input)
	return &out, nil
}
