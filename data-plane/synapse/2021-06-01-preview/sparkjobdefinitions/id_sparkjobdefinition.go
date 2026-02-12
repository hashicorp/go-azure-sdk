package sparkjobdefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SparkJobDefinitionId{}

// SparkJobDefinitionId is a struct representing the Resource ID for a Spark Job Definition
type SparkJobDefinitionId struct {
	BaseURI                string
	SparkJobDefinitionName string
}

// NewSparkJobDefinitionID returns a new SparkJobDefinitionId struct
func NewSparkJobDefinitionID(baseURI string, sparkJobDefinitionName string) SparkJobDefinitionId {
	return SparkJobDefinitionId{
		BaseURI:                strings.TrimSuffix(baseURI, "/"),
		SparkJobDefinitionName: sparkJobDefinitionName,
	}
}

// ParseSparkJobDefinitionID parses 'input' into a SparkJobDefinitionId
func ParseSparkJobDefinitionID(input string) (*SparkJobDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SparkJobDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SparkJobDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSparkJobDefinitionIDInsensitively parses 'input' case-insensitively into a SparkJobDefinitionId
// note: this method should only be used for API response data and not user input
func ParseSparkJobDefinitionIDInsensitively(input string) (*SparkJobDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SparkJobDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SparkJobDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SparkJobDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.SparkJobDefinitionName, ok = input.Parsed["sparkJobDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sparkJobDefinitionName", input)
	}

	return nil
}

// ValidateSparkJobDefinitionID checks that 'input' can be parsed as a Spark Job Definition ID
func ValidateSparkJobDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSparkJobDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Spark Job Definition ID
func (id SparkJobDefinitionId) ID() string {
	fmtString := "%s/sparkJobDefinitions/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.SparkJobDefinitionName)
}

// Path returns the formatted Spark Job Definition ID without the BaseURI
func (id SparkJobDefinitionId) Path() string {
	fmtString := "/sparkJobDefinitions/%s"
	return fmt.Sprintf(fmtString, id.SparkJobDefinitionName)
}

// PathElements returns the values of Spark Job Definition ID Segments without the BaseURI
func (id SparkJobDefinitionId) PathElements() []any {
	return []any{id.SparkJobDefinitionName}
}

// Segments returns a slice of Resource ID Segments which comprise this Spark Job Definition ID
func (id SparkJobDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticSparkJobDefinitions", "sparkJobDefinitions", "sparkJobDefinitions"),
		resourceids.UserSpecifiedSegment("sparkJobDefinitionName", "sparkJobDefinitionName"),
	}
}

// String returns a human-readable description of this Spark Job Definition ID
func (id SparkJobDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Spark Job Definition Name: %q", id.SparkJobDefinitionName),
	}
	return fmt.Sprintf("Spark Job Definition (%s)", strings.Join(components, "\n"))
}
