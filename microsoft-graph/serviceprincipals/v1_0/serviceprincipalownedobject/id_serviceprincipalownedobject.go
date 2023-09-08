package serviceprincipalownedobject

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalOwnedObjectId{}

// ServicePrincipalOwnedObjectId is a struct representing the Resource ID for a Service Principal Owned Object
type ServicePrincipalOwnedObjectId struct {
	ServicePrincipalId string
	DirectoryObjectId  string
}

// NewServicePrincipalOwnedObjectID returns a new ServicePrincipalOwnedObjectId struct
func NewServicePrincipalOwnedObjectID(servicePrincipalId string, directoryObjectId string) ServicePrincipalOwnedObjectId {
	return ServicePrincipalOwnedObjectId{
		ServicePrincipalId: servicePrincipalId,
		DirectoryObjectId:  directoryObjectId,
	}
}

// ParseServicePrincipalOwnedObjectID parses 'input' into a ServicePrincipalOwnedObjectId
func ParseServicePrincipalOwnedObjectID(input string) (*ServicePrincipalOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalOwnedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalOwnedObjectId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalOwnedObjectIDInsensitively parses 'input' case-insensitively into a ServicePrincipalOwnedObjectId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalOwnedObjectIDInsensitively(input string) (*ServicePrincipalOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalOwnedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalOwnedObjectId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalOwnedObjectID checks that 'input' can be parsed as a Service Principal Owned Object ID
func ValidateServicePrincipalOwnedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalOwnedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Owned Object ID
func (id ServicePrincipalOwnedObjectId) ID() string {
	fmtString := "/servicePrincipals/%s/ownedObjects/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Owned Object ID
func (id ServicePrincipalOwnedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticOwnedObjects", "ownedObjects", "ownedObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Owned Object ID
func (id ServicePrincipalOwnedObjectId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Service Principal Owned Object (%s)", strings.Join(components, "\n"))
}
