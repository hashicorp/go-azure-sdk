package userjoinedteamscheduleshift

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleShiftId{}

// UserJoinedTeamScheduleShiftId is a struct representing the Resource ID for a User Joined Team Schedule Shift
type UserJoinedTeamScheduleShiftId struct {
	UserId  string
	TeamId  string
	ShiftId string
}

// NewUserJoinedTeamScheduleShiftID returns a new UserJoinedTeamScheduleShiftId struct
func NewUserJoinedTeamScheduleShiftID(userId string, teamId string, shiftId string) UserJoinedTeamScheduleShiftId {
	return UserJoinedTeamScheduleShiftId{
		UserId:  userId,
		TeamId:  teamId,
		ShiftId: shiftId,
	}
}

// ParseUserJoinedTeamScheduleShiftID parses 'input' into a UserJoinedTeamScheduleShiftId
func ParseUserJoinedTeamScheduleShiftID(input string) (*UserJoinedTeamScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleShiftId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ShiftId, ok = parsed.Parsed["shiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "shiftId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleShiftIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleShiftId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleShiftIDInsensitively(input string) (*UserJoinedTeamScheduleShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleShiftId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.ShiftId, ok = parsed.Parsed["shiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "shiftId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleShiftID checks that 'input' can be parsed as a User Joined Team Schedule Shift ID
func ValidateUserJoinedTeamScheduleShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Shift ID
func (id UserJoinedTeamScheduleShiftId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/shifts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.ShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Shift ID
func (id UserJoinedTeamScheduleShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticShifts", "shifts", "shifts"),
		resourceids.UserSpecifiedSegment("shiftId", "shiftIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Shift ID
func (id UserJoinedTeamScheduleShiftId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Shift: %q", id.ShiftId),
	}
	return fmt.Sprintf("User Joined Team Schedule Shift (%s)", strings.Join(components, "\n"))
}
