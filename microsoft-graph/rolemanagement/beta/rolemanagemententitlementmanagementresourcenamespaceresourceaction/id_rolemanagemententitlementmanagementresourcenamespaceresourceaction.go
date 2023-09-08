package rolemanagemententitlementmanagementresourcenamespaceresourceaction

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEntitlementManagementResourceNamespaceResourceActionId{}

// RoleManagementEntitlementManagementResourceNamespaceResourceActionId is a struct representing the Resource ID for a Role Management Entitlement Management Resource Namespace Resource Action
type RoleManagementEntitlementManagementResourceNamespaceResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementEntitlementManagementResourceNamespaceResourceActionID returns a new RoleManagementEntitlementManagementResourceNamespaceResourceActionId struct
func NewRoleManagementEntitlementManagementResourceNamespaceResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementEntitlementManagementResourceNamespaceResourceActionId {
	return RoleManagementEntitlementManagementResourceNamespaceResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementEntitlementManagementResourceNamespaceResourceActionID parses 'input' into a RoleManagementEntitlementManagementResourceNamespaceResourceActionId
func ParseRoleManagementEntitlementManagementResourceNamespaceResourceActionID(input string) (*RoleManagementEntitlementManagementResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementResourceNamespaceResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementResourceNamespaceResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementResourceNamespaceResourceActionIDInsensitively(input string) (*RoleManagementEntitlementManagementResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEntitlementManagementResourceNamespaceResourceActionID checks that 'input' can be parsed as a Role Management Entitlement Management Resource Namespace Resource Action ID
func ValidateRoleManagementEntitlementManagementResourceNamespaceResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementResourceNamespaceResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Resource Namespace Resource Action ID
func (id RoleManagementEntitlementManagementResourceNamespaceResourceActionId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Resource Namespace Resource Action ID
func (id RoleManagementEntitlementManagementResourceNamespaceResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEntitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
		resourceids.StaticSegment("staticResourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Resource Namespace Resource Action ID
func (id RoleManagementEntitlementManagementResourceNamespaceResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Resource Namespace Resource Action (%s)", strings.Join(components, "\n"))
}
