package userdevicemanagementtroubleshootingevent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserDeviceManagementTroubleshootingEventId{}

// UserDeviceManagementTroubleshootingEventId is a struct representing the Resource ID for a User Device Management Troubleshooting Event
type UserDeviceManagementTroubleshootingEventId struct {
	UserId                                 string
	DeviceManagementTroubleshootingEventId string
}

// NewUserDeviceManagementTroubleshootingEventID returns a new UserDeviceManagementTroubleshootingEventId struct
func NewUserDeviceManagementTroubleshootingEventID(userId string, deviceManagementTroubleshootingEventId string) UserDeviceManagementTroubleshootingEventId {
	return UserDeviceManagementTroubleshootingEventId{
		UserId:                                 userId,
		DeviceManagementTroubleshootingEventId: deviceManagementTroubleshootingEventId,
	}
}

// ParseUserDeviceManagementTroubleshootingEventID parses 'input' into a UserDeviceManagementTroubleshootingEventId
func ParseUserDeviceManagementTroubleshootingEventID(input string) (*UserDeviceManagementTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceManagementTroubleshootingEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceManagementTroubleshootingEventId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceManagementTroubleshootingEventId, ok = parsed.Parsed["deviceManagementTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTroubleshootingEventId", *parsed)
	}

	return &id, nil
}

// ParseUserDeviceManagementTroubleshootingEventIDInsensitively parses 'input' case-insensitively into a UserDeviceManagementTroubleshootingEventId
// note: this method should only be used for API response data and not user input
func ParseUserDeviceManagementTroubleshootingEventIDInsensitively(input string) (*UserDeviceManagementTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserDeviceManagementTroubleshootingEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserDeviceManagementTroubleshootingEventId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DeviceManagementTroubleshootingEventId, ok = parsed.Parsed["deviceManagementTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTroubleshootingEventId", *parsed)
	}

	return &id, nil
}

// ValidateUserDeviceManagementTroubleshootingEventID checks that 'input' can be parsed as a User Device Management Troubleshooting Event ID
func ValidateUserDeviceManagementTroubleshootingEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserDeviceManagementTroubleshootingEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Device Management Troubleshooting Event ID
func (id UserDeviceManagementTroubleshootingEventId) ID() string {
	fmtString := "/users/%s/deviceManagementTroubleshootingEvents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DeviceManagementTroubleshootingEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Device Management Troubleshooting Event ID
func (id UserDeviceManagementTroubleshootingEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticDeviceManagementTroubleshootingEvents", "deviceManagementTroubleshootingEvents", "deviceManagementTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("deviceManagementTroubleshootingEventId", "deviceManagementTroubleshootingEventIdValue"),
	}
}

// String returns a human-readable description of this User Device Management Troubleshooting Event ID
func (id UserDeviceManagementTroubleshootingEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Device Management Troubleshooting Event: %q", id.DeviceManagementTroubleshootingEventId),
	}
	return fmt.Sprintf("User Device Management Troubleshooting Event (%s)", strings.Join(components, "\n"))
}
