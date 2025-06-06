package dnsresolverpolicyvirtualnetworklinks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DnsResolverPolicyVirtualNetworkLinkId{})
}

var _ resourceids.ResourceId = &DnsResolverPolicyVirtualNetworkLinkId{}

// DnsResolverPolicyVirtualNetworkLinkId is a struct representing the Resource ID for a Dns Resolver Policy Virtual Network Link
type DnsResolverPolicyVirtualNetworkLinkId struct {
	SubscriptionId         string
	ResourceGroupName      string
	DnsResolverPolicyName  string
	VirtualNetworkLinkName string
}

// NewDnsResolverPolicyVirtualNetworkLinkID returns a new DnsResolverPolicyVirtualNetworkLinkId struct
func NewDnsResolverPolicyVirtualNetworkLinkID(subscriptionId string, resourceGroupName string, dnsResolverPolicyName string, virtualNetworkLinkName string) DnsResolverPolicyVirtualNetworkLinkId {
	return DnsResolverPolicyVirtualNetworkLinkId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		DnsResolverPolicyName:  dnsResolverPolicyName,
		VirtualNetworkLinkName: virtualNetworkLinkName,
	}
}

// ParseDnsResolverPolicyVirtualNetworkLinkID parses 'input' into a DnsResolverPolicyVirtualNetworkLinkId
func ParseDnsResolverPolicyVirtualNetworkLinkID(input string) (*DnsResolverPolicyVirtualNetworkLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DnsResolverPolicyVirtualNetworkLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DnsResolverPolicyVirtualNetworkLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDnsResolverPolicyVirtualNetworkLinkIDInsensitively parses 'input' case-insensitively into a DnsResolverPolicyVirtualNetworkLinkId
// note: this method should only be used for API response data and not user input
func ParseDnsResolverPolicyVirtualNetworkLinkIDInsensitively(input string) (*DnsResolverPolicyVirtualNetworkLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DnsResolverPolicyVirtualNetworkLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DnsResolverPolicyVirtualNetworkLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DnsResolverPolicyVirtualNetworkLinkId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.VirtualNetworkLinkName, ok = input.Parsed["virtualNetworkLinkName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualNetworkLinkName", input)
	}

	return nil
}

// ValidateDnsResolverPolicyVirtualNetworkLinkID checks that 'input' can be parsed as a Dns Resolver Policy Virtual Network Link ID
func ValidateDnsResolverPolicyVirtualNetworkLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDnsResolverPolicyVirtualNetworkLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dns Resolver Policy Virtual Network Link ID
func (id DnsResolverPolicyVirtualNetworkLinkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/dnsResolverPolicies/%s/virtualNetworkLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DnsResolverPolicyName, id.VirtualNetworkLinkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dns Resolver Policy Virtual Network Link ID
func (id DnsResolverPolicyVirtualNetworkLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticDnsResolverPolicies", "dnsResolverPolicies", "dnsResolverPolicies"),
		resourceids.UserSpecifiedSegment("dnsResolverPolicyName", "dnsResolverPolicyName"),
		resourceids.StaticSegment("staticVirtualNetworkLinks", "virtualNetworkLinks", "virtualNetworkLinks"),
		resourceids.UserSpecifiedSegment("virtualNetworkLinkName", "virtualNetworkLinkName"),
	}
}

// String returns a human-readable description of this Dns Resolver Policy Virtual Network Link ID
func (id DnsResolverPolicyVirtualNetworkLinkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Dns Resolver Policy Name: %q", id.DnsResolverPolicyName),
		fmt.Sprintf("Virtual Network Link Name: %q", id.VirtualNetworkLinkName),
	}
	return fmt.Sprintf("Dns Resolver Policy Virtual Network Link (%s)", strings.Join(components, "\n"))
}
