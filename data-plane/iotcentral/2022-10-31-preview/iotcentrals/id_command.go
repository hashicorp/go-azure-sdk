package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CommandId{}

// CommandId is a struct representing the Resource ID for a Command
type CommandId struct {
	BaseURI     string
	DeviceId    string
	CommandName string
}

// NewCommandID returns a new CommandId struct
func NewCommandID(baseURI string, deviceId string, commandName string) CommandId {
	return CommandId{
		BaseURI:     strings.TrimSuffix(baseURI, "/"),
		DeviceId:    deviceId,
		CommandName: commandName,
	}
}

// ParseCommandID parses 'input' into a CommandId
func ParseCommandID(input string) (*CommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CommandId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCommandIDInsensitively parses 'input' case-insensitively into a CommandId
// note: this method should only be used for API response data and not user input
func ParseCommandIDInsensitively(input string) (*CommandId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CommandId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CommandId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CommandId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.DeviceId, ok = input.Parsed["deviceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deviceId", input)
	}

	if id.CommandName, ok = input.Parsed["commandName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "commandName", input)
	}

	return nil
}

// ValidateCommandID checks that 'input' can be parsed as a Command ID
func ValidateCommandID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCommandID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Command ID
func (id CommandId) ID() string {
	fmtString := "%s/devices/%s/commands/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.DeviceId, id.CommandName)
}

// Path returns the formatted Command ID without the BaseURI
func (id CommandId) Path() string {
	fmtString := "/devices/%s/commands/%s"
	return fmt.Sprintf(fmtString, id.DeviceId, id.CommandName)
}

// PathElements returns the values of Command ID Segments without the BaseURI
func (id CommandId) PathElements() []any {
	return []any{id.DeviceId, id.CommandName}
}

// Segments returns a slice of Resource ID Segments which comprise this Command ID
func (id CommandId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticDevices", "devices", "devices"),
		resourceids.UserSpecifiedSegment("deviceId", "deviceId"),
		resourceids.StaticSegment("staticCommands", "commands", "commands"),
		resourceids.UserSpecifiedSegment("commandName", "commandName"),
	}
}

// String returns a human-readable description of this Command ID
func (id CommandId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Device: %q", id.DeviceId),
		fmt.Sprintf("Command Name: %q", id.CommandName),
	}
	return fmt.Sprintf("Command (%s)", strings.Join(components, "\n"))
}
