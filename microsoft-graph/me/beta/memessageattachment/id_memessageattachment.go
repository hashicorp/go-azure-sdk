package memessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMessageAttachmentId{}

// MeMessageAttachmentId is a struct representing the Resource ID for a Me Message Attachment
type MeMessageAttachmentId struct {
	MessageId    string
	AttachmentId string
}

// NewMeMessageAttachmentID returns a new MeMessageAttachmentId struct
func NewMeMessageAttachmentID(messageId string, attachmentId string) MeMessageAttachmentId {
	return MeMessageAttachmentId{
		MessageId:    messageId,
		AttachmentId: attachmentId,
	}
}

// ParseMeMessageAttachmentID parses 'input' into a MeMessageAttachmentId
func ParseMeMessageAttachmentID(input string) (*MeMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMessageAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMessageAttachmentId{}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeMessageAttachmentIDInsensitively parses 'input' case-insensitively into a MeMessageAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeMessageAttachmentIDInsensitively(input string) (*MeMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMessageAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMessageAttachmentId{}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeMessageAttachmentID checks that 'input' can be parsed as a Me Message Attachment ID
func ValidateMeMessageAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMessageAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Message Attachment ID
func (id MeMessageAttachmentId) ID() string {
	fmtString := "/me/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Message Attachment ID
func (id MeMessageAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Message Attachment ID
func (id MeMessageAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Message Attachment (%s)", strings.Join(components, "\n"))
}
