package groupteamscheduleopenshift

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleOpenShiftId{}

// GroupTeamScheduleOpenShiftId is a struct representing the Resource ID for a Group Team Schedule Open Shift
type GroupTeamScheduleOpenShiftId struct {
	GroupId     string
	OpenShiftId string
}

// NewGroupTeamScheduleOpenShiftID returns a new GroupTeamScheduleOpenShiftId struct
func NewGroupTeamScheduleOpenShiftID(groupId string, openShiftId string) GroupTeamScheduleOpenShiftId {
	return GroupTeamScheduleOpenShiftId{
		GroupId:     groupId,
		OpenShiftId: openShiftId,
	}
}

// ParseGroupTeamScheduleOpenShiftID parses 'input' into a GroupTeamScheduleOpenShiftId
func ParseGroupTeamScheduleOpenShiftID(input string) (*GroupTeamScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleOpenShiftId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OpenShiftId, ok = parsed.Parsed["openShiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleOpenShiftIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleOpenShiftId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleOpenShiftIDInsensitively(input string) (*GroupTeamScheduleOpenShiftId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleOpenShiftId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleOpenShiftId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OpenShiftId, ok = parsed.Parsed["openShiftId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleOpenShiftID checks that 'input' can be parsed as a Group Team Schedule Open Shift ID
func ValidateGroupTeamScheduleOpenShiftID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleOpenShiftID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Open Shift ID
func (id GroupTeamScheduleOpenShiftId) ID() string {
	fmtString := "/groups/%s/team/schedule/openShifts/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OpenShiftId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Open Shift ID
func (id GroupTeamScheduleOpenShiftId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOpenShifts", "openShifts", "openShifts"),
		resourceids.UserSpecifiedSegment("openShiftId", "openShiftIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Open Shift ID
func (id GroupTeamScheduleOpenShiftId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Open Shift: %q", id.OpenShiftId),
	}
	return fmt.Sprintf("Group Team Schedule Open Shift (%s)", strings.Join(components, "\n"))
}
