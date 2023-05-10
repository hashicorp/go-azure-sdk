package guestconfigurationassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GuestConfigurationAssignmentId{}

// GuestConfigurationAssignmentId is a struct representing the Resource ID for a Guest Configuration Assignment
type GuestConfigurationAssignmentId struct {
	SubscriptionId                   string
	ResourceGroupName                string
	VirtualMachineScaleSetName       string
	GuestConfigurationAssignmentName string
}

// NewGuestConfigurationAssignmentID returns a new GuestConfigurationAssignmentId struct
func NewGuestConfigurationAssignmentID(subscriptionId string, resourceGroupName string, virtualMachineScaleSetName string, guestConfigurationAssignmentName string) GuestConfigurationAssignmentId {
	return GuestConfigurationAssignmentId{
		SubscriptionId:                   subscriptionId,
		ResourceGroupName:                resourceGroupName,
		VirtualMachineScaleSetName:       virtualMachineScaleSetName,
		GuestConfigurationAssignmentName: guestConfigurationAssignmentName,
	}
}

// ParseGuestConfigurationAssignmentID parses 'input' into a GuestConfigurationAssignmentId
func ParseGuestConfigurationAssignmentID(input string) (*GuestConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GuestConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GuestConfigurationAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineScaleSetName, ok = parsed.Parsed["virtualMachineScaleSetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineScaleSetName", *parsed)
	}

	if id.GuestConfigurationAssignmentName, ok = parsed.Parsed["guestConfigurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "guestConfigurationAssignmentName", *parsed)
	}

	return &id, nil
}

// ParseGuestConfigurationAssignmentIDInsensitively parses 'input' case-insensitively into a GuestConfigurationAssignmentId
// note: this method should only be used for API response data and not user input
func ParseGuestConfigurationAssignmentIDInsensitively(input string) (*GuestConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GuestConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GuestConfigurationAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineScaleSetName, ok = parsed.Parsed["virtualMachineScaleSetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineScaleSetName", *parsed)
	}

	if id.GuestConfigurationAssignmentName, ok = parsed.Parsed["guestConfigurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "guestConfigurationAssignmentName", *parsed)
	}

	return &id, nil
}

// ValidateGuestConfigurationAssignmentID checks that 'input' can be parsed as a Guest Configuration Assignment ID
func ValidateGuestConfigurationAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGuestConfigurationAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Guest Configuration Assignment ID
func (id GuestConfigurationAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Compute/virtualMachineScaleSets/%s/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineScaleSetName, id.GuestConfigurationAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Guest Configuration Assignment ID
func (id GuestConfigurationAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticVirtualMachineScaleSets", "virtualMachineScaleSets", "virtualMachineScaleSets"),
		resourceids.UserSpecifiedSegment("virtualMachineScaleSetName", "virtualMachineScaleSetValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftGuestConfiguration", "Microsoft.GuestConfiguration", "Microsoft.GuestConfiguration"),
		resourceids.StaticSegment("staticGuestConfigurationAssignments", "guestConfigurationAssignments", "guestConfigurationAssignments"),
		resourceids.UserSpecifiedSegment("guestConfigurationAssignmentName", "guestConfigurationAssignmentValue"),
	}
}

// String returns a human-readable description of this Guest Configuration Assignment ID
func (id GuestConfigurationAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Scale Set Name: %q", id.VirtualMachineScaleSetName),
		fmt.Sprintf("Guest Configuration Assignment Name: %q", id.GuestConfigurationAssignmentName),
	}
	return fmt.Sprintf("Guest Configuration Assignment (%s)", strings.Join(components, "\n"))
}
