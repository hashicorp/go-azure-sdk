package directoryadministrativeunitmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryAdministrativeUnitMemberId{}

// DirectoryAdministrativeUnitMemberId is a struct representing the Resource ID for a Directory Administrative Unit Member
type DirectoryAdministrativeUnitMemberId struct {
	AdministrativeUnitId string
	DirectoryObjectId    string
}

// NewDirectoryAdministrativeUnitMemberID returns a new DirectoryAdministrativeUnitMemberId struct
func NewDirectoryAdministrativeUnitMemberID(administrativeUnitId string, directoryObjectId string) DirectoryAdministrativeUnitMemberId {
	return DirectoryAdministrativeUnitMemberId{
		AdministrativeUnitId: administrativeUnitId,
		DirectoryObjectId:    directoryObjectId,
	}
}

// ParseDirectoryAdministrativeUnitMemberID parses 'input' into a DirectoryAdministrativeUnitMemberId
func ParseDirectoryAdministrativeUnitMemberID(input string) (*DirectoryAdministrativeUnitMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryAdministrativeUnitMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryAdministrativeUnitMemberId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryAdministrativeUnitMemberIDInsensitively parses 'input' case-insensitively into a DirectoryAdministrativeUnitMemberId
// note: this method should only be used for API response data and not user input
func ParseDirectoryAdministrativeUnitMemberIDInsensitively(input string) (*DirectoryAdministrativeUnitMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryAdministrativeUnitMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryAdministrativeUnitMemberId{}

	if id.AdministrativeUnitId, ok = parsed.Parsed["administrativeUnitId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "administrativeUnitId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryAdministrativeUnitMemberID checks that 'input' can be parsed as a Directory Administrative Unit Member ID
func ValidateDirectoryAdministrativeUnitMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryAdministrativeUnitMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Administrative Unit Member ID
func (id DirectoryAdministrativeUnitMemberId) ID() string {
	fmtString := "/directory/administrativeUnits/%s/members/%s"
	return fmt.Sprintf(fmtString, id.AdministrativeUnitId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Administrative Unit Member ID
func (id DirectoryAdministrativeUnitMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticAdministrativeUnits", "administrativeUnits", "administrativeUnits"),
		resourceids.UserSpecifiedSegment("administrativeUnitId", "administrativeUnitIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Directory Administrative Unit Member ID
func (id DirectoryAdministrativeUnitMemberId) String() string {
	components := []string{
		fmt.Sprintf("Administrative Unit: %q", id.AdministrativeUnitId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Directory Administrative Unit Member (%s)", strings.Join(components, "\n"))
}
