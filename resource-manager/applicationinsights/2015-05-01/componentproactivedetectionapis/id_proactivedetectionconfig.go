package componentproactivedetectionapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProactiveDetectionConfigId{}

// ProactiveDetectionConfigId is a struct representing the Resource ID for a Proactive Detection Config
type ProactiveDetectionConfigId struct {
	SubscriptionId    string
	ResourceGroupName string
	ComponentName     string
	ConfigurationId   string
}

// NewProactiveDetectionConfigID returns a new ProactiveDetectionConfigId struct
func NewProactiveDetectionConfigID(subscriptionId string, resourceGroupName string, componentName string, configurationId string) ProactiveDetectionConfigId {
	return ProactiveDetectionConfigId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ComponentName:     componentName,
		ConfigurationId:   configurationId,
	}
}

// ParseProactiveDetectionConfigID parses 'input' into a ProactiveDetectionConfigId
func ParseProactiveDetectionConfigID(input string) (*ProactiveDetectionConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProactiveDetectionConfigId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProactiveDetectionConfigId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ComponentName, ok = parsed.Parsed["componentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "componentName", *parsed)
	}

	if id.ConfigurationId, ok = parsed.Parsed["configurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationId", *parsed)
	}

	return &id, nil
}

// ParseProactiveDetectionConfigIDInsensitively parses 'input' case-insensitively into a ProactiveDetectionConfigId
// note: this method should only be used for API response data and not user input
func ParseProactiveDetectionConfigIDInsensitively(input string) (*ProactiveDetectionConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProactiveDetectionConfigId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProactiveDetectionConfigId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ComponentName, ok = parsed.Parsed["componentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "componentName", *parsed)
	}

	if id.ConfigurationId, ok = parsed.Parsed["configurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationId", *parsed)
	}

	return &id, nil
}

// ValidateProactiveDetectionConfigID checks that 'input' can be parsed as a Proactive Detection Config ID
func ValidateProactiveDetectionConfigID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProactiveDetectionConfigID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Proactive Detection Config ID
func (id ProactiveDetectionConfigId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/components/%s/proactiveDetectionConfigs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ComponentName, id.ConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Proactive Detection Config ID
func (id ProactiveDetectionConfigId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentValue"),
		resourceids.StaticSegment("staticProactiveDetectionConfigs", "proactiveDetectionConfigs", "proactiveDetectionConfigs"),
		resourceids.UserSpecifiedSegment("configurationId", "configurationIdValue"),
	}
}

// String returns a human-readable description of this Proactive Detection Config ID
func (id ProactiveDetectionConfigId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Configuration: %q", id.ConfigurationId),
	}
	return fmt.Sprintf("Proactive Detection Config (%s)", strings.Join(components, "\n"))
}
