package userdeviceregisteredowner

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceRegisteredOwnerId{}

// UserDeviceRegisteredOwnerId is a struct representing the Resource ID for a User Device Registered Owner
type UserDeviceRegisteredOwnerId struct {
	UserId            string
	DeviceId          string
	DirectoryObjectId string
}

// NewUserDeviceRegisteredOwnerID returns a new UserDeviceRegisteredOwnerId struct
func NewUserDeviceRegisteredOwnerID(userId string, deviceId string, directoryObjectId string) UserDeviceRegisteredOwnerId {
	return UserDeviceRegisteredOwnerId{
		UserId:            userId,
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserDeviceRegisteredOwnerID parses 'input' into a UserDeviceRegisteredOwnerId
func ParseUserDeviceRegisteredOwnerID(input string) (*UserDeviceRegisteredOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceRegisteredOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceRegisteredOwnerId{}

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

// ParseUserDeviceRegisteredOwnerIDInsensitively parses 'input' case-insensitively into a UserDeviceRegisteredOwnerId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceRegisteredOwnerIDInsensitively(input string) (*UserDeviceRegisteredOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceRegisteredOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceRegisteredOwnerId{}

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

// ValidateUserDeviceRegisteredOwnerID checks that 'input' can be parsed as a User Device Registered Owner ID
func ValidateUserDeviceRegisteredOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceRegisteredOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Registered Owner ID
func (id UserDeviceRegisteredOwnerId) ID() string {
	fmtString := "/users/%s/devices/%s/registeredOwners/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Registered Owner ID
func (id UserDeviceRegisteredOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticRegisteredOwners", "registeredOwners", "registeredOwners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Device Registered Owner ID
func (id UserDeviceRegisteredOwnerId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Device Registered Owner (%s)", strings.Join(components, "\n"))
}
