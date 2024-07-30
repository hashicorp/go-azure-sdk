package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions string

const (
	Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForEdge          Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions = "enabledForEdge"
	Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForEdgeAndOffice Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions = "enabledForEdgeAndOffice"
	Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForOffice        Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions = "enabledForOffice"
	Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_NotConfigured           Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationApplicationGuardEnabledOptions() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForEdge),
		string(Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForEdgeAndOffice),
		string(Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForOffice),
		string(Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationApplicationGuardEnabledOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationApplicationGuardEnabledOptions(input string) (*Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions{
		"enabledforedge":          Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForEdge,
		"enabledforedgeandoffice": Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForEdgeAndOffice,
		"enabledforoffice":        Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_EnabledForOffice,
		"notconfigured":           Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationApplicationGuardEnabledOptions(input)
	return &out, nil
}
