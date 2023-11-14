package workflowtriggers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TriggerId{}

// TriggerId is a struct representing the Resource ID for a Trigger
type TriggerId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	WorkflowName      string
	TriggerName       string
}

// NewTriggerID returns a new TriggerId struct
func NewTriggerID(subscriptionId string, resourceGroupName string, siteName string, workflowName string, triggerName string) TriggerId {
	return TriggerId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		WorkflowName:      workflowName,
		TriggerName:       triggerName,
	}
}

// ParseTriggerID parses 'input' into a TriggerId
func ParseTriggerID(input string) (*TriggerId, error) {
	parser := resourceids.NewParserFromResourceIdType(TriggerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TriggerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.WorkflowName, ok = parsed.Parsed["workflowName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workflowName", *parsed)
	}

	if id.TriggerName, ok = parsed.Parsed["triggerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "triggerName", *parsed)
	}

	return &id, nil
}

// ParseTriggerIDInsensitively parses 'input' case-insensitively into a TriggerId
// note: this method should only be used for API response data and not user input
func ParseTriggerIDInsensitively(input string) (*TriggerId, error) {
	parser := resourceids.NewParserFromResourceIdType(TriggerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TriggerId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.WorkflowName, ok = parsed.Parsed["workflowName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workflowName", *parsed)
	}

	if id.TriggerName, ok = parsed.Parsed["triggerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "triggerName", *parsed)
	}

	return &id, nil
}

// ValidateTriggerID checks that 'input' can be parsed as a Trigger ID
func ValidateTriggerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTriggerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Trigger ID
func (id TriggerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/hostRuntime/runtime/webHooks/workflow/api/management/workflows/%s/triggers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.WorkflowName, id.TriggerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Trigger ID
func (id TriggerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticHostRuntime", "hostRuntime", "hostRuntime"),
		resourceids.StaticSegment("staticRuntime", "runtime", "runtime"),
		resourceids.StaticSegment("staticWebHooks", "webHooks", "webHooks"),
		resourceids.StaticSegment("staticWorkflow", "workflow", "workflow"),
		resourceids.StaticSegment("staticApi", "api", "api"),
		resourceids.StaticSegment("staticManagement", "management", "management"),
		resourceids.StaticSegment("staticWorkflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowName", "workflowValue"),
		resourceids.StaticSegment("staticTriggers", "triggers", "triggers"),
		resourceids.UserSpecifiedSegment("triggerName", "triggerValue"),
	}
}

// String returns a human-readable description of this Trigger ID
func (id TriggerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Workflow Name: %q", id.WorkflowName),
		fmt.Sprintf("Trigger Name: %q", id.TriggerName),
	}
	return fmt.Sprintf("Trigger (%s)", strings.Join(components, "\n"))
}
