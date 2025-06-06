package dnsresolverdomainlists

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DnsResolverDomainListId{})
}

var _ resourceids.ResourceId = &DnsResolverDomainListId{}

// DnsResolverDomainListId is a struct representing the Resource ID for a Dns Resolver Domain List
type DnsResolverDomainListId struct {
	SubscriptionId            string
	ResourceGroupName         string
	DnsResolverDomainListName string
}

// NewDnsResolverDomainListID returns a new DnsResolverDomainListId struct
func NewDnsResolverDomainListID(subscriptionId string, resourceGroupName string, dnsResolverDomainListName string) DnsResolverDomainListId {
	return DnsResolverDomainListId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		DnsResolverDomainListName: dnsResolverDomainListName,
	}
}

// ParseDnsResolverDomainListID parses 'input' into a DnsResolverDomainListId
func ParseDnsResolverDomainListID(input string) (*DnsResolverDomainListId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DnsResolverDomainListId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DnsResolverDomainListId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDnsResolverDomainListIDInsensitively parses 'input' case-insensitively into a DnsResolverDomainListId
// note: this method should only be used for API response data and not user input
func ParseDnsResolverDomainListIDInsensitively(input string) (*DnsResolverDomainListId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DnsResolverDomainListId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DnsResolverDomainListId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DnsResolverDomainListId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DnsResolverDomainListName, ok = input.Parsed["dnsResolverDomainListName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dnsResolverDomainListName", input)
	}

	return nil
}

// ValidateDnsResolverDomainListID checks that 'input' can be parsed as a Dns Resolver Domain List ID
func ValidateDnsResolverDomainListID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDnsResolverDomainListID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dns Resolver Domain List ID
func (id DnsResolverDomainListId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/dnsResolverDomainLists/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DnsResolverDomainListName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dns Resolver Domain List ID
func (id DnsResolverDomainListId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticDnsResolverDomainLists", "dnsResolverDomainLists", "dnsResolverDomainLists"),
		resourceids.UserSpecifiedSegment("dnsResolverDomainListName", "dnsResolverDomainListName"),
	}
}

// String returns a human-readable description of this Dns Resolver Domain List ID
func (id DnsResolverDomainListId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Dns Resolver Domain List Name: %q", id.DnsResolverDomainListName),
	}
	return fmt.Sprintf("Dns Resolver Domain List (%s)", strings.Join(components, "\n"))
}
