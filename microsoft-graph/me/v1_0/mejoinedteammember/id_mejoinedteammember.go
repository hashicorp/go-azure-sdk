package mejoinedteammember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamMemberId{}

// MeJoinedTeamMemberId is a struct representing the Resource ID for a Me Joined Team Member
type MeJoinedTeamMemberId struct {
	TeamId               string
	ConversationMemberId string
}

// NewMeJoinedTeamMemberID returns a new MeJoinedTeamMemberId struct
func NewMeJoinedTeamMemberID(teamId string, conversationMemberId string) MeJoinedTeamMemberId {
	return MeJoinedTeamMemberId{
		TeamId:               teamId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeJoinedTeamMemberID parses 'input' into a MeJoinedTeamMemberId
func ParseMeJoinedTeamMemberID(input string) (*MeJoinedTeamMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamMemberIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamMemberId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamMemberIDInsensitively(input string) (*MeJoinedTeamMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamMemberId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamMemberID checks that 'input' can be parsed as a Me Joined Team Member ID
func ValidateMeJoinedTeamMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Member ID
func (id MeJoinedTeamMemberId) ID() string {
	fmtString := "/me/joinedTeams/%s/members/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Member ID
func (id MeJoinedTeamMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Member ID
func (id MeJoinedTeamMemberId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Joined Team Member (%s)", strings.Join(components, "\n"))
}
