package medevicecommand

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceCommandId{}

// MeDeviceCommandId is a struct representing the Resource ID for a Me Device Command
type MeDeviceCommandId struct {
	DeviceId  string
	CommandId string
}

// NewMeDeviceCommandID returns a new MeDeviceCommandId struct
func NewMeDeviceCommandID(deviceId string, commandId string) MeDeviceCommandId {
	return MeDeviceCommandId{
		DeviceId:  deviceId,
		CommandId: commandId,
	}
}

// ParseMeDeviceCommandID parses 'input' into a MeDeviceCommandId
func ParseMeDeviceCommandID(input string) (*MeDeviceCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceCommandId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.CommandId, ok = parsed.Parsed["commandId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "commandId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceCommandIDInsensitively parses 'input' case-insensitively into a MeDeviceCommandId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceCommandIDInsensitively(input string) (*MeDeviceCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceCommandId{}

	if id.DeviceId, ok = parsed.Parsed["deviceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceId", *parsed)
	}

	if id.CommandId, ok = parsed.Parsed["commandId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "commandId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceCommandID checks that 'input' can be parsed as a Me Device Command ID
func ValidateMeDeviceCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Command ID
func (id MeDeviceCommandId) ID() string {
	fmtString := "/me/devices/%s/commands/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.CommandId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Command ID
func (id MeDeviceCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceIdValue"),
		resourceids.StaticSegment("staticCommands", "commands", "commands"),
		resourceids.UserSpecifiedSegment("commandId", "commandIdValue"),
	}
}

// String returns a human-readable description of this Me Device Command ID
func (id MeDeviceCommandId) String() string {
	components := []string{
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Command: %q", id.CommandId),
	}
	return fmt.Sprintf("Me Device Command (%s)", strings.Join(components, "\n"))
}
