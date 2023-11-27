package communitygalleryimageversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = CommunityGalleryImageVersionId{}

// CommunityGalleryImageVersionId is a struct representing the Resource ID for a Community Gallery Image Version
type CommunityGalleryImageVersionId struct {
	SubscriptionId       string
	LocationName         string
	CommunityGalleryName string
	ImageName            string
	VersionName          string
}

// NewCommunityGalleryImageVersionID returns a new CommunityGalleryImageVersionId struct
func NewCommunityGalleryImageVersionID(subscriptionId string, locationName string, communityGalleryName string, imageName string, versionName string) CommunityGalleryImageVersionId {
	return CommunityGalleryImageVersionId{
		SubscriptionId:       subscriptionId,
		LocationName:         locationName,
		CommunityGalleryName: communityGalleryName,
		ImageName:            imageName,
		VersionName:          versionName,
	}
}

// ParseCommunityGalleryImageVersionID parses 'input' into a CommunityGalleryImageVersionId
func ParseCommunityGalleryImageVersionID(input string) (*CommunityGalleryImageVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(CommunityGalleryImageVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CommunityGalleryImageVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCommunityGalleryImageVersionIDInsensitively parses 'input' case-insensitively into a CommunityGalleryImageVersionId
// note: this method should only be used for API response data and not user input
func ParseCommunityGalleryImageVersionIDInsensitively(input string) (*CommunityGalleryImageVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(CommunityGalleryImageVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CommunityGalleryImageVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CommunityGalleryImageVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateCommunityGalleryImageVersionID checks that 'input' can be parsed as a Community Gallery Image Version ID
func ValidateCommunityGalleryImageVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCommunityGalleryImageVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Community Gallery Image Version ID
func (id CommunityGalleryImageVersionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/communityGalleries/%s/images/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.CommunityGalleryName, id.ImageName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Community Gallery Image Version ID
func (id CommunityGalleryImageVersionId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionValue"),
	}
}

// String returns a human-readable description of this Community Gallery Image Version ID
func (id CommunityGalleryImageVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Community Gallery Name: %q", id.CommunityGalleryName),
		fmt.Sprintf("Image Name: %q", id.ImageName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Community Gallery Image Version (%s)", strings.Join(components, "\n"))
}
