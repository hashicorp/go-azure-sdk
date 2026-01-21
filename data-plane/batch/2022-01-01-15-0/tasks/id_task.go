package tasks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TaskId{}

// TaskId is a struct representing the Resource ID for a Task
type TaskId struct {
	BaseURI string
	JobId   string
	TaskId  string
}

// NewTaskID returns a new TaskId struct
func NewTaskID(baseURI string, jobId string, taskId string) TaskId {
	return TaskId{
		BaseURI: baseURI,
		JobId:   jobId,
		TaskId:  taskId,
	}
}

// ParseTaskID parses 'input' into a TaskId
func ParseTaskID(input string) (*TaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TaskId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTaskIDInsensitively parses 'input' case-insensitively into a TaskId
// note: this method should only be used for API response data and not user input
func ParseTaskIDInsensitively(input string) (*TaskId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TaskId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TaskId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TaskId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.JobId, ok = input.Parsed["jobId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "jobId", input)
	}

	if id.TaskId, ok = input.Parsed["taskId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "taskId", input)
	}

	return nil
}

// ValidateTaskID checks that 'input' can be parsed as a Task ID
func ValidateTaskID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTaskID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Task ID
func (id TaskId) ID() string {
	fmtString := "%s/jobs/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.BaseURI, id.JobId, id.TaskId)
}

// Path returns the formatted Task ID without the Scope / BaseURI
func (id TaskId) Path() string {
	fmtString := "/jobs/%s/tasks/%s"
	return fmt.Sprintf(fmtString, id.JobId, id.TaskId)
}

// Segments returns a slice of Resource ID Segments which comprise this Task ID
func (id TaskId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint_url"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("jobId", "jobId"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("taskId", "taskId"),
	}
}

// String returns a human-readable description of this Task ID
func (id TaskId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Job: %q", id.JobId),
		fmt.Sprintf("Task: %q", id.TaskId),
	}
	return fmt.Sprintf("Task (%s)", strings.Join(components, "\n"))
}
