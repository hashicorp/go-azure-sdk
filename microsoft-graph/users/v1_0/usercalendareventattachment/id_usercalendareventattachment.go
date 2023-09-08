package usercalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventAttachmentId{}

// UserCalendarEventAttachmentId is a struct representing the Resource ID for a User Calendar Event Attachment
type UserCalendarEventAttachmentId struct {
	UserId       string
	CalendarId   string
	EventId      string
	AttachmentId string
}

// NewUserCalendarEventAttachmentID returns a new UserCalendarEventAttachmentId struct
func NewUserCalendarEventAttachmentID(userId string, calendarId string, eventId string, attachmentId string) UserCalendarEventAttachmentId {
	return UserCalendarEventAttachmentId{
		UserId:       userId,
		CalendarId:   calendarId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseUserCalendarEventAttachmentID parses 'input' into a UserCalendarEventAttachmentId
func ParseUserCalendarEventAttachmentID(input string) (*UserCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarEventAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarEventAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventAttachmentIDInsensitively(input string) (*UserCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarEventAttachmentID checks that 'input' can be parsed as a User Calendar Event Attachment ID
func ValidateUserCalendarEventAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event Attachment ID
func (id UserCalendarEventAttachmentId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event Attachment ID
func (id UserCalendarEventAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Event Attachment ID
func (id UserCalendarEventAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar Event Attachment (%s)", strings.Join(components, "\n"))
}
