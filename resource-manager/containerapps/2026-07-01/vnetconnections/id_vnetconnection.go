package vnetconnections

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&VnetConnectionId{})
}

var _ resourceids.ResourceId = &VnetConnectionId{}

// VnetConnectionId is a struct representing the Resource ID for a Vnet Connection
type VnetConnectionId struct {
	SubscriptionId     string
	ResourceGroupName  string
	SandboxGroupName   string
	VnetConnectionName string
}

// NewVnetConnectionID returns a new VnetConnectionId struct
func NewVnetConnectionID(subscriptionId string, resourceGroupName string, sandboxGroupName string, vnetConnectionName string) VnetConnectionId {
	return VnetConnectionId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		SandboxGroupName:   sandboxGroupName,
		VnetConnectionName: vnetConnectionName,
	}
}

// ParseVnetConnectionID parses 'input' into a VnetConnectionId
func ParseVnetConnectionID(input string) (*VnetConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VnetConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VnetConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVnetConnectionIDInsensitively parses 'input' case-insensitively into a VnetConnectionId
// note: this method should only be used for API response data and not user input
func ParseVnetConnectionIDInsensitively(input string) (*VnetConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VnetConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VnetConnectionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VnetConnectionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SandboxGroupName, ok = input.Parsed["sandboxGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sandboxGroupName", input)
	}

	if id.VnetConnectionName, ok = input.Parsed["vnetConnectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "vnetConnectionName", input)
	}

	return nil
}

// ValidateVnetConnectionID checks that 'input' can be parsed as a Vnet Connection ID
func ValidateVnetConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVnetConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Vnet Connection ID
func (id VnetConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/sandboxGroups/%s/vnetConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SandboxGroupName, id.VnetConnectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Vnet Connection ID
func (id VnetConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticSandboxGroups", "sandboxGroups", "sandboxGroups"),
		resourceids.UserSpecifiedSegment("sandboxGroupName", "sandboxGroupName"),
		resourceids.StaticSegment("staticVnetConnections", "vnetConnections", "vnetConnections"),
		resourceids.UserSpecifiedSegment("vnetConnectionName", "vnetConnectionName"),
	}
}

// String returns a human-readable description of this Vnet Connection ID
func (id VnetConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Sandbox Group Name: %q", id.SandboxGroupName),
		fmt.Sprintf("Vnet Connection Name: %q", id.VnetConnectionName),
	}
	return fmt.Sprintf("Vnet Connection (%s)", strings.Join(components, "\n"))
}
