package administrativeunitscopedrolemember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AdministrativeUnitScopedRoleMemberId{}

// AdministrativeUnitScopedRoleMemberId is a struct representing the Resource ID for a Administrative Unit Scoped Role Member
type AdministrativeUnitScopedRoleMemberId struct {
	AdministrativeUnitId   string
	ScopedRoleMembershipId string
}

// NewAdministrativeUnitScopedRoleMemberID returns a new AdministrativeUnitScopedRoleMemberId struct
func NewAdministrativeUnitScopedRoleMemberID(administrativeUnitId string, scopedRoleMembershipId string) AdministrativeUnitScopedRoleMemberId {
	return AdministrativeUnitScopedRoleMemberId{
		AdministrativeUnitId:   administrativeUnitId,
		ScopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// ParseAdministrativeUnitScopedRoleMemberID parses 'input' into a AdministrativeUnitScopedRoleMemberId
func ParseAdministrativeUnitScopedRoleMemberID(input string) (*AdministrativeUnitScopedRoleMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdministrativeUnitScopedRoleMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AdministrativeUnitScopedRoleMemberId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.ScopedRoleMembershipId, ok = parsed.Parsed["scopedRoleMembershipId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", *parsed)
	}

	return &id, nil
}

// ParseAdministrativeUnitScopedRoleMemberIDInsensitively parses 'input' case-insensitively into a AdministrativeUnitScopedRoleMemberId
// note: this method should only be used for API response data and not user input
func ParseAdministrativeUnitScopedRoleMemberIDInsensitively(input string) (*AdministrativeUnitScopedRoleMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdministrativeUnitScopedRoleMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AdministrativeUnitScopedRoleMemberId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.ScopedRoleMembershipId, ok = parsed.Parsed["scopedRoleMembershipId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", *parsed)
	}

	return &id, nil
}

// ValidateAdministrativeUnitScopedRoleMemberID checks that 'input' can be parsed as a Administrative Unit Scoped Role Member ID
func ValidateAdministrativeUnitScopedRoleMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdministrativeUnitScopedRoleMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Administrative Unit Scoped Role Member ID
func (id AdministrativeUnitScopedRoleMemberId) ID() string {
	fmtString := "/administrativeUnits/%s/scopedRoleMembers/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.ScopedRoleMembershipId)
}

// Segments returns a slice of Resource ID Segments which comprise this Administrative Unit Scoped Role Member ID
func (id AdministrativeUnitScopedRoleMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticAdministrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitIdValue"),
		resourceids.StaticSegment("staticScopedRoleMembers", "scopedRoleMembers", "scopedRoleMembers"),
		resourceids.UserSpecifiedSegment("scopedRoleMembershipId", "scopedRoleMembershipIdValue"),
	}
}

// String returns a human-readable description of this Administrative Unit Scoped Role Member ID
func (id AdministrativeUnitScopedRoleMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Scoped Role Membership: %q", id.ScopedRoleMembershipId),
	}
	return fmt.Sprintf("Administrative Unit Scoped Role Member (%s)", strings.Join(components, "\n"))
}
