package mejoinedteamprimarychannelmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelMessageReplyId{}

// MeJoinedTeamPrimaryChannelMessageReplyId is a struct representing the Resource ID for a Me Joined Team Primary Channel Message Reply
type MeJoinedTeamPrimaryChannelMessageReplyId struct {
	TeamId         string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewMeJoinedTeamPrimaryChannelMessageReplyID returns a new MeJoinedTeamPrimaryChannelMessageReplyId struct
func NewMeJoinedTeamPrimaryChannelMessageReplyID(teamId string, chatMessageId string, chatMessageId1 string) MeJoinedTeamPrimaryChannelMessageReplyId {
	return MeJoinedTeamPrimaryChannelMessageReplyId{
		TeamId:         teamId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseMeJoinedTeamPrimaryChannelMessageReplyID parses 'input' into a MeJoinedTeamPrimaryChannelMessageReplyId
func ParseMeJoinedTeamPrimaryChannelMessageReplyID(input string) (*MeJoinedTeamPrimaryChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMessageReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMessageReplyId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPrimaryChannelMessageReplyIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPrimaryChannelMessageReplyId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPrimaryChannelMessageReplyIDInsensitively(input string) (*MeJoinedTeamPrimaryChannelMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMessageReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMessageReplyId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPrimaryChannelMessageReplyID checks that 'input' can be parsed as a Me Joined Team Primary Channel Message Reply ID
func ValidateMeJoinedTeamPrimaryChannelMessageReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPrimaryChannelMessageReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Primary Channel Message Reply ID
func (id MeJoinedTeamPrimaryChannelMessageReplyId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Primary Channel Message Reply ID
func (id MeJoinedTeamPrimaryChannelMessageReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
	}
}

// String returns a human-readable description of this Me Joined Team Primary Channel Message Reply ID
func (id MeJoinedTeamPrimaryChannelMessageReplyId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Me Joined Team Primary Channel Message Reply (%s)", strings.Join(components, "\n"))
}
