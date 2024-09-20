package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BareMetalMachineId{})
}

var _ resourceids.ResourceId = &BareMetalMachineId{}

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
	parser := resourceids.NewParserFromResourceIdType(&BareMetalMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BareMetalMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBareMetalMachineIDInsensitively parses 'input' case-insensitively into a BareMetalMachineId
// note: this method should only be used for API response data and not user input
func ParseBareMetalMachineIDInsensitively(input string) (*BareMetalMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BareMetalMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BareMetalMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BareMetalMachineId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.BareMetalMachineName, ok = input.Parsed["bareMetalMachineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bareMetalMachineName", input)
	}

	return nil
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
		resourceids.UserSpecifiedSegment("bareMetalMachineName", "bareMetalMachineName"),
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
