package applicationsynchronizationtemplate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationSynchronizationTemplateId{}

// ApplicationSynchronizationTemplateId is a struct representing the Resource ID for a Application Synchronization Template
type ApplicationSynchronizationTemplateId struct {
	ApplicationId             string
	SynchronizationTemplateId string
}

// NewApplicationSynchronizationTemplateID returns a new ApplicationSynchronizationTemplateId struct
func NewApplicationSynchronizationTemplateID(applicationId string, synchronizationTemplateId string) ApplicationSynchronizationTemplateId {
	return ApplicationSynchronizationTemplateId{
		ApplicationId:             applicationId,
		SynchronizationTemplateId: synchronizationTemplateId,
	}
}

// ParseApplicationSynchronizationTemplateID parses 'input' into a ApplicationSynchronizationTemplateId
func ParseApplicationSynchronizationTemplateID(input string) (*ApplicationSynchronizationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationSynchronizationTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationSynchronizationTemplateId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.SynchronizationTemplateId, ok = parsed.Parsed["synchronizationTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", *parsed)
	}

	return &id, nil
}

// ParseApplicationSynchronizationTemplateIDInsensitively parses 'input' case-insensitively into a ApplicationSynchronizationTemplateId
// note: this method should only be used for API response data and not user input
func ParseApplicationSynchronizationTemplateIDInsensitively(input string) (*ApplicationSynchronizationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationSynchronizationTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationSynchronizationTemplateId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.SynchronizationTemplateId, ok = parsed.Parsed["synchronizationTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationSynchronizationTemplateID checks that 'input' can be parsed as a Application Synchronization Template ID
func ValidateApplicationSynchronizationTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationSynchronizationTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Synchronization Template ID
func (id ApplicationSynchronizationTemplateId) ID() string {
	fmtString := "/applications/%s/synchronization/templates/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.SynchronizationTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Synchronization Template ID
func (id ApplicationSynchronizationTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticSynchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("staticTemplates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("synchronizationTemplateId", "synchronizationTemplateIdValue"),
	}
}

// String returns a human-readable description of this Application Synchronization Template ID
func (id ApplicationSynchronizationTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Synchronization Template: %q", id.SynchronizationTemplateId),
	}
	return fmt.Sprintf("Application Synchronization Template (%s)", strings.Join(components, "\n"))
}
