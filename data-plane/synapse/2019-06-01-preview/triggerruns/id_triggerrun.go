package triggerruns

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TriggerRunId{}

// TriggerRunId is a struct representing the Resource ID for a Trigger Run
type TriggerRunId struct {
	BaseURI     string
	TriggerName string
	RunId       string
}

// NewTriggerRunID returns a new TriggerRunId struct
func NewTriggerRunID(baseURI string, triggerName string, runId string) TriggerRunId {
	return TriggerRunId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		TriggerName: triggerName,
		RunId:       runId,
	}
}

// ParseTriggerRunID parses 'input' into a TriggerRunId
func ParseTriggerRunID(input string) (*TriggerRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TriggerRunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggerRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTriggerRunIDInsensitively parses 'input' case-insensitively into a TriggerRunId
// note: this method should only be used for API response data and not user input
func ParseTriggerRunIDInsensitively(input string) (*TriggerRunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TriggerRunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TriggerRunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TriggerRunId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.TriggerName, ok = input.Parsed["triggerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "triggerName", input)
	}

	if id.RunId, ok = input.Parsed["runId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "runId", input)
	}

	return nil
}

// ValidateTriggerRunID checks that 'input' can be parsed as a Trigger Run ID
func ValidateTriggerRunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTriggerRunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Trigger Run ID
func (id TriggerRunId) ID() string {
	fmtString := "%s/triggers/%s/triggerRuns/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.TriggerName, id.RunId)
}

// Path returns the formatted Trigger Run ID without the BaseURI
func (id TriggerRunId) Path() string {
	fmtString := "/triggers/%s/triggerRuns/%s"
	return fmt.Sprintf(fmtString, id.TriggerName, id.RunId)
}

// PathElements returns the values of Trigger Run ID Segments without the BaseURI
func (id TriggerRunId) PathElements() []any {
	return []any{id.TriggerName, id.RunId}
}

// Segments returns a slice of Resource ID Segments which comprise this Trigger Run ID
func (id TriggerRunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticTriggers", "triggers", "triggers"),
		resourceids.UserSpecifiedSegment("triggerName", "triggerName"),
		resourceids.StaticSegment("staticTriggerRuns", "triggerRuns", "triggerRuns"),
		resourceids.UserSpecifiedSegment("runId", "runId"),
	}
}

// String returns a human-readable description of this Trigger Run ID
func (id TriggerRunId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Trigger Name: %q", id.TriggerName),
		fmt.Sprintf("Run: %q", id.RunId),
	}
	return fmt.Sprintf("Trigger Run (%s)", strings.Join(components, "\n"))
}
