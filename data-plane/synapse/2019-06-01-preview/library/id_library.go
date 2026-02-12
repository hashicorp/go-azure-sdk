package library

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LibraryId{}

// LibraryId is a struct representing the Resource ID for a Library
type LibraryId struct {
	BaseURI     string
	LibraryName string
}

// NewLibraryID returns a new LibraryId struct
func NewLibraryID(baseURI string, libraryName string) LibraryId {
	return LibraryId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		LibraryName: libraryName,
	}
}

// ParseLibraryID parses 'input' into a LibraryId
func ParseLibraryID(input string) (*LibraryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LibraryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LibraryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLibraryIDInsensitively parses 'input' case-insensitively into a LibraryId
// note: this method should only be used for API response data and not user input
func ParseLibraryIDInsensitively(input string) (*LibraryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LibraryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LibraryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LibraryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.LibraryName, ok = input.Parsed["libraryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "libraryName", input)
	}

	return nil
}

// ValidateLibraryID checks that 'input' can be parsed as a Library ID
func ValidateLibraryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLibraryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Library ID
func (id LibraryId) ID() string {
	fmtString := "%s/libraries/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.LibraryName)
}

// Path returns the formatted Library ID without the BaseURI
func (id LibraryId) Path() string {
	fmtString := "/libraries/%s"
	return fmt.Sprintf(fmtString, id.LibraryName)
}

// PathElements returns the values of Library ID Segments without the BaseURI
func (id LibraryId) PathElements() []any {
	return []any{id.LibraryName}
}

// Segments returns a slice of Resource ID Segments which comprise this Library ID
func (id LibraryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticLibraries", "libraries", "libraries"),
		resourceids.UserSpecifiedSegment("libraryName", "libraryName"),
	}
}

// String returns a human-readable description of this Library ID
func (id LibraryId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Library Name: %q", id.LibraryName),
	}
	return fmt.Sprintf("Library (%s)", strings.Join(components, "\n"))
}
