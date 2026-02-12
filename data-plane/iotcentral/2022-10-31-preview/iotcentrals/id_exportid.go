package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ExportIdId{}

// ExportIdId is a struct representing the Resource ID for a Export Id
type ExportIdId struct {
	BaseURI  string
	ExportId string
}

// NewExportIdID returns a new ExportIdId struct
func NewExportIdID(baseURI string, exportId string) ExportIdId {
	return ExportIdId{
		BaseURI:  strings.TrimSuffix(baseURI, "/"),
		ExportId: exportId,
	}
}

// ParseExportIdID parses 'input' into a ExportIdId
func ParseExportIdID(input string) (*ExportIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExportIdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExportIdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseExportIdIDInsensitively parses 'input' case-insensitively into a ExportIdId
// note: this method should only be used for API response data and not user input
func ParseExportIdIDInsensitively(input string) (*ExportIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ExportIdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ExportIdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ExportIdId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.ExportId, ok = input.Parsed["exportId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "exportId", input)
	}

	return nil
}

// ValidateExportIdID checks that 'input' can be parsed as a Export Id ID
func ValidateExportIdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExportIdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Export Id ID
func (id ExportIdId) ID() string {
	fmtString := "%s/dataExport/exports/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.ExportId)
}

// Path returns the formatted Export Id ID without the BaseURI
func (id ExportIdId) Path() string {
	fmtString := "/dataExport/exports/%s"
	return fmt.Sprintf(fmtString, id.ExportId)
}

// PathElements returns the values of Export Id ID Segments without the BaseURI
func (id ExportIdId) PathElements() []any {
	return []any{id.ExportId}
}

// Segments returns a slice of Resource ID Segments which comprise this Export Id ID
func (id ExportIdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDataExport", "dataExport", "dataExport"),
		resourceids.StaticSegment("staticExports", "exports", "exports"),
		resourceids.UserSpecifiedSegment("exportId", "exportId"),
	}
}

// String returns a human-readable description of this Export Id ID
func (id ExportIdId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Export: %q", id.ExportId),
	}
	return fmt.Sprintf("Export Id (%s)", strings.Join(components, "\n"))
}
