package activityruns

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PipelinePipelineRunId{}

// PipelinePipelineRunId is a struct representing the Resource ID for a Pipeline Pipeline Run
type PipelinePipelineRunId struct {
	BaseURI      string
	PipelineName string
	RunId        string
}

// NewPipelinePipelineRunID returns a new PipelinePipelineRunId struct
func NewPipelinePipelineRunID(baseURI string, pipelineName string, runId string) PipelinePipelineRunId {
	return PipelinePipelineRunId{
		BaseURI:      strings.TrimSuffix(baseURI, "/"),
		PipelineName: pipelineName,
		RunId:        runId,
	}
}

// ParsePipelinePipelineRunID parses 'input' into a PipelinePipelineRunId
func ParsePipelinePipelineRunID(input string) (*PipelinePipelineRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PipelinePipelineRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PipelinePipelineRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePipelinePipelineRunIDInsensitively parses 'input' case-insensitively into a PipelinePipelineRunId
// note: this method should only be used for API response data and not user input
func ParsePipelinePipelineRunIDInsensitively(input string) (*PipelinePipelineRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PipelinePipelineRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PipelinePipelineRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PipelinePipelineRunId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.PipelineName, ok = input.Parsed["pipelineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pipelineName", input)
	}

	if id.RunId, ok = input.Parsed["runId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runId", input)
	}

	return nil
}

// ValidatePipelinePipelineRunID checks that 'input' can be parsed as a Pipeline Pipeline Run ID
func ValidatePipelinePipelineRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePipelinePipelineRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pipeline Pipeline Run ID
func (id PipelinePipelineRunId) ID() string {
	fmtString := "%s/pipelines/%s/pipelineRuns/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.PipelineName, id.RunId)
}

// Path returns the formatted Pipeline Pipeline Run ID without the BaseURI
func (id PipelinePipelineRunId) Path() string {
	fmtString := "/pipelines/%s/pipelineRuns/%s"
	return fmt.Sprintf(fmtString, id.PipelineName, id.RunId)
}

// PathElements returns the values of Pipeline Pipeline Run ID Segments without the BaseURI
func (id PipelinePipelineRunId) PathElements() []any {
	return []any{id.PipelineName, id.RunId}
}

// Segments returns a slice of Resource ID Segments which comprise this Pipeline Pipeline Run ID
func (id PipelinePipelineRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticPipelines", "pipelines", "pipelines"),
		resourceids.UserSpecifiedSegment("pipelineName", "pipelineName"),
		resourceids.StaticSegment("staticPipelineRuns", "pipelineRuns", "pipelineRuns"),
		resourceids.UserSpecifiedSegment("runId", "runId"),
	}
}

// String returns a human-readable description of this Pipeline Pipeline Run ID
func (id PipelinePipelineRunId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Pipeline Name: %q", id.PipelineName),
		fmt.Sprintf("Run: %q", id.RunId),
	}
	return fmt.Sprintf("Pipeline Pipeline Run (%s)", strings.Join(components, "\n"))
}
