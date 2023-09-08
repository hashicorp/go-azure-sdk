package groupteamscheduletimecard

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleTimeCardId{}

// GroupTeamScheduleTimeCardId is a struct representing the Resource ID for a Group Team Schedule Time Card
type GroupTeamScheduleTimeCardId struct {
	GroupId    string
	TimeCardId string
}

// NewGroupTeamScheduleTimeCardID returns a new GroupTeamScheduleTimeCardId struct
func NewGroupTeamScheduleTimeCardID(groupId string, timeCardId string) GroupTeamScheduleTimeCardId {
	return GroupTeamScheduleTimeCardId{
		GroupId:    groupId,
		TimeCardId: timeCardId,
	}
}

// ParseGroupTeamScheduleTimeCardID parses 'input' into a GroupTeamScheduleTimeCardId
func ParseGroupTeamScheduleTimeCardID(input string) (*GroupTeamScheduleTimeCardId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleTimeCardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleTimeCardId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TimeCardId, ok = parsed.Parsed["timeCardId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeCardId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleTimeCardIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleTimeCardId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleTimeCardIDInsensitively(input string) (*GroupTeamScheduleTimeCardId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleTimeCardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleTimeCardId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TimeCardId, ok = parsed.Parsed["timeCardId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeCardId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleTimeCardID checks that 'input' can be parsed as a Group Team Schedule Time Card ID
func ValidateGroupTeamScheduleTimeCardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleTimeCardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Time Card ID
func (id GroupTeamScheduleTimeCardId) ID() string {
	fmtString := "/groups/%s/team/schedule/timeCards/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TimeCardId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Time Card ID
func (id GroupTeamScheduleTimeCardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimeCards", "timeCards", "timeCards"),
		resourceids.UserSpecifiedSegment("timeCardId", "timeCardIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Time Card ID
func (id GroupTeamScheduleTimeCardId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Time Card: %q", id.TimeCardId),
	}
	return fmt.Sprintf("Group Team Schedule Time Card (%s)", strings.Join(components, "\n"))
}
