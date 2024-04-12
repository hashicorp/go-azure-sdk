package servicefabricschedules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ServiceFabricScheduleId{})
}

var _ resourceids.ResourceId = &ServiceFabricScheduleId{}

// ServiceFabricScheduleId is a struct representing the Resource ID for a Service Fabric Schedule
type ServiceFabricScheduleId struct {
	SubscriptionId    string
	ResourceGroupName string
	LabName           string
	UserName          string
	ServiceFabricName string
	ScheduleName      string
}

// NewServiceFabricScheduleID returns a new ServiceFabricScheduleId struct
func NewServiceFabricScheduleID(subscriptionId string, resourceGroupName string, labName string, userName string, serviceFabricName string, scheduleName string) ServiceFabricScheduleId {
	return ServiceFabricScheduleId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LabName:           labName,
		UserName:          userName,
		ServiceFabricName: serviceFabricName,
		ScheduleName:      scheduleName,
	}
}

// ParseServiceFabricScheduleID parses 'input' into a ServiceFabricScheduleId
func ParseServiceFabricScheduleID(input string) (*ServiceFabricScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServiceFabricScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServiceFabricScheduleId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServiceFabricScheduleIDInsensitively parses 'input' case-insensitively into a ServiceFabricScheduleId
// note: this method should only be used for API response data and not user input
func ParseServiceFabricScheduleIDInsensitively(input string) (*ServiceFabricScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServiceFabricScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServiceFabricScheduleId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServiceFabricScheduleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LabName, ok = input.Parsed["labName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "labName", input)
	}

	if id.UserName, ok = input.Parsed["userName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userName", input)
	}

	if id.ServiceFabricName, ok = input.Parsed["serviceFabricName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serviceFabricName", input)
	}

	if id.ScheduleName, ok = input.Parsed["scheduleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scheduleName", input)
	}

	return nil
}

// ValidateServiceFabricScheduleID checks that 'input' can be parsed as a Service Fabric Schedule ID
func ValidateServiceFabricScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServiceFabricScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Fabric Schedule ID
func (id ServiceFabricScheduleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/users/%s/serviceFabrics/%s/schedules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.UserName, id.ServiceFabricName, id.ScheduleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Fabric Schedule ID
func (id ServiceFabricScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userName", "userValue"),
		resourceids.StaticSegment("staticServiceFabrics", "serviceFabrics", "serviceFabrics"),
		resourceids.UserSpecifiedSegment("serviceFabricName", "serviceFabricValue"),
		resourceids.StaticSegment("staticSchedules", "schedules", "schedules"),
		resourceids.UserSpecifiedSegment("scheduleName", "scheduleValue"),
	}
}

// String returns a human-readable description of this Service Fabric Schedule ID
func (id ServiceFabricScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("User Name: %q", id.UserName),
		fmt.Sprintf("Service Fabric Name: %q", id.ServiceFabricName),
		fmt.Sprintf("Schedule Name: %q", id.ScheduleName),
	}
	return fmt.Sprintf("Service Fabric Schedule (%s)", strings.Join(components, "\n"))
}
