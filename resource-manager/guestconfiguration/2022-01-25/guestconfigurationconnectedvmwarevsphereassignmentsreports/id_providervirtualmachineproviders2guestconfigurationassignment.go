package guestconfigurationconnectedvmwarevsphereassignmentsreports

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProviderVirtualMachineProviders2GuestConfigurationAssignmentId{}

// ProviderVirtualMachineProviders2GuestConfigurationAssignmentId is a struct representing the Resource ID for a Provider Virtual Machine Providers 2 Guest Configuration Assignment
type ProviderVirtualMachineProviders2GuestConfigurationAssignmentId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	VirtualMachineName               string
	GuestConfigurationAssignmentName string
}

// NewProviderVirtualMachineProviders2GuestConfigurationAssignmentID returns a new ProviderVirtualMachineProviders2GuestConfigurationAssignmentId struct
func NewProviderVirtualMachineProviders2GuestConfigurationAssignmentID(subscriptionId string, resourceGroupName string, virtualMachineName string, guestConfigurationAssignmentName string) ProviderVirtualMachineProviders2GuestConfigurationAssignmentId {
	return ProviderVirtualMachineProviders2GuestConfigurationAssignmentId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		VirtualMachineName:               virtualMachineName,
		GuestConfigurationAssignmentName: guestConfigurationAssignmentName,
	}
}

// ParseProviderVirtualMachineProviders2GuestConfigurationAssignmentID parses 'input' into a ProviderVirtualMachineProviders2GuestConfigurationAssignmentId
func ParseProviderVirtualMachineProviders2GuestConfigurationAssignmentID(input string) (*ProviderVirtualMachineProviders2GuestConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderVirtualMachineProviders2GuestConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderVirtualMachineProviders2GuestConfigurationAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.GuestConfigurationAssignmentName, ok = parsed.Parsed["guestConfigurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "guestConfigurationAssignmentName", *parsed)
	}

	return &id, nil
}

// ParseProviderVirtualMachineProviders2GuestConfigurationAssignmentIDInsensitively parses 'input' case-insensitively into a ProviderVirtualMachineProviders2GuestConfigurationAssignmentId
// note: this method should only be used for API response data and not user input
func ParseProviderVirtualMachineProviders2GuestConfigurationAssignmentIDInsensitively(input string) (*ProviderVirtualMachineProviders2GuestConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderVirtualMachineProviders2GuestConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderVirtualMachineProviders2GuestConfigurationAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.GuestConfigurationAssignmentName, ok = parsed.Parsed["guestConfigurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "guestConfigurationAssignmentName", *parsed)
	}

	return &id, nil
}

// ValidateProviderVirtualMachineProviders2GuestConfigurationAssignmentID checks that 'input' can be parsed as a Provider Virtual Machine Providers 2 Guest Configuration Assignment ID
func ValidateProviderVirtualMachineProviders2GuestConfigurationAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderVirtualMachineProviders2GuestConfigurationAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Virtual Machine Providers 2 Guest Configuration Assignment ID
func (id ProviderVirtualMachineProviders2GuestConfigurationAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/%s/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName, id.GuestConfigurationAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Virtual Machine Providers 2 Guest Configuration Assignment ID
func (id ProviderVirtualMachineProviders2GuestConfigurationAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftGuestConfiguration", "Microsoft.GuestConfiguration", "Microsoft.GuestConfiguration"),
		resourceids.StaticSegment("staticGuestConfigurationAssignments", "guestConfigurationAssignments", "guestConfigurationAssignments"),
		resourceids.UserSpecifiedSegment("guestConfigurationAssignmentName", "guestConfigurationAssignmentValue"),
	}
}

// String returns a human-readable description of this Provider Virtual Machine Providers 2 Guest Configuration Assignment ID
func (id ProviderVirtualMachineProviders2GuestConfigurationAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Guest Configuration Assignment Name: %q", id.GuestConfigurationAssignmentName),
	}
	return fmt.Sprintf("Provider Virtual Machine Providers 2 Guest Configuration Assignment (%s)", strings.Join(components, "\n"))
}
