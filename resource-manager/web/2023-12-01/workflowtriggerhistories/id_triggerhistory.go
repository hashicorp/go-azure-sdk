package workflowtriggerhistories

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TriggerHistoryId{})
}

var _ resourceids.ResourceId = &TriggerHistoryId{}

// TriggerHistoryId is a struct representing the Resource ID for a Trigger History
type TriggerHistoryId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	WorkflowName      string
	TriggerName       string
	HistoryName       string
}

// NewTriggerHistoryID returns a new TriggerHistoryId struct
func NewTriggerHistoryID(subscriptionId string, resourceGroupName string, siteName string, workflowName string, triggerName string, historyName string) TriggerHistoryId {
	return TriggerHistoryId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		WorkflowName:      workflowName,
		TriggerName:       triggerName,
		HistoryName:       historyName,
	}
}

// ParseTriggerHistoryID parses 'input' into a TriggerHistoryId
func ParseTriggerHistoryID(input string) (*TriggerHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TriggerHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggerHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTriggerHistoryIDInsensitively parses 'input' case-insensitively into a TriggerHistoryId
// note: this method should only be used for API response data and not user input
func ParseTriggerHistoryIDInsensitively(input string) (*TriggerHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TriggerHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggerHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TriggerHistoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SiteName, ok = input.Parsed["siteName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteName", input)
	}

	if id.WorkflowName, ok = input.Parsed["workflowName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workflowName", input)
	}

	if id.TriggerName, ok = input.Parsed["triggerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "triggerName", input)
	}

	if id.HistoryName, ok = input.Parsed["historyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "historyName", input)
	}

	return nil
}

// ValidateTriggerHistoryID checks that 'input' can be parsed as a Trigger History ID
func ValidateTriggerHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTriggerHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Trigger History ID
func (id TriggerHistoryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/hostRuntime/runtime/webHooks/workflow/api/management/workflows/%s/triggers/%s/histories/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.WorkflowName, id.TriggerName, id.HistoryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Trigger History ID
func (id TriggerHistoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "name"),
		resourceids.StaticSegment("staticHostRuntime", "hostRuntime", "hostRuntime"),
		resourceids.StaticSegment("staticRuntime", "runtime", "runtime"),
		resourceids.StaticSegment("staticWebHooks", "webHooks", "webHooks"),
		resourceids.StaticSegment("staticWorkflow", "workflow", "workflow"),
		resourceids.StaticSegment("staticApi", "api", "api"),
		resourceids.StaticSegment("staticManagement", "management", "management"),
		resourceids.StaticSegment("staticWorkflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowName", "workflowName"),
		resourceids.StaticSegment("staticTriggers", "triggers", "triggers"),
		resourceids.UserSpecifiedSegment("triggerName", "triggerName"),
		resourceids.StaticSegment("staticHistories", "histories", "histories"),
		resourceids.UserSpecifiedSegment("historyName", "historyName"),
	}
}

// String returns a human-readable description of this Trigger History ID
func (id TriggerHistoryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Workflow Name: %q", id.WorkflowName),
		fmt.Sprintf("Trigger Name: %q", id.TriggerName),
		fmt.Sprintf("History Name: %q", id.HistoryName),
	}
	return fmt.Sprintf("Trigger History (%s)", strings.Join(components, "\n"))
}
