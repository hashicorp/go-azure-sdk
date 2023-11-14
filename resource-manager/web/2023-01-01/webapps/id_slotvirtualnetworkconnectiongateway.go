package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotVirtualNetworkConnectionGatewayId{}

// SlotVirtualNetworkConnectionGatewayId is a struct representing the Resource ID for a Slot Virtual Network Connection Gateway
type SlotVirtualNetworkConnectionGatewayId struct {
	SubscriptionId               string
	ResourceGroupName            string
	SiteName                     string
	SlotName                     string
	VirtualNetworkConnectionName string
	GatewayName                  string
}

// NewSlotVirtualNetworkConnectionGatewayID returns a new SlotVirtualNetworkConnectionGatewayId struct
func NewSlotVirtualNetworkConnectionGatewayID(subscriptionId string, resourceGroupName string, siteName string, slotName string, virtualNetworkConnectionName string, gatewayName string) SlotVirtualNetworkConnectionGatewayId {
	return SlotVirtualNetworkConnectionGatewayId{
		SubscriptionId:               subscriptionId,
		ResourceGroupName:            resourceGroupName,
		SiteName:                     siteName,
		SlotName:                     slotName,
		VirtualNetworkConnectionName: virtualNetworkConnectionName,
		GatewayName:                  gatewayName,
	}
}

// ParseSlotVirtualNetworkConnectionGatewayID parses 'input' into a SlotVirtualNetworkConnectionGatewayId
func ParseSlotVirtualNetworkConnectionGatewayID(input string) (*SlotVirtualNetworkConnectionGatewayId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotVirtualNetworkConnectionGatewayId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotVirtualNetworkConnectionGatewayId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.SlotName, ok = parsed.Parsed["slotName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "slotName", *parsed)
	}

	if id.VirtualNetworkConnectionName, ok = parsed.Parsed["virtualNetworkConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualNetworkConnectionName", *parsed)
	}

	if id.GatewayName, ok = parsed.Parsed["gatewayName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "gatewayName", *parsed)
	}

	return &id, nil
}

// ParseSlotVirtualNetworkConnectionGatewayIDInsensitively parses 'input' case-insensitively into a SlotVirtualNetworkConnectionGatewayId
// note: this method should only be used for API response data and not user input
func ParseSlotVirtualNetworkConnectionGatewayIDInsensitively(input string) (*SlotVirtualNetworkConnectionGatewayId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotVirtualNetworkConnectionGatewayId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotVirtualNetworkConnectionGatewayId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.SlotName, ok = parsed.Parsed["slotName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "slotName", *parsed)
	}

	if id.VirtualNetworkConnectionName, ok = parsed.Parsed["virtualNetworkConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualNetworkConnectionName", *parsed)
	}

	if id.GatewayName, ok = parsed.Parsed["gatewayName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "gatewayName", *parsed)
	}

	return &id, nil
}

// ValidateSlotVirtualNetworkConnectionGatewayID checks that 'input' can be parsed as a Slot Virtual Network Connection Gateway ID
func ValidateSlotVirtualNetworkConnectionGatewayID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotVirtualNetworkConnectionGatewayID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Virtual Network Connection Gateway ID
func (id SlotVirtualNetworkConnectionGatewayId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/virtualNetworkConnections/%s/gateways/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.VirtualNetworkConnectionName, id.GatewayName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Virtual Network Connection Gateway ID
func (id SlotVirtualNetworkConnectionGatewayId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticSlots", "slots", "slots"),
		resourceids.UserSpecifiedSegment("slotName", "slotValue"),
		resourceids.StaticSegment("staticVirtualNetworkConnections", "virtualNetworkConnections", "virtualNetworkConnections"),
		resourceids.UserSpecifiedSegment("virtualNetworkConnectionName", "virtualNetworkConnectionValue"),
		resourceids.StaticSegment("staticGateways", "gateways", "gateways"),
		resourceids.UserSpecifiedSegment("gatewayName", "gatewayValue"),
	}
}

// String returns a human-readable description of this Slot Virtual Network Connection Gateway ID
func (id SlotVirtualNetworkConnectionGatewayId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Virtual Network Connection Name: %q", id.VirtualNetworkConnectionName),
		fmt.Sprintf("Gateway Name: %q", id.GatewayName),
	}
	return fmt.Sprintf("Slot Virtual Network Connection Gateway (%s)", strings.Join(components, "\n"))
}
