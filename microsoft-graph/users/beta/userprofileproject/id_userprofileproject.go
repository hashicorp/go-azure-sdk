package userprofileproject

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserProfileProjectId{}

// UserProfileProjectId is a struct representing the Resource ID for a User Profile Project
type UserProfileProjectId struct {
	UserId                 string
	ProjectParticipationId string
}

// NewUserProfileProjectID returns a new UserProfileProjectId struct
func NewUserProfileProjectID(userId string, projectParticipationId string) UserProfileProjectId {
	return UserProfileProjectId{
		UserId:                 userId,
		ProjectParticipationId: projectParticipationId,
	}
}

// ParseUserProfileProjectID parses 'input' into a UserProfileProjectId
func ParseUserProfileProjectID(input string) (*UserProfileProjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileProjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileProjectId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ProjectParticipationId, ok = parsed.Parsed["projectParticipationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "projectParticipationId", *parsed)
	}

	return &id, nil
}

// ParseUserProfileProjectIDInsensitively parses 'input' case-insensitively into a UserProfileProjectId
// note: this method should only be used for API response data and not user input
func ParseUserProfileProjectIDInsensitively(input string) (*UserProfileProjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserProfileProjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserProfileProjectId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ProjectParticipationId, ok = parsed.Parsed["projectParticipationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "projectParticipationId", *parsed)
	}

	return &id, nil
}

// ValidateUserProfileProjectID checks that 'input' can be parsed as a User Profile Project ID
func ValidateUserProfileProjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserProfileProjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Profile Project ID
func (id UserProfileProjectId) ID() string {
	fmtString := "/users/%s/profile/projects/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ProjectParticipationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Profile Project ID
func (id UserProfileProjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticProfile", "profile", "profile"),
		resourceids.StaticSegment("staticProjects", "projects", "projects"),
		resourceids.UserSpecifiedSegment("projectParticipationId", "projectParticipationIdValue"),
	}
}

// String returns a human-readable description of this User Profile Project ID
func (id UserProfileProjectId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Project Participation: %q", id.ProjectParticipationId),
	}
	return fmt.Sprintf("User Profile Project (%s)", strings.Join(components, "\n"))
}
