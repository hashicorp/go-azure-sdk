package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RackId{}

// RackId is a struct representing the Resource ID for a Rack
type RackId struct {
	SubscriptionId    string
	ResourceGroupName string
	RackName          string
}

// NewRackID returns a new RackId struct
func NewRackID(subscriptionId string, resourceGroupName string, rackName string) RackId {
	return RackId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		RackName:          rackName,
	}
}

// ParseRackID parses 'input' into a RackId
func ParseRackID(input string) (*RackId, error) {
	parser := resourceids.NewParserFromResourceIdType(RackId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RackId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.RackName, ok = parsed.Parsed["rackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rackName", *parsed)
	}

	return &id, nil
}

// ParseRackIDInsensitively parses 'input' case-insensitively into a RackId
// note: this method should only be used for API response data and not user input
func ParseRackIDInsensitively(input string) (*RackId, error) {
	parser := resourceids.NewParserFromResourceIdType(RackId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RackId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.RackName, ok = parsed.Parsed["rackName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rackName", *parsed)
	}

	return &id, nil
}

// ValidateRackID checks that 'input' can be parsed as a Rack ID
func ValidateRackID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRackID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rack ID
func (id RackId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/racks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RackName)
}

// Segments returns a slice of Resource ID Segments which comprise this Rack ID
func (id RackId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticRacks", "racks", "racks"),
		resourceids.UserSpecifiedSegment("rackName", "rackValue"),
	}
}

// String returns a human-readable description of this Rack ID
func (id RackId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Rack Name: %q", id.RackName),
	}
	return fmt.Sprintf("Rack (%s)", strings.Join(components, "\n"))
}
