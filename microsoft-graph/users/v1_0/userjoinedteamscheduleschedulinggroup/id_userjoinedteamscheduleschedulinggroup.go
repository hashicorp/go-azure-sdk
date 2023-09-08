package userjoinedteamscheduleschedulinggroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleSchedulingGroupId{}

// UserJoinedTeamScheduleSchedulingGroupId is a struct representing the Resource ID for a User Joined Team Schedule Scheduling Group
type UserJoinedTeamScheduleSchedulingGroupId struct {
	UserId            string
	TeamId            string
	SchedulingGroupId string
}

// NewUserJoinedTeamScheduleSchedulingGroupID returns a new UserJoinedTeamScheduleSchedulingGroupId struct
func NewUserJoinedTeamScheduleSchedulingGroupID(userId string, teamId string, schedulingGroupId string) UserJoinedTeamScheduleSchedulingGroupId {
	return UserJoinedTeamScheduleSchedulingGroupId{
		UserId:            userId,
		TeamId:            teamId,
		SchedulingGroupId: schedulingGroupId,
	}
}

// ParseUserJoinedTeamScheduleSchedulingGroupID parses 'input' into a UserJoinedTeamScheduleSchedulingGroupId
func ParseUserJoinedTeamScheduleSchedulingGroupID(input string) (*UserJoinedTeamScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleSchedulingGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SchedulingGroupId, ok = parsed.Parsed["schedulingGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleSchedulingGroupIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleSchedulingGroupId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleSchedulingGroupIDInsensitively(input string) (*UserJoinedTeamScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleSchedulingGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SchedulingGroupId, ok = parsed.Parsed["schedulingGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleSchedulingGroupID checks that 'input' can be parsed as a User Joined Team Schedule Scheduling Group ID
func ValidateUserJoinedTeamScheduleSchedulingGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleSchedulingGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Scheduling Group ID
func (id UserJoinedTeamScheduleSchedulingGroupId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/schedulingGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.SchedulingGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Scheduling Group ID
func (id UserJoinedTeamScheduleSchedulingGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticSchedulingGroups", "schedulingGroups", "schedulingGroups"),
		resourceids.UserSpecifiedSegment("schedulingGroupId", "schedulingGroupIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Scheduling Group ID
func (id UserJoinedTeamScheduleSchedulingGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Scheduling Group: %q", id.SchedulingGroupId),
	}
	return fmt.Sprintf("User Joined Team Schedule Scheduling Group (%s)", strings.Join(components, "\n"))
}
