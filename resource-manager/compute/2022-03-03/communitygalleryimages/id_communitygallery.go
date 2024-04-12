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
	recaser.RegisterResourceId(&CommunityGalleryId{})
}

var _ resourceids.ResourceId = &CommunityGalleryId{}

// CommunityGalleryId is a struct representing the Resource ID for a Community Gallery
type CommunityGalleryId struct {
	SubscriptionId       string
	LocationName         string
	CommunityGalleryName string
}

// NewCommunityGalleryID returns a new CommunityGalleryId struct
func NewCommunityGalleryID(subscriptionId string, locationName string, communityGalleryName string) CommunityGalleryId {
	return CommunityGalleryId{
		SubscriptionId:       subscriptionId,
		LocationName:         locationName,
		CommunityGalleryName: communityGalleryName,
	}
}

// ParseCommunityGalleryID parses 'input' into a CommunityGalleryId
func ParseCommunityGalleryID(input string) (*CommunityGalleryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CommunityGalleryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CommunityGalleryId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCommunityGalleryIDInsensitively parses 'input' case-insensitively into a CommunityGalleryId
// note: this method should only be used for API response data and not user input
func ParseCommunityGalleryIDInsensitively(input string) (*CommunityGalleryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CommunityGalleryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CommunityGalleryId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CommunityGalleryId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateCommunityGalleryID checks that 'input' can be parsed as a Community Gallery ID
func ValidateCommunityGalleryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCommunityGalleryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Community Gallery ID
func (id CommunityGalleryId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/communityGalleries/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.CommunityGalleryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Community Gallery ID
func (id CommunityGalleryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticCommunityGalleries", "communityGalleries", "communityGalleries"),
		resourceids.UserSpecifiedSegment("communityGalleryName", "communityGalleryValue"),
	}
}

// String returns a human-readable description of this Community Gallery ID
func (id CommunityGalleryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Community Gallery Name: %q", id.CommunityGalleryName),
	}
	return fmt.Sprintf("Community Gallery (%s)", strings.Join(components, "\n"))
}
