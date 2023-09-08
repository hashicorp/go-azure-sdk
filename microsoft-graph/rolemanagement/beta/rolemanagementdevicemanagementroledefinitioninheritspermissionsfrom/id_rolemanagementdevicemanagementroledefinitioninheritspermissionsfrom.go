package rolemanagementdevicemanagementroledefinitioninheritspermissionsfrom

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId{}

// RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId is a struct representing the Resource ID for a Role Management Device Management Role Definition Inherits Permissions From
type RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId struct {
	UnifiedRoleDefinitionId  string
	UnifiedRoleDefinitionId1 string
}

// NewRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromID returns a new RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId struct
func NewRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromID(unifiedRoleDefinitionId string, unifiedRoleDefinitionId1 string) RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId {
	return RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId{
		UnifiedRoleDefinitionId:  unifiedRoleDefinitionId,
		UnifiedRoleDefinitionId1: unifiedRoleDefinitionId1,
	}
}

// ParseRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromID parses 'input' into a RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId
func ParseRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromID(input string) (*RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId{}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	if id.UnifiedRoleDefinitionId1, ok = parsed.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromIDInsensitively(input string) (*RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId{}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	if id.UnifiedRoleDefinitionId1, ok = parsed.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromID checks that 'input' can be parsed as a Role Management Device Management Role Definition Inherits Permissions From ID
func ValidateRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Definition Inherits Permissions From ID
func (id RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleDefinitions/%s/inheritsPermissionsFrom/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId, id.UnifiedRoleDefinitionId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Definition Inherits Permissions From ID
func (id RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDeviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("staticRoleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionIdValue"),
		resourceids.StaticSegment("staticInheritsPermissionsFrom", "inheritsPermissionsFrom", "inheritsPermissionsFrom"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId1", "unifiedRoleDefinitionId1Value"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Definition Inherits Permissions From ID
func (id RoleManagementDeviceManagementRoleDefinitionInheritsPermissionsFromId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
		fmt.Sprintf("Unified Role Definition Id 1: %q", id.UnifiedRoleDefinitionId1),
	}
	return fmt.Sprintf("Role Management Device Management Role Definition Inherits Permissions From (%s)", strings.Join(components, "\n"))
}
