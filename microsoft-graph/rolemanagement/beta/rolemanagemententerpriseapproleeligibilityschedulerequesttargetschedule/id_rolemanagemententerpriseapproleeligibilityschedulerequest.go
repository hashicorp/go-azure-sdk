package rolemanagemententerpriseapproleeligibilityschedulerequesttargetschedule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId{}

// RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId is a struct representing the Resource ID for a Role Management Enterprise App Role Eligibility Schedule Request
type RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId struct {
	RbacApplicationId                       string
	UnifiedRoleEligibilityScheduleRequestId string
}

// NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestID returns a new RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId struct
func NewRoleManagementEnterpriseAppRoleEligibilityScheduleRequestID(rbacApplicationId string, unifiedRoleEligibilityScheduleRequestId string) RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId {
	return RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId{
		RbacApplicationId:                       rbacApplicationId,
		UnifiedRoleEligibilityScheduleRequestId: unifiedRoleEligibilityScheduleRequestId,
	}
}

// ParseRoleManagementEnterpriseAppRoleEligibilityScheduleRequestID parses 'input' into a RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId
func ParseRoleManagementEnterpriseAppRoleEligibilityScheduleRequestID(input string) (*RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleEligibilityScheduleRequestId, ok = parsed.Parsed["unifiedRoleEligibilityScheduleRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleRequestId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleEligibilityScheduleRequestIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleEligibilityScheduleRequestIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleEligibilityScheduleRequestId, ok = parsed.Parsed["unifiedRoleEligibilityScheduleRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleEligibilityScheduleRequestId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleEligibilityScheduleRequestID checks that 'input' can be parsed as a Role Management Enterprise App Role Eligibility Schedule Request ID
func ValidateRoleManagementEnterpriseAppRoleEligibilityScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleEligibilityScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Eligibility Schedule Request ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleEligibilityScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleEligibilityScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Eligibility Schedule Request ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleEligibilityScheduleRequests", "roleEligibilityScheduleRequests", "roleEligibilityScheduleRequests"),
		resourceids.UserSpecifiedSegment("unifiedRoleEligibilityScheduleRequestId", "unifiedRoleEligibilityScheduleRequestIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Eligibility Schedule Request ID
func (id RoleManagementEnterpriseAppRoleEligibilityScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Eligibility Schedule Request: %q", id.UnifiedRoleEligibilityScheduleRequestId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Eligibility Schedule Request (%s)", strings.Join(components, "\n"))
}
