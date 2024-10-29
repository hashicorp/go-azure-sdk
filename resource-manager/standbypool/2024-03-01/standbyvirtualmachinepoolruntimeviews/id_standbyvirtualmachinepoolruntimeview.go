package standbyvirtualmachinepoolruntimeviews

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&StandbyVirtualMachinePoolRuntimeViewId{})
}

var _ resourceids.ResourceId = &StandbyVirtualMachinePoolRuntimeViewId{}

// StandbyVirtualMachinePoolRuntimeViewId is a struct representing the Resource ID for a Standby Virtual Machine Pool Runtime View
type StandbyVirtualMachinePoolRuntimeViewId struct {
	SubscriptionId                string
	ResourceGroupName             string
	StandbyVirtualMachinePoolName string
	RuntimeViewName               string
}

// NewStandbyVirtualMachinePoolRuntimeViewID returns a new StandbyVirtualMachinePoolRuntimeViewId struct
func NewStandbyVirtualMachinePoolRuntimeViewID(subscriptionId string, resourceGroupName string, standbyVirtualMachinePoolName string, runtimeViewName string) StandbyVirtualMachinePoolRuntimeViewId {
	return StandbyVirtualMachinePoolRuntimeViewId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		StandbyVirtualMachinePoolName: standbyVirtualMachinePoolName,
		RuntimeViewName:               runtimeViewName,
	}
}

// ParseStandbyVirtualMachinePoolRuntimeViewID parses 'input' into a StandbyVirtualMachinePoolRuntimeViewId
func ParseStandbyVirtualMachinePoolRuntimeViewID(input string) (*StandbyVirtualMachinePoolRuntimeViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StandbyVirtualMachinePoolRuntimeViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StandbyVirtualMachinePoolRuntimeViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseStandbyVirtualMachinePoolRuntimeViewIDInsensitively parses 'input' case-insensitively into a StandbyVirtualMachinePoolRuntimeViewId
// note: this method should only be used for API response data and not user input
func ParseStandbyVirtualMachinePoolRuntimeViewIDInsensitively(input string) (*StandbyVirtualMachinePoolRuntimeViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StandbyVirtualMachinePoolRuntimeViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StandbyVirtualMachinePoolRuntimeViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *StandbyVirtualMachinePoolRuntimeViewId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RuntimeViewName, ok = input.Parsed["runtimeViewName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runtimeViewName", input)
	}

	return nil
}

// ValidateStandbyVirtualMachinePoolRuntimeViewID checks that 'input' can be parsed as a Standby Virtual Machine Pool Runtime View ID
func ValidateStandbyVirtualMachinePoolRuntimeViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStandbyVirtualMachinePoolRuntimeViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Standby Virtual Machine Pool Runtime View ID
func (id StandbyVirtualMachinePoolRuntimeViewId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/%s/runtimeViews/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StandbyVirtualMachinePoolName, id.RuntimeViewName)
}

// Segments returns a slice of Resource ID Segments which comprise this Standby Virtual Machine Pool Runtime View ID
func (id StandbyVirtualMachinePoolRuntimeViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStandbyPool", "Microsoft.StandbyPool", "Microsoft.StandbyPool"),
		resourceids.StaticSegment("staticStandbyVirtualMachinePools", "standbyVirtualMachinePools", "standbyVirtualMachinePools"),
		resourceids.UserSpecifiedSegment("standbyVirtualMachinePoolName", "standbyVirtualMachinePoolName"),
		resourceids.StaticSegment("staticRuntimeViews", "runtimeViews", "runtimeViews"),
		resourceids.UserSpecifiedSegment("runtimeViewName", "runtimeViewName"),
	}
}

// String returns a human-readable description of this Standby Virtual Machine Pool Runtime View ID
func (id StandbyVirtualMachinePoolRuntimeViewId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Standby Virtual Machine Pool Name: %q", id.StandbyVirtualMachinePoolName),
		fmt.Sprintf("Runtime View Name: %q", id.RuntimeViewName),
	}
	return fmt.Sprintf("Standby Virtual Machine Pool Runtime View (%s)", strings.Join(components, "\n"))
}
