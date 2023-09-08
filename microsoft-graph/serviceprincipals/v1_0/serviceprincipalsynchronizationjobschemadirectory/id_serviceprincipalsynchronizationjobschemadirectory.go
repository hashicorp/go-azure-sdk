package serviceprincipalsynchronizationjobschemadirectory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalSynchronizationJobSchemaDirectoryId{}

// ServicePrincipalSynchronizationJobSchemaDirectoryId is a struct representing the Resource ID for a Service Principal Synchronization Job Schema Directory
type ServicePrincipalSynchronizationJobSchemaDirectoryId struct {
	ServicePrincipalId    string
	SynchronizationJobId  string
	DirectoryDefinitionId string
}

// NewServicePrincipalSynchronizationJobSchemaDirectoryID returns a new ServicePrincipalSynchronizationJobSchemaDirectoryId struct
func NewServicePrincipalSynchronizationJobSchemaDirectoryID(servicePrincipalId string, synchronizationJobId string, directoryDefinitionId string) ServicePrincipalSynchronizationJobSchemaDirectoryId {
	return ServicePrincipalSynchronizationJobSchemaDirectoryId{
		ServicePrincipalId:    servicePrincipalId,
		SynchronizationJobId:  synchronizationJobId,
		DirectoryDefinitionId: directoryDefinitionId,
	}
}

// ParseServicePrincipalSynchronizationJobSchemaDirectoryID parses 'input' into a ServicePrincipalSynchronizationJobSchemaDirectoryId
func ParseServicePrincipalSynchronizationJobSchemaDirectoryID(input string) (*ServicePrincipalSynchronizationJobSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalSynchronizationJobSchemaDirectoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalSynchronizationJobSchemaDirectoryId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.SynchronizationJobId, ok = parsed.Parsed["synchronizationJobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", *parsed)
	}

	if id.DirectoryDefinitionId, ok = parsed.Parsed["directoryDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalSynchronizationJobSchemaDirectoryIDInsensitively parses 'input' case-insensitively into a ServicePrincipalSynchronizationJobSchemaDirectoryId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalSynchronizationJobSchemaDirectoryIDInsensitively(input string) (*ServicePrincipalSynchronizationJobSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalSynchronizationJobSchemaDirectoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalSynchronizationJobSchemaDirectoryId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.SynchronizationJobId, ok = parsed.Parsed["synchronizationJobId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationJobId", *parsed)
	}

	if id.DirectoryDefinitionId, ok = parsed.Parsed["directoryDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalSynchronizationJobSchemaDirectoryID checks that 'input' can be parsed as a Service Principal Synchronization Job Schema Directory ID
func ValidateServicePrincipalSynchronizationJobSchemaDirectoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalSynchronizationJobSchemaDirectoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Synchronization Job Schema Directory ID
func (id ServicePrincipalSynchronizationJobSchemaDirectoryId) ID() string {
	fmtString := "/servicePrincipals/%s/synchronization/jobs/%s/schema/directories/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.SynchronizationJobId, id.DirectoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Synchronization Job Schema Directory ID
func (id ServicePrincipalSynchronizationJobSchemaDirectoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticSynchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("staticJobs", "jobs", "jobs"),
		resourceids.UserSpecifiedSegment("synchronizationJobId", "synchronizationJobIdValue"),
		resourceids.StaticSegment("staticSchema", "schema", "schema"),
		resourceids.StaticSegment("staticDirectories", "directories", "directories"),
		resourceids.UserSpecifiedSegment("directoryDefinitionId", "directoryDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Synchronization Job Schema Directory ID
func (id ServicePrincipalSynchronizationJobSchemaDirectoryId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Synchronization Job: %q", id.SynchronizationJobId),
		fmt.Sprintf("Directory Definition: %q", id.DirectoryDefinitionId),
	}
	return fmt.Sprintf("Service Principal Synchronization Job Schema Directory (%s)", strings.Join(components, "\n"))
}
