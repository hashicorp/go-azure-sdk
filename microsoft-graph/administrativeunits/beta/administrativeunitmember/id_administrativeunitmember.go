package administrativeunitmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AdministrativeUnitMemberId{}

// AdministrativeUnitMemberId is a struct representing the Resource ID for a Administrative Unit Member
type AdministrativeUnitMemberId struct {
	AdministrativeUnitId string
	DirectoryObjectId    string
}

// NewAdministrativeUnitMemberID returns a new AdministrativeUnitMemberId struct
func NewAdministrativeUnitMemberID(administrativeUnitId string, directoryObjectId string) AdministrativeUnitMemberId {
	return AdministrativeUnitMemberId{
		AdministrativeUnitId: administrativeUnitId,
		DirectoryObjectId:    directoryObjectId,
	}
}

// ParseAdministrativeUnitMemberID parses 'input' into a AdministrativeUnitMemberId
func ParseAdministrativeUnitMemberID(input string) (*AdministrativeUnitMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdministrativeUnitMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AdministrativeUnitMemberId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseAdministrativeUnitMemberIDInsensitively parses 'input' case-insensitively into a AdministrativeUnitMemberId
// note: this method should only be used for API response data and not user input
func ParseAdministrativeUnitMemberIDInsensitively(input string) (*AdministrativeUnitMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdministrativeUnitMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AdministrativeUnitMemberId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateAdministrativeUnitMemberID checks that 'input' can be parsed as a Administrative Unit Member ID
func ValidateAdministrativeUnitMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdministrativeUnitMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Administrative Unit Member ID
func (id AdministrativeUnitMemberId) ID() string {
	fmtString := "/administrativeUnits/%s/members/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Administrative Unit Member ID
func (id AdministrativeUnitMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticAdministrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Administrative Unit Member ID
func (id AdministrativeUnitMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Administrative Unit Member (%s)", strings.Join(components, "\n"))
}
