package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ModuleId{}

// ModuleId is a struct representing the Resource ID for a Module
type ModuleId struct {
	BaseURI    string
	DeviceId   string
	ModuleName string
}

// NewModuleID returns a new ModuleId struct
func NewModuleID(baseURI string, deviceId string, moduleName string) ModuleId {
	return ModuleId{
		BaseURI:    strings.TrimSuffix(baseURI, "/"),
		DeviceId:   deviceId,
		ModuleName: moduleName,
	}
}

// ParseModuleID parses 'input' into a ModuleId
func ParseModuleID(input string) (*ModuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseModuleIDInsensitively parses 'input' case-insensitively into a ModuleId
// note: this method should only be used for API response data and not user input
func ParseModuleIDInsensitively(input string) (*ModuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ModuleId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateModuleID checks that 'input' can be parsed as a Module ID
func ValidateModuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseModuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Module ID
func (id ModuleId) ID() string {
	fmtString := "%s/devices/%s/modules/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.ModuleName)
}

// Path returns the formatted Module ID without the BaseURI
func (id ModuleId) Path() string {
	fmtString := "/devices/%s/modules/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.ModuleName)
}

// PathElements returns the values of Module ID Segments without the BaseURI
func (id ModuleId) PathElements() []any {
	return []any{id.DeviceId, id.ModuleName}
}

// Segments returns a slice of Resource ID Segments which comprise this Module ID
func (id ModuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleName"),
	}
}

// String returns a human-readable description of this Module ID
func (id ModuleId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
	}
	return fmt.Sprintf("Module (%s)", strings.Join(components, "\n"))
}
