package workloadnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DnsServiceId{}

// DnsServiceId is a struct representing the Resource ID for a Dns Service
type DnsServiceId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	DnsServiceId      string
}

// NewDnsServiceID returns a new DnsServiceId struct
func NewDnsServiceID(subscriptionId string, resourceGroupName string, privateCloudName string, dnsServiceId string) DnsServiceId {
	return DnsServiceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		DnsServiceId:      dnsServiceId,
	}
}

// ParseDnsServiceID parses 'input' into a DnsServiceId
func ParseDnsServiceID(input string) (*DnsServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DnsServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DnsServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.DnsServiceId, ok = parsed.Parsed["dnsServiceId"]; !ok {
		return nil, fmt.Errorf("the segment 'dnsServiceId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDnsServiceIDInsensitively parses 'input' case-insensitively into a DnsServiceId
// note: this method should only be used for API response data and not user input
func ParseDnsServiceIDInsensitively(input string) (*DnsServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DnsServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DnsServiceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.DnsServiceId, ok = parsed.Parsed["dnsServiceId"]; !ok {
		return nil, fmt.Errorf("the segment 'dnsServiceId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDnsServiceID checks that 'input' can be parsed as a Dns Service ID
func ValidateDnsServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDnsServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dns Service ID
func (id DnsServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/dnsServices/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.DnsServiceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Dns Service ID
func (id DnsServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticWorkloadNetworks", "workloadNetworks", "workloadNetworks"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticDnsServices", "dnsServices", "dnsServices"),
		resourceids.UserSpecifiedSegment("dnsServiceId", "dnsServiceIdValue"),
	}
}

// String returns a human-readable description of this Dns Service ID
func (id DnsServiceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Dns Service: %q", id.DnsServiceId),
	}
	return fmt.Sprintf("Dns Service (%s)", strings.Join(components, "\n"))
}
