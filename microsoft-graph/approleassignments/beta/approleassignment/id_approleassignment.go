package approleassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AppRoleAssignmentId{}

// AppRoleAssignmentId is a struct representing the Resource ID for a App Role Assignment
type AppRoleAssignmentId struct {
	AppRoleAssignmentId string
}

// NewAppRoleAssignmentID returns a new AppRoleAssignmentId struct
func NewAppRoleAssignmentID(appRoleAssignmentId string) AppRoleAssignmentId {
	return AppRoleAssignmentId{
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseAppRoleAssignmentID parses 'input' into a AppRoleAssignmentId
func ParseAppRoleAssignmentID(input string) (*AppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(AppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AppRoleAssignmentId{}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a AppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseAppRoleAssignmentIDInsensitively(input string) (*AppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(AppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AppRoleAssignmentId{}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateAppRoleAssignmentID checks that 'input' can be parsed as a App Role Assignment ID
func ValidateAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted App Role Assignment ID
func (id AppRoleAssignmentId) ID() string {
	fmtString := "/appRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this App Role Assignment ID
func (id AppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticAppRoleAssignments", "appRoleAssignments", "appRoleAssignments"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentIdValue"),
	}
}

// String returns a human-readable description of this App Role Assignment ID
func (id AppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("App Role Assignment (%s)", strings.Join(components, "\n"))
}
