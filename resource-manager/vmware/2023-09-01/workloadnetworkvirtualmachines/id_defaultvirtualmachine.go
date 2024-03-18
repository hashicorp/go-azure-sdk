package workloadnetworkvirtualmachines

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DefaultVirtualMachineId{}

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
	parser := resourceids.NewParserFromResourceIdType(&DefaultVirtualMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DefaultVirtualMachineId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDefaultVirtualMachineIDInsensitively parses 'input' case-insensitively into a DefaultVirtualMachineId
// note: this method should only be used for API response data and not user input
func ParseDefaultVirtualMachineIDInsensitively(input string) (*DefaultVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DefaultVirtualMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DefaultVirtualMachineId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DefaultVirtualMachineId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PrivateCloudName, ok = input.Parsed["privateCloudName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateCloudName", input)
	}

	if id.VirtualMachineId, ok = input.Parsed["virtualMachineId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineId", input)
	}

	return nil
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
