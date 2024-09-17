package guestconfigurationconnectedvmwarevsphereassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ProviderVirtualMachineId{})
}

var _ resourceids.ResourceId = &ProviderVirtualMachineId{}

// ProviderVirtualMachineId is a struct representing the Resource ID for a Provider Virtual Machine
type ProviderVirtualMachineId struct {
	SubscriptionId     string
	ResourceGroupName  string
	VirtualMachineName string
}

// NewProviderVirtualMachineID returns a new ProviderVirtualMachineId struct
func NewProviderVirtualMachineID(subscriptionId string, resourceGroupName string, virtualMachineName string) ProviderVirtualMachineId {
	return ProviderVirtualMachineId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		VirtualMachineName: virtualMachineName,
	}
}

// ParseProviderVirtualMachineID parses 'input' into a ProviderVirtualMachineId
func ParseProviderVirtualMachineID(input string) (*ProviderVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderVirtualMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderVirtualMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviderVirtualMachineIDInsensitively parses 'input' case-insensitively into a ProviderVirtualMachineId
// note: this method should only be used for API response data and not user input
func ParseProviderVirtualMachineIDInsensitively(input string) (*ProviderVirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProviderVirtualMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderVirtualMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProviderVirtualMachineId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.VirtualMachineName, ok = input.Parsed["virtualMachineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", input)
	}

	return nil
}

// ValidateProviderVirtualMachineID checks that 'input' can be parsed as a Provider Virtual Machine ID
func ValidateProviderVirtualMachineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderVirtualMachineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Virtual Machine ID
func (id ProviderVirtualMachineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Virtual Machine ID
func (id ProviderVirtualMachineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineValue"),
	}
}

// String returns a human-readable description of this Provider Virtual Machine ID
func (id ProviderVirtualMachineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
	}
	return fmt.Sprintf("Provider Virtual Machine (%s)", strings.Join(components, "\n"))
}
