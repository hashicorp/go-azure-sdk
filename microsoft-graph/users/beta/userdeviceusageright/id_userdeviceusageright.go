package userdeviceusageright

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceUsageRightId{}

// UserDeviceUsageRightId is a struct representing the Resource ID for a User Device Usage Right
type UserDeviceUsageRightId struct {
	UserId       string
	DeviceId     string
	UsageRightId string
}

// NewUserDeviceUsageRightID returns a new UserDeviceUsageRightId struct
func NewUserDeviceUsageRightID(userId string, deviceId string, usageRightId string) UserDeviceUsageRightId {
	return UserDeviceUsageRightId{
		UserId:       userId,
		DeviceId:     deviceId,
		UsageRightId: usageRightId,
	}
}

// ParseUserDeviceUsageRightID parses 'input' into a UserDeviceUsageRightId
func ParseUserDeviceUsageRightID(input string) (*UserDeviceUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceUsageRightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceUsageRightId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.UsageRightId, ok = parsed.Parsed["usageRightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceUsageRightIDInsensitively parses 'input' case-insensitively into a UserDeviceUsageRightId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceUsageRightIDInsensitively(input string) (*UserDeviceUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceUsageRightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceUsageRightId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.UsageRightId, ok = parsed.Parsed["usageRightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceUsageRightID checks that 'input' can be parsed as a User Device Usage Right ID
func ValidateUserDeviceUsageRightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceUsageRightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Usage Right ID
func (id UserDeviceUsageRightId) ID() string {
	fmtString := "/users/%s/devices/%s/usageRights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.UsageRightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Usage Right ID
func (id UserDeviceUsageRightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticUsageRights", "usageRights", "usageRights"),
		resourceids.UserSpecifiedSegment("usageRightId", "usageRightIdValue"),
	}
}

// String returns a human-readable description of this User Device Usage Right ID
func (id UserDeviceUsageRightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Usage Right: %q", id.UsageRightId),
	}
	return fmt.Sprintf("User Device Usage Right (%s)", strings.Join(components, "\n"))
}
