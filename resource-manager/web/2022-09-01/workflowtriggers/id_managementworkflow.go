package workflowtriggers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ManagementWorkflowId{})
}

var _ resourceids.ResourceId = &ManagementWorkflowId{}

// ManagementWorkflowId is a struct representing the Resource ID for a Management Workflow
type ManagementWorkflowId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	WorkflowName      string
}

// NewManagementWorkflowID returns a new ManagementWorkflowId struct
func NewManagementWorkflowID(subscriptionId string, resourceGroupName string, siteName string, workflowName string) ManagementWorkflowId {
	return ManagementWorkflowId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		WorkflowName:      workflowName,
	}
}

// ParseManagementWorkflowID parses 'input' into a ManagementWorkflowId
func ParseManagementWorkflowID(input string) (*ManagementWorkflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagementWorkflowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagementWorkflowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagementWorkflowIDInsensitively parses 'input' case-insensitively into a ManagementWorkflowId
// note: this method should only be used for API response data and not user input
func ParseManagementWorkflowIDInsensitively(input string) (*ManagementWorkflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagementWorkflowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagementWorkflowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagementWorkflowId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateManagementWorkflowID checks that 'input' can be parsed as a Management Workflow ID
func ValidateManagementWorkflowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagementWorkflowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Management Workflow ID
func (id ManagementWorkflowId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/hostRuntime/runtime/webHooks/workflow/api/management/workflows/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.WorkflowName)
}

// Segments returns a slice of Resource ID Segments which comprise this Management Workflow ID
func (id ManagementWorkflowId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Management Workflow ID
func (id ManagementWorkflowId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Workflow Name: %q", id.WorkflowName),
	}
	return fmt.Sprintf("Management Workflow (%s)", strings.Join(components, "\n"))
}
