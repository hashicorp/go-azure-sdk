package applicationsynchronizationjobschemadirectory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationSynchronizationJobSchemaDirectoryId{}

// ApplicationSynchronizationJobSchemaDirectoryId is a struct representing the Resource ID for a Application Synchronization Job Schema Directory
type ApplicationSynchronizationJobSchemaDirectoryId struct {
	ApplicationId         string
	SynchronizationJobId  string
	DirectoryDefinitionId string
}

// NewApplicationSynchronizationJobSchemaDirectoryID returns a new ApplicationSynchronizationJobSchemaDirectoryId struct
func NewApplicationSynchronizationJobSchemaDirectoryID(applicationId string, synchronizationJobId string, directoryDefinitionId string) ApplicationSynchronizationJobSchemaDirectoryId {
	return ApplicationSynchronizationJobSchemaDirectoryId{
		ApplicationId:         applicationId,
		SynchronizationJobId:  synchronizationJobId,
		DirectoryDefinitionId: directoryDefinitionId,
	}
}

// ParseApplicationSynchronizationJobSchemaDirectoryID parses 'input' into a ApplicationSynchronizationJobSchemaDirectoryId
func ParseApplicationSynchronizationJobSchemaDirectoryID(input string) (*ApplicationSynchronizationJobSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationSynchronizationJobSchemaDirectoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationSynchronizationJobSchemaDirectoryId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.SynchronizationJobId, ok = parsed.Parsed["synchronizationJobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", *parsed)
	}

	if id.DirectoryDefinitionId, ok = parsed.Parsed["directoryDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseApplicationSynchronizationJobSchemaDirectoryIDInsensitively parses 'input' case-insensitively into a ApplicationSynchronizationJobSchemaDirectoryId
// note: this method should only be used for API response data and not user input
func ParseApplicationSynchronizationJobSchemaDirectoryIDInsensitively(input string) (*ApplicationSynchronizationJobSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationSynchronizationJobSchemaDirectoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationSynchronizationJobSchemaDirectoryId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.SynchronizationJobId, ok = parsed.Parsed["synchronizationJobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", *parsed)
	}

	if id.DirectoryDefinitionId, ok = parsed.Parsed["directoryDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationSynchronizationJobSchemaDirectoryID checks that 'input' can be parsed as a Application Synchronization Job Schema Directory ID
func ValidateApplicationSynchronizationJobSchemaDirectoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationSynchronizationJobSchemaDirectoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Synchronization Job Schema Directory ID
func (id ApplicationSynchronizationJobSchemaDirectoryId) ID() string {
	fmtString := "/applications/%s/synchronization/jobs/%s/schema/directories/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.SynchronizationJobId, id.DirectoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Synchronization Job Schema Directory ID
func (id ApplicationSynchronizationJobSchemaDirectoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticSynchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("synchronizationJobId", "synchronizationJobIdValue"),
		resourceids.StaticSegment("staticSchema", "schema", "schema"),
		resourceids.StaticSegment("staticDirectories", "directories", "directories"),
		resourceids.UserSpecifiedSegment("directoryDefinitionId", "directoryDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Application Synchronization Job Schema Directory ID
func (id ApplicationSynchronizationJobSchemaDirectoryId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Synchronization Job: %q", id.SynchronizationJobId),
		fmt.Sprintf("Directory Definition: %q", id.DirectoryDefinitionId),
	}
	return fmt.Sprintf("Application Synchronization Job Schema Directory (%s)", strings.Join(components, "\n"))
}
