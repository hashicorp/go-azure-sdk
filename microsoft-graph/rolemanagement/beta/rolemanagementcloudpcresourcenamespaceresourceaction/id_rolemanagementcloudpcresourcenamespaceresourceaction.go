package rolemanagementcloudpcresourcenamespaceresourceaction

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementCloudPCResourceNamespaceResourceActionId{}

// RoleManagementCloudPCResourceNamespaceResourceActionId is a struct representing the Resource ID for a Role Management Cloud P C Resource Namespace Resource Action
type RoleManagementCloudPCResourceNamespaceResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementCloudPCResourceNamespaceResourceActionID returns a new RoleManagementCloudPCResourceNamespaceResourceActionId struct
func NewRoleManagementCloudPCResourceNamespaceResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementCloudPCResourceNamespaceResourceActionId {
	return RoleManagementCloudPCResourceNamespaceResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementCloudPCResourceNamespaceResourceActionID parses 'input' into a RoleManagementCloudPCResourceNamespaceResourceActionId
func ParseRoleManagementCloudPCResourceNamespaceResourceActionID(input string) (*RoleManagementCloudPCResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementCloudPCResourceNamespaceResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCResourceNamespaceResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCResourceNamespaceResourceActionIDInsensitively(input string) (*RoleManagementCloudPCResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementCloudPCResourceNamespaceResourceActionID checks that 'input' can be parsed as a Role Management Cloud P C Resource Namespace Resource Action ID
func ValidateRoleManagementCloudPCResourceNamespaceResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCResourceNamespaceResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud P C Resource Namespace Resource Action ID
func (id RoleManagementCloudPCResourceNamespaceResourceActionId) ID() string {
	fmtString := "/roleManagement/cloudPC/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud P C Resource Namespace Resource Action ID
func (id RoleManagementCloudPCResourceNamespaceResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticCloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
		resourceids.StaticSegment("staticResourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Cloud P C Resource Namespace Resource Action ID
func (id RoleManagementCloudPCResourceNamespaceResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Cloud P C Resource Namespace Resource Action (%s)", strings.Join(components, "\n"))
}
