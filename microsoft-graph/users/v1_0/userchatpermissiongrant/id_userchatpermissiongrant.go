package userchatpermissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatPermissionGrantId{}

// UserChatPermissionGrantId is a struct representing the Resource ID for a User Chat Permission Grant
type UserChatPermissionGrantId struct {
	UserId                            string
	ChatId                            string
	ResourceSpecificPermissionGrantId string
}

// NewUserChatPermissionGrantID returns a new UserChatPermissionGrantId struct
func NewUserChatPermissionGrantID(userId string, chatId string, resourceSpecificPermissionGrantId string) UserChatPermissionGrantId {
	return UserChatPermissionGrantId{
		UserId:                            userId,
		ChatId:                            chatId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseUserChatPermissionGrantID parses 'input' into a UserChatPermissionGrantId
func ParseUserChatPermissionGrantID(input string) (*UserChatPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatPermissionGrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseUserChatPermissionGrantIDInsensitively parses 'input' case-insensitively into a UserChatPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserChatPermissionGrantIDInsensitively(input string) (*UserChatPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatPermissionGrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateUserChatPermissionGrantID checks that 'input' can be parsed as a User Chat Permission Grant ID
func ValidateUserChatPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Permission Grant ID
func (id UserChatPermissionGrantId) ID() string {
	fmtString := "/users/%s/chats/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Permission Grant ID
func (id UserChatPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticPermissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this User Chat Permission Grant ID
func (id UserChatPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("User Chat Permission Grant (%s)", strings.Join(components, "\n"))
}
