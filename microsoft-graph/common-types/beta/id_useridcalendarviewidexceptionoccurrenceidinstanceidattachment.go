package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{}

// UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId is a struct representing the Resource ID for a User Id Calendar View Id Exception Occurrence Id Instance Id Attachment
type UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId struct {
	UserId       string
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID returns a new UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId struct
func NewUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID(userId string, eventId string, eventId1 string, eventId2 string, attachmentId string) UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId {
	return UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID parses 'input' into a UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId
func ParseUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID(input string) (*UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentIDInsensitively(input string) (*UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID checks that 'input' can be parsed as a User Id Calendar View Id Exception Occurrence Id Instance Id Attachment ID
func ValidateUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar View Id Exception Occurrence Id Instance Id Attachment ID
func (id UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId) ID() string {
	fmtString := "/users/%s/calendarView/%s/exceptionOccurrences/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar View Id Exception Occurrence Id Instance Id Attachment ID
func (id UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar View Id Exception Occurrence Id Instance Id Attachment ID
func (id UserIdCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Calendar View Id Exception Occurrence Id Instance Id Attachment (%s)", strings.Join(components, "\n"))
}
