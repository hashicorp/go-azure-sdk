package rolemanagementdirectoryresourcenamespaceresourceaction

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDirectoryResourceNamespaceResourceActionId{}

// RoleManagementDirectoryResourceNamespaceResourceActionId is a struct representing the Resource ID for a Role Management Directory Resource Namespace Resource Action
type RoleManagementDirectoryResourceNamespaceResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementDirectoryResourceNamespaceResourceActionID returns a new RoleManagementDirectoryResourceNamespaceResourceActionId struct
func NewRoleManagementDirectoryResourceNamespaceResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementDirectoryResourceNamespaceResourceActionId {
	return RoleManagementDirectoryResourceNamespaceResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementDirectoryResourceNamespaceResourceActionID parses 'input' into a RoleManagementDirectoryResourceNamespaceResourceActionId
func ParseRoleManagementDirectoryResourceNamespaceResourceActionID(input string) (*RoleManagementDirectoryResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDirectoryResourceNamespaceResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryResourceNamespaceResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryResourceNamespaceResourceActionIDInsensitively(input string) (*RoleManagementDirectoryResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDirectoryResourceNamespaceResourceActionID checks that 'input' can be parsed as a Role Management Directory Resource Namespace Resource Action ID
func ValidateRoleManagementDirectoryResourceNamespaceResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryResourceNamespaceResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Resource Namespace Resource Action ID
func (id RoleManagementDirectoryResourceNamespaceResourceActionId) ID() string {
	fmtString := "/roleManagement/directory/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Resource Namespace Resource Action ID
func (id RoleManagementDirectoryResourceNamespaceResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
		resourceids.StaticSegment("staticResourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Directory Resource Namespace Resource Action ID
func (id RoleManagementDirectoryResourceNamespaceResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Directory Resource Namespace Resource Action (%s)", strings.Join(components, "\n"))
}
