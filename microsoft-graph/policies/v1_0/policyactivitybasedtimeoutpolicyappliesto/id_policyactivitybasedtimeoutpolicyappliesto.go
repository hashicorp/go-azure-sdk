package policyactivitybasedtimeoutpolicyappliesto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyActivityBasedTimeoutPolicyAppliesToId{}

// PolicyActivityBasedTimeoutPolicyAppliesToId is a struct representing the Resource ID for a Policy Activity Based Timeout Policy Applies To
type PolicyActivityBasedTimeoutPolicyAppliesToId struct {
	ActivityBasedTimeoutPolicyId string
	DirectoryObjectId            string
}

// NewPolicyActivityBasedTimeoutPolicyAppliesToID returns a new PolicyActivityBasedTimeoutPolicyAppliesToId struct
func NewPolicyActivityBasedTimeoutPolicyAppliesToID(activityBasedTimeoutPolicyId string, directoryObjectId string) PolicyActivityBasedTimeoutPolicyAppliesToId {
	return PolicyActivityBasedTimeoutPolicyAppliesToId{
		ActivityBasedTimeoutPolicyId: activityBasedTimeoutPolicyId,
		DirectoryObjectId:            directoryObjectId,
	}
}

// ParsePolicyActivityBasedTimeoutPolicyAppliesToID parses 'input' into a PolicyActivityBasedTimeoutPolicyAppliesToId
func ParsePolicyActivityBasedTimeoutPolicyAppliesToID(input string) (*PolicyActivityBasedTimeoutPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyActivityBasedTimeoutPolicyAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyActivityBasedTimeoutPolicyAppliesToId{}

	if id.ActivityBasedTimeoutPolicyId, ok = parsed.Parsed["activityBasedTimeoutPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "activityBasedTimeoutPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParsePolicyActivityBasedTimeoutPolicyAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyActivityBasedTimeoutPolicyAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyActivityBasedTimeoutPolicyAppliesToIDInsensitively(input string) (*PolicyActivityBasedTimeoutPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyActivityBasedTimeoutPolicyAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyActivityBasedTimeoutPolicyAppliesToId{}

	if id.ActivityBasedTimeoutPolicyId, ok = parsed.Parsed["activityBasedTimeoutPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "activityBasedTimeoutPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyActivityBasedTimeoutPolicyAppliesToID checks that 'input' can be parsed as a Policy Activity Based Timeout Policy Applies To ID
func ValidatePolicyActivityBasedTimeoutPolicyAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyActivityBasedTimeoutPolicyAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Activity Based Timeout Policy Applies To ID
func (id PolicyActivityBasedTimeoutPolicyAppliesToId) ID() string {
	fmtString := "/policies/activityBasedTimeoutPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.ActivityBasedTimeoutPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Activity Based Timeout Policy Applies To ID
func (id PolicyActivityBasedTimeoutPolicyAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticActivityBasedTimeoutPolicies", "activityBasedTimeoutPolicies", "activityBasedTimeoutPolicies"),
		resourceids.UserSpecifiedSegment("activityBasedTimeoutPolicyId", "activityBasedTimeoutPolicyIdValue"),
		resourceids.StaticSegment("staticAppliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Policy Activity Based Timeout Policy Applies To ID
func (id PolicyActivityBasedTimeoutPolicyAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Activity Based Timeout Policy: %q", id.ActivityBasedTimeoutPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Activity Based Timeout Policy Applies To (%s)", strings.Join(components, "\n"))
}
