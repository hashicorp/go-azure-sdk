package analyticsitemsapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProviderComponentId{}

// ProviderComponentId is a struct representing the Resource ID for a Provider Component
type ProviderComponentId struct {
	SubscriptionId    string
	ResourceGroupName string
	ComponentName     string
	ScopePath         string
}

// NewProviderComponentID returns a new ProviderComponentId struct
func NewProviderComponentID(subscriptionId string, resourceGroupName string, componentName string, scopePath string) ProviderComponentId {
	return ProviderComponentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ComponentName:     componentName,
		ScopePath:         scopePath,
	}
}

// ParseProviderComponentID parses 'input' into a ProviderComponentId
func ParseProviderComponentID(input string) (*ProviderComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderComponentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderComponentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ComponentName, ok = parsed.Parsed["componentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "componentName", *parsed)
	}

	if id.ScopePath, ok = parsed.Parsed["scopePath"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopePath", *parsed)
	}

	return &id, nil
}

// ParseProviderComponentIDInsensitively parses 'input' case-insensitively into a ProviderComponentId
// note: this method should only be used for API response data and not user input
func ParseProviderComponentIDInsensitively(input string) (*ProviderComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderComponentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderComponentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ComponentName, ok = parsed.Parsed["componentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "componentName", *parsed)
	}

	if id.ScopePath, ok = parsed.Parsed["scopePath"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopePath", *parsed)
	}

	return &id, nil
}

// ValidateProviderComponentID checks that 'input' can be parsed as a Provider Component ID
func ValidateProviderComponentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderComponentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Component ID
func (id ProviderComponentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/components/%s/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ComponentName, strings.TrimPrefix(id.ScopePath, "/"))
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Component ID
func (id ProviderComponentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentValue"),
		resourceids.ScopeSegment("scopePath", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
	}
}

// String returns a human-readable description of this Provider Component ID
func (id ProviderComponentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Scope Path: %q", id.ScopePath),
	}
	return fmt.Sprintf("Provider Component (%s)", strings.Join(components, "\n"))
}
