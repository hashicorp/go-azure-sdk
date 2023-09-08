package policymobiledevicemanagementpolicyincludedgroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyMobileDeviceManagementPolicyIncludedGroupId{}

// PolicyMobileDeviceManagementPolicyIncludedGroupId is a struct representing the Resource ID for a Policy Mobile Device Management Policy Included Group
type PolicyMobileDeviceManagementPolicyIncludedGroupId struct {
	MobilityManagementPolicyId string
	GroupId                    string
}

// NewPolicyMobileDeviceManagementPolicyIncludedGroupID returns a new PolicyMobileDeviceManagementPolicyIncludedGroupId struct
func NewPolicyMobileDeviceManagementPolicyIncludedGroupID(mobilityManagementPolicyId string, groupId string) PolicyMobileDeviceManagementPolicyIncludedGroupId {
	return PolicyMobileDeviceManagementPolicyIncludedGroupId{
		MobilityManagementPolicyId: mobilityManagementPolicyId,
		GroupId:                    groupId,
	}
}

// ParsePolicyMobileDeviceManagementPolicyIncludedGroupID parses 'input' into a PolicyMobileDeviceManagementPolicyIncludedGroupId
func ParsePolicyMobileDeviceManagementPolicyIncludedGroupID(input string) (*PolicyMobileDeviceManagementPolicyIncludedGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyMobileDeviceManagementPolicyIncludedGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyMobileDeviceManagementPolicyIncludedGroupId{}

	if id.MobilityManagementPolicyId, ok = parsed.Parsed["mobilityManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobilityManagementPolicyId", *parsed)
	}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	return &id, nil
}

// ParsePolicyMobileDeviceManagementPolicyIncludedGroupIDInsensitively parses 'input' case-insensitively into a PolicyMobileDeviceManagementPolicyIncludedGroupId
// note: this method should only be used for API response data and not user input
func ParsePolicyMobileDeviceManagementPolicyIncludedGroupIDInsensitively(input string) (*PolicyMobileDeviceManagementPolicyIncludedGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyMobileDeviceManagementPolicyIncludedGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyMobileDeviceManagementPolicyIncludedGroupId{}

	if id.MobilityManagementPolicyId, ok = parsed.Parsed["mobilityManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mobilityManagementPolicyId", *parsed)
	}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyMobileDeviceManagementPolicyIncludedGroupID checks that 'input' can be parsed as a Policy Mobile Device Management Policy Included Group ID
func ValidatePolicyMobileDeviceManagementPolicyIncludedGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyMobileDeviceManagementPolicyIncludedGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Mobile Device Management Policy Included Group ID
func (id PolicyMobileDeviceManagementPolicyIncludedGroupId) ID() string {
	fmtString := "/policies/mobileDeviceManagementPolicies/%s/includedGroups/%s"
	return fmt.Sprintf(fmtString, id.MobilityManagementPolicyId, id.GroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Mobile Device Management Policy Included Group ID
func (id PolicyMobileDeviceManagementPolicyIncludedGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticMobileDeviceManagementPolicies", "mobileDeviceManagementPolicies", "mobileDeviceManagementPolicies"),
		resourceids.UserSpecifiedSegment("mobilityManagementPolicyId", "mobilityManagementPolicyIdValue"),
		resourceids.StaticSegment("staticIncludedGroups", "includedGroups", "includedGroups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
	}
}

// String returns a human-readable description of this Policy Mobile Device Management Policy Included Group ID
func (id PolicyMobileDeviceManagementPolicyIncludedGroupId) String() string {
	components := []string{
		fmt.Sprintf("Mobility Management Policy: %q", id.MobilityManagementPolicyId),
		fmt.Sprintf("Group: %q", id.GroupId),
	}
	return fmt.Sprintf("Policy Mobile Device Management Policy Included Group (%s)", strings.Join(components, "\n"))
}
