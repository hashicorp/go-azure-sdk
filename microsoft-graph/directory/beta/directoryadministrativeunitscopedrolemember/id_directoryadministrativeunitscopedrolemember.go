package directoryadministrativeunitscopedrolemember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryAdministrativeUnitScopedRoleMemberId{}

// DirectoryAdministrativeUnitScopedRoleMemberId is a struct representing the Resource ID for a Directory Administrative Unit Scoped Role Member
type DirectoryAdministrativeUnitScopedRoleMemberId struct {
	AdministrativeUnitId   string
	ScopedRoleMembershipId string
}

// NewDirectoryAdministrativeUnitScopedRoleMemberID returns a new DirectoryAdministrativeUnitScopedRoleMemberId struct
func NewDirectoryAdministrativeUnitScopedRoleMemberID(administrativeUnitId string, scopedRoleMembershipId string) DirectoryAdministrativeUnitScopedRoleMemberId {
	return DirectoryAdministrativeUnitScopedRoleMemberId{
		AdministrativeUnitId:   administrativeUnitId,
		ScopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// ParseDirectoryAdministrativeUnitScopedRoleMemberID parses 'input' into a DirectoryAdministrativeUnitScopedRoleMemberId
func ParseDirectoryAdministrativeUnitScopedRoleMemberID(input string) (*DirectoryAdministrativeUnitScopedRoleMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryAdministrativeUnitScopedRoleMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryAdministrativeUnitScopedRoleMemberId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.ScopedRoleMembershipId, ok = parsed.Parsed["scopedRoleMembershipId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryAdministrativeUnitScopedRoleMemberIDInsensitively parses 'input' case-insensitively into a DirectoryAdministrativeUnitScopedRoleMemberId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAdministrativeUnitScopedRoleMemberIDInsensitively(input string) (*DirectoryAdministrativeUnitScopedRoleMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryAdministrativeUnitScopedRoleMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryAdministrativeUnitScopedRoleMemberId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.ScopedRoleMembershipId, ok = parsed.Parsed["scopedRoleMembershipId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryAdministrativeUnitScopedRoleMemberID checks that 'input' can be parsed as a Directory Administrative Unit Scoped Role Member ID
func ValidateDirectoryAdministrativeUnitScopedRoleMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAdministrativeUnitScopedRoleMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Administrative Unit Scoped Role Member ID
func (id DirectoryAdministrativeUnitScopedRoleMemberId) ID() string {
	fmtString := "/directory/administrativeUnits/%s/scopedRoleMembers/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.ScopedRoleMembershipId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Administrative Unit Scoped Role Member ID
func (id DirectoryAdministrativeUnitScopedRoleMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticAdministrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitIdValue"),
		resourceids.StaticSegment("staticScopedRoleMembers", "scopedRoleMembers", "scopedRoleMembers"),
		resourceids.UserSpecifiedSegment("scopedRoleMembershipId", "scopedRoleMembershipIdValue"),
	}
}

// String returns a human-readable description of this Directory Administrative Unit Scoped Role Member ID
func (id DirectoryAdministrativeUnitScopedRoleMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Scoped Role Membership: %q", id.ScopedRoleMembershipId),
	}
	return fmt.Sprintf("Directory Administrative Unit Scoped Role Member (%s)", strings.Join(components, "\n"))
}
