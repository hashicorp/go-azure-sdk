package sharedgalleryimageversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ImageId{}

// ImageId is a struct representing the Resource ID for a Image
type ImageId struct {
	SubscriptionId    string
	LocationName      string
	SharedGalleryName string
	ImageName         string
}

// NewImageID returns a new ImageId struct
func NewImageID(subscriptionId string, locationName string, sharedGalleryName string, imageName string) ImageId {
	return ImageId{
		SubscriptionId:    subscriptionId,
		LocationName:      locationName,
		SharedGalleryName: sharedGalleryName,
		ImageName:         imageName,
	}
}

// ParseImageID parses 'input' into a ImageId
func ParseImageID(input string) (*ImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(ImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ImageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseImageIDInsensitively parses 'input' case-insensitively into a ImageId
// note: this method should only be used for API response data and not user input
func ParseImageIDInsensitively(input string) (*ImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(ImageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ImageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ImageId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateImageID checks that 'input' can be parsed as a Image ID
func ValidateImageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseImageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Image ID
func (id ImageId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/sharedGalleries/%s/images/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.SharedGalleryName, id.ImageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Image ID
func (id ImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticSharedGalleries", "sharedGalleries", "sharedGalleries"),
		resourceids.UserSpecifiedSegment("sharedGalleryName", "sharedGalleryValue"),
		resourceids.StaticSegment("staticImages", "images", "images"),
		resourceids.UserSpecifiedSegment("imageName", "imageValue"),
	}
}

// String returns a human-readable description of this Image ID
func (id ImageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Shared Gallery Name: %q", id.SharedGalleryName),
		fmt.Sprintf("Image Name: %q", id.ImageName),
	}
	return fmt.Sprintf("Image (%s)", strings.Join(components, "\n"))
}
