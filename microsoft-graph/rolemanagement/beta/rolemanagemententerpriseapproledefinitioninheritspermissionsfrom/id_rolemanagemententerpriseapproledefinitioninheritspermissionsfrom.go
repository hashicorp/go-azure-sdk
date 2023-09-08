package rolemanagemententerpriseapproledefinitioninheritspermissionsfrom

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId{}

// RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId is a struct representing the Resource ID for a Role Management Enterprise App Role Definition Inherits Permissions From
type RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId struct {
	RbacApplicationId        string
	UnifiedRoleDefinitionId  string
	UnifiedRoleDefinitionId1 string
}

// NewRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromID returns a new RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId struct
func NewRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromID(rbacApplicationId string, unifiedRoleDefinitionId string, unifiedRoleDefinitionId1 string) RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId {
	return RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId{
		RbacApplicationId:        rbacApplicationId,
		UnifiedRoleDefinitionId:  unifiedRoleDefinitionId,
		UnifiedRoleDefinitionId1: unifiedRoleDefinitionId1,
	}
}

// ParseRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromID parses 'input' into a RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId
func ParseRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromID(input string) (*RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	if id.UnifiedRoleDefinitionId1, ok = parsed.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	if id.UnifiedRoleDefinitionId1, ok = parsed.Parsed["unifiedRoleDefinitionId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId1", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromID checks that 'input' can be parsed as a Role Management Enterprise App Role Definition Inherits Permissions From ID
func ValidateRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Definition Inherits Permissions From ID
func (id RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleDefinitions/%s/inheritsPermissionsFrom/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleDefinitionId, id.UnifiedRoleDefinitionId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Definition Inherits Permissions From ID
func (id RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionIdValue"),
		resourceids.StaticSegment("staticInheritsPermissionsFrom", "inheritsPermissionsFrom", "inheritsPermissionsFrom"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId1", "unifiedRoleDefinitionId1Value"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Definition Inherits Permissions From ID
func (id RoleManagementEnterpriseAppRoleDefinitionInheritsPermissionsFromId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
		fmt.Sprintf("Unified Role Definition Id 1: %q", id.UnifiedRoleDefinitionId1),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Definition Inherits Permissions From (%s)", strings.Join(components, "\n"))
}
