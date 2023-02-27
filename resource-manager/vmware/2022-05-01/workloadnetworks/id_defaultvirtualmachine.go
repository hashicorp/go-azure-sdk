// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DefaultVirtualMachineId{}

// DefaultVirtualMachineId is a struct representing the Resource ID for a Default Virtual Machine
type DefaultVirtualMachineId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	VirtualMachineId  string
}

// NewDefaultVirtualMachineID returns a new DefaultVirtualMachineId struct
func NewDefaultVirtualMachineID(subscriptionId string, resourceGroupName string, privateCloudName string, virtualMachineId string) DefaultVirtualMachineId {
	return DefaultVirtualMachineId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		VirtualMachineId:  virtualMachineId,
	}
}

// ParseDefaultVirtualMachineID parses 'input' into a DefaultVirtualMachineId
func ParseDefaultVirtualMachineID(input string) (*DefaultVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(DefaultVirtualMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DefaultVirtualMachineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineId, ok = parsed.Parsed["virtualMachineId"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDefaultVirtualMachineIDInsensitively parses 'input' case-insensitively into a DefaultVirtualMachineId
// note: this method should only be used for API response data and not user input
func ParseDefaultVirtualMachineIDInsensitively(input string) (*DefaultVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(DefaultVirtualMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DefaultVirtualMachineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineId, ok = parsed.Parsed["virtualMachineId"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDefaultVirtualMachineID checks that 'input' can be parsed as a Default Virtual Machine ID
func ValidateDefaultVirtualMachineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDefaultVirtualMachineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Default Virtual Machine ID
func (id DefaultVirtualMachineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/virtualMachines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.VirtualMachineId)
}

// Segments returns a slice of Resource ID Segments which comprise this Default Virtual Machine ID
func (id DefaultVirtualMachineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticWorkloadNetworks", "workloadNetworks", "workloadNetworks"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineId", "virtualMachineIdValue"),
	}
}

// String returns a human-readable description of this Default Virtual Machine ID
func (id DefaultVirtualMachineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Virtual Machine: %q", id.VirtualMachineId),
	}
	return fmt.Sprintf("Default Virtual Machine (%s)", strings.Join(components, "\n"))
}
