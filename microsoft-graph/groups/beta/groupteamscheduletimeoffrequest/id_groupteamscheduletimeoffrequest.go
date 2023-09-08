package groupteamscheduletimeoffrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleTimeOffRequestId{}

// GroupTeamScheduleTimeOffRequestId is a struct representing the Resource ID for a Group Team Schedule Time Off Request
type GroupTeamScheduleTimeOffRequestId struct {
	GroupId          string
	TimeOffRequestId string
}

// NewGroupTeamScheduleTimeOffRequestID returns a new GroupTeamScheduleTimeOffRequestId struct
func NewGroupTeamScheduleTimeOffRequestID(groupId string, timeOffRequestId string) GroupTeamScheduleTimeOffRequestId {
	return GroupTeamScheduleTimeOffRequestId{
		GroupId:          groupId,
		TimeOffRequestId: timeOffRequestId,
	}
}

// ParseGroupTeamScheduleTimeOffRequestID parses 'input' into a GroupTeamScheduleTimeOffRequestId
func ParseGroupTeamScheduleTimeOffRequestID(input string) (*GroupTeamScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleTimeOffRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TimeOffRequestId, ok = parsed.Parsed["timeOffRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleTimeOffRequestIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleTimeOffRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleTimeOffRequestIDInsensitively(input string) (*GroupTeamScheduleTimeOffRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleTimeOffRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleTimeOffRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TimeOffRequestId, ok = parsed.Parsed["timeOffRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffRequestId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleTimeOffRequestID checks that 'input' can be parsed as a Group Team Schedule Time Off Request ID
func ValidateGroupTeamScheduleTimeOffRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleTimeOffRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Time Off Request ID
func (id GroupTeamScheduleTimeOffRequestId) ID() string {
	fmtString := "/groups/%s/team/schedule/timeOffRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TimeOffRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Time Off Request ID
func (id GroupTeamScheduleTimeOffRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimeOffRequests", "timeOffRequests", "timeOffRequests"),
		resourceids.UserSpecifiedSegment("timeOffRequestId", "timeOffRequestIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Time Off Request ID
func (id GroupTeamScheduleTimeOffRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Time Off Request: %q", id.TimeOffRequestId),
	}
	return fmt.Sprintf("Group Team Schedule Time Off Request (%s)", strings.Join(components, "\n"))
}
