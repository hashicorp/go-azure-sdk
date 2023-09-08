package memanageddevicedeviceconfigurationstate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeManagedDeviceDeviceConfigurationStateId{}

// MeManagedDeviceDeviceConfigurationStateId is a struct representing the Resource ID for a Me Managed Device Device Configuration State
type MeManagedDeviceDeviceConfigurationStateId struct {
	ManagedDeviceId            string
	DeviceConfigurationStateId string
}

// NewMeManagedDeviceDeviceConfigurationStateID returns a new MeManagedDeviceDeviceConfigurationStateId struct
func NewMeManagedDeviceDeviceConfigurationStateID(managedDeviceId string, deviceConfigurationStateId string) MeManagedDeviceDeviceConfigurationStateId {
	return MeManagedDeviceDeviceConfigurationStateId{
		ManagedDeviceId:            managedDeviceId,
		DeviceConfigurationStateId: deviceConfigurationStateId,
	}
}

// ParseMeManagedDeviceDeviceConfigurationStateID parses 'input' into a MeManagedDeviceDeviceConfigurationStateId
func ParseMeManagedDeviceDeviceConfigurationStateID(input string) (*MeManagedDeviceDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceDeviceConfigurationStateId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceConfigurationStateId, ok = parsed.Parsed["deviceConfigurationStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationStateId", *parsed)
	}

	return &id, nil
}

// ParseMeManagedDeviceDeviceConfigurationStateIDInsensitively parses 'input' case-insensitively into a MeManagedDeviceDeviceConfigurationStateId
// note: this method should only be used for API response data and not user input
func ParseMeManagedDeviceDeviceConfigurationStateIDInsensitively(input string) (*MeManagedDeviceDeviceConfigurationStateId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeManagedDeviceDeviceConfigurationStateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeManagedDeviceDeviceConfigurationStateId{}

	if id.ManagedDeviceId, ok = parsed.Parsed["managedDeviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", *parsed)
	}

	if id.DeviceConfigurationStateId, ok = parsed.Parsed["deviceConfigurationStateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceConfigurationStateId", *parsed)
	}

	return &id, nil
}

// ValidateMeManagedDeviceDeviceConfigurationStateID checks that 'input' can be parsed as a Me Managed Device Device Configuration State ID
func ValidateMeManagedDeviceDeviceConfigurationStateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeManagedDeviceDeviceConfigurationStateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Managed Device Device Configuration State ID
func (id MeManagedDeviceDeviceConfigurationStateId) ID() string {
	fmtString := "/me/managedDevices/%s/deviceConfigurationStates/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DeviceConfigurationStateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Managed Device Device Configuration State ID
func (id MeManagedDeviceDeviceConfigurationStateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticManagedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceIdValue"),
		resourceids.StaticSegment("staticDeviceConfigurationStates", "deviceConfigurationStates", "deviceConfigurationStates"),
		resourceids.UserSpecifiedSegment("deviceConfigurationStateId", "deviceConfigurationStateIdValue"),
	}
}

// String returns a human-readable description of this Me Managed Device Device Configuration State ID
func (id MeManagedDeviceDeviceConfigurationStateId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Device Configuration State: %q", id.DeviceConfigurationStateId),
	}
	return fmt.Sprintf("Me Managed Device Device Configuration State (%s)", strings.Join(components, "\n"))
}
