package pipelineruns

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PipelineRunId{}

// PipelineRunId is a struct representing the Resource ID for a Pipeline Run
type PipelineRunId struct {
	BaseURI string
	RunId   string
}

// NewPipelineRunID returns a new PipelineRunId struct
func NewPipelineRunID(baseURI string, runId string) PipelineRunId {
	return PipelineRunId{
		BaseURI: strings.TrimSuffix(baseURI, "/"),
		RunId:   runId,
	}
}

// ParsePipelineRunID parses 'input' into a PipelineRunId
func ParsePipelineRunID(input string) (*PipelineRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PipelineRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PipelineRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePipelineRunIDInsensitively parses 'input' case-insensitively into a PipelineRunId
// note: this method should only be used for API response data and not user input
func ParsePipelineRunIDInsensitively(input string) (*PipelineRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PipelineRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PipelineRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PipelineRunId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.RunId, ok = input.Parsed["runId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runId", input)
	}

	return nil
}

// ValidatePipelineRunID checks that 'input' can be parsed as a Pipeline Run ID
func ValidatePipelineRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePipelineRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pipeline Run ID
func (id PipelineRunId) ID() string {
	fmtString := "%s/pipelineRuns/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.RunId)
}

// Path returns the formatted Pipeline Run ID without the BaseURI
func (id PipelineRunId) Path() string {
	fmtString := "/pipelineRuns/%s"
	return fmt.Sprintf(fmtString, id.RunId)
}

// PathElements returns the values of Pipeline Run ID Segments without the BaseURI
func (id PipelineRunId) PathElements() []any {
	return []any{id.RunId}
}

// Segments returns a slice of Resource ID Segments which comprise this Pipeline Run ID
func (id PipelineRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticPipelineRuns", "pipelineRuns", "pipelineRuns"),
		resourceids.UserSpecifiedSegment("runId", "runId"),
	}
}

// String returns a human-readable description of this Pipeline Run ID
func (id PipelineRunId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Run: %q", id.RunId),
	}
	return fmt.Sprintf("Pipeline Run (%s)", strings.Join(components, "\n"))
}
