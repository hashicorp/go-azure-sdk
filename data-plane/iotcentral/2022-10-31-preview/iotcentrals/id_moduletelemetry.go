package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ModuleTelemetryId{}

// ModuleTelemetryId is a struct representing the Resource ID for a Module Telemetry
type ModuleTelemetryId struct {
	BaseURI       string
	DeviceId      string
	ModuleName    string
	TelemetryName string
}

// NewModuleTelemetryID returns a new ModuleTelemetryId struct
func NewModuleTelemetryID(baseURI string, deviceId string, moduleName string, telemetryName string) ModuleTelemetryId {
	return ModuleTelemetryId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceId:      deviceId,
		ModuleName:    moduleName,
		TelemetryName: telemetryName,
	}
}

// ParseModuleTelemetryID parses 'input' into a ModuleTelemetryId
func ParseModuleTelemetryID(input string) (*ModuleTelemetryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleTelemetryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleTelemetryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseModuleTelemetryIDInsensitively parses 'input' case-insensitively into a ModuleTelemetryId
// note: this method should only be used for API response data and not user input
func ParseModuleTelemetryIDInsensitively(input string) (*ModuleTelemetryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleTelemetryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleTelemetryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ModuleTelemetryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.ModuleName, ok = input.Parsed["moduleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "moduleName", input)
	}

	if id.TelemetryName, ok = input.Parsed["telemetryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "telemetryName", input)
	}

	return nil
}

// ValidateModuleTelemetryID checks that 'input' can be parsed as a Module Telemetry ID
func ValidateModuleTelemetryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseModuleTelemetryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Module Telemetry ID
func (id ModuleTelemetryId) ID() string {
	fmtString := "%s/devices/%s/modules/%s/telemetry/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ModuleName, id.TelemetryName)
}

// Path returns the formatted Module Telemetry ID without the BaseURI
func (id ModuleTelemetryId) Path() string {
	fmtString := "/devices/%s/modules/%s/telemetry/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ModuleName, id.TelemetryName)
}

// PathElements returns the values of Module Telemetry ID Segments without the BaseURI
func (id ModuleTelemetryId) PathElements() []any {
	return []any{id.DeviceId, id.ModuleName, id.TelemetryName}
}

// Segments returns a slice of Resource ID Segments which comprise this Module Telemetry ID
func (id ModuleTelemetryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleName"),
		resourceids.StaticSegment("staticTelemetry", "telemetry", "telemetry"),
		resourceids.UserSpecifiedSegment("telemetryName", "telemetryName"),
	}
}

// String returns a human-readable description of this Module Telemetry ID
func (id ModuleTelemetryId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
		fmt.Sprintf("Telemetry Name: %q", id.TelemetryName),
	}
	return fmt.Sprintf("Module Telemetry (%s)", strings.Join(components, "\n"))
}
