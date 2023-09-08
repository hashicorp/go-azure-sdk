package serviceprincipalsynchronizationtemplateschemadirectory

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalSynchronizationTemplateId{}

// ServicePrincipalSynchronizationTemplateId is a struct representing the Resource ID for a Service Principal Synchronization Template
type ServicePrincipalSynchronizationTemplateId struct {
	ServicePrincipalId        string
	SynchronizationTemplateId string
}

// NewServicePrincipalSynchronizationTemplateID returns a new ServicePrincipalSynchronizationTemplateId struct
func NewServicePrincipalSynchronizationTemplateID(servicePrincipalId string, synchronizationTemplateId string) ServicePrincipalSynchronizationTemplateId {
	return ServicePrincipalSynchronizationTemplateId{
		ServicePrincipalId:        servicePrincipalId,
		SynchronizationTemplateId: synchronizationTemplateId,
	}
}

// ParseServicePrincipalSynchronizationTemplateID parses 'input' into a ServicePrincipalSynchronizationTemplateId
func ParseServicePrincipalSynchronizationTemplateID(input string) (*ServicePrincipalSynchronizationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalSynchronizationTemplateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalSynchronizationTemplateId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.SynchronizationTemplateId, ok = parsed.Parsed["synchronizationTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalSynchronizationTemplateIDInsensitively parses 'input' case-insensitively into a ServicePrincipalSynchronizationTemplateId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalSynchronizationTemplateIDInsensitively(input string) (*ServicePrincipalSynchronizationTemplateId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalSynchronizationTemplateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalSynchronizationTemplateId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.SynchronizationTemplateId, ok = parsed.Parsed["synchronizationTemplateId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "synchronizationTemplateId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalSynchronizationTemplateID checks that 'input' can be parsed as a Service Principal Synchronization Template ID
func ValidateServicePrincipalSynchronizationTemplateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalSynchronizationTemplateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Synchronization Template ID
func (id ServicePrincipalSynchronizationTemplateId) ID() string {
	fmtString := "/servicePrincipals/%s/synchronization/templates/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.SynchronizationTemplateId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Synchronization Template ID
func (id ServicePrincipalSynchronizationTemplateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticSynchronization", "synchronization", "synchronization"),
		resourceids.StaticSegment("staticTemplates", "templates", "templates"),
		resourceids.UserSpecifiedSegment("synchronizationTemplateId", "synchronizationTemplateIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Synchronization Template ID
func (id ServicePrincipalSynchronizationTemplateId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Synchronization Template: %q", id.SynchronizationTemplateId),
	}
	return fmt.Sprintf("Service Principal Synchronization Template (%s)", strings.Join(components, "\n"))
}
