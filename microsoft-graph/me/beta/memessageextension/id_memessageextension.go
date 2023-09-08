package memessageextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMessageExtensionId{}

// MeMessageExtensionId is a struct representing the Resource ID for a Me Message Extension
type MeMessageExtensionId struct {
	MessageId   string
	ExtensionId string
}

// NewMeMessageExtensionID returns a new MeMessageExtensionId struct
func NewMeMessageExtensionID(messageId string, extensionId string) MeMessageExtensionId {
	return MeMessageExtensionId{
		MessageId:   messageId,
		ExtensionId: extensionId,
	}
}

// ParseMeMessageExtensionID parses 'input' into a MeMessageExtensionId
func ParseMeMessageExtensionID(input string) (*MeMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMessageExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMessageExtensionId{}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeMessageExtensionIDInsensitively parses 'input' case-insensitively into a MeMessageExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeMessageExtensionIDInsensitively(input string) (*MeMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeMessageExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeMessageExtensionId{}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeMessageExtensionID checks that 'input' can be parsed as a Me Message Extension ID
func ValidateMeMessageExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeMessageExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Message Extension ID
func (id MeMessageExtensionId) ID() string {
	fmtString := "/me/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Message Extension ID
func (id MeMessageExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Message Extension ID
func (id MeMessageExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Message Extension (%s)", strings.Join(components, "\n"))
}
