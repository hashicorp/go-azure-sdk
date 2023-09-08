package serviceprincipalowner

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalOwnerId{}

// ServicePrincipalOwnerId is a struct representing the Resource ID for a Service Principal Owner
type ServicePrincipalOwnerId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalOwnerID returns a new ServicePrincipalOwnerId struct
func NewServicePrincipalOwnerID(servicePrincipalId string, directoryObjectId string) ServicePrincipalOwnerId {
	return ServicePrincipalOwnerId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalOwnerID parses 'input' into a ServicePrincipalOwnerId
func ParseServicePrincipalOwnerID(input string) (*ServicePrincipalOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalOwnerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalOwnerId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalOwnerIDInsensitively parses 'input' case-insensitively into a ServicePrincipalOwnerId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalOwnerIDInsensitively(input string) (*ServicePrincipalOwnerId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalOwnerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalOwnerId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalOwnerID checks that 'input' can be parsed as a Service Principal Owner ID
func ValidateServicePrincipalOwnerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalOwnerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Owner ID
func (id ServicePrincipalOwnerId) ID() string {
	fmtString := "/servicePrincipals/%s/owners/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Owner ID
func (id ServicePrincipalOwnerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticOwners", "owners", "owners"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Owner ID
func (id ServicePrincipalOwnerId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Owner (%s)", strings.Join(components, "\n"))
}
