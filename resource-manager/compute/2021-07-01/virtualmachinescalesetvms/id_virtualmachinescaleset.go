// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachinescalesetvms

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VirtualMachineScaleSetId{}

// VirtualMachineScaleSetId is a struct representing the Resource ID for a Virtual Machine Scale Set
type VirtualMachineScaleSetId struct {
	SubscriptionId             string
	ResourceGroupName          string
	VirtualMachineScaleSetName string
}

// NewVirtualMachineScaleSetID returns a new VirtualMachineScaleSetId struct
func NewVirtualMachineScaleSetID(subscriptionId string, resourceGroupName string, virtualMachineScaleSetName string) VirtualMachineScaleSetId {
	return VirtualMachineScaleSetId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		VirtualMachineScaleSetName: virtualMachineScaleSetName,
	}
}

// ParseVirtualMachineScaleSetID parses 'input' into a VirtualMachineScaleSetId
func ParseVirtualMachineScaleSetID(input string) (*VirtualMachineScaleSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineScaleSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineScaleSetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineScaleSetName, ok = parsed.Parsed["virtualMachineScaleSetName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineScaleSetName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVirtualMachineScaleSetIDInsensitively parses 'input' case-insensitively into a VirtualMachineScaleSetId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineScaleSetIDInsensitively(input string) (*VirtualMachineScaleSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineScaleSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineScaleSetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineScaleSetName, ok = parsed.Parsed["virtualMachineScaleSetName"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineScaleSetName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateVirtualMachineScaleSetID checks that 'input' can be parsed as a Virtual Machine Scale Set ID
func ValidateVirtualMachineScaleSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineScaleSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine Scale Set ID
func (id VirtualMachineScaleSetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineScaleSetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine Scale Set ID
func (id VirtualMachineScaleSetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticVirtualMachineScaleSets", "virtualMachineScaleSets", "virtualMachineScaleSets"),
		resourceids.UserSpecifiedSegment("virtualMachineScaleSetName", "virtualMachineScaleSetValue"),
	}
}

// String returns a human-readable description of this Virtual Machine Scale Set ID
func (id VirtualMachineScaleSetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Scale Set Name: %q", id.VirtualMachineScaleSetName),
	}
	return fmt.Sprintf("Virtual Machine Scale Set (%s)", strings.Join(components, "\n"))
}
