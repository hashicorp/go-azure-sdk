package certificates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ThumbprintId{}

// ThumbprintId is a struct representing the Resource ID for a Thumbprint
type ThumbprintId struct {
	BaseURI                 string
	ThumbprintAlgorithmName string
	ThumbprintName          string
}

// NewThumbprintID returns a new ThumbprintId struct
func NewThumbprintID(baseURI string, thumbprintAlgorithmName string, thumbprintName string) ThumbprintId {
	return ThumbprintId{
		BaseURI:                 strings.TrimSuffix(baseURI, "/"),
		ThumbprintAlgorithmName: thumbprintAlgorithmName,
		ThumbprintName:          thumbprintName,
	}
}

// ParseThumbprintID parses 'input' into a ThumbprintId
func ParseThumbprintID(input string) (*ThumbprintId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ThumbprintId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ThumbprintId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseThumbprintIDInsensitively parses 'input' case-insensitively into a ThumbprintId
// note: this method should only be used for API response data and not user input
func ParseThumbprintIDInsensitively(input string) (*ThumbprintId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ThumbprintId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ThumbprintId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ThumbprintId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.ThumbprintAlgorithmName, ok = input.Parsed["thumbprintAlgorithmName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "thumbprintAlgorithmName", input)
	}

	if id.ThumbprintName, ok = input.Parsed["thumbprintName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "thumbprintName", input)
	}

	return nil
}

// ValidateThumbprintID checks that 'input' can be parsed as a Thumbprint ID
func ValidateThumbprintID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseThumbprintID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Thumbprint ID
func (id ThumbprintId) ID() string {
	fmtString := "%s/thumbprintAlgorithm/%s/thumbprint/%s"
	return fmt.Sprintf(fmtString, id.BaseURI, id.ThumbprintAlgorithmName, id.ThumbprintName)
}

// Path returns the formatted Thumbprint ID without the BaseURI
func (id ThumbprintId) Path() string {
	fmtString := "/thumbprintAlgorithm/%s/thumbprint/%s"
	return fmt.Sprintf(fmtString, id.ThumbprintAlgorithmName, id.ThumbprintName)
}

// PathElements returns the values of Thumbprint ID Segments without the BaseURI
func (id ThumbprintId) PathElements() []any {
	return []any{id.ThumbprintAlgorithmName, id.ThumbprintName}
}

// Segments returns a slice of Resource ID Segments which comprise this Thumbprint ID
func (id ThumbprintId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint_url"),
		resourceids.StaticSegment("thumbprintAlgorithm", "thumbprintAlgorithm", "thumbprintAlgorithm"),
		resourceids.UserSpecifiedSegment("thumbprintAlgorithmName", "thumbprintAlgorithmName"),
		resourceids.StaticSegment("thumbprint", "thumbprint", "thumbprint"),
		resourceids.UserSpecifiedSegment("thumbprintName", "thumbprintName"),
	}
}

// String returns a human-readable description of this Thumbprint ID
func (id ThumbprintId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Thumbprint Algorithm Name: %q", id.ThumbprintAlgorithmName),
		fmt.Sprintf("Thumbprint Name: %q", id.ThumbprintName),
	}
	return fmt.Sprintf("Thumbprint (%s)", strings.Join(components, "\n"))
}
