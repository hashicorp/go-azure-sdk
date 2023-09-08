package groupteamprimarychannelsharedwithteamallowedmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

// GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId is a struct representing the Resource ID for a Group Team Primary Channel Shared With Team Allowed Member
type GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId struct {
	GroupId                     string
	SharedWithChannelTeamInfoId string
	ConversationMemberId        string
}

// NewGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID returns a new GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId struct
func NewGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID(groupId string, sharedWithChannelTeamInfoId string, conversationMemberId string) GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId {
	return GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{
		GroupId:                     groupId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
		ConversationMemberId:        conversationMemberId,
	}
}

// ParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID parses 'input' into a GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId
func ParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID(input string) (*GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively parses 'input' case-insensitively into a GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberIDInsensitively(input string) (*GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID checks that 'input' can be parsed as a Group Team Primary Channel Shared With Team Allowed Member ID
func ValidateGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPrimaryChannelSharedWithTeamAllowedMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Primary Channel Shared With Team Allowed Member ID
func (id GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/sharedWithTeams/%s/allowedMembers/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SharedWithChannelTeamInfoId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Primary Channel Shared With Team Allowed Member ID
func (id GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
		resourceids.StaticSegment("staticAllowedMembers", "allowedMembers", "allowedMembers"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Group Team Primary Channel Shared With Team Allowed Member ID
func (id GroupTeamPrimaryChannelSharedWithTeamAllowedMemberId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Group Team Primary Channel Shared With Team Allowed Member (%s)", strings.Join(components, "\n"))
}
