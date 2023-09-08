package userjoinedteamscheduletimeoffreason

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleTimeOffReasonId{}

// UserJoinedTeamScheduleTimeOffReasonId is a struct representing the Resource ID for a User Joined Team Schedule Time Off Reason
type UserJoinedTeamScheduleTimeOffReasonId struct {
	UserId          string
	TeamId          string
	TimeOffReasonId string
}

// NewUserJoinedTeamScheduleTimeOffReasonID returns a new UserJoinedTeamScheduleTimeOffReasonId struct
func NewUserJoinedTeamScheduleTimeOffReasonID(userId string, teamId string, timeOffReasonId string) UserJoinedTeamScheduleTimeOffReasonId {
	return UserJoinedTeamScheduleTimeOffReasonId{
		UserId:          userId,
		TeamId:          teamId,
		TimeOffReasonId: timeOffReasonId,
	}
}

// ParseUserJoinedTeamScheduleTimeOffReasonID parses 'input' into a UserJoinedTeamScheduleTimeOffReasonId
func ParseUserJoinedTeamScheduleTimeOffReasonID(input string) (*UserJoinedTeamScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleTimeOffReasonId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffReasonId, ok = parsed.Parsed["timeOffReasonId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffReasonId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleTimeOffReasonIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleTimeOffReasonId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleTimeOffReasonIDInsensitively(input string) (*UserJoinedTeamScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleTimeOffReasonId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffReasonId, ok = parsed.Parsed["timeOffReasonId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffReasonId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleTimeOffReasonID checks that 'input' can be parsed as a User Joined Team Schedule Time Off Reason ID
func ValidateUserJoinedTeamScheduleTimeOffReasonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleTimeOffReasonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Time Off Reason ID
func (id UserJoinedTeamScheduleTimeOffReasonId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/timeOffReasons/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TimeOffReasonId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Time Off Reason ID
func (id UserJoinedTeamScheduleTimeOffReasonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimeOffReasons", "timeOffReasons", "timeOffReasons"),
		resourceids.UserSpecifiedSegment("timeOffReasonId", "timeOffReasonIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Time Off Reason ID
func (id UserJoinedTeamScheduleTimeOffReasonId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off Reason: %q", id.TimeOffReasonId),
	}
	return fmt.Sprintf("User Joined Team Schedule Time Off Reason (%s)", strings.Join(components, "\n"))
}
