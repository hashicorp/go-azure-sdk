package guestagents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GuestAgentId{}

// GuestAgentId is a struct representing the Resource ID for a Guest Agent
type GuestAgentId struct {
	SubscriptionId     string
	ResourceGroupName  string
	VirtualMachineName string
	GuestAgentName     string
}

// NewGuestAgentID returns a new GuestAgentId struct
func NewGuestAgentID(subscriptionId string, resourceGroupName string, virtualMachineName string, guestAgentName string) GuestAgentId {
	return GuestAgentId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		VirtualMachineName: virtualMachineName,
		GuestAgentName:     guestAgentName,
	}
}

// ParseGuestAgentID parses 'input' into a GuestAgentId
func ParseGuestAgentID(input string) (*GuestAgentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GuestAgentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GuestAgentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.GuestAgentName, ok = parsed.Parsed["guestAgentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "guestAgentName", *parsed)
	}

	return &id, nil
}

// ParseGuestAgentIDInsensitively parses 'input' case-insensitively into a GuestAgentId
// note: this method should only be used for API response data and not user input
func ParseGuestAgentIDInsensitively(input string) (*GuestAgentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GuestAgentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GuestAgentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.GuestAgentName, ok = parsed.Parsed["guestAgentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "guestAgentName", *parsed)
	}

	return &id, nil
}

// ValidateGuestAgentID checks that 'input' can be parsed as a Guest Agent ID
func ValidateGuestAgentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGuestAgentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Guest Agent ID
func (id GuestAgentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/%s/guestAgents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName, id.GuestAgentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Guest Agent ID
func (id GuestAgentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineValue"),
		resourceids.StaticSegment("staticGuestAgents", "guestAgents", "guestAgents"),
		resourceids.UserSpecifiedSegment("guestAgentName", "guestAgentValue"),
	}
}

// String returns a human-readable description of this Guest Agent ID
func (id GuestAgentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Guest Agent Name: %q", id.GuestAgentName),
	}
	return fmt.Sprintf("Guest Agent (%s)", strings.Join(components, "\n"))
}
