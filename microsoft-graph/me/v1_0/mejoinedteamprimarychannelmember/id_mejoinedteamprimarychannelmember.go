package mejoinedteamprimarychannelmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelMemberId{}

// MeJoinedTeamPrimaryChannelMemberId is a struct representing the Resource ID for a Me Joined Team Primary Channel Member
type MeJoinedTeamPrimaryChannelMemberId struct {
	TeamId               string
	ConversationMemberId string
}

// NewMeJoinedTeamPrimaryChannelMemberID returns a new MeJoinedTeamPrimaryChannelMemberId struct
func NewMeJoinedTeamPrimaryChannelMemberID(teamId string, conversationMemberId string) MeJoinedTeamPrimaryChannelMemberId {
	return MeJoinedTeamPrimaryChannelMemberId{
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeJoinedTeamPrimaryChannelMemberID parses 'input' into a MeJoinedTeamPrimaryChannelMemberId
func ParseMeJoinedTeamPrimaryChannelMemberID(input string) (*MeJoinedTeamPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPrimaryChannelMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPrimaryChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPrimaryChannelMemberIDInsensitively(input string) (*MeJoinedTeamPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPrimaryChannelMemberID checks that 'input' can be parsed as a Me Joined Team Primary Channel Member ID
func ValidateMeJoinedTeamPrimaryChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPrimaryChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Primary Channel Member ID
func (id MeJoinedTeamPrimaryChannelMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/members/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Primary Channel Member ID
func (id MeJoinedTeamPrimaryChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Primary Channel Member ID
func (id MeJoinedTeamPrimaryChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Primary Channel Member (%s)", strings.Join(components, "\n"))
}
