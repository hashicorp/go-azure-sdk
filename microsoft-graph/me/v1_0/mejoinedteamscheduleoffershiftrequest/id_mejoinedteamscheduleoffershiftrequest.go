package mejoinedteamscheduleoffershiftrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeJoinedTeamScheduleOfferShiftRequestId{}

// MeJoinedTeamScheduleOfferShiftRequestId is a struct representing the Resource ID for a Me Joined Team Schedule Offer Shift Request
type MeJoinedTeamScheduleOfferShiftRequestId struct {
	TeamId              string
	OfferShiftRequestId string
}

// NewMeJoinedTeamScheduleOfferShiftRequestID returns a new MeJoinedTeamScheduleOfferShiftRequestId struct
func NewMeJoinedTeamScheduleOfferShiftRequestID(teamId string, offerShiftRequestId string) MeJoinedTeamScheduleOfferShiftRequestId {
	return MeJoinedTeamScheduleOfferShiftRequestId{
		TeamId:              teamId,
		OfferShiftRequestId: offerShiftRequestId,
	}
}

// ParseMeJoinedTeamScheduleOfferShiftRequestID parses 'input' into a MeJoinedTeamScheduleOfferShiftRequestId
func ParseMeJoinedTeamScheduleOfferShiftRequestID(input string) (*MeJoinedTeamScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleOfferShiftRequestId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OfferShiftRequestId, ok = parsed.Parsed["offerShiftRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", *parsed)
	}

	return &id, nil
}

// ParseMeJoinedTeamScheduleOfferShiftRequestIDInsensitively parses 'input' case-insensitively into a MeJoinedTeamScheduleOfferShiftRequestId
// note: this method should only be used for API response data and not user input
func ParseMeJoinedTeamScheduleOfferShiftRequestIDInsensitively(input string) (*MeJoinedTeamScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeJoinedTeamScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeJoinedTeamScheduleOfferShiftRequestId{}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OfferShiftRequestId, ok = parsed.Parsed["offerShiftRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", *parsed)
	}

	return &id, nil
}

// ValidateMeJoinedTeamScheduleOfferShiftRequestID checks that 'input' can be parsed as a Me Joined Team Schedule Offer Shift Request ID
func ValidateMeJoinedTeamScheduleOfferShiftRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeJoinedTeamScheduleOfferShiftRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Joined Team Schedule Offer Shift Request ID
func (id MeJoinedTeamScheduleOfferShiftRequestId) ID() string {
	fmtString := "/me/joinedTeams/%s/schedule/offerShiftRequests/%s"
	return fmt.Sprintf(fmtString, id.TeamId, id.OfferShiftRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Joined Team Schedule Offer Shift Request ID
func (id MeJoinedTeamScheduleOfferShiftRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOfferShiftRequests", "offerShiftRequests", "offerShiftRequests"),
		resourceids.UserSpecifiedSegment("offerShiftRequestId", "offerShiftRequestIdValue"),
	}
}

// String returns a human-readable description of this Me Joined Team Schedule Offer Shift Request ID
func (id MeJoinedTeamScheduleOfferShiftRequestId) String() string {
	components := []string{
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Offer Shift Request: %q", id.OfferShiftRequestId),
	}
	return fmt.Sprintf("Me Joined Team Schedule Offer Shift Request (%s)", strings.Join(components, "\n"))
}
