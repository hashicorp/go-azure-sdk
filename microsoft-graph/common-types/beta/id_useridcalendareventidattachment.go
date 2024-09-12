package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarEventIdAttachmentId{}

// UserIdCalendarEventIdAttachmentId is a struct representing the Resource ID for a User Id Calendar Event Id Attachment
type UserIdCalendarEventIdAttachmentId struct {
	UserId       string
	EventId      string
	AttachmentId string
}

// NewUserIdCalendarEventIdAttachmentID returns a new UserIdCalendarEventIdAttachmentId struct
func NewUserIdCalendarEventIdAttachmentID(userId string, eventId string, attachmentId string) UserIdCalendarEventIdAttachmentId {
	return UserIdCalendarEventIdAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdCalendarEventIdAttachmentID parses 'input' into a UserIdCalendarEventIdAttachmentId
func ParseUserIdCalendarEventIdAttachmentID(input string) (*UserIdCalendarEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarEventIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdCalendarEventIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarEventIdAttachmentIDInsensitively(input string) (*UserIdCalendarEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarEventIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateUserIdCalendarEventIdAttachmentID checks that 'input' can be parsed as a User Id Calendar Event Id Attachment ID
func ValidateUserIdCalendarEventIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarEventIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Event Id Attachment ID
func (id UserIdCalendarEventIdAttachmentId) ID() string {
	fmtString := "/users/%s/calendar/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Event Id Attachment ID
func (id UserIdCalendarEventIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Event Id Attachment ID
func (id UserIdCalendarEventIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Calendar Event Id Attachment (%s)", strings.Join(components, "\n"))
}
