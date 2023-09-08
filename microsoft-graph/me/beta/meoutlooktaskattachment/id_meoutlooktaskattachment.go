package meoutlooktaskattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskAttachmentId{}

// MeOutlookTaskAttachmentId is a struct representing the Resource ID for a Me Outlook Task Attachment
type MeOutlookTaskAttachmentId struct {
	OutlookTaskId string
	AttachmentId  string
}

// NewMeOutlookTaskAttachmentID returns a new MeOutlookTaskAttachmentId struct
func NewMeOutlookTaskAttachmentID(outlookTaskId string, attachmentId string) MeOutlookTaskAttachmentId {
	return MeOutlookTaskAttachmentId{
		OutlookTaskId: outlookTaskId,
		AttachmentId:  attachmentId,
	}
}

// ParseMeOutlookTaskAttachmentID parses 'input' into a MeOutlookTaskAttachmentId
func ParseMeOutlookTaskAttachmentID(input string) (*MeOutlookTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskAttachmentId{}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeOutlookTaskAttachmentIDInsensitively parses 'input' case-insensitively into a MeOutlookTaskAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeOutlookTaskAttachmentIDInsensitively(input string) (*MeOutlookTaskAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeOutlookTaskAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeOutlookTaskAttachmentId{}

	if id.OutlookTaskId, ok = parsed.Parsed["outlookTaskId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "outlookTaskId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeOutlookTaskAttachmentID checks that 'input' can be parsed as a Me Outlook Task Attachment ID
func ValidateMeOutlookTaskAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeOutlookTaskAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Outlook Task Attachment ID
func (id MeOutlookTaskAttachmentId) ID() string {
	fmtString := "/me/outlook/tasks/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.OutlookTaskId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Outlook Task Attachment ID
func (id MeOutlookTaskAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticOutlook", "outlook", "outlook"),
		resourceids.StaticSegment("staticTasks", "tasks", "tasks"),
		resourceids.UserSpecifiedSegment("outlookTaskId", "outlookTaskIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Outlook Task Attachment ID
func (id MeOutlookTaskAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Outlook Task: %q", id.OutlookTaskId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Outlook Task Attachment (%s)", strings.Join(components, "\n"))
}
