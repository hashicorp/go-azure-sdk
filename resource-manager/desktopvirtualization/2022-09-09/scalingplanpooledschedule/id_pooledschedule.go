package scalingplanpooledschedule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PooledScheduleId{})
}

var _ resourceids.ResourceId = &PooledScheduleId{}

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
	parser := resourceids.NewParserFromResourceIdType(&PooledScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PooledScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePooledScheduleIDInsensitively parses 'input' case-insensitively into a PooledScheduleId
// note: this method should only be used for API response data and not user input
func ParsePooledScheduleIDInsensitively(input string) (*PooledScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PooledScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PooledScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PooledScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ScalingPlanName, ok = input.Parsed["scalingPlanName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scalingPlanName", input)
	}

	if id.PooledScheduleName, ok = input.Parsed["pooledScheduleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pooledScheduleName", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("scalingPlanName", "scalingPlanName"),
		resourceids.StaticSegment("staticPooledSchedules", "pooledSchedules", "pooledSchedules"),
		resourceids.UserSpecifiedSegment("pooledScheduleName", "pooledScheduleName"),
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
