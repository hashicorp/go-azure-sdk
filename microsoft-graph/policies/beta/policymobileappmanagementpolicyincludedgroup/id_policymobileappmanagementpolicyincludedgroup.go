package policymobileappmanagementpolicyincludedgroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyMobileAppManagementPolicyIncludedGroupId{}

// PolicyMobileAppManagementPolicyIncludedGroupId is a struct representing the Resource ID for a Policy Mobile App Management Policy Included Group
type PolicyMobileAppManagementPolicyIncludedGroupId struct {
	MobilityManagementPolicyId string
	GroupId                    string
}

// NewPolicyMobileAppManagementPolicyIncludedGroupID returns a new PolicyMobileAppManagementPolicyIncludedGroupId struct
func NewPolicyMobileAppManagementPolicyIncludedGroupID(mobilityManagementPolicyId string, groupId string) PolicyMobileAppManagementPolicyIncludedGroupId {
	return PolicyMobileAppManagementPolicyIncludedGroupId{
		MobilityManagementPolicyId: mobilityManagementPolicyId,
		GroupId:                    groupId,
	}
}

// ParsePolicyMobileAppManagementPolicyIncludedGroupID parses 'input' into a PolicyMobileAppManagementPolicyIncludedGroupId
func ParsePolicyMobileAppManagementPolicyIncludedGroupID(input string) (*PolicyMobileAppManagementPolicyIncludedGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyMobileAppManagementPolicyIncludedGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyMobileAppManagementPolicyIncludedGroupId{}

	if id.MobilityManagementPolicyId, ok = parsed.Parsed["mobilityManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobilityManagementPolicyId", *parsed)
	}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	return &id, nil
}

// ParsePolicyMobileAppManagementPolicyIncludedGroupIDInsensitively parses 'input' case-insensitively into a PolicyMobileAppManagementPolicyIncludedGroupId
// note: this method should only be used for API response data and not user input
func ParsePolicyMobileAppManagementPolicyIncludedGroupIDInsensitively(input string) (*PolicyMobileAppManagementPolicyIncludedGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyMobileAppManagementPolicyIncludedGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyMobileAppManagementPolicyIncludedGroupId{}

	if id.MobilityManagementPolicyId, ok = parsed.Parsed["mobilityManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobilityManagementPolicyId", *parsed)
	}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyMobileAppManagementPolicyIncludedGroupID checks that 'input' can be parsed as a Policy Mobile App Management Policy Included Group ID
func ValidatePolicyMobileAppManagementPolicyIncludedGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyMobileAppManagementPolicyIncludedGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Mobile App Management Policy Included Group ID
func (id PolicyMobileAppManagementPolicyIncludedGroupId) ID() string {
	fmtString := "/policies/mobileAppManagementPolicies/%s/includedGroups/%s"
	return fmt.Sprintf(fmtString, id.MobilityManagementPolicyId, id.GroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Mobile App Management Policy Included Group ID
func (id PolicyMobileAppManagementPolicyIncludedGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticMobileAppManagementPolicies", "mobileAppManagementPolicies", "mobileAppManagementPolicies"),
		resourceids.UserSpecifiedSegment("mobilityManagementPolicyId", "mobilityManagementPolicyIdValue"),
		resourceids.StaticSegment("staticIncludedGroups", "includedGroups", "includedGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
	}
}

// String returns a human-readable description of this Policy Mobile App Management Policy Included Group ID
func (id PolicyMobileAppManagementPolicyIncludedGroupId) String() string {
	components := []string{
		fmt.Sprintf("Mobility Management Policy: %q", id.MobilityManagementPolicyId),
		fmt.Sprintf("Group: %q", id.GroupId),
	}
	return fmt.Sprintf("Policy Mobile App Management Policy Included Group (%s)", strings.Join(components, "\n"))
}
