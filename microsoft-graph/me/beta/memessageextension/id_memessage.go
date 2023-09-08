package memessageextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMessageId{}

// MeMessageId is a struct representing the Resource ID for a Me Message
type MeMessageId struct {
	MessageId string
}

// NewMeMessageID returns a new MeMessageId struct
func NewMeMessageID(messageId string) MeMessageId {
	return MeMessageId{
		MessageId: messageId,
	}
}

// ParseMeMessageID parses 'input' into a MeMessageId
func ParseMeMessageID(input string) (*MeMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMessageId{}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ParseMeMessageIDInsensitively parses 'input' case-insensitively into a MeMessageId
// note: this method should only be used for API response data and not user input
func ParseMeMessageIDInsensitively(input string) (*MeMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMessageId{}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ValidateMeMessageID checks that 'input' can be parsed as a Me Message ID
func ValidateMeMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Message ID
func (id MeMessageId) ID() string {
	fmtString := "/me/messages/%s"
	return fmt.Sprintf(fmtString, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Message ID
func (id MeMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
	}
}

// String returns a human-readable description of this Me Message ID
func (id MeMessageId) String() string {
	components := []string{
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("Me Message (%s)", strings.Join(components, "\n"))
}
