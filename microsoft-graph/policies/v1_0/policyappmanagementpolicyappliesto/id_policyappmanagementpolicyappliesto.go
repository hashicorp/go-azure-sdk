package policyappmanagementpolicyappliesto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyAppManagementPolicyAppliesToId{}

// PolicyAppManagementPolicyAppliesToId is a struct representing the Resource ID for a Policy App Management Policy Applies To
type PolicyAppManagementPolicyAppliesToId struct {
	AppManagementPolicyId string
	DirectoryObjectId     string
}

// NewPolicyAppManagementPolicyAppliesToID returns a new PolicyAppManagementPolicyAppliesToId struct
func NewPolicyAppManagementPolicyAppliesToID(appManagementPolicyId string, directoryObjectId string) PolicyAppManagementPolicyAppliesToId {
	return PolicyAppManagementPolicyAppliesToId{
		AppManagementPolicyId: appManagementPolicyId,
		DirectoryObjectId:     directoryObjectId,
	}
}

// ParsePolicyAppManagementPolicyAppliesToID parses 'input' into a PolicyAppManagementPolicyAppliesToId
func ParsePolicyAppManagementPolicyAppliesToID(input string) (*PolicyAppManagementPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyAppManagementPolicyAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyAppManagementPolicyAppliesToId{}

	if id.AppManagementPolicyId, ok = parsed.Parsed["appManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParsePolicyAppManagementPolicyAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyAppManagementPolicyAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyAppManagementPolicyAppliesToIDInsensitively(input string) (*PolicyAppManagementPolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyAppManagementPolicyAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyAppManagementPolicyAppliesToId{}

	if id.AppManagementPolicyId, ok = parsed.Parsed["appManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "appManagementPolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyAppManagementPolicyAppliesToID checks that 'input' can be parsed as a Policy App Management Policy Applies To ID
func ValidatePolicyAppManagementPolicyAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAppManagementPolicyAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy App Management Policy Applies To ID
func (id PolicyAppManagementPolicyAppliesToId) ID() string {
	fmtString := "/policies/appManagementPolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.AppManagementPolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy App Management Policy Applies To ID
func (id PolicyAppManagementPolicyAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticAppManagementPolicies", "appManagementPolicies", "appManagementPolicies"),
		resourceids.UserSpecifiedSegment("appManagementPolicyId", "appManagementPolicyIdValue"),
		resourceids.StaticSegment("staticAppliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Policy App Management Policy Applies To ID
func (id PolicyAppManagementPolicyAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("App Management Policy: %q", id.AppManagementPolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy App Management Policy Applies To (%s)", strings.Join(components, "\n"))
}
