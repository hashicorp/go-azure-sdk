package userjoinedteamchannelsharedwithteamteam

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelSharedWithTeamId{}

// UserJoinedTeamChannelSharedWithTeamId is a struct representing the Resource ID for a User Joined Team Channel Shared With Team
type UserJoinedTeamChannelSharedWithTeamId struct {
	UserId                      string
	TeamId                      string
	ChannelId                   string
	SharedWithChannelTeamInfoId string
}

// NewUserJoinedTeamChannelSharedWithTeamID returns a new UserJoinedTeamChannelSharedWithTeamId struct
func NewUserJoinedTeamChannelSharedWithTeamID(userId string, teamId string, channelId string, sharedWithChannelTeamInfoId string) UserJoinedTeamChannelSharedWithTeamId {
	return UserJoinedTeamChannelSharedWithTeamId{
		UserId:                      userId,
		TeamId:                      teamId,
		ChannelId:                   channelId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseUserJoinedTeamChannelSharedWithTeamID parses 'input' into a UserJoinedTeamChannelSharedWithTeamId
func ParseUserJoinedTeamChannelSharedWithTeamID(input string) (*UserJoinedTeamChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelSharedWithTeamId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelSharedWithTeamIDInsensitively(input string) (*UserJoinedTeamChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelSharedWithTeamId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamChannelSharedWithTeamID checks that 'input' can be parsed as a User Joined Team Channel Shared With Team ID
func ValidateUserJoinedTeamChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel Shared With Team ID
func (id UserJoinedTeamChannelSharedWithTeamId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel Shared With Team ID
func (id UserJoinedTeamChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Channel Shared With Team ID
func (id UserJoinedTeamChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("User Joined Team Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
