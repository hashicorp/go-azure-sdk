package sparkconfigurations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &SparkConfigurationId{}

// SparkConfigurationId is a struct representing the Resource ID for a Spark Configuration
type SparkConfigurationId struct {
	BaseURI                string
	SparkConfigurationName string
}

// NewSparkConfigurationID returns a new SparkConfigurationId struct
func NewSparkConfigurationID(baseURI string, sparkConfigurationName string) SparkConfigurationId {
	return SparkConfigurationId{
		BaseURI:                strings.TrimSuffix(baseURI, "/"),
		SparkConfigurationName: sparkConfigurationName,
	}
}

// ParseSparkConfigurationID parses 'input' into a SparkConfigurationId
func ParseSparkConfigurationID(input string) (*SparkConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SparkConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SparkConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSparkConfigurationIDInsensitively parses 'input' case-insensitively into a SparkConfigurationId
// note: this method should only be used for API response data and not user input
func ParseSparkConfigurationIDInsensitively(input string) (*SparkConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SparkConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SparkConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SparkConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.SparkConfigurationName, ok = input.Parsed["sparkConfigurationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sparkConfigurationName", input)
	}

	return nil
}

// ValidateSparkConfigurationID checks that 'input' can be parsed as a Spark Configuration ID
func ValidateSparkConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSparkConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Spark Configuration ID
func (id SparkConfigurationId) ID() string {
	fmtString := "%s/sparkConfigurations/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.SparkConfigurationName)
}

// Path returns the formatted Spark Configuration ID without the BaseURI
func (id SparkConfigurationId) Path() string {
	fmtString := "/sparkConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SparkConfigurationName)
}

// PathElements returns the values of Spark Configuration ID Segments without the BaseURI
func (id SparkConfigurationId) PathElements() []any {
	return []any{id.SparkConfigurationName}
}

// Segments returns a slice of Resource ID Segments which comprise this Spark Configuration ID
func (id SparkConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticSparkConfigurations", "sparkConfigurations", "sparkConfigurations"),
		resourceids.UserSpecifiedSegment("sparkConfigurationName", "sparkConfigurationName"),
	}
}

// String returns a human-readable description of this Spark Configuration ID
func (id SparkConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Spark Configuration Name: %q", id.SparkConfigurationName),
	}
	return fmt.Sprintf("Spark Configuration (%s)", strings.Join(components, "\n"))
}
