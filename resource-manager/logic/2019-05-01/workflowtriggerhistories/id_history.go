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
	recaser.RegisterResourceId(&HistoryId{})
}

var _ resourceids.ResourceId = &HistoryId{}

// HistoryId is a struct representing the Resource ID for a History
type HistoryId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkflowName      string
	TriggerName       string
	HistoryName       string
}

// NewHistoryID returns a new HistoryId struct
func NewHistoryID(subscriptionId string, resourceGroupName string, workflowName string, triggerName string, historyName string) HistoryId {
	return HistoryId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkflowName:      workflowName,
		TriggerName:       triggerName,
		HistoryName:       historyName,
	}
}

// ParseHistoryID parses 'input' into a HistoryId
func ParseHistoryID(input string) (*HistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHistoryIDInsensitively parses 'input' case-insensitively into a HistoryId
// note: this method should only be used for API response data and not user input
func ParseHistoryIDInsensitively(input string) (*HistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HistoryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
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

// ValidateHistoryID checks that 'input' can be parsed as a History ID
func ValidateHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted History ID
func (id HistoryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Logic/workflows/%s/triggers/%s/histories/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkflowName, id.TriggerName, id.HistoryName)
}

// Segments returns a slice of Resource ID Segments which comprise this History ID
func (id HistoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftLogic", "Microsoft.Logic", "Microsoft.Logic"),
		resourceids.StaticSegment("staticWorkflows", "workflows", "workflows"),
		resourceids.UserSpecifiedSegment("workflowName", "workflowName"),
		resourceids.StaticSegment("staticTriggers", "triggers", "triggers"),
		resourceids.UserSpecifiedSegment("triggerName", "triggerName"),
		resourceids.StaticSegment("staticHistories", "histories", "histories"),
		resourceids.UserSpecifiedSegment("historyName", "historyName"),
	}
}

// String returns a human-readable description of this History ID
func (id HistoryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workflow Name: %q", id.WorkflowName),
		fmt.Sprintf("Trigger Name: %q", id.TriggerName),
		fmt.Sprintf("History Name: %q", id.HistoryName),
	}
	return fmt.Sprintf("History (%s)", strings.Join(components, "\n"))
}
