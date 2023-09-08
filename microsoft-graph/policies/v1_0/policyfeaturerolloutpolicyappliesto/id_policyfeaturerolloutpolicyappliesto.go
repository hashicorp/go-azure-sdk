package policyfeaturerolloutpolicyappliesto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyFeatureRolloutPolicyAppliesToId{}

// PolicyFeatureRolloutPolicyAppliesToId is a struct representing the Resource ID for a Policy Feature Rollout Policy Applies To
type PolicyFeatureRolloutPolicyAppliesToId struct {
	FeatureRolloutPolicyId string
	DirectoryObjectId      string
}

// NewPolicyFeatureRolloutPolicyAppliesToID returns a new PolicyFeatureRolloutPolicyAppliesToId struct
func NewPolicyFeatureRolloutPolicyAppliesToID(featureRolloutPolicyId string, directoryObjectId string) PolicyFeatureRolloutPolicyAppliesToId {
	return PolicyFeatureRolloutPolicyAppliesToId{
		FeatureRolloutPolicyId: featureRolloutPolicyId,
		DirectoryObjectId:      directoryObjectId,
	}
}

// ParsePolicyFeatureRolloutPolicyAppliesToID parses 'input' into a PolicyFeatureRolloutPolicyAppliesToId
func ParsePolicyFeatureRolloutPolicyAppliesToID(input string) (*PolicyFeatureRolloutPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyFeatureRolloutPolicyAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyFeatureRolloutPolicyAppliesToId{}

	if id.FeatureRolloutPolicyId, ok = parsed.Parsed["featureRolloutPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureRolloutPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParsePolicyFeatureRolloutPolicyAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyFeatureRolloutPolicyAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyFeatureRolloutPolicyAppliesToIDInsensitively(input string) (*PolicyFeatureRolloutPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyFeatureRolloutPolicyAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyFeatureRolloutPolicyAppliesToId{}

	if id.FeatureRolloutPolicyId, ok = parsed.Parsed["featureRolloutPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureRolloutPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyFeatureRolloutPolicyAppliesToID checks that 'input' can be parsed as a Policy Feature Rollout Policy Applies To ID
func ValidatePolicyFeatureRolloutPolicyAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyFeatureRolloutPolicyAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Feature Rollout Policy Applies To ID
func (id PolicyFeatureRolloutPolicyAppliesToId) ID() string {
	fmtString := "/policies/featureRolloutPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.FeatureRolloutPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Feature Rollout Policy Applies To ID
func (id PolicyFeatureRolloutPolicyAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticFeatureRolloutPolicies", "featureRolloutPolicies", "featureRolloutPolicies"),
		resourceids.UserSpecifiedSegment("featureRolloutPolicyId", "featureRolloutPolicyIdValue"),
		resourceids.StaticSegment("staticAppliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Policy Feature Rollout Policy Applies To ID
func (id PolicyFeatureRolloutPolicyAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Feature Rollout Policy: %q", id.FeatureRolloutPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Feature Rollout Policy Applies To (%s)", strings.Join(components, "\n"))
}
