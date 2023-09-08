package groupteammember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamMemberId{}

// GroupTeamMemberId is a struct representing the Resource ID for a Group Team Member
type GroupTeamMemberId struct {
	GroupId              string
	ConversationMemberId string
}

// NewGroupTeamMemberID returns a new GroupTeamMemberId struct
func NewGroupTeamMemberID(groupId string, conversationMemberId string) GroupTeamMemberId {
	return GroupTeamMemberId{
		GroupId:              groupId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseGroupTeamMemberID parses 'input' into a GroupTeamMemberId
func ParseGroupTeamMemberID(input string) (*GroupTeamMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamMemberIDInsensitively parses 'input' case-insensitively into a GroupTeamMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamMemberIDInsensitively(input string) (*GroupTeamMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamMemberID checks that 'input' can be parsed as a Group Team Member ID
func ValidateGroupTeamMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Member ID
func (id GroupTeamMemberId) ID() string {
	fmtString := "/groups/%s/team/members/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Member ID
func (id GroupTeamMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Group Team Member ID
func (id GroupTeamMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Team Member (%s)", strings.Join(components, "\n"))
}
