package sharedgalleryimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SharedGalleryImageId{})
}

var _ resourceids.ResourceId = &SharedGalleryImageId{}

// SharedGalleryImageId is a struct representing the Resource ID for a Shared Gallery Image
type SharedGalleryImageId struct {
	SubscriptionId    string
	LocationName      string
	SharedGalleryName string
	ImageName         string
}

// NewSharedGalleryImageID returns a new SharedGalleryImageId struct
func NewSharedGalleryImageID(subscriptionId string, locationName string, sharedGalleryName string, imageName string) SharedGalleryImageId {
	return SharedGalleryImageId{
		SubscriptionId:    subscriptionId,
		LocationName:      locationName,
		SharedGalleryName: sharedGalleryName,
		ImageName:         imageName,
	}
}

// ParseSharedGalleryImageID parses 'input' into a SharedGalleryImageId
func ParseSharedGalleryImageID(input string) (*SharedGalleryImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SharedGalleryImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SharedGalleryImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSharedGalleryImageIDInsensitively parses 'input' case-insensitively into a SharedGalleryImageId
// note: this method should only be used for API response data and not user input
func ParseSharedGalleryImageIDInsensitively(input string) (*SharedGalleryImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SharedGalleryImageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SharedGalleryImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SharedGalleryImageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.SharedGalleryName, ok = input.Parsed["sharedGalleryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sharedGalleryName", input)
	}

	if id.ImageName, ok = input.Parsed["imageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "imageName", input)
	}

	return nil
}

// ValidateSharedGalleryImageID checks that 'input' can be parsed as a Shared Gallery Image ID
func ValidateSharedGalleryImageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSharedGalleryImageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Shared Gallery Image ID
func (id SharedGalleryImageId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/sharedGalleries/%s/images/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.SharedGalleryName, id.ImageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Shared Gallery Image ID
func (id SharedGalleryImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticSharedGalleries", "sharedGalleries", "sharedGalleries"),
		resourceids.UserSpecifiedSegment("sharedGalleryName", "sharedGalleryName"),
		resourceids.StaticSegment("staticImages", "images", "images"),
		resourceids.UserSpecifiedSegment("imageName", "imageName"),
	}
}

// String returns a human-readable description of this Shared Gallery Image ID
func (id SharedGalleryImageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Shared Gallery Name: %q", id.SharedGalleryName),
		fmt.Sprintf("Image Name: %q", id.ImageName),
	}
	return fmt.Sprintf("Shared Gallery Image (%s)", strings.Join(components, "\n"))
}
