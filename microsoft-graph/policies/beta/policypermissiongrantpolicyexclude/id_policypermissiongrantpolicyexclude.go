package policypermissiongrantpolicyexclude

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyPermissionGrantPolicyExcludeId{}

// PolicyPermissionGrantPolicyExcludeId is a struct representing the Resource ID for a Policy Permission Grant Policy Exclude
type PolicyPermissionGrantPolicyExcludeId struct {
	PermissionGrantPolicyId       string
	PermissionGrantConditionSetId string
}

// NewPolicyPermissionGrantPolicyExcludeID returns a new PolicyPermissionGrantPolicyExcludeId struct
func NewPolicyPermissionGrantPolicyExcludeID(permissionGrantPolicyId string, permissionGrantConditionSetId string) PolicyPermissionGrantPolicyExcludeId {
	return PolicyPermissionGrantPolicyExcludeId{
		PermissionGrantPolicyId:       permissionGrantPolicyId,
		PermissionGrantConditionSetId: permissionGrantConditionSetId,
	}
}

// ParsePolicyPermissionGrantPolicyExcludeID parses 'input' into a PolicyPermissionGrantPolicyExcludeId
func ParsePolicyPermissionGrantPolicyExcludeID(input string) (*PolicyPermissionGrantPolicyExcludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyPermissionGrantPolicyExcludeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyPermissionGrantPolicyExcludeId{}

	if id.PermissionGrantPolicyId, ok = parsed.Parsed["permissionGrantPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPolicyId", *parsed)
	}

	if id.PermissionGrantConditionSetId, ok = parsed.Parsed["permissionGrantConditionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantConditionSetId", *parsed)
	}

	return &id, nil
}

// ParsePolicyPermissionGrantPolicyExcludeIDInsensitively parses 'input' case-insensitively into a PolicyPermissionGrantPolicyExcludeId
// note: this method should only be used for API response data and not user input
func ParsePolicyPermissionGrantPolicyExcludeIDInsensitively(input string) (*PolicyPermissionGrantPolicyExcludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyPermissionGrantPolicyExcludeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyPermissionGrantPolicyExcludeId{}

	if id.PermissionGrantPolicyId, ok = parsed.Parsed["permissionGrantPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPolicyId", *parsed)
	}

	if id.PermissionGrantConditionSetId, ok = parsed.Parsed["permissionGrantConditionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantConditionSetId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyPermissionGrantPolicyExcludeID checks that 'input' can be parsed as a Policy Permission Grant Policy Exclude ID
func ValidatePolicyPermissionGrantPolicyExcludeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyPermissionGrantPolicyExcludeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Permission Grant Policy Exclude ID
func (id PolicyPermissionGrantPolicyExcludeId) ID() string {
	fmtString := "/policies/permissionGrantPolicies/%s/excludes/%s"
	return fmt.Sprintf(fmtString, id.PermissionGrantPolicyId, id.PermissionGrantConditionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Permission Grant Policy Exclude ID
func (id PolicyPermissionGrantPolicyExcludeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticPermissionGrantPolicies", "permissionGrantPolicies", "permissionGrantPolicies"),
		resourceids.UserSpecifiedSegment("permissionGrantPolicyId", "permissionGrantPolicyIdValue"),
		resourceids.StaticSegment("staticExcludes", "excludes", "excludes"),
		resourceids.UserSpecifiedSegment("permissionGrantConditionSetId", "permissionGrantConditionSetIdValue"),
	}
}

// String returns a human-readable description of this Policy Permission Grant Policy Exclude ID
func (id PolicyPermissionGrantPolicyExcludeId) String() string {
	components := []string{
		fmt.Sprintf("Permission Grant Policy: %q", id.PermissionGrantPolicyId),
		fmt.Sprintf("Permission Grant Condition Set: %q", id.PermissionGrantConditionSetId),
	}
	return fmt.Sprintf("Policy Permission Grant Policy Exclude (%s)", strings.Join(components, "\n"))
}
