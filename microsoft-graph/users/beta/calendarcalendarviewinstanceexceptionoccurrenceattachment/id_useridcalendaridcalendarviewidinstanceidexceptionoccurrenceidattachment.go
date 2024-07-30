package calendarcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{}

// UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId is a struct representing the Resource ID for a User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Attachment
type UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId struct {
	UserId       string
	CalendarId   string
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID returns a new UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId struct
func NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(userId string, calendarId string, eventId string, eventId1 string, eventId2 string, attachmentId string) UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId {
	return UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{
		UserId:       userId,
		CalendarId:   calendarId,
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID parses 'input' into a UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId
func ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(input string) (*UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentIDInsensitively(input string) (*UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
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

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID checks that 'input' can be parsed as a User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Attachment ID
func ValidateUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarView/%s/instances/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Attachment (%s)", strings.Join(components, "\n"))
}
