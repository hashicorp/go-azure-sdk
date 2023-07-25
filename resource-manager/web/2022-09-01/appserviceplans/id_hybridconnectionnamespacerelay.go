package appserviceplans

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = HybridConnectionNamespaceRelayId{}

// HybridConnectionNamespaceRelayId is a struct representing the Resource ID for a Hybrid Connection Namespace Relay
type HybridConnectionNamespaceRelayId struct {
	SubscriptionId                string
	ResourceGroupName             string
	ServerFarmName                string
	HybridConnectionNamespaceName string
	RelayName                     string
}

// NewHybridConnectionNamespaceRelayID returns a new HybridConnectionNamespaceRelayId struct
func NewHybridConnectionNamespaceRelayID(subscriptionId string, resourceGroupName string, serverFarmName string, hybridConnectionNamespaceName string, relayName string) HybridConnectionNamespaceRelayId {
	return HybridConnectionNamespaceRelayId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		ServerFarmName:                serverFarmName,
		HybridConnectionNamespaceName: hybridConnectionNamespaceName,
		RelayName:                     relayName,
	}
}

// ParseHybridConnectionNamespaceRelayID parses 'input' into a HybridConnectionNamespaceRelayId
func ParseHybridConnectionNamespaceRelayID(input string) (*HybridConnectionNamespaceRelayId, error) {
	parser := resourceids.NewParserFromResourceIdType(HybridConnectionNamespaceRelayId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HybridConnectionNamespaceRelayId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerFarmName, ok = parsed.Parsed["serverFarmName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverFarmName", *parsed)
	}

	if id.HybridConnectionNamespaceName, ok = parsed.Parsed["hybridConnectionNamespaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hybridConnectionNamespaceName", *parsed)
	}

	if id.RelayName, ok = parsed.Parsed["relayName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relayName", *parsed)
	}

	return &id, nil
}

// ParseHybridConnectionNamespaceRelayIDInsensitively parses 'input' case-insensitively into a HybridConnectionNamespaceRelayId
// note: this method should only be used for API response data and not user input
func ParseHybridConnectionNamespaceRelayIDInsensitively(input string) (*HybridConnectionNamespaceRelayId, error) {
	parser := resourceids.NewParserFromResourceIdType(HybridConnectionNamespaceRelayId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HybridConnectionNamespaceRelayId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerFarmName, ok = parsed.Parsed["serverFarmName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverFarmName", *parsed)
	}

	if id.HybridConnectionNamespaceName, ok = parsed.Parsed["hybridConnectionNamespaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hybridConnectionNamespaceName", *parsed)
	}

	if id.RelayName, ok = parsed.Parsed["relayName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "relayName", *parsed)
	}

	return &id, nil
}

// ValidateHybridConnectionNamespaceRelayID checks that 'input' can be parsed as a Hybrid Connection Namespace Relay ID
func ValidateHybridConnectionNamespaceRelayID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHybridConnectionNamespaceRelayID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hybrid Connection Namespace Relay ID
func (id HybridConnectionNamespaceRelayId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/serverFarms/%s/hybridConnectionNamespaces/%s/relays/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerFarmName, id.HybridConnectionNamespaceName, id.RelayName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hybrid Connection Namespace Relay ID
func (id HybridConnectionNamespaceRelayId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticServerFarms", "serverFarms", "serverFarms"),
		resourceids.UserSpecifiedSegment("serverFarmName", "serverFarmValue"),
		resourceids.StaticSegment("staticHybridConnectionNamespaces", "hybridConnectionNamespaces", "hybridConnectionNamespaces"),
		resourceids.UserSpecifiedSegment("hybridConnectionNamespaceName", "hybridConnectionNamespaceValue"),
		resourceids.StaticSegment("staticRelays", "relays", "relays"),
		resourceids.UserSpecifiedSegment("relayName", "relayValue"),
	}
}

// String returns a human-readable description of this Hybrid Connection Namespace Relay ID
func (id HybridConnectionNamespaceRelayId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Farm Name: %q", id.ServerFarmName),
		fmt.Sprintf("Hybrid Connection Namespace Name: %q", id.HybridConnectionNamespaceName),
		fmt.Sprintf("Relay Name: %q", id.RelayName),
	}
	return fmt.Sprintf("Hybrid Connection Namespace Relay (%s)", strings.Join(components, "\n"))
}
