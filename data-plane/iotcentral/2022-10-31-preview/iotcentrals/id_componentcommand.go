package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ComponentCommandId{}

// ComponentCommandId is a struct representing the Resource ID for a Component Command
type ComponentCommandId struct {
	BaseURI       string
	DeviceId      string
	ComponentName string
	CommandName   string
}

// NewComponentCommandID returns a new ComponentCommandId struct
func NewComponentCommandID(baseURI string, deviceId string, componentName string, commandName string) ComponentCommandId {
	return ComponentCommandId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceId:      deviceId,
		ComponentName: componentName,
		CommandName:   commandName,
	}
}

// ParseComponentCommandID parses 'input' into a ComponentCommandId
func ParseComponentCommandID(input string) (*ComponentCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ComponentCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ComponentCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseComponentCommandIDInsensitively parses 'input' case-insensitively into a ComponentCommandId
// note: this method should only be used for API response data and not user input
func ParseComponentCommandIDInsensitively(input string) (*ComponentCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ComponentCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ComponentCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ComponentCommandId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.CommandName, ok = input.Parsed["commandName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "commandName", input)
	}

	return nil
}

// ValidateComponentCommandID checks that 'input' can be parsed as a Component Command ID
func ValidateComponentCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseComponentCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Component Command ID
func (id ComponentCommandId) ID() string {
	fmtString := "%s/devices/%s/components/%s/commands/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ComponentName, id.CommandName)
}

// Path returns the formatted Component Command ID without the BaseURI
func (id ComponentCommandId) Path() string {
	fmtString := "/devices/%s/components/%s/commands/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ComponentName, id.CommandName)
}

// PathElements returns the values of Component Command ID Segments without the BaseURI
func (id ComponentCommandId) PathElements() []any {
	return []any{id.DeviceId, id.ComponentName, id.CommandName}
}

// Segments returns a slice of Resource ID Segments which comprise this Component Command ID
func (id ComponentCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
		resourceids.StaticSegment("staticCommands", "commands", "commands"),
		resourceids.UserSpecifiedSegment("commandName", "commandName"),
	}
}

// String returns a human-readable description of this Component Command ID
func (id ComponentCommandId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Command Name: %q", id.CommandName),
	}
	return fmt.Sprintf("Component Command (%s)", strings.Join(components, "\n"))
}
