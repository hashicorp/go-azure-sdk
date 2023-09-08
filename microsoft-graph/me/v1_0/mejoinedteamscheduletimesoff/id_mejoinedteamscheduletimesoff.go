package mejoinedteamscheduletimesoff

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleTimesOffId{}

// MeJoinedTeamScheduleTimesOffId is a struct representing the Resource ID for a Me Joined Team Schedule Times Off
type MeJoinedTeamScheduleTimesOffId struct {
	TeamId    string
	TimeOffId string
}

// NewMeJoinedTeamScheduleTimesOffID returns a new MeJoinedTeamScheduleTimesOffId struct
func NewMeJoinedTeamScheduleTimesOffID(teamId string, timeOffId string) MeJoinedTeamScheduleTimesOffId {
	return MeJoinedTeamScheduleTimesOffId{
		TeamId:    teamId,
		TimeOffId: timeOffId,
	}
}

// ParseMeJoinedTeamScheduleTimesOffID parses 'input' into a MeJoinedTeamScheduleTimesOffId
func ParseMeJoinedTeamScheduleTimesOffID(input string) (*MeJoinedTeamScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleTimesOffId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleTimesOffId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffId, ok = parsed.Parsed["timeOffId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamScheduleTimesOffIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamScheduleTimesOffId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamScheduleTimesOffIDInsensitively(input string) (*MeJoinedTeamScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleTimesOffId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleTimesOffId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.TimeOffId, ok = parsed.Parsed["timeOffId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamScheduleTimesOffID checks that 'input' can be parsed as a Me Joined Team Schedule Times Off ID
func ValidateMeJoinedTeamScheduleTimesOffID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamScheduleTimesOffID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Schedule Times Off ID
func (id MeJoinedTeamScheduleTimesOffId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/timesOff/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.TimeOffId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Schedule Times Off ID
func (id MeJoinedTeamScheduleTimesOffId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimesOff", "timesOff", "timesOff"),
		resourceids.UserSpecifiedSegment("timeOffId", "timeOffIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Schedule Times Off ID
func (id MeJoinedTeamScheduleTimesOffId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Time Off: %q", id.TimeOffId),
	}
	return fmt.Sprintf("Me Joined Team Schedule Times Off (%s)", strings.Join(components, "\n"))
}
