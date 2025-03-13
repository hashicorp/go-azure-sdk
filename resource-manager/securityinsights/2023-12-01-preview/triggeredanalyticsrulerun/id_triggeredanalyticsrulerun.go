package triggeredanalyticsrulerun

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TriggeredAnalyticsRuleRunId{})
}

var _ resourceids.ResourceId = &TriggeredAnalyticsRuleRunId{}

// TriggeredAnalyticsRuleRunId is a struct representing the Resource ID for a Triggered Analytics Rule Run
type TriggeredAnalyticsRuleRunId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	RuleRunId         string
}

// NewTriggeredAnalyticsRuleRunID returns a new TriggeredAnalyticsRuleRunId struct
func NewTriggeredAnalyticsRuleRunID(subscriptionId string, resourceGroupName string, workspaceName string, ruleRunId string) TriggeredAnalyticsRuleRunId {
	return TriggeredAnalyticsRuleRunId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		RuleRunId:         ruleRunId,
	}
}

// ParseTriggeredAnalyticsRuleRunID parses 'input' into a TriggeredAnalyticsRuleRunId
func ParseTriggeredAnalyticsRuleRunID(input string) (*TriggeredAnalyticsRuleRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TriggeredAnalyticsRuleRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggeredAnalyticsRuleRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTriggeredAnalyticsRuleRunIDInsensitively parses 'input' case-insensitively into a TriggeredAnalyticsRuleRunId
// note: this method should only be used for API response data and not user input
func ParseTriggeredAnalyticsRuleRunIDInsensitively(input string) (*TriggeredAnalyticsRuleRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TriggeredAnalyticsRuleRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggeredAnalyticsRuleRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TriggeredAnalyticsRuleRunId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RuleRunId, ok = input.Parsed["ruleRunId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "ruleRunId", input)
	}

	return nil
}

// ValidateTriggeredAnalyticsRuleRunID checks that 'input' can be parsed as a Triggered Analytics Rule Run ID
func ValidateTriggeredAnalyticsRuleRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTriggeredAnalyticsRuleRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Triggered Analytics Rule Run ID
func (id TriggeredAnalyticsRuleRunId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/triggeredAnalyticsRuleRuns/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.RuleRunId)
}

// Segments returns a slice of Resource ID Segments which comprise this Triggered Analytics Rule Run ID
func (id TriggeredAnalyticsRuleRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticTriggeredAnalyticsRuleRuns", "triggeredAnalyticsRuleRuns", "triggeredAnalyticsRuleRuns"),
		resourceids.UserSpecifiedSegment("ruleRunId", "ruleRunId"),
	}
}

// String returns a human-readable description of this Triggered Analytics Rule Run ID
func (id TriggeredAnalyticsRuleRunId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Rule Run: %q", id.RuleRunId),
	}
	return fmt.Sprintf("Triggered Analytics Rule Run (%s)", strings.Join(components, "\n"))
}
