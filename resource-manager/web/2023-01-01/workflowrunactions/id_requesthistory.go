package workflowrunactions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RequestHistoryId{})
}

var _ resourceids.ResourceId = &RequestHistoryId{}

// RequestHistoryId is a struct representing the Resource ID for a Request History
type RequestHistoryId struct {
	SubscriptionId     string
	ResourceGroupName  string
	SiteName           string
	WorkflowName       string
	RunName            string
	ActionName         string
	RepetitionName     string
	RequestHistoryName string
}

// NewRequestHistoryID returns a new RequestHistoryId struct
func NewRequestHistoryID(subscriptionId string, resourceGroupName string, siteName string, workflowName string, runName string, actionName string, repetitionName string, requestHistoryName string) RequestHistoryId {
	return RequestHistoryId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		SiteName:           siteName,
		WorkflowName:       workflowName,
		RunName:            runName,
		ActionName:         actionName,
		RepetitionName:     repetitionName,
		RequestHistoryName: requestHistoryName,
	}
}

// ParseRequestHistoryID parses 'input' into a RequestHistoryId
func ParseRequestHistoryID(input string) (*RequestHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RequestHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RequestHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRequestHistoryIDInsensitively parses 'input' case-insensitively into a RequestHistoryId
// note: this method should only be used for API response data and not user input
func ParseRequestHistoryIDInsensitively(input string) (*RequestHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RequestHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RequestHistoryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RequestHistoryId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RunName, ok = input.Parsed["runName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runName", input)
	}

	if id.ActionName, ok = input.Parsed["actionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "actionName", input)
	}

	if id.RepetitionName, ok = input.Parsed["repetitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "repetitionName", input)
	}

	if id.RequestHistoryName, ok = input.Parsed["requestHistoryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "requestHistoryName", input)
	}

	return nil
}

// ValidateRequestHistoryID checks that 'input' can be parsed as a Request History ID
func ValidateRequestHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRequestHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Request History ID
func (id RequestHistoryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/hostRuntime/runtime/webHooks/workflow/api/management/workflows/%s/runs/%s/actions/%s/repetitions/%s/requestHistories/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.WorkflowName, id.RunName, id.ActionName, id.RepetitionName, id.RequestHistoryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Request History ID
func (id RequestHistoryId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticRuns", "runs", "runs"),
		resourceids.UserSpecifiedSegment("runName", "runValue"),
		resourceids.StaticSegment("staticActions", "actions", "actions"),
		resourceids.UserSpecifiedSegment("actionName", "actionValue"),
		resourceids.StaticSegment("staticRepetitions", "repetitions", "repetitions"),
		resourceids.UserSpecifiedSegment("repetitionName", "repetitionValue"),
		resourceids.StaticSegment("staticRequestHistories", "requestHistories", "requestHistories"),
		resourceids.UserSpecifiedSegment("requestHistoryName", "requestHistoryValue"),
	}
}

// String returns a human-readable description of this Request History ID
func (id RequestHistoryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Workflow Name: %q", id.WorkflowName),
		fmt.Sprintf("Run Name: %q", id.RunName),
		fmt.Sprintf("Action Name: %q", id.ActionName),
		fmt.Sprintf("Repetition Name: %q", id.RepetitionName),
		fmt.Sprintf("Request History Name: %q", id.RequestHistoryName),
	}
	return fmt.Sprintf("Request History (%s)", strings.Join(components, "\n"))
}
