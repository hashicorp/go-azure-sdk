package componentcontainer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RegistryComponentId{})
}

var _ resourceids.ResourceId = &RegistryComponentId{}

// RegistryComponentId is a struct representing the Resource ID for a Registry Component
type RegistryComponentId struct {
	SubscriptionId    string
	ResourceGroupName string
	RegistryName      string
	ComponentName     string
}

// NewRegistryComponentID returns a new RegistryComponentId struct
func NewRegistryComponentID(subscriptionId string, resourceGroupName string, registryName string, componentName string) RegistryComponentId {
	return RegistryComponentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		RegistryName:      registryName,
		ComponentName:     componentName,
	}
}

// ParseRegistryComponentID parses 'input' into a RegistryComponentId
func ParseRegistryComponentID(input string) (*RegistryComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegistryComponentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegistryComponentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRegistryComponentIDInsensitively parses 'input' case-insensitively into a RegistryComponentId
// note: this method should only be used for API response data and not user input
func ParseRegistryComponentIDInsensitively(input string) (*RegistryComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegistryComponentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegistryComponentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RegistryComponentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.RegistryName, ok = input.Parsed["registryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "registryName", input)
	}

	if id.ComponentName, ok = input.Parsed["componentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "componentName", input)
	}

	return nil
}

// ValidateRegistryComponentID checks that 'input' can be parsed as a Registry Component ID
func ValidateRegistryComponentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegistryComponentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Registry Component ID
func (id RegistryComponentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/registries/%s/components/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RegistryName, id.ComponentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Registry Component ID
func (id RegistryComponentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticRegistries", "registries", "registries"),
		resourceids.UserSpecifiedSegment("registryName", "registryName"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
	}
}

// String returns a human-readable description of this Registry Component ID
func (id RegistryComponentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Registry Name: %q", id.RegistryName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
	}
	return fmt.Sprintf("Registry Component (%s)", strings.Join(components, "\n"))
}
