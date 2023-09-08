package userapproleassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppRoleAssignmentId{}

// UserAppRoleAssignmentId is a struct representing the Resource ID for a User App Role Assignment
type UserAppRoleAssignmentId struct {
	UserId              string
	AppRoleAssignmentId string
}

// NewUserAppRoleAssignmentID returns a new UserAppRoleAssignmentId struct
func NewUserAppRoleAssignmentID(userId string, appRoleAssignmentId string) UserAppRoleAssignmentId {
	return UserAppRoleAssignmentId{
		UserId:              userId,
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseUserAppRoleAssignmentID parses 'input' into a UserAppRoleAssignmentId
func ParseUserAppRoleAssignmentID(input string) (*UserAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppRoleAssignmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseUserAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a UserAppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseUserAppRoleAssignmentIDInsensitively(input string) (*UserAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppRoleAssignmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserAppRoleAssignmentID checks that 'input' can be parsed as a User App Role Assignment ID
func ValidateUserAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User App Role Assignment ID
func (id UserAppRoleAssignmentId) ID() string {
	fmtString := "/users/%s/appRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User App Role Assignment ID
func (id UserAppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAppRoleAssignments", "appRoleAssignments", "appRoleAssignments"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentIdValue"),
	}
}

// String returns a human-readable description of this User App Role Assignment ID
func (id UserAppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("User App Role Assignment (%s)", strings.Join(components, "\n"))
}
