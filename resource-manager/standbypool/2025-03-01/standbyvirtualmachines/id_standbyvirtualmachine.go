package standbyvirtualmachines

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&StandbyVirtualMachineId{})
}

var _ resourceids.ResourceId = &StandbyVirtualMachineId{}

// StandbyVirtualMachineId is a struct representing the Resource ID for a Standby Virtual Machine
type StandbyVirtualMachineId struct {
	SubscriptionId                string
	ResourceGroupName             string
	StandbyVirtualMachinePoolName string
	StandbyVirtualMachineName     string
}

// NewStandbyVirtualMachineID returns a new StandbyVirtualMachineId struct
func NewStandbyVirtualMachineID(subscriptionId string, resourceGroupName string, standbyVirtualMachinePoolName string, standbyVirtualMachineName string) StandbyVirtualMachineId {
	return StandbyVirtualMachineId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		StandbyVirtualMachinePoolName: standbyVirtualMachinePoolName,
		StandbyVirtualMachineName:     standbyVirtualMachineName,
	}
}

// ParseStandbyVirtualMachineID parses 'input' into a StandbyVirtualMachineId
func ParseStandbyVirtualMachineID(input string) (*StandbyVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StandbyVirtualMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StandbyVirtualMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseStandbyVirtualMachineIDInsensitively parses 'input' case-insensitively into a StandbyVirtualMachineId
// note: this method should only be used for API response data and not user input
func ParseStandbyVirtualMachineIDInsensitively(input string) (*StandbyVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StandbyVirtualMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StandbyVirtualMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *StandbyVirtualMachineId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.StandbyVirtualMachinePoolName, ok = input.Parsed["standbyVirtualMachinePoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "standbyVirtualMachinePoolName", input)
	}

	if id.StandbyVirtualMachineName, ok = input.Parsed["standbyVirtualMachineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "standbyVirtualMachineName", input)
	}

	return nil
}

// ValidateStandbyVirtualMachineID checks that 'input' can be parsed as a Standby Virtual Machine ID
func ValidateStandbyVirtualMachineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStandbyVirtualMachineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Standby Virtual Machine ID
func (id StandbyVirtualMachineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/%s/standbyVirtualMachines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StandbyVirtualMachinePoolName, id.StandbyVirtualMachineName)
}

// Segments returns a slice of Resource ID Segments which comprise this Standby Virtual Machine ID
func (id StandbyVirtualMachineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStandbyPool", "Microsoft.StandbyPool", "Microsoft.StandbyPool"),
		resourceids.StaticSegment("staticStandbyVirtualMachinePools", "standbyVirtualMachinePools", "standbyVirtualMachinePools"),
		resourceids.UserSpecifiedSegment("standbyVirtualMachinePoolName", "standbyVirtualMachinePoolName"),
		resourceids.StaticSegment("staticStandbyVirtualMachines", "standbyVirtualMachines", "standbyVirtualMachines"),
		resourceids.UserSpecifiedSegment("standbyVirtualMachineName", "standbyVirtualMachineName"),
	}
}

// String returns a human-readable description of this Standby Virtual Machine ID
func (id StandbyVirtualMachineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Standby Virtual Machine Pool Name: %q", id.StandbyVirtualMachinePoolName),
		fmt.Sprintf("Standby Virtual Machine Name: %q", id.StandbyVirtualMachineName),
	}
	return fmt.Sprintf("Standby Virtual Machine (%s)", strings.Join(components, "\n"))
}
