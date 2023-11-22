package fleetupdatestrategies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UpdateStrategyId{}

// UpdateStrategyId is a struct representing the Resource ID for a Update Strategy
type UpdateStrategyId struct {
	SubscriptionId     string
	ResourceGroupName  string
	FleetName          string
	UpdateStrategyName string
}

// NewUpdateStrategyID returns a new UpdateStrategyId struct
func NewUpdateStrategyID(subscriptionId string, resourceGroupName string, fleetName string, updateStrategyName string) UpdateStrategyId {
	return UpdateStrategyId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		FleetName:          fleetName,
		UpdateStrategyName: updateStrategyName,
	}
}

// ParseUpdateStrategyID parses 'input' into a UpdateStrategyId
func ParseUpdateStrategyID(input string) (*UpdateStrategyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UpdateStrategyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UpdateStrategyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.FleetName, ok = parsed.Parsed["fleetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fleetName", *parsed)
	}

	if id.UpdateStrategyName, ok = parsed.Parsed["updateStrategyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "updateStrategyName", *parsed)
	}

	return &id, nil
}

// ParseUpdateStrategyIDInsensitively parses 'input' case-insensitively into a UpdateStrategyId
// note: this method should only be used for API response data and not user input
func ParseUpdateStrategyIDInsensitively(input string) (*UpdateStrategyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UpdateStrategyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UpdateStrategyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.FleetName, ok = parsed.Parsed["fleetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fleetName", *parsed)
	}

	if id.UpdateStrategyName, ok = parsed.Parsed["updateStrategyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "updateStrategyName", *parsed)
	}

	return &id, nil
}

// ValidateUpdateStrategyID checks that 'input' can be parsed as a Update Strategy ID
func ValidateUpdateStrategyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUpdateStrategyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Update Strategy ID
func (id UpdateStrategyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ContainerService/fleets/%s/updateStrategies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FleetName, id.UpdateStrategyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Update Strategy ID
func (id UpdateStrategyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftContainerService", "Microsoft.ContainerService", "Microsoft.ContainerService"),
		resourceids.StaticSegment("staticFleets", "fleets", "fleets"),
		resourceids.UserSpecifiedSegment("fleetName", "fleetValue"),
		resourceids.StaticSegment("staticUpdateStrategies", "updateStrategies", "updateStrategies"),
		resourceids.UserSpecifiedSegment("updateStrategyName", "updateStrategyValue"),
	}
}

// String returns a human-readable description of this Update Strategy ID
func (id UpdateStrategyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Fleet Name: %q", id.FleetName),
		fmt.Sprintf("Update Strategy Name: %q", id.UpdateStrategyName),
	}
	return fmt.Sprintf("Update Strategy (%s)", strings.Join(components, "\n"))
}
