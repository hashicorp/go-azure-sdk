package alertrulesapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AlertRuleId{})
}

var _ resourceids.ResourceId = &AlertRuleId{}

// AlertRuleId is a struct representing the Resource ID for a Alert Rule
type AlertRuleId struct {
	SubscriptionId    string
	ResourceGroupName string
	AlertRuleName     string
}

// NewAlertRuleID returns a new AlertRuleId struct
func NewAlertRuleID(subscriptionId string, resourceGroupName string, alertRuleName string) AlertRuleId {
	return AlertRuleId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AlertRuleName:     alertRuleName,
	}
}

// ParseAlertRuleID parses 'input' into a AlertRuleId
func ParseAlertRuleID(input string) (*AlertRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AlertRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AlertRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAlertRuleIDInsensitively parses 'input' case-insensitively into a AlertRuleId
// note: this method should only be used for API response data and not user input
func ParseAlertRuleIDInsensitively(input string) (*AlertRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AlertRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AlertRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AlertRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AlertRuleName, ok = input.Parsed["alertRuleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "alertRuleName", input)
	}

	return nil
}

// ValidateAlertRuleID checks that 'input' can be parsed as a Alert Rule ID
func ValidateAlertRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAlertRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Alert Rule ID
func (id AlertRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/alertRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AlertRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Alert Rule ID
func (id AlertRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticAlertRules", "alertRules", "alertRules"),
		resourceids.UserSpecifiedSegment("alertRuleName", "alertRuleName"),
	}
}

// String returns a human-readable description of this Alert Rule ID
func (id AlertRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Alert Rule Name: %q", id.AlertRuleName),
	}
	return fmt.Sprintf("Alert Rule (%s)", strings.Join(components, "\n"))
}
