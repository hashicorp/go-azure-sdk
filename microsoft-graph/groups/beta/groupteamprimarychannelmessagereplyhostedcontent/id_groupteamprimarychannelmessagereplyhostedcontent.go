package groupteamprimarychannelmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelMessageReplyHostedContentId{}

// GroupTeamPrimaryChannelMessageReplyHostedContentId is a struct representing the Resource ID for a Group Team Primary Channel Message Reply Hosted Content
type GroupTeamPrimaryChannelMessageReplyHostedContentId struct {
	GroupId                    string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewGroupTeamPrimaryChannelMessageReplyHostedContentID returns a new GroupTeamPrimaryChannelMessageReplyHostedContentId struct
func NewGroupTeamPrimaryChannelMessageReplyHostedContentID(groupId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) GroupTeamPrimaryChannelMessageReplyHostedContentId {
	return GroupTeamPrimaryChannelMessageReplyHostedContentId{
		GroupId:                    groupId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseGroupTeamPrimaryChannelMessageReplyHostedContentID parses 'input' into a GroupTeamPrimaryChannelMessageReplyHostedContentId
func ParseGroupTeamPrimaryChannelMessageReplyHostedContentID(input string) (*GroupTeamPrimaryChannelMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMessageReplyHostedContentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ParseGroupTeamPrimaryChannelMessageReplyHostedContentIDInsensitively parses 'input' case-insensitively into a GroupTeamPrimaryChannelMessageReplyHostedContentId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPrimaryChannelMessageReplyHostedContentIDInsensitively(input string) (*GroupTeamPrimaryChannelMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMessageReplyHostedContentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
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

// ValidateGroupTeamPrimaryChannelMessageReplyHostedContentID checks that 'input' can be parsed as a Group Team Primary Channel Message Reply Hosted Content ID
func ValidateGroupTeamPrimaryChannelMessageReplyHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPrimaryChannelMessageReplyHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Primary Channel Message Reply Hosted Content ID
func (id GroupTeamPrimaryChannelMessageReplyHostedContentId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Primary Channel Message Reply Hosted Content ID
func (id GroupTeamPrimaryChannelMessageReplyHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this Group Team Primary Channel Message Reply Hosted Content ID
func (id GroupTeamPrimaryChannelMessageReplyHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Group Team Primary Channel Message Reply Hosted Content (%s)", strings.Join(components, "\n"))
}
