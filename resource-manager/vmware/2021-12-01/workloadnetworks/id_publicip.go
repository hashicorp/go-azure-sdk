// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadnetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PublicIPId{}

// PublicIPId is a struct representing the Resource ID for a Public I P
type PublicIPId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	PublicIPId        string
}

// NewPublicIPID returns a new PublicIPId struct
func NewPublicIPID(subscriptionId string, resourceGroupName string, privateCloudName string, publicIPId string) PublicIPId {
	return PublicIPId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		PublicIPId:        publicIPId,
	}
}

// ParsePublicIPID parses 'input' into a PublicIPId
func ParsePublicIPID(input string) (*PublicIPId, error) {
	parser := resourceids.NewParserFromResourceIdType(PublicIPId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PublicIPId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.PublicIPId, ok = parsed.Parsed["publicIPId"]; !ok {
		return nil, fmt.Errorf("the segment 'publicIPId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePublicIPIDInsensitively parses 'input' case-insensitively into a PublicIPId
// note: this method should only be used for API response data and not user input
func ParsePublicIPIDInsensitively(input string) (*PublicIPId, error) {
	parser := resourceids.NewParserFromResourceIdType(PublicIPId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PublicIPId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.PublicIPId, ok = parsed.Parsed["publicIPId"]; !ok {
		return nil, fmt.Errorf("the segment 'publicIPId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePublicIPID checks that 'input' can be parsed as a Public I P ID
func ValidatePublicIPID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePublicIPID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Public I P ID
func (id PublicIPId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/publicIPs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.PublicIPId)
}

// Segments returns a slice of Resource ID Segments which comprise this Public I P ID
func (id PublicIPId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticPublicIPs", "publicIPs", "publicIPs"),
		resourceids.UserSpecifiedSegment("publicIPId", "publicIPIdValue"),
	}
}

// String returns a human-readable description of this Public I P ID
func (id PublicIPId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Public I P: %q", id.PublicIPId),
	}
	return fmt.Sprintf("Public I P (%s)", strings.Join(components, "\n"))
}
