package rolemanagemententerpriseappresourcenamespace

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppResourceNamespaceId{}

// RoleManagementEnterpriseAppResourceNamespaceId is a struct representing the Resource ID for a Role Management Enterprise App Resource Namespace
type RoleManagementEnterpriseAppResourceNamespaceId struct {
	RbacApplicationId              string
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementEnterpriseAppResourceNamespaceID returns a new RoleManagementEnterpriseAppResourceNamespaceId struct
func NewRoleManagementEnterpriseAppResourceNamespaceID(rbacApplicationId string, unifiedRbacResourceNamespaceId string) RoleManagementEnterpriseAppResourceNamespaceId {
	return RoleManagementEnterpriseAppResourceNamespaceId{
		RbacApplicationId:              rbacApplicationId,
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementEnterpriseAppResourceNamespaceID parses 'input' into a RoleManagementEnterpriseAppResourceNamespaceId
func ParseRoleManagementEnterpriseAppResourceNamespaceID(input string) (*RoleManagementEnterpriseAppResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppResourceNamespaceId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppResourceNamespaceIDInsensitively(input string) (*RoleManagementEnterpriseAppResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppResourceNamespaceId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppResourceNamespaceID checks that 'input' can be parsed as a Role Management Enterprise App Resource Namespace ID
func ValidateRoleManagementEnterpriseAppResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Resource Namespace ID
func (id RoleManagementEnterpriseAppResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Resource Namespace ID
func (id RoleManagementEnterpriseAppResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Resource Namespace ID
func (id RoleManagementEnterpriseAppResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Enterprise App Resource Namespace (%s)", strings.Join(components, "\n"))
}
