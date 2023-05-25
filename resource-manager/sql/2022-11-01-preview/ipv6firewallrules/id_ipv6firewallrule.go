package ipv6firewallrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IPv6FirewallRuleId{}

// IPv6FirewallRuleId is a struct representing the Resource ID for a I Pv 6 Firewall Rule
type IPv6FirewallRuleId struct {
	SubscriptionId       string
	ResourceGroupName    string
	ServerName           string
	Ipv6FirewallRuleName string
}

// NewIPv6FirewallRuleID returns a new IPv6FirewallRuleId struct
func NewIPv6FirewallRuleID(subscriptionId string, resourceGroupName string, serverName string, ipv6FirewallRuleName string) IPv6FirewallRuleId {
	return IPv6FirewallRuleId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		ServerName:           serverName,
		Ipv6FirewallRuleName: ipv6FirewallRuleName,
	}
}

// ParseIPv6FirewallRuleID parses 'input' into a IPv6FirewallRuleId
func ParseIPv6FirewallRuleID(input string) (*IPv6FirewallRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(IPv6FirewallRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IPv6FirewallRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.Ipv6FirewallRuleName, ok = parsed.Parsed["ipv6FirewallRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "ipv6FirewallRuleName", *parsed)
	}

	return &id, nil
}

// ParseIPv6FirewallRuleIDInsensitively parses 'input' case-insensitively into a IPv6FirewallRuleId
// note: this method should only be used for API response data and not user input
func ParseIPv6FirewallRuleIDInsensitively(input string) (*IPv6FirewallRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(IPv6FirewallRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IPv6FirewallRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.Ipv6FirewallRuleName, ok = parsed.Parsed["ipv6FirewallRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "ipv6FirewallRuleName", *parsed)
	}

	return &id, nil
}

// ValidateIPv6FirewallRuleID checks that 'input' can be parsed as a I Pv 6 Firewall Rule ID
func ValidateIPv6FirewallRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIPv6FirewallRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted I Pv 6 Firewall Rule ID
func (id IPv6FirewallRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/ipv6FirewallRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.Ipv6FirewallRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this I Pv 6 Firewall Rule ID
func (id IPv6FirewallRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticIpv6FirewallRules", "ipv6FirewallRules", "ipv6FirewallRules"),
		resourceids.UserSpecifiedSegment("ipv6FirewallRuleName", "ipv6FirewallRuleValue"),
	}
}

// String returns a human-readable description of this I Pv 6 Firewall Rule ID
func (id IPv6FirewallRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Ipv 6 Firewall Rule Name: %q", id.Ipv6FirewallRuleName),
	}
	return fmt.Sprintf("I Pv 6 Firewall Rule (%s)", strings.Join(components, "\n"))
}
