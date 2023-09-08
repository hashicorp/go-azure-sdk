package userjoinedteamscheduletimesoff

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleTimesOffId{}

// UserJoinedTeamScheduleTimesOffId is a struct representing the Resource ID for a User Joined Team Schedule Times Off
type UserJoinedTeamScheduleTimesOffId struct {
	UserId    string
	TeamId    string
	TimeOffId string
}

// NewUserJoinedTeamScheduleTimesOffID returns a new UserJoinedTeamScheduleTimesOffId struct
func NewUserJoinedTeamScheduleTimesOffID(userId string, teamId string, timeOffId string) UserJoinedTeamScheduleTimesOffId {
	return UserJoinedTeamScheduleTimesOffId{
		UserId:    userId,
		TeamId:    teamId,
		TimeOffId: timeOffId,
	}
}

// ParseUserJoinedTeamScheduleTimesOffID parses 'input' into a UserJoinedTeamScheduleTimesOffId
func ParseUserJoinedTeamScheduleTimesOffID(input string) (*UserJoinedTeamScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleTimesOffId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleTimesOffId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffId, ok = parsed.Parsed["timeOffId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleTimesOffIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleTimesOffId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleTimesOffIDInsensitively(input string) (*UserJoinedTeamScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleTimesOffId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleTimesOffId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffId, ok = parsed.Parsed["timeOffId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleTimesOffID checks that 'input' can be parsed as a User Joined Team Schedule Times Off ID
func ValidateUserJoinedTeamScheduleTimesOffID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleTimesOffID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Times Off ID
func (id UserJoinedTeamScheduleTimesOffId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/timesOff/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TimeOffId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Times Off ID
func (id UserJoinedTeamScheduleTimesOffId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimesOff", "timesOff", "timesOff"),
		resourceids.UserSpecifiedSegment("timeOffId", "timeOffIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Times Off ID
func (id UserJoinedTeamScheduleTimesOffId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off: %q", id.TimeOffId),
	}
	return fmt.Sprintf("User Joined Team Schedule Times Off (%s)", strings.Join(components, "\n"))
}
