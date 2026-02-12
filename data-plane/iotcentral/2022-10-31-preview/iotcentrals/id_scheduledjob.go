package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScheduledJobId{}

// ScheduledJobId is a struct representing the Resource ID for a Scheduled Job
type ScheduledJobId struct {
	BaseURI        string
	ScheduledJobId string
}

// NewScheduledJobID returns a new ScheduledJobId struct
func NewScheduledJobID(baseURI string, scheduledJobId string) ScheduledJobId {
	return ScheduledJobId{
		BaseURI:        strings.TrimSuffix(baseURI, "/"),
		ScheduledJobId: scheduledJobId,
	}
}

// ParseScheduledJobID parses 'input' into a ScheduledJobId
func ParseScheduledJobID(input string) (*ScheduledJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScheduledJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScheduledJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScheduledJobIDInsensitively parses 'input' case-insensitively into a ScheduledJobId
// note: this method should only be used for API response data and not user input
func ParseScheduledJobIDInsensitively(input string) (*ScheduledJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScheduledJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScheduledJobId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScheduledJobId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.ScheduledJobId, ok = input.Parsed["scheduledJobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scheduledJobId", input)
	}

	return nil
}

// ValidateScheduledJobID checks that 'input' can be parsed as a Scheduled Job ID
func ValidateScheduledJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScheduledJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scheduled Job ID
func (id ScheduledJobId) ID() string {
	fmtString := "%s/scheduledJobs/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.ScheduledJobId)
}

// Path returns the formatted Scheduled Job ID without the BaseURI
func (id ScheduledJobId) Path() string {
	fmtString := "/scheduledJobs/%s"
	return fmt.Sprintf(fmtString, id.ScheduledJobId)
}

// PathElements returns the values of Scheduled Job ID Segments without the BaseURI
func (id ScheduledJobId) PathElements() []any {
	return []any{id.ScheduledJobId}
}

// Segments returns a slice of Resource ID Segments which comprise this Scheduled Job ID
func (id ScheduledJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticScheduledJobs", "scheduledJobs", "scheduledJobs"),
		resourceids.UserSpecifiedSegment("scheduledJobId", "scheduledJobId"),
	}
}

// String returns a human-readable description of this Scheduled Job ID
func (id ScheduledJobId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Scheduled Job: %q", id.ScheduledJobId),
	}
	return fmt.Sprintf("Scheduled Job (%s)", strings.Join(components, "\n"))
}
