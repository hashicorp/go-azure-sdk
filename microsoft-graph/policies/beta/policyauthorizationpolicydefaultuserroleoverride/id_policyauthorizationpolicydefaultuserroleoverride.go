package policyauthorizationpolicydefaultuserroleoverride

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyAuthorizationPolicyDefaultUserRoleOverrideId{}

// PolicyAuthorizationPolicyDefaultUserRoleOverrideId is a struct representing the Resource ID for a Policy Authorization Policy Default User Role Override
type PolicyAuthorizationPolicyDefaultUserRoleOverrideId struct {
	AuthorizationPolicyId     string
	DefaultUserRoleOverrideId string
}

// NewPolicyAuthorizationPolicyDefaultUserRoleOverrideID returns a new PolicyAuthorizationPolicyDefaultUserRoleOverrideId struct
func NewPolicyAuthorizationPolicyDefaultUserRoleOverrideID(authorizationPolicyId string, defaultUserRoleOverrideId string) PolicyAuthorizationPolicyDefaultUserRoleOverrideId {
	return PolicyAuthorizationPolicyDefaultUserRoleOverrideId{
		AuthorizationPolicyId:     authorizationPolicyId,
		DefaultUserRoleOverrideId: defaultUserRoleOverrideId,
	}
}

// ParsePolicyAuthorizationPolicyDefaultUserRoleOverrideID parses 'input' into a PolicyAuthorizationPolicyDefaultUserRoleOverrideId
func ParsePolicyAuthorizationPolicyDefaultUserRoleOverrideID(input string) (*PolicyAuthorizationPolicyDefaultUserRoleOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyAuthorizationPolicyDefaultUserRoleOverrideId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyAuthorizationPolicyDefaultUserRoleOverrideId{}

	if id.AuthorizationPolicyId, ok = parsed.Parsed["authorizationPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authorizationPolicyId", *parsed)
	}

	if id.DefaultUserRoleOverrideId, ok = parsed.Parsed["defaultUserRoleOverrideId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "defaultUserRoleOverrideId", *parsed)
	}

	return &id, nil
}

// ParsePolicyAuthorizationPolicyDefaultUserRoleOverrideIDInsensitively parses 'input' case-insensitively into a PolicyAuthorizationPolicyDefaultUserRoleOverrideId
// note: this method should only be used for API response data and not user input
func ParsePolicyAuthorizationPolicyDefaultUserRoleOverrideIDInsensitively(input string) (*PolicyAuthorizationPolicyDefaultUserRoleOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyAuthorizationPolicyDefaultUserRoleOverrideId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyAuthorizationPolicyDefaultUserRoleOverrideId{}

	if id.AuthorizationPolicyId, ok = parsed.Parsed["authorizationPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authorizationPolicyId", *parsed)
	}

	if id.DefaultUserRoleOverrideId, ok = parsed.Parsed["defaultUserRoleOverrideId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "defaultUserRoleOverrideId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyAuthorizationPolicyDefaultUserRoleOverrideID checks that 'input' can be parsed as a Policy Authorization Policy Default User Role Override ID
func ValidatePolicyAuthorizationPolicyDefaultUserRoleOverrideID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAuthorizationPolicyDefaultUserRoleOverrideID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Authorization Policy Default User Role Override ID
func (id PolicyAuthorizationPolicyDefaultUserRoleOverrideId) ID() string {
	fmtString := "/policies/authorizationPolicy/%s/defaultUserRoleOverrides/%s"
	return fmt.Sprintf(fmtString, id.AuthorizationPolicyId, id.DefaultUserRoleOverrideId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Authorization Policy Default User Role Override ID
func (id PolicyAuthorizationPolicyDefaultUserRoleOverrideId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticAuthorizationPolicy", "authorizationPolicy", "authorizationPolicy"),
		resourceids.UserSpecifiedSegment("authorizationPolicyId", "authorizationPolicyIdValue"),
		resourceids.StaticSegment("staticDefaultUserRoleOverrides", "defaultUserRoleOverrides", "defaultUserRoleOverrides"),
		resourceids.UserSpecifiedSegment("defaultUserRoleOverrideId", "defaultUserRoleOverrideIdValue"),
	}
}

// String returns a human-readable description of this Policy Authorization Policy Default User Role Override ID
func (id PolicyAuthorizationPolicyDefaultUserRoleOverrideId) String() string {
	components := []string{
		fmt.Sprintf("Authorization Policy: %q", id.AuthorizationPolicyId),
		fmt.Sprintf("Default User Role Override: %q", id.DefaultUserRoleOverrideId),
	}
	return fmt.Sprintf("Policy Authorization Policy Default User Role Override (%s)", strings.Join(components, "\n"))
}
