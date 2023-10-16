package componentworkitemconfigsapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkItemConfigId{}

// WorkItemConfigId is a struct representing the Resource ID for a Work Item Config
type WorkItemConfigId struct {
	SubscriptionId    string
	ResourceGroupName string
	ComponentName     string
	WorkItemConfigId  string
}

// NewWorkItemConfigID returns a new WorkItemConfigId struct
func NewWorkItemConfigID(subscriptionId string, resourceGroupName string, componentName string, workItemConfigId string) WorkItemConfigId {
	return WorkItemConfigId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ComponentName:     componentName,
		WorkItemConfigId:  workItemConfigId,
	}
}

// ParseWorkItemConfigID parses 'input' into a WorkItemConfigId
func ParseWorkItemConfigID(input string) (*WorkItemConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkItemConfigId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkItemConfigId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ComponentName, ok = parsed.Parsed["componentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "componentName", *parsed)
	}

	if id.WorkItemConfigId, ok = parsed.Parsed["workItemConfigId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workItemConfigId", *parsed)
	}

	return &id, nil
}

// ParseWorkItemConfigIDInsensitively parses 'input' case-insensitively into a WorkItemConfigId
// note: this method should only be used for API response data and not user input
func ParseWorkItemConfigIDInsensitively(input string) (*WorkItemConfigId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkItemConfigId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkItemConfigId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ComponentName, ok = parsed.Parsed["componentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "componentName", *parsed)
	}

	if id.WorkItemConfigId, ok = parsed.Parsed["workItemConfigId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workItemConfigId", *parsed)
	}

	return &id, nil
}

// ValidateWorkItemConfigID checks that 'input' can be parsed as a Work Item Config ID
func ValidateWorkItemConfigID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkItemConfigID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Work Item Config ID
func (id WorkItemConfigId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/components/%s/workItemConfigs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ComponentName, id.WorkItemConfigId)
}

// Segments returns a slice of Resource ID Segments which comprise this Work Item Config ID
func (id WorkItemConfigId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentValue"),
		resourceids.StaticSegment("staticWorkItemConfigs", "workItemConfigs", "workItemConfigs"),
		resourceids.UserSpecifiedSegment("workItemConfigId", "workItemConfigIdValue"),
	}
}

// String returns a human-readable description of this Work Item Config ID
func (id WorkItemConfigId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Work Item Config: %q", id.WorkItemConfigId),
	}
	return fmt.Sprintf("Work Item Config (%s)", strings.Join(components, "\n"))
}
