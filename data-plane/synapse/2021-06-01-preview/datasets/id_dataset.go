package datasets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DatasetId{}

// DatasetId is a struct representing the Resource ID for a Dataset
type DatasetId struct {
	BaseURI     string
	DatasetName string
}

// NewDatasetID returns a new DatasetId struct
func NewDatasetID(baseURI string, datasetName string) DatasetId {
	return DatasetId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		DatasetName: datasetName,
	}
}

// ParseDatasetID parses 'input' into a DatasetId
func ParseDatasetID(input string) (*DatasetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DatasetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DatasetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDatasetIDInsensitively parses 'input' case-insensitively into a DatasetId
// note: this method should only be used for API response data and not user input
func ParseDatasetIDInsensitively(input string) (*DatasetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DatasetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DatasetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DatasetId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DatasetName, ok = input.Parsed["datasetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "datasetName", input)
	}

	return nil
}

// ValidateDatasetID checks that 'input' can be parsed as a Dataset ID
func ValidateDatasetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDatasetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dataset ID
func (id DatasetId) ID() string {
	fmtString := "%s/datasets/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DatasetName)
}

// Path returns the formatted Dataset ID without the BaseURI
func (id DatasetId) Path() string {
	fmtString := "/datasets/%s"
	return fmt.Sprintf(fmtString, id.DatasetName)
}

// PathElements returns the values of Dataset ID Segments without the BaseURI
func (id DatasetId) PathElements() []any {
	return []any{id.DatasetName}
}

// Segments returns a slice of Resource ID Segments which comprise this Dataset ID
func (id DatasetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDatasets", "datasets", "datasets"),
		resourceids.UserSpecifiedSegment("datasetName", "datasetName"),
	}
}

// String returns a human-readable description of this Dataset ID
func (id DatasetId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Dataset Name: %q", id.DatasetName),
	}
	return fmt.Sprintf("Dataset (%s)", strings.Join(components, "\n"))
}
