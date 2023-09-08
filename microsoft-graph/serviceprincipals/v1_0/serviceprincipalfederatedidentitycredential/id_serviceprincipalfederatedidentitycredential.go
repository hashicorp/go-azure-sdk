package serviceprincipalfederatedidentitycredential

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalFederatedIdentityCredentialId{}

// ServicePrincipalFederatedIdentityCredentialId is a struct representing the Resource ID for a Service Principal Federated Identity Credential
type ServicePrincipalFederatedIdentityCredentialId struct {
	ServicePrincipalId            string
	FederatedIdentityCredentialId string
}

// NewServicePrincipalFederatedIdentityCredentialID returns a new ServicePrincipalFederatedIdentityCredentialId struct
func NewServicePrincipalFederatedIdentityCredentialID(servicePrincipalId string, federatedIdentityCredentialId string) ServicePrincipalFederatedIdentityCredentialId {
	return ServicePrincipalFederatedIdentityCredentialId{
		ServicePrincipalId:            servicePrincipalId,
		FederatedIdentityCredentialId: federatedIdentityCredentialId,
	}
}

// ParseServicePrincipalFederatedIdentityCredentialID parses 'input' into a ServicePrincipalFederatedIdentityCredentialId
func ParseServicePrincipalFederatedIdentityCredentialID(input string) (*ServicePrincipalFederatedIdentityCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalFederatedIdentityCredentialId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalFederatedIdentityCredentialId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.FederatedIdentityCredentialId, ok = parsed.Parsed["federatedIdentityCredentialId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "federatedIdentityCredentialId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalFederatedIdentityCredentialIDInsensitively parses 'input' case-insensitively into a ServicePrincipalFederatedIdentityCredentialId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalFederatedIdentityCredentialIDInsensitively(input string) (*ServicePrincipalFederatedIdentityCredentialId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalFederatedIdentityCredentialId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalFederatedIdentityCredentialId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.FederatedIdentityCredentialId, ok = parsed.Parsed["federatedIdentityCredentialId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "federatedIdentityCredentialId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalFederatedIdentityCredentialID checks that 'input' can be parsed as a Service Principal Federated Identity Credential ID
func ValidateServicePrincipalFederatedIdentityCredentialID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalFederatedIdentityCredentialID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Federated Identity Credential ID
func (id ServicePrincipalFederatedIdentityCredentialId) ID() string {
	fmtString := "/servicePrincipals/%s/federatedIdentityCredentials/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.FederatedIdentityCredentialId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Federated Identity Credential ID
func (id ServicePrincipalFederatedIdentityCredentialId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticFederatedIdentityCredentials", "federatedIdentityCredentials", "federatedIdentityCredentials"),
		resourceids.UserSpecifiedSegment("federatedIdentityCredentialId", "federatedIdentityCredentialIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Federated Identity Credential ID
func (id ServicePrincipalFederatedIdentityCredentialId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Federated Identity Credential: %q", id.FederatedIdentityCredentialId),
	}
	return fmt.Sprintf("Service Principal Federated Identity Credential (%s)", strings.Join(components, "\n"))
}
