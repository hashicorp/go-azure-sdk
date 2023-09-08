package groupteamprimarychannelsharedwithteamteam

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamPrimaryChannelSharedWithTeamId{}

// GroupTeamPrimaryChannelSharedWithTeamId is a struct representing the Resource ID for a Group Team Primary Channel Shared With Team
type GroupTeamPrimaryChannelSharedWithTeamId struct {
	GroupId                     string
	SharedWithChannelTeamInfoId string
}

// NewGroupTeamPrimaryChannelSharedWithTeamID returns a new GroupTeamPrimaryChannelSharedWithTeamId struct
func NewGroupTeamPrimaryChannelSharedWithTeamID(groupId string, sharedWithChannelTeamInfoId string) GroupTeamPrimaryChannelSharedWithTeamId {
	return GroupTeamPrimaryChannelSharedWithTeamId{
		GroupId:                     groupId,
		SharedWithChannelTeamInfoId: sharedWithChannelTeamInfoId,
	}
}

// ParseGroupTeamPrimaryChannelSharedWithTeamID parses 'input' into a GroupTeamPrimaryChannelSharedWithTeamId
func ParseGroupTeamPrimaryChannelSharedWithTeamID(input string) (*GroupTeamPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelSharedWithTeamId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamPrimaryChannelSharedWithTeamIDInsensitively parses 'input' case-insensitively into a GroupTeamPrimaryChannelSharedWithTeamId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamPrimaryChannelSharedWithTeamIDInsensitively(input string) (*GroupTeamPrimaryChannelSharedWithTeamId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamPrimaryChannelSharedWithTeamId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamPrimaryChannelSharedWithTeamId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SharedWithChannelTeamInfoId, ok = parsed.Parsed["sharedWithChannelTeamInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sharedWithChannelTeamInfoId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamPrimaryChannelSharedWithTeamID checks that 'input' can be parsed as a Group Team Primary Channel Shared With Team ID
func ValidateGroupTeamPrimaryChannelSharedWithTeamID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamPrimaryChannelSharedWithTeamID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Primary Channel Shared With Team ID
func (id GroupTeamPrimaryChannelSharedWithTeamId) ID() string {
	fmtString := "/groups/%s/team/primaryChannel/sharedWithTeams/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SharedWithChannelTeamInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Primary Channel Shared With Team ID
func (id GroupTeamPrimaryChannelSharedWithTeamId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticPrimaryChannel", "primaryChannel", "primaryChannel"),
		resourceids.StaticSegment("staticSharedWithTeams", "sharedWithTeams", "sharedWithTeams"),
		resourceids.UserSpecifiedSegment("sharedWithChannelTeamInfoId", "sharedWithChannelTeamInfoIdValue"),
	}
}

// String returns a human-readable description of this Group Team Primary Channel Shared With Team ID
func (id GroupTeamPrimaryChannelSharedWithTeamId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Shared With Channel Team Info: %q", id.SharedWithChannelTeamInfoId),
	}
	return fmt.Sprintf("Group Team Primary Channel Shared With Team (%s)", strings.Join(components, "\n"))
}
