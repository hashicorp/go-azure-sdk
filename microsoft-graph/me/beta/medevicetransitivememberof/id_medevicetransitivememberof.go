package medevicetransitivememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceTransitiveMemberOfId{}

// MeDeviceTransitiveMemberOfId is a struct representing the Resource ID for a Me Device Transitive Member Of
type MeDeviceTransitiveMemberOfId struct {
	DeviceId          string
	DirectoryObjectId string
}

// NewMeDeviceTransitiveMemberOfID returns a new MeDeviceTransitiveMemberOfId struct
func NewMeDeviceTransitiveMemberOfID(deviceId string, directoryObjectId string) MeDeviceTransitiveMemberOfId {
	return MeDeviceTransitiveMemberOfId{
		DeviceId:          deviceId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseMeDeviceTransitiveMemberOfID parses 'input' into a MeDeviceTransitiveMemberOfId
func ParseMeDeviceTransitiveMemberOfID(input string) (*MeDeviceTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceTransitiveMemberOfId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a MeDeviceTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceTransitiveMemberOfIDInsensitively(input string) (*MeDeviceTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceTransitiveMemberOfId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceTransitiveMemberOfID checks that 'input' can be parsed as a Me Device Transitive Member Of ID
func ValidateMeDeviceTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Transitive Member Of ID
func (id MeDeviceTransitiveMemberOfId) ID() string {
	fmtString := "/me/devices/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Transitive Member Of ID
func (id MeDeviceTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticTransitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Me Device Transitive Member Of ID
func (id MeDeviceTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Me Device Transitive Member Of (%s)", strings.Join(components, "\n"))
}
