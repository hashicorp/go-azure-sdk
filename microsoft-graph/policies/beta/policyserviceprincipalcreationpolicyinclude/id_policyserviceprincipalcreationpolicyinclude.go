package policyserviceprincipalcreationpolicyinclude

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyServicePrincipalCreationPolicyIncludeId{}

// PolicyServicePrincipalCreationPolicyIncludeId is a struct representing the Resource ID for a Policy Service Principal Creation Policy Include
type PolicyServicePrincipalCreationPolicyIncludeId struct {
	ServicePrincipalCreationPolicyId       string
	ServicePrincipalCreationConditionSetId string
}

// NewPolicyServicePrincipalCreationPolicyIncludeID returns a new PolicyServicePrincipalCreationPolicyIncludeId struct
func NewPolicyServicePrincipalCreationPolicyIncludeID(servicePrincipalCreationPolicyId string, servicePrincipalCreationConditionSetId string) PolicyServicePrincipalCreationPolicyIncludeId {
	return PolicyServicePrincipalCreationPolicyIncludeId{
		ServicePrincipalCreationPolicyId:       servicePrincipalCreationPolicyId,
		ServicePrincipalCreationConditionSetId: servicePrincipalCreationConditionSetId,
	}
}

// ParsePolicyServicePrincipalCreationPolicyIncludeID parses 'input' into a PolicyServicePrincipalCreationPolicyIncludeId
func ParsePolicyServicePrincipalCreationPolicyIncludeID(input string) (*PolicyServicePrincipalCreationPolicyIncludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyServicePrincipalCreationPolicyIncludeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyServicePrincipalCreationPolicyIncludeId{}

	if id.ServicePrincipalCreationPolicyId, ok = parsed.Parsed["servicePrincipalCreationPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationPolicyId", *parsed)
	}

	if id.ServicePrincipalCreationConditionSetId, ok = parsed.Parsed["servicePrincipalCreationConditionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationConditionSetId", *parsed)
	}

	return &id, nil
}

// ParsePolicyServicePrincipalCreationPolicyIncludeIDInsensitively parses 'input' case-insensitively into a PolicyServicePrincipalCreationPolicyIncludeId
// note: this method should only be used for API response data and not user input
func ParsePolicyServicePrincipalCreationPolicyIncludeIDInsensitively(input string) (*PolicyServicePrincipalCreationPolicyIncludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyServicePrincipalCreationPolicyIncludeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyServicePrincipalCreationPolicyIncludeId{}

	if id.ServicePrincipalCreationPolicyId, ok = parsed.Parsed["servicePrincipalCreationPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationPolicyId", *parsed)
	}

	if id.ServicePrincipalCreationConditionSetId, ok = parsed.Parsed["servicePrincipalCreationConditionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationConditionSetId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyServicePrincipalCreationPolicyIncludeID checks that 'input' can be parsed as a Policy Service Principal Creation Policy Include ID
func ValidatePolicyServicePrincipalCreationPolicyIncludeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyServicePrincipalCreationPolicyIncludeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Service Principal Creation Policy Include ID
func (id PolicyServicePrincipalCreationPolicyIncludeId) ID() string {
	fmtString := "/policies/servicePrincipalCreationPolicies/%s/includes/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalCreationPolicyId, id.ServicePrincipalCreationConditionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Service Principal Creation Policy Include ID
func (id PolicyServicePrincipalCreationPolicyIncludeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticServicePrincipalCreationPolicies", "servicePrincipalCreationPolicies", "servicePrincipalCreationPolicies"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationPolicyId", "servicePrincipalCreationPolicyIdValue"),
		resourceids.StaticSegment("staticIncludes", "includes", "includes"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationConditionSetId", "servicePrincipalCreationConditionSetIdValue"),
	}
}

// String returns a human-readable description of this Policy Service Principal Creation Policy Include ID
func (id PolicyServicePrincipalCreationPolicyIncludeId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal Creation Policy: %q", id.ServicePrincipalCreationPolicyId),
		fmt.Sprintf("Service Principal Creation Condition Set: %q", id.ServicePrincipalCreationConditionSetId),
	}
	return fmt.Sprintf("Policy Service Principal Creation Policy Include (%s)", strings.Join(components, "\n"))
}
