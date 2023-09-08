package groupeventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupEventAttachmentId{}

// GroupEventAttachmentId is a struct representing the Resource ID for a Group Event Attachment
type GroupEventAttachmentId struct {
	GroupId      string
	EventId      string
	AttachmentId string
}

// NewGroupEventAttachmentID returns a new GroupEventAttachmentId struct
func NewGroupEventAttachmentID(groupId string, eventId string, attachmentId string) GroupEventAttachmentId {
	return GroupEventAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseGroupEventAttachmentID parses 'input' into a GroupEventAttachmentId
func ParseGroupEventAttachmentID(input string) (*GroupEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseGroupEventAttachmentIDInsensitively parses 'input' case-insensitively into a GroupEventAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupEventAttachmentIDInsensitively(input string) (*GroupEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupEventAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupEventAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateGroupEventAttachmentID checks that 'input' can be parsed as a Group Event Attachment ID
func ValidateGroupEventAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupEventAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Event Attachment ID
func (id GroupEventAttachmentId) ID() string {
	fmtString := "/groups/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Event Attachment ID
func (id GroupEventAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Event Attachment ID
func (id GroupEventAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Event Attachment (%s)", strings.Join(components, "\n"))
}
