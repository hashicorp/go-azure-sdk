package userchatinstalledappteamsapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatInstalledAppId{}

// UserChatInstalledAppId is a struct representing the Resource ID for a User Chat Installed App
type UserChatInstalledAppId struct {
	UserId                 string
	ChatId                 string
	TeamsAppInstallationId string
}

// NewUserChatInstalledAppID returns a new UserChatInstalledAppId struct
func NewUserChatInstalledAppID(userId string, chatId string, teamsAppInstallationId string) UserChatInstalledAppId {
	return UserChatInstalledAppId{
		UserId:                 userId,
		ChatId:                 chatId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseUserChatInstalledAppID parses 'input' into a UserChatInstalledAppId
func ParseUserChatInstalledAppID(input string) (*UserChatInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatInstalledAppId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ParseUserChatInstalledAppIDInsensitively parses 'input' case-insensitively into a UserChatInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseUserChatInstalledAppIDInsensitively(input string) (*UserChatInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatInstalledAppId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ValidateUserChatInstalledAppID checks that 'input' can be parsed as a User Chat Installed App ID
func ValidateUserChatInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Installed App ID
func (id UserChatInstalledAppId) ID() string {
	fmtString := "/users/%s/chats/%s/installedApps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Installed App ID
func (id UserChatInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticInstalledApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationIdValue"),
	}
}

// String returns a human-readable description of this User Chat Installed App ID
func (id UserChatInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("User Chat Installed App (%s)", strings.Join(components, "\n"))
}
