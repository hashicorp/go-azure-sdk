package groupteamscheduleschedulinggroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleSchedulingGroupId{}

// GroupTeamScheduleSchedulingGroupId is a struct representing the Resource ID for a Group Team Schedule Scheduling Group
type GroupTeamScheduleSchedulingGroupId struct {
	GroupId           string
	SchedulingGroupId string
}

// NewGroupTeamScheduleSchedulingGroupID returns a new GroupTeamScheduleSchedulingGroupId struct
func NewGroupTeamScheduleSchedulingGroupID(groupId string, schedulingGroupId string) GroupTeamScheduleSchedulingGroupId {
	return GroupTeamScheduleSchedulingGroupId{
		GroupId:           groupId,
		SchedulingGroupId: schedulingGroupId,
	}
}

// ParseGroupTeamScheduleSchedulingGroupID parses 'input' into a GroupTeamScheduleSchedulingGroupId
func ParseGroupTeamScheduleSchedulingGroupID(input string) (*GroupTeamScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleSchedulingGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SchedulingGroupId, ok = parsed.Parsed["schedulingGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleSchedulingGroupIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleSchedulingGroupId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleSchedulingGroupIDInsensitively(input string) (*GroupTeamScheduleSchedulingGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleSchedulingGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleSchedulingGroupId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SchedulingGroupId, ok = parsed.Parsed["schedulingGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "schedulingGroupId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleSchedulingGroupID checks that 'input' can be parsed as a Group Team Schedule Scheduling Group ID
func ValidateGroupTeamScheduleSchedulingGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleSchedulingGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Scheduling Group ID
func (id GroupTeamScheduleSchedulingGroupId) ID() string {
	fmtString := "/groups/%s/team/schedule/schedulingGroups/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SchedulingGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Scheduling Group ID
func (id GroupTeamScheduleSchedulingGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticSchedulingGroups", "schedulingGroups", "schedulingGroups"),
		resourceids.UserSpecifiedSegment("schedulingGroupId", "schedulingGroupIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Scheduling Group ID
func (id GroupTeamScheduleSchedulingGroupId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Scheduling Group: %q", id.SchedulingGroupId),
	}
	return fmt.Sprintf("Group Team Schedule Scheduling Group (%s)", strings.Join(components, "\n"))
}
