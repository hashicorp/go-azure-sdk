package galleryimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ImageId{}

// ImageId is a struct representing the Resource ID for a Image
type ImageId struct {
	SubscriptionId    string
	Location          string
	GalleryUniqueName string
	GalleryImageName  string
}

// NewImageID returns a new ImageId struct
func NewImageID(subscriptionId string, location string, galleryUniqueName string, galleryImageName string) ImageId {
	return ImageId{
		SubscriptionId:    subscriptionId,
		Location:          location,
		GalleryUniqueName: galleryUniqueName,
		GalleryImageName:  galleryImageName,
	}
}

// ParseImageID parses 'input' into a ImageId
func ParseImageID(input string) (*ImageId, error) {
	parser := resourceids.NewParserFromResourceIdType(ImageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ImageId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.GalleryUniqueName, ok = parsed.Parsed["galleryUniqueName"]; !ok {
		return nil, fmt.Errorf("the segment 'galleryUniqueName' was not found in the resource id %q", input)
	}

	if id.GalleryImageName, ok = parsed.Parsed["galleryImageName"]; !ok {
		return nil, fmt.Errorf("the segment 'galleryImageName' was not found in the resource id %q", input)
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

	var ok bool
	id := ImageId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.GalleryUniqueName, ok = parsed.Parsed["galleryUniqueName"]; !ok {
		return nil, fmt.Errorf("the segment 'galleryUniqueName' was not found in the resource id %q", input)
	}

	if id.GalleryImageName, ok = parsed.Parsed["galleryImageName"]; !ok {
		return nil, fmt.Errorf("the segment 'galleryImageName' was not found in the resource id %q", input)
	}

	return &id, nil
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
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.GalleryUniqueName, id.GalleryImageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Image ID
func (id ImageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticSharedGalleries", "sharedGalleries", "sharedGalleries"),
		resourceids.UserSpecifiedSegment("galleryUniqueName", "galleryUniqueValue"),
		resourceids.StaticSegment("staticImages", "images", "images"),
		resourceids.UserSpecifiedSegment("galleryImageName", "galleryImageValue"),
	}
}

// String returns a human-readable description of this Image ID
func (id ImageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Gallery Unique Name: %q", id.GalleryUniqueName),
		fmt.Sprintf("Gallery Image Name: %q", id.GalleryImageName),
	}
	return fmt.Sprintf("Image (%s)", strings.Join(components, "\n"))
}
