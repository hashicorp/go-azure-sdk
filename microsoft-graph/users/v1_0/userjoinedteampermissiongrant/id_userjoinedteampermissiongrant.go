package userjoinedteampermissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPermissionGrantId{}

// UserJoinedTeamPermissionGrantId is a struct representing the Resource ID for a User Joined Team Permission Grant
type UserJoinedTeamPermissionGrantId struct {
	UserId                            string
	TeamId                            string
	ResourceSpecificPermissionGrantId string
}

// NewUserJoinedTeamPermissionGrantID returns a new UserJoinedTeamPermissionGrantId struct
func NewUserJoinedTeamPermissionGrantID(userId string, teamId string, resourceSpecificPermissionGrantId string) UserJoinedTeamPermissionGrantId {
	return UserJoinedTeamPermissionGrantId{
		UserId:                            userId,
		TeamId:                            teamId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseUserJoinedTeamPermissionGrantID parses 'input' into a UserJoinedTeamPermissionGrantId
func ParseUserJoinedTeamPermissionGrantID(input string) (*UserJoinedTeamPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPermissionGrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamPermissionGrantIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamPermissionGrantIDInsensitively(input string) (*UserJoinedTeamPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPermissionGrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamPermissionGrantID checks that 'input' can be parsed as a User Joined Team Permission Grant ID
func ValidateUserJoinedTeamPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Permission Grant ID
func (id UserJoinedTeamPermissionGrantId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Permission Grant ID
func (id UserJoinedTeamPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPermissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Permission Grant ID
func (id UserJoinedTeamPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("User Joined Team Permission Grant (%s)", strings.Join(components, "\n"))
}
