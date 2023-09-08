package userdevicecommand

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceCommandId{}

// UserDeviceCommandId is a struct representing the Resource ID for a User Device Command
type UserDeviceCommandId struct {
	UserId    string
	DeviceId  string
	CommandId string
}

// NewUserDeviceCommandID returns a new UserDeviceCommandId struct
func NewUserDeviceCommandID(userId string, deviceId string, commandId string) UserDeviceCommandId {
	return UserDeviceCommandId{
		UserId:    userId,
		DeviceId:  deviceId,
		CommandId: commandId,
	}
}

// ParseUserDeviceCommandID parses 'input' into a UserDeviceCommandId
func ParseUserDeviceCommandID(input string) (*UserDeviceCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceCommandId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.CommandId, ok = parsed.Parsed["commandId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "commandId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceCommandIDInsensitively parses 'input' case-insensitively into a UserDeviceCommandId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceCommandIDInsensitively(input string) (*UserDeviceCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceCommandId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.CommandId, ok = parsed.Parsed["commandId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "commandId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceCommandID checks that 'input' can be parsed as a User Device Command ID
func ValidateUserDeviceCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Command ID
func (id UserDeviceCommandId) ID() string {
	fmtString := "/users/%s/devices/%s/commands/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceId, id.CommandId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Command ID
func (id UserDeviceCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticCommands", "commands", "commands"),
		resourceids.UserSpecifiedSegment("commandId", "commandIdValue"),
	}
}

// String returns a human-readable description of this User Device Command ID
func (id UserDeviceCommandId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Command: %q", id.CommandId),
	}
	return fmt.Sprintf("User Device Command (%s)", strings.Join(components, "\n"))
}
