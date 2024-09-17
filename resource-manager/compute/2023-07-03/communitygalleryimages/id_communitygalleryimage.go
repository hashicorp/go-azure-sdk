package communitygalleryimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CommunityGalleryImageId{})
}

var _ resourceids.ResourceId = &CommunityGalleryImageId{}

// CommunityGalleryImageId is a struct representing the Resource ID for a Community Gallery Image
type CommunityGalleryImageId struct {
	SubscriptionId       string
	LocationName         string
	CommunityGalleryName string
	ImageName            string
}

// NewCommunityGalleryImageID returns a new CommunityGalleryImageId struct
func NewCommunityGalleryImageID(subscriptionId string, locationName string, communityGalleryName string, imageName string) CommunityGalleryImageId {
	return CommunityGalleryImageId{
		SubscriptionId:       subscriptionId,
		LocationName:         locationName,
		CommunityGalleryName: communityGalleryName,
		ImageName:            imageName,
	}
}

// ParseCommunityGalleryImageID parses 'input' into a CommunityGalleryImageId
func ParseCommunityGalleryImageID(input string) (*CommunityGalleryImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CommunityGalleryImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CommunityGalleryImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCommunityGalleryImageIDInsensitively parses 'input' case-insensitively into a CommunityGalleryImageId
// note: this method should only be used for API response data and not user input
func ParseCommunityGalleryImageIDInsensitively(input string) (*CommunityGalleryImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CommunityGalleryImageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CommunityGalleryImageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CommunityGalleryImageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.CommunityGalleryName, ok = input.Parsed["communityGalleryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "communityGalleryName", input)
	}

	if id.ImageName, ok = input.Parsed["imageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "imageName", input)
	}

	return nil
}

// ValidateCommunityGalleryImageID checks that 'input' can be parsed as a Community Gallery Image ID
func ValidateCommunityGalleryImageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCommunityGalleryImageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Community Gallery Image ID
func (id CommunityGalleryImageId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/communityGalleries/%s/images/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.CommunityGalleryName, id.ImageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Community Gallery Image ID
func (id CommunityGalleryImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticCommunityGalleries", "communityGalleries", "communityGalleries"),
		resourceids.UserSpecifiedSegment("communityGalleryName", "communityGalleryValue"),
		resourceids.StaticSegment("staticImages", "images", "images"),
		resourceids.UserSpecifiedSegment("imageName", "imageValue"),
	}
}

// String returns a human-readable description of this Community Gallery Image ID
func (id CommunityGalleryImageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Community Gallery Name: %q", id.CommunityGalleryName),
		fmt.Sprintf("Image Name: %q", id.ImageName),
	}
	return fmt.Sprintf("Community Gallery Image (%s)", strings.Join(components, "\n"))
}
