package guestagents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&GuestAgentId{})
}

var _ resourceids.ResourceId = &GuestAgentId{}

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
	parser := resourceids.NewParserFromResourceIdType(&GuestAgentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GuestAgentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGuestAgentIDInsensitively parses 'input' case-insensitively into a GuestAgentId
// note: this method should only be used for API response data and not user input
func ParseGuestAgentIDInsensitively(input string) (*GuestAgentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GuestAgentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GuestAgentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GuestAgentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.GuestAgentName, ok = input.Parsed["guestAgentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "guestAgentName", input)
	}

	return nil
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
