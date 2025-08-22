package vmwares

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PublicIPId{})
}

var _ resourceids.ResourceId = &PublicIPId{}

// PublicIPId is a struct representing the Resource ID for a Public IP
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
	parser := resourceids.NewParserFromResourceIdType(&PublicIPId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PublicIPId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePublicIPIDInsensitively parses 'input' case-insensitively into a PublicIPId
// note: this method should only be used for API response data and not user input
func ParsePublicIPIDInsensitively(input string) (*PublicIPId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PublicIPId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PublicIPId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PublicIPId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PrivateCloudName, ok = input.Parsed["privateCloudName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateCloudName", input)
	}

	if id.PublicIPId, ok = input.Parsed["publicIPId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "publicIPId", input)
	}

	return nil
}

// ValidatePublicIPID checks that 'input' can be parsed as a Public IP ID
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

// ID returns the formatted Public IP ID
func (id PublicIPId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/publicIPs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.PublicIPId)
}

// Segments returns a slice of Resource ID Segments which comprise this Public IP ID
func (id PublicIPId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudName"),
		resourceids.StaticSegment("staticWorkloadNetworks", "workloadNetworks", "workloadNetworks"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticPublicIPs", "publicIPs", "publicIPs"),
		resourceids.UserSpecifiedSegment("publicIPId", "publicIPId"),
	}
}

// String returns a human-readable description of this Public IP ID
func (id PublicIPId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Public IP: %q", id.PublicIPId),
	}
	return fmt.Sprintf("Public IP (%s)", strings.Join(components, "\n"))
}
