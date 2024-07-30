package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection string

const (
	Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_Disable       Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection = "disable"
	Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_Enable        Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection = "enable"
	Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_NotConfigured Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationWindowsDefenderTamperProtection() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_Disable),
		string(Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_Enable),
		string(Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationWindowsDefenderTamperProtection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationWindowsDefenderTamperProtection(input string) (*Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection{
		"disable":       Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_Disable,
		"enable":        Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_Enable,
		"notconfigured": Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationWindowsDefenderTamperProtection(input)
	return &out, nil
}
