package serviceprincipalcreatedobject

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalCreatedObjectId{}

// ServicePrincipalCreatedObjectId is a struct representing the Resource ID for a Service Principal Created Object
type ServicePrincipalCreatedObjectId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalCreatedObjectID returns a new ServicePrincipalCreatedObjectId struct
func NewServicePrincipalCreatedObjectID(servicePrincipalId string, directoryObjectId string) ServicePrincipalCreatedObjectId {
	return ServicePrincipalCreatedObjectId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalCreatedObjectID parses 'input' into a ServicePrincipalCreatedObjectId
func ParseServicePrincipalCreatedObjectID(input string) (*ServicePrincipalCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalCreatedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalCreatedObjectId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalCreatedObjectIDInsensitively parses 'input' case-insensitively into a ServicePrincipalCreatedObjectId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalCreatedObjectIDInsensitively(input string) (*ServicePrincipalCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalCreatedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalCreatedObjectId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalCreatedObjectID checks that 'input' can be parsed as a Service Principal Created Object ID
func ValidateServicePrincipalCreatedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalCreatedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Created Object ID
func (id ServicePrincipalCreatedObjectId) ID() string {
	fmtString := "/servicePrincipals/%s/createdObjects/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Created Object ID
func (id ServicePrincipalCreatedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticCreatedObjects", "createdObjects", "createdObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Created Object ID
func (id ServicePrincipalCreatedObjectId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Created Object (%s)", strings.Join(components, "\n"))
}
