package dotnetcomponents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DotNetComponentId{})
}

var _ resourceids.ResourceId = &DotNetComponentId{}

// DotNetComponentId is a struct representing the Resource ID for a Dot Net Component
type DotNetComponentId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ManagedEnvironmentName string
	DotNetComponentName    string
}

// NewDotNetComponentID returns a new DotNetComponentId struct
func NewDotNetComponentID(subscriptionId string, resourceGroupName string, managedEnvironmentName string, dotNetComponentName string) DotNetComponentId {
	return DotNetComponentId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ManagedEnvironmentName: managedEnvironmentName,
		DotNetComponentName:    dotNetComponentName,
	}
}

// ParseDotNetComponentID parses 'input' into a DotNetComponentId
func ParseDotNetComponentID(input string) (*DotNetComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DotNetComponentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DotNetComponentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDotNetComponentIDInsensitively parses 'input' case-insensitively into a DotNetComponentId
// note: this method should only be used for API response data and not user input
func ParseDotNetComponentIDInsensitively(input string) (*DotNetComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DotNetComponentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DotNetComponentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DotNetComponentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedEnvironmentName, ok = input.Parsed["managedEnvironmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedEnvironmentName", input)
	}

	if id.DotNetComponentName, ok = input.Parsed["dotNetComponentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dotNetComponentName", input)
	}

	return nil
}

// ValidateDotNetComponentID checks that 'input' can be parsed as a Dot Net Component ID
func ValidateDotNetComponentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDotNetComponentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dot Net Component ID
func (id DotNetComponentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/managedEnvironments/%s/dotNetComponents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedEnvironmentName, id.DotNetComponentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dot Net Component ID
func (id DotNetComponentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticManagedEnvironments", "managedEnvironments", "managedEnvironments"),
		resourceids.UserSpecifiedSegment("managedEnvironmentName", "managedEnvironmentValue"),
		resourceids.StaticSegment("staticDotNetComponents", "dotNetComponents", "dotNetComponents"),
		resourceids.UserSpecifiedSegment("dotNetComponentName", "dotNetComponentValue"),
	}
}

// String returns a human-readable description of this Dot Net Component ID
func (id DotNetComponentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Environment Name: %q", id.ManagedEnvironmentName),
		fmt.Sprintf("Dot Net Component Name: %q", id.DotNetComponentName),
	}
	return fmt.Sprintf("Dot Net Component (%s)", strings.Join(components, "\n"))
}
