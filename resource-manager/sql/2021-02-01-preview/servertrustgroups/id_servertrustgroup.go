package servertrustgroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServerTrustGroupId{}

// ServerTrustGroupId is a struct representing the Resource ID for a Server Trust Group
type ServerTrustGroupId struct {
	SubscriptionId       string
	ResourceGroupName    string
	LocationName         string
	ServerTrustGroupName string
}

// NewServerTrustGroupID returns a new ServerTrustGroupId struct
func NewServerTrustGroupID(subscriptionId string, resourceGroupName string, locationName string, serverTrustGroupName string) ServerTrustGroupId {
	return ServerTrustGroupId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		LocationName:         locationName,
		ServerTrustGroupName: serverTrustGroupName,
	}
}

// ParseServerTrustGroupID parses 'input' into a ServerTrustGroupId
func ParseServerTrustGroupID(input string) (*ServerTrustGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerTrustGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerTrustGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.ServerTrustGroupName, ok = parsed.Parsed["serverTrustGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverTrustGroupName", *parsed)
	}

	return &id, nil
}

// ParseServerTrustGroupIDInsensitively parses 'input' case-insensitively into a ServerTrustGroupId
// note: this method should only be used for API response data and not user input
func ParseServerTrustGroupIDInsensitively(input string) (*ServerTrustGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerTrustGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerTrustGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.ServerTrustGroupName, ok = parsed.Parsed["serverTrustGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverTrustGroupName", *parsed)
	}

	return &id, nil
}

// ValidateServerTrustGroupID checks that 'input' can be parsed as a Server Trust Group ID
func ValidateServerTrustGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServerTrustGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Server Trust Group ID
func (id ServerTrustGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/locations/%s/serverTrustGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.ServerTrustGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Server Trust Group ID
func (id ServerTrustGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticServerTrustGroups", "serverTrustGroups", "serverTrustGroups"),
		resourceids.UserSpecifiedSegment("serverTrustGroupName", "serverTrustGroupValue"),
	}
}

// String returns a human-readable description of this Server Trust Group ID
func (id ServerTrustGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Server Trust Group Name: %q", id.ServerTrustGroupName),
	}
	return fmt.Sprintf("Server Trust Group (%s)", strings.Join(components, "\n"))
}
