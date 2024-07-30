package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod string

const (
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_DeviceDefault Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "deviceDefault"
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_Disabled      Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "disabled"
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueBoth     Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "queueBoth"
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueInbound  Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "queueInbound"
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueOutbound Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "queueOutbound"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationFirewallPacketQueueingMethod() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_DeviceDefault),
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_Disabled),
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueBoth),
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueInbound),
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueOutbound),
	}
}

func (s *Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationFirewallPacketQueueingMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationFirewallPacketQueueingMethod(input string) (*Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod{
		"devicedefault": Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_DeviceDefault,
		"disabled":      Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_Disabled,
		"queueboth":     Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueBoth,
		"queueinbound":  Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueInbound,
		"queueoutbound": Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod_QueueOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod(input)
	return &out, nil
}
