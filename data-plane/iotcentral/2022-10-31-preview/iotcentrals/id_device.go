package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceId{}

// DeviceId is a struct representing the Resource ID for a Device
type DeviceId struct {
	BaseURI  string
	DeviceId string
}

// NewDeviceID returns a new DeviceId struct
func NewDeviceID(baseURI string, deviceId string) DeviceId {
	return DeviceId{
		BaseURI:  strings.TrimSuffix(baseURI, "/"),
		DeviceId: deviceId,
	}
}

// ParseDeviceID parses 'input' into a DeviceId
func ParseDeviceID(input string) (*DeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceIDInsensitively parses 'input' case-insensitively into a DeviceId
// note: this method should only be used for API response data and not user input
func ParseDeviceIDInsensitively(input string) (*DeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	return nil
}

// ValidateDeviceID checks that 'input' can be parsed as a Device ID
func ValidateDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device ID
func (id DeviceId) ID() string {
	fmtString := "%s/devices/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId)
}

// Path returns the formatted Device ID without the BaseURI
func (id DeviceId) Path() string {
	fmtString := "/devices/%s"
	return fmt.Sprintf(fmtString, id.DeviceId)
}

// PathElements returns the values of Device ID Segments without the BaseURI
func (id DeviceId) PathElements() []any {
	return []any{id.DeviceId}
}

// Segments returns a slice of Resource ID Segments which comprise this Device ID
func (id DeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
	}
}

// String returns a human-readable description of this Device ID
func (id DeviceId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
	}
	return fmt.Sprintf("Device (%s)", strings.Join(components, "\n"))
}
