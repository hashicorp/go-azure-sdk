package rolemanagementcloudpcroleassignmentprincipal

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementCloudPCRoleAssignmentPrincipalId{}

// RoleManagementCloudPCRoleAssignmentPrincipalId is a struct representing the Resource ID for a Role Management Cloud P C Role Assignment Principal
type RoleManagementCloudPCRoleAssignmentPrincipalId struct {
	UnifiedRoleAssignmentMultipleId string
	DirectoryObjectId               string
}

// NewRoleManagementCloudPCRoleAssignmentPrincipalID returns a new RoleManagementCloudPCRoleAssignmentPrincipalId struct
func NewRoleManagementCloudPCRoleAssignmentPrincipalID(unifiedRoleAssignmentMultipleId string, directoryObjectId string) RoleManagementCloudPCRoleAssignmentPrincipalId {
	return RoleManagementCloudPCRoleAssignmentPrincipalId{
		UnifiedRoleAssignmentMultipleId: unifiedRoleAssignmentMultipleId,
		DirectoryObjectId:               directoryObjectId,
	}
}

// ParseRoleManagementCloudPCRoleAssignmentPrincipalID parses 'input' into a RoleManagementCloudPCRoleAssignmentPrincipalId
func ParseRoleManagementCloudPCRoleAssignmentPrincipalID(input string) (*RoleManagementCloudPCRoleAssignmentPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCRoleAssignmentPrincipalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCRoleAssignmentPrincipalId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementCloudPCRoleAssignmentPrincipalIDInsensitively parses 'input' case-insensitively into a RoleManagementCloudPCRoleAssignmentPrincipalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementCloudPCRoleAssignmentPrincipalIDInsensitively(input string) (*RoleManagementCloudPCRoleAssignmentPrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementCloudPCRoleAssignmentPrincipalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementCloudPCRoleAssignmentPrincipalId{}

	if id.UnifiedRoleAssignmentMultipleId, ok = parsed.Parsed["unifiedRoleAssignmentMultipleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentMultipleId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementCloudPCRoleAssignmentPrincipalID checks that 'input' can be parsed as a Role Management Cloud P C Role Assignment Principal ID
func ValidateRoleManagementCloudPCRoleAssignmentPrincipalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementCloudPCRoleAssignmentPrincipalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Cloud P C Role Assignment Principal ID
func (id RoleManagementCloudPCRoleAssignmentPrincipalId) ID() string {
	fmtString := "/roleManagement/cloudPC/roleAssignments/%s/principals/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentMultipleId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Cloud P C Role Assignment Principal ID
func (id RoleManagementCloudPCRoleAssignmentPrincipalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticCloudPC", "cloudPC", "cloudPC"),
		resourceids.StaticSegment("staticRoleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentMultipleId", "unifiedRoleAssignmentMultipleIdValue"),
		resourceids.StaticSegment("staticPrincipals", "principals", "principals"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Role Management Cloud P C Role Assignment Principal ID
func (id RoleManagementCloudPCRoleAssignmentPrincipalId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Multiple: %q", id.UnifiedRoleAssignmentMultipleId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Role Management Cloud P C Role Assignment Principal (%s)", strings.Join(components, "\n"))
}
