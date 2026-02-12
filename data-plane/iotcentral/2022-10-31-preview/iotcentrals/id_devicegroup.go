package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceGroupId{}

// DeviceGroupId is a struct representing the Resource ID for a Device Group
type DeviceGroupId struct {
	BaseURI       string
	DeviceGroupId string
}

// NewDeviceGroupID returns a new DeviceGroupId struct
func NewDeviceGroupID(baseURI string, deviceGroupId string) DeviceGroupId {
	return DeviceGroupId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceGroupId: deviceGroupId,
	}
}

// ParseDeviceGroupID parses 'input' into a DeviceGroupId
func ParseDeviceGroupID(input string) (*DeviceGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceGroupIDInsensitively parses 'input' case-insensitively into a DeviceGroupId
// note: this method should only be used for API response data and not user input
func ParseDeviceGroupIDInsensitively(input string) (*DeviceGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceGroupId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeviceGroupId, ok = input.Parsed["deviceGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceGroupId", input)
	}

	return nil
}

// ValidateDeviceGroupID checks that 'input' can be parsed as a Device Group ID
func ValidateDeviceGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Group ID
func (id DeviceGroupId) ID() string {
	fmtString := "%s/deviceGroups/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceGroupId)
}

// Path returns the formatted Device Group ID without the BaseURI
func (id DeviceGroupId) Path() string {
	fmtString := "/deviceGroups/%s"
	return fmt.Sprintf(fmtString, id.DeviceGroupId)
}

// PathElements returns the values of Device Group ID Segments without the BaseURI
func (id DeviceGroupId) PathElements() []any {
	return []any{id.DeviceGroupId}
}

// Segments returns a slice of Resource ID Segments which comprise this Device Group ID
func (id DeviceGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDeviceGroups", "deviceGroups", "deviceGroups"),
		resourceids.UserSpecifiedSegment("deviceGroupId", "deviceGroupId"),
	}
}

// String returns a human-readable description of this Device Group ID
func (id DeviceGroupId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device Group: %q", id.DeviceGroupId),
	}
	return fmt.Sprintf("Device Group (%s)", strings.Join(components, "\n"))
}
