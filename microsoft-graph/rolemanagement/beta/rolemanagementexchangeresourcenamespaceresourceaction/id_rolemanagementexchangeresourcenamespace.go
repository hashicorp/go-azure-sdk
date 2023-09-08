package rolemanagementexchangeresourcenamespaceresourceaction

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementExchangeResourceNamespaceId{}

// RoleManagementExchangeResourceNamespaceId is a struct representing the Resource ID for a Role Management Exchange Resource Namespace
type RoleManagementExchangeResourceNamespaceId struct {
	UnifiedRbacResourceNamespaceId string
}

// NewRoleManagementExchangeResourceNamespaceID returns a new RoleManagementExchangeResourceNamespaceId struct
func NewRoleManagementExchangeResourceNamespaceID(unifiedRbacResourceNamespaceId string) RoleManagementExchangeResourceNamespaceId {
	return RoleManagementExchangeResourceNamespaceId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
	}
}

// ParseRoleManagementExchangeResourceNamespaceID parses 'input' into a RoleManagementExchangeResourceNamespaceId
func ParseRoleManagementExchangeResourceNamespaceID(input string) (*RoleManagementExchangeResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementExchangeResourceNamespaceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementExchangeResourceNamespaceId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementExchangeResourceNamespaceIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeResourceNamespaceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeResourceNamespaceIDInsensitively(input string) (*RoleManagementExchangeResourceNamespaceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementExchangeResourceNamespaceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementExchangeResourceNamespaceId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementExchangeResourceNamespaceID checks that 'input' can be parsed as a Role Management Exchange Resource Namespace ID
func ValidateRoleManagementExchangeResourceNamespaceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeResourceNamespaceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Resource Namespace ID
func (id RoleManagementExchangeResourceNamespaceId) ID() string {
	fmtString := "/roleManagement/exchange/resourceNamespaces/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Resource Namespace ID
func (id RoleManagementExchangeResourceNamespaceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticExchange", "exchange", "exchange"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
	}
}

// String returns a human-readable description of this Role Management Exchange Resource Namespace ID
func (id RoleManagementExchangeResourceNamespaceId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
	}
	return fmt.Sprintf("Role Management Exchange Resource Namespace (%s)", strings.Join(components, "\n"))
}
