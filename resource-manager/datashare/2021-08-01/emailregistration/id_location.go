package emailregistration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = LocationId{}

// LocationId is a struct representing the Resource ID for a Location
type LocationId struct {
	LocationName string
}

// NewLocationID returns a new LocationId struct
func NewLocationID(locationName string) LocationId {
	return LocationId{
		LocationName: locationName,
	}
}

// ParseLocationID parses 'input' into a LocationId
func ParseLocationID(input string) (*LocationId, error) {
	parser := resourceids.NewParserFromResourceIdType(LocationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LocationId{}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseLocationIDInsensitively parses 'input' case-insensitively into a LocationId
// note: this method should only be used for API response data and not user input
func ParseLocationIDInsensitively(input string) (*LocationId, error) {
	parser := resourceids.NewParserFromResourceIdType(LocationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LocationId{}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateLocationID checks that 'input' can be parsed as a Location ID
func ValidateLocationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLocationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Location ID
func (id LocationId) ID() string {
	fmtString := "/providers/Microsoft.DataShare/locations/%s"
	return fmt.Sprintf(fmtString, id.LocationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Location ID
func (id LocationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataShare", "Microsoft.DataShare", "Microsoft.DataShare"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
	}
}

// String returns a human-readable description of this Location ID
func (id LocationId) String() string {
	components := []string{
		fmt.Sprintf("Location Name: %q", id.LocationName),
	}
	return fmt.Sprintf("Location (%s)", strings.Join(components, "\n"))
}
