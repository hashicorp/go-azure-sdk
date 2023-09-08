package rolemanagementexchangeroledefinitioninheritspermissionsfrom

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementExchangeRoleDefinitionId{}

// RoleManagementExchangeRoleDefinitionId is a struct representing the Resource ID for a Role Management Exchange Role Definition
type RoleManagementExchangeRoleDefinitionId struct {
	UnifiedRoleDefinitionId string
}

// NewRoleManagementExchangeRoleDefinitionID returns a new RoleManagementExchangeRoleDefinitionId struct
func NewRoleManagementExchangeRoleDefinitionID(unifiedRoleDefinitionId string) RoleManagementExchangeRoleDefinitionId {
	return RoleManagementExchangeRoleDefinitionId{
		UnifiedRoleDefinitionId: unifiedRoleDefinitionId,
	}
}

// ParseRoleManagementExchangeRoleDefinitionID parses 'input' into a RoleManagementExchangeRoleDefinitionId
func ParseRoleManagementExchangeRoleDefinitionID(input string) (*RoleManagementExchangeRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementExchangeRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementExchangeRoleDefinitionId{}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementExchangeRoleDefinitionIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeRoleDefinitionIDInsensitively(input string) (*RoleManagementExchangeRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementExchangeRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementExchangeRoleDefinitionId{}

	if id.UnifiedRoleDefinitionId, ok = parsed.Parsed["unifiedRoleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementExchangeRoleDefinitionID checks that 'input' can be parsed as a Role Management Exchange Role Definition ID
func ValidateRoleManagementExchangeRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Role Definition ID
func (id RoleManagementExchangeRoleDefinitionId) ID() string {
	fmtString := "/roleManagement/exchange/roleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Role Definition ID
func (id RoleManagementExchangeRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticExchange", "exchange", "exchange"),
		resourceids.StaticSegment("staticRoleDefinitions", "roleDefinitions", "roleDefinitions"),
		resourceids.UserSpecifiedSegment("unifiedRoleDefinitionId", "unifiedRoleDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Role Management Exchange Role Definition ID
func (id RoleManagementExchangeRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Definition: %q", id.UnifiedRoleDefinitionId),
	}
	return fmt.Sprintf("Role Management Exchange Role Definition (%s)", strings.Join(components, "\n"))
}
