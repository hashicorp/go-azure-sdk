package userjoinedteamchanneltab

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamChannelTabId{}

// UserJoinedTeamChannelTabId is a struct representing the Resource ID for a User Joined Team Channel Tab
type UserJoinedTeamChannelTabId struct {
	UserId     string
	TeamId     string
	ChannelId  string
	TeamsTabId string
}

// NewUserJoinedTeamChannelTabID returns a new UserJoinedTeamChannelTabId struct
func NewUserJoinedTeamChannelTabID(userId string, teamId string, channelId string, teamsTabId string) UserJoinedTeamChannelTabId {
	return UserJoinedTeamChannelTabId{
		UserId:     userId,
		TeamId:     teamId,
		ChannelId:  channelId,
		TeamsTabId: teamsTabId,
	}
}

// ParseUserJoinedTeamChannelTabID parses 'input' into a UserJoinedTeamChannelTabId
func ParseUserJoinedTeamChannelTabID(input string) (*UserJoinedTeamChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelTabId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamChannelTabIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamChannelTabId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamChannelTabIDInsensitively(input string) (*UserJoinedTeamChannelTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamChannelTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamChannelTabId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChannelId, ok = parsed.Parsed["channelId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "channelId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamChannelTabID checks that 'input' can be parsed as a User Joined Team Channel Tab ID
func ValidateUserJoinedTeamChannelTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamChannelTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Channel Tab ID
func (id UserJoinedTeamChannelTabId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/channels/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Channel Tab ID
func (id UserJoinedTeamChannelTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticChannels", "channels", "channels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
		resourceids.StaticSegment("staticTabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Channel Tab ID
func (id UserJoinedTeamChannelTabId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("User Joined Team Channel Tab (%s)", strings.Join(components, "\n"))
}
