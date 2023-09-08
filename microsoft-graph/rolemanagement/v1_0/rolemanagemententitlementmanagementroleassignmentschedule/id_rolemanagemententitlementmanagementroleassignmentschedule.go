package rolemanagemententitlementmanagementroleassignmentschedule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEntitlementManagementRoleAssignmentScheduleId{}

// RoleManagementEntitlementManagementRoleAssignmentScheduleId is a struct representing the Resource ID for a Role Management Entitlement Management Role Assignment Schedule
type RoleManagementEntitlementManagementRoleAssignmentScheduleId struct {
	UnifiedRoleAssignmentScheduleId string
}

// NewRoleManagementEntitlementManagementRoleAssignmentScheduleID returns a new RoleManagementEntitlementManagementRoleAssignmentScheduleId struct
func NewRoleManagementEntitlementManagementRoleAssignmentScheduleID(unifiedRoleAssignmentScheduleId string) RoleManagementEntitlementManagementRoleAssignmentScheduleId {
	return RoleManagementEntitlementManagementRoleAssignmentScheduleId{
		UnifiedRoleAssignmentScheduleId: unifiedRoleAssignmentScheduleId,
	}
}

// ParseRoleManagementEntitlementManagementRoleAssignmentScheduleID parses 'input' into a RoleManagementEntitlementManagementRoleAssignmentScheduleId
func ParseRoleManagementEntitlementManagementRoleAssignmentScheduleID(input string) (*RoleManagementEntitlementManagementRoleAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementRoleAssignmentScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementRoleAssignmentScheduleId{}

	if id.UnifiedRoleAssignmentScheduleId, ok = parsed.Parsed["unifiedRoleAssignmentScheduleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleAssignmentScheduleIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleAssignmentScheduleId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleAssignmentScheduleIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementRoleAssignmentScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementRoleAssignmentScheduleId{}

	if id.UnifiedRoleAssignmentScheduleId, ok = parsed.Parsed["unifiedRoleAssignmentScheduleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEntitlementManagementRoleAssignmentScheduleID checks that 'input' can be parsed as a Role Management Entitlement Management Role Assignment Schedule ID
func ValidateRoleManagementEntitlementManagementRoleAssignmentScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleAssignmentScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Assignment Schedule ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleAssignmentSchedules/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleAssignmentScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Assignment Schedule ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEntitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("staticRoleAssignmentSchedules", "roleAssignmentSchedules", "roleAssignmentSchedules"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleId", "unifiedRoleAssignmentScheduleIdValue"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Assignment Schedule ID
func (id RoleManagementEntitlementManagementRoleAssignmentScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Assignment Schedule: %q", id.UnifiedRoleAssignmentScheduleId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Assignment Schedule (%s)", strings.Join(components, "\n"))
}
