package meeventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventAttachmentId{}

// MeEventAttachmentId is a struct representing the Resource ID for a Me Event Attachment
type MeEventAttachmentId struct {
	EventId      string
	AttachmentId string
}

// NewMeEventAttachmentID returns a new MeEventAttachmentId struct
func NewMeEventAttachmentID(eventId string, attachmentId string) MeEventAttachmentId {
	return MeEventAttachmentId{
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeEventAttachmentID parses 'input' into a MeEventAttachmentId
func ParseMeEventAttachmentID(input string) (*MeEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeEventAttachmentIDInsensitively parses 'input' case-insensitively into a MeEventAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeEventAttachmentIDInsensitively(input string) (*MeEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeEventAttachmentID checks that 'input' can be parsed as a Me Event Attachment ID
func ValidateMeEventAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Attachment ID
func (id MeEventAttachmentId) ID() string {
	fmtString := "/me/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Attachment ID
func (id MeEventAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Event Attachment ID
func (id MeEventAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Event Attachment (%s)", strings.Join(components, "\n"))
}
