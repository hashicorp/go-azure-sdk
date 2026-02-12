package notebooks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NotebookId{}

// NotebookId is a struct representing the Resource ID for a Notebook
type NotebookId struct {
	BaseURI      string
	NotebookName string
}

// NewNotebookID returns a new NotebookId struct
func NewNotebookID(baseURI string, notebookName string) NotebookId {
	return NotebookId{
		BaseURI:      strings.TrimSuffix(baseURI, "/"),
		NotebookName: notebookName,
	}
}

// ParseNotebookID parses 'input' into a NotebookId
func ParseNotebookID(input string) (*NotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NotebookId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseNotebookIDInsensitively parses 'input' case-insensitively into a NotebookId
// note: this method should only be used for API response data and not user input
func ParseNotebookIDInsensitively(input string) (*NotebookId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NotebookId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NotebookId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *NotebookId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.NotebookName, ok = input.Parsed["notebookName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "notebookName", input)
	}

	return nil
}

// ValidateNotebookID checks that 'input' can be parsed as a Notebook ID
func ValidateNotebookID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNotebookID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Notebook ID
func (id NotebookId) ID() string {
	fmtString := "%s/notebooks/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.NotebookName)
}

// Path returns the formatted Notebook ID without the BaseURI
func (id NotebookId) Path() string {
	fmtString := "/notebooks/%s"
	return fmt.Sprintf(fmtString, id.NotebookName)
}

// PathElements returns the values of Notebook ID Segments without the BaseURI
func (id NotebookId) PathElements() []any {
	return []any{id.NotebookName}
}

// Segments returns a slice of Resource ID Segments which comprise this Notebook ID
func (id NotebookId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticNotebooks", "notebooks", "notebooks"),
		resourceids.UserSpecifiedSegment("notebookName", "notebookName"),
	}
}

// String returns a human-readable description of this Notebook ID
func (id NotebookId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Notebook Name: %q", id.NotebookName),
	}
	return fmt.Sprintf("Notebook (%s)", strings.Join(components, "\n"))
}
