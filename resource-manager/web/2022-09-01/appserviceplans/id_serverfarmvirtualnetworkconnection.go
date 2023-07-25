package appserviceplans

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServerFarmVirtualNetworkConnectionId{}

// ServerFarmVirtualNetworkConnectionId is a struct representing the Resource ID for a Server Farm Virtual Network Connection
type ServerFarmVirtualNetworkConnectionId struct {
	SubscriptionId               string
	ResourceGroupName            string
	ServerFarmName               string
	VirtualNetworkConnectionName string
}

// NewServerFarmVirtualNetworkConnectionID returns a new ServerFarmVirtualNetworkConnectionId struct
func NewServerFarmVirtualNetworkConnectionID(subscriptionId string, resourceGroupName string, serverFarmName string, virtualNetworkConnectionName string) ServerFarmVirtualNetworkConnectionId {
	return ServerFarmVirtualNetworkConnectionId{
		SubscriptionId:               subscriptionId,
		ResourceGroupName:            resourceGroupName,
		ServerFarmName:               serverFarmName,
		VirtualNetworkConnectionName: virtualNetworkConnectionName,
	}
}

// ParseServerFarmVirtualNetworkConnectionID parses 'input' into a ServerFarmVirtualNetworkConnectionId
func ParseServerFarmVirtualNetworkConnectionID(input string) (*ServerFarmVirtualNetworkConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerFarmVirtualNetworkConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerFarmVirtualNetworkConnectionId{}

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

	return &id, nil
}

// ParseServerFarmVirtualNetworkConnectionIDInsensitively parses 'input' case-insensitively into a ServerFarmVirtualNetworkConnectionId
// note: this method should only be used for API response data and not user input
func ParseServerFarmVirtualNetworkConnectionIDInsensitively(input string) (*ServerFarmVirtualNetworkConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerFarmVirtualNetworkConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerFarmVirtualNetworkConnectionId{}

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

	return &id, nil
}

// ValidateServerFarmVirtualNetworkConnectionID checks that 'input' can be parsed as a Server Farm Virtual Network Connection ID
func ValidateServerFarmVirtualNetworkConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServerFarmVirtualNetworkConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Server Farm Virtual Network Connection ID
func (id ServerFarmVirtualNetworkConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/serverFarms/%s/virtualNetworkConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerFarmName, id.VirtualNetworkConnectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Server Farm Virtual Network Connection ID
func (id ServerFarmVirtualNetworkConnectionId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Server Farm Virtual Network Connection ID
func (id ServerFarmVirtualNetworkConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Farm Name: %q", id.ServerFarmName),
		fmt.Sprintf("Virtual Network Connection Name: %q", id.VirtualNetworkConnectionName),
	}
	return fmt.Sprintf("Server Farm Virtual Network Connection (%s)", strings.Join(components, "\n"))
}
