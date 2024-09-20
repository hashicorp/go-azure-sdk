package virtualmachineschedules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&VirtualMachineScheduleId{})
}

var _ resourceids.ResourceId = &VirtualMachineScheduleId{}

// VirtualMachineScheduleId is a struct representing the Resource ID for a Virtual Machine Schedule
type VirtualMachineScheduleId struct {
	SubscriptionId     string
	ResourceGroupName  string
	LabName            string
	VirtualMachineName string
	ScheduleName       string
}

// NewVirtualMachineScheduleID returns a new VirtualMachineScheduleId struct
func NewVirtualMachineScheduleID(subscriptionId string, resourceGroupName string, labName string, virtualMachineName string, scheduleName string) VirtualMachineScheduleId {
	return VirtualMachineScheduleId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		LabName:            labName,
		VirtualMachineName: virtualMachineName,
		ScheduleName:       scheduleName,
	}
}

// ParseVirtualMachineScheduleID parses 'input' into a VirtualMachineScheduleId
func ParseVirtualMachineScheduleID(input string) (*VirtualMachineScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualMachineScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualMachineScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVirtualMachineScheduleIDInsensitively parses 'input' case-insensitively into a VirtualMachineScheduleId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineScheduleIDInsensitively(input string) (*VirtualMachineScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualMachineScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualMachineScheduleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VirtualMachineScheduleId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.VirtualMachineName, ok = input.Parsed["virtualMachineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", input)
	}

	if id.ScheduleName, ok = input.Parsed["scheduleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scheduleName", input)
	}

	return nil
}

// ValidateVirtualMachineScheduleID checks that 'input' can be parsed as a Virtual Machine Schedule ID
func ValidateVirtualMachineScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Schedule ID
func (id VirtualMachineScheduleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/virtualMachines/%s/schedules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.VirtualMachineName, id.ScheduleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Schedule ID
func (id VirtualMachineScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labName"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineName"),
		resourceids.StaticSegment("staticSchedules", "schedules", "schedules"),
		resourceids.UserSpecifiedSegment("scheduleName", "name"),
	}
}

// String returns a human-readable description of this Virtual Machine Schedule ID
func (id VirtualMachineScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Schedule Name: %q", id.ScheduleName),
	}
	return fmt.Sprintf("Virtual Machine Schedule (%s)", strings.Join(components, "\n"))
}
