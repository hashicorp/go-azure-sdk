package calendarcalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdCalendarViewIdAttachmentId{}

// UserIdCalendarIdCalendarViewIdAttachmentId is a struct representing the Resource ID for a User Id Calendar Id Calendar View Id Attachment
type UserIdCalendarIdCalendarViewIdAttachmentId struct {
	UserId       string
	CalendarId   string
	EventId      string
	AttachmentId string
}

// NewUserIdCalendarIdCalendarViewIdAttachmentID returns a new UserIdCalendarIdCalendarViewIdAttachmentId struct
func NewUserIdCalendarIdCalendarViewIdAttachmentID(userId string, calendarId string, eventId string, attachmentId string) UserIdCalendarIdCalendarViewIdAttachmentId {
	return UserIdCalendarIdCalendarViewIdAttachmentId{
		UserId:       userId,
		CalendarId:   calendarId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdCalendarIdCalendarViewIdAttachmentID parses 'input' into a UserIdCalendarIdCalendarViewIdAttachmentId
func ParseUserIdCalendarIdCalendarViewIdAttachmentID(input string) (*UserIdCalendarIdCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdCalendarViewIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdCalendarViewIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdCalendarViewIdAttachmentIDInsensitively(input string) (*UserIdCalendarIdCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdCalendarViewIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdCalendarIdCalendarViewIdAttachmentID checks that 'input' can be parsed as a User Id Calendar Id Calendar View Id Attachment ID
func ValidateUserIdCalendarIdCalendarViewIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdCalendarViewIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Calendar View Id Attachment ID
func (id UserIdCalendarIdCalendarViewIdAttachmentId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Calendar View Id Attachment ID
func (id UserIdCalendarIdCalendarViewIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Calendar View Id Attachment ID
func (id UserIdCalendarIdCalendarViewIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Calendar Id Calendar View Id Attachment (%s)", strings.Join(components, "\n"))
}
