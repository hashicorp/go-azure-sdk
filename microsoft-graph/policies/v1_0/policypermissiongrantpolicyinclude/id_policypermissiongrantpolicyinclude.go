package policypermissiongrantpolicyinclude

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyPermissionGrantPolicyIncludeId{}

// PolicyPermissionGrantPolicyIncludeId is a struct representing the Resource ID for a Policy Permission Grant Policy Include
type PolicyPermissionGrantPolicyIncludeId struct {
	PermissionGrantPolicyId       string
	PermissionGrantConditionSetId string
}

// NewPolicyPermissionGrantPolicyIncludeID returns a new PolicyPermissionGrantPolicyIncludeId struct
func NewPolicyPermissionGrantPolicyIncludeID(permissionGrantPolicyId string, permissionGrantConditionSetId string) PolicyPermissionGrantPolicyIncludeId {
	return PolicyPermissionGrantPolicyIncludeId{
		PermissionGrantPolicyId:       permissionGrantPolicyId,
		PermissionGrantConditionSetId: permissionGrantConditionSetId,
	}
}

// ParsePolicyPermissionGrantPolicyIncludeID parses 'input' into a PolicyPermissionGrantPolicyIncludeId
func ParsePolicyPermissionGrantPolicyIncludeID(input string) (*PolicyPermissionGrantPolicyIncludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyPermissionGrantPolicyIncludeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyPermissionGrantPolicyIncludeId{}

	if id.PermissionGrantPolicyId, ok = parsed.Parsed["permissionGrantPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPolicyId", *parsed)
	}

	if id.PermissionGrantConditionSetId, ok = parsed.Parsed["permissionGrantConditionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantConditionSetId", *parsed)
	}

	return &id, nil
}

// ParsePolicyPermissionGrantPolicyIncludeIDInsensitively parses 'input' case-insensitively into a PolicyPermissionGrantPolicyIncludeId
// note: this method should only be used for API response data and not user input
func ParsePolicyPermissionGrantPolicyIncludeIDInsensitively(input string) (*PolicyPermissionGrantPolicyIncludeId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyPermissionGrantPolicyIncludeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyPermissionGrantPolicyIncludeId{}

	if id.PermissionGrantPolicyId, ok = parsed.Parsed["permissionGrantPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantPolicyId", *parsed)
	}

	if id.PermissionGrantConditionSetId, ok = parsed.Parsed["permissionGrantConditionSetId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "permissionGrantConditionSetId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyPermissionGrantPolicyIncludeID checks that 'input' can be parsed as a Policy Permission Grant Policy Include ID
func ValidatePolicyPermissionGrantPolicyIncludeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyPermissionGrantPolicyIncludeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Permission Grant Policy Include ID
func (id PolicyPermissionGrantPolicyIncludeId) ID() string {
	fmtString := "/policies/permissionGrantPolicies/%s/includes/%s"
	return fmt.Sprintf(fmtString, id.PermissionGrantPolicyId, id.PermissionGrantConditionSetId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Permission Grant Policy Include ID
func (id PolicyPermissionGrantPolicyIncludeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticPermissionGrantPolicies", "permissionGrantPolicies", "permissionGrantPolicies"),
		resourceids.UserSpecifiedSegment("permissionGrantPolicyId", "permissionGrantPolicyIdValue"),
		resourceids.StaticSegment("staticIncludes", "includes", "includes"),
		resourceids.UserSpecifiedSegment("permissionGrantConditionSetId", "permissionGrantConditionSetIdValue"),
	}
}

// String returns a human-readable description of this Policy Permission Grant Policy Include ID
func (id PolicyPermissionGrantPolicyIncludeId) String() string {
	components := []string{
		fmt.Sprintf("Permission Grant Policy: %q", id.PermissionGrantPolicyId),
		fmt.Sprintf("Permission Grant Condition Set: %q", id.PermissionGrantConditionSetId),
	}
	return fmt.Sprintf("Policy Permission Grant Policy Include (%s)", strings.Join(components, "\n"))
}
