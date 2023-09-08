package applicationowner

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplicationOwnerId{}

// ApplicationOwnerId is a struct representing the Resource ID for a Application Owner
type ApplicationOwnerId struct {
	ApplicationId     string
	DirectoryObjectId string
}

// NewApplicationOwnerID returns a new ApplicationOwnerId struct
func NewApplicationOwnerID(applicationId string, directoryObjectId string) ApplicationOwnerId {
	return ApplicationOwnerId{
		ApplicationId:     applicationId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseApplicationOwnerID parses 'input' into a ApplicationOwnerId
func ParseApplicationOwnerID(input string) (*ApplicationOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationOwnerId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseApplicationOwnerIDInsensitively parses 'input' case-insensitively into a ApplicationOwnerId
// note: this method should only be used for API response data and not user input
func ParseApplicationOwnerIDInsensitively(input string) (*ApplicationOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplicationOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplicationOwnerId{}

	if id.ApplicationId, ok = parsed.Parsed["applicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applicationId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateApplicationOwnerID checks that 'input' can be parsed as a Application Owner ID
func ValidateApplicationOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplicationOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Application Owner ID
func (id ApplicationOwnerId) ID() string {
	fmtString := "/applications/%s/owners/%s"
	return fmt.Sprintf(fmtString, id.ApplicationId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Application Owner ID
func (id ApplicationOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticApplications", "applications", "applications"),
		resourceids.UserSpecifiedSegment("applicationId", "applicationIdValue"),
		resourceids.StaticSegment("staticOwners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Application Owner ID
func (id ApplicationOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Application: %q", id.ApplicationId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Application Owner (%s)", strings.Join(components, "\n"))
}
