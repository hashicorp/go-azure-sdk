package serviceprincipalappmanagementpolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalAppManagementPolicyId{}

// ServicePrincipalAppManagementPolicyId is a struct representing the Resource ID for a Service Principal App Management Policy
type ServicePrincipalAppManagementPolicyId struct {
	ServicePrincipalId    string
	AppManagementPolicyId string
}

// NewServicePrincipalAppManagementPolicyID returns a new ServicePrincipalAppManagementPolicyId struct
func NewServicePrincipalAppManagementPolicyID(servicePrincipalId string, appManagementPolicyId string) ServicePrincipalAppManagementPolicyId {
	return ServicePrincipalAppManagementPolicyId{
		ServicePrincipalId:    servicePrincipalId,
		AppManagementPolicyId: appManagementPolicyId,
	}
}

// ParseServicePrincipalAppManagementPolicyID parses 'input' into a ServicePrincipalAppManagementPolicyId
func ParseServicePrincipalAppManagementPolicyID(input string) (*ServicePrincipalAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalAppManagementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalAppManagementPolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.AppManagementPolicyId, ok = parsed.Parsed["appManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalAppManagementPolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalAppManagementPolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalAppManagementPolicyIDInsensitively(input string) (*ServicePrincipalAppManagementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalAppManagementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalAppManagementPolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.AppManagementPolicyId, ok = parsed.Parsed["appManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalAppManagementPolicyID checks that 'input' can be parsed as a Service Principal App Management Policy ID
func ValidateServicePrincipalAppManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalAppManagementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal App Management Policy ID
func (id ServicePrincipalAppManagementPolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/appManagementPolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.AppManagementPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal App Management Policy ID
func (id ServicePrincipalAppManagementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticAppManagementPolicies", "appManagementPolicies", "appManagementPolicies"),
		resourceids.UserSpecifiedSegment("appManagementPolicyId", "appManagementPolicyIdValue"),
	}
}

// String returns a human-readable description of this Service Principal App Management Policy ID
func (id ServicePrincipalAppManagementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("App Management Policy: %q", id.AppManagementPolicyId),
	}
	return fmt.Sprintf("Service Principal App Management Policy (%s)", strings.Join(components, "\n"))
}
