package iotcentrals

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleId{}

// RoleId is a struct representing the Resource ID for a Role
type RoleId struct {
	BaseURI string
	RoleId  string
}

// NewRoleID returns a new RoleId struct
func NewRoleID(baseURI string, roleId string) RoleId {
	return RoleId{
		BaseURI: strings.TrimSuffix(baseURI, "/"),
		RoleId:  roleId,
	}
}

// ParseRoleID parses 'input' into a RoleId
func ParseRoleID(input string) (*RoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleIDInsensitively parses 'input' case-insensitively into a RoleId
// note: this method should only be used for API response data and not user input
func ParseRoleIDInsensitively(input string) (*RoleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if id.RoleId, ok = input.Parsed["roleId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "roleId", input)
	}

	return nil
}

// ValidateRoleID checks that 'input' can be parsed as a Role ID
func ValidateRoleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role ID
func (id RoleId) ID() string {
	fmtString := "%s/roles/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), id.RoleId)
}

// Path returns the formatted Role ID without the BaseURI
func (id RoleId) Path() string {
	fmtString := "/roles/%s"
	return fmt.Sprintf(fmtString, id.RoleId)
}

// PathElements returns the values of Role ID Segments without the BaseURI
func (id RoleId) PathElements() []any {
	return []any{id.RoleId}
}

// Segments returns a slice of Resource ID Segments which comprise this Role ID
func (id RoleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticRoles", "roles", "roles"),
		resourceids.UserSpecifiedSegment("roleId", "roleId"),
	}
}

// String returns a human-readable description of this Role ID
func (id RoleId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Role: %q", id.RoleId),
	}
	return fmt.Sprintf("Role (%s)", strings.Join(components, "\n"))
}
