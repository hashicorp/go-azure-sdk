package serviceprincipalapproleassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalAppRoleAssignmentId{}

// ServicePrincipalAppRoleAssignmentId is a struct representing the Resource ID for a Service Principal App Role Assignment
type ServicePrincipalAppRoleAssignmentId struct {
	ServicePrincipalId  string
	AppRoleAssignmentId string
}

// NewServicePrincipalAppRoleAssignmentID returns a new ServicePrincipalAppRoleAssignmentId struct
func NewServicePrincipalAppRoleAssignmentID(servicePrincipalId string, appRoleAssignmentId string) ServicePrincipalAppRoleAssignmentId {
	return ServicePrincipalAppRoleAssignmentId{
		ServicePrincipalId:  servicePrincipalId,
		AppRoleAssignmentId: appRoleAssignmentId,
	}
}

// ParseServicePrincipalAppRoleAssignmentID parses 'input' into a ServicePrincipalAppRoleAssignmentId
func ParseServicePrincipalAppRoleAssignmentID(input string) (*ServicePrincipalAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalAppRoleAssignmentId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a ServicePrincipalAppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalAppRoleAssignmentIDInsensitively(input string) (*ServicePrincipalAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalAppRoleAssignmentId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.AppRoleAssignmentId, ok = parsed.Parsed["appRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalAppRoleAssignmentID checks that 'input' can be parsed as a Service Principal App Role Assignment ID
func ValidateServicePrincipalAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal App Role Assignment ID
func (id ServicePrincipalAppRoleAssignmentId) ID() string {
	fmtString := "/servicePrincipals/%s/appRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.AppRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal App Role Assignment ID
func (id ServicePrincipalAppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticAppRoleAssignments", "appRoleAssignments", "appRoleAssignments"),
		resourceids.UserSpecifiedSegment("appRoleAssignmentId", "appRoleAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Service Principal App Role Assignment ID
func (id ServicePrincipalAppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("App Role Assignment: %q", id.AppRoleAssignmentId),
	}
	return fmt.Sprintf("Service Principal App Role Assignment (%s)", strings.Join(components, "\n"))
}
