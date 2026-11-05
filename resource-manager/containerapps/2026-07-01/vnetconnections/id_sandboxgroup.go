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
	recaser.RegisterResourceId(&SandboxGroupId{})
}

var _ resourceids.ResourceId = &SandboxGroupId{}

// SandboxGroupId is a struct representing the Resource ID for a Sandbox Group
type SandboxGroupId struct {
	SubscriptionId    string
	ResourceGroupName string
	SandboxGroupName  string
}

// NewSandboxGroupID returns a new SandboxGroupId struct
func NewSandboxGroupID(subscriptionId string, resourceGroupName string, sandboxGroupName string) SandboxGroupId {
	return SandboxGroupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SandboxGroupName:  sandboxGroupName,
	}
}

// ParseSandboxGroupID parses 'input' into a SandboxGroupId
func ParseSandboxGroupID(input string) (*SandboxGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SandboxGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SandboxGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSandboxGroupIDInsensitively parses 'input' case-insensitively into a SandboxGroupId
// note: this method should only be used for API response data and not user input
func ParseSandboxGroupIDInsensitively(input string) (*SandboxGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SandboxGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SandboxGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SandboxGroupId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateSandboxGroupID checks that 'input' can be parsed as a Sandbox Group ID
func ValidateSandboxGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSandboxGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sandbox Group ID
func (id SandboxGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/sandboxGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SandboxGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sandbox Group ID
func (id SandboxGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticSandboxGroups", "sandboxGroups", "sandboxGroups"),
		resourceids.UserSpecifiedSegment("sandboxGroupName", "sandboxGroupName"),
	}
}

// String returns a human-readable description of this Sandbox Group ID
func (id SandboxGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Sandbox Group Name: %q", id.SandboxGroupName),
	}
	return fmt.Sprintf("Sandbox Group (%s)", strings.Join(components, "\n"))
}
