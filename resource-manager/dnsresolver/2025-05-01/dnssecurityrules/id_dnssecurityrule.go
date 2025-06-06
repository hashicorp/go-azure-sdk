package dnssecurityrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DnsSecurityRuleId{})
}

var _ resourceids.ResourceId = &DnsSecurityRuleId{}

// DnsSecurityRuleId is a struct representing the Resource ID for a Dns Security Rule
type DnsSecurityRuleId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DnsResolverPolicyName string
	DnsSecurityRuleName   string
}

// NewDnsSecurityRuleID returns a new DnsSecurityRuleId struct
func NewDnsSecurityRuleID(subscriptionId string, resourceGroupName string, dnsResolverPolicyName string, dnsSecurityRuleName string) DnsSecurityRuleId {
	return DnsSecurityRuleId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DnsResolverPolicyName: dnsResolverPolicyName,
		DnsSecurityRuleName:   dnsSecurityRuleName,
	}
}

// ParseDnsSecurityRuleID parses 'input' into a DnsSecurityRuleId
func ParseDnsSecurityRuleID(input string) (*DnsSecurityRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DnsSecurityRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DnsSecurityRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDnsSecurityRuleIDInsensitively parses 'input' case-insensitively into a DnsSecurityRuleId
// note: this method should only be used for API response data and not user input
func ParseDnsSecurityRuleIDInsensitively(input string) (*DnsSecurityRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DnsSecurityRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DnsSecurityRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DnsSecurityRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DnsResolverPolicyName, ok = input.Parsed["dnsResolverPolicyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dnsResolverPolicyName", input)
	}

	if id.DnsSecurityRuleName, ok = input.Parsed["dnsSecurityRuleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dnsSecurityRuleName", input)
	}

	return nil
}

// ValidateDnsSecurityRuleID checks that 'input' can be parsed as a Dns Security Rule ID
func ValidateDnsSecurityRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDnsSecurityRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dns Security Rule ID
func (id DnsSecurityRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/dnsResolverPolicies/%s/dnsSecurityRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DnsResolverPolicyName, id.DnsSecurityRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dns Security Rule ID
func (id DnsSecurityRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticDnsResolverPolicies", "dnsResolverPolicies", "dnsResolverPolicies"),
		resourceids.UserSpecifiedSegment("dnsResolverPolicyName", "dnsResolverPolicyName"),
		resourceids.StaticSegment("staticDnsSecurityRules", "dnsSecurityRules", "dnsSecurityRules"),
		resourceids.UserSpecifiedSegment("dnsSecurityRuleName", "dnsSecurityRuleName"),
	}
}

// String returns a human-readable description of this Dns Security Rule ID
func (id DnsSecurityRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Dns Resolver Policy Name: %q", id.DnsResolverPolicyName),
		fmt.Sprintf("Dns Security Rule Name: %q", id.DnsSecurityRuleName),
	}
	return fmt.Sprintf("Dns Security Rule (%s)", strings.Join(components, "\n"))
}
