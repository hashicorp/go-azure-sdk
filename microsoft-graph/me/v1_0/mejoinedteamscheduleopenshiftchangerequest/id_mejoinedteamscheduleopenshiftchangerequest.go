package mejoinedteamscheduleopenshiftchangerequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleOpenShiftChangeRequestId{}

// MeJoinedTeamScheduleOpenShiftChangeRequestId is a struct representing the Resource ID for a Me Joined Team Schedule Open Shift Change Request
type MeJoinedTeamScheduleOpenShiftChangeRequestId struct {
	TeamId                   string
	OpenShiftChangeRequestId string
}

// NewMeJoinedTeamScheduleOpenShiftChangeRequestID returns a new MeJoinedTeamScheduleOpenShiftChangeRequestId struct
func NewMeJoinedTeamScheduleOpenShiftChangeRequestID(teamId string, openShiftChangeRequestId string) MeJoinedTeamScheduleOpenShiftChangeRequestId {
	return MeJoinedTeamScheduleOpenShiftChangeRequestId{
		TeamId:                   teamId,
		OpenShiftChangeRequestId: openShiftChangeRequestId,
	}
}

// ParseMeJoinedTeamScheduleOpenShiftChangeRequestID parses 'input' into a MeJoinedTeamScheduleOpenShiftChangeRequestId
func ParseMeJoinedTeamScheduleOpenShiftChangeRequestID(input string) (*MeJoinedTeamScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleOpenShiftChangeRequestId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OpenShiftChangeRequestId, ok = parsed.Parsed["openShiftChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamScheduleOpenShiftChangeRequestIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamScheduleOpenShiftChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamScheduleOpenShiftChangeRequestIDInsensitively(input string) (*MeJoinedTeamScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleOpenShiftChangeRequestId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OpenShiftChangeRequestId, ok = parsed.Parsed["openShiftChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamScheduleOpenShiftChangeRequestID checks that 'input' can be parsed as a Me Joined Team Schedule Open Shift Change Request ID
func ValidateMeJoinedTeamScheduleOpenShiftChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamScheduleOpenShiftChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Schedule Open Shift Change Request ID
func (id MeJoinedTeamScheduleOpenShiftChangeRequestId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/openShiftChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.OpenShiftChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Schedule Open Shift Change Request ID
func (id MeJoinedTeamScheduleOpenShiftChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOpenShiftChangeRequests", "openShiftChangeRequests", "openShiftChangeRequests"),
		resourceids.UserSpecifiedSegment("openShiftChangeRequestId", "openShiftChangeRequestIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Schedule Open Shift Change Request ID
func (id MeJoinedTeamScheduleOpenShiftChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Open Shift Change Request: %q", id.OpenShiftChangeRequestId),
	}
	return fmt.Sprintf("Me Joined Team Schedule Open Shift Change Request (%s)", strings.Join(components, "\n"))
}
