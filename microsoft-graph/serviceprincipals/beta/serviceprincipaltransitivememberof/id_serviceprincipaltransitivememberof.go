package serviceprincipaltransitivememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalTransitiveMemberOfId{}

// ServicePrincipalTransitiveMemberOfId is a struct representing the Resource ID for a Service Principal Transitive Member Of
type ServicePrincipalTransitiveMemberOfId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalTransitiveMemberOfID returns a new ServicePrincipalTransitiveMemberOfId struct
func NewServicePrincipalTransitiveMemberOfID(servicePrincipalId string, directoryObjectId string) ServicePrincipalTransitiveMemberOfId {
	return ServicePrincipalTransitiveMemberOfId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalTransitiveMemberOfID parses 'input' into a ServicePrincipalTransitiveMemberOfId
func ParseServicePrincipalTransitiveMemberOfID(input string) (*ServicePrincipalTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalTransitiveMemberOfId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a ServicePrincipalTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalTransitiveMemberOfIDInsensitively(input string) (*ServicePrincipalTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalTransitiveMemberOfId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalTransitiveMemberOfID checks that 'input' can be parsed as a Service Principal Transitive Member Of ID
func ValidateServicePrincipalTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Transitive Member Of ID
func (id ServicePrincipalTransitiveMemberOfId) ID() string {
	fmtString := "/servicePrincipals/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Transitive Member Of ID
func (id ServicePrincipalTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticTransitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Transitive Member Of ID
func (id ServicePrincipalTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Transitive Member Of (%s)", strings.Join(components, "\n"))
}
