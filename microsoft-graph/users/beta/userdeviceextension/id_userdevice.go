package userdeviceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceId{}

// UserDeviceId is a struct representing the Resource ID for a User Device
type UserDeviceId struct {
	UserId   string
	DeviceId string
}

// NewUserDeviceID returns a new UserDeviceId struct
func NewUserDeviceID(userId string, deviceId string) UserDeviceId {
	return UserDeviceId{
		UserId:   userId,
		DeviceId: deviceId,
	}
}

// ParseUserDeviceID parses 'input' into a UserDeviceId
func ParseUserDeviceID(input string) (*UserDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceIDInsensitively parses 'input' case-insensitively into a UserDeviceId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceIDInsensitively(input string) (*UserDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceID checks that 'input' can be parsed as a User Device ID
func ValidateUserDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device ID
func (id UserDeviceId) ID() string {
	fmtString := "/users/%s/devices/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device ID
func (id UserDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
	}
}

// String returns a human-readable description of this User Device ID
func (id UserDeviceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
	}
	return fmt.Sprintf("User Device (%s)", strings.Join(components, "\n"))
}
