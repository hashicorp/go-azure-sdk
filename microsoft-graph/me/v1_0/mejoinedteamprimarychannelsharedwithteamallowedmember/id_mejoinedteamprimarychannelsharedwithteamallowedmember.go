package mejoinedteamprimarychannelsharedwithteamallowedmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

// MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId is a struct representing the Resource ID for a Me Joined Team Primary Channel Shared With Team Allowed Member
type MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId struct {
	TeamId                      string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID returns a new MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId struct
func NewMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(teamId string, sharedWithChannelTeamInfoId string, conversationMemberId string) MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId {
	return MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{
		TeamId:                      teamId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID parses 'input' into a MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId
func ParseMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(input string) (*MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively(input string) (*MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID checks that 'input' can be parsed as a Me Joined Team Primary Channel Shared With Team Allowed Member ID
func ValidateMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Primary Channel Shared With Team Allowed Member ID
func (id MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Primary Channel Shared With Team Allowed Member ID
func (id MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
		resourceids.StaticSegment("staticAllowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Primary Channel Shared With Team Allowed Member ID
func (id MeJoinedTeamPrimaryChannelSharedWithTeamAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Primary Channel Shared With Team Allowed Member (%s)", strings.Join(components, "\n"))
}
