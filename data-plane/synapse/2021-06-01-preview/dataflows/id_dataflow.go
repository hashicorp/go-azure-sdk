package dataflows

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DataflowId{}

// DataflowId is a struct representing the Resource ID for a Dataflow
type DataflowId struct {
	BaseURI      string
	DataflowName string
}

// NewDataflowID returns a new DataflowId struct
func NewDataflowID(baseURI string, dataflowName string) DataflowId {
	return DataflowId{
		BaseURI:      strings.TrimSuffix(baseURI, "/"),
		DataflowName: dataflowName,
	}
}

// ParseDataflowID parses 'input' into a DataflowId
func ParseDataflowID(input string) (*DataflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataflowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataflowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDataflowIDInsensitively parses 'input' case-insensitively into a DataflowId
// note: this method should only be used for API response data and not user input
func ParseDataflowIDInsensitively(input string) (*DataflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataflowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataflowId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DataflowId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DataflowName, ok = input.Parsed["dataflowName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataflowName", input)
	}

	return nil
}

// ValidateDataflowID checks that 'input' can be parsed as a Dataflow ID
func ValidateDataflowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataflowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dataflow ID
func (id DataflowId) ID() string {
	fmtString := "%s/dataflows/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DataflowName)
}

// Path returns the formatted Dataflow ID without the BaseURI
func (id DataflowId) Path() string {
	fmtString := "/dataflows/%s"
	return fmt.Sprintf(fmtString, id.DataflowName)
}

// PathElements returns the values of Dataflow ID Segments without the BaseURI
func (id DataflowId) PathElements() []any {
	return []any{id.DataflowName}
}

// Segments returns a slice of Resource ID Segments which comprise this Dataflow ID
func (id DataflowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDataflows", "dataflows", "dataflows"),
		resourceids.UserSpecifiedSegment("dataflowName", "dataflowName"),
	}
}

// String returns a human-readable description of this Dataflow ID
func (id DataflowId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Dataflow Name: %q", id.DataflowName),
	}
	return fmt.Sprintf("Dataflow (%s)", strings.Join(components, "\n"))
}
