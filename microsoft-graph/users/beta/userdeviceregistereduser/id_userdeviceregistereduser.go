package userdeviceregistereduser

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceRegisteredUserId{}

// UserDeviceRegisteredUserId is a struct representing the Resource ID for a User Device Registered User
type UserDeviceRegisteredUserId struct {
	UserId            string
	DeviceId          string
	DirectoryObjectId string
}

// NewUserDeviceRegisteredUserID returns a new UserDeviceRegisteredUserId struct
func NewUserDeviceRegisteredUserID(userId string, deviceId string, directoryObjectId string) UserDeviceRegisteredUserId {
	return UserDeviceRegisteredUserId{
		UserId:            userId,
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserDeviceRegisteredUserID parses 'input' into a UserDeviceRegisteredUserId
func ParseUserDeviceRegisteredUserID(input string) (*UserDeviceRegisteredUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceRegisteredUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceRegisteredUserId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceRegisteredUserIDInsensitively parses 'input' case-insensitively into a UserDeviceRegisteredUserId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceRegisteredUserIDInsensitively(input string) (*UserDeviceRegisteredUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceRegisteredUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceRegisteredUserId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceRegisteredUserID checks that 'input' can be parsed as a User Device Registered User ID
func ValidateUserDeviceRegisteredUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceRegisteredUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Registered User ID
func (id UserDeviceRegisteredUserId) ID() string {
	fmtString := "/users/%s/devices/%s/registeredUsers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Registered User ID
func (id UserDeviceRegisteredUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticRegisteredUsers", "registeredUsers", "registeredUsers"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Device Registered User ID
func (id UserDeviceRegisteredUserId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Device Registered User (%s)", strings.Join(components, "\n"))
}
