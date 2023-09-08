package groupteamownermailboxsetting

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamOwnerId{}

// GroupTeamOwnerId is a struct representing the Resource ID for a Group Team Owner
type GroupTeamOwnerId struct {
	GroupId string
	UserId  string
}

// NewGroupTeamOwnerID returns a new GroupTeamOwnerId struct
func NewGroupTeamOwnerID(groupId string, userId string) GroupTeamOwnerId {
	return GroupTeamOwnerId{
		GroupId: groupId,
		UserId:  userId,
	}
}

// ParseGroupTeamOwnerID parses 'input' into a GroupTeamOwnerId
func ParseGroupTeamOwnerID(input string) (*GroupTeamOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamOwnerId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamOwnerIDInsensitively parses 'input' case-insensitively into a GroupTeamOwnerId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamOwnerIDInsensitively(input string) (*GroupTeamOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamOwnerId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamOwnerID checks that 'input' can be parsed as a Group Team Owner ID
func ValidateGroupTeamOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Owner ID
func (id GroupTeamOwnerId) ID() string {
	fmtString := "/groups/%s/team/owners/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.UserId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Owner ID
func (id GroupTeamOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticOwners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
	}
}

// String returns a human-readable description of this Group Team Owner ID
func (id GroupTeamOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("User: %q", id.UserId),
	}
	return fmt.Sprintf("Group Team Owner (%s)", strings.Join(components, "\n"))
}
