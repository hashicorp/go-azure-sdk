package serviceprincipalmemberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalMemberOfId{}

// ServicePrincipalMemberOfId is a struct representing the Resource ID for a Service Principal Member Of
type ServicePrincipalMemberOfId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalMemberOfID returns a new ServicePrincipalMemberOfId struct
func NewServicePrincipalMemberOfID(servicePrincipalId string, directoryObjectId string) ServicePrincipalMemberOfId {
	return ServicePrincipalMemberOfId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalMemberOfID parses 'input' into a ServicePrincipalMemberOfId
func ParseServicePrincipalMemberOfID(input string) (*ServicePrincipalMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalMemberOfId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalMemberOfIDInsensitively parses 'input' case-insensitively into a ServicePrincipalMemberOfId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalMemberOfIDInsensitively(input string) (*ServicePrincipalMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalMemberOfId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalMemberOfID checks that 'input' can be parsed as a Service Principal Member Of ID
func ValidateServicePrincipalMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Member Of ID
func (id ServicePrincipalMemberOfId) ID() string {
	fmtString := "/servicePrincipals/%s/memberOf/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Member Of ID
func (id ServicePrincipalMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticMemberOf", "memberOf", "memberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Member Of ID
func (id ServicePrincipalMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Member Of (%s)", strings.Join(components, "\n"))
}
