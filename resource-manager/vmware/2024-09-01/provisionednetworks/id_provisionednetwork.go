package provisionednetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ProvisionedNetworkId{})
}

var _ resourceids.ResourceId = &ProvisionedNetworkId{}

// ProvisionedNetworkId is a struct representing the Resource ID for a Provisioned Network
type ProvisionedNetworkId struct {
	SubscriptionId         string
	ResourceGroupName      string
	PrivateCloudName       string
	ProvisionedNetworkName string
}

// NewProvisionedNetworkID returns a new ProvisionedNetworkId struct
func NewProvisionedNetworkID(subscriptionId string, resourceGroupName string, privateCloudName string, provisionedNetworkName string) ProvisionedNetworkId {
	return ProvisionedNetworkId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		PrivateCloudName:       privateCloudName,
		ProvisionedNetworkName: provisionedNetworkName,
	}
}

// ParseProvisionedNetworkID parses 'input' into a ProvisionedNetworkId
func ParseProvisionedNetworkID(input string) (*ProvisionedNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProvisionedNetworkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProvisionedNetworkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProvisionedNetworkIDInsensitively parses 'input' case-insensitively into a ProvisionedNetworkId
// note: this method should only be used for API response data and not user input
func ParseProvisionedNetworkIDInsensitively(input string) (*ProvisionedNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProvisionedNetworkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProvisionedNetworkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProvisionedNetworkId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ProvisionedNetworkName, ok = input.Parsed["provisionedNetworkName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "provisionedNetworkName", input)
	}

	return nil
}

// ValidateProvisionedNetworkID checks that 'input' can be parsed as a Provisioned Network ID
func ValidateProvisionedNetworkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProvisionedNetworkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provisioned Network ID
func (id ProvisionedNetworkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/provisionedNetworks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.ProvisionedNetworkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provisioned Network ID
func (id ProvisionedNetworkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudName"),
		resourceids.StaticSegment("staticProvisionedNetworks", "provisionedNetworks", "provisionedNetworks"),
		resourceids.UserSpecifiedSegment("provisionedNetworkName", "provisionedNetworkName"),
	}
}

// String returns a human-readable description of this Provisioned Network ID
func (id ProvisionedNetworkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Provisioned Network Name: %q", id.ProvisionedNetworkName),
	}
	return fmt.Sprintf("Provisioned Network (%s)", strings.Join(components, "\n"))
}
