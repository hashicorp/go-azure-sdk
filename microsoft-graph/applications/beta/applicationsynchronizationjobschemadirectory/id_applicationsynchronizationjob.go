package applicationsynchronizationjobschemadirectory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationSynchronizationJobId{}

// ApplicationSynchronizationJobId is a struct representing the Resource ID for a Application Synchronization Job
type ApplicationSynchronizationJobId struct {
	ApplicationId        string
	SynchronizationJobId string
}

// NewApplicationSynchronizationJobID returns a new ApplicationSynchronizationJobId struct
func NewApplicationSynchronizationJobID(applicationId string, synchronizationJobId string) ApplicationSynchronizationJobId {
	return ApplicationSynchronizationJobId{
		ApplicationId:        applicationId,
		SynchronizationJobId: synchronizationJobId,
	}
}

// ParseApplicationSynchronizationJobID parses 'input' into a ApplicationSynchronizationJobId
func ParseApplicationSynchronizationJobID(input string) (*ApplicationSynchronizationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationSynchronizationJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationSynchronizationJobId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.SynchronizationJobId, ok = parsed.Parsed["synchronizationJobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", *parsed)
	}

	return &id, nil
}

// ParseApplicationSynchronizationJobIDInsensitively parses 'input' case-insensitively into a ApplicationSynchronizationJobId
// note: this method should only be used for API response data and not user input
func ParseApplicationSynchronizationJobIDInsensitively(input string) (*ApplicationSynchronizationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationSynchronizationJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationSynchronizationJobId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.SynchronizationJobId, ok = parsed.Parsed["synchronizationJobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationSynchronizationJobID checks that 'input' can be parsed as a Application Synchronization Job ID
func ValidateApplicationSynchronizationJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationSynchronizationJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Synchronization Job ID
func (id ApplicationSynchronizationJobId) ID() string {
	fmtString := "/applications/%s/synchronization/jobs/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.SynchronizationJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Synchronization Job ID
func (id ApplicationSynchronizationJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticSynchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("synchronizationJobId", "synchronizationJobIdValue"),
	}
}

// String returns a human-readable description of this Application Synchronization Job ID
func (id ApplicationSynchronizationJobId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Synchronization Job: %q", id.SynchronizationJobId),
	}
	return fmt.Sprintf("Application Synchronization Job (%s)", strings.Join(components, "\n"))
}
