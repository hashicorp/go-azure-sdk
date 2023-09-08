package serviceprincipaltokenlifetimepolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalTokenLifetimePolicyId{}

// ServicePrincipalTokenLifetimePolicyId is a struct representing the Resource ID for a Service Principal Token Lifetime Policy
type ServicePrincipalTokenLifetimePolicyId struct {
	ServicePrincipalId    string
	TokenLifetimePolicyId string
}

// NewServicePrincipalTokenLifetimePolicyID returns a new ServicePrincipalTokenLifetimePolicyId struct
func NewServicePrincipalTokenLifetimePolicyID(servicePrincipalId string, tokenLifetimePolicyId string) ServicePrincipalTokenLifetimePolicyId {
	return ServicePrincipalTokenLifetimePolicyId{
		ServicePrincipalId:    servicePrincipalId,
		TokenLifetimePolicyId: tokenLifetimePolicyId,
	}
}

// ParseServicePrincipalTokenLifetimePolicyID parses 'input' into a ServicePrincipalTokenLifetimePolicyId
func ParseServicePrincipalTokenLifetimePolicyID(input string) (*ServicePrincipalTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalTokenLifetimePolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.TokenLifetimePolicyId, ok = parsed.Parsed["tokenLifetimePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalTokenLifetimePolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalTokenLifetimePolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalTokenLifetimePolicyIDInsensitively(input string) (*ServicePrincipalTokenLifetimePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalTokenLifetimePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalTokenLifetimePolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.TokenLifetimePolicyId, ok = parsed.Parsed["tokenLifetimePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalTokenLifetimePolicyID checks that 'input' can be parsed as a Service Principal Token Lifetime Policy ID
func ValidateServicePrincipalTokenLifetimePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalTokenLifetimePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Token Lifetime Policy ID
func (id ServicePrincipalTokenLifetimePolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/tokenLifetimePolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.TokenLifetimePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Token Lifetime Policy ID
func (id ServicePrincipalTokenLifetimePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticTokenLifetimePolicies", "tokenLifetimePolicies", "tokenLifetimePolicies"),
		resourceids.UserSpecifiedSegment("tokenLifetimePolicyId", "tokenLifetimePolicyIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Token Lifetime Policy ID
func (id ServicePrincipalTokenLifetimePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Token Lifetime Policy: %q", id.TokenLifetimePolicyId),
	}
	return fmt.Sprintf("Service Principal Token Lifetime Policy (%s)", strings.Join(components, "\n"))
}
