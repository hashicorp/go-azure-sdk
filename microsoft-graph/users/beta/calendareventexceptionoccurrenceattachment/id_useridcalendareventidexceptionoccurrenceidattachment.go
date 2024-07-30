package calendareventexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId{}

// UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId is a struct representing the Resource ID for a User Id Calendar Event Id Exception Occurrence Id Attachment
type UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId struct {
	UserId       string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewUserIdCalendarEventIdExceptionOccurrenceIdAttachmentID returns a new UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId struct
func NewUserIdCalendarEventIdExceptionOccurrenceIdAttachmentID(userId string, eventId string, eventId1 string, attachmentId string) UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId {
	return UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdCalendarEventIdExceptionOccurrenceIdAttachmentID parses 'input' into a UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId
func ParseUserIdCalendarEventIdExceptionOccurrenceIdAttachmentID(input string) (*UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarEventIdExceptionOccurrenceIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarEventIdExceptionOccurrenceIdAttachmentIDInsensitively(input string) (*UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
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

// ValidateUserIdCalendarEventIdExceptionOccurrenceIdAttachmentID checks that 'input' can be parsed as a User Id Calendar Event Id Exception Occurrence Id Attachment ID
func ValidateUserIdCalendarEventIdExceptionOccurrenceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarEventIdExceptionOccurrenceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Event Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId) ID() string {
	fmtString := "/users/%s/calendar/events/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Event Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Event Id Exception Occurrence Id Attachment ID
func (id UserIdCalendarEventIdExceptionOccurrenceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Calendar Event Id Exception Occurrence Id Attachment (%s)", strings.Join(components, "\n"))
}
