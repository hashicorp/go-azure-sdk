package meonlinemeetingregistrationcustomquestion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingRegistrationCustomQuestionId{}

// MeOnlineMeetingRegistrationCustomQuestionId is a struct representing the Resource ID for a Me Online Meeting Registration Custom Question
type MeOnlineMeetingRegistrationCustomQuestionId struct {
	OnlineMeetingId               string
	MeetingRegistrationQuestionId string
}

// NewMeOnlineMeetingRegistrationCustomQuestionID returns a new MeOnlineMeetingRegistrationCustomQuestionId struct
func NewMeOnlineMeetingRegistrationCustomQuestionID(onlineMeetingId string, meetingRegistrationQuestionId string) MeOnlineMeetingRegistrationCustomQuestionId {
	return MeOnlineMeetingRegistrationCustomQuestionId{
		OnlineMeetingId:               onlineMeetingId,
		MeetingRegistrationQuestionId: meetingRegistrationQuestionId,
	}
}

// ParseMeOnlineMeetingRegistrationCustomQuestionID parses 'input' into a MeOnlineMeetingRegistrationCustomQuestionId
func ParseMeOnlineMeetingRegistrationCustomQuestionID(input string) (*MeOnlineMeetingRegistrationCustomQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingRegistrationCustomQuestionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingRegistrationCustomQuestionId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingRegistrationQuestionId, ok = parsed.Parsed["meetingRegistrationQuestionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrationQuestionId", *parsed)
	}

	return &id, nil
}

// ParseMeOnlineMeetingRegistrationCustomQuestionIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingRegistrationCustomQuestionId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingRegistrationCustomQuestionIDInsensitively(input string) (*MeOnlineMeetingRegistrationCustomQuestionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingRegistrationCustomQuestionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingRegistrationCustomQuestionId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingRegistrationQuestionId, ok = parsed.Parsed["meetingRegistrationQuestionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrationQuestionId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnlineMeetingRegistrationCustomQuestionID checks that 'input' can be parsed as a Me Online Meeting Registration Custom Question ID
func ValidateMeOnlineMeetingRegistrationCustomQuestionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingRegistrationCustomQuestionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Registration Custom Question ID
func (id MeOnlineMeetingRegistrationCustomQuestionId) ID() string {
	fmtString := "/me/onlineMeetings/%s/registration/customQuestions/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.MeetingRegistrationQuestionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Registration Custom Question ID
func (id MeOnlineMeetingRegistrationCustomQuestionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticRegistration", "registration", "registration"),
		resourceids.StaticSegment("staticCustomQuestions", "customQuestions", "customQuestions"),
		resourceids.UserSpecifiedSegment("meetingRegistrationQuestionId", "meetingRegistrationQuestionIdValue"),
	}
}

// String returns a human-readable description of this Me Online Meeting Registration Custom Question ID
func (id MeOnlineMeetingRegistrationCustomQuestionId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Registration Question: %q", id.MeetingRegistrationQuestionId),
	}
	return fmt.Sprintf("Me Online Meeting Registration Custom Question (%s)", strings.Join(components, "\n"))
}
