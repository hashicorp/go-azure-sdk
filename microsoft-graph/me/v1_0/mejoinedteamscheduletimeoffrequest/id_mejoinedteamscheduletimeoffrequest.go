package mejoinedteamscheduletimeoffrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleTimeOffRequestId{}

// MeJoinedTeamScheduleTimeOffRequestId is a struct representing the Resource ID for a Me Joined Team Schedule Time Off Request
type MeJoinedTeamScheduleTimeOffRequestId struct {
	TeamId           string
	TimeOffRequestId string
}

// NewMeJoinedTeamScheduleTimeOffRequestID returns a new MeJoinedTeamScheduleTimeOffRequestId struct
func NewMeJoinedTeamScheduleTimeOffRequestID(teamId string, timeOffRequestId string) MeJoinedTeamScheduleTimeOffRequestId {
	return MeJoinedTeamScheduleTimeOffRequestId{
		TeamId:           teamId,
		TimeOffRequestId: timeOffRequestId,
	}
}

// ParseMeJoinedTeamScheduleTimeOffRequestID parses 'input' into a MeJoinedTeamScheduleTimeOffRequestId
func ParseMeJoinedTeamScheduleTimeOffRequestID(input string) (*MeJoinedTeamScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleTimeOffRequestId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffRequestId, ok = parsed.Parsed["timeOffRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamScheduleTimeOffRequestIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamScheduleTimeOffRequestId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamScheduleTimeOffRequestIDInsensitively(input string) (*MeJoinedTeamScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleTimeOffRequestId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffRequestId, ok = parsed.Parsed["timeOffRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamScheduleTimeOffRequestID checks that 'input' can be parsed as a Me Joined Team Schedule Time Off Request ID
func ValidateMeJoinedTeamScheduleTimeOffRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamScheduleTimeOffRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Schedule Time Off Request ID
func (id MeJoinedTeamScheduleTimeOffRequestId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/timeOffRequests/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TimeOffRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Schedule Time Off Request ID
func (id MeJoinedTeamScheduleTimeOffRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimeOffRequests", "timeOffRequests", "timeOffRequests"),
		resourceids.UserSpecifiedSegment("timeOffRequestId", "timeOffRequestIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Schedule Time Off Request ID
func (id MeJoinedTeamScheduleTimeOffRequestId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off Request: %q", id.TimeOffRequestId),
	}
	return fmt.Sprintf("Me Joined Team Schedule Time Off Request (%s)", strings.Join(components, "\n"))
}
