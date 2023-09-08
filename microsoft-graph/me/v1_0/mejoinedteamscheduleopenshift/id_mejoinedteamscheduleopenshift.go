package mejoinedteamscheduleopenshift

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleOpenShiftId{}

// MeJoinedTeamScheduleOpenShiftId is a struct representing the Resource ID for a Me Joined Team Schedule Open Shift
type MeJoinedTeamScheduleOpenShiftId struct {
	TeamId      string
	OpenShiftId string
}

// NewMeJoinedTeamScheduleOpenShiftID returns a new MeJoinedTeamScheduleOpenShiftId struct
func NewMeJoinedTeamScheduleOpenShiftID(teamId string, openShiftId string) MeJoinedTeamScheduleOpenShiftId {
	return MeJoinedTeamScheduleOpenShiftId{
		TeamId:      teamId,
		OpenShiftId: openShiftId,
	}
}

// ParseMeJoinedTeamScheduleOpenShiftID parses 'input' into a MeJoinedTeamScheduleOpenShiftId
func ParseMeJoinedTeamScheduleOpenShiftID(input string) (*MeJoinedTeamScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleOpenShiftId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OpenShiftId, ok = parsed.Parsed["openShiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamScheduleOpenShiftIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamScheduleOpenShiftId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamScheduleOpenShiftIDInsensitively(input string) (*MeJoinedTeamScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleOpenShiftId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OpenShiftId, ok = parsed.Parsed["openShiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamScheduleOpenShiftID checks that 'input' can be parsed as a Me Joined Team Schedule Open Shift ID
func ValidateMeJoinedTeamScheduleOpenShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamScheduleOpenShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Schedule Open Shift ID
func (id MeJoinedTeamScheduleOpenShiftId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/openShifts/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.OpenShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Schedule Open Shift ID
func (id MeJoinedTeamScheduleOpenShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOpenShifts", "openShifts", "openShifts"),
		resourceids.UserSpecifiedSegment("openShiftId", "openShiftIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Schedule Open Shift ID
func (id MeJoinedTeamScheduleOpenShiftId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Open Shift: %q", id.OpenShiftId),
	}
	return fmt.Sprintf("Me Joined Team Schedule Open Shift (%s)", strings.Join(components, "\n"))
}
