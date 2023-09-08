package userjoinedteamscheduleoffershiftrequest

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserJoinedTeamScheduleOfferShiftRequestId{}

// UserJoinedTeamScheduleOfferShiftRequestId is a struct representing the Resource ID for a User Joined Team Schedule Offer Shift Request
type UserJoinedTeamScheduleOfferShiftRequestId struct {
	UserId              string
	TeamId              string
	OfferShiftRequestId string
}

// NewUserJoinedTeamScheduleOfferShiftRequestID returns a new UserJoinedTeamScheduleOfferShiftRequestId struct
func NewUserJoinedTeamScheduleOfferShiftRequestID(userId string, teamId string, offerShiftRequestId string) UserJoinedTeamScheduleOfferShiftRequestId {
	return UserJoinedTeamScheduleOfferShiftRequestId{
		UserId:              userId,
		TeamId:              teamId,
		OfferShiftRequestId: offerShiftRequestId,
	}
}

// ParseUserJoinedTeamScheduleOfferShiftRequestID parses 'input' into a UserJoinedTeamScheduleOfferShiftRequestId
func ParseUserJoinedTeamScheduleOfferShiftRequestID(input string) (*UserJoinedTeamScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleOfferShiftRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OfferShiftRequestId, ok = parsed.Parsed["offerShiftRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", *parsed)
	}

	return &id, nil
}

// ParseUserJoinedTeamScheduleOfferShiftRequestIDInsensitively parses 'input' case-insensitively into a UserJoinedTeamScheduleOfferShiftRequestId
// note: this method should only be used for API response data and not user input
func ParseUserJoinedTeamScheduleOfferShiftRequestIDInsensitively(input string) (*UserJoinedTeamScheduleOfferShiftRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserJoinedTeamScheduleOfferShiftRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserJoinedTeamScheduleOfferShiftRequestId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.TeamId, ok = parsed.Parsed["teamId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamId", *parsed)
	}

	if id.OfferShiftRequestId, ok = parsed.Parsed["offerShiftRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "offerShiftRequestId", *parsed)
	}

	return &id, nil
}

// ValidateUserJoinedTeamScheduleOfferShiftRequestID checks that 'input' can be parsed as a User Joined Team Schedule Offer Shift Request ID
func ValidateUserJoinedTeamScheduleOfferShiftRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserJoinedTeamScheduleOfferShiftRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Joined Team Schedule Offer Shift Request ID
func (id UserJoinedTeamScheduleOfferShiftRequestId) ID() string {
	fmtString := "/users/%s/joinedTeams/%s/schedule/offerShiftRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.TeamId, id.OfferShiftRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Joined Team Schedule Offer Shift Request ID
func (id UserJoinedTeamScheduleOfferShiftRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticJoinedTeams", "joinedTeams", "joinedTeams"),
		resourceids.UserSpecifiedSegment("teamId", "teamIdValue"),
		resourceids.StaticSegment("staticSchedule", "schedule", "schedule"),
		resourceids.StaticSegment("staticOfferShiftRequests", "offerShiftRequests", "offerShiftRequests"),
		resourceids.UserSpecifiedSegment("offerShiftRequestId", "offerShiftRequestIdValue"),
	}
}

// String returns a human-readable description of this User Joined Team Schedule Offer Shift Request ID
func (id UserJoinedTeamScheduleOfferShiftRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Team: %q", id.TeamId),
		fmt.Sprintf("Offer Shift Request: %q", id.OfferShiftRequestId),
	}
	return fmt.Sprintf("User Joined Team Schedule Offer Shift Request (%s)", strings.Join(components, "\n"))
}
