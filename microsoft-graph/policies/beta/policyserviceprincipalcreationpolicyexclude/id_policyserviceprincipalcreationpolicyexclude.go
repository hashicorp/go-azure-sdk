package policyserviceprincipalcreationpolicyexclude

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyServicePrincipalCreationPolicyExcludeId{}

// PolicyServicePrincipalCreationPolicyExcludeId is a struct representing the Resource ID for a Policy Service Principal Creation Policy Exclude
type PolicyServicePrincipalCreationPolicyExcludeId struct {
	ServicePrincipalCreationPolicyId       string
	ServicePrincipalCreationConditionSetId string
}

// NewPolicyServicePrincipalCreationPolicyExcludeID returns a new PolicyServicePrincipalCreationPolicyExcludeId struct
func NewPolicyServicePrincipalCreationPolicyExcludeID(servicePrincipalCreationPolicyId string, servicePrincipalCreationConditionSetId string) PolicyServicePrincipalCreationPolicyExcludeId {
	return PolicyServicePrincipalCreationPolicyExcludeId{
		ServicePrincipalCreationPolicyId:       servicePrincipalCreationPolicyId,
		ServicePrincipalCreationConditionSetId: servicePrincipalCreationConditionSetId,
	}
}

// ParsePolicyServicePrincipalCreationPolicyExcludeID parses 'input' into a PolicyServicePrincipalCreationPolicyExcludeId
func ParsePolicyServicePrincipalCreationPolicyExcludeID(input string) (*PolicyServicePrincipalCreationPolicyExcludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyServicePrincipalCreationPolicyExcludeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyServicePrincipalCreationPolicyExcludeId{}

	if id.ServicePrincipalCreationPolicyId, ok = parsed.Parsed["servicePrincipalCreationPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationPolicyId", *parsed)
	}

	if id.ServicePrincipalCreationConditionSetId, ok = parsed.Parsed["servicePrincipalCreationConditionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationConditionSetId", *parsed)
	}

	return &id, nil
}

// ParsePolicyServicePrincipalCreationPolicyExcludeIDInsensitively parses 'input' case-insensitively into a PolicyServicePrincipalCreationPolicyExcludeId
// note: this method should only be used for API response data and not user input
func ParsePolicyServicePrincipalCreationPolicyExcludeIDInsensitively(input string) (*PolicyServicePrincipalCreationPolicyExcludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyServicePrincipalCreationPolicyExcludeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyServicePrincipalCreationPolicyExcludeId{}

	if id.ServicePrincipalCreationPolicyId, ok = parsed.Parsed["servicePrincipalCreationPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationPolicyId", *parsed)
	}

	if id.ServicePrincipalCreationConditionSetId, ok = parsed.Parsed["servicePrincipalCreationConditionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalCreationConditionSetId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyServicePrincipalCreationPolicyExcludeID checks that 'input' can be parsed as a Policy Service Principal Creation Policy Exclude ID
func ValidatePolicyServicePrincipalCreationPolicyExcludeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyServicePrincipalCreationPolicyExcludeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Service Principal Creation Policy Exclude ID
func (id PolicyServicePrincipalCreationPolicyExcludeId) ID() string {
	fmtString := "/policies/servicePrincipalCreationPolicies/%s/excludes/%s"
	return fmt.Sprintf(fmtString, id.ServicePrincipalCreationPolicyId, id.ServicePrincipalCreationConditionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Service Principal Creation Policy Exclude ID
func (id PolicyServicePrincipalCreationPolicyExcludeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticServicePrincipalCreationPolicies", "servicePrincipalCreationPolicies", "servicePrincipalCreationPolicies"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationPolicyId", "servicePrincipalCreationPolicyIdValue"),
		resourceids.StaticSegment("staticExcludes", "excludes", "excludes"),
		resourceids.UserSpecifiedSegment("servicePrincipalCreationConditionSetId", "servicePrincipalCreationConditionSetIdValue"),
	}
}

// String returns a human-readable description of this Policy Service Principal Creation Policy Exclude ID
func (id PolicyServicePrincipalCreationPolicyExcludeId) String() string {
	components := []string{
		fmt.Sprintf("Service Principal Creation Policy: %q", id.ServicePrincipalCreationPolicyId),
		fmt.Sprintf("Service Principal Creation Condition Set: %q", id.ServicePrincipalCreationConditionSetId),
	}
	return fmt.Sprintf("Policy Service Principal Creation Policy Exclude (%s)", strings.Join(components, "\n"))
}
