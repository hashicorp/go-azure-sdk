package usercalendargroupcalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarEventAttachmentId{}

// UserCalendarGroupCalendarEventAttachmentId is a struct representing the Resource ID for a User Calendar Group Calendar Event Attachment
type UserCalendarGroupCalendarEventAttachmentId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
	AttachmentId    string
}

// NewUserCalendarGroupCalendarEventAttachmentID returns a new UserCalendarGroupCalendarEventAttachmentId struct
func NewUserCalendarGroupCalendarEventAttachmentID(userId string, calendarGroupId string, calendarId string, eventId string, attachmentId string) UserCalendarGroupCalendarEventAttachmentId {
	return UserCalendarGroupCalendarEventAttachmentId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		AttachmentId:    attachmentId,
	}
}

// ParseUserCalendarGroupCalendarEventAttachmentID parses 'input' into a UserCalendarGroupCalendarEventAttachmentId
func ParseUserCalendarGroupCalendarEventAttachmentID(input string) (*UserCalendarGroupCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarEventAttachmentId{}

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

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarGroupCalendarEventAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupCalendarEventAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupCalendarEventAttachmentIDInsensitively(input string) (*UserCalendarGroupCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarEventAttachmentId{}

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

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarGroupCalendarEventAttachmentID checks that 'input' can be parsed as a User Calendar Group Calendar Event Attachment ID
func ValidateUserCalendarGroupCalendarEventAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupCalendarEventAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group Calendar Event Attachment ID
func (id UserCalendarGroupCalendarEventAttachmentId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group Calendar Event Attachment ID
func (id UserCalendarGroupCalendarEventAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Group Calendar Event Attachment ID
func (id UserCalendarGroupCalendarEventAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar Group Calendar Event Attachment (%s)", strings.Join(components, "\n"))
}
