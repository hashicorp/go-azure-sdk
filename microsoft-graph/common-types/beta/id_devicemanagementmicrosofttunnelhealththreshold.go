package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementMicrosoftTunnelHealthThresholdId{}

// DeviceManagementMicrosoftTunnelHealthThresholdId is a struct representing the Resource ID for a Device Management Microsoft Tunnel Health Threshold
type DeviceManagementMicrosoftTunnelHealthThresholdId struct {
	MicrosoftTunnelHealthThresholdId string
}

// NewDeviceManagementMicrosoftTunnelHealthThresholdID returns a new DeviceManagementMicrosoftTunnelHealthThresholdId struct
func NewDeviceManagementMicrosoftTunnelHealthThresholdID(microsoftTunnelHealthThresholdId string) DeviceManagementMicrosoftTunnelHealthThresholdId {
	return DeviceManagementMicrosoftTunnelHealthThresholdId{
		MicrosoftTunnelHealthThresholdId: microsoftTunnelHealthThresholdId,
	}
}

// ParseDeviceManagementMicrosoftTunnelHealthThresholdID parses 'input' into a DeviceManagementMicrosoftTunnelHealthThresholdId
func ParseDeviceManagementMicrosoftTunnelHealthThresholdID(input string) (*DeviceManagementMicrosoftTunnelHealthThresholdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelHealthThresholdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelHealthThresholdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementMicrosoftTunnelHealthThresholdIDInsensitively parses 'input' case-insensitively into a DeviceManagementMicrosoftTunnelHealthThresholdId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementMicrosoftTunnelHealthThresholdIDInsensitively(input string) (*DeviceManagementMicrosoftTunnelHealthThresholdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementMicrosoftTunnelHealthThresholdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementMicrosoftTunnelHealthThresholdId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementMicrosoftTunnelHealthThresholdId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.MicrosoftTunnelHealthThresholdId, ok = input.Parsed["microsoftTunnelHealthThresholdId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "microsoftTunnelHealthThresholdId", input)
	}

	return nil
}

// ValidateDeviceManagementMicrosoftTunnelHealthThresholdID checks that 'input' can be parsed as a Device Management Microsoft Tunnel Health Threshold ID
func ValidateDeviceManagementMicrosoftTunnelHealthThresholdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementMicrosoftTunnelHealthThresholdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Microsoft Tunnel Health Threshold ID
func (id DeviceManagementMicrosoftTunnelHealthThresholdId) ID() string {
	fmtString := "/deviceManagement/microsoftTunnelHealthThresholds/%s"
	return fmt.Sprintf(fmtString, id.MicrosoftTunnelHealthThresholdId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Microsoft Tunnel Health Threshold ID
func (id DeviceManagementMicrosoftTunnelHealthThresholdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("microsoftTunnelHealthThresholds", "microsoftTunnelHealthThresholds", "microsoftTunnelHealthThresholds"),
		resourceids.UserSpecifiedSegment("microsoftTunnelHealthThresholdId", "microsoftTunnelHealthThresholdId"),
	}
}

// String returns a human-readable description of this Device Management Microsoft Tunnel Health Threshold ID
func (id DeviceManagementMicrosoftTunnelHealthThresholdId) String() string {
	components := []string{
		fmt.Sprintf("Microsoft Tunnel Health Threshold: %q", id.MicrosoftTunnelHealthThresholdId),
	}
	return fmt.Sprintf("Device Management Microsoft Tunnel Health Threshold (%s)", strings.Join(components, "\n"))
}
