package extensions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PublisherId{})
}

var _ resourceids.ResourceId = &PublisherId{}

// PublisherId is a struct representing the Resource ID for a Publisher
type PublisherId struct {
	LocationName  string
	PublisherName string
}

// NewPublisherID returns a new PublisherId struct
func NewPublisherID(locationName string, publisherName string) PublisherId {
	return PublisherId{
		LocationName:  locationName,
		PublisherName: publisherName,
	}
}

// ParsePublisherID parses 'input' into a PublisherId
func ParsePublisherID(input string) (*PublisherId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PublisherId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PublisherId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePublisherIDInsensitively parses 'input' case-insensitively into a PublisherId
// note: this method should only be used for API response data and not user input
func ParsePublisherIDInsensitively(input string) (*PublisherId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PublisherId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PublisherId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PublisherId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.PublisherName, ok = input.Parsed["publisherName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "publisherName", input)
	}

	return nil
}

// ValidatePublisherID checks that 'input' can be parsed as a Publisher ID
func ValidatePublisherID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePublisherID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Publisher ID
func (id PublisherId) ID() string {
	fmtString := "/providers/Microsoft.HybridCompute/locations/%s/publishers/%s"
	return fmt.Sprintf(fmtString, id.LocationName, id.PublisherName)
}

// Segments returns a slice of Resource ID Segments which comprise this Publisher ID
func (id PublisherId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHybridCompute", "Microsoft.HybridCompute", "Microsoft.HybridCompute"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticPublishers", "publishers", "publishers"),
		resourceids.UserSpecifiedSegment("publisherName", "publisherName"),
	}
}

// String returns a human-readable description of this Publisher ID
func (id PublisherId) String() string {
	components := []string{
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Publisher Name: %q", id.PublisherName),
	}
	return fmt.Sprintf("Publisher (%s)", strings.Join(components, "\n"))
}
