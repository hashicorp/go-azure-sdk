package sharedgalleryimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SharedGalleryId{}

// SharedGalleryId is a struct representing the Resource ID for a Shared Gallery
type SharedGalleryId struct {
	SubscriptionId    string
	Location          string
	GalleryUniqueName string
}

// NewSharedGalleryID returns a new SharedGalleryId struct
func NewSharedGalleryID(subscriptionId string, location string, galleryUniqueName string) SharedGalleryId {
	return SharedGalleryId{
		SubscriptionId:    subscriptionId,
		Location:          location,
		GalleryUniqueName: galleryUniqueName,
	}
}

// ParseSharedGalleryID parses 'input' into a SharedGalleryId
func ParseSharedGalleryID(input string) (*SharedGalleryId, error) {
	parser := resourceids.NewParserFromResourceIdType(SharedGalleryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SharedGalleryId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.GalleryUniqueName, ok = parsed.Parsed["galleryUniqueName"]; !ok {
		return nil, fmt.Errorf("the segment 'galleryUniqueName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseSharedGalleryIDInsensitively parses 'input' case-insensitively into a SharedGalleryId
// note: this method should only be used for API response data and not user input
func ParseSharedGalleryIDInsensitively(input string) (*SharedGalleryId, error) {
	parser := resourceids.NewParserFromResourceIdType(SharedGalleryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SharedGalleryId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.GalleryUniqueName, ok = parsed.Parsed["galleryUniqueName"]; !ok {
		return nil, fmt.Errorf("the segment 'galleryUniqueName' was not found in the resource id %q", input)
	}

	return &id, nil
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
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.GalleryUniqueName)
}

// Segments returns a slice of Resource ID Segments which comprise this Shared Gallery ID
func (id SharedGalleryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticSharedGalleries", "sharedGalleries", "sharedGalleries"),
		resourceids.UserSpecifiedSegment("galleryUniqueName", "galleryUniqueValue"),
	}
}

// String returns a human-readable description of this Shared Gallery ID
func (id SharedGalleryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Gallery Unique Name: %q", id.GalleryUniqueName),
	}
	return fmt.Sprintf("Shared Gallery (%s)", strings.Join(components, "\n"))
}
