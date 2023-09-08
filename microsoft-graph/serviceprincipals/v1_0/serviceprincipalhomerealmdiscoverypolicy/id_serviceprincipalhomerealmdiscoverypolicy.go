package serviceprincipalhomerealmdiscoverypolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServicePrincipalHomeRealmDiscoveryPolicyId{}

// ServicePrincipalHomeRealmDiscoveryPolicyId is a struct representing the Resource ID for a Service Principal Home Realm Discovery Policy
type ServicePrincipalHomeRealmDiscoveryPolicyId struct {
	ServicePrincipalId         string
	HomeRealmDiscoveryPolicyId string
}

// NewServicePrincipalHomeRealmDiscoveryPolicyID returns a new ServicePrincipalHomeRealmDiscoveryPolicyId struct
func NewServicePrincipalHomeRealmDiscoveryPolicyID(servicePrincipalId string, homeRealmDiscoveryPolicyId string) ServicePrincipalHomeRealmDiscoveryPolicyId {
	return ServicePrincipalHomeRealmDiscoveryPolicyId{
		ServicePrincipalId:         servicePrincipalId,
		HomeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
	}
}

// ParseServicePrincipalHomeRealmDiscoveryPolicyID parses 'input' into a ServicePrincipalHomeRealmDiscoveryPolicyId
func ParseServicePrincipalHomeRealmDiscoveryPolicyID(input string) (*ServicePrincipalHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalHomeRealmDiscoveryPolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.HomeRealmDiscoveryPolicyId, ok = parsed.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", *parsed)
	}

	return &id, nil
}

// ParseServicePrincipalHomeRealmDiscoveryPolicyIDInsensitively parses 'input' case-insensitively into a ServicePrincipalHomeRealmDiscoveryPolicyId
// note: this method should only be used for API response data and not user input
func ParseServicePrincipalHomeRealmDiscoveryPolicyIDInsensitively(input string) (*ServicePrincipalHomeRealmDiscoveryPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServicePrincipalHomeRealmDiscoveryPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServicePrincipalHomeRealmDiscoveryPolicyId{}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	if id.HomeRealmDiscoveryPolicyId, ok = parsed.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", *parsed)
	}

	return &id, nil
}

// ValidateServicePrincipalHomeRealmDiscoveryPolicyID checks that 'input' can be parsed as a Service Principal Home Realm Discovery Policy ID
func ValidateServicePrincipalHomeRealmDiscoveryPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServicePrincipalHomeRealmDiscoveryPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Service Principal Home Realm Discovery Policy ID
func (id ServicePrincipalHomeRealmDiscoveryPolicyId) ID() string {
	fmtString := "/servicePrincipals/%s/homeRealmDiscoveryPolicies/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalId, id.HomeRealmDiscoveryPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Service Principal Home Realm Discovery Policy ID
func (id ServicePrincipalHomeRealmDiscoveryPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticServicePrincipals", "servicePrincipals", "servicePrincipals"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
		resourceids.StaticSegment("staticHomeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies"),
		resourceids.UserSpecifiedSegment("homeRealmDiscoveryPolicyId", "homeRealmDiscoveryPolicyIdValue"),
	}
}

// String returns a human-readable description of this Service Principal Home Realm Discovery Policy ID
func (id ServicePrincipalHomeRealmDiscoveryPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
		fmt.Sprintf("Home Realm Discovery Policy: %q", id.HomeRealmDiscoveryPolicyId),
	}
	return fmt.Sprintf("Service Principal Home Realm Discovery Policy (%s)", strings.Join(components, "\n"))
}
