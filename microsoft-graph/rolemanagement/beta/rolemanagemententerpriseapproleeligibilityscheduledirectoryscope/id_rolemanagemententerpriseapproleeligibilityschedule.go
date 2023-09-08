package rolemanagemententerpriseapproleeligibilityscheduledirectoryscope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleEligibilityScheduleId{}

// RoleManagementEnterpriseAppRoleEligibilityScheduleId is a struct representing the Resource ID for a Role Management Enterprise App Role Eligibility Schedule
type RoleManagementEnterpriseAppRoleEligibilityScheduleId struct {
	RbacApplicationId                string
	UnifiedRoleEligibilityScheduleId string
}

// NewRoleManagementEnterpriseAppRoleEligibilityScheduleID returns a new RoleManagementEnterpriseAppRoleEligibilityScheduleId struct
func NewRoleManagementEnterpriseAppRoleEligibilityScheduleID(rbacApplicationId string, unifiedRoleEligibilityScheduleId string) RoleManagementEnterpriseAppRoleEligibilityScheduleId {
	return RoleManagementEnterpriseAppRoleEligibilityScheduleId{
		RbacApplicationId:                rbacApplicationId,
		UnifiedRoleEligibilityScheduleId: unifiedRoleEligibilityScheduleId,
	}
}

// ParseRoleManagementEnterpriseAppRoleEligibilityScheduleID parses 'input' into a RoleManagementEnterpriseAppRoleEligibilityScheduleId
func ParseRoleManagementEnterpriseAppRoleEligibilityScheduleID(input string) (*RoleManagementEnterpriseAppRoleEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleEligibilityScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleEligibilityScheduleId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleEligibilityScheduleId, ok = parsed.Parsed["unifiedRoleEligibilityScheduleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleEligibilityScheduleIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleEligibilityScheduleId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleEligibilityScheduleIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleEligibilityScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleEligibilityScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleEligibilityScheduleId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleEligibilityScheduleId, ok = parsed.Parsed["unifiedRoleEligibilityScheduleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleEligibilityScheduleID checks that 'input' can be parsed as a Role Management Enterprise App Role Eligibility Schedule ID
func ValidateRoleManagementEnterpriseAppRoleEligibilityScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleEligibilityScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Eligibility Schedule ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleEligibilitySchedules/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleEligibilityScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Eligibility Schedule ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleEligibilitySchedules", "roleEligibilitySchedules", "roleEligibilitySchedules"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleId", "unifiedRoleEligibilityScheduleIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Eligibility Schedule ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Eligibility Schedule: %q", id.UnifiedRoleEligibilityScheduleId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Eligibility Schedule (%s)", strings.Join(components, "\n"))
}
