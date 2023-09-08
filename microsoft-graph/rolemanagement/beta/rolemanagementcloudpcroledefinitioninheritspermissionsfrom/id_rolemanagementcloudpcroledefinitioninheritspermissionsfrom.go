package rolemanagementcloudpcroledefinitioninheritspermissionsfrom

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{}

// RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId is a struct representing the Resource ID for a Role Management Cloud P C Role Definition Inherits Permissions From
type RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId struct {
	UnifiedRoleDefinitionId  string
	UnifiedRoleDefinitionId1 string
}

// NewRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID returns a new RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId struct
func NewRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID(unifiedRoleDefinitionId string, unifiedRoleDefinitionId1 string) RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId {
	return RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{
		UnifiedRoleDefinitionId:  unifiedRoleDefinitionId,
		UnifiedRoleDefinitionId1: unifiedRoleDefinitionId1,
	}
}

// ParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID parses 'input' into a RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId
func ParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID(input string) (*RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	if id.UnifiedRoleDefinitionId1, ok = parsed.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromIDInsensitively(input string) (*RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId{}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	if id.UnifiedRoleDefinitionId1, ok = parsed.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID checks that 'input' can be parsed as a Role Management Cloud P C Role Definition Inherits Permissions From ID
func ValidateRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleDefinitionInheritsPermissionsFromID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud P C Role Definition Inherits Permissions From ID
func (id RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleDefinitions/%s/inheritsPermissionsFrom/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId, id.UnifiedRoleDefinitionId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud P C Role Definition Inherits Permissions From ID
func (id RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticCloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("staticRoleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionIdValue"),
		resourceids.StaticSegment("staticInheritsPermissionsFrom", "inheritsPermissionsFrom", "inheritsPermissionsFrom"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId1", "unifiedRoleDefinitionId1Value"),
	}
}

// String returns a human-readable description of this Role Management Cloud P C Role Definition Inherits Permissions From ID
func (id RoleManagementCloudPCRoleDefinitionInheritsPermissionsFromId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
		fmt.Sprintf("Unified Role Definition Id 1: %q", id.UnifiedRoleDefinitionId1),
	}
	return fmt.Sprintf("Role Management Cloud P C Role Definition Inherits Permissions From (%s)", strings.Join(components, "\n"))
}
