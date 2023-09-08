package meeventinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventInstanceExceptionOccurrenceAttachmentId{}

// MeEventInstanceExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a Me Event Instance Exception Occurrence Attachment
type MeEventInstanceExceptionOccurrenceAttachmentId struct {
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewMeEventInstanceExceptionOccurrenceAttachmentID returns a new MeEventInstanceExceptionOccurrenceAttachmentId struct
func NewMeEventInstanceExceptionOccurrenceAttachmentID(eventId string, eventId1 string, eventId2 string, attachmentId string) MeEventInstanceExceptionOccurrenceAttachmentId {
	return MeEventInstanceExceptionOccurrenceAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseMeEventInstanceExceptionOccurrenceAttachmentID parses 'input' into a MeEventInstanceExceptionOccurrenceAttachmentId
func ParseMeEventInstanceExceptionOccurrenceAttachmentID(input string) (*MeEventInstanceExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceExceptionOccurrenceAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeEventInstanceExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a MeEventInstanceExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeEventInstanceExceptionOccurrenceAttachmentIDInsensitively(input string) (*MeEventInstanceExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceExceptionOccurrenceAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeEventInstanceExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a Me Event Instance Exception Occurrence Attachment ID
func ValidateMeEventInstanceExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventInstanceExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Instance Exception Occurrence Attachment ID
func (id MeEventInstanceExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/me/events/%s/instances/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Instance Exception Occurrence Attachment ID
func (id MeEventInstanceExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Event Instance Exception Occurrence Attachment ID
func (id MeEventInstanceExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Event Instance Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
