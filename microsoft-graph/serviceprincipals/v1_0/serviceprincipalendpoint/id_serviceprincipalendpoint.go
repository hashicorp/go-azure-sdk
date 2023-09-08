package serviceprincipalendpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalEndpointId{}

// ServicePrincipalEndpointId is a struct representing the Resource ID for a Service Principal Endpoint
type ServicePrincipalEndpointId struct {
	ServicePrincipalId string
	EndpointId         string
}

// NewServicePrincipalEndpointID returns a new ServicePrincipalEndpointId struct
func NewServicePrincipalEndpointID(servicePrincipalId string, endpointId string) ServicePrincipalEndpointId {
	return ServicePrincipalEndpointId{
		ServicePrincipalId: servicePrincipalId,
		EndpointId:         endpointId,
	}
}

// ParseServicePrincipalEndpointID parses 'input' into a ServicePrincipalEndpointId
func ParseServicePrincipalEndpointID(input string) (*ServicePrincipalEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalEndpointId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.EndpointId, ok = parsed.Parsed["endpointId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "endpointId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalEndpointIDInsensitively parses 'input' case-insensitively into a ServicePrincipalEndpointId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalEndpointIDInsensitively(input string) (*ServicePrincipalEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalEndpointId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.EndpointId, ok = parsed.Parsed["endpointId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "endpointId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalEndpointID checks that 'input' can be parsed as a Service Principal Endpoint ID
func ValidateServicePrincipalEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Endpoint ID
func (id ServicePrincipalEndpointId) ID() string {
	fmtString := "/servicePrincipals/%s/endpoints/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.EndpointId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Endpoint ID
func (id ServicePrincipalEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticEndpoints", "endpoints", "endpoints"),
		resourceids.UserSpecifiedSegment("endpointId", "endpointIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Endpoint ID
func (id ServicePrincipalEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Endpoint: %q", id.EndpointId),
	}
	return fmt.Sprintf("Service Principal Endpoint (%s)", strings.Join(components, "\n"))
}
