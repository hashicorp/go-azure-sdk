package groupteamprimarychannelmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelMemberId{}

// GroupTeamPrimaryChannelMemberId is a struct representing the Resource ID for a Group Team Primary Channel Member
type GroupTeamPrimaryChannelMemberId struct {
	GroupId              string
	ConversationMemberId string
}

// NewGroupTeamPrimaryChannelMemberID returns a new GroupTeamPrimaryChannelMemberId struct
func NewGroupTeamPrimaryChannelMemberID(groupId string, conversationMemberId string) GroupTeamPrimaryChannelMemberId {
	return GroupTeamPrimaryChannelMemberId{
		GroupId:              groupId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseGroupTeamPrimaryChannelMemberID parses 'input' into a GroupTeamPrimaryChannelMemberId
func ParseGroupTeamPrimaryChannelMemberID(input string) (*GroupTeamPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamPrimaryChannelMemberIDInsensitively parses 'input' case-insensitively into a GroupTeamPrimaryChannelMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPrimaryChannelMemberIDInsensitively(input string) (*GroupTeamPrimaryChannelMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamPrimaryChannelMemberID checks that 'input' can be parsed as a Group Team Primary Channel Member ID
func ValidateGroupTeamPrimaryChannelMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPrimaryChannelMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Primary Channel Member ID
func (id GroupTeamPrimaryChannelMemberId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Primary Channel Member ID
func (id GroupTeamPrimaryChannelMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Group Team Primary Channel Member ID
func (id GroupTeamPrimaryChannelMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Team Primary Channel Member (%s)", strings.Join(components, "\n"))
}
