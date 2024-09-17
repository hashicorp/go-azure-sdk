package managementassociation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedManagementAssociationId{})
}

var _ resourceids.ResourceId = &ScopedManagementAssociationId{}

// ScopedManagementAssociationId is a struct representing the Resource ID for a Scoped Management Association
type ScopedManagementAssociationId struct {
	Scope                     string
	ManagementAssociationName string
}

// NewScopedManagementAssociationID returns a new ScopedManagementAssociationId struct
func NewScopedManagementAssociationID(scope string, managementAssociationName string) ScopedManagementAssociationId {
	return ScopedManagementAssociationId{
		Scope:                     scope,
		ManagementAssociationName: managementAssociationName,
	}
}

// ParseScopedManagementAssociationID parses 'input' into a ScopedManagementAssociationId
func ParseScopedManagementAssociationID(input string) (*ScopedManagementAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedManagementAssociationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedManagementAssociationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedManagementAssociationIDInsensitively parses 'input' case-insensitively into a ScopedManagementAssociationId
// note: this method should only be used for API response data and not user input
func ParseScopedManagementAssociationIDInsensitively(input string) (*ScopedManagementAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedManagementAssociationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedManagementAssociationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedManagementAssociationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.ManagementAssociationName, ok = input.Parsed["managementAssociationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementAssociationName", input)
	}

	return nil
}

// ValidateScopedManagementAssociationID checks that 'input' can be parsed as a Scoped Management Association ID
func ValidateScopedManagementAssociationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedManagementAssociationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Management Association ID
func (id ScopedManagementAssociationId) ID() string {
	fmtString := "/%s/providers/Microsoft.OperationsManagement/managementAssociations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.ManagementAssociationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Management Association ID
func (id ScopedManagementAssociationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationsManagement", "Microsoft.OperationsManagement", "Microsoft.OperationsManagement"),
		resourceids.StaticSegment("staticManagementAssociations", "managementAssociations", "managementAssociations"),
		resourceids.UserSpecifiedSegment("managementAssociationName", "managementAssociationValue"),
	}
}

// String returns a human-readable description of this Scoped Management Association ID
func (id ScopedManagementAssociationId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Management Association Name: %q", id.ManagementAssociationName),
	}
	return fmt.Sprintf("Scoped Management Association (%s)", strings.Join(components, "\n"))
}
