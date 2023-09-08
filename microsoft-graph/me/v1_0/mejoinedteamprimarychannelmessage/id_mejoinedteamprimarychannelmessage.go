package mejoinedteamprimarychannelmessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelMessageId{}

// MeJoinedTeamPrimaryChannelMessageId is a struct representing the Resource ID for a Me Joined Team Primary Channel Message
type MeJoinedTeamPrimaryChannelMessageId struct {
	TeamId        string
	ChatMessageId string
}

// NewMeJoinedTeamPrimaryChannelMessageID returns a new MeJoinedTeamPrimaryChannelMessageId struct
func NewMeJoinedTeamPrimaryChannelMessageID(teamId string, chatMessageId string) MeJoinedTeamPrimaryChannelMessageId {
	return MeJoinedTeamPrimaryChannelMessageId{
		TeamId:        teamId,
		ChatMessageId: chatMessageId,
	}
}

// ParseMeJoinedTeamPrimaryChannelMessageID parses 'input' into a MeJoinedTeamPrimaryChannelMessageId
func ParseMeJoinedTeamPrimaryChannelMessageID(input string) (*MeJoinedTeamPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMessageId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPrimaryChannelMessageIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPrimaryChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPrimaryChannelMessageIDInsensitively(input string) (*MeJoinedTeamPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMessageId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPrimaryChannelMessageID checks that 'input' can be parsed as a Me Joined Team Primary Channel Message ID
func ValidateMeJoinedTeamPrimaryChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPrimaryChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Primary Channel Message ID
func (id MeJoinedTeamPrimaryChannelMessageId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/messages/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Primary Channel Message ID
func (id MeJoinedTeamPrimaryChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Primary Channel Message ID
func (id MeJoinedTeamPrimaryChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Me Joined Team Primary Channel Message (%s)", strings.Join(components, "\n"))
}
