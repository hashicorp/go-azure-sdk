package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceManagementManagedDeviceIdDetectedAppId{}

// DeviceManagementManagedDeviceIdDetectedAppId is a struct representing the Resource ID for a Device Management Managed Device Id Detected App
type DeviceManagementManagedDeviceIdDetectedAppId struct {
	ManagedDeviceId string
	DetectedAppId   string
}

// NewDeviceManagementManagedDeviceIdDetectedAppID returns a new DeviceManagementManagedDeviceIdDetectedAppId struct
func NewDeviceManagementManagedDeviceIdDetectedAppID(managedDeviceId string, detectedAppId string) DeviceManagementManagedDeviceIdDetectedAppId {
	return DeviceManagementManagedDeviceIdDetectedAppId{
		ManagedDeviceId: managedDeviceId,
		DetectedAppId:   detectedAppId,
	}
}

// ParseDeviceManagementManagedDeviceIdDetectedAppID parses 'input' into a DeviceManagementManagedDeviceIdDetectedAppId
func ParseDeviceManagementManagedDeviceIdDetectedAppID(input string) (*DeviceManagementManagedDeviceIdDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdDetectedAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceManagementManagedDeviceIdDetectedAppIDInsensitively parses 'input' case-insensitively into a DeviceManagementManagedDeviceIdDetectedAppId
// note: this method should only be used for API response data and not user input
func ParseDeviceManagementManagedDeviceIdDetectedAppIDInsensitively(input string) (*DeviceManagementManagedDeviceIdDetectedAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceManagementManagedDeviceIdDetectedAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceManagementManagedDeviceIdDetectedAppId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceManagementManagedDeviceIdDetectedAppId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagedDeviceId, ok = input.Parsed["managedDeviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedDeviceId", input)
	}

	if id.DetectedAppId, ok = input.Parsed["detectedAppId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectedAppId", input)
	}

	return nil
}

// ValidateDeviceManagementManagedDeviceIdDetectedAppID checks that 'input' can be parsed as a Device Management Managed Device Id Detected App ID
func ValidateDeviceManagementManagedDeviceIdDetectedAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceManagementManagedDeviceIdDetectedAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Management Managed Device Id Detected App ID
func (id DeviceManagementManagedDeviceIdDetectedAppId) ID() string {
	fmtString := "/deviceManagement/managedDevices/%s/detectedApps/%s"
	return fmt.Sprintf(fmtString, id.ManagedDeviceId, id.DetectedAppId)
}

// Segments returns a slice of Resource ID Segments which comprise this Device Management Managed Device Id Detected App ID
func (id DeviceManagementManagedDeviceIdDetectedAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("deviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("managedDevices", "managedDevices", "managedDevices"),
		resourceids.UserSpecifiedSegment("managedDeviceId", "managedDeviceId"),
		resourceids.StaticSegment("detectedApps", "detectedApps", "detectedApps"),
		resourceids.UserSpecifiedSegment("detectedAppId", "detectedAppId"),
	}
}

// String returns a human-readable description of this Device Management Managed Device Id Detected App ID
func (id DeviceManagementManagedDeviceIdDetectedAppId) String() string {
	components := []string{
		fmt.Sprintf("Managed Device: %q", id.ManagedDeviceId),
		fmt.Sprintf("Detected App: %q", id.DetectedAppId),
	}
	return fmt.Sprintf("Device Management Managed Device Id Detected App (%s)", strings.Join(components, "\n"))
}
