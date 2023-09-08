package useronlinemeetingregistrationcustomquestion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOnlineMeetingRegistrationCustomQuestionId{}

// UserOnlineMeetingRegistrationCustomQuestionId is a struct representing the Resource ID for a User Online Meeting Registration Custom Question
type UserOnlineMeetingRegistrationCustomQuestionId struct {
	UserId                        string
	OnlineMeetingId               string
	MeetingRegistrationQuestionId string
}

// NewUserOnlineMeetingRegistrationCustomQuestionID returns a new UserOnlineMeetingRegistrationCustomQuestionId struct
func NewUserOnlineMeetingRegistrationCustomQuestionID(userId string, onlineMeetingId string, meetingRegistrationQuestionId string) UserOnlineMeetingRegistrationCustomQuestionId {
	return UserOnlineMeetingRegistrationCustomQuestionId{
		UserId:                        userId,
		OnlineMeetingId:               onlineMeetingId,
		MeetingRegistrationQuestionId: meetingRegistrationQuestionId,
	}
}

// ParseUserOnlineMeetingRegistrationCustomQuestionID parses 'input' into a UserOnlineMeetingRegistrationCustomQuestionId
func ParseUserOnlineMeetingRegistrationCustomQuestionID(input string) (*UserOnlineMeetingRegistrationCustomQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingRegistrationCustomQuestionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingRegistrationCustomQuestionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingRegistrationQuestionId, ok = parsed.Parsed["meetingRegistrationQuestionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrationQuestionId", *parsed)
	}

	return &id, nil
}

// ParseUserOnlineMeetingRegistrationCustomQuestionIDInsensitively parses 'input' case-insensitively into a UserOnlineMeetingRegistrationCustomQuestionId
// note: this method should only be used for API response data and not user input
func ParseUserOnlineMeetingRegistrationCustomQuestionIDInsensitively(input string) (*UserOnlineMeetingRegistrationCustomQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOnlineMeetingRegistrationCustomQuestionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOnlineMeetingRegistrationCustomQuestionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingRegistrationQuestionId, ok = parsed.Parsed["meetingRegistrationQuestionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrationQuestionId", *parsed)
	}

	return &id, nil
}

// ValidateUserOnlineMeetingRegistrationCustomQuestionID checks that 'input' can be parsed as a User Online Meeting Registration Custom Question ID
func ValidateUserOnlineMeetingRegistrationCustomQuestionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOnlineMeetingRegistrationCustomQuestionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Online Meeting Registration Custom Question ID
func (id UserOnlineMeetingRegistrationCustomQuestionId) ID() string {
	fmtString := "/users/%s/onlineMeetings/%s/registration/customQuestions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnlineMeetingId, id.MeetingRegistrationQuestionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Online Meeting Registration Custom Question ID
func (id UserOnlineMeetingRegistrationCustomQuestionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticRegistration", "registration", "registration"),
		resourceids.StaticSegment("staticCustomQuestions", "customQuestions", "customQuestions"),
		resourceids.UserSpecifiedSegment("meetingRegistrationQuestionId", "meetingRegistrationQuestionIdValue"),
	}
}

// String returns a human-readable description of this User Online Meeting Registration Custom Question ID
func (id UserOnlineMeetingRegistrationCustomQuestionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Registration Question: %q", id.MeetingRegistrationQuestionId),
	}
	return fmt.Sprintf("User Online Meeting Registration Custom Question (%s)", strings.Join(components, "\n"))
}
