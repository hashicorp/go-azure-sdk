package usercalendargroupcalendarcalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId{}

// UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId is a struct representing the Resource ID for a User Calendar Group Calendar Calendar View Exception Occurrence Instance Attachment
type UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	EventId2        string
	AttachmentId    string
}

// NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID returns a new UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId struct
func NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID(userId string, calendarGroupId string, calendarId string, eventId string, eventId1 string, eventId2 string, attachmentId string) UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId {
	return UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		EventId2:        eventId2,
		AttachmentId:    attachmentId,
	}
}

// ParseUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID parses 'input' into a UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId
func ParseUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID(input string) (*UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentIDInsensitively(input string) (*UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID checks that 'input' can be parsed as a User Calendar Group Calendar Calendar View Exception Occurrence Instance Attachment ID
func ValidateUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group Calendar Calendar View Exception Occurrence Instance Attachment ID
func (id UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/calendarView/%s/exceptionOccurrences/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group Calendar Calendar View Exception Occurrence Instance Attachment ID
func (id UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Group Calendar Calendar View Exception Occurrence Instance Attachment ID
func (id UserCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar Group Calendar Calendar View Exception Occurrence Instance Attachment (%s)", strings.Join(components, "\n"))
}
