package appserviceplans

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServerFarmId{}

// ServerFarmId is a struct representing the Resource ID for a Server Farm
type ServerFarmId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerFarmName    string
}

// NewServerFarmID returns a new ServerFarmId struct
func NewServerFarmID(subscriptionId string, resourceGroupName string, serverFarmName string) ServerFarmId {
	return ServerFarmId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerFarmName:    serverFarmName,
	}
}

// ParseServerFarmID parses 'input' into a ServerFarmId
func ParseServerFarmID(input string) (*ServerFarmId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerFarmId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerFarmId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerFarmName, ok = parsed.Parsed["serverFarmName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverFarmName", *parsed)
	}

	return &id, nil
}

// ParseServerFarmIDInsensitively parses 'input' case-insensitively into a ServerFarmId
// note: this method should only be used for API response data and not user input
func ParseServerFarmIDInsensitively(input string) (*ServerFarmId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerFarmId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerFarmId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerFarmName, ok = parsed.Parsed["serverFarmName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverFarmName", *parsed)
	}

	return &id, nil
}

// ValidateServerFarmID checks that 'input' can be parsed as a Server Farm ID
func ValidateServerFarmID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServerFarmID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Server Farm ID
func (id ServerFarmId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/serverFarms/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerFarmName)
}

// Segments returns a slice of Resource ID Segments which comprise this Server Farm ID
func (id ServerFarmId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticServerFarms", "serverFarms", "serverFarms"),
		resourceids.UserSpecifiedSegment("serverFarmName", "serverFarmValue"),
	}
}

// String returns a human-readable description of this Server Farm ID
func (id ServerFarmId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Farm Name: %q", id.ServerFarmName),
	}
	return fmt.Sprintf("Server Farm (%s)", strings.Join(components, "\n"))
}
