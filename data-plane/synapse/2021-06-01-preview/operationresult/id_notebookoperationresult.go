package operationresult

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &NotebookOperationResultId{}

// NotebookOperationResultId is a struct representing the Resource ID for a Notebook Operation Result
type NotebookOperationResultId struct {
	BaseURI     string
	OperationId string
}

// NewNotebookOperationResultID returns a new NotebookOperationResultId struct
func NewNotebookOperationResultID(baseURI string, operationId string) NotebookOperationResultId {
	return NotebookOperationResultId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		OperationId: operationId,
	}
}

// ParseNotebookOperationResultID parses 'input' into a NotebookOperationResultId
func ParseNotebookOperationResultID(input string) (*NotebookOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NotebookOperationResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NotebookOperationResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseNotebookOperationResultIDInsensitively parses 'input' case-insensitively into a NotebookOperationResultId
// note: this method should only be used for API response data and not user input
func ParseNotebookOperationResultIDInsensitively(input string) (*NotebookOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NotebookOperationResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NotebookOperationResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *NotebookOperationResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
}

// ValidateNotebookOperationResultID checks that 'input' can be parsed as a Notebook Operation Result ID
func ValidateNotebookOperationResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNotebookOperationResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Notebook Operation Result ID
func (id NotebookOperationResultId) ID() string {
	fmtString := "%s/notebookOperationResults/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.OperationId)
}

// Path returns the formatted Notebook Operation Result ID without the BaseURI
func (id NotebookOperationResultId) Path() string {
	fmtString := "/notebookOperationResults/%s"
	return fmt.Sprintf(fmtString, id.OperationId)
}

// PathElements returns the values of Notebook Operation Result ID Segments without the BaseURI
func (id NotebookOperationResultId) PathElements() []any {
	return []any{id.OperationId}
}

// Segments returns a slice of Resource ID Segments which comprise this Notebook Operation Result ID
func (id NotebookOperationResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticNotebookOperationResults", "notebookOperationResults", "notebookOperationResults"),
		resourceids.UserSpecifiedSegment("operationId", "operationId"),
	}
}

// String returns a human-readable description of this Notebook Operation Result ID
func (id NotebookOperationResultId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Notebook Operation Result (%s)", strings.Join(components, "\n"))
}
