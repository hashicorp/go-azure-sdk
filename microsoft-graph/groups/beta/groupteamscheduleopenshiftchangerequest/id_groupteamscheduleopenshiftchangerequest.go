package groupteamscheduleopenshiftchangerequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleOpenShiftChangeRequestId{}

// GroupTeamScheduleOpenShiftChangeRequestId is a struct representing the Resource ID for a Group Team Schedule Open Shift Change Request
type GroupTeamScheduleOpenShiftChangeRequestId struct {
	GroupId                  string
	OpenShiftChangeRequestId string
}

// NewGroupTeamScheduleOpenShiftChangeRequestID returns a new GroupTeamScheduleOpenShiftChangeRequestId struct
func NewGroupTeamScheduleOpenShiftChangeRequestID(groupId string, openShiftChangeRequestId string) GroupTeamScheduleOpenShiftChangeRequestId {
	return GroupTeamScheduleOpenShiftChangeRequestId{
		GroupId:                  groupId,
		OpenShiftChangeRequestId: openShiftChangeRequestId,
	}
}

// ParseGroupTeamScheduleOpenShiftChangeRequestID parses 'input' into a GroupTeamScheduleOpenShiftChangeRequestId
func ParseGroupTeamScheduleOpenShiftChangeRequestID(input string) (*GroupTeamScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleOpenShiftChangeRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OpenShiftChangeRequestId, ok = parsed.Parsed["openShiftChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleOpenShiftChangeRequestIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleOpenShiftChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleOpenShiftChangeRequestIDInsensitively(input string) (*GroupTeamScheduleOpenShiftChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleOpenShiftChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleOpenShiftChangeRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OpenShiftChangeRequestId, ok = parsed.Parsed["openShiftChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "openShiftChangeRequestId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleOpenShiftChangeRequestID checks that 'input' can be parsed as a Group Team Schedule Open Shift Change Request ID
func ValidateGroupTeamScheduleOpenShiftChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleOpenShiftChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Open Shift Change Request ID
func (id GroupTeamScheduleOpenShiftChangeRequestId) ID() string {
	fmtString := "/groups/%s/team/schedule/openShiftChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OpenShiftChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Open Shift Change Request ID
func (id GroupTeamScheduleOpenShiftChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOpenShiftChangeRequests", "openShiftChangeRequests", "openShiftChangeRequests"),
		resourceids.UserSpecifiedSegment("openShiftChangeRequestId", "openShiftChangeRequestIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Open Shift Change Request ID
func (id GroupTeamScheduleOpenShiftChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Open Shift Change Request: %q", id.OpenShiftChangeRequestId),
	}
	return fmt.Sprintf("Group Team Schedule Open Shift Change Request (%s)", strings.Join(components, "\n"))
}
