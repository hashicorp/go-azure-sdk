package workloadnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DhcpConfigurationId{}

// DhcpConfigurationId is a struct representing the Resource ID for a Dhcp Configuration
type DhcpConfigurationId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	DhcpId            string
}

// NewDhcpConfigurationID returns a new DhcpConfigurationId struct
func NewDhcpConfigurationID(subscriptionId string, resourceGroupName string, privateCloudName string, dhcpId string) DhcpConfigurationId {
	return DhcpConfigurationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		DhcpId:            dhcpId,
	}
}

// ParseDhcpConfigurationID parses 'input' into a DhcpConfigurationId
func ParseDhcpConfigurationID(input string) (*DhcpConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DhcpConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DhcpConfigurationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.DhcpId, ok = parsed.Parsed["dhcpId"]; !ok {
		return nil, fmt.Errorf("the segment 'dhcpId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDhcpConfigurationIDInsensitively parses 'input' case-insensitively into a DhcpConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDhcpConfigurationIDInsensitively(input string) (*DhcpConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DhcpConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DhcpConfigurationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.DhcpId, ok = parsed.Parsed["dhcpId"]; !ok {
		return nil, fmt.Errorf("the segment 'dhcpId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDhcpConfigurationID checks that 'input' can be parsed as a Dhcp Configuration ID
func ValidateDhcpConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDhcpConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dhcp Configuration ID
func (id DhcpConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/dhcpConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.DhcpId)
}

// Segments returns a slice of Resource ID Segments which comprise this Dhcp Configuration ID
func (id DhcpConfigurationId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticDhcpConfigurations", "dhcpConfigurations", "dhcpConfigurations"),
		resourceids.UserSpecifiedSegment("dhcpId", "dhcpIdValue"),
	}
}

// String returns a human-readable description of this Dhcp Configuration ID
func (id DhcpConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Dhcp: %q", id.DhcpId),
	}
	return fmt.Sprintf("Dhcp Configuration (%s)", strings.Join(components, "\n"))
}
