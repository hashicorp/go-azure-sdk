package cloudlinks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CloudLinkId{})
}

var _ resourceids.ResourceId = &CloudLinkId{}

// CloudLinkId is a struct representing the Resource ID for a Cloud Link
type CloudLinkId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	CloudLinkName     string
}

// NewCloudLinkID returns a new CloudLinkId struct
func NewCloudLinkID(subscriptionId string, resourceGroupName string, privateCloudName string, cloudLinkName string) CloudLinkId {
	return CloudLinkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		CloudLinkName:     cloudLinkName,
	}
}

// ParseCloudLinkID parses 'input' into a CloudLinkId
func ParseCloudLinkID(input string) (*CloudLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CloudLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CloudLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCloudLinkIDInsensitively parses 'input' case-insensitively into a CloudLinkId
// note: this method should only be used for API response data and not user input
func ParseCloudLinkIDInsensitively(input string) (*CloudLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CloudLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CloudLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CloudLinkId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.CloudLinkName, ok = input.Parsed["cloudLinkName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudLinkName", input)
	}

	return nil
}

// ValidateCloudLinkID checks that 'input' can be parsed as a Cloud Link ID
func ValidateCloudLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCloudLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cloud Link ID
func (id CloudLinkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/cloudLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.CloudLinkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cloud Link ID
func (id CloudLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudName"),
		resourceids.StaticSegment("staticCloudLinks", "cloudLinks", "cloudLinks"),
		resourceids.UserSpecifiedSegment("cloudLinkName", "cloudLinkName"),
	}
}

// String returns a human-readable description of this Cloud Link ID
func (id CloudLinkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Cloud Link Name: %q", id.CloudLinkName),
	}
	return fmt.Sprintf("Cloud Link (%s)", strings.Join(components, "\n"))
}
