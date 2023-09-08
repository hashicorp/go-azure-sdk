package usermanageddevicedeviceconfigurationstate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserManagedDeviceDeviceConfigurationStateId{}

// UserManagedDeviceDeviceConfigurationStateId is a struct representing the Resource ID for a User Managed Device Device Configuration State
type UserManagedDeviceDeviceConfigurationStateId struct {
	UserId                     string
	ManagedDeviceId            string
	DeviceConfigurationStateId string
}

// NewUserManagedDeviceDeviceConfigurationStateID returns a new UserManagedDeviceDeviceConfigurationStateId struct
func NewUserManagedDeviceDeviceConfigurationStateID(userId string, managedDeviceId string, deviceConfigurationStateId string) UserManagedDeviceDeviceConfigurationStateId {
	return UserManagedDeviceDeviceConfigurationStateId{
		UserId:                     userId,
		ManagedDeviceId:            managedDeviceId,
		DeviceConfigurationStateId: deviceConfigurationStateId,
	}
}

// ParseUserManagedDeviceDeviceConfigurationStateID parses 'input' into a UserManagedDeviceDeviceConfigurationStateId
func ParseUserManagedDeviceDeviceConfigurationStateID(input string) (*UserManagedDeviceDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceDeviceConfigurationStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceConfigurationStateId, ok = parsed.Parsed["deviceConfigurationStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationStateId", *parsed)
	}

	return &id, nil
}

// ParseUserManagedDeviceDeviceConfigurationStateIDInsensitively parses 'input' case-insensitively into a UserManagedDeviceDeviceConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseUserManagedDeviceDeviceConfigurationStateIDInsensitively(input string) (*UserManagedDeviceDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserManagedDeviceDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserManagedDeviceDeviceConfigurationStateId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceConfigurationStateId, ok = parsed.Parsed["deviceConfigurationStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationStateId", *parsed)
	}

	return &id, nil
}

// ValidateUserManagedDeviceDeviceConfigurationStateID checks that 'input' can be parsed as a User Managed Device Device Configuration State ID
func ValidateUserManagedDeviceDeviceConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserManagedDeviceDeviceConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Managed Device Device Configuration State ID
func (id UserManagedDeviceDeviceConfigurationStateId) ID() string {
	fmtString := "/users/%s/managedDevices/%s/deviceConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedDeviceId, id.DeviceConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Managed Device Device Configuration State ID
func (id UserManagedDeviceDeviceConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticDeviceConfigurationStates", "deviceConfigurationStates", "deviceConfigurationStates"),
		resourceids.UserSpecifiedSegment("deviceConfigurationStateId", "deviceConfigurationStateIdValue"),
	}
}

// String returns a human-readable description of this User Managed Device Device Configuration State ID
func (id UserManagedDeviceDeviceConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Configuration State: %q", id.DeviceConfigurationStateId),
	}
	return fmt.Sprintf("User Managed Device Device Configuration State (%s)", strings.Join(components, "\n"))
}
