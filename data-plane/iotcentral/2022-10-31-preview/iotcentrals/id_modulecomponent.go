package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ModuleComponentId{}

// ModuleComponentId is a struct representing the Resource ID for a Module Component
type ModuleComponentId struct {
	BaseURI       string
	DeviceId      string
	ModuleName    string
	ComponentName string
}

// NewModuleComponentID returns a new ModuleComponentId struct
func NewModuleComponentID(baseURI string, deviceId string, moduleName string, componentName string) ModuleComponentId {
	return ModuleComponentId{
		BaseURI:       strings.TrimSuffix(baseURI, "/"),
		DeviceId:      deviceId,
		ModuleName:    moduleName,
		ComponentName: componentName,
	}
}

// ParseModuleComponentID parses 'input' into a ModuleComponentId
func ParseModuleComponentID(input string) (*ModuleComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleComponentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleComponentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseModuleComponentIDInsensitively parses 'input' case-insensitively into a ModuleComponentId
// note: this method should only be used for API response data and not user input
func ParseModuleComponentIDInsensitively(input string) (*ModuleComponentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleComponentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleComponentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ModuleComponentId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateModuleComponentID checks that 'input' can be parsed as a Module Component ID
func ValidateModuleComponentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseModuleComponentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Module Component ID
func (id ModuleComponentId) ID() string {
	fmtString := "%s/devices/%s/modules/%s/components/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ModuleName, id.ComponentName)
}

// Path returns the formatted Module Component ID without the BaseURI
func (id ModuleComponentId) Path() string {
	fmtString := "/devices/%s/modules/%s/components/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ModuleName, id.ComponentName)
}

// PathElements returns the values of Module Component ID Segments without the BaseURI
func (id ModuleComponentId) PathElements() []any {
	return []any{id.DeviceId, id.ModuleName, id.ComponentName}
}

// Segments returns a slice of Resource ID Segments which comprise this Module Component ID
func (id ModuleComponentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleName"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
	}
}

// String returns a human-readable description of this Module Component ID
func (id ModuleComponentId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
	}
	return fmt.Sprintf("Module Component (%s)", strings.Join(components, "\n"))
}
