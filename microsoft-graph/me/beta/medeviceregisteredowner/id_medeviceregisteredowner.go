package medeviceregisteredowner

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceRegisteredOwnerId{}

// MeDeviceRegisteredOwnerId is a struct representing the Resource ID for a Me Device Registered Owner
type MeDeviceRegisteredOwnerId struct {
	DeviceId          string
	DirectoryObjectId string
}

// NewMeDeviceRegisteredOwnerID returns a new MeDeviceRegisteredOwnerId struct
func NewMeDeviceRegisteredOwnerID(deviceId string, directoryObjectId string) MeDeviceRegisteredOwnerId {
	return MeDeviceRegisteredOwnerId{
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDeviceRegisteredOwnerID parses 'input' into a MeDeviceRegisteredOwnerId
func ParseMeDeviceRegisteredOwnerID(input string) (*MeDeviceRegisteredOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceRegisteredOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceRegisteredOwnerId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceRegisteredOwnerIDInsensitively parses 'input' case-insensitively into a MeDeviceRegisteredOwnerId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceRegisteredOwnerIDInsensitively(input string) (*MeDeviceRegisteredOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceRegisteredOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceRegisteredOwnerId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceRegisteredOwnerID checks that 'input' can be parsed as a Me Device Registered Owner ID
func ValidateMeDeviceRegisteredOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceRegisteredOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Registered Owner ID
func (id MeDeviceRegisteredOwnerId) ID() string {
	fmtString := "/me/devices/%s/registeredOwners/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Registered Owner ID
func (id MeDeviceRegisteredOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticRegisteredOwners", "registeredOwners", "registeredOwners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Me Device Registered Owner ID
func (id MeDeviceRegisteredOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Device Registered Owner (%s)", strings.Join(components, "\n"))
}
