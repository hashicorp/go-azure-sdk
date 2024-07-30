package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection string

const (
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_Disabled      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection = "disabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_Enabled       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection = "enabled"
	MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_NotConfigured MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection = "notConfigured"
)

func PossibleValuesForMacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection() []string {
	return []string{
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_Disabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_Enabled),
		string(MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_NotConfigured),
	}
}

func (s *MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection(input string) (*MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection, error) {
	vals := map[string]MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection{
		"disabled":      MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_Disabled,
		"enabled":       MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_Enabled,
		"notconfigured": MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEndpointProtectionConfigurationAdvancedThreatProtectionDiagnosticDataCollection(input)
	return &out, nil
}
