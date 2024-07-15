package javacomponents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&JavaComponentId{})
}

var _ resourceids.ResourceId = &JavaComponentId{}

// JavaComponentId is a struct representing the Resource ID for a Java Component
type JavaComponentId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ManagedEnvironmentName string
	JavaComponentName      string
}

// NewJavaComponentID returns a new JavaComponentId struct
func NewJavaComponentID(subscriptionId string, resourceGroupName string, managedEnvironmentName string, javaComponentName string) JavaComponentId {
	return JavaComponentId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ManagedEnvironmentName: managedEnvironmentName,
		JavaComponentName:      javaComponentName,
	}
}

// ParseJavaComponentID parses 'input' into a JavaComponentId
func ParseJavaComponentID(input string) (*JavaComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&JavaComponentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := JavaComponentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseJavaComponentIDInsensitively parses 'input' case-insensitively into a JavaComponentId
// note: this method should only be used for API response data and not user input
func ParseJavaComponentIDInsensitively(input string) (*JavaComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&JavaComponentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := JavaComponentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *JavaComponentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.JavaComponentName, ok = input.Parsed["javaComponentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "javaComponentName", input)
	}

	return nil
}

// ValidateJavaComponentID checks that 'input' can be parsed as a Java Component ID
func ValidateJavaComponentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseJavaComponentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Java Component ID
func (id JavaComponentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/managedEnvironments/%s/javaComponents/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedEnvironmentName, id.JavaComponentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Java Component ID
func (id JavaComponentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticManagedEnvironments", "managedEnvironments", "managedEnvironments"),
		resourceids.UserSpecifiedSegment("managedEnvironmentName", "managedEnvironmentValue"),
		resourceids.StaticSegment("staticJavaComponents", "javaComponents", "javaComponents"),
		resourceids.UserSpecifiedSegment("javaComponentName", "javaComponentValue"),
	}
}

// String returns a human-readable description of this Java Component ID
func (id JavaComponentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Environment Name: %q", id.ManagedEnvironmentName),
		fmt.Sprintf("Java Component Name: %q", id.JavaComponentName),
	}
	return fmt.Sprintf("Java Component (%s)", strings.Join(components, "\n"))
}
