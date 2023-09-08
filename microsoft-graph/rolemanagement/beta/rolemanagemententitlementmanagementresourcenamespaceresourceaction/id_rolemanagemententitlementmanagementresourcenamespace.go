package rolemanagemententitlementmanagementresourcenamespaceresourceaction

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEntitlementManagementResourceNamespaceId{}

// RoleManagementEntitlementManagementResourceNamespaceId is a struct representing the Resource ID for a Role Management Entitlement Management Resource Namespace
type RoleManagementEntitlementManagementResourceNamespaceId struct {
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementEntitlementManagementResourceNamespaceID returns a new RoleManagementEntitlementManagementResourceNamespaceId struct
func NewRoleManagementEntitlementManagementResourceNamespaceID(unifiedRbacResourceNamespaceId string) RoleManagementEntitlementManagementResourceNamespaceId {
	return RoleManagementEntitlementManagementResourceNamespaceId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementEntitlementManagementResourceNamespaceID parses 'input' into a RoleManagementEntitlementManagementResourceNamespaceId
func ParseRoleManagementEntitlementManagementResourceNamespaceID(input string) (*RoleManagementEntitlementManagementResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementResourceNamespaceId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementResourceNamespaceIDInsensitively(input string) (*RoleManagementEntitlementManagementResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementResourceNamespaceId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEntitlementManagementResourceNamespaceID checks that 'input' can be parsed as a Role Management Entitlement Management Resource Namespace ID
func ValidateRoleManagementEntitlementManagementResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Resource Namespace ID
func (id RoleManagementEntitlementManagementResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Resource Namespace ID
func (id RoleManagementEntitlementManagementResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEntitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Resource Namespace ID
func (id RoleManagementEntitlementManagementResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Resource Namespace (%s)", strings.Join(components, "\n"))
}
