package workloadnetworksegments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SegmentId{})
}

var _ resourceids.ResourceId = &SegmentId{}

// SegmentId is a struct representing the Resource ID for a Segment
type SegmentId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	SegmentId         string
}

// NewSegmentID returns a new SegmentId struct
func NewSegmentID(subscriptionId string, resourceGroupName string, privateCloudName string, segmentId string) SegmentId {
	return SegmentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		SegmentId:         segmentId,
	}
}

// ParseSegmentID parses 'input' into a SegmentId
func ParseSegmentID(input string) (*SegmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SegmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SegmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSegmentIDInsensitively parses 'input' case-insensitively into a SegmentId
// note: this method should only be used for API response data and not user input
func ParseSegmentIDInsensitively(input string) (*SegmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SegmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SegmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SegmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SegmentId, ok = input.Parsed["segmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "segmentId", input)
	}

	return nil
}

// ValidateSegmentID checks that 'input' can be parsed as a Segment ID
func ValidateSegmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSegmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Segment ID
func (id SegmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/workloadNetworks/default/segments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.SegmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Segment ID
func (id SegmentId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticSegments", "segments", "segments"),
		resourceids.UserSpecifiedSegment("segmentId", "segmentId"),
	}
}

// String returns a human-readable description of this Segment ID
func (id SegmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Segment: %q", id.SegmentId),
	}
	return fmt.Sprintf("Segment (%s)", strings.Join(components, "\n"))
}
