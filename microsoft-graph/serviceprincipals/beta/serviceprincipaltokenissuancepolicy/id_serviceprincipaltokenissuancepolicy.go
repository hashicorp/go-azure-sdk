package serviceprincipaltokenissuancepolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalTokenIssuancePolicyId{}

// ServicePrincipalTokenIssuancePolicyId is a struct representing the Resource ID for a Service Principal Token Issuance Policy
type ServicePrincipalTokenIssuancePolicyId struct {
	ServicePrincipalId    string
	TokenIssuancePolicyId string
}

// NewServicePrincipalTokenIssuancePolicyID returns a new ServicePrincipalTokenIssuancePolicyId struct
func NewServicePrincipalTokenIssuancePolicyID(servicePrincipalId string, tokenIssuancePolicyId string) ServicePrincipalTokenIssuancePolicyId {
	return ServicePrincipalTokenIssuancePolicyId{
		ServicePrincipalId:    servicePrincipalId,
		TokenIssuancePolicyId: tokenIssuancePolicyId,
	}
}

// ParseServicePrincipalTokenIssuancePolicyID parses 'input' into a ServicePrincipalTokenIssuancePolicyId
func ParseServicePrincipalTokenIssuancePolicyID(input string) (*ServicePrincipalTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalTokenIssuancePolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.TokenIssuancePolicyId, ok = parsed.Parsed["tokenIssuancePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalTokenIssuancePolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalTokenIssuancePolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalTokenIssuancePolicyIDInsensitively(input string) (*ServicePrincipalTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalTokenIssuancePolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.TokenIssuancePolicyId, ok = parsed.Parsed["tokenIssuancePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalTokenIssuancePolicyID checks that 'input' can be parsed as a Service Principal Token Issuance Policy ID
func ValidateServicePrincipalTokenIssuancePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalTokenIssuancePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Token Issuance Policy ID
func (id ServicePrincipalTokenIssuancePolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/tokenIssuancePolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.TokenIssuancePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Token Issuance Policy ID
func (id ServicePrincipalTokenIssuancePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticTokenIssuancePolicies", "tokenIssuancePolicies", "tokenIssuancePolicies"),
		resourceids.UserSpecifiedSegment("tokenIssuancePolicyId", "tokenIssuancePolicyIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Token Issuance Policy ID
func (id ServicePrincipalTokenIssuancePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Token Issuance Policy: %q", id.TokenIssuancePolicyId),
	}
	return fmt.Sprintf("Service Principal Token Issuance Policy (%s)", strings.Join(components, "\n"))
}
