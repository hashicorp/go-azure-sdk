package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CloudServicesNetworkId{}

// CloudServicesNetworkId is a struct representing the Resource ID for a Cloud Services Network
type CloudServicesNetworkId struct {
	SubscriptionId           string
	ResourceGroupName        string
	CloudServicesNetworkName string
}

// NewCloudServicesNetworkID returns a new CloudServicesNetworkId struct
func NewCloudServicesNetworkID(subscriptionId string, resourceGroupName string, cloudServicesNetworkName string) CloudServicesNetworkId {
	return CloudServicesNetworkId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		CloudServicesNetworkName: cloudServicesNetworkName,
	}
}

// ParseCloudServicesNetworkID parses 'input' into a CloudServicesNetworkId
func ParseCloudServicesNetworkID(input string) (*CloudServicesNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(CloudServicesNetworkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CloudServicesNetworkId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.CloudServicesNetworkName, ok = parsed.Parsed["cloudServicesNetworkName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "cloudServicesNetworkName", *parsed)
	}

	return &id, nil
}

// ParseCloudServicesNetworkIDInsensitively parses 'input' case-insensitively into a CloudServicesNetworkId
// note: this method should only be used for API response data and not user input
func ParseCloudServicesNetworkIDInsensitively(input string) (*CloudServicesNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(CloudServicesNetworkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CloudServicesNetworkId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.CloudServicesNetworkName, ok = parsed.Parsed["cloudServicesNetworkName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "cloudServicesNetworkName", *parsed)
	}

	return &id, nil
}

// ValidateCloudServicesNetworkID checks that 'input' can be parsed as a Cloud Services Network ID
func ValidateCloudServicesNetworkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCloudServicesNetworkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cloud Services Network ID
func (id CloudServicesNetworkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/cloudServicesNetworks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CloudServicesNetworkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cloud Services Network ID
func (id CloudServicesNetworkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticCloudServicesNetworks", "cloudServicesNetworks", "cloudServicesNetworks"),
		resourceids.UserSpecifiedSegment("cloudServicesNetworkName", "cloudServicesNetworkValue"),
	}
}

// String returns a human-readable description of this Cloud Services Network ID
func (id CloudServicesNetworkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cloud Services Network Name: %q", id.CloudServicesNetworkName),
	}
	return fmt.Sprintf("Cloud Services Network (%s)", strings.Join(components, "\n"))
}
