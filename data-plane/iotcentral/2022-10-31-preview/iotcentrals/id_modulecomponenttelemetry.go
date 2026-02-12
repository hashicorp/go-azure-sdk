package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ModuleComponentTelemetryId{}

// ModuleComponentTelemetryId is a struct representing the Resource ID for a Module Component Telemetry
type ModuleComponentTelemetryId struct {
	BaseURI       string
	DeviceId      string
	ModuleName    string
	ComponentName string
	TelemetryName string
}

// NewModuleComponentTelemetryID returns a new ModuleComponentTelemetryId struct
func NewModuleComponentTelemetryID(baseURI string, deviceId string, moduleName string, componentName string, telemetryName string) ModuleComponentTelemetryId {
	return ModuleComponentTelemetryId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceId:      deviceId,
		ModuleName:    moduleName,
		ComponentName: componentName,
		TelemetryName: telemetryName,
	}
}

// ParseModuleComponentTelemetryID parses 'input' into a ModuleComponentTelemetryId
func ParseModuleComponentTelemetryID(input string) (*ModuleComponentTelemetryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleComponentTelemetryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleComponentTelemetryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseModuleComponentTelemetryIDInsensitively parses 'input' case-insensitively into a ModuleComponentTelemetryId
// note: this method should only be used for API response data and not user input
func ParseModuleComponentTelemetryIDInsensitively(input string) (*ModuleComponentTelemetryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleComponentTelemetryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleComponentTelemetryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ModuleComponentTelemetryId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ComponentName, ok = input.Parsed["componentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "componentName", input)
	}

	if id.TelemetryName, ok = input.Parsed["telemetryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "telemetryName", input)
	}

	return nil
}

// ValidateModuleComponentTelemetryID checks that 'input' can be parsed as a Module Component Telemetry ID
func ValidateModuleComponentTelemetryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseModuleComponentTelemetryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Module Component Telemetry ID
func (id ModuleComponentTelemetryId) ID() string {
	fmtString := "%s/devices/%s/modules/%s/components/%s/telemetry/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ModuleName, id.ComponentName, id.TelemetryName)
}

// Path returns the formatted Module Component Telemetry ID without the BaseURI
func (id ModuleComponentTelemetryId) Path() string {
	fmtString := "/devices/%s/modules/%s/components/%s/telemetry/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ModuleName, id.ComponentName, id.TelemetryName)
}

// PathElements returns the values of Module Component Telemetry ID Segments without the BaseURI
func (id ModuleComponentTelemetryId) PathElements() []any {
	return []any{id.DeviceId, id.ModuleName, id.ComponentName, id.TelemetryName}
}

// Segments returns a slice of Resource ID Segments which comprise this Module Component Telemetry ID
func (id ModuleComponentTelemetryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleName"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
		resourceids.StaticSegment("staticTelemetry", "telemetry", "telemetry"),
		resourceids.UserSpecifiedSegment("telemetryName", "telemetryName"),
	}
}

// String returns a human-readable description of this Module Component Telemetry ID
func (id ModuleComponentTelemetryId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Telemetry Name: %q", id.TelemetryName),
	}
	return fmt.Sprintf("Module Component Telemetry (%s)", strings.Join(components, "\n"))
}
