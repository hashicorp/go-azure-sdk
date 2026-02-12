package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RerunId{}

// RerunId is a struct representing the Resource ID for a Rerun
type RerunId struct {
	BaseURI string
	JobId   string
	RerunId string
}

// NewRerunID returns a new RerunId struct
func NewRerunID(baseURI string, jobId string, rerunId string) RerunId {
	return RerunId{
		BaseURI: strings.TrimSuffix(baseURI, "/"),
		JobId:   jobId,
		RerunId: rerunId,
	}
}

// ParseRerunID parses 'input' into a RerunId
func ParseRerunID(input string) (*RerunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RerunId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RerunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRerunIDInsensitively parses 'input' case-insensitively into a RerunId
// note: this method should only be used for API response data and not user input
func ParseRerunIDInsensitively(input string) (*RerunId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RerunId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RerunId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RerunId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.JobId, ok = input.Parsed["jobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "jobId", input)
	}

	if id.RerunId, ok = input.Parsed["rerunId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "rerunId", input)
	}

	return nil
}

// ValidateRerunID checks that 'input' can be parsed as a Rerun ID
func ValidateRerunID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRerunID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rerun ID
func (id RerunId) ID() string {
	fmtString := "%s/jobs/%s/rerun/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.JobId, id.RerunId)
}

// Path returns the formatted Rerun ID without the BaseURI
func (id RerunId) Path() string {
	fmtString := "/jobs/%s/rerun/%s"
	return fmt.Sprintf(fmtString, id.JobId, id.RerunId)
}

// PathElements returns the values of Rerun ID Segments without the BaseURI
func (id RerunId) PathElements() []any {
	return []any{id.JobId, id.RerunId}
}

// Segments returns a slice of Resource ID Segments which comprise this Rerun ID
func (id RerunId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("jobId", "jobId"),
		resourceids.StaticSegment("staticRerun", "rerun", "rerun"),
		resourceids.UserSpecifiedSegment("rerunId", "rerunId"),
	}
}

// String returns a human-readable description of this Rerun ID
func (id RerunId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Job: %q", id.JobId),
		fmt.Sprintf("Rerun: %q", id.RerunId),
	}
	return fmt.Sprintf("Rerun (%s)", strings.Join(components, "\n"))
}
