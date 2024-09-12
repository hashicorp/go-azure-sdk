package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarViewIdAttachmentId{}

// UserIdCalendarViewIdAttachmentId is a struct representing the Resource ID for a User Id Calendar View Id Attachment
type UserIdCalendarViewIdAttachmentId struct {
	UserId       string
	EventId      string
	AttachmentId string
}

// NewUserIdCalendarViewIdAttachmentID returns a new UserIdCalendarViewIdAttachmentId struct
func NewUserIdCalendarViewIdAttachmentID(userId string, eventId string, attachmentId string) UserIdCalendarViewIdAttachmentId {
	return UserIdCalendarViewIdAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseUserIdCalendarViewIdAttachmentID parses 'input' into a UserIdCalendarViewIdAttachmentId
func ParseUserIdCalendarViewIdAttachmentID(input string) (*UserIdCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarViewIdAttachmentIDInsensitively parses 'input' case-insensitively into a UserIdCalendarViewIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarViewIdAttachmentIDInsensitively(input string) (*UserIdCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarViewIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdCalendarViewIdAttachmentID checks that 'input' can be parsed as a User Id Calendar View Id Attachment ID
func ValidateUserIdCalendarViewIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarViewIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar View Id Attachment ID
func (id UserIdCalendarViewIdAttachmentId) ID() string {
	fmtString := "/users/%s/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar View Id Attachment ID
func (id UserIdCalendarViewIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar View Id Attachment ID
func (id UserIdCalendarViewIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Id Calendar View Id Attachment (%s)", strings.Join(components, "\n"))
}
