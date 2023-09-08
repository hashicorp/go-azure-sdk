package serviceprincipalsynchronizationtemplateschemadirectory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalSynchronizationTemplateSchemaDirectoryId{}

// ServicePrincipalSynchronizationTemplateSchemaDirectoryId is a struct representing the Resource ID for a Service Principal Synchronization Template Schema Directory
type ServicePrincipalSynchronizationTemplateSchemaDirectoryId struct {
	ServicePrincipalId        string
	SynchronizationTemplateId string
	DirectoryDefinitionId     string
}

// NewServicePrincipalSynchronizationTemplateSchemaDirectoryID returns a new ServicePrincipalSynchronizationTemplateSchemaDirectoryId struct
func NewServicePrincipalSynchronizationTemplateSchemaDirectoryID(servicePrincipalId string, synchronizationTemplateId string, directoryDefinitionId string) ServicePrincipalSynchronizationTemplateSchemaDirectoryId {
	return ServicePrincipalSynchronizationTemplateSchemaDirectoryId{
		ServicePrincipalId:        servicePrincipalId,
		SynchronizationTemplateId: synchronizationTemplateId,
		DirectoryDefinitionId:     directoryDefinitionId,
	}
}

// ParseServicePrincipalSynchronizationTemplateSchemaDirectoryID parses 'input' into a ServicePrincipalSynchronizationTemplateSchemaDirectoryId
func ParseServicePrincipalSynchronizationTemplateSchemaDirectoryID(input string) (*ServicePrincipalSynchronizationTemplateSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalSynchronizationTemplateSchemaDirectoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalSynchronizationTemplateSchemaDirectoryId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.SynchronizationTemplateId, ok = parsed.Parsed["synchronizationTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", *parsed)
	}

	if id.DirectoryDefinitionId, ok = parsed.Parsed["directoryDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalSynchronizationTemplateSchemaDirectoryIDInsensitively parses 'input' case-insensitively into a ServicePrincipalSynchronizationTemplateSchemaDirectoryId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalSynchronizationTemplateSchemaDirectoryIDInsensitively(input string) (*ServicePrincipalSynchronizationTemplateSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalSynchronizationTemplateSchemaDirectoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalSynchronizationTemplateSchemaDirectoryId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.SynchronizationTemplateId, ok = parsed.Parsed["synchronizationTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", *parsed)
	}

	if id.DirectoryDefinitionId, ok = parsed.Parsed["directoryDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalSynchronizationTemplateSchemaDirectoryID checks that 'input' can be parsed as a Service Principal Synchronization Template Schema Directory ID
func ValidateServicePrincipalSynchronizationTemplateSchemaDirectoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalSynchronizationTemplateSchemaDirectoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Synchronization Template Schema Directory ID
func (id ServicePrincipalSynchronizationTemplateSchemaDirectoryId) ID() string {
	fmtString := "/servicePrincipals/%s/synchronization/templates/%s/schema/directories/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.SynchronizationTemplateId, id.DirectoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Synchronization Template Schema Directory ID
func (id ServicePrincipalSynchronizationTemplateSchemaDirectoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticSynchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("staticTemplates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("synchronizationTemplateId", "synchronizationTemplateIdValue"),
		resourceids.StaticSegment("staticSchema", "schema", "schema"),
		resourceids.StaticSegment("staticDirectories", "directories", "directories"),
		resourceids.UserSpecifiedSegment("directoryDefinitionId", "directoryDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Synchronization Template Schema Directory ID
func (id ServicePrincipalSynchronizationTemplateSchemaDirectoryId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Synchronization Template: %q", id.SynchronizationTemplateId),
		fmt.Sprintf("Directory Definition: %q", id.DirectoryDefinitionId),
	}
	return fmt.Sprintf("Service Principal Synchronization Template Schema Directory (%s)", strings.Join(components, "\n"))
}
