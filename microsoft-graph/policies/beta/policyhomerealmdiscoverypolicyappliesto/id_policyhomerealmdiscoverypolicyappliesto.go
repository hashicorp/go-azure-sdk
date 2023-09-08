package policyhomerealmdiscoverypolicyappliesto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyHomeRealmDiscoveryPolicyAppliesToId{}

// PolicyHomeRealmDiscoveryPolicyAppliesToId is a struct representing the Resource ID for a Policy Home Realm Discovery Policy Applies To
type PolicyHomeRealmDiscoveryPolicyAppliesToId struct {
	HomeRealmDiscoveryPolicyId string
	DirectoryObjectId          string
}

// NewPolicyHomeRealmDiscoveryPolicyAppliesToID returns a new PolicyHomeRealmDiscoveryPolicyAppliesToId struct
func NewPolicyHomeRealmDiscoveryPolicyAppliesToID(homeRealmDiscoveryPolicyId string, directoryObjectId string) PolicyHomeRealmDiscoveryPolicyAppliesToId {
	return PolicyHomeRealmDiscoveryPolicyAppliesToId{
		HomeRealmDiscoveryPolicyId: homeRealmDiscoveryPolicyId,
		DirectoryObjectId:          directoryObjectId,
	}
}

// ParsePolicyHomeRealmDiscoveryPolicyAppliesToID parses 'input' into a PolicyHomeRealmDiscoveryPolicyAppliesToId
func ParsePolicyHomeRealmDiscoveryPolicyAppliesToID(input string) (*PolicyHomeRealmDiscoveryPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyHomeRealmDiscoveryPolicyAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyHomeRealmDiscoveryPolicyAppliesToId{}

	if id.HomeRealmDiscoveryPolicyId, ok = parsed.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParsePolicyHomeRealmDiscoveryPolicyAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyHomeRealmDiscoveryPolicyAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyHomeRealmDiscoveryPolicyAppliesToIDInsensitively(input string) (*PolicyHomeRealmDiscoveryPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyHomeRealmDiscoveryPolicyAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyHomeRealmDiscoveryPolicyAppliesToId{}

	if id.HomeRealmDiscoveryPolicyId, ok = parsed.Parsed["homeRealmDiscoveryPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "homeRealmDiscoveryPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyHomeRealmDiscoveryPolicyAppliesToID checks that 'input' can be parsed as a Policy Home Realm Discovery Policy Applies To ID
func ValidatePolicyHomeRealmDiscoveryPolicyAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyHomeRealmDiscoveryPolicyAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Home Realm Discovery Policy Applies To ID
func (id PolicyHomeRealmDiscoveryPolicyAppliesToId) ID() string {
	fmtString := "/policies/homeRealmDiscoveryPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.HomeRealmDiscoveryPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Home Realm Discovery Policy Applies To ID
func (id PolicyHomeRealmDiscoveryPolicyAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticHomeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies", "homeRealmDiscoveryPolicies"),
		resourceids.UserSpecifiedSegment("homeRealmDiscoveryPolicyId", "homeRealmDiscoveryPolicyIdValue"),
		resourceids.StaticSegment("staticAppliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Policy Home Realm Discovery Policy Applies To ID
func (id PolicyHomeRealmDiscoveryPolicyAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Home Realm Discovery Policy: %q", id.HomeRealmDiscoveryPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Home Realm Discovery Policy Applies To (%s)", strings.Join(components, "\n"))
}
