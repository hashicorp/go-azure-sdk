package usernotification

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserNotificationId{}

// UserNotificationId is a struct representing the Resource ID for a User Notification
type UserNotificationId struct {
	UserId         string
	NotificationId string
}

// NewUserNotificationID returns a new UserNotificationId struct
func NewUserNotificationID(userId string, notificationId string) UserNotificationId {
	return UserNotificationId{
		UserId:         userId,
		NotificationId: notificationId,
	}
}

// ParseUserNotificationID parses 'input' into a UserNotificationId
func ParseUserNotificationID(input string) (*UserNotificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserNotificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserNotificationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotificationId, ok = parsed.Parsed["notificationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notificationId", *parsed)
	}

	return &id, nil
}

// ParseUserNotificationIDInsensitively parses 'input' case-insensitively into a UserNotificationId
// note: this method should only be used for API response data and not user input
func ParseUserNotificationIDInsensitively(input string) (*UserNotificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserNotificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserNotificationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.NotificationId, ok = parsed.Parsed["notificationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "notificationId", *parsed)
	}

	return &id, nil
}

// ValidateUserNotificationID checks that 'input' can be parsed as a User Notification ID
func ValidateUserNotificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserNotificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Notification ID
func (id UserNotificationId) ID() string {
	fmtString := "/users/%s/notifications/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.NotificationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Notification ID
func (id UserNotificationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticNotifications", "notifications", "notifications"),
		resourceids.UserSpecifiedSegment("notificationId", "notificationIdValue"),
	}
}

// String returns a human-readable description of this User Notification ID
func (id UserNotificationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Notification: %q", id.NotificationId),
	}
	return fmt.Sprintf("User Notification (%s)", strings.Join(components, "\n"))
}
