package userjoinedteamprimarychannelsharedwithteamteam

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamPrimaryChannelSharedWithTeamId{}

// UserJoinedTeamPrimaryChannelSharedWithTeamId is a struct representing the Resource ID for a User Joined Team Primary Channel Shared With Team
type UserJoinedTeamPrimaryChannelSharedWithTeamId struct {
	UserId                      string
	TeamId                      string
	SharedWithChannelTeamInfoId string
}

// NewUserJoinedTeamPrimaryChannelSharedWithTeamID returns a new UserJoinedTeamPrimaryChannelSharedWithTeamId struct
func NewUserJoinedTeamPrimaryChannelSharedWithTeamID(userId string, teamId string, sharedWithChannelTeamInfoId string) UserJoinedTeamPrimaryChannelSharedWithTeamId {
	return UserJoinedTeamPrimaryChannelSharedWithTeamId{
		UserId:                      userId,
		TeamId:                      teamId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseUserJoinedTeamPrimaryChannelSharedWithTeamID parses 'input' into a UserJoinedTeamPrimaryChannelSharedWithTeamId
func ParseUserJoinedTeamPrimaryChannelSharedWithTeamID(input string) (*UserJoinedTeamPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelSharedWithTeamId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamPrimaryChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamPrimaryChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamPrimaryChannelSharedWithTeamIDInsensitively(input string) (*UserJoinedTeamPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamPrimaryChannelSharedWithTeamId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamPrimaryChannelSharedWithTeamID checks that 'input' can be parsed as a User Joined Team Primary Channel Shared With Team ID
func ValidateUserJoinedTeamPrimaryChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamPrimaryChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Primary Channel Shared With Team ID
func (id UserJoinedTeamPrimaryChannelSharedWithTeamId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/primaryChannel/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Primary Channel Shared With Team ID
func (id UserJoinedTeamPrimaryChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Primary Channel Shared With Team ID
func (id UserJoinedTeamPrimaryChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("User Joined Team Primary Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
