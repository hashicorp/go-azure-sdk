package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DestinationIdId{}

// DestinationIdId is a struct representing the Resource ID for a Destination Id
type DestinationIdId struct {
	BaseURI       string
	DestinationId string
}

// NewDestinationIdID returns a new DestinationIdId struct
func NewDestinationIdID(baseURI string, destinationId string) DestinationIdId {
	return DestinationIdId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DestinationId: destinationId,
	}
}

// ParseDestinationIdID parses 'input' into a DestinationIdId
func ParseDestinationIdID(input string) (*DestinationIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DestinationIdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DestinationIdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDestinationIdIDInsensitively parses 'input' case-insensitively into a DestinationIdId
// note: this method should only be used for API response data and not user input
func ParseDestinationIdIDInsensitively(input string) (*DestinationIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DestinationIdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DestinationIdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DestinationIdId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DestinationId, ok = input.Parsed["destinationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "destinationId", input)
	}

	return nil
}

// ValidateDestinationIdID checks that 'input' can be parsed as a Destination Id ID
func ValidateDestinationIdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDestinationIdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Destination Id ID
func (id DestinationIdId) ID() string {
	fmtString := "%s/dataExport/destinations/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DestinationId)
}

// Path returns the formatted Destination Id ID without the BaseURI
func (id DestinationIdId) Path() string {
	fmtString := "/dataExport/destinations/%s"
	return fmt.Sprintf(fmtString, id.DestinationId)
}

// PathElements returns the values of Destination Id ID Segments without the BaseURI
func (id DestinationIdId) PathElements() []any {
	return []any{id.DestinationId}
}

// Segments returns a slice of Resource ID Segments which comprise this Destination Id ID
func (id DestinationIdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDataExport", "dataExport", "dataExport"),
		resourceids.StaticSegment("staticDestinations", "destinations", "destinations"),
		resourceids.UserSpecifiedSegment("destinationId", "destinationId"),
	}
}

// String returns a human-readable description of this Destination Id ID
func (id DestinationIdId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Destination: %q", id.DestinationId),
	}
	return fmt.Sprintf("Destination Id (%s)", strings.Join(components, "\n"))
}
