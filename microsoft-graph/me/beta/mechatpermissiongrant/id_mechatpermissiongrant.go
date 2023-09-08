package mechatpermissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatPermissionGrantId{}

// MeChatPermissionGrantId is a struct representing the Resource ID for a Me Chat Permission Grant
type MeChatPermissionGrantId struct {
	ChatId                            string
	ResourceSpecificPermissionGrantId string
}

// NewMeChatPermissionGrantID returns a new MeChatPermissionGrantId struct
func NewMeChatPermissionGrantID(chatId string, resourceSpecificPermissionGrantId string) MeChatPermissionGrantId {
	return MeChatPermissionGrantId{
		ChatId:                            chatId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseMeChatPermissionGrantID parses 'input' into a MeChatPermissionGrantId
func ParseMeChatPermissionGrantID(input string) (*MeChatPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatPermissionGrantId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseMeChatPermissionGrantIDInsensitively parses 'input' case-insensitively into a MeChatPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseMeChatPermissionGrantIDInsensitively(input string) (*MeChatPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatPermissionGrantId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatPermissionGrantID checks that 'input' can be parsed as a Me Chat Permission Grant ID
func ValidateMeChatPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Permission Grant ID
func (id MeChatPermissionGrantId) ID() string {
	fmtString := "/me/chats/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Permission Grant ID
func (id MeChatPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticPermissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Permission Grant ID
func (id MeChatPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("Me Chat Permission Grant (%s)", strings.Join(components, "\n"))
}
