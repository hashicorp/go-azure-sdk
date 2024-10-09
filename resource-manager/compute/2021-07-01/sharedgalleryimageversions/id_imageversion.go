package sharedgalleryimageversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ImageVersionId{})
}

var _ resourceids.ResourceId = &ImageVersionId{}

// ImageVersionId is a struct representing the Resource ID for a Image Version
type ImageVersionId struct {
	SubscriptionId    string
	LocationName      string
	SharedGalleryName string
	ImageName         string
	VersionName       string
}

// NewImageVersionID returns a new ImageVersionId struct
func NewImageVersionID(subscriptionId string, locationName string, sharedGalleryName string, imageName string, versionName string) ImageVersionId {
	return ImageVersionId{
		SubscriptionId:    subscriptionId,
		LocationName:      locationName,
		SharedGalleryName: sharedGalleryName,
		ImageName:         imageName,
		VersionName:       versionName,
	}
}

// ParseImageVersionID parses 'input' into a ImageVersionId
func ParseImageVersionID(input string) (*ImageVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ImageVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ImageVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseImageVersionIDInsensitively parses 'input' case-insensitively into a ImageVersionId
// note: this method should only be used for API response data and not user input
func ParseImageVersionIDInsensitively(input string) (*ImageVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ImageVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ImageVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ImageVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateImageVersionID checks that 'input' can be parsed as a Image Version ID
func ValidateImageVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseImageVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Image Version ID
func (id ImageVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/sharedGalleries/%s/images/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.SharedGalleryName, id.ImageName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Image Version ID
func (id ImageVersionId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Image Version ID
func (id ImageVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Shared Gallery Name: %q", id.SharedGalleryName),
		fmt.Sprintf("Image Name: %q", id.ImageName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Image Version (%s)", strings.Join(components, "\n"))
}
