package userjoinedteamscheduleopenshift

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleOpenShiftId{}

// UserJoinedTeamScheduleOpenShiftId is a struct representing the Resource ID for a User Joined Team Schedule Open Shift
type UserJoinedTeamScheduleOpenShiftId struct {
	UserId      string
	TeamId      string
	OpenShiftId string
}

// NewUserJoinedTeamScheduleOpenShiftID returns a new UserJoinedTeamScheduleOpenShiftId struct
func NewUserJoinedTeamScheduleOpenShiftID(userId string, teamId string, openShiftId string) UserJoinedTeamScheduleOpenShiftId {
	return UserJoinedTeamScheduleOpenShiftId{
		UserId:      userId,
		TeamId:      teamId,
		OpenShiftId: openShiftId,
	}
}

// ParseUserJoinedTeamScheduleOpenShiftID parses 'input' into a UserJoinedTeamScheduleOpenShiftId
func ParseUserJoinedTeamScheduleOpenShiftID(input string) (*UserJoinedTeamScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleOpenShiftId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OpenShiftId, ok = parsed.Parsed["openShiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleOpenShiftIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleOpenShiftId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleOpenShiftIDInsensitively(input string) (*UserJoinedTeamScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleOpenShiftId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OpenShiftId, ok = parsed.Parsed["openShiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleOpenShiftID checks that 'input' can be parsed as a User Joined Team Schedule Open Shift ID
func ValidateUserJoinedTeamScheduleOpenShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleOpenShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Open Shift ID
func (id UserJoinedTeamScheduleOpenShiftId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/openShifts/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.OpenShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Open Shift ID
func (id UserJoinedTeamScheduleOpenShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOpenShifts", "openShifts", "openShifts"),
		resourceids.UserSpecifiedSegment("openShiftId", "openShiftIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Open Shift ID
func (id UserJoinedTeamScheduleOpenShiftId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Open Shift: %q", id.OpenShiftId),
	}
	return fmt.Sprintf("User Joined Team Schedule Open Shift (%s)", strings.Join(components, "\n"))
}
