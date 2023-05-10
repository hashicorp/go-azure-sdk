package scalingplanpooledschedule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PooledScheduleId{}

// PooledScheduleId is a struct representing the Resource ID for a Pooled Schedule
type PooledScheduleId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ScalingPlanName    string
	PooledScheduleName string
}

// NewPooledScheduleID returns a new PooledScheduleId struct
func NewPooledScheduleID(subscriptionId string, resourceGroupName string, scalingPlanName string, pooledScheduleName string) PooledScheduleId {
	return PooledScheduleId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ScalingPlanName:    scalingPlanName,
		PooledScheduleName: pooledScheduleName,
	}
}

// ParsePooledScheduleID parses 'input' into a PooledScheduleId
func ParsePooledScheduleID(input string) (*PooledScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(PooledScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PooledScheduleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ScalingPlanName, ok = parsed.Parsed["scalingPlanName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scalingPlanName", *parsed)
	}

	if id.PooledScheduleName, ok = parsed.Parsed["pooledScheduleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "pooledScheduleName", *parsed)
	}

	return &id, nil
}

// ParsePooledScheduleIDInsensitively parses 'input' case-insensitively into a PooledScheduleId
// note: this method should only be used for API response data and not user input
func ParsePooledScheduleIDInsensitively(input string) (*PooledScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(PooledScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PooledScheduleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ScalingPlanName, ok = parsed.Parsed["scalingPlanName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scalingPlanName", *parsed)
	}

	if id.PooledScheduleName, ok = parsed.Parsed["pooledScheduleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "pooledScheduleName", *parsed)
	}

	return &id, nil
}

// ValidatePooledScheduleID checks that 'input' can be parsed as a Pooled Schedule ID
func ValidatePooledScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePooledScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pooled Schedule ID
func (id PooledScheduleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/scalingPlans/%s/pooledSchedules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ScalingPlanName, id.PooledScheduleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Pooled Schedule ID
func (id PooledScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDesktopVirtualization", "Microsoft.DesktopVirtualization", "Microsoft.DesktopVirtualization"),
		resourceids.StaticSegment("staticScalingPlans", "scalingPlans", "scalingPlans"),
		resourceids.UserSpecifiedSegment("scalingPlanName", "scalingPlanValue"),
		resourceids.StaticSegment("staticPooledSchedules", "pooledSchedules", "pooledSchedules"),
		resourceids.UserSpecifiedSegment("pooledScheduleName", "pooledScheduleValue"),
	}
}

// String returns a human-readable description of this Pooled Schedule ID
func (id PooledScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Scaling Plan Name: %q", id.ScalingPlanName),
		fmt.Sprintf("Pooled Schedule Name: %q", id.PooledScheduleName),
	}
	return fmt.Sprintf("Pooled Schedule (%s)", strings.Join(components, "\n"))
}
