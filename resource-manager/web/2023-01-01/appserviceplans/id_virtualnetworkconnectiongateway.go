package appserviceplans

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VirtualNetworkConnectionGatewayId{}

// VirtualNetworkConnectionGatewayId is a struct representing the Resource ID for a Virtual Network Connection Gateway
type VirtualNetworkConnectionGatewayId struct {
	SubscriptionId               string
	ResourceGroupName            string
	ServerFarmName               string
	VirtualNetworkConnectionName string
	GatewayName                  string
}

// NewVirtualNetworkConnectionGatewayID returns a new VirtualNetworkConnectionGatewayId struct
func NewVirtualNetworkConnectionGatewayID(subscriptionId string, resourceGroupName string, serverFarmName string, virtualNetworkConnectionName string, gatewayName string) VirtualNetworkConnectionGatewayId {
	return VirtualNetworkConnectionGatewayId{
		SubscriptionId:               subscriptionId,
		ResourceGroupName:            resourceGroupName,
		ServerFarmName:               serverFarmName,
		VirtualNetworkConnectionName: virtualNetworkConnectionName,
		GatewayName:                  gatewayName,
	}
}

// ParseVirtualNetworkConnectionGatewayID parses 'input' into a VirtualNetworkConnectionGatewayId
func ParseVirtualNetworkConnectionGatewayID(input string) (*VirtualNetworkConnectionGatewayId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualNetworkConnectionGatewayId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualNetworkConnectionGatewayId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerFarmName, ok = parsed.Parsed["serverFarmName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverFarmName", *parsed)
	}

	if id.VirtualNetworkConnectionName, ok = parsed.Parsed["virtualNetworkConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualNetworkConnectionName", *parsed)
	}

	if id.GatewayName, ok = parsed.Parsed["gatewayName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "gatewayName", *parsed)
	}

	return &id, nil
}

// ParseVirtualNetworkConnectionGatewayIDInsensitively parses 'input' case-insensitively into a VirtualNetworkConnectionGatewayId
// note: this method should only be used for API response data and not user input
func ParseVirtualNetworkConnectionGatewayIDInsensitively(input string) (*VirtualNetworkConnectionGatewayId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualNetworkConnectionGatewayId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualNetworkConnectionGatewayId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerFarmName, ok = parsed.Parsed["serverFarmName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverFarmName", *parsed)
	}

	if id.VirtualNetworkConnectionName, ok = parsed.Parsed["virtualNetworkConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualNetworkConnectionName", *parsed)
	}

	if id.GatewayName, ok = parsed.Parsed["gatewayName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "gatewayName", *parsed)
	}

	return &id, nil
}

// ValidateVirtualNetworkConnectionGatewayID checks that 'input' can be parsed as a Virtual Network Connection Gateway ID
func ValidateVirtualNetworkConnectionGatewayID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualNetworkConnectionGatewayID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Network Connection Gateway ID
func (id VirtualNetworkConnectionGatewayId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/serverFarms/%s/virtualNetworkConnections/%s/gateways/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerFarmName, id.VirtualNetworkConnectionName, id.GatewayName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Network Connection Gateway ID
func (id VirtualNetworkConnectionGatewayId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticServerFarms", "serverFarms", "serverFarms"),
		resourceids.UserSpecifiedSegment("serverFarmName", "serverFarmValue"),
		resourceids.StaticSegment("staticVirtualNetworkConnections", "virtualNetworkConnections", "virtualNetworkConnections"),
		resourceids.UserSpecifiedSegment("virtualNetworkConnectionName", "virtualNetworkConnectionValue"),
		resourceids.StaticSegment("staticGateways", "gateways", "gateways"),
		resourceids.UserSpecifiedSegment("gatewayName", "gatewayValue"),
	}
}

// String returns a human-readable description of this Virtual Network Connection Gateway ID
func (id VirtualNetworkConnectionGatewayId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Farm Name: %q", id.ServerFarmName),
		fmt.Sprintf("Virtual Network Connection Name: %q", id.VirtualNetworkConnectionName),
		fmt.Sprintf("Gateway Name: %q", id.GatewayName),
	}
	return fmt.Sprintf("Virtual Network Connection Gateway (%s)", strings.Join(components, "\n"))
}
