package usercalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarViewAttachmentId{}

// UserCalendarViewAttachmentId is a struct representing the Resource ID for a User Calendar View Attachment
type UserCalendarViewAttachmentId struct {
	UserId       string
	EventId      string
	AttachmentId string
}

// NewUserCalendarViewAttachmentID returns a new UserCalendarViewAttachmentId struct
func NewUserCalendarViewAttachmentID(userId string, eventId string, attachmentId string) UserCalendarViewAttachmentId {
	return UserCalendarViewAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseUserCalendarViewAttachmentID parses 'input' into a UserCalendarViewAttachmentId
func ParseUserCalendarViewAttachmentID(input string) (*UserCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewAttachmentId{}

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

// ParseUserCalendarViewAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarViewAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarViewAttachmentIDInsensitively(input string) (*UserCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewAttachmentId{}

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

// ValidateUserCalendarViewAttachmentID checks that 'input' can be parsed as a User Calendar View Attachment ID
func ValidateUserCalendarViewAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarViewAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar View Attachment ID
func (id UserCalendarViewAttachmentId) ID() string {
	fmtString := "/users/%s/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar View Attachment ID
func (id UserCalendarViewAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar View Attachment ID
func (id UserCalendarViewAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar View Attachment (%s)", strings.Join(components, "\n"))
}
