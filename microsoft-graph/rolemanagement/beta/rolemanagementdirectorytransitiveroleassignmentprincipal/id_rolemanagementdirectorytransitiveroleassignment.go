package rolemanagementdirectorytransitiveroleassignmentprincipal

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDirectoryTransitiveRoleAssignmentId{}

// RoleManagementDirectoryTransitiveRoleAssignmentId is a struct representing the Resource ID for a Role Management Directory Transitive Role Assignment
type RoleManagementDirectoryTransitiveRoleAssignmentId struct {
	UnifiedRoleAssignmentId string
}

// NewRoleManagementDirectoryTransitiveRoleAssignmentID returns a new RoleManagementDirectoryTransitiveRoleAssignmentId struct
func NewRoleManagementDirectoryTransitiveRoleAssignmentID(unifiedRoleAssignmentId string) RoleManagementDirectoryTransitiveRoleAssignmentId {
	return RoleManagementDirectoryTransitiveRoleAssignmentId{
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementDirectoryTransitiveRoleAssignmentID parses 'input' into a RoleManagementDirectoryTransitiveRoleAssignmentId
func ParseRoleManagementDirectoryTransitiveRoleAssignmentID(input string) (*RoleManagementDirectoryTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryTransitiveRoleAssignmentId{}

	if id.UnifiedRoleAssignmentId, ok = parsed.Parsed["unifiedRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDirectoryTransitiveRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryTransitiveRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryTransitiveRoleAssignmentIDInsensitively(input string) (*RoleManagementDirectoryTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryTransitiveRoleAssignmentId{}

	if id.UnifiedRoleAssignmentId, ok = parsed.Parsed["unifiedRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDirectoryTransitiveRoleAssignmentID checks that 'input' can be parsed as a Role Management Directory Transitive Role Assignment ID
func ValidateRoleManagementDirectoryTransitiveRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryTransitiveRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Transitive Role Assignment ID
func (id RoleManagementDirectoryTransitiveRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/directory/transitiveRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Transitive Role Assignment ID
func (id RoleManagementDirectoryTransitiveRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticTransitiveRoleAssignments", "transitiveRoleAssignments", "transitiveRoleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Role Management Directory Transitive Role Assignment ID
func (id RoleManagementDirectoryTransitiveRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Directory Transitive Role Assignment (%s)", strings.Join(components, "\n"))
}
