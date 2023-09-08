package rolemanagementdevicemanagementroleassignmentappscope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDeviceManagementRoleAssignmentAppScopeId{}

// RoleManagementDeviceManagementRoleAssignmentAppScopeId is a struct representing the Resource ID for a Role Management Device Management Role Assignment App Scope
type RoleManagementDeviceManagementRoleAssignmentAppScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	AppScopeId                      string
}

// NewRoleManagementDeviceManagementRoleAssignmentAppScopeID returns a new RoleManagementDeviceManagementRoleAssignmentAppScopeId struct
func NewRoleManagementDeviceManagementRoleAssignmentAppScopeID(unifiedRoleAssignmentMultipleId string, appScopeId string) RoleManagementDeviceManagementRoleAssignmentAppScopeId {
	return RoleManagementDeviceManagementRoleAssignmentAppScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		AppScopeId:                      appScopeId,
	}
}

// ParseRoleManagementDeviceManagementRoleAssignmentAppScopeID parses 'input' into a RoleManagementDeviceManagementRoleAssignmentAppScopeId
func ParseRoleManagementDeviceManagementRoleAssignmentAppScopeID(input string) (*RoleManagementDeviceManagementRoleAssignmentAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementRoleAssignmentAppScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementRoleAssignmentAppScopeId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.AppScopeId, ok = parsed.Parsed["appScopeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appScopeId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleAssignmentAppScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleAssignmentAppScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleAssignmentAppScopeIDInsensitively(input string) (*RoleManagementDeviceManagementRoleAssignmentAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementRoleAssignmentAppScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementRoleAssignmentAppScopeId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.AppScopeId, ok = parsed.Parsed["appScopeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appScopeId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDeviceManagementRoleAssignmentAppScopeID checks that 'input' can be parsed as a Role Management Device Management Role Assignment App Scope ID
func ValidateRoleManagementDeviceManagementRoleAssignmentAppScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleAssignmentAppScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Assignment App Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentAppScopeId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleAssignments/%s/appScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.AppScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Assignment App Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentAppScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDeviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("staticRoleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleIdValue"),
		resourceids.StaticSegment("staticAppScopes", "appScopes", "appScopes"),
		resourceids.UserSpecifiedSegment("appScopeId", "appScopeIdValue"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Assignment App Scope ID
func (id RoleManagementDeviceManagementRoleAssignmentAppScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("App Scope: %q", id.AppScopeId),
	}
	return fmt.Sprintf("Role Management Device Management Role Assignment App Scope (%s)", strings.Join(components, "\n"))
}
