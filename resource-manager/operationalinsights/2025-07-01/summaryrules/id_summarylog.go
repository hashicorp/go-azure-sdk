package summaryrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SummaryLogId{})
}

var _ resourceids.ResourceId = &SummaryLogId{}

// SummaryLogId is a struct representing the Resource ID for a Summary Log
type SummaryLogId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	SummaryLogName    string
}

// NewSummaryLogID returns a new SummaryLogId struct
func NewSummaryLogID(subscriptionId string, resourceGroupName string, workspaceName string, summaryLogName string) SummaryLogId {
	return SummaryLogId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		SummaryLogName:    summaryLogName,
	}
}

// ParseSummaryLogID parses 'input' into a SummaryLogId
func ParseSummaryLogID(input string) (*SummaryLogId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SummaryLogId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SummaryLogId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSummaryLogIDInsensitively parses 'input' case-insensitively into a SummaryLogId
// note: this method should only be used for API response data and not user input
func ParseSummaryLogIDInsensitively(input string) (*SummaryLogId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SummaryLogId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SummaryLogId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SummaryLogId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.SummaryLogName, ok = input.Parsed["summaryLogName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "summaryLogName", input)
	}

	return nil
}

// ValidateSummaryLogID checks that 'input' can be parsed as a Summary Log ID
func ValidateSummaryLogID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSummaryLogID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Summary Log ID
func (id SummaryLogId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/summaryLogs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SummaryLogName)
}

// Segments returns a slice of Resource ID Segments which comprise this Summary Log ID
func (id SummaryLogId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticSummaryLogs", "summaryLogs", "summaryLogs"),
		resourceids.UserSpecifiedSegment("summaryLogName", "summaryLogName"),
	}
}

// String returns a human-readable description of this Summary Log ID
func (id SummaryLogId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Summary Log Name: %q", id.SummaryLogName),
	}
	return fmt.Sprintf("Summary Log (%s)", strings.Join(components, "\n"))
}
