package roleassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleAssignmentId{}

// RoleAssignmentId is a struct representing the Resource ID for a Role Assignment
type RoleAssignmentId struct {
	BaseURI          string
	RoleAssignmentId string
}

// NewRoleAssignmentID returns a new RoleAssignmentId struct
func NewRoleAssignmentID(baseURI string, roleAssignmentId string) RoleAssignmentId {
	return RoleAssignmentId{
		BaseURI:          strings.TrimSuffix(baseURI, "/"),
		RoleAssignmentId: roleAssignmentId,
	}
}

// ParseRoleAssignmentID parses 'input' into a RoleAssignmentId
func ParseRoleAssignmentID(input string) (*RoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleAssignmentIDInsensitively(input string) (*RoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.RoleAssignmentId, ok = input.Parsed["roleAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleAssignmentId", input)
	}

	return nil
}

// ValidateRoleAssignmentID checks that 'input' can be parsed as a Role Assignment ID
func ValidateRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Assignment ID
func (id RoleAssignmentId) ID() string {
	fmtString := "%s/roleAssignments/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.RoleAssignmentId)
}

// Path returns the formatted Role Assignment ID without the BaseURI
func (id RoleAssignmentId) Path() string {
	fmtString := "/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.RoleAssignmentId)
}

// PathElements returns the values of Role Assignment ID Segments without the BaseURI
func (id RoleAssignmentId) PathElements() []any {
	return []any{id.RoleAssignmentId}
}

// Segments returns a slice of Resource ID Segments which comprise this Role Assignment ID
func (id RoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticRoleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("roleAssignmentId", "roleAssignmentId"),
	}
}

// String returns a human-readable description of this Role Assignment ID
func (id RoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Role Assignment: %q", id.RoleAssignmentId),
	}
	return fmt.Sprintf("Role Assignment (%s)", strings.Join(components, "\n"))
}
