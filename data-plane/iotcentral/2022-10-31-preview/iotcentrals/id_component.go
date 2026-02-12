package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ComponentId{}

// ComponentId is a struct representing the Resource ID for a Component
type ComponentId struct {
	BaseURI       string
	DeviceId      string
	ComponentName string
}

// NewComponentID returns a new ComponentId struct
func NewComponentID(baseURI string, deviceId string, componentName string) ComponentId {
	return ComponentId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceId:      deviceId,
		ComponentName: componentName,
	}
}

// ParseComponentID parses 'input' into a ComponentId
func ParseComponentID(input string) (*ComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ComponentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ComponentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseComponentIDInsensitively parses 'input' case-insensitively into a ComponentId
// note: this method should only be used for API response data and not user input
func ParseComponentIDInsensitively(input string) (*ComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ComponentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ComponentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ComponentId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateComponentID checks that 'input' can be parsed as a Component ID
func ValidateComponentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseComponentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Component ID
func (id ComponentId) ID() string {
	fmtString := "%s/devices/%s/components/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ComponentName)
}

// Path returns the formatted Component ID without the BaseURI
func (id ComponentId) Path() string {
	fmtString := "/devices/%s/components/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ComponentName)
}

// PathElements returns the values of Component ID Segments without the BaseURI
func (id ComponentId) PathElements() []any {
	return []any{id.DeviceId, id.ComponentName}
}

// Segments returns a slice of Resource ID Segments which comprise this Component ID
func (id ComponentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
	}
}

// String returns a human-readable description of this Component ID
func (id ComponentId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
	}
	return fmt.Sprintf("Component (%s)", strings.Join(components, "\n"))
}
