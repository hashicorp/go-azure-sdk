package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod string

const (
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethoddeviceDefault Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "DeviceDefault"
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethoddisabled      Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "Disabled"
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueBoth     Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "QueueBoth"
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueInbound  Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "QueueInbound"
	Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueOutbound Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod = "QueueOutbound"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationFirewallPacketQueueingMethod() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethoddeviceDefault),
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethoddisabled),
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueBoth),
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueInbound),
		string(Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueOutbound),
	}
}

func parseWindows10EndpointProtectionConfigurationFirewallPacketQueueingMethod(input string) (*Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod{
		"devicedefault": Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethoddeviceDefault,
		"disabled":      Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethoddisabled,
		"queueboth":     Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueBoth,
		"queueinbound":  Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueInbound,
		"queueoutbound": Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethodqueueOutbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod(input)
	return &out, nil
}
