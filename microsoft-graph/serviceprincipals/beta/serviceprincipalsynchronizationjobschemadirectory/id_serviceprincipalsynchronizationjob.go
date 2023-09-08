package serviceprincipalsynchronizationjobschemadirectory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalSynchronizationJobId{}

// ServicePrincipalSynchronizationJobId is a struct representing the Resource ID for a Service Principal Synchronization Job
type ServicePrincipalSynchronizationJobId struct {
	ServicePrincipalId   string
	SynchronizationJobId string
}

// NewServicePrincipalSynchronizationJobID returns a new ServicePrincipalSynchronizationJobId struct
func NewServicePrincipalSynchronizationJobID(servicePrincipalId string, synchronizationJobId string) ServicePrincipalSynchronizationJobId {
	return ServicePrincipalSynchronizationJobId{
		ServicePrincipalId:   servicePrincipalId,
		SynchronizationJobId: synchronizationJobId,
	}
}

// ParseServicePrincipalSynchronizationJobID parses 'input' into a ServicePrincipalSynchronizationJobId
func ParseServicePrincipalSynchronizationJobID(input string) (*ServicePrincipalSynchronizationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalSynchronizationJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalSynchronizationJobId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.SynchronizationJobId, ok = parsed.Parsed["synchronizationJobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalSynchronizationJobIDInsensitively parses 'input' case-insensitively into a ServicePrincipalSynchronizationJobId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalSynchronizationJobIDInsensitively(input string) (*ServicePrincipalSynchronizationJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalSynchronizationJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalSynchronizationJobId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.SynchronizationJobId, ok = parsed.Parsed["synchronizationJobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalSynchronizationJobID checks that 'input' can be parsed as a Service Principal Synchronization Job ID
func ValidateServicePrincipalSynchronizationJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalSynchronizationJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Synchronization Job ID
func (id ServicePrincipalSynchronizationJobId) ID() string {
	fmtString := "/servicePrincipals/%s/synchronization/jobs/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.SynchronizationJobId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Synchronization Job ID
func (id ServicePrincipalSynchronizationJobId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticSynchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("synchronizationJobId", "synchronizationJobIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Synchronization Job ID
func (id ServicePrincipalSynchronizationJobId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Synchronization Job: %q", id.SynchronizationJobId),
	}
	return fmt.Sprintf("Service Principal Synchronization Job (%s)", strings.Join(components, "\n"))
}
