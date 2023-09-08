package serviceprincipalapproleassignedto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalAppRoleAssignedToId{}

// ServicePrincipalAppRoleAssignedToId is a struct representing the Resource ID for a Service Principal App Role Assigned To
type ServicePrincipalAppRoleAssignedToId struct {
	ServicePrincipalId  string
	AppRoleAssignmentId string
}

// NewServicePrincipalAppRoleAssignedToID returns a new ServicePrincipalAppRoleAssignedToId struct
func NewServicePrincipalAppRoleAssignedToID(servicePrincipalId string, appRoleAssignmentId string) ServicePrincipalAppRoleAssignedToId {
	return ServicePrincipalAppRoleAssignedToId{
		ServicePrincipalId:  servicePrincipalId,
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseServicePrincipalAppRoleAssignedToID parses 'input' into a ServicePrincipalAppRoleAssignedToId
func ParseServicePrincipalAppRoleAssignedToID(input string) (*ServicePrincipalAppRoleAssignedToId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalAppRoleAssignedToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalAppRoleAssignedToId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalAppRoleAssignedToIDInsensitively parses 'input' case-insensitively into a ServicePrincipalAppRoleAssignedToId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalAppRoleAssignedToIDInsensitively(input string) (*ServicePrincipalAppRoleAssignedToId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalAppRoleAssignedToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalAppRoleAssignedToId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalAppRoleAssignedToID checks that 'input' can be parsed as a Service Principal App Role Assigned To ID
func ValidateServicePrincipalAppRoleAssignedToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalAppRoleAssignedToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal App Role Assigned To ID
func (id ServicePrincipalAppRoleAssignedToId) ID() string {
	fmtString := "/servicePrincipals/%s/appRoleAssignedTo/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal App Role Assigned To ID
func (id ServicePrincipalAppRoleAssignedToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticAppRoleAssignedTo", "appRoleAssignedTo", "appRoleAssignedTo"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Service Principal App Role Assigned To ID
func (id ServicePrincipalAppRoleAssignedToId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("Service Principal App Role Assigned To (%s)", strings.Join(components, "\n"))
}
