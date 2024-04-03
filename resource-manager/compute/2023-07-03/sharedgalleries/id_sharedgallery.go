package sharedgalleries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SharedGalleryId{}

// SharedGalleryId is a struct representing the Resource ID for a Shared Gallery
type SharedGalleryId struct {
	SubscriptionId    string
	LocationName      string
	SharedGalleryName string
}

// NewSharedGalleryID returns a new SharedGalleryId struct
func NewSharedGalleryID(subscriptionId string, locationName string, sharedGalleryName string) SharedGalleryId {
	return SharedGalleryId{
		SubscriptionId:    subscriptionId,
		LocationName:      locationName,
		SharedGalleryName: sharedGalleryName,
	}
}

// ParseSharedGalleryID parses 'input' into a SharedGalleryId
func ParseSharedGalleryID(input string) (*SharedGalleryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SharedGalleryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SharedGalleryId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSharedGalleryIDInsensitively parses 'input' case-insensitively into a SharedGalleryId
// note: this method should only be used for API response data and not user input
func ParseSharedGalleryIDInsensitively(input string) (*SharedGalleryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SharedGalleryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SharedGalleryId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SharedGalleryId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateSharedGalleryID checks that 'input' can be parsed as a Shared Gallery ID
func ValidateSharedGalleryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSharedGalleryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Shared Gallery ID
func (id SharedGalleryId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/sharedGalleries/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.SharedGalleryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Shared Gallery ID
func (id SharedGalleryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticSharedGalleries", "sharedGalleries", "sharedGalleries"),
		resourceids.UserSpecifiedSegment("sharedGalleryName", "sharedGalleryValue"),
	}
}

// String returns a human-readable description of this Shared Gallery ID
func (id SharedGalleryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Shared Gallery Name: %q", id.SharedGalleryName),
	}
	return fmt.Sprintf("Shared Gallery (%s)", strings.Join(components, "\n"))
}
