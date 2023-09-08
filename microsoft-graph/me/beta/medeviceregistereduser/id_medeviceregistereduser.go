package medeviceregistereduser

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceRegisteredUserId{}

// MeDeviceRegisteredUserId is a struct representing the Resource ID for a Me Device Registered User
type MeDeviceRegisteredUserId struct {
	DeviceId          string
	DirectoryObjectId string
}

// NewMeDeviceRegisteredUserID returns a new MeDeviceRegisteredUserId struct
func NewMeDeviceRegisteredUserID(deviceId string, directoryObjectId string) MeDeviceRegisteredUserId {
	return MeDeviceRegisteredUserId{
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDeviceRegisteredUserID parses 'input' into a MeDeviceRegisteredUserId
func ParseMeDeviceRegisteredUserID(input string) (*MeDeviceRegisteredUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceRegisteredUserId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceRegisteredUserId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceRegisteredUserIDInsensitively parses 'input' case-insensitively into a MeDeviceRegisteredUserId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceRegisteredUserIDInsensitively(input string) (*MeDeviceRegisteredUserId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceRegisteredUserId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceRegisteredUserId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceRegisteredUserID checks that 'input' can be parsed as a Me Device Registered User ID
func ValidateMeDeviceRegisteredUserID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceRegisteredUserID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Registered User ID
func (id MeDeviceRegisteredUserId) ID() string {
	fmtString := "/me/devices/%s/registeredUsers/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Registered User ID
func (id MeDeviceRegisteredUserId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticRegisteredUsers", "registeredUsers", "registeredUsers"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Me Device Registered User ID
func (id MeDeviceRegisteredUserId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Device Registered User (%s)", strings.Join(components, "\n"))
}
