package userjoinedteamincomingchannel

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamIncomingChannelId{}

// UserJoinedTeamIncomingChannelId is a struct representing the Resource ID for a User Joined Team Incoming Channel
type UserJoinedTeamIncomingChannelId struct {
	UserId    string
	TeamId    string
	ChannelId string
}

// NewUserJoinedTeamIncomingChannelID returns a new UserJoinedTeamIncomingChannelId struct
func NewUserJoinedTeamIncomingChannelID(userId string, teamId string, channelId string) UserJoinedTeamIncomingChannelId {
	return UserJoinedTeamIncomingChannelId{
		UserId:    userId,
		TeamId:    teamId,
		ChannelId: channelId,
	}
}

// ParseUserJoinedTeamIncomingChannelID parses 'input' into a UserJoinedTeamIncomingChannelId
func ParseUserJoinedTeamIncomingChannelID(input string) (*UserJoinedTeamIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamIncomingChannelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamIncomingChannelId{}

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

// ParseUserJoinedTeamIncomingChannelIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamIncomingChannelId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamIncomingChannelIDInsensitively(input string) (*UserJoinedTeamIncomingChannelId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamIncomingChannelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamIncomingChannelId{}

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

// ValidateUserJoinedTeamIncomingChannelID checks that 'input' can be parsed as a User Joined Team Incoming Channel ID
func ValidateUserJoinedTeamIncomingChannelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamIncomingChannelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Incoming Channel ID
func (id UserJoinedTeamIncomingChannelId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/incomingChannels/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ChannelId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Incoming Channel ID
func (id UserJoinedTeamIncomingChannelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticIncomingChannels", "incomingChannels", "incomingChannels"),
		resourceids.UserSpecifiedSegment("channelId", "channelIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Incoming Channel ID
func (id UserJoinedTeamIncomingChannelId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Channel: %q", id.ChannelId),
	}
	return fmt.Sprintf("User Joined Team Incoming Channel (%s)", strings.Join(components, "\n"))
}
