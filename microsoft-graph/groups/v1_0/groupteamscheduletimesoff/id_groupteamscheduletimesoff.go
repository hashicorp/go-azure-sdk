package groupteamscheduletimesoff

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleTimesOffId{}

// GroupTeamScheduleTimesOffId is a struct representing the Resource ID for a Group Team Schedule Times Off
type GroupTeamScheduleTimesOffId struct {
	GroupId   string
	TimeOffId string
}

// NewGroupTeamScheduleTimesOffID returns a new GroupTeamScheduleTimesOffId struct
func NewGroupTeamScheduleTimesOffID(groupId string, timeOffId string) GroupTeamScheduleTimesOffId {
	return GroupTeamScheduleTimesOffId{
		GroupId:   groupId,
		TimeOffId: timeOffId,
	}
}

// ParseGroupTeamScheduleTimesOffID parses 'input' into a GroupTeamScheduleTimesOffId
func ParseGroupTeamScheduleTimesOffID(input string) (*GroupTeamScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleTimesOffId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleTimesOffId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TimeOffId, ok = parsed.Parsed["timeOffId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleTimesOffIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleTimesOffId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleTimesOffIDInsensitively(input string) (*GroupTeamScheduleTimesOffId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleTimesOffId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleTimesOffId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.TimeOffId, ok = parsed.Parsed["timeOffId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "timeOffId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleTimesOffID checks that 'input' can be parsed as a Group Team Schedule Times Off ID
func ValidateGroupTeamScheduleTimesOffID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleTimesOffID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Times Off ID
func (id GroupTeamScheduleTimesOffId) ID() string {
	fmtString := "/groups/%s/team/schedule/timesOff/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.TimeOffId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Times Off ID
func (id GroupTeamScheduleTimesOffId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticTimesOff", "timesOff", "timesOff"),
		resourceids.UserSpecifiedSegment("timeOffId", "timeOffIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Times Off ID
func (id GroupTeamScheduleTimesOffId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Time Off: %q", id.TimeOffId),
	}
	return fmt.Sprintf("Group Team Schedule Times Off (%s)", strings.Join(components, "\n"))
}
