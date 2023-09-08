package rolemanagementdirectoryroledefinitioninheritspermissionsfrom

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDirectoryRoleDefinitionId{}

// RoleManagementDirectoryRoleDefinitionId is a struct representing the Resource ID for a Role Management Directory Role Definition
type RoleManagementDirectoryRoleDefinitionId struct {
	UnifiedRoleDefinitionId string
}

// NewRoleManagementDirectoryRoleDefinitionID returns a new RoleManagementDirectoryRoleDefinitionId struct
func NewRoleManagementDirectoryRoleDefinitionID(unifiedRoleDefinitionId string) RoleManagementDirectoryRoleDefinitionId {
	return RoleManagementDirectoryRoleDefinitionId{
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementDirectoryRoleDefinitionID parses 'input' into a RoleManagementDirectoryRoleDefinitionId
func ParseRoleManagementDirectoryRoleDefinitionID(input string) (*RoleManagementDirectoryRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryRoleDefinitionId{}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleDefinitionIDInsensitively(input string) (*RoleManagementDirectoryRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryRoleDefinitionId{}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDirectoryRoleDefinitionID checks that 'input' can be parsed as a Role Management Directory Role Definition ID
func ValidateRoleManagementDirectoryRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Definition ID
func (id RoleManagementDirectoryRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/directory/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Definition ID
func (id RoleManagementDirectoryRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticRoleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Definition ID
func (id RoleManagementDirectoryRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Directory Role Definition (%s)", strings.Join(components, "\n"))
}
