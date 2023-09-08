package mechatinstalledappteamsappdefinition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatInstalledAppId{}

// MeChatInstalledAppId is a struct representing the Resource ID for a Me Chat Installed App
type MeChatInstalledAppId struct {
	ChatId                 string
	TeamsAppInstallationId string
}

// NewMeChatInstalledAppID returns a new MeChatInstalledAppId struct
func NewMeChatInstalledAppID(chatId string, teamsAppInstallationId string) MeChatInstalledAppId {
	return MeChatInstalledAppId{
		ChatId:                 chatId,
		TeamsAppInstallationId: teamsAppInstallationId,
	}
}

// ParseMeChatInstalledAppID parses 'input' into a MeChatInstalledAppId
func ParseMeChatInstalledAppID(input string) (*MeChatInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatInstalledAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatInstalledAppId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ParseMeChatInstalledAppIDInsensitively parses 'input' case-insensitively into a MeChatInstalledAppId
// note: this method should only be used for API response data and not user input
func ParseMeChatInstalledAppIDInsensitively(input string) (*MeChatInstalledAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatInstalledAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatInstalledAppId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsAppInstallationId, ok = parsed.Parsed["teamsAppInstallationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAppInstallationId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatInstalledAppID checks that 'input' can be parsed as a Me Chat Installed App ID
func ValidateMeChatInstalledAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatInstalledAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Installed App ID
func (id MeChatInstalledAppId) ID() string {
	fmtString := "/me/chats/%s/installedApps/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.TeamsAppInstallationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Installed App ID
func (id MeChatInstalledAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticInstalledApps", "installedApps", "installedApps"),
		resourceids.UserSpecifiedSegment("teamsAppInstallationId", "teamsAppInstallationIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Installed App ID
func (id MeChatInstalledAppId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams App Installation: %q", id.TeamsAppInstallationId),
	}
	return fmt.Sprintf("Me Chat Installed App (%s)", strings.Join(components, "\n"))
}
