package pipelines

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PipelineId{}

// PipelineId is a struct representing the Resource ID for a Pipeline
type PipelineId struct {
	BaseURI      string
	PipelineName string
}

// NewPipelineID returns a new PipelineId struct
func NewPipelineID(baseURI string, pipelineName string) PipelineId {
	return PipelineId{
		BaseURI:      strings.TrimSuffix(baseURI, "/"),
		PipelineName: pipelineName,
	}
}

// ParsePipelineID parses 'input' into a PipelineId
func ParsePipelineID(input string) (*PipelineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PipelineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PipelineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePipelineIDInsensitively parses 'input' case-insensitively into a PipelineId
// note: this method should only be used for API response data and not user input
func ParsePipelineIDInsensitively(input string) (*PipelineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PipelineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PipelineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PipelineId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.PipelineName, ok = input.Parsed["pipelineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pipelineName", input)
	}

	return nil
}

// ValidatePipelineID checks that 'input' can be parsed as a Pipeline ID
func ValidatePipelineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePipelineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pipeline ID
func (id PipelineId) ID() string {
	fmtString := "%s/pipelines/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.PipelineName)
}

// Path returns the formatted Pipeline ID without the BaseURI
func (id PipelineId) Path() string {
	fmtString := "/pipelines/%s"
	return fmt.Sprintf(fmtString, id.PipelineName)
}

// PathElements returns the values of Pipeline ID Segments without the BaseURI
func (id PipelineId) PathElements() []any {
	return []any{id.PipelineName}
}

// Segments returns a slice of Resource ID Segments which comprise this Pipeline ID
func (id PipelineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticPipelines", "pipelines", "pipelines"),
		resourceids.UserSpecifiedSegment("pipelineName", "pipelineName"),
	}
}

// String returns a human-readable description of this Pipeline ID
func (id PipelineId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Pipeline Name: %q", id.PipelineName),
	}
	return fmt.Sprintf("Pipeline (%s)", strings.Join(components, "\n"))
}
