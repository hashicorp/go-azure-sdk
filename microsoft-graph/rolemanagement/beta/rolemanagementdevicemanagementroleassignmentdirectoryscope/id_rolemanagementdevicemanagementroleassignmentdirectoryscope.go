package rolemanagementdevicemanagementroleassignmentdirectoryscope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId{}

// RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId is a struct representing the Resource ID for a Role Management Device Management Role Assignment Directory Scope
type RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementDeviceManagementRoleAssignmentDirectoryScopeID returns a new RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId struct
func NewRoleManagementDeviceManagementRoleAssignmentDirectoryScopeID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId {
	return RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementDeviceManagementRoleAssignmentDirectoryScopeID parses 'input' into a RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId
func ParseRoleManagementDeviceManagementRoleAssignmentDirectoryScopeID(input string) (*RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleAssignmentDirectoryScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleAssignmentDirectoryScopeIDInsensitively(input string) (*RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDeviceManagementRoleAssignmentDirectoryScopeID checks that 'input' can be parsed as a Role Management Device Management Role Assignment Directory Scope ID
func ValidateRoleManagementDeviceManagementRoleAssignmentDirectoryScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleAssignmentDirectoryScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Assignment Directory Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleAssignments/%s/directoryScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Assignment Directory Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDeviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("staticRoleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleIdValue"),
		resourceids.StaticSegment("staticDirectoryScopes", "directoryScopes", "directoryScopes"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Assignment Directory Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentDirectoryScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Device Management Role Assignment Directory Scope (%s)", strings.Join(components, "\n"))
}
