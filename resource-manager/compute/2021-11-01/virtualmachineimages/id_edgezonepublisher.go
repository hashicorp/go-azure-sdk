package virtualmachineimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = EdgeZonePublisherId{}

// EdgeZonePublisherId is a struct representing the Resource ID for a Edge Zone Publisher
type EdgeZonePublisherId struct {
	SubscriptionId string
	Location       string
	EdgeZone       string
	PublisherName  string
}

// NewEdgeZonePublisherID returns a new EdgeZonePublisherId struct
func NewEdgeZonePublisherID(subscriptionId string, location string, edgeZone string, publisherName string) EdgeZonePublisherId {
	return EdgeZonePublisherId{
		SubscriptionId: subscriptionId,
		Location:       location,
		EdgeZone:       edgeZone,
		PublisherName:  publisherName,
	}
}

// ParseEdgeZonePublisherID parses 'input' into a EdgeZonePublisherId
func ParseEdgeZonePublisherID(input string) (*EdgeZonePublisherId, error) {
	parser := resourceids.NewParserFromResourceIdType(EdgeZonePublisherId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EdgeZonePublisherId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.EdgeZone, ok = parsed.Parsed["edgeZone"]; !ok {
		return nil, fmt.Errorf("the segment 'edgeZone' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseEdgeZonePublisherIDInsensitively parses 'input' case-insensitively into a EdgeZonePublisherId
// note: this method should only be used for API response data and not user input
func ParseEdgeZonePublisherIDInsensitively(input string) (*EdgeZonePublisherId, error) {
	parser := resourceids.NewParserFromResourceIdType(EdgeZonePublisherId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EdgeZonePublisherId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.EdgeZone, ok = parsed.Parsed["edgeZone"]; !ok {
		return nil, fmt.Errorf("the segment 'edgeZone' was not found in the resource id %q", input)
	}

	if id.PublisherName, ok = parsed.Parsed["publisherName"]; !ok {
		return nil, fmt.Errorf("the segment 'publisherName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateEdgeZonePublisherID checks that 'input' can be parsed as a Edge Zone Publisher ID
func ValidateEdgeZonePublisherID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEdgeZonePublisherID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Edge Zone Publisher ID
func (id EdgeZonePublisherId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/edgeZones/%s/publishers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.EdgeZone, id.PublisherName)
}

// Segments returns a slice of Resource ID Segments which comprise this Edge Zone Publisher ID
func (id EdgeZonePublisherId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticEdgeZones", "edgeZones", "edgeZones"),
		resourceids.UserSpecifiedSegment("edgeZone", "edgeZoneValue"),
		resourceids.StaticSegment("staticPublishers", "publishers", "publishers"),
		resourceids.UserSpecifiedSegment("publisherName", "publisherValue"),
	}
}

// String returns a human-readable description of this Edge Zone Publisher ID
func (id EdgeZonePublisherId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Edge Zone: %q", id.EdgeZone),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
	}
	return fmt.Sprintf("Edge Zone Publisher (%s)", strings.Join(components, "\n"))
}
