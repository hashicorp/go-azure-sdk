package mejoinedteamscheduleswapshiftschangerequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleSwapShiftsChangeRequestId{}

// MeJoinedTeamScheduleSwapShiftsChangeRequestId is a struct representing the Resource ID for a Me Joined Team Schedule Swap Shifts Change Request
type MeJoinedTeamScheduleSwapShiftsChangeRequestId struct {
	TeamId                    string
	SwapShiftsChangeRequestId string
}

// NewMeJoinedTeamScheduleSwapShiftsChangeRequestID returns a new MeJoinedTeamScheduleSwapShiftsChangeRequestId struct
func NewMeJoinedTeamScheduleSwapShiftsChangeRequestID(teamId string, swapShiftsChangeRequestId string) MeJoinedTeamScheduleSwapShiftsChangeRequestId {
	return MeJoinedTeamScheduleSwapShiftsChangeRequestId{
		TeamId:                    teamId,
		SwapShiftsChangeRequestId: swapShiftsChangeRequestId,
	}
}

// ParseMeJoinedTeamScheduleSwapShiftsChangeRequestID parses 'input' into a MeJoinedTeamScheduleSwapShiftsChangeRequestId
func ParseMeJoinedTeamScheduleSwapShiftsChangeRequestID(input string) (*MeJoinedTeamScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleSwapShiftsChangeRequestId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SwapShiftsChangeRequestId, ok = parsed.Parsed["swapShiftsChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamScheduleSwapShiftsChangeRequestIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamScheduleSwapShiftsChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamScheduleSwapShiftsChangeRequestIDInsensitively(input string) (*MeJoinedTeamScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleSwapShiftsChangeRequestId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SwapShiftsChangeRequestId, ok = parsed.Parsed["swapShiftsChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamScheduleSwapShiftsChangeRequestID checks that 'input' can be parsed as a Me Joined Team Schedule Swap Shifts Change Request ID
func ValidateMeJoinedTeamScheduleSwapShiftsChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamScheduleSwapShiftsChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Schedule Swap Shifts Change Request ID
func (id MeJoinedTeamScheduleSwapShiftsChangeRequestId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/swapShiftsChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.SwapShiftsChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Schedule Swap Shifts Change Request ID
func (id MeJoinedTeamScheduleSwapShiftsChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticSwapShiftsChangeRequests", "swapShiftsChangeRequests", "swapShiftsChangeRequests"),
		resourceids.UserSpecifiedSegment("swapShiftsChangeRequestId", "swapShiftsChangeRequestIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Schedule Swap Shifts Change Request ID
func (id MeJoinedTeamScheduleSwapShiftsChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Swap Shifts Change Request: %q", id.SwapShiftsChangeRequestId),
	}
	return fmt.Sprintf("Me Joined Team Schedule Swap Shifts Change Request (%s)", strings.Join(components, "\n"))
}
