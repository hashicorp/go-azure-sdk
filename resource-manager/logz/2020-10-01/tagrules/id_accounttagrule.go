// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tagrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = AccountTagRuleId{}

// AccountTagRuleId is a struct representing the Resource ID for a Account Tag Rule
type AccountTagRuleId struct {
	SubscriptionId    string
	ResourceGroupName string
	MonitorName       string
	AccountName       string
	TagRuleName       string
}

// NewAccountTagRuleID returns a new AccountTagRuleId struct
func NewAccountTagRuleID(subscriptionId string, resourceGroupName string, monitorName string, accountName string, tagRuleName string) AccountTagRuleId {
	return AccountTagRuleId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		MonitorName:       monitorName,
		AccountName:       accountName,
		TagRuleName:       tagRuleName,
	}
}

// ParseAccountTagRuleID parses 'input' into a AccountTagRuleId
func ParseAccountTagRuleID(input string) (*AccountTagRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountTagRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccountTagRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.MonitorName, ok = parsed.Parsed["monitorName"]; !ok {
		return nil, fmt.Errorf("the segment 'monitorName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.TagRuleName, ok = parsed.Parsed["tagRuleName"]; !ok {
		return nil, fmt.Errorf("the segment 'tagRuleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseAccountTagRuleIDInsensitively parses 'input' case-insensitively into a AccountTagRuleId
// note: this method should only be used for API response data and not user input
func ParseAccountTagRuleIDInsensitively(input string) (*AccountTagRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountTagRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccountTagRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.MonitorName, ok = parsed.Parsed["monitorName"]; !ok {
		return nil, fmt.Errorf("the segment 'monitorName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.TagRuleName, ok = parsed.Parsed["tagRuleName"]; !ok {
		return nil, fmt.Errorf("the segment 'tagRuleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateAccountTagRuleID checks that 'input' can be parsed as a Account Tag Rule ID
func ValidateAccountTagRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAccountTagRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Account Tag Rule ID
func (id AccountTagRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Logz/monitors/%s/accounts/%s/tagRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MonitorName, id.AccountName, id.TagRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Account Tag Rule ID
func (id AccountTagRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftLogz", "Microsoft.Logz", "Microsoft.Logz"),
		resourceids.StaticSegment("staticMonitors", "monitors", "monitors"),
		resourceids.UserSpecifiedSegment("monitorName", "monitorValue"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticTagRules", "tagRules", "tagRules"),
		resourceids.UserSpecifiedSegment("tagRuleName", "tagRuleValue"),
	}
}

// String returns a human-readable description of this Account Tag Rule ID
func (id AccountTagRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Monitor Name: %q", id.MonitorName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Tag Rule Name: %q", id.TagRuleName),
	}
	return fmt.Sprintf("Account Tag Rule (%s)", strings.Join(components, "\n"))
}
