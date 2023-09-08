package userjoinedteamscheduletimeoffrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleTimeOffRequestId{}

// UserJoinedTeamScheduleTimeOffRequestId is a struct representing the Resource ID for a User Joined Team Schedule Time Off Request
type UserJoinedTeamScheduleTimeOffRequestId struct {
	UserId           string
	TeamId           string
	TimeOffRequestId string
}

// NewUserJoinedTeamScheduleTimeOffRequestID returns a new UserJoinedTeamScheduleTimeOffRequestId struct
func NewUserJoinedTeamScheduleTimeOffRequestID(userId string, teamId string, timeOffRequestId string) UserJoinedTeamScheduleTimeOffRequestId {
	return UserJoinedTeamScheduleTimeOffRequestId{
		UserId:           userId,
		TeamId:           teamId,
		TimeOffRequestId: timeOffRequestId,
	}
}

// ParseUserJoinedTeamScheduleTimeOffRequestID parses 'input' into a UserJoinedTeamScheduleTimeOffRequestId
func ParseUserJoinedTeamScheduleTimeOffRequestID(input string) (*UserJoinedTeamScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleTimeOffRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffRequestId, ok = parsed.Parsed["timeOffRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleTimeOffRequestIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleTimeOffRequestId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleTimeOffRequestIDInsensitively(input string) (*UserJoinedTeamScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleTimeOffRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffRequestId, ok = parsed.Parsed["timeOffRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleTimeOffRequestID checks that 'input' can be parsed as a User Joined Team Schedule Time Off Request ID
func ValidateUserJoinedTeamScheduleTimeOffRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleTimeOffRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Time Off Request ID
func (id UserJoinedTeamScheduleTimeOffRequestId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/timeOffRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.TimeOffRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Time Off Request ID
func (id UserJoinedTeamScheduleTimeOffRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimeOffRequests", "timeOffRequests", "timeOffRequests"),
		resourceids.UserSpecifiedSegment("timeOffRequestId", "timeOffRequestIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Time Off Request ID
func (id UserJoinedTeamScheduleTimeOffRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off Request: %q", id.TimeOffRequestId),
	}
	return fmt.Sprintf("User Joined Team Schedule Time Off Request (%s)", strings.Join(components, "\n"))
}
