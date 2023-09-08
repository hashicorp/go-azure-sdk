package groupteamoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamOperationId{}

// GroupTeamOperationId is a struct representing the Resource ID for a Group Team Operation
type GroupTeamOperationId struct {
	GroupId               string
	TeamsAsyncOperationId string
}

// NewGroupTeamOperationID returns a new GroupTeamOperationId struct
func NewGroupTeamOperationID(groupId string, teamsAsyncOperationId string) GroupTeamOperationId {
	return GroupTeamOperationId{
		GroupId:               groupId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseGroupTeamOperationID parses 'input' into a GroupTeamOperationId
func ParseGroupTeamOperationID(input string) (*GroupTeamOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamOperationIDInsensitively parses 'input' case-insensitively into a GroupTeamOperationId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamOperationIDInsensitively(input string) (*GroupTeamOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamOperationId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamOperationID checks that 'input' can be parsed as a Group Team Operation ID
func ValidateGroupTeamOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Operation ID
func (id GroupTeamOperationId) ID() string {
	fmtString := "/groups/%s/team/operations/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Operation ID
func (id GroupTeamOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationIdValue"),
	}
}

// String returns a human-readable description of this Group Team Operation ID
func (id GroupTeamOperationId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("Group Team Operation (%s)", strings.Join(components, "\n"))
}
