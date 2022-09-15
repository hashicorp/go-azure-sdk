package communitygalleries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CommunityGalleryId{}

// CommunityGalleryId is a struct representing the Resource ID for a Community Gallery
type CommunityGalleryId struct {
	SubscriptionId    string
	Location          string
	PublicGalleryName string
}

// NewCommunityGalleryID returns a new CommunityGalleryId struct
func NewCommunityGalleryID(subscriptionId string, location string, publicGalleryName string) CommunityGalleryId {
	return CommunityGalleryId{
		SubscriptionId:    subscriptionId,
		Location:          location,
		PublicGalleryName: publicGalleryName,
	}
}

// ParseCommunityGalleryID parses 'input' into a CommunityGalleryId
func ParseCommunityGalleryID(input string) (*CommunityGalleryId, error) {
	parser := resourceids.NewParserFromResourceIdType(CommunityGalleryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CommunityGalleryId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.PublicGalleryName, ok = parsed.Parsed["publicGalleryName"]; !ok {
		return nil, fmt.Errorf("the segment 'publicGalleryName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCommunityGalleryIDInsensitively parses 'input' case-insensitively into a CommunityGalleryId
// note: this method should only be used for API response data and not user input
func ParseCommunityGalleryIDInsensitively(input string) (*CommunityGalleryId, error) {
	parser := resourceids.NewParserFromResourceIdType(CommunityGalleryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CommunityGalleryId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.PublicGalleryName, ok = parsed.Parsed["publicGalleryName"]; !ok {
		return nil, fmt.Errorf("the segment 'publicGalleryName' was not found in the resource id %q", input)
	}

	return &id, nil
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
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.PublicGalleryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Community Gallery ID
func (id CommunityGalleryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticCommunityGalleries", "communityGalleries", "communityGalleries"),
		resourceids.UserSpecifiedSegment("publicGalleryName", "publicGalleryValue"),
	}
}

// String returns a human-readable description of this Community Gallery ID
func (id CommunityGalleryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Public Gallery Name: %q", id.PublicGalleryName),
	}
	return fmt.Sprintf("Community Gallery (%s)", strings.Join(components, "\n"))
}
