package serviceprincipalclaimsmappingpolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalClaimsMappingPolicyId{}

// ServicePrincipalClaimsMappingPolicyId is a struct representing the Resource ID for a Service Principal Claims Mapping Policy
type ServicePrincipalClaimsMappingPolicyId struct {
	ServicePrincipalId    string
	ClaimsMappingPolicyId string
}

// NewServicePrincipalClaimsMappingPolicyID returns a new ServicePrincipalClaimsMappingPolicyId struct
func NewServicePrincipalClaimsMappingPolicyID(servicePrincipalId string, claimsMappingPolicyId string) ServicePrincipalClaimsMappingPolicyId {
	return ServicePrincipalClaimsMappingPolicyId{
		ServicePrincipalId:    servicePrincipalId,
		ClaimsMappingPolicyId: claimsMappingPolicyId,
	}
}

// ParseServicePrincipalClaimsMappingPolicyID parses 'input' into a ServicePrincipalClaimsMappingPolicyId
func ParseServicePrincipalClaimsMappingPolicyID(input string) (*ServicePrincipalClaimsMappingPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalClaimsMappingPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalClaimsMappingPolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.ClaimsMappingPolicyId, ok = parsed.Parsed["claimsMappingPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "claimsMappingPolicyId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalClaimsMappingPolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalClaimsMappingPolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalClaimsMappingPolicyIDInsensitively(input string) (*ServicePrincipalClaimsMappingPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalClaimsMappingPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalClaimsMappingPolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.ClaimsMappingPolicyId, ok = parsed.Parsed["claimsMappingPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "claimsMappingPolicyId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalClaimsMappingPolicyID checks that 'input' can be parsed as a Service Principal Claims Mapping Policy ID
func ValidateServicePrincipalClaimsMappingPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalClaimsMappingPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Claims Mapping Policy ID
func (id ServicePrincipalClaimsMappingPolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/claimsMappingPolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.ClaimsMappingPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Claims Mapping Policy ID
func (id ServicePrincipalClaimsMappingPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticClaimsMappingPolicies", "claimsMappingPolicies", "claimsMappingPolicies"),
		resourceids.UserSpecifiedSegment("claimsMappingPolicyId", "claimsMappingPolicyIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Claims Mapping Policy ID
func (id ServicePrincipalClaimsMappingPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Claims Mapping Policy: %q", id.ClaimsMappingPolicyId),
	}
	return fmt.Sprintf("Service Principal Claims Mapping Policy (%s)", strings.Join(components, "\n"))
}
