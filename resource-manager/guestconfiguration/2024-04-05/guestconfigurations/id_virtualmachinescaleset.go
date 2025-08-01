package guestconfigurations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&VirtualMachineScaleSetId{})
}

var _ resourceids.ResourceId = &VirtualMachineScaleSetId{}

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
	parser := resourceids.NewParserFromResourceIdType(&VirtualMachineScaleSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualMachineScaleSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVirtualMachineScaleSetIDInsensitively parses 'input' case-insensitively into a VirtualMachineScaleSetId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineScaleSetIDInsensitively(input string) (*VirtualMachineScaleSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualMachineScaleSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualMachineScaleSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VirtualMachineScaleSetId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.VirtualMachineScaleSetName, ok = input.Parsed["virtualMachineScaleSetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineScaleSetName", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("virtualMachineScaleSetName", "virtualMachineScaleSetName"),
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
