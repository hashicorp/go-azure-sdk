package applicationextensionproperty

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationExtensionPropertyId{}

// ApplicationExtensionPropertyId is a struct representing the Resource ID for a Application Extension Property
type ApplicationExtensionPropertyId struct {
	ApplicationId       string
	ExtensionPropertyId string
}

// NewApplicationExtensionPropertyID returns a new ApplicationExtensionPropertyId struct
func NewApplicationExtensionPropertyID(applicationId string, extensionPropertyId string) ApplicationExtensionPropertyId {
	return ApplicationExtensionPropertyId{
		ApplicationId:       applicationId,
		ExtensionPropertyId: extensionPropertyId,
	}
}

// ParseApplicationExtensionPropertyID parses 'input' into a ApplicationExtensionPropertyId
func ParseApplicationExtensionPropertyID(input string) (*ApplicationExtensionPropertyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationExtensionPropertyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationExtensionPropertyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.ExtensionPropertyId, ok = parsed.Parsed["extensionPropertyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionPropertyId", *parsed)
	}

	return &id, nil
}

// ParseApplicationExtensionPropertyIDInsensitively parses 'input' case-insensitively into a ApplicationExtensionPropertyId
// note: this method should only be used for API response data and not user input
func ParseApplicationExtensionPropertyIDInsensitively(input string) (*ApplicationExtensionPropertyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationExtensionPropertyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationExtensionPropertyId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.ExtensionPropertyId, ok = parsed.Parsed["extensionPropertyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionPropertyId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationExtensionPropertyID checks that 'input' can be parsed as a Application Extension Property ID
func ValidateApplicationExtensionPropertyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationExtensionPropertyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Extension Property ID
func (id ApplicationExtensionPropertyId) ID() string {
	fmtString := "/applications/%s/extensionProperties/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.ExtensionPropertyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Extension Property ID
func (id ApplicationExtensionPropertyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticExtensionProperties", "extensionProperties", "extensionProperties"),
		resourceids.UserSpecifiedSegment("extensionPropertyId", "extensionPropertyIdValue"),
	}
}

// String returns a human-readable description of this Application Extension Property ID
func (id ApplicationExtensionPropertyId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Extension Property: %q", id.ExtensionPropertyId),
	}
	return fmt.Sprintf("Application Extension Property (%s)", strings.Join(components, "\n"))
}
