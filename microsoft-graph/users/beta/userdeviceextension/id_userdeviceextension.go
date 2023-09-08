package userdeviceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceExtensionId{}

// UserDeviceExtensionId is a struct representing the Resource ID for a User Device Extension
type UserDeviceExtensionId struct {
	UserId      string
	DeviceId    string
	ExtensionId string
}

// NewUserDeviceExtensionID returns a new UserDeviceExtensionId struct
func NewUserDeviceExtensionID(userId string, deviceId string, extensionId string) UserDeviceExtensionId {
	return UserDeviceExtensionId{
		UserId:      userId,
		DeviceId:    deviceId,
		ExtensionId: extensionId,
	}
}

// ParseUserDeviceExtensionID parses 'input' into a UserDeviceExtensionId
func ParseUserDeviceExtensionID(input string) (*UserDeviceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceExtensionIDInsensitively parses 'input' case-insensitively into a UserDeviceExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceExtensionIDInsensitively(input string) (*UserDeviceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceExtensionID checks that 'input' can be parsed as a User Device Extension ID
func ValidateUserDeviceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Extension ID
func (id UserDeviceExtensionId) ID() string {
	fmtString := "/users/%s/devices/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Extension ID
func (id UserDeviceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Device Extension ID
func (id UserDeviceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Device Extension (%s)", strings.Join(components, "\n"))
}
