package rolemanagementcloudpcroleassignmentappscope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementCloudPCRoleAssignmentAppScopeId{}

// RoleManagementCloudPCRoleAssignmentAppScopeId is a struct representing the Resource ID for a Role Management Cloud P C Role Assignment App Scope
type RoleManagementCloudPCRoleAssignmentAppScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	AppScopeId                      string
}

// NewRoleManagementCloudPCRoleAssignmentAppScopeID returns a new RoleManagementCloudPCRoleAssignmentAppScopeId struct
func NewRoleManagementCloudPCRoleAssignmentAppScopeID(unifiedRoleAssignmentMultipleId string, appScopeId string) RoleManagementCloudPCRoleAssignmentAppScopeId {
	return RoleManagementCloudPCRoleAssignmentAppScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		AppScopeId:                      appScopeId,
	}
}

// ParseRoleManagementCloudPCRoleAssignmentAppScopeID parses 'input' into a RoleManagementCloudPCRoleAssignmentAppScopeId
func ParseRoleManagementCloudPCRoleAssignmentAppScopeID(input string) (*RoleManagementCloudPCRoleAssignmentAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCRoleAssignmentAppScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCRoleAssignmentAppScopeId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.AppScopeId, ok = parsed.Parsed["appScopeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appScopeId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleAssignmentAppScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleAssignmentAppScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleAssignmentAppScopeIDInsensitively(input string) (*RoleManagementCloudPCRoleAssignmentAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCRoleAssignmentAppScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCRoleAssignmentAppScopeId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.AppScopeId, ok = parsed.Parsed["appScopeId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appScopeId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementCloudPCRoleAssignmentAppScopeID checks that 'input' can be parsed as a Role Management Cloud P C Role Assignment App Scope ID
func ValidateRoleManagementCloudPCRoleAssignmentAppScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleAssignmentAppScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud P C Role Assignment App Scope ID
func (id RoleManagementCloudPCRoleAssignmentAppScopeId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleAssignments/%s/appScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.AppScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud P C Role Assignment App Scope ID
func (id RoleManagementCloudPCRoleAssignmentAppScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticCloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("staticRoleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleIdValue"),
		resourceids.StaticSegment("staticAppScopes", "appScopes", "appScopes"),
		resourceids.UserSpecifiedSegment("appScopeId", "appScopeIdValue"),
	}
}

// String returns a human-readable description of this Role Management Cloud P C Role Assignment App Scope ID
func (id RoleManagementCloudPCRoleAssignmentAppScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("App Scope: %q", id.AppScopeId),
	}
	return fmt.Sprintf("Role Management Cloud P C Role Assignment App Scope (%s)", strings.Join(components, "\n"))
}
