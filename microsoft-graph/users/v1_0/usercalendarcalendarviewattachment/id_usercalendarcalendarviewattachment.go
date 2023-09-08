package usercalendarcalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarViewAttachmentId{}

// UserCalendarCalendarViewAttachmentId is a struct representing the Resource ID for a User Calendar Calendar View Attachment
type UserCalendarCalendarViewAttachmentId struct {
	UserId       string
	EventId      string
	AttachmentId string
}

// NewUserCalendarCalendarViewAttachmentID returns a new UserCalendarCalendarViewAttachmentId struct
func NewUserCalendarCalendarViewAttachmentID(userId string, eventId string, attachmentId string) UserCalendarCalendarViewAttachmentId {
	return UserCalendarCalendarViewAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseUserCalendarCalendarViewAttachmentID parses 'input' into a UserCalendarCalendarViewAttachmentId
func ParseUserCalendarCalendarViewAttachmentID(input string) (*UserCalendarCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarCalendarViewAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarViewAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarViewAttachmentIDInsensitively(input string) (*UserCalendarCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarCalendarViewAttachmentID checks that 'input' can be parsed as a User Calendar Calendar View Attachment ID
func ValidateUserCalendarCalendarViewAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarViewAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar View Attachment ID
func (id UserCalendarCalendarViewAttachmentId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar View Attachment ID
func (id UserCalendarCalendarViewAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Calendar View Attachment ID
func (id UserCalendarCalendarViewAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar Calendar View Attachment (%s)", strings.Join(components, "\n"))
}
