package applicationsynchronizationtemplateschemadirectory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationSynchronizationTemplateSchemaDirectoryId{}

// ApplicationSynchronizationTemplateSchemaDirectoryId is a struct representing the Resource ID for a Application Synchronization Template Schema Directory
type ApplicationSynchronizationTemplateSchemaDirectoryId struct {
	ApplicationId             string
	SynchronizationTemplateId string
	DirectoryDefinitionId     string
}

// NewApplicationSynchronizationTemplateSchemaDirectoryID returns a new ApplicationSynchronizationTemplateSchemaDirectoryId struct
func NewApplicationSynchronizationTemplateSchemaDirectoryID(applicationId string, synchronizationTemplateId string, directoryDefinitionId string) ApplicationSynchronizationTemplateSchemaDirectoryId {
	return ApplicationSynchronizationTemplateSchemaDirectoryId{
		ApplicationId:             applicationId,
		SynchronizationTemplateId: synchronizationTemplateId,
		DirectoryDefinitionId:     directoryDefinitionId,
	}
}

// ParseApplicationSynchronizationTemplateSchemaDirectoryID parses 'input' into a ApplicationSynchronizationTemplateSchemaDirectoryId
func ParseApplicationSynchronizationTemplateSchemaDirectoryID(input string) (*ApplicationSynchronizationTemplateSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationSynchronizationTemplateSchemaDirectoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationSynchronizationTemplateSchemaDirectoryId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.SynchronizationTemplateId, ok = parsed.Parsed["synchronizationTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", *parsed)
	}

	if id.DirectoryDefinitionId, ok = parsed.Parsed["directoryDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseApplicationSynchronizationTemplateSchemaDirectoryIDInsensitively parses 'input' case-insensitively into a ApplicationSynchronizationTemplateSchemaDirectoryId
// note: this method should only be used for API response data and not user input
func ParseApplicationSynchronizationTemplateSchemaDirectoryIDInsensitively(input string) (*ApplicationSynchronizationTemplateSchemaDirectoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationSynchronizationTemplateSchemaDirectoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationSynchronizationTemplateSchemaDirectoryId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.SynchronizationTemplateId, ok = parsed.Parsed["synchronizationTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", *parsed)
	}

	if id.DirectoryDefinitionId, ok = parsed.Parsed["directoryDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationSynchronizationTemplateSchemaDirectoryID checks that 'input' can be parsed as a Application Synchronization Template Schema Directory ID
func ValidateApplicationSynchronizationTemplateSchemaDirectoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationSynchronizationTemplateSchemaDirectoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Synchronization Template Schema Directory ID
func (id ApplicationSynchronizationTemplateSchemaDirectoryId) ID() string {
	fmtString := "/applications/%s/synchronization/templates/%s/schema/directories/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.SynchronizationTemplateId, id.DirectoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Synchronization Template Schema Directory ID
func (id ApplicationSynchronizationTemplateSchemaDirectoryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticSynchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("staticTemplates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("synchronizationTemplateId", "synchronizationTemplateIdValue"),
		resourceids.StaticSegment("staticSchema", "schema", "schema"),
		resourceids.StaticSegment("staticDirectories", "directories", "directories"),
		resourceids.UserSpecifiedSegment("directoryDefinitionId", "directoryDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Application Synchronization Template Schema Directory ID
func (id ApplicationSynchronizationTemplateSchemaDirectoryId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Synchronization Template: %q", id.SynchronizationTemplateId),
		fmt.Sprintf("Directory Definition: %q", id.DirectoryDefinitionId),
	}
	return fmt.Sprintf("Application Synchronization Template Schema Directory (%s)", strings.Join(components, "\n"))
}
