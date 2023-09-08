package userjoinedteamchanneltab

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelId{}

// UserJoinedTeamChannelId is a struct representing the Resource ID for a User Joined Team Channel
type UserJoinedTeamChannelId struct {
	UserId    string
	TeamId    string
	ChannelId string
}

// NewUserJoinedTeamChannelID returns a new UserJoinedTeamChannelId struct
func NewUserJoinedTeamChannelID(userId string, teamId string, channelId string) UserJoinedTeamChannelId {
	return UserJoinedTeamChannelId{
		UserId:    userId,
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseUserJoinedTeamChannelID parses 'input' into a UserJoinedTeamChannelId
func ParseUserJoinedTeamChannelID(input string) (*UserJoinedTeamChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamChannelIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelIDInsensitively(input string) (*UserJoinedTeamChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamChannelID checks that 'input' can be parsed as a User Joined Team Channel ID
func ValidateUserJoinedTeamChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel ID
func (id UserJoinedTeamChannelId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel ID
func (id UserJoinedTeamChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Channel ID
func (id UserJoinedTeamChannelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("User Joined Team Channel (%s)", strings.Join(components, "\n"))
}
