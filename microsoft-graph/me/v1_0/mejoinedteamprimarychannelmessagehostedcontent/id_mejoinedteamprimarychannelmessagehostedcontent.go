package mejoinedteamprimarychannelmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamPrimaryChannelMessageHostedContentId{}

// MeJoinedTeamPrimaryChannelMessageHostedContentId is a struct representing the Resource ID for a Me Joined Team Primary Channel Message Hosted Content
type MeJoinedTeamPrimaryChannelMessageHostedContentId struct {
	TeamId                     string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewMeJoinedTeamPrimaryChannelMessageHostedContentID returns a new MeJoinedTeamPrimaryChannelMessageHostedContentId struct
func NewMeJoinedTeamPrimaryChannelMessageHostedContentID(teamId string, chatMessageId string, chatMessageHostedContentId string) MeJoinedTeamPrimaryChannelMessageHostedContentId {
	return MeJoinedTeamPrimaryChannelMessageHostedContentId{
		TeamId:                     teamId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeJoinedTeamPrimaryChannelMessageHostedContentID parses 'input' into a MeJoinedTeamPrimaryChannelMessageHostedContentId
func ParseMeJoinedTeamPrimaryChannelMessageHostedContentID(input string) (*MeJoinedTeamPrimaryChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMessageHostedContentId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamPrimaryChannelMessageHostedContentIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamPrimaryChannelMessageHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamPrimaryChannelMessageHostedContentIDInsensitively(input string) (*MeJoinedTeamPrimaryChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamPrimaryChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamPrimaryChannelMessageHostedContentId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamPrimaryChannelMessageHostedContentID checks that 'input' can be parsed as a Me Joined Team Primary Channel Message Hosted Content ID
func ValidateMeJoinedTeamPrimaryChannelMessageHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamPrimaryChannelMessageHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Primary Channel Message Hosted Content ID
func (id MeJoinedTeamPrimaryChannelMessageHostedContentId) ID() string {
	fmtString := "/me/joinedTeams/%s/primaryChannel/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Primary Channel Message Hosted Content ID
func (id MeJoinedTeamPrimaryChannelMessageHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Primary Channel Message Hosted Content ID
func (id MeJoinedTeamPrimaryChannelMessageHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Joined Team Primary Channel Message Hosted Content (%s)", strings.Join(components, "\n"))
}
