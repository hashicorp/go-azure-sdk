package scalingplanpersonalschedule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PersonalScheduleId{})
}

var _ resourceids.ResourceId = &PersonalScheduleId{}

// PersonalScheduleId is a struct representing the Resource ID for a Personal Schedule
type PersonalScheduleId struct {
	SubscriptionId       string
	ResourceGroupName    string
	ScalingPlanName      string
	PersonalScheduleName string
}

// NewPersonalScheduleID returns a new PersonalScheduleId struct
func NewPersonalScheduleID(subscriptionId string, resourceGroupName string, scalingPlanName string, personalScheduleName string) PersonalScheduleId {
	return PersonalScheduleId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		ScalingPlanName:      scalingPlanName,
		PersonalScheduleName: personalScheduleName,
	}
}

// ParsePersonalScheduleID parses 'input' into a PersonalScheduleId
func ParsePersonalScheduleID(input string) (*PersonalScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PersonalScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PersonalScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePersonalScheduleIDInsensitively parses 'input' case-insensitively into a PersonalScheduleId
// note: this method should only be used for API response data and not user input
func ParsePersonalScheduleIDInsensitively(input string) (*PersonalScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PersonalScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PersonalScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PersonalScheduleId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PersonalScheduleName, ok = input.Parsed["personalScheduleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "personalScheduleName", input)
	}

	return nil
}

// ValidatePersonalScheduleID checks that 'input' can be parsed as a Personal Schedule ID
func ValidatePersonalScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePersonalScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Personal Schedule ID
func (id PersonalScheduleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/scalingPlans/%s/personalSchedules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ScalingPlanName, id.PersonalScheduleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Personal Schedule ID
func (id PersonalScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDesktopVirtualization", "Microsoft.DesktopVirtualization", "Microsoft.DesktopVirtualization"),
		resourceids.StaticSegment("staticScalingPlans", "scalingPlans", "scalingPlans"),
		resourceids.UserSpecifiedSegment("scalingPlanName", "scalingPlanName"),
		resourceids.StaticSegment("staticPersonalSchedules", "personalSchedules", "personalSchedules"),
		resourceids.UserSpecifiedSegment("personalScheduleName", "personalScheduleName"),
	}
}

// String returns a human-readable description of this Personal Schedule ID
func (id PersonalScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Scaling Plan Name: %q", id.ScalingPlanName),
		fmt.Sprintf("Personal Schedule Name: %q", id.PersonalScheduleName),
	}
	return fmt.Sprintf("Personal Schedule (%s)", strings.Join(components, "\n"))
}
