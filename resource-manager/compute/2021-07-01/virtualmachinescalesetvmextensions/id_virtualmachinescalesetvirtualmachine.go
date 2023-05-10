package virtualmachinescalesetvmextensions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VirtualMachineScaleSetVirtualMachineId{}

// VirtualMachineScaleSetVirtualMachineId is a struct representing the Resource ID for a Virtual Machine Scale Set Virtual Machine
type VirtualMachineScaleSetVirtualMachineId struct {
	SubscriptionId             string
	ResourceGroupName          string
	VirtualMachineScaleSetName string
	InstanceId                 string
}

// NewVirtualMachineScaleSetVirtualMachineID returns a new VirtualMachineScaleSetVirtualMachineId struct
func NewVirtualMachineScaleSetVirtualMachineID(subscriptionId string, resourceGroupName string, virtualMachineScaleSetName string, instanceId string) VirtualMachineScaleSetVirtualMachineId {
	return VirtualMachineScaleSetVirtualMachineId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		VirtualMachineScaleSetName: virtualMachineScaleSetName,
		InstanceId:                 instanceId,
	}
}

// ParseVirtualMachineScaleSetVirtualMachineID parses 'input' into a VirtualMachineScaleSetVirtualMachineId
func ParseVirtualMachineScaleSetVirtualMachineID(input string) (*VirtualMachineScaleSetVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineScaleSetVirtualMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineScaleSetVirtualMachineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineScaleSetName, ok = parsed.Parsed["virtualMachineScaleSetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineScaleSetName", *parsed)
	}

	if id.InstanceId, ok = parsed.Parsed["instanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "instanceId", *parsed)
	}

	return &id, nil
}

// ParseVirtualMachineScaleSetVirtualMachineIDInsensitively parses 'input' case-insensitively into a VirtualMachineScaleSetVirtualMachineId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineScaleSetVirtualMachineIDInsensitively(input string) (*VirtualMachineScaleSetVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineScaleSetVirtualMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineScaleSetVirtualMachineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineScaleSetName, ok = parsed.Parsed["virtualMachineScaleSetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineScaleSetName", *parsed)
	}

	if id.InstanceId, ok = parsed.Parsed["instanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "instanceId", *parsed)
	}

	return &id, nil
}

// ValidateVirtualMachineScaleSetVirtualMachineID checks that 'input' can be parsed as a Virtual Machine Scale Set Virtual Machine ID
func ValidateVirtualMachineScaleSetVirtualMachineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineScaleSetVirtualMachineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Scale Set Virtual Machine ID
func (id VirtualMachineScaleSetVirtualMachineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s/virtualMachines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineScaleSetName, id.InstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Scale Set Virtual Machine ID
func (id VirtualMachineScaleSetVirtualMachineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticVirtualMachineScaleSets", "virtualMachineScaleSets", "virtualMachineScaleSets"),
		resourceids.UserSpecifiedSegment("virtualMachineScaleSetName", "virtualMachineScaleSetValue"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("instanceId", "instanceIdValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Scale Set Virtual Machine ID
func (id VirtualMachineScaleSetVirtualMachineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Scale Set Name: %q", id.VirtualMachineScaleSetName),
		fmt.Sprintf("Instance: %q", id.InstanceId),
	}
	return fmt.Sprintf("Virtual Machine Scale Set Virtual Machine (%s)", strings.Join(components, "\n"))
}
