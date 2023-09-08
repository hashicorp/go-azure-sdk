package userjoinedteamscheduleopenshiftchangerequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleOpenShiftChangeRequestId{}

// UserJoinedTeamScheduleOpenShiftChangeRequestId is a struct representing the Resource ID for a User Joined Team Schedule Open Shift Change Request
type UserJoinedTeamScheduleOpenShiftChangeRequestId struct {
	UserId                   string
	TeamId                   string
	OpenShiftChangeRequestId string
}

// NewUserJoinedTeamScheduleOpenShiftChangeRequestID returns a new UserJoinedTeamScheduleOpenShiftChangeRequestId struct
func NewUserJoinedTeamScheduleOpenShiftChangeRequestID(userId string, teamId string, openShiftChangeRequestId string) UserJoinedTeamScheduleOpenShiftChangeRequestId {
	return UserJoinedTeamScheduleOpenShiftChangeRequestId{
		UserId:                   userId,
		TeamId:                   teamId,
		OpenShiftChangeRequestId: openShiftChangeRequestId,
	}
}

// ParseUserJoinedTeamScheduleOpenShiftChangeRequestID parses 'input' into a UserJoinedTeamScheduleOpenShiftChangeRequestId
func ParseUserJoinedTeamScheduleOpenShiftChangeRequestID(input string) (*UserJoinedTeamScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleOpenShiftChangeRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OpenShiftChangeRequestId, ok = parsed.Parsed["openShiftChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleOpenShiftChangeRequestIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleOpenShiftChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleOpenShiftChangeRequestIDInsensitively(input string) (*UserJoinedTeamScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleOpenShiftChangeRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OpenShiftChangeRequestId, ok = parsed.Parsed["openShiftChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleOpenShiftChangeRequestID checks that 'input' can be parsed as a User Joined Team Schedule Open Shift Change Request ID
func ValidateUserJoinedTeamScheduleOpenShiftChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleOpenShiftChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Open Shift Change Request ID
func (id UserJoinedTeamScheduleOpenShiftChangeRequestId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/openShiftChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.OpenShiftChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Open Shift Change Request ID
func (id UserJoinedTeamScheduleOpenShiftChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOpenShiftChangeRequests", "openShiftChangeRequests", "openShiftChangeRequests"),
		resourceids.UserSpecifiedSegment("openShiftChangeRequestId", "openShiftChangeRequestIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Open Shift Change Request ID
func (id UserJoinedTeamScheduleOpenShiftChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Open Shift Change Request: %q", id.OpenShiftChangeRequestId),
	}
	return fmt.Sprintf("User Joined Team Schedule Open Shift Change Request (%s)", strings.Join(components, "\n"))
}
