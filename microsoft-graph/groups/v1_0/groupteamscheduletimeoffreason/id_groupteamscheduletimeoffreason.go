package groupteamscheduletimeoffreason

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleTimeOffReasonId{}

// GroupTeamScheduleTimeOffReasonId is a struct representing the Resource ID for a Group Team Schedule Time Off Reason
type GroupTeamScheduleTimeOffReasonId struct {
	GroupId         string
	TimeOffReasonId string
}

// NewGroupTeamScheduleTimeOffReasonID returns a new GroupTeamScheduleTimeOffReasonId struct
func NewGroupTeamScheduleTimeOffReasonID(groupId string, timeOffReasonId string) GroupTeamScheduleTimeOffReasonId {
	return GroupTeamScheduleTimeOffReasonId{
		GroupId:         groupId,
		TimeOffReasonId: timeOffReasonId,
	}
}

// ParseGroupTeamScheduleTimeOffReasonID parses 'input' into a GroupTeamScheduleTimeOffReasonId
func ParseGroupTeamScheduleTimeOffReasonID(input string) (*GroupTeamScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleTimeOffReasonId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TimeOffReasonId, ok = parsed.Parsed["timeOffReasonId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffReasonId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleTimeOffReasonIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleTimeOffReasonId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleTimeOffReasonIDInsensitively(input string) (*GroupTeamScheduleTimeOffReasonId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleTimeOffReasonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleTimeOffReasonId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TimeOffReasonId, ok = parsed.Parsed["timeOffReasonId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffReasonId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleTimeOffReasonID checks that 'input' can be parsed as a Group Team Schedule Time Off Reason ID
func ValidateGroupTeamScheduleTimeOffReasonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleTimeOffReasonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Time Off Reason ID
func (id GroupTeamScheduleTimeOffReasonId) ID() string {
	fmtString := "/groups/%s/team/schedule/timeOffReasons/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TimeOffReasonId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Time Off Reason ID
func (id GroupTeamScheduleTimeOffReasonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimeOffReasons", "timeOffReasons", "timeOffReasons"),
		resourceids.UserSpecifiedSegment("timeOffReasonId", "timeOffReasonIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Time Off Reason ID
func (id GroupTeamScheduleTimeOffReasonId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Time Off Reason: %q", id.TimeOffReasonId),
	}
	return fmt.Sprintf("Group Team Schedule Time Off Reason (%s)", strings.Join(components, "\n"))
}
