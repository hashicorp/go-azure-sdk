package rolemanagementdevicemanagementresourcenamespaceresourceactionauthenticationcontext

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDeviceManagementResourceNamespaceResourceActionId{}

// RoleManagementDeviceManagementResourceNamespaceResourceActionId is a struct representing the Resource ID for a Role Management Device Management Resource Namespace Resource Action
type RoleManagementDeviceManagementResourceNamespaceResourceActionId struct {
	UnifiedRbacResourceNamespaceId string
	UnifiedRbacResourceActionId    string
}

// NewRoleManagementDeviceManagementResourceNamespaceResourceActionID returns a new RoleManagementDeviceManagementResourceNamespaceResourceActionId struct
func NewRoleManagementDeviceManagementResourceNamespaceResourceActionID(unifiedRbacResourceNamespaceId string, unifiedRbacResourceActionId string) RoleManagementDeviceManagementResourceNamespaceResourceActionId {
	return RoleManagementDeviceManagementResourceNamespaceResourceActionId{
		UnifiedRbacResourceNamespaceId: unifiedRbacResourceNamespaceId,
		UnifiedRbacResourceActionId:    unifiedRbacResourceActionId,
	}
}

// ParseRoleManagementDeviceManagementResourceNamespaceResourceActionID parses 'input' into a RoleManagementDeviceManagementResourceNamespaceResourceActionId
func ParseRoleManagementDeviceManagementResourceNamespaceResourceActionID(input string) (*RoleManagementDeviceManagementResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDeviceManagementResourceNamespaceResourceActionIDInsensitively parses 'input' case-insensitively into a RoleManagementDeviceManagementResourceNamespaceResourceActionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDeviceManagementResourceNamespaceResourceActionIDInsensitively(input string) (*RoleManagementDeviceManagementResourceNamespaceResourceActionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDeviceManagementResourceNamespaceResourceActionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDeviceManagementResourceNamespaceResourceActionId{}

	if id.UnifiedRbacResourceNamespaceId, ok = parsed.Parsed["unifiedRbacResourceNamespaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceNamespaceId", *parsed)
	}

	if id.UnifiedRbacResourceActionId, ok = parsed.Parsed["unifiedRbacResourceActionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRbacResourceActionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDeviceManagementResourceNamespaceResourceActionID checks that 'input' can be parsed as a Role Management Device Management Resource Namespace Resource Action ID
func ValidateRoleManagementDeviceManagementResourceNamespaceResourceActionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDeviceManagementResourceNamespaceResourceActionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Device Management Resource Namespace Resource Action ID
func (id RoleManagementDeviceManagementResourceNamespaceResourceActionId) ID() string {
	fmtString := "/roleManagement/deviceManagement/resourceNamespaces/%s/resourceActions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRbacResourceNamespaceId, id.UnifiedRbacResourceActionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Device Management Resource Namespace Resource Action ID
func (id RoleManagementDeviceManagementResourceNamespaceResourceActionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDeviceManagement", "deviceManagement", "deviceManagement"),
		resourceids.StaticSegment("staticResourceNamespaces", "resourceNamespaces", "resourceNamespaces"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceNamespaceId", "unifiedRbacResourceNamespaceIdValue"),
		resourceids.StaticSegment("staticResourceActions", "resourceActions", "resourceActions"),
		resourceids.UserSpecifiedSegment("unifiedRbacResourceActionId", "unifiedRbacResourceActionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Device Management Resource Namespace Resource Action ID
func (id RoleManagementDeviceManagementResourceNamespaceResourceActionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Rbac Resource Namespace: %q", id.UnifiedRbacResourceNamespaceId),
		fmt.Sprintf("Unified Rbac Resource Action: %q", id.UnifiedRbacResourceActionId),
	}
	return fmt.Sprintf("Role Management Device Management Resource Namespace Resource Action (%s)", strings.Join(components, "\n"))
}
