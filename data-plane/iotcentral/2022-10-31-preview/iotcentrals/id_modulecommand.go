package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ModuleCommandId{}

// ModuleCommandId is a struct representing the Resource ID for a Module Command
type ModuleCommandId struct {
	BaseURI     string
	DeviceId    string
	ModuleName  string
	CommandName string
}

// NewModuleCommandID returns a new ModuleCommandId struct
func NewModuleCommandID(baseURI string, deviceId string, moduleName string, commandName string) ModuleCommandId {
	return ModuleCommandId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		DeviceId:    deviceId,
		ModuleName:  moduleName,
		CommandName: commandName,
	}
}

// ParseModuleCommandID parses 'input' into a ModuleCommandId
func ParseModuleCommandID(input string) (*ModuleCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleCommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseModuleCommandIDInsensitively parses 'input' case-insensitively into a ModuleCommandId
// note: this method should only be used for API response data and not user input
func ParseModuleCommandIDInsensitively(input string) (*ModuleCommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleCommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleCommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ModuleCommandId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.CommandName, ok = input.Parsed["commandName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "commandName", input)
	}

	return nil
}

// ValidateModuleCommandID checks that 'input' can be parsed as a Module Command ID
func ValidateModuleCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseModuleCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Module Command ID
func (id ModuleCommandId) ID() string {
	fmtString := "%s/devices/%s/modules/%s/commands/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ModuleName, id.CommandName)
}

// Path returns the formatted Module Command ID without the BaseURI
func (id ModuleCommandId) Path() string {
	fmtString := "/devices/%s/modules/%s/commands/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ModuleName, id.CommandName)
}

// PathElements returns the values of Module Command ID Segments without the BaseURI
func (id ModuleCommandId) PathElements() []any {
	return []any{id.DeviceId, id.ModuleName, id.CommandName}
}

// Segments returns a slice of Resource ID Segments which comprise this Module Command ID
func (id ModuleCommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleName"),
		resourceids.StaticSegment("staticCommands", "commands", "commands"),
		resourceids.UserSpecifiedSegment("commandName", "commandName"),
	}
}

// String returns a human-readable description of this Module Command ID
func (id ModuleCommandId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
		fmt.Sprintf("Command Name: %q", id.CommandName),
	}
	return fmt.Sprintf("Module Command (%s)", strings.Join(components, "\n"))
}
