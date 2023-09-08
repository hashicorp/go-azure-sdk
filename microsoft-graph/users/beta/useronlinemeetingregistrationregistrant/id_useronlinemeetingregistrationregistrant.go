package useronlinemeetingregistrationregistrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingRegistrationRegistrantId{}

// UserOnlineMeetingRegistrationRegistrantId is a struct representing the Resource ID for a User Online Meeting Registration Registrant
type UserOnlineMeetingRegistrationRegistrantId struct {
	UserId                  string
	OnlineMeetingId         string
	MeetingRegistrantBaseId string
}

// NewUserOnlineMeetingRegistrationRegistrantID returns a new UserOnlineMeetingRegistrationRegistrantId struct
func NewUserOnlineMeetingRegistrationRegistrantID(userId string, onlineMeetingId string, meetingRegistrantBaseId string) UserOnlineMeetingRegistrationRegistrantId {
	return UserOnlineMeetingRegistrationRegistrantId{
		UserId:                  userId,
		OnlineMeetingId:         onlineMeetingId,
		MeetingRegistrantBaseId: meetingRegistrantBaseId,
	}
}

// ParseUserOnlineMeetingRegistrationRegistrantID parses 'input' into a UserOnlineMeetingRegistrationRegistrantId
func ParseUserOnlineMeetingRegistrationRegistrantID(input string) (*UserOnlineMeetingRegistrationRegistrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingRegistrationRegistrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingRegistrationRegistrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingRegistrantBaseId, ok = parsed.Parsed["meetingRegistrantBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrantBaseId", *parsed)
	}

	return &id, nil
}

// ParseUserOnlineMeetingRegistrationRegistrantIDInsensitively parses 'input' case-insensitively into a UserOnlineMeetingRegistrationRegistrantId
// note: this method should only be used for API response data and not user input
func ParseUserOnlineMeetingRegistrationRegistrantIDInsensitively(input string) (*UserOnlineMeetingRegistrationRegistrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingRegistrationRegistrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingRegistrationRegistrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingRegistrantBaseId, ok = parsed.Parsed["meetingRegistrantBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrantBaseId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnlineMeetingRegistrationRegistrantID checks that 'input' can be parsed as a User Online Meeting Registration Registrant ID
func ValidateUserOnlineMeetingRegistrationRegistrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnlineMeetingRegistrationRegistrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Online Meeting Registration Registrant ID
func (id UserOnlineMeetingRegistrationRegistrantId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/registration/registrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.MeetingRegistrantBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Online Meeting Registration Registrant ID
func (id UserOnlineMeetingRegistrationRegistrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticRegistration", "registration", "registration"),
		resourceids.StaticSegment("staticRegistrants", "registrants", "registrants"),
		resourceids.UserSpecifiedSegment("meetingRegistrantBaseId", "meetingRegistrantBaseIdValue"),
	}
}

// String returns a human-readable description of this User Online Meeting Registration Registrant ID
func (id UserOnlineMeetingRegistrationRegistrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Registrant Base: %q", id.MeetingRegistrantBaseId),
	}
	return fmt.Sprintf("User Online Meeting Registration Registrant (%s)", strings.Join(components, "\n"))
}
