package medeviceusageright

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceUsageRightId{}

// MeDeviceUsageRightId is a struct representing the Resource ID for a Me Device Usage Right
type MeDeviceUsageRightId struct {
	DeviceId     string
	UsageRightId string
}

// NewMeDeviceUsageRightID returns a new MeDeviceUsageRightId struct
func NewMeDeviceUsageRightID(deviceId string, usageRightId string) MeDeviceUsageRightId {
	return MeDeviceUsageRightId{
		DeviceId:     deviceId,
		UsageRightId: usageRightId,
	}
}

// ParseMeDeviceUsageRightID parses 'input' into a MeDeviceUsageRightId
func ParseMeDeviceUsageRightID(input string) (*MeDeviceUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceUsageRightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceUsageRightId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.UsageRightId, ok = parsed.Parsed["usageRightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceUsageRightIDInsensitively parses 'input' case-insensitively into a MeDeviceUsageRightId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceUsageRightIDInsensitively(input string) (*MeDeviceUsageRightId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceUsageRightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceUsageRightId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.UsageRightId, ok = parsed.Parsed["usageRightId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usageRightId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceUsageRightID checks that 'input' can be parsed as a Me Device Usage Right ID
func ValidateMeDeviceUsageRightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceUsageRightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Usage Right ID
func (id MeDeviceUsageRightId) ID() string {
	fmtString := "/me/devices/%s/usageRights/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.UsageRightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Usage Right ID
func (id MeDeviceUsageRightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticUsageRights", "usageRights", "usageRights"),
		resourceids.UserSpecifiedSegment("usageRightId", "usageRightIdValue"),
	}
}

// String returns a human-readable description of this Me Device Usage Right ID
func (id MeDeviceUsageRightId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Usage Right: %q", id.UsageRightId),
	}
	return fmt.Sprintf("Me Device Usage Right (%s)", strings.Join(components, "\n"))
}
