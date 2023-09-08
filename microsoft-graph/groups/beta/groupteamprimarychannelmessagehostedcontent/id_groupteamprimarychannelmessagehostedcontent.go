package groupteamprimarychannelmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelMessageHostedContentId{}

// GroupTeamPrimaryChannelMessageHostedContentId is a struct representing the Resource ID for a Group Team Primary Channel Message Hosted Content
type GroupTeamPrimaryChannelMessageHostedContentId struct {
	GroupId                    string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewGroupTeamPrimaryChannelMessageHostedContentID returns a new GroupTeamPrimaryChannelMessageHostedContentId struct
func NewGroupTeamPrimaryChannelMessageHostedContentID(groupId string, chatMessageId string, chatMessageHostedContentId string) GroupTeamPrimaryChannelMessageHostedContentId {
	return GroupTeamPrimaryChannelMessageHostedContentId{
		GroupId:                    groupId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseGroupTeamPrimaryChannelMessageHostedContentID parses 'input' into a GroupTeamPrimaryChannelMessageHostedContentId
func ParseGroupTeamPrimaryChannelMessageHostedContentID(input string) (*GroupTeamPrimaryChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMessageHostedContentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamPrimaryChannelMessageHostedContentIDInsensitively parses 'input' case-insensitively into a GroupTeamPrimaryChannelMessageHostedContentId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPrimaryChannelMessageHostedContentIDInsensitively(input string) (*GroupTeamPrimaryChannelMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMessageHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMessageHostedContentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamPrimaryChannelMessageHostedContentID checks that 'input' can be parsed as a Group Team Primary Channel Message Hosted Content ID
func ValidateGroupTeamPrimaryChannelMessageHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPrimaryChannelMessageHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Primary Channel Message Hosted Content ID
func (id GroupTeamPrimaryChannelMessageHostedContentId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Primary Channel Message Hosted Content ID
func (id GroupTeamPrimaryChannelMessageHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this Group Team Primary Channel Message Hosted Content ID
func (id GroupTeamPrimaryChannelMessageHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Group Team Primary Channel Message Hosted Content (%s)", strings.Join(components, "\n"))
}
