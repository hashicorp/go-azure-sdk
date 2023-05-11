package alertssuppressionrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AlertsSuppressionRuleId{}

// AlertsSuppressionRuleId is a struct representing the Resource ID for a Alerts Suppression Rule
type AlertsSuppressionRuleId struct {
	SubscriptionId            string
	AlertsSuppressionRuleName string
}

// NewAlertsSuppressionRuleID returns a new AlertsSuppressionRuleId struct
func NewAlertsSuppressionRuleID(subscriptionId string, alertsSuppressionRuleName string) AlertsSuppressionRuleId {
	return AlertsSuppressionRuleId{
		SubscriptionId:            subscriptionId,
		AlertsSuppressionRuleName: alertsSuppressionRuleName,
	}
}

// ParseAlertsSuppressionRuleID parses 'input' into a AlertsSuppressionRuleId
func ParseAlertsSuppressionRuleID(input string) (*AlertsSuppressionRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(AlertsSuppressionRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AlertsSuppressionRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.AlertsSuppressionRuleName, ok = parsed.Parsed["alertsSuppressionRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "alertsSuppressionRuleName", *parsed)
	}

	return &id, nil
}

// ParseAlertsSuppressionRuleIDInsensitively parses 'input' case-insensitively into a AlertsSuppressionRuleId
// note: this method should only be used for API response data and not user input
func ParseAlertsSuppressionRuleIDInsensitively(input string) (*AlertsSuppressionRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(AlertsSuppressionRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AlertsSuppressionRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.AlertsSuppressionRuleName, ok = parsed.Parsed["alertsSuppressionRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "alertsSuppressionRuleName", *parsed)
	}

	return &id, nil
}

// ValidateAlertsSuppressionRuleID checks that 'input' can be parsed as a Alerts Suppression Rule ID
func ValidateAlertsSuppressionRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAlertsSuppressionRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Alerts Suppression Rule ID
func (id AlertsSuppressionRuleId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Security/alertsSuppressionRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.AlertsSuppressionRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Alerts Suppression Rule ID
func (id AlertsSuppressionRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticAlertsSuppressionRules", "alertsSuppressionRules", "alertsSuppressionRules"),
		resourceids.UserSpecifiedSegment("alertsSuppressionRuleName", "alertsSuppressionRuleValue"),
	}
}

// String returns a human-readable description of this Alerts Suppression Rule ID
func (id AlertsSuppressionRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Alerts Suppression Rule Name: %q", id.AlertsSuppressionRuleName),
	}
	return fmt.Sprintf("Alerts Suppression Rule (%s)", strings.Join(components, "\n"))
}
