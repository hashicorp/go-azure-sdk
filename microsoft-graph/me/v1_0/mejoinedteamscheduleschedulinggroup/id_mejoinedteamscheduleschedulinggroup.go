package mejoinedteamscheduleschedulinggroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleSchedulingGroupId{}

// MeJoinedTeamScheduleSchedulingGroupId is a struct representing the Resource ID for a Me Joined Team Schedule Scheduling Group
type MeJoinedTeamScheduleSchedulingGroupId struct {
	TeamId            string
	SchedulingGroupId string
}

// NewMeJoinedTeamScheduleSchedulingGroupID returns a new MeJoinedTeamScheduleSchedulingGroupId struct
func NewMeJoinedTeamScheduleSchedulingGroupID(teamId string, schedulingGroupId string) MeJoinedTeamScheduleSchedulingGroupId {
	return MeJoinedTeamScheduleSchedulingGroupId{
		TeamId:            teamId,
		SchedulingGroupId: schedulingGroupId,
	}
}

// ParseMeJoinedTeamScheduleSchedulingGroupID parses 'input' into a MeJoinedTeamScheduleSchedulingGroupId
func ParseMeJoinedTeamScheduleSchedulingGroupID(input string) (*MeJoinedTeamScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleSchedulingGroupId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SchedulingGroupId, ok = parsed.Parsed["schedulingGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamScheduleSchedulingGroupIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamScheduleSchedulingGroupId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamScheduleSchedulingGroupIDInsensitively(input string) (*MeJoinedTeamScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleSchedulingGroupId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SchedulingGroupId, ok = parsed.Parsed["schedulingGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamScheduleSchedulingGroupID checks that 'input' can be parsed as a Me Joined Team Schedule Scheduling Group ID
func ValidateMeJoinedTeamScheduleSchedulingGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamScheduleSchedulingGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Schedule Scheduling Group ID
func (id MeJoinedTeamScheduleSchedulingGroupId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/schedulingGroups/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.SchedulingGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Schedule Scheduling Group ID
func (id MeJoinedTeamScheduleSchedulingGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticSchedulingGroups", "schedulingGroups", "schedulingGroups"),
		resourceids.UserSpecifiedSegment("schedulingGroupId", "schedulingGroupIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Schedule Scheduling Group ID
func (id MeJoinedTeamScheduleSchedulingGroupId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Scheduling Group: %q", id.SchedulingGroupId),
	}
	return fmt.Sprintf("Me Joined Team Schedule Scheduling Group (%s)", strings.Join(components, "\n"))
}
