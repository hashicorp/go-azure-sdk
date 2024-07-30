package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod string

const (
	Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_DeviceDefault Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod = "deviceDefault"
	Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_None          Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod = "none"
	Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_UtF8          Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod = "utF8"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_DeviceDefault),
		string(Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_None),
		string(Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_UtF8),
	}
}

func (s *Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod(input string) (*Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod{
		"devicedefault": Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_DeviceDefault,
		"none":          Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_None,
		"utf8":          Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod_UtF8,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod(input)
	return &out, nil
}
