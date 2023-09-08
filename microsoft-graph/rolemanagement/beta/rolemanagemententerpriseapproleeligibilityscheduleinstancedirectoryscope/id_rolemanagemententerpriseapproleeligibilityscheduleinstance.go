package rolemanagemententerpriseapproleeligibilityscheduleinstancedirectoryscope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId{}

// RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId is a struct representing the Resource ID for a Role Management Enterprise App Role Eligibility Schedule Instance
type RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId struct {
	RbacApplicationId                        string
	UnifiedRoleEligibilityScheduleInstanceId string
}

// NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceID returns a new RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId struct
func NewRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceID(rbacApplicationId string, unifiedRoleEligibilityScheduleInstanceId string) RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId {
	return RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId{
		RbacApplicationId:                        rbacApplicationId,
		UnifiedRoleEligibilityScheduleInstanceId: unifiedRoleEligibilityScheduleInstanceId,
	}
}

// ParseRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceID parses 'input' into a RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId
func ParseRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceID(input string) (*RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleEligibilityScheduleInstanceId, ok = parsed.Parsed["unifiedRoleEligibilityScheduleInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleInstanceId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleEligibilityScheduleInstanceId, ok = parsed.Parsed["unifiedRoleEligibilityScheduleInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleInstanceId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceID checks that 'input' can be parsed as a Role Management Enterprise App Role Eligibility Schedule Instance ID
func ValidateRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleEligibilityScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Eligibility Schedule Instance ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleEligibilityScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleEligibilityScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Eligibility Schedule Instance ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleEligibilityScheduleInstances", "roleEligibilityScheduleInstances", "roleEligibilityScheduleInstances"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleInstanceId", "unifiedRoleEligibilityScheduleInstanceIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Eligibility Schedule Instance ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Eligibility Schedule Instance: %q", id.UnifiedRoleEligibilityScheduleInstanceId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Eligibility Schedule Instance (%s)", strings.Join(components, "\n"))
}
