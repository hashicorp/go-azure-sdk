package library

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &LibraryOperationResultId{}

// LibraryOperationResultId is a struct representing the Resource ID for a Library Operation Result
type LibraryOperationResultId struct {
	BaseURI     string
	OperationId string
}

// NewLibraryOperationResultID returns a new LibraryOperationResultId struct
func NewLibraryOperationResultID(baseURI string, operationId string) LibraryOperationResultId {
	return LibraryOperationResultId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		OperationId: operationId,
	}
}

// ParseLibraryOperationResultID parses 'input' into a LibraryOperationResultId
func ParseLibraryOperationResultID(input string) (*LibraryOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LibraryOperationResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LibraryOperationResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLibraryOperationResultIDInsensitively parses 'input' case-insensitively into a LibraryOperationResultId
// note: this method should only be used for API response data and not user input
func ParseLibraryOperationResultIDInsensitively(input string) (*LibraryOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LibraryOperationResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LibraryOperationResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LibraryOperationResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
}

// ValidateLibraryOperationResultID checks that 'input' can be parsed as a Library Operation Result ID
func ValidateLibraryOperationResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLibraryOperationResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Library Operation Result ID
func (id LibraryOperationResultId) ID() string {
	fmtString := "%s/libraryOperationResults/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.OperationId)
}

// Path returns the formatted Library Operation Result ID without the BaseURI
func (id LibraryOperationResultId) Path() string {
	fmtString := "/libraryOperationResults/%s"
	return fmt.Sprintf(fmtString, id.OperationId)
}

// PathElements returns the values of Library Operation Result ID Segments without the BaseURI
func (id LibraryOperationResultId) PathElements() []any {
	return []any{id.OperationId}
}

// Segments returns a slice of Resource ID Segments which comprise this Library Operation Result ID
func (id LibraryOperationResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticLibraryOperationResults", "libraryOperationResults", "libraryOperationResults"),
		resourceids.UserSpecifiedSegment("operationId", "operationId"),
	}
}

// String returns a human-readable description of this Library Operation Result ID
func (id LibraryOperationResultId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Library Operation Result (%s)", strings.Join(components, "\n"))
}
