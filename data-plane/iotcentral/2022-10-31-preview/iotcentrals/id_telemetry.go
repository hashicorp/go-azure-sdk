package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &TelemetryId{}

// TelemetryId is a struct representing the Resource ID for a Telemetry
type TelemetryId struct {
	BaseURI       string
	DeviceId      string
	TelemetryName string
}

// NewTelemetryID returns a new TelemetryId struct
func NewTelemetryID(baseURI string, deviceId string, telemetryName string) TelemetryId {
	return TelemetryId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceId:      deviceId,
		TelemetryName: telemetryName,
	}
}

// ParseTelemetryID parses 'input' into a TelemetryId
func ParseTelemetryID(input string) (*TelemetryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TelemetryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TelemetryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTelemetryIDInsensitively parses 'input' case-insensitively into a TelemetryId
// note: this method should only be used for API response data and not user input
func ParseTelemetryIDInsensitively(input string) (*TelemetryId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TelemetryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TelemetryId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TelemetryId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.TelemetryName, ok = input.Parsed["telemetryName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "telemetryName", input)
	}

	return nil
}

// ValidateTelemetryID checks that 'input' can be parsed as a Telemetry ID
func ValidateTelemetryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTelemetryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Telemetry ID
func (id TelemetryId) ID() string {
	fmtString := "%s/devices/%s/telemetry/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.TelemetryName)
}

// Path returns the formatted Telemetry ID without the BaseURI
func (id TelemetryId) Path() string {
	fmtString := "/devices/%s/telemetry/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.TelemetryName)
}

// PathElements returns the values of Telemetry ID Segments without the BaseURI
func (id TelemetryId) PathElements() []any {
	return []any{id.DeviceId, id.TelemetryName}
}

// Segments returns a slice of Resource ID Segments which comprise this Telemetry ID
func (id TelemetryId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticTelemetry", "telemetry", "telemetry"),
		resourceids.UserSpecifiedSegment("telemetryName", "telemetryName"),
	}
}

// String returns a human-readable description of this Telemetry ID
func (id TelemetryId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Telemetry Name: %q", id.TelemetryName),
	}
	return fmt.Sprintf("Telemetry (%s)", strings.Join(components, "\n"))
}
