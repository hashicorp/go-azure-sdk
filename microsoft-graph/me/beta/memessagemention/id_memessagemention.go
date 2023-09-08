package memessagemention

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMessageMentionId{}

// MeMessageMentionId is a struct representing the Resource ID for a Me Message Mention
type MeMessageMentionId struct {
	MessageId string
	MentionId string
}

// NewMeMessageMentionID returns a new MeMessageMentionId struct
func NewMeMessageMentionID(messageId string, mentionId string) MeMessageMentionId {
	return MeMessageMentionId{
		MessageId: messageId,
		MentionId: mentionId,
	}
}

// ParseMeMessageMentionID parses 'input' into a MeMessageMentionId
func ParseMeMessageMentionID(input string) (*MeMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMessageMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMessageMentionId{}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ParseMeMessageMentionIDInsensitively parses 'input' case-insensitively into a MeMessageMentionId
// note: this method should only be used for API response data and not user input
func ParseMeMessageMentionIDInsensitively(input string) (*MeMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMessageMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMessageMentionId{}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ValidateMeMessageMentionID checks that 'input' can be parsed as a Me Message Mention ID
func ValidateMeMessageMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMessageMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Message Mention ID
func (id MeMessageMentionId) ID() string {
	fmtString := "/me/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Message Mention ID
func (id MeMessageMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticMentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionIdValue"),
	}
}

// String returns a human-readable description of this Me Message Mention ID
func (id MeMessageMentionId) String() string {
	components := []string{
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("Me Message Mention (%s)", strings.Join(components, "\n"))
}
