package meeventinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventInstanceAttachmentId{}

// MeEventInstanceAttachmentId is a struct representing the Resource ID for a Me Event Instance Attachment
type MeEventInstanceAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeEventInstanceAttachmentID returns a new MeEventInstanceAttachmentId struct
func NewMeEventInstanceAttachmentID(eventId string, eventId1 string, attachmentId string) MeEventInstanceAttachmentId {
	return MeEventInstanceAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeEventInstanceAttachmentID parses 'input' into a MeEventInstanceAttachmentId
func ParseMeEventInstanceAttachmentID(input string) (*MeEventInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceAttachmentId{}

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

// ParseMeEventInstanceAttachmentIDInsensitively parses 'input' case-insensitively into a MeEventInstanceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeEventInstanceAttachmentIDInsensitively(input string) (*MeEventInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceAttachmentId{}

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

// ValidateMeEventInstanceAttachmentID checks that 'input' can be parsed as a Me Event Instance Attachment ID
func ValidateMeEventInstanceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventInstanceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Instance Attachment ID
func (id MeEventInstanceAttachmentId) ID() string {
	fmtString := "/me/events/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Instance Attachment ID
func (id MeEventInstanceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Event Instance Attachment ID
func (id MeEventInstanceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Event Instance Attachment (%s)", strings.Join(components, "\n"))
}
