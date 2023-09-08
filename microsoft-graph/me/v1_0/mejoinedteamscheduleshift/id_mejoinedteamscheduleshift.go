package mejoinedteamscheduleshift

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleShiftId{}

// MeJoinedTeamScheduleShiftId is a struct representing the Resource ID for a Me Joined Team Schedule Shift
type MeJoinedTeamScheduleShiftId struct {
	TeamId  string
	ShiftId string
}

// NewMeJoinedTeamScheduleShiftID returns a new MeJoinedTeamScheduleShiftId struct
func NewMeJoinedTeamScheduleShiftID(teamId string, shiftId string) MeJoinedTeamScheduleShiftId {
	return MeJoinedTeamScheduleShiftId{
		TeamId:  teamId,
		ShiftId: shiftId,
	}
}

// ParseMeJoinedTeamScheduleShiftID parses 'input' into a MeJoinedTeamScheduleShiftId
func ParseMeJoinedTeamScheduleShiftID(input string) (*MeJoinedTeamScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleShiftId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ShiftId, ok = parsed.Parsed["shiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "shiftId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamScheduleShiftIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamScheduleShiftId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamScheduleShiftIDInsensitively(input string) (*MeJoinedTeamScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleShiftId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ShiftId, ok = parsed.Parsed["shiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "shiftId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamScheduleShiftID checks that 'input' can be parsed as a Me Joined Team Schedule Shift ID
func ValidateMeJoinedTeamScheduleShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamScheduleShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Schedule Shift ID
func (id MeJoinedTeamScheduleShiftId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/shifts/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.ShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Schedule Shift ID
func (id MeJoinedTeamScheduleShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticShifts", "shifts", "shifts"),
		resourceids.UserSpecifiedSegment("shiftId", "shiftIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Schedule Shift ID
func (id MeJoinedTeamScheduleShiftId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shift: %q", id.ShiftId),
	}
	return fmt.Sprintf("Me Joined Team Schedule Shift (%s)", strings.Join(components, "\n"))
}
