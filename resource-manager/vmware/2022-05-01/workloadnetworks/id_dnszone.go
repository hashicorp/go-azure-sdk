// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DnsZoneId{}

// DnsZoneId is a struct representing the Resource ID for a Dns Zone
type DnsZoneId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	DnsZoneId         string
}

// NewDnsZoneID returns a new DnsZoneId struct
func NewDnsZoneID(subscriptionId string, resourceGroupName string, privateCloudName string, dnsZoneId string) DnsZoneId {
	return DnsZoneId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		DnsZoneId:         dnsZoneId,
	}
}

// ParseDnsZoneID parses 'input' into a DnsZoneId
func ParseDnsZoneID(input string) (*DnsZoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(DnsZoneId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DnsZoneId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.DnsZoneId, ok = parsed.Parsed["dnsZoneId"]; !ok {
		return nil, fmt.Errorf("the segment 'dnsZoneId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDnsZoneIDInsensitively parses 'input' case-insensitively into a DnsZoneId
// note: this method should only be used for API response data and not user input
func ParseDnsZoneIDInsensitively(input string) (*DnsZoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(DnsZoneId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DnsZoneId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.DnsZoneId, ok = parsed.Parsed["dnsZoneId"]; !ok {
		return nil, fmt.Errorf("the segment 'dnsZoneId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDnsZoneID checks that 'input' can be parsed as a Dns Zone ID
func ValidateDnsZoneID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDnsZoneID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dns Zone ID
func (id DnsZoneId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/dnsZones/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.DnsZoneId)
}

// Segments returns a slice of Resource ID Segments which comprise this Dns Zone ID
func (id DnsZoneId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticDnsZones", "dnsZones", "dnsZones"),
		resourceids.UserSpecifiedSegment("dnsZoneId", "dnsZoneIdValue"),
	}
}

// String returns a human-readable description of this Dns Zone ID
func (id DnsZoneId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Dns Zone: %q", id.DnsZoneId),
	}
	return fmt.Sprintf("Dns Zone (%s)", strings.Join(components, "\n"))
}
