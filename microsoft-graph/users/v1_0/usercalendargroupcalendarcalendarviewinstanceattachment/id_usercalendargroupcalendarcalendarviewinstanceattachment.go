package usercalendargroupcalendarcalendarviewinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarCalendarViewInstanceAttachmentId{}

// UserCalendarGroupCalendarCalendarViewInstanceAttachmentId is a struct representing the Resource ID for a User Calendar Group Calendar Calendar View Instance Attachment
type UserCalendarGroupCalendarCalendarViewInstanceAttachmentId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	AttachmentId    string
}

// NewUserCalendarGroupCalendarCalendarViewInstanceAttachmentID returns a new UserCalendarGroupCalendarCalendarViewInstanceAttachmentId struct
func NewUserCalendarGroupCalendarCalendarViewInstanceAttachmentID(userId string, calendarGroupId string, calendarId string, eventId string, eventId1 string, attachmentId string) UserCalendarGroupCalendarCalendarViewInstanceAttachmentId {
	return UserCalendarGroupCalendarCalendarViewInstanceAttachmentId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		AttachmentId:    attachmentId,
	}
}

// ParseUserCalendarGroupCalendarCalendarViewInstanceAttachmentID parses 'input' into a UserCalendarGroupCalendarCalendarViewInstanceAttachmentId
func ParseUserCalendarGroupCalendarCalendarViewInstanceAttachmentID(input string) (*UserCalendarGroupCalendarCalendarViewInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarCalendarViewInstanceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarCalendarViewInstanceAttachmentId{}

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

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarGroupCalendarCalendarViewInstanceAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupCalendarCalendarViewInstanceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupCalendarCalendarViewInstanceAttachmentIDInsensitively(input string) (*UserCalendarGroupCalendarCalendarViewInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarCalendarViewInstanceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarCalendarViewInstanceAttachmentId{}

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

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarGroupCalendarCalendarViewInstanceAttachmentID checks that 'input' can be parsed as a User Calendar Group Calendar Calendar View Instance Attachment ID
func ValidateUserCalendarGroupCalendarCalendarViewInstanceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupCalendarCalendarViewInstanceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group Calendar Calendar View Instance Attachment ID
func (id UserCalendarGroupCalendarCalendarViewInstanceAttachmentId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/calendarView/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group Calendar Calendar View Instance Attachment ID
func (id UserCalendarGroupCalendarCalendarViewInstanceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Group Calendar Calendar View Instance Attachment ID
func (id UserCalendarGroupCalendarCalendarViewInstanceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar Group Calendar Calendar View Instance Attachment (%s)", strings.Join(components, "\n"))
}
