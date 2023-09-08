package rolemanagementdevicemanagementroleassignmentprincipal

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDeviceManagementRoleAssignmentPrincipalId{}

// RoleManagementDeviceManagementRoleAssignmentPrincipalId is a struct representing the Resource ID for a Role Management Device Management Role Assignment Principal
type RoleManagementDeviceManagementRoleAssignmentPrincipalId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementDeviceManagementRoleAssignmentPrincipalID returns a new RoleManagementDeviceManagementRoleAssignmentPrincipalId struct
func NewRoleManagementDeviceManagementRoleAssignmentPrincipalID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementDeviceManagementRoleAssignmentPrincipalId {
	return RoleManagementDeviceManagementRoleAssignmentPrincipalId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementDeviceManagementRoleAssignmentPrincipalID parses 'input' into a RoleManagementDeviceManagementRoleAssignmentPrincipalId
func ParseRoleManagementDeviceManagementRoleAssignmentPrincipalID(input string) (*RoleManagementDeviceManagementRoleAssignmentPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementRoleAssignmentPrincipalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementRoleAssignmentPrincipalId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementRoleAssignmentPrincipalIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementRoleAssignmentPrincipalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementRoleAssignmentPrincipalIDInsensitively(input string) (*RoleManagementDeviceManagementRoleAssignmentPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementRoleAssignmentPrincipalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementRoleAssignmentPrincipalId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDeviceManagementRoleAssignmentPrincipalID checks that 'input' can be parsed as a Role Management Device Management Role Assignment Principal ID
func ValidateRoleManagementDeviceManagementRoleAssignmentPrincipalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementRoleAssignmentPrincipalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Role Assignment Principal ID
func (id RoleManagementDeviceManagementRoleAssignmentPrincipalId) ID() string {
	fmtString := "/roleManagement/deviceManagement/roleAssignments/%s/principals/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Role Assignment Principal ID
func (id RoleManagementDeviceManagementRoleAssignmentPrincipalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDeviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("staticRoleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleIdValue"),
		resourceids.StaticSegment("staticPrincipals", "principals", "principals"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Role Management Device Management Role Assignment Principal ID
func (id RoleManagementDeviceManagementRoleAssignmentPrincipalId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Device Management Role Assignment Principal (%s)", strings.Join(components, "\n"))
}
