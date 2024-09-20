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
	recaser.RegisterResourceId(&ConsoleId{})
}

var _ resourceids.ResourceId = &ConsoleId{}

// ConsoleId is a struct representing the Resource ID for a Console
type ConsoleId struct {
	SubscriptionId     string
	ResourceGroupName  string
	VirtualMachineName string
	ConsoleName        string
}

// NewConsoleID returns a new ConsoleId struct
func NewConsoleID(subscriptionId string, resourceGroupName string, virtualMachineName string, consoleName string) ConsoleId {
	return ConsoleId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		VirtualMachineName: virtualMachineName,
		ConsoleName:        consoleName,
	}
}

// ParseConsoleID parses 'input' into a ConsoleId
func ParseConsoleID(input string) (*ConsoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConsoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConsoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseConsoleIDInsensitively parses 'input' case-insensitively into a ConsoleId
// note: this method should only be used for API response data and not user input
func ParseConsoleIDInsensitively(input string) (*ConsoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConsoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConsoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ConsoleId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ConsoleName, ok = input.Parsed["consoleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "consoleName", input)
	}

	return nil
}

// ValidateConsoleID checks that 'input' can be parsed as a Console ID
func ValidateConsoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConsoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Console ID
func (id ConsoleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/virtualMachines/%s/consoles/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName, id.ConsoleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Console ID
func (id ConsoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineName"),
		resourceids.StaticSegment("staticConsoles", "consoles", "consoles"),
		resourceids.UserSpecifiedSegment("consoleName", "consoleName"),
	}
}

// String returns a human-readable description of this Console ID
func (id ConsoleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Console Name: %q", id.ConsoleName),
	}
	return fmt.Sprintf("Console (%s)", strings.Join(components, "\n"))
}
