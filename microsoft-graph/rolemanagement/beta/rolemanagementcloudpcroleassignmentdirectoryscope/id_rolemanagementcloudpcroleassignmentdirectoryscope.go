package rolemanagementcloudpcroleassignmentdirectoryscope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementCloudPCRoleAssignmentDirectoryScopeId{}

// RoleManagementCloudPCRoleAssignmentDirectoryScopeId is a struct representing the Resource ID for a Role Management Cloud P C Role Assignment Directory Scope
type RoleManagementCloudPCRoleAssignmentDirectoryScopeId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementCloudPCRoleAssignmentDirectoryScopeID returns a new RoleManagementCloudPCRoleAssignmentDirectoryScopeId struct
func NewRoleManagementCloudPCRoleAssignmentDirectoryScopeID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementCloudPCRoleAssignmentDirectoryScopeId {
	return RoleManagementCloudPCRoleAssignmentDirectoryScopeId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementCloudPCRoleAssignmentDirectoryScopeID parses 'input' into a RoleManagementCloudPCRoleAssignmentDirectoryScopeId
func ParseRoleManagementCloudPCRoleAssignmentDirectoryScopeID(input string) (*RoleManagementCloudPCRoleAssignmentDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCRoleAssignmentDirectoryScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCRoleAssignmentDirectoryScopeId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleAssignmentDirectoryScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleAssignmentDirectoryScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleAssignmentDirectoryScopeIDInsensitively(input string) (*RoleManagementCloudPCRoleAssignmentDirectoryScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCRoleAssignmentDirectoryScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCRoleAssignmentDirectoryScopeId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementCloudPCRoleAssignmentDirectoryScopeID checks that 'input' can be parsed as a Role Management Cloud P C Role Assignment Directory Scope ID
func ValidateRoleManagementCloudPCRoleAssignmentDirectoryScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleAssignmentDirectoryScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud P C Role Assignment Directory Scope ID
func (id RoleManagementCloudPCRoleAssignmentDirectoryScopeId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleAssignments/%s/directoryScopes/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud P C Role Assignment Directory Scope ID
func (id RoleManagementCloudPCRoleAssignmentDirectoryScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticCloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("staticRoleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleIdValue"),
		resourceids.StaticSegment("staticDirectoryScopes", "directoryScopes", "directoryScopes"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Role Management Cloud P C Role Assignment Directory Scope ID
func (id RoleManagementCloudPCRoleAssignmentDirectoryScopeId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Cloud P C Role Assignment Directory Scope (%s)", strings.Join(components, "\n"))
}
