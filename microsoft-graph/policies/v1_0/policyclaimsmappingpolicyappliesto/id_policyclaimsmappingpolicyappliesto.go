package policyclaimsmappingpolicyappliesto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyClaimsMappingPolicyAppliesToId{}

// PolicyClaimsMappingPolicyAppliesToId is a struct representing the Resource ID for a Policy Claims Mapping Policy Applies To
type PolicyClaimsMappingPolicyAppliesToId struct {
	ClaimsMappingPolicyId string
	DirectoryObjectId     string
}

// NewPolicyClaimsMappingPolicyAppliesToID returns a new PolicyClaimsMappingPolicyAppliesToId struct
func NewPolicyClaimsMappingPolicyAppliesToID(claimsMappingPolicyId string, directoryObjectId string) PolicyClaimsMappingPolicyAppliesToId {
	return PolicyClaimsMappingPolicyAppliesToId{
		ClaimsMappingPolicyId: claimsMappingPolicyId,
		DirectoryObjectId:     directoryObjectId,
	}
}

// ParsePolicyClaimsMappingPolicyAppliesToID parses 'input' into a PolicyClaimsMappingPolicyAppliesToId
func ParsePolicyClaimsMappingPolicyAppliesToID(input string) (*PolicyClaimsMappingPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyClaimsMappingPolicyAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyClaimsMappingPolicyAppliesToId{}

	if id.ClaimsMappingPolicyId, ok = parsed.Parsed["claimsMappingPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "claimsMappingPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParsePolicyClaimsMappingPolicyAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyClaimsMappingPolicyAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyClaimsMappingPolicyAppliesToIDInsensitively(input string) (*PolicyClaimsMappingPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyClaimsMappingPolicyAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyClaimsMappingPolicyAppliesToId{}

	if id.ClaimsMappingPolicyId, ok = parsed.Parsed["claimsMappingPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "claimsMappingPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyClaimsMappingPolicyAppliesToID checks that 'input' can be parsed as a Policy Claims Mapping Policy Applies To ID
func ValidatePolicyClaimsMappingPolicyAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyClaimsMappingPolicyAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Claims Mapping Policy Applies To ID
func (id PolicyClaimsMappingPolicyAppliesToId) ID() string {
	fmtString := "/policies/claimsMappingPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.ClaimsMappingPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Claims Mapping Policy Applies To ID
func (id PolicyClaimsMappingPolicyAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticClaimsMappingPolicies", "claimsMappingPolicies", "claimsMappingPolicies"),
		resourceids.UserSpecifiedSegment("claimsMappingPolicyId", "claimsMappingPolicyIdValue"),
		resourceids.StaticSegment("staticAppliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Policy Claims Mapping Policy Applies To ID
func (id PolicyClaimsMappingPolicyAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Claims Mapping Policy: %q", id.ClaimsMappingPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Claims Mapping Policy Applies To (%s)", strings.Join(components, "\n"))
}
