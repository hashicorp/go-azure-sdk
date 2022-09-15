package virtualmachineimages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = EdgeZoneId{}

// EdgeZoneId is a struct representing the Resource ID for a Edge Zone
type EdgeZoneId struct {
	SubscriptionId string
	Location       string
	EdgeZone       string
}

// NewEdgeZoneID returns a new EdgeZoneId struct
func NewEdgeZoneID(subscriptionId string, location string, edgeZone string) EdgeZoneId {
	return EdgeZoneId{
		SubscriptionId: subscriptionId,
		Location:       location,
		EdgeZone:       edgeZone,
	}
}

// ParseEdgeZoneID parses 'input' into a EdgeZoneId
func ParseEdgeZoneID(input string) (*EdgeZoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(EdgeZoneId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EdgeZoneId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.EdgeZone, ok = parsed.Parsed["edgeZone"]; !ok {
		return nil, fmt.Errorf("the segment 'edgeZone' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseEdgeZoneIDInsensitively parses 'input' case-insensitively into a EdgeZoneId
// note: this method should only be used for API response data and not user input
func ParseEdgeZoneIDInsensitively(input string) (*EdgeZoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(EdgeZoneId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := EdgeZoneId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.Location, ok = parsed.Parsed["location"]; !ok {
		return nil, fmt.Errorf("the segment 'location' was not found in the resource id %q", input)
	}

	if id.EdgeZone, ok = parsed.Parsed["edgeZone"]; !ok {
		return nil, fmt.Errorf("the segment 'edgeZone' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateEdgeZoneID checks that 'input' can be parsed as a Edge Zone ID
func ValidateEdgeZoneID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEdgeZoneID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Edge Zone ID
func (id EdgeZoneId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Compute/locations/%s/edgeZones/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.Location, id.EdgeZone)
}

// Segments returns a slice of Resource ID Segments which comprise this Edge Zone ID
func (id EdgeZoneId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCompute", "Microsoft.Compute", "Microsoft.Compute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("location", "locationValue"),
		resourceids.StaticSegment("staticEdgeZones", "edgeZones", "edgeZones"),
		resourceids.UserSpecifiedSegment("edgeZone", "edgeZoneValue"),
	}
}

// String returns a human-readable description of this Edge Zone ID
func (id EdgeZoneId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location: %q", id.Location),
		fmt.Sprintf("Edge Zone: %q", id.EdgeZone),
	}
	return fmt.Sprintf("Edge Zone (%s)", strings.Join(components, "\n"))
}
