package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{}

// IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId is a struct representing the Resource ID for a Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Role
type IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId struct {
	AccessPackageResourceRoleScopeId string
	AccessPackageResourceRoleId      string
}

// NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID returns a new IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId struct
func NewIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID(accessPackageResourceRoleScopeId string, accessPackageResourceRoleId string) IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId {
	return IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
		AccessPackageResourceRoleId:      accessPackageResourceRoleId,
	}
}

// ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID parses 'input' into a IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId
func ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID(input string) (*IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	if id.AccessPackageResourceRoleId, ok = input.Parsed["accessPackageResourceRoleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID checks that 'input' can be parsed as a Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Role ID
func ValidateIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/resourceRoleScopes/%s/scope/resource/roles/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageResourceRoleScopeId, id.AccessPackageResourceRoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("resourceRoleScopes", "resourceRoleScopes", "resourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
		resourceids.StaticSegment("scope", "scope", "scope"),
		resourceids.StaticSegment("resource", "resource", "resource"),
		resourceids.StaticSegment("roles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleId", "accessPackageResourceRoleId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Role ID
func (id IdentityGovernanceEntitlementManagementResourceRoleScopeIdScopeResourceRoleId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
		fmt.Sprintf("Access Package Resource Role: %q", id.AccessPackageResourceRoleId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Resource Role Scope Id Scope Resource Role (%s)", strings.Join(components, "\n"))
}
