package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ComponentTelemetryId{}

// ComponentTelemetryId is a struct representing the Resource ID for a Component Telemetry
type ComponentTelemetryId struct {
	BaseURI       string
	DeviceId      string
	ComponentName string
	TelemetryName string
}

// NewComponentTelemetryID returns a new ComponentTelemetryId struct
func NewComponentTelemetryID(baseURI string, deviceId string, componentName string, telemetryName string) ComponentTelemetryId {
	return ComponentTelemetryId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceId:      deviceId,
		ComponentName: componentName,
		TelemetryName: telemetryName,
	}
}

// ParseComponentTelemetryID parses 'input' into a ComponentTelemetryId
func ParseComponentTelemetryID(input string) (*ComponentTelemetryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ComponentTelemetryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ComponentTelemetryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseComponentTelemetryIDInsensitively parses 'input' case-insensitively into a ComponentTelemetryId
// note: this method should only be used for API response data and not user input
func ParseComponentTelemetryIDInsensitively(input string) (*ComponentTelemetryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ComponentTelemetryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ComponentTelemetryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ComponentTelemetryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.ComponentName, ok = input.Parsed["componentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "componentName", input)
	}

	if id.TelemetryName, ok = input.Parsed["telemetryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "telemetryName", input)
	}

	return nil
}

// ValidateComponentTelemetryID checks that 'input' can be parsed as a Component Telemetry ID
func ValidateComponentTelemetryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseComponentTelemetryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Component Telemetry ID
func (id ComponentTelemetryId) ID() string {
	fmtString := "%s/devices/%s/components/%s/telemetry/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ComponentName, id.TelemetryName)
}

// Path returns the formatted Component Telemetry ID without the BaseURI
func (id ComponentTelemetryId) Path() string {
	fmtString := "/devices/%s/components/%s/telemetry/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ComponentName, id.TelemetryName)
}

// PathElements returns the values of Component Telemetry ID Segments without the BaseURI
func (id ComponentTelemetryId) PathElements() []any {
	return []any{id.DeviceId, id.ComponentName, id.TelemetryName}
}

// Segments returns a slice of Resource ID Segments which comprise this Component Telemetry ID
func (id ComponentTelemetryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
		resourceids.StaticSegment("staticTelemetry", "telemetry", "telemetry"),
		resourceids.UserSpecifiedSegment("telemetryName", "telemetryName"),
	}
}

// String returns a human-readable description of this Component Telemetry ID
func (id ComponentTelemetryId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Telemetry Name: %q", id.TelemetryName),
	}
	return fmt.Sprintf("Component Telemetry (%s)", strings.Join(components, "\n"))
}
