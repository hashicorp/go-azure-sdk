package topology

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TopologyId{})
}

var _ resourceids.ResourceId = &TopologyId{}

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
	parser := resourceids.NewParserFromResourceIdType(&TopologyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TopologyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTopologyIDInsensitively parses 'input' case-insensitively into a TopologyId
// note: this method should only be used for API response data and not user input
func ParseTopologyIDInsensitively(input string) (*TopologyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TopologyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TopologyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TopologyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.TopologyName, ok = input.Parsed["topologyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "topologyName", input)
	}

	return nil
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
