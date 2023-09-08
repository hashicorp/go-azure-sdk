package mejoinedteamprimarychannelmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelMessageReplyHostedContentId{}

// MeJoinedTeamPrimaryChannelMessageReplyHostedContentId is a struct representing the Resource ID for a Me Joined Team Primary Channel Message Reply Hosted Content
type MeJoinedTeamPrimaryChannelMessageReplyHostedContentId struct {
	TeamId                     string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewMeJoinedTeamPrimaryChannelMessageReplyHostedContentID returns a new MeJoinedTeamPrimaryChannelMessageReplyHostedContentId struct
func NewMeJoinedTeamPrimaryChannelMessageReplyHostedContentID(teamId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) MeJoinedTeamPrimaryChannelMessageReplyHostedContentId {
	return MeJoinedTeamPrimaryChannelMessageReplyHostedContentId{
		TeamId:                     teamId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeJoinedTeamPrimaryChannelMessageReplyHostedContentID parses 'input' into a MeJoinedTeamPrimaryChannelMessageReplyHostedContentId
func ParseMeJoinedTeamPrimaryChannelMessageReplyHostedContentID(input string) (*MeJoinedTeamPrimaryChannelMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMessageReplyHostedContentId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPrimaryChannelMessageReplyHostedContentIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPrimaryChannelMessageReplyHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPrimaryChannelMessageReplyHostedContentIDInsensitively(input string) (*MeJoinedTeamPrimaryChannelMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMessageReplyHostedContentId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPrimaryChannelMessageReplyHostedContentID checks that 'input' can be parsed as a Me Joined Team Primary Channel Message Reply Hosted Content ID
func ValidateMeJoinedTeamPrimaryChannelMessageReplyHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPrimaryChannelMessageReplyHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Primary Channel Message Reply Hosted Content ID
func (id MeJoinedTeamPrimaryChannelMessageReplyHostedContentId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Primary Channel Message Reply Hosted Content ID
func (id MeJoinedTeamPrimaryChannelMessageReplyHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Primary Channel Message Reply Hosted Content ID
func (id MeJoinedTeamPrimaryChannelMessageReplyHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Joined Team Primary Channel Message Reply Hosted Content (%s)", strings.Join(components, "\n"))
}
