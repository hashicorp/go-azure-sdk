package serviceprincipalsynchronizationtemplate

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalId{}

// ServicePrincipalId is a struct representing the Resource ID for a Service Principal
type ServicePrincipalId struct {
	ServicePrincipalId string
}

// NewServicePrincipalID returns a new ServicePrincipalId struct
func NewServicePrincipalID(servicePrincipalId string) ServicePrincipalId {
	return ServicePrincipalId{
		ServicePrincipalId: servicePrincipalId,
	}
}

// ParseServicePrincipalID parses 'input' into a ServicePrincipalId
func ParseServicePrincipalID(input string) (*ServicePrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalIDInsensitively parses 'input' case-insensitively into a ServicePrincipalId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalIDInsensitively(input string) (*ServicePrincipalId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalID checks that 'input' can be parsed as a Service Principal ID
func ValidateServicePrincipalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal ID
func (id ServicePrincipalId) ID() string {
	fmtString := "/servicePrincipals/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal ID
func (id ServicePrincipalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
	}
}

// String returns a human-readable description of this Service Principal ID
func (id ServicePrincipalId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
	}
	return fmt.Sprintf("Service Principal (%s)", strings.Join(components, "\n"))
}
