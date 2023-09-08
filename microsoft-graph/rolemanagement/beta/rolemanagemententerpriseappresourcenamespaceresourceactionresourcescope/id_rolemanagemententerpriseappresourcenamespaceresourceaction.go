package rolemanagemententerpriseappresourcenamespaceresourceactionresourcescope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppResourceNamespaceResourceActionId{}

// RoleManagementEnterpriseAppResourceNamespaceResourceActionId is a struct representing the Resource ID for a Role Management Enterprise App Resource Namespace Resource Action
type RoleManagementEnterpriseAppResourceNamespaceResourceActionId struct {
	RbacApplicationId              string
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementEnterpriseAppResourceNamespaceResourceActionID returns a new RoleManagementEnterpriseAppResourceNamespaceResourceActionId struct
func NewRoleManagementEnterpriseAppResourceNamespaceResourceActionID(rbacApplicationId string, unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementEnterpriseAppResourceNamespaceResourceActionId {
	return RoleManagementEnterpriseAppResourceNamespaceResourceActionId{
		RbacApplicationId:              rbacApplicationId,
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementEnterpriseAppResourceNamespaceResourceActionID parses 'input' into a RoleManagementEnterpriseAppResourceNamespaceResourceActionId
func ParseRoleManagementEnterpriseAppResourceNamespaceResourceActionID(input string) (*RoleManagementEnterpriseAppResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppResourceNamespaceResourceActionId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppResourceNamespaceResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppResourceNamespaceResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppResourceNamespaceResourceActionIDInsensitively(input string) (*RoleManagementEnterpriseAppResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppResourceNamespaceResourceActionId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppResourceNamespaceResourceActionID checks that 'input' can be parsed as a Role Management Enterprise App Resource Namespace Resource Action ID
func ValidateRoleManagementEnterpriseAppResourceNamespaceResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppResourceNamespaceResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Resource Namespace Resource Action ID
func (id RoleManagementEnterpriseAppResourceNamespaceResourceActionId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Resource Namespace Resource Action ID
func (id RoleManagementEnterpriseAppResourceNamespaceResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
		resourceids.StaticSegment("staticResourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Resource Namespace Resource Action ID
func (id RoleManagementEnterpriseAppResourceNamespaceResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Enterprise App Resource Namespace Resource Action (%s)", strings.Join(components, "\n"))
}
