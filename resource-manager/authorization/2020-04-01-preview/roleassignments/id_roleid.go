package roleassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleIdId{}

// RoleIdId is a struct representing the Resource ID for a Role Id
type RoleIdId struct {
	RoleId string
}

// NewRoleIdID returns a new RoleIdId struct
func NewRoleIdID(roleId string) RoleIdId {
	return RoleIdId{
		RoleId: roleId,
	}
}

// ParseRoleIdID parses 'input' into a RoleIdId
func ParseRoleIdID(input string) (*RoleIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleIdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleIdId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleIdIDInsensitively parses 'input' case-insensitively into a RoleIdId
// note: this method should only be used for API response data and not user input
func ParseRoleIdIDInsensitively(input string) (*RoleIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleIdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleIdId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleIdId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RoleId, ok = input.Parsed["roleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleId", input)
	}

	return nil
}

// ValidateRoleIdID checks that 'input' can be parsed as a Role Id ID
func ValidateRoleIdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleIdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Id ID
func (id RoleIdId) ID() string {
	fmtString := "/%s"
	return fmt.Sprintf(fmtString, id.RoleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Id ID
func (id RoleIdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.UserSpecifiedSegment("roleId", "roleIdValue"),
	}
}

// String returns a human-readable description of this Role Id ID
func (id RoleIdId) String() string {
	components := []string{
		fmt.Sprintf("Role: %q", id.RoleId),
	}
	return fmt.Sprintf("Role Id (%s)", strings.Join(components, "\n"))
}
