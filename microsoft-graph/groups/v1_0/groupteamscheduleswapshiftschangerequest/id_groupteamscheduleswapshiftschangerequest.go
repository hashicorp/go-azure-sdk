package groupteamscheduleswapshiftschangerequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleSwapShiftsChangeRequestId{}

// GroupTeamScheduleSwapShiftsChangeRequestId is a struct representing the Resource ID for a Group Team Schedule Swap Shifts Change Request
type GroupTeamScheduleSwapShiftsChangeRequestId struct {
	GroupId                   string
	SwapShiftsChangeRequestId string
}

// NewGroupTeamScheduleSwapShiftsChangeRequestID returns a new GroupTeamScheduleSwapShiftsChangeRequestId struct
func NewGroupTeamScheduleSwapShiftsChangeRequestID(groupId string, swapShiftsChangeRequestId string) GroupTeamScheduleSwapShiftsChangeRequestId {
	return GroupTeamScheduleSwapShiftsChangeRequestId{
		GroupId:                   groupId,
		SwapShiftsChangeRequestId: swapShiftsChangeRequestId,
	}
}

// ParseGroupTeamScheduleSwapShiftsChangeRequestID parses 'input' into a GroupTeamScheduleSwapShiftsChangeRequestId
func ParseGroupTeamScheduleSwapShiftsChangeRequestID(input string) (*GroupTeamScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleSwapShiftsChangeRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SwapShiftsChangeRequestId, ok = parsed.Parsed["swapShiftsChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleSwapShiftsChangeRequestIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleSwapShiftsChangeRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleSwapShiftsChangeRequestIDInsensitively(input string) (*GroupTeamScheduleSwapShiftsChangeRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleSwapShiftsChangeRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleSwapShiftsChangeRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.SwapShiftsChangeRequestId, ok = parsed.Parsed["swapShiftsChangeRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "swapShiftsChangeRequestId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleSwapShiftsChangeRequestID checks that 'input' can be parsed as a Group Team Schedule Swap Shifts Change Request ID
func ValidateGroupTeamScheduleSwapShiftsChangeRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleSwapShiftsChangeRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Swap Shifts Change Request ID
func (id GroupTeamScheduleSwapShiftsChangeRequestId) ID() string {
	fmtString := "/groups/%s/team/schedule/swapShiftsChangeRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SwapShiftsChangeRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Swap Shifts Change Request ID
func (id GroupTeamScheduleSwapShiftsChangeRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticSwapShiftsChangeRequests", "swapShiftsChangeRequests", "swapShiftsChangeRequests"),
		resourceids.UserSpecifiedSegment("swapShiftsChangeRequestId", "swapShiftsChangeRequestIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Swap Shifts Change Request ID
func (id GroupTeamScheduleSwapShiftsChangeRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Swap Shifts Change Request: %q", id.SwapShiftsChangeRequestId),
	}
	return fmt.Sprintf("Group Team Schedule Swap Shifts Change Request (%s)", strings.Join(components, "\n"))
}
