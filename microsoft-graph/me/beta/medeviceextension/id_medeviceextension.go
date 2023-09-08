package medeviceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceExtensionId{}

// MeDeviceExtensionId is a struct representing the Resource ID for a Me Device Extension
type MeDeviceExtensionId struct {
	DeviceId    string
	ExtensionId string
}

// NewMeDeviceExtensionID returns a new MeDeviceExtensionId struct
func NewMeDeviceExtensionID(deviceId string, extensionId string) MeDeviceExtensionId {
	return MeDeviceExtensionId{
		DeviceId:    deviceId,
		ExtensionId: extensionId,
	}
}

// ParseMeDeviceExtensionID parses 'input' into a MeDeviceExtensionId
func ParseMeDeviceExtensionID(input string) (*MeDeviceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceExtensionId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceExtensionIDInsensitively parses 'input' case-insensitively into a MeDeviceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceExtensionIDInsensitively(input string) (*MeDeviceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceExtensionId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceExtensionID checks that 'input' can be parsed as a Me Device Extension ID
func ValidateMeDeviceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Extension ID
func (id MeDeviceExtensionId) ID() string {
	fmtString := "/me/devices/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Extension ID
func (id MeDeviceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Device Extension ID
func (id MeDeviceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Device Extension (%s)", strings.Join(components, "\n"))
}
