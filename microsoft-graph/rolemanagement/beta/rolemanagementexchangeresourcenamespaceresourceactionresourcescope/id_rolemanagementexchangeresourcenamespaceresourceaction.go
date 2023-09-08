package rolemanagementexchangeresourcenamespaceresourceactionresourcescope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementExchangeResourceNamespaceResourceActionId{}

// RoleManagementExchangeResourceNamespaceResourceActionId is a struct representing the Resource ID for a Role Management Exchange Resource Namespace Resource Action
type RoleManagementExchangeResourceNamespaceResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementExchangeResourceNamespaceResourceActionID returns a new RoleManagementExchangeResourceNamespaceResourceActionId struct
func NewRoleManagementExchangeResourceNamespaceResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementExchangeResourceNamespaceResourceActionId {
	return RoleManagementExchangeResourceNamespaceResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementExchangeResourceNamespaceResourceActionID parses 'input' into a RoleManagementExchangeResourceNamespaceResourceActionId
func ParseRoleManagementExchangeResourceNamespaceResourceActionID(input string) (*RoleManagementExchangeResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementExchangeResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementExchangeResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementExchangeResourceNamespaceResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeResourceNamespaceResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeResourceNamespaceResourceActionIDInsensitively(input string) (*RoleManagementExchangeResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementExchangeResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementExchangeResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementExchangeResourceNamespaceResourceActionID checks that 'input' can be parsed as a Role Management Exchange Resource Namespace Resource Action ID
func ValidateRoleManagementExchangeResourceNamespaceResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeResourceNamespaceResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Resource Namespace Resource Action ID
func (id RoleManagementExchangeResourceNamespaceResourceActionId) ID() string {
	fmtString := "/roleManagement/exchange/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Resource Namespace Resource Action ID
func (id RoleManagementExchangeResourceNamespaceResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticExchange", "exchange", "exchange"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
		resourceids.StaticSegment("staticResourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Exchange Resource Namespace Resource Action ID
func (id RoleManagementExchangeResourceNamespaceResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Exchange Resource Namespace Resource Action (%s)", strings.Join(components, "\n"))
}
