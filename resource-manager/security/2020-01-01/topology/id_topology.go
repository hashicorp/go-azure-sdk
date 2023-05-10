package topology

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TopologyId{}

// TopologyId is a struct representing the Resource ID for a Topology
type TopologyId struct {
	SubscriptionId    string
	ResourceGroupName string
	LocationName      string
	TopologyName      string
}

// NewTopologyID returns a new TopologyId struct
func NewTopologyID(subscriptionId string, resourceGroupName string, locationName string, topologyName string) TopologyId {
	return TopologyId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LocationName:      locationName,
		TopologyName:      topologyName,
	}
}

// ParseTopologyID parses 'input' into a TopologyId
func ParseTopologyID(input string) (*TopologyId, error) {
	parser := resourceids.NewParserFromResourceIdType(TopologyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TopologyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.TopologyName, ok = parsed.Parsed["topologyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "topologyName", *parsed)
	}

	return &id, nil
}

// ParseTopologyIDInsensitively parses 'input' case-insensitively into a TopologyId
// note: this method should only be used for API response data and not user input
func ParseTopologyIDInsensitively(input string) (*TopologyId, error) {
	parser := resourceids.NewParserFromResourceIdType(TopologyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TopologyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.TopologyName, ok = parsed.Parsed["topologyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "topologyName", *parsed)
	}

	return &id, nil
}

// ValidateTopologyID checks that 'input' can be parsed as a Topology ID
func ValidateTopologyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTopologyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Topology ID
func (id TopologyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/locations/%s/topologies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.TopologyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Topology ID
func (id TopologyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticTopologies", "topologies", "topologies"),
		resourceids.UserSpecifiedSegment("topologyName", "topologyValue"),
	}
}

// String returns a human-readable description of this Topology ID
func (id TopologyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Topology Name: %q", id.TopologyName),
	}
	return fmt.Sprintf("Topology (%s)", strings.Join(components, "\n"))
}
