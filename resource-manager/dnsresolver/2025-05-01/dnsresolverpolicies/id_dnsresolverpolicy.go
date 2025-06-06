package dnsresolverpolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DnsResolverPolicyId{})
}

var _ resourceids.ResourceId = &DnsResolverPolicyId{}

// DnsResolverPolicyId is a struct representing the Resource ID for a Dns Resolver Policy
type DnsResolverPolicyId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DnsResolverPolicyName string
}

// NewDnsResolverPolicyID returns a new DnsResolverPolicyId struct
func NewDnsResolverPolicyID(subscriptionId string, resourceGroupName string, dnsResolverPolicyName string) DnsResolverPolicyId {
	return DnsResolverPolicyId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DnsResolverPolicyName: dnsResolverPolicyName,
	}
}

// ParseDnsResolverPolicyID parses 'input' into a DnsResolverPolicyId
func ParseDnsResolverPolicyID(input string) (*DnsResolverPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DnsResolverPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DnsResolverPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDnsResolverPolicyIDInsensitively parses 'input' case-insensitively into a DnsResolverPolicyId
// note: this method should only be used for API response data and not user input
func ParseDnsResolverPolicyIDInsensitively(input string) (*DnsResolverPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DnsResolverPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DnsResolverPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DnsResolverPolicyId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateDnsResolverPolicyID checks that 'input' can be parsed as a Dns Resolver Policy ID
func ValidateDnsResolverPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDnsResolverPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dns Resolver Policy ID
func (id DnsResolverPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/dnsResolverPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DnsResolverPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dns Resolver Policy ID
func (id DnsResolverPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticDnsResolverPolicies", "dnsResolverPolicies", "dnsResolverPolicies"),
		resourceids.UserSpecifiedSegment("dnsResolverPolicyName", "dnsResolverPolicyName"),
	}
}

// String returns a human-readable description of this Dns Resolver Policy ID
func (id DnsResolverPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Dns Resolver Policy Name: %q", id.DnsResolverPolicyName),
	}
	return fmt.Sprintf("Dns Resolver Policy (%s)", strings.Join(components, "\n"))
}
