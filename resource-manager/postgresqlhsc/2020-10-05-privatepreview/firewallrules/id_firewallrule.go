// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firewallrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = FirewallRuleId{}

// FirewallRuleId is a struct representing the Resource ID for a Firewall Rule
type FirewallRuleId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ServerGroupsv2Name string
	FirewallRuleName   string
}

// NewFirewallRuleID returns a new FirewallRuleId struct
func NewFirewallRuleID(subscriptionId string, resourceGroupName string, serverGroupsv2Name string, firewallRuleName string) FirewallRuleId {
	return FirewallRuleId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ServerGroupsv2Name: serverGroupsv2Name,
		FirewallRuleName:   firewallRuleName,
	}
}

// ParseFirewallRuleID parses 'input' into a FirewallRuleId
func ParseFirewallRuleID(input string) (*FirewallRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(FirewallRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FirewallRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServerGroupsv2Name, ok = parsed.Parsed["serverGroupsv2Name"]; !ok {
		return nil, fmt.Errorf("the segment 'serverGroupsv2Name' was not found in the resource id %q", input)
	}

	if id.FirewallRuleName, ok = parsed.Parsed["firewallRuleName"]; !ok {
		return nil, fmt.Errorf("the segment 'firewallRuleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseFirewallRuleIDInsensitively parses 'input' case-insensitively into a FirewallRuleId
// note: this method should only be used for API response data and not user input
func ParseFirewallRuleIDInsensitively(input string) (*FirewallRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(FirewallRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FirewallRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ServerGroupsv2Name, ok = parsed.Parsed["serverGroupsv2Name"]; !ok {
		return nil, fmt.Errorf("the segment 'serverGroupsv2Name' was not found in the resource id %q", input)
	}

	if id.FirewallRuleName, ok = parsed.Parsed["firewallRuleName"]; !ok {
		return nil, fmt.Errorf("the segment 'firewallRuleName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateFirewallRuleID checks that 'input' can be parsed as a Firewall Rule ID
func ValidateFirewallRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFirewallRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Firewall Rule ID
func (id FirewallRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/%s/firewallRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerGroupsv2Name, id.FirewallRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Firewall Rule ID
func (id FirewallRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticServerGroupsv2", "serverGroupsv2", "serverGroupsv2"),
		resourceids.UserSpecifiedSegment("serverGroupsv2Name", "serverGroupsv2Value"),
		resourceids.StaticSegment("staticFirewallRules", "firewallRules", "firewallRules"),
		resourceids.UserSpecifiedSegment("firewallRuleName", "firewallRuleValue"),
	}
}

// String returns a human-readable description of this Firewall Rule ID
func (id FirewallRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Groupsv 2 Name: %q", id.ServerGroupsv2Name),
		fmt.Sprintf("Firewall Rule Name: %q", id.FirewallRuleName),
	}
	return fmt.Sprintf("Firewall Rule (%s)", strings.Join(components, "\n"))
}
