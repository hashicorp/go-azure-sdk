package medevicemanagementtroubleshootingevent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeDeviceManagementTroubleshootingEventId{}

// MeDeviceManagementTroubleshootingEventId is a struct representing the Resource ID for a Me Device Management Troubleshooting Event
type MeDeviceManagementTroubleshootingEventId struct {
	DeviceManagementTroubleshootingEventId string
}

// NewMeDeviceManagementTroubleshootingEventID returns a new MeDeviceManagementTroubleshootingEventId struct
func NewMeDeviceManagementTroubleshootingEventID(deviceManagementTroubleshootingEventId string) MeDeviceManagementTroubleshootingEventId {
	return MeDeviceManagementTroubleshootingEventId{
		DeviceManagementTroubleshootingEventId: deviceManagementTroubleshootingEventId,
	}
}

// ParseMeDeviceManagementTroubleshootingEventID parses 'input' into a MeDeviceManagementTroubleshootingEventId
func ParseMeDeviceManagementTroubleshootingEventID(input string) (*MeDeviceManagementTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceManagementTroubleshootingEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceManagementTroubleshootingEventId{}

	if id.DeviceManagementTroubleshootingEventId, ok = parsed.Parsed["deviceManagementTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTroubleshootingEventId", *parsed)
	}

	return &id, nil
}

// ParseMeDeviceManagementTroubleshootingEventIDInsensitively parses 'input' case-insensitively into a MeDeviceManagementTroubleshootingEventId
// note: this method should only be used for API response data and not user input
func ParseMeDeviceManagementTroubleshootingEventIDInsensitively(input string) (*MeDeviceManagementTroubleshootingEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeDeviceManagementTroubleshootingEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeDeviceManagementTroubleshootingEventId{}

	if id.DeviceManagementTroubleshootingEventId, ok = parsed.Parsed["deviceManagementTroubleshootingEventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "deviceManagementTroubleshootingEventId", *parsed)
	}

	return &id, nil
}

// ValidateMeDeviceManagementTroubleshootingEventID checks that 'input' can be parsed as a Me Device Management Troubleshooting Event ID
func ValidateMeDeviceManagementTroubleshootingEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDeviceManagementTroubleshootingEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Device Management Troubleshooting Event ID
func (id MeDeviceManagementTroubleshootingEventId) ID() string {
	fmtString := "/me/deviceManagementTroubleshootingEvents/%s"
	return fmt.Sprintf(fmtString, id.DeviceManagementTroubleshootingEventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Device Management Troubleshooting Event ID
func (id MeDeviceManagementTroubleshootingEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticDeviceManagementTroubleshootingEvents", "deviceManagementTroubleshootingEvents", "deviceManagementTroubleshootingEvents"),
		resourceids.UserSpecifiedSegment("deviceManagementTroubleshootingEventId", "deviceManagementTroubleshootingEventIdValue"),
	}
}

// String returns a human-readable description of this Me Device Management Troubleshooting Event ID
func (id MeDeviceManagementTroubleshootingEventId) String() string {
	components := []string{
		fmt.Sprintf("Device Management Troubleshooting Event: %q", id.DeviceManagementTroubleshootingEventId),
	}
	return fmt.Sprintf("Me Device Management Troubleshooting Event (%s)", strings.Join(components, "\n"))
}
