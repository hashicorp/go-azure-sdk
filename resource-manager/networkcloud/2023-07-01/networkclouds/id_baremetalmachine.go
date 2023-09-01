package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = BareMetalMachineId{}

// BareMetalMachineId is a struct representing the Resource ID for a Bare Metal Machine
type BareMetalMachineId struct {
	SubscriptionId       string
	ResourceGroupName    string
	BareMetalMachineName string
}

// NewBareMetalMachineID returns a new BareMetalMachineId struct
func NewBareMetalMachineID(subscriptionId string, resourceGroupName string, bareMetalMachineName string) BareMetalMachineId {
	return BareMetalMachineId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		BareMetalMachineName: bareMetalMachineName,
	}
}

// ParseBareMetalMachineID parses 'input' into a BareMetalMachineId
func ParseBareMetalMachineID(input string) (*BareMetalMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(BareMetalMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BareMetalMachineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.BareMetalMachineName, ok = parsed.Parsed["bareMetalMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "bareMetalMachineName", *parsed)
	}

	return &id, nil
}

// ParseBareMetalMachineIDInsensitively parses 'input' case-insensitively into a BareMetalMachineId
// note: this method should only be used for API response data and not user input
func ParseBareMetalMachineIDInsensitively(input string) (*BareMetalMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(BareMetalMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BareMetalMachineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.BareMetalMachineName, ok = parsed.Parsed["bareMetalMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "bareMetalMachineName", *parsed)
	}

	return &id, nil
}

// ValidateBareMetalMachineID checks that 'input' can be parsed as a Bare Metal Machine ID
func ValidateBareMetalMachineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBareMetalMachineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Bare Metal Machine ID
func (id BareMetalMachineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/bareMetalMachines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.BareMetalMachineName)
}

// Segments returns a slice of Resource ID Segments which comprise this Bare Metal Machine ID
func (id BareMetalMachineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticBareMetalMachines", "bareMetalMachines", "bareMetalMachines"),
		resourceids.UserSpecifiedSegment("bareMetalMachineName", "bareMetalMachineValue"),
	}
}

// String returns a human-readable description of this Bare Metal Machine ID
func (id BareMetalMachineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Bare Metal Machine Name: %q", id.BareMetalMachineName),
	}
	return fmt.Sprintf("Bare Metal Machine (%s)", strings.Join(components, "\n"))
}
