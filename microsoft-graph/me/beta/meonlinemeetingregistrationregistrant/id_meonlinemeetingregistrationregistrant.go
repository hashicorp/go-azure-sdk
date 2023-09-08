package meonlinemeetingregistrationregistrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOnlineMeetingRegistrationRegistrantId{}

// MeOnlineMeetingRegistrationRegistrantId is a struct representing the Resource ID for a Me Online Meeting Registration Registrant
type MeOnlineMeetingRegistrationRegistrantId struct {
	OnlineMeetingId         string
	MeetingRegistrantBaseId string
}

// NewMeOnlineMeetingRegistrationRegistrantID returns a new MeOnlineMeetingRegistrationRegistrantId struct
func NewMeOnlineMeetingRegistrationRegistrantID(onlineMeetingId string, meetingRegistrantBaseId string) MeOnlineMeetingRegistrationRegistrantId {
	return MeOnlineMeetingRegistrationRegistrantId{
		OnlineMeetingId:         onlineMeetingId,
		MeetingRegistrantBaseId: meetingRegistrantBaseId,
	}
}

// ParseMeOnlineMeetingRegistrationRegistrantID parses 'input' into a MeOnlineMeetingRegistrationRegistrantId
func ParseMeOnlineMeetingRegistrationRegistrantID(input string) (*MeOnlineMeetingRegistrationRegistrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingRegistrationRegistrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingRegistrationRegistrantId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingRegistrantBaseId, ok = parsed.Parsed["meetingRegistrantBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrantBaseId", *parsed)
	}

	return &id, nil
}

// ParseMeOnlineMeetingRegistrationRegistrantIDInsensitively parses 'input' case-insensitively into a MeOnlineMeetingRegistrationRegistrantId
// note: this method should only be used for API response data and not user input
func ParseMeOnlineMeetingRegistrationRegistrantIDInsensitively(input string) (*MeOnlineMeetingRegistrationRegistrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOnlineMeetingRegistrationRegistrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOnlineMeetingRegistrationRegistrantId{}

	if id.OnlineMeetingId, ok = parsed.Parsed["onlineMeetingId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "onlineMeetingId", *parsed)
	}

	if id.MeetingRegistrantBaseId, ok = parsed.Parsed["meetingRegistrantBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "meetingRegistrantBaseId", *parsed)
	}

	return &id, nil
}

// ValidateMeOnlineMeetingRegistrationRegistrantID checks that 'input' can be parsed as a Me Online Meeting Registration Registrant ID
func ValidateMeOnlineMeetingRegistrationRegistrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOnlineMeetingRegistrationRegistrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Online Meeting Registration Registrant ID
func (id MeOnlineMeetingRegistrationRegistrantId) ID() string {
	fmtString := "/me/onlineMeetings/%s/registration/registrants/%s"
	return fmt.Sprintf(fmtString, id.OnlineMeetingId, id.MeetingRegistrantBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Online Meeting Registration Registrant ID
func (id MeOnlineMeetingRegistrationRegistrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOnlineMeetings", "onlineMeetings", "onlineMeetings"),
		resourceids.UserSpecifiedSegment("onlineMeetingId", "onlineMeetingIdValue"),
		resourceids.StaticSegment("staticRegistration", "registration", "registration"),
		resourceids.StaticSegment("staticRegistrants", "registrants", "registrants"),
		resourceids.UserSpecifiedSegment("meetingRegistrantBaseId", "meetingRegistrantBaseIdValue"),
	}
}

// String returns a human-readable description of this Me Online Meeting Registration Registrant ID
func (id MeOnlineMeetingRegistrationRegistrantId) String() string {
	components := []string{
		fmt.Sprintf("Online Meeting: %q", id.OnlineMeetingId),
		fmt.Sprintf("Meeting Registrant Base: %q", id.MeetingRegistrantBaseId),
	}
	return fmt.Sprintf("Me Online Meeting Registration Registrant (%s)", strings.Join(components, "\n"))
}
