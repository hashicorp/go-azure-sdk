package groupteamscheduleoffershiftrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupTeamScheduleOfferShiftRequestId{}

// GroupTeamScheduleOfferShiftRequestId is a struct representing the Resource ID for a Group Team Schedule Offer Shift Request
type GroupTeamScheduleOfferShiftRequestId struct {
	GroupId             string
	OfferShiftRequestId string
}

// NewGroupTeamScheduleOfferShiftRequestID returns a new GroupTeamScheduleOfferShiftRequestId struct
func NewGroupTeamScheduleOfferShiftRequestID(groupId string, offerShiftRequestId string) GroupTeamScheduleOfferShiftRequestId {
	return GroupTeamScheduleOfferShiftRequestId{
		GroupId:             groupId,
		OfferShiftRequestId: offerShiftRequestId,
	}
}

// ParseGroupTeamScheduleOfferShiftRequestID parses 'input' into a GroupTeamScheduleOfferShiftRequestId
func ParseGroupTeamScheduleOfferShiftRequestID(input string) (*GroupTeamScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleOfferShiftRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OfferShiftRequestId, ok = parsed.Parsed["offerShiftRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", *parsed)
	}

	return &id, nil
}

// ParseGroupTeamScheduleOfferShiftRequestIDInsensitively parses 'input' case-insensitively into a GroupTeamScheduleOfferShiftRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupTeamScheduleOfferShiftRequestIDInsensitively(input string) (*GroupTeamScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupTeamScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupTeamScheduleOfferShiftRequestId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.OfferShiftRequestId, ok = parsed.Parsed["offerShiftRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", *parsed)
	}

	return &id, nil
}

// ValidateGroupTeamScheduleOfferShiftRequestID checks that 'input' can be parsed as a Group Team Schedule Offer Shift Request ID
func ValidateGroupTeamScheduleOfferShiftRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupTeamScheduleOfferShiftRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Team Schedule Offer Shift Request ID
func (id GroupTeamScheduleOfferShiftRequestId) ID() string {
	fmtString := "/groups/%s/team/schedule/offerShiftRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.OfferShiftRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Team Schedule Offer Shift Request ID
func (id GroupTeamScheduleOfferShiftRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticTeam", "team", "team"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOfferShiftRequests", "offerShiftRequests", "offerShiftRequests"),
		resourceids.UserSpecifiedSegment("offerShiftRequestId", "offerShiftRequestIdValue"),
	}
}

// String returns a human-readable description of this Group Team Schedule Offer Shift Request ID
func (id GroupTeamScheduleOfferShiftRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Offer Shift Request: %q", id.OfferShiftRequestId),
	}
	return fmt.Sprintf("Group Team Schedule Offer Shift Request (%s)", strings.Join(components, "\n"))
}
