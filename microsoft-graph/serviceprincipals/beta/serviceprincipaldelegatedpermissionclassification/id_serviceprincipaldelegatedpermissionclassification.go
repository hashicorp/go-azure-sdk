package serviceprincipaldelegatedpermissionclassification

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalDelegatedPermissionClassificationId{}

// ServicePrincipalDelegatedPermissionClassificationId is a struct representing the Resource ID for a Service Principal Delegated Permission Classification
type ServicePrincipalDelegatedPermissionClassificationId struct {
	ServicePrincipalId                  string
	DelegatedPermissionClassificationId string
}

// NewServicePrincipalDelegatedPermissionClassificationID returns a new ServicePrincipalDelegatedPermissionClassificationId struct
func NewServicePrincipalDelegatedPermissionClassificationID(servicePrincipalId string, delegatedPermissionClassificationId string) ServicePrincipalDelegatedPermissionClassificationId {
	return ServicePrincipalDelegatedPermissionClassificationId{
		ServicePrincipalId:                  servicePrincipalId,
		DelegatedPermissionClassificationId: delegatedPermissionClassificationId,
	}
}

// ParseServicePrincipalDelegatedPermissionClassificationID parses 'input' into a ServicePrincipalDelegatedPermissionClassificationId
func ParseServicePrincipalDelegatedPermissionClassificationID(input string) (*ServicePrincipalDelegatedPermissionClassificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalDelegatedPermissionClassificationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalDelegatedPermissionClassificationId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DelegatedPermissionClassificationId, ok = parsed.Parsed["delegatedPermissionClassificationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "delegatedPermissionClassificationId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalDelegatedPermissionClassificationIDInsensitively parses 'input' case-insensitively into a ServicePrincipalDelegatedPermissionClassificationId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalDelegatedPermissionClassificationIDInsensitively(input string) (*ServicePrincipalDelegatedPermissionClassificationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalDelegatedPermissionClassificationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalDelegatedPermissionClassificationId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DelegatedPermissionClassificationId, ok = parsed.Parsed["delegatedPermissionClassificationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "delegatedPermissionClassificationId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalDelegatedPermissionClassificationID checks that 'input' can be parsed as a Service Principal Delegated Permission Classification ID
func ValidateServicePrincipalDelegatedPermissionClassificationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalDelegatedPermissionClassificationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Delegated Permission Classification ID
func (id ServicePrincipalDelegatedPermissionClassificationId) ID() string {
	fmtString := "/servicePrincipals/%s/delegatedPermissionClassifications/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DelegatedPermissionClassificationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Delegated Permission Classification ID
func (id ServicePrincipalDelegatedPermissionClassificationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticDelegatedPermissionClassifications", "delegatedPermissionClassifications", "delegatedPermissionClassifications"),
		resourceids.UserSpecifiedSegment("delegatedPermissionClassificationId", "delegatedPermissionClassificationIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Delegated Permission Classification ID
func (id ServicePrincipalDelegatedPermissionClassificationId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Delegated Permission Classification: %q", id.DelegatedPermissionClassificationId),
	}
	return fmt.Sprintf("Service Principal Delegated Permission Classification (%s)", strings.Join(components, "\n"))
}
