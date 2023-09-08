package rolemanagemententerpriseapproledefinitioninheritspermissionsfrom

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleDefinitionId{}

// RoleManagementEnterpriseAppRoleDefinitionId is a struct representing the Resource ID for a Role Management Enterprise App Role Definition
type RoleManagementEnterpriseAppRoleDefinitionId struct {
	RbacApplicationId       string
	UnifiedRoleDefinitionId string
}

// NewRoleManagementEnterpriseAppRoleDefinitionID returns a new RoleManagementEnterpriseAppRoleDefinitionId struct
func NewRoleManagementEnterpriseAppRoleDefinitionID(rbacApplicationId string, unifiedRoleDefinitionId string) RoleManagementEnterpriseAppRoleDefinitionId {
	return RoleManagementEnterpriseAppRoleDefinitionId{
		RbacApplicationId:       rbacApplicationId,
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementEnterpriseAppRoleDefinitionID parses 'input' into a RoleManagementEnterpriseAppRoleDefinitionId
func ParseRoleManagementEnterpriseAppRoleDefinitionID(input string) (*RoleManagementEnterpriseAppRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleDefinitionId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleDefinitionIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleDefinitionId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleDefinitionID checks that 'input' can be parsed as a Role Management Enterprise App Role Definition ID
func ValidateRoleManagementEnterpriseAppRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Definition ID
func (id RoleManagementEnterpriseAppRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Definition ID
func (id RoleManagementEnterpriseAppRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Definition ID
func (id RoleManagementEnterpriseAppRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Definition (%s)", strings.Join(components, "\n"))
}
