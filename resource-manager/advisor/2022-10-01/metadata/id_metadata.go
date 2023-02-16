package metadata

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = MetadataId{}

// MetadataId is a struct representing the Resource ID for a Metadata
type MetadataId struct {
	MetadataName string
}

// NewMetadataID returns a new MetadataId struct
func NewMetadataID(metadataName string) MetadataId {
	return MetadataId{
		MetadataName: metadataName,
	}
}

// ParseMetadataID parses 'input' into a MetadataId
func ParseMetadataID(input string) (*MetadataId, error) {
	parser := resourceids.NewParserFromResourceIdType(MetadataId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MetadataId{}

	if id.MetadataName, ok = parsed.Parsed["metadataName"]; !ok {
		return nil, fmt.Errorf("the segment 'metadataName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseMetadataIDInsensitively parses 'input' case-insensitively into a MetadataId
// note: this method should only be used for API response data and not user input
func ParseMetadataIDInsensitively(input string) (*MetadataId, error) {
	parser := resourceids.NewParserFromResourceIdType(MetadataId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MetadataId{}

	if id.MetadataName, ok = parsed.Parsed["metadataName"]; !ok {
		return nil, fmt.Errorf("the segment 'metadataName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateMetadataID checks that 'input' can be parsed as a Metadata ID
func ValidateMetadataID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMetadataID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Metadata ID
func (id MetadataId) ID() string {
	fmtString := "/providers/Microsoft.Advisor/metadata/%s"
	return fmt.Sprintf(fmtString, id.MetadataName)
}

// Segments returns a slice of Resource ID Segments which comprise this Metadata ID
func (id MetadataId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAdvisor", "Microsoft.Advisor", "Microsoft.Advisor"),
		resourceids.StaticSegment("staticMetadata", "metadata", "metadata"),
		resourceids.UserSpecifiedSegment("metadataName", "metadataValue"),
	}
}

// String returns a human-readable description of this Metadata ID
func (id MetadataId) String() string {
	components := []string{
		fmt.Sprintf("Metadata Name: %q", id.MetadataName),
	}
	return fmt.Sprintf("Metadata (%s)", strings.Join(components, "\n"))
}
