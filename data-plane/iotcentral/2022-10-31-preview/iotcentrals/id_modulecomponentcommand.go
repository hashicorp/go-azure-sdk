package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ModuleComponentCommandId{}

// ModuleComponentCommandId is a struct representing the Resource ID for a Module Component Command
type ModuleComponentCommandId struct {
	BaseURI       string
	DeviceId      string
	ModuleName    string
	ComponentName string
	CommandName   string
}

// NewModuleComponentCommandID returns a new ModuleComponentCommandId struct
func NewModuleComponentCommandID(baseURI string, deviceId string, moduleName string, componentName string, commandName string) ModuleComponentCommandId {
	return ModuleComponentCommandId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceId:      deviceId,
		ModuleName:    moduleName,
		ComponentName: componentName,
		CommandName:   commandName,
	}
}

// ParseModuleComponentCommandID parses 'input' into a ModuleComponentCommandId
func ParseModuleComponentCommandID(input string) (*ModuleComponentCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleComponentCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleComponentCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseModuleComponentCommandIDInsensitively parses 'input' case-insensitively into a ModuleComponentCommandId
// note: this method should only be used for API response data and not user input
func ParseModuleComponentCommandIDInsensitively(input string) (*ModuleComponentCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleComponentCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleComponentCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ModuleComponentCommandId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.CommandName, ok = input.Parsed["commandName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "commandName", input)
	}

	return nil
}

// ValidateModuleComponentCommandID checks that 'input' can be parsed as a Module Component Command ID
func ValidateModuleComponentCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseModuleComponentCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Module Component Command ID
func (id ModuleComponentCommandId) ID() string {
	fmtString := "%s/devices/%s/modules/%s/components/%s/commands/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ModuleName, id.ComponentName, id.CommandName)
}

// Path returns the formatted Module Component Command ID without the BaseURI
func (id ModuleComponentCommandId) Path() string {
	fmtString := "/devices/%s/modules/%s/components/%s/commands/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ModuleName, id.ComponentName, id.CommandName)
}

// PathElements returns the values of Module Component Command ID Segments without the BaseURI
func (id ModuleComponentCommandId) PathElements() []any {
	return []any{id.DeviceId, id.ModuleName, id.ComponentName, id.CommandName}
}

// Segments returns a slice of Resource ID Segments which comprise this Module Component Command ID
func (id ModuleComponentCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleName"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
		resourceids.StaticSegment("staticCommands", "commands", "commands"),
		resourceids.UserSpecifiedSegment("commandName", "commandName"),
	}
}

// String returns a human-readable description of this Module Component Command ID
func (id ModuleComponentCommandId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Command Name: %q", id.CommandName),
	}
	return fmt.Sprintf("Module Component Command (%s)", strings.Join(components, "\n"))
}
