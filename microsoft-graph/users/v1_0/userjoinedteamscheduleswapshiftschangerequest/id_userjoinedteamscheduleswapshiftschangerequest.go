package userjoinedteamscheduleswapshiftschangerequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleSwapShiftsChangeRequestId{}

// UserJoinedTeamScheduleSwapShiftsChangeRequestId is a struct representing the Resource ID for a User Joined Team Schedule Swap Shifts Change Request
type UserJoinedTeamScheduleSwapShiftsChangeRequestId struct {
	UserId                    string
	TeamId                    string
	SwapShiftsChangeRequestId string
}

// NewUserJoinedTeamScheduleSwapShiftsChangeRequestID returns a new UserJoinedTeamScheduleSwapShiftsChangeRequestId struct
func NewUserJoinedTeamScheduleSwapShiftsChangeRequestID(userId string, teamId string, swapShiftsChangeRequestId string) UserJoinedTeamScheduleSwapShiftsChangeRequestId {
	return UserJoinedTeamScheduleSwapShiftsChangeRequestId{
		UserId:                    userId,
		TeamId:                    teamId,
		SwapShiftsChangeRequestId: swapShiftsChangeRequestId,
	}
}

// ParseUserJoinedTeamScheduleSwapShiftsChangeRequestID parses 'input' into a UserJoinedTeamScheduleSwapShiftsChangeRequestId
func ParseUserJoinedTeamScheduleSwapShiftsChangeRequestID(input string) (*UserJoinedTeamScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleSwapShiftsChangeRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SwapShiftsChangeRequestId, ok = parsed.Parsed["swapShiftsChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleSwapShiftsChangeRequestIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleSwapShiftsChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleSwapShiftsChangeRequestIDInsensitively(input string) (*UserJoinedTeamScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleSwapShiftsChangeRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.SwapShiftsChangeRequestId, ok = parsed.Parsed["swapShiftsChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleSwapShiftsChangeRequestID checks that 'input' can be parsed as a User Joined Team Schedule Swap Shifts Change Request ID
func ValidateUserJoinedTeamScheduleSwapShiftsChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleSwapShiftsChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Swap Shifts Change Request ID
func (id UserJoinedTeamScheduleSwapShiftsChangeRequestId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/swapShiftsChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.SwapShiftsChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Swap Shifts Change Request ID
func (id UserJoinedTeamScheduleSwapShiftsChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticSwapShiftsChangeRequests", "swapShiftsChangeRequests", "swapShiftsChangeRequests"),
		resourceids.UserSpecifiedSegment("swapShiftsChangeRequestId", "swapShiftsChangeRequestIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Swap Shifts Change Request ID
func (id UserJoinedTeamScheduleSwapShiftsChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Swap Shifts Change Request: %q", id.SwapShiftsChangeRequestId),
	}
	return fmt.Sprintf("User Joined Team Schedule Swap Shifts Change Request (%s)", strings.Join(components, "\n"))
}
