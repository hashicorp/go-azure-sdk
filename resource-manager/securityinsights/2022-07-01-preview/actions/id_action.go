// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ActionId{}

// ActionId is a struct representing the Resource ID for a Action
type ActionId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	RuleId            string
	ActionId          string
}

// NewActionID returns a new ActionId struct
func NewActionID(subscriptionId string, resourceGroupName string, workspaceName string, ruleId string, actionId string) ActionId {
	return ActionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		RuleId:            ruleId,
		ActionId:          actionId,
	}
}

// ParseActionID parses 'input' into a ActionId
func ParseActionID(input string) (*ActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ActionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.RuleId, ok = parsed.Parsed["ruleId"]; !ok {
		return nil, fmt.Errorf("the segment 'ruleId' was not found in the resource id %q", input)
	}

	if id.ActionId, ok = parsed.Parsed["actionId"]; !ok {
		return nil, fmt.Errorf("the segment 'actionId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseActionIDInsensitively parses 'input' case-insensitively into a ActionId
// note: this method should only be used for API response data and not user input
func ParseActionIDInsensitively(input string) (*ActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ActionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.RuleId, ok = parsed.Parsed["ruleId"]; !ok {
		return nil, fmt.Errorf("the segment 'ruleId' was not found in the resource id %q", input)
	}

	if id.ActionId, ok = parsed.Parsed["actionId"]; !ok {
		return nil, fmt.Errorf("the segment 'actionId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateActionID checks that 'input' can be parsed as a Action ID
func ValidateActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Action ID
func (id ActionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/alertRules/%s/actions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.RuleId, id.ActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Action ID
func (id ActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticAlertRules", "alertRules", "alertRules"),
		resourceids.UserSpecifiedSegment("ruleId", "ruleIdValue"),
		resourceids.StaticSegment("staticActions", "actions", "actions"),
		resourceids.UserSpecifiedSegment("actionId", "actionIdValue"),
	}
}

// String returns a human-readable description of this Action ID
func (id ActionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Rule: %q", id.RuleId),
		fmt.Sprintf("Action: %q", id.ActionId),
	}
	return fmt.Sprintf("Action (%s)", strings.Join(components, "\n"))
}
