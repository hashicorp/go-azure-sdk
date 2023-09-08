package groupteamscheduleshift

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleShiftId{}

// GroupTeamScheduleShiftId is a struct representing the Resource ID for a Group Team Schedule Shift
type GroupTeamScheduleShiftId struct {
	GroupId string
	ShiftId string
}

// NewGroupTeamScheduleShiftID returns a new GroupTeamScheduleShiftId struct
func NewGroupTeamScheduleShiftID(groupId string, shiftId string) GroupTeamScheduleShiftId {
	return GroupTeamScheduleShiftId{
		GroupId: groupId,
		ShiftId: shiftId,
	}
}

// ParseGroupTeamScheduleShiftID parses 'input' into a GroupTeamScheduleShiftId
func ParseGroupTeamScheduleShiftID(input string) (*GroupTeamScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleShiftId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ShiftId, ok = parsed.Parsed["shiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "shiftId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleShiftIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleShiftId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleShiftIDInsensitively(input string) (*GroupTeamScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleShiftId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.ShiftId, ok = parsed.Parsed["shiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "shiftId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleShiftID checks that 'input' can be parsed as a Group Team Schedule Shift ID
func ValidateGroupTeamScheduleShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Shift ID
func (id GroupTeamScheduleShiftId) ID() string {
	fmtString := "/groups/%s/team/schedule/shifts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Shift ID
func (id GroupTeamScheduleShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticShifts", "shifts", "shifts"),
		resourceids.UserSpecifiedSegment("shiftId", "shiftIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Shift ID
func (id GroupTeamScheduleShiftId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Shift: %q", id.ShiftId),
	}
	return fmt.Sprintf("Group Team Schedule Shift (%s)", strings.Join(components, "\n"))
}
