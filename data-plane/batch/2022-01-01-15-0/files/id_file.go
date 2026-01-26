package files

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &FileId{}

// FileId is a struct representing the Resource ID for a File
type FileId struct {
	BaseURI  string
	JobId    string
	TaskId   string
	FileName string
}

// NewFileID returns a new FileId struct
func NewFileID(baseURI string, jobId string, taskId string, fileName string) FileId {
	return FileId{
		BaseURI:  strings.TrimSuffix(baseURI, "/"),
		JobId:    jobId,
		TaskId:   taskId,
		FileName: fileName,
	}
}

// ParseFileID parses 'input' into a FileId
func ParseFileID(input string) (*FileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FileId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFileIDInsensitively parses 'input' case-insensitively into a FileId
// note: this method should only be used for API response data and not user input
func ParseFileIDInsensitively(input string) (*FileId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FileId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FileId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FileId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.FileName, ok = input.Parsed["fileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "fileName", input)
	}

	return nil
}

// ValidateFileID checks that 'input' can be parsed as a File ID
func ValidateFileID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFileID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted File ID
func (id FileId) ID() string {
	fmtString := "%s/jobs/%s/tasks/%s/files/%s"
	return fmt.Sprintf(fmtString, id.BaseURI, id.JobId, id.TaskId, id.FileName)
}

// Path returns the formatted File ID without the BaseURI
func (id FileId) Path() string {
	fmtString := "/jobs/%s/tasks/%s/files/%s"
	return fmt.Sprintf(fmtString, id.JobId, id.TaskId, id.FileName)
}

// Segments returns a slice of Resource ID Segments which comprise this File ID
func (id FileId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint_url"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("jobId", "jobId"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("taskId", "taskId"),
		resourceids.StaticSegment("staticFiles", "files", "files"),
		resourceids.UserSpecifiedSegment("fileName", "fileName"),
	}
}

// String returns a human-readable description of this File ID
func (id FileId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Job: %q", id.JobId),
		fmt.Sprintf("Task: %q", id.TaskId),
		fmt.Sprintf("File Name: %q", id.FileName),
	}
	return fmt.Sprintf("File (%s)", strings.Join(components, "\n"))
}
