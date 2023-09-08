package groupteamprimarychannelmessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelMessageId{}

// GroupTeamPrimaryChannelMessageId is a struct representing the Resource ID for a Group Team Primary Channel Message
type GroupTeamPrimaryChannelMessageId struct {
	GroupId       string
	ChatMessageId string
}

// NewGroupTeamPrimaryChannelMessageID returns a new GroupTeamPrimaryChannelMessageId struct
func NewGroupTeamPrimaryChannelMessageID(groupId string, chatMessageId string) GroupTeamPrimaryChannelMessageId {
	return GroupTeamPrimaryChannelMessageId{
		GroupId:       groupId,
		ChatMessageId: chatMessageId,
	}
}

// ParseGroupTeamPrimaryChannelMessageID parses 'input' into a GroupTeamPrimaryChannelMessageId
func ParseGroupTeamPrimaryChannelMessageID(input string) (*GroupTeamPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMessageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamPrimaryChannelMessageIDInsensitively parses 'input' case-insensitively into a GroupTeamPrimaryChannelMessageId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPrimaryChannelMessageIDInsensitively(input string) (*GroupTeamPrimaryChannelMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMessageId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamPrimaryChannelMessageID checks that 'input' can be parsed as a Group Team Primary Channel Message ID
func ValidateGroupTeamPrimaryChannelMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPrimaryChannelMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Primary Channel Message ID
func (id GroupTeamPrimaryChannelMessageId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/messages/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Primary Channel Message ID
func (id GroupTeamPrimaryChannelMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
	}
}

// String returns a human-readable description of this Group Team Primary Channel Message ID
func (id GroupTeamPrimaryChannelMessageId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Group Team Primary Channel Message (%s)", strings.Join(components, "\n"))
}
