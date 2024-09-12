package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId{}

// UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId is a struct representing the Resource ID for a User Id Calendar Group Id Calendar Id Event Id Exception Occurrence Id Attachment
type UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	AttachmentId    string
}

// NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID returns a new UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId struct
func NewUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID(userId string, calendarGroupId string, calendarId string, eventId string, eventId1 string, attachmentId string) UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId {
	return UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		AttachmentId:    attachmentId,
	}
}

// ParseUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID parses 'input' into a UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId
func ParseUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID(input string) (*UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentIDInsensitively(input string) (*UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarGroupId, ok = input.Parsed["calendarGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID checks that 'input' can be parsed as a User Id Calendar Group Id Calendar Id Event Id Exception Occurrence Id Attachment ID
func ValidateUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Group Id Calendar Id Event Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/events/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Group Id Calendar Id Event Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Group Id Calendar Id Event Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarGroupIdCalendarIdEventIdExceptionOccurrenceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Calendar Group Id Calendar Id Event Id Exception Occurrence Id Attachment (%s)", strings.Join(components, "\n"))
}
