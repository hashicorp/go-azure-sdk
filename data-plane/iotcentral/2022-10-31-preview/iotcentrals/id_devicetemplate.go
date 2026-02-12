package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DeviceTemplateId{}

// DeviceTemplateId is a struct representing the Resource ID for a Device Template
type DeviceTemplateId struct {
	BaseURI          string
	DeviceTemplateId string
}

// NewDeviceTemplateID returns a new DeviceTemplateId struct
func NewDeviceTemplateID(baseURI string, deviceTemplateId string) DeviceTemplateId {
	return DeviceTemplateId{
		BaseURI:          strings.TrimSuffix(baseURI, "/"),
		DeviceTemplateId: deviceTemplateId,
	}
}

// ParseDeviceTemplateID parses 'input' into a DeviceTemplateId
func ParseDeviceTemplateID(input string) (*DeviceTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeviceTemplateIDInsensitively parses 'input' case-insensitively into a DeviceTemplateId
// note: this method should only be used for API response data and not user input
func ParseDeviceTemplateIDInsensitively(input string) (*DeviceTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeviceTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeviceTemplateId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeviceTemplateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeviceTemplateId, ok = input.Parsed["deviceTemplateId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceTemplateId", input)
	}

	return nil
}

// ValidateDeviceTemplateID checks that 'input' can be parsed as a Device Template ID
func ValidateDeviceTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeviceTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Device Template ID
func (id DeviceTemplateId) ID() string {
	fmtString := "%s/deviceTemplates/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceTemplateId)
}

// Path returns the formatted Device Template ID without the BaseURI
func (id DeviceTemplateId) Path() string {
	fmtString := "/deviceTemplates/%s"
	return fmt.Sprintf(fmtString, id.DeviceTemplateId)
}

// PathElements returns the values of Device Template ID Segments without the BaseURI
func (id DeviceTemplateId) PathElements() []any {
	return []any{id.DeviceTemplateId}
}

// Segments returns a slice of Resource ID Segments which comprise this Device Template ID
func (id DeviceTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDeviceTemplates", "deviceTemplates", "deviceTemplates"),
		resourceids.UserSpecifiedSegment("deviceTemplateId", "deviceTemplateId"),
	}
}

// String returns a human-readable description of this Device Template ID
func (id DeviceTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device Template: %q", id.DeviceTemplateId),
	}
	return fmt.Sprintf("Device Template (%s)", strings.Join(components, "\n"))
}
