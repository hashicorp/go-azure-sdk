package denyassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedDenyAssignmentIdId{})
}

var _ resourceids.ResourceId = &ScopedDenyAssignmentIdId{}

// ScopedDenyAssignmentIdId is a struct representing the Resource ID for a Scoped Deny Assignment Id
type ScopedDenyAssignmentIdId struct {
	Scope            string
	DenyAssignmentId string
}

// NewScopedDenyAssignmentIdID returns a new ScopedDenyAssignmentIdId struct
func NewScopedDenyAssignmentIdID(scope string, denyAssignmentId string) ScopedDenyAssignmentIdId {
	return ScopedDenyAssignmentIdId{
		Scope:            scope,
		DenyAssignmentId: denyAssignmentId,
	}
}

// ParseScopedDenyAssignmentIdID parses 'input' into a ScopedDenyAssignmentIdId
func ParseScopedDenyAssignmentIdID(input string) (*ScopedDenyAssignmentIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedDenyAssignmentIdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedDenyAssignmentIdId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedDenyAssignmentIdIDInsensitively parses 'input' case-insensitively into a ScopedDenyAssignmentIdId
// note: this method should only be used for API response data and not user input
func ParseScopedDenyAssignmentIdIDInsensitively(input string) (*ScopedDenyAssignmentIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedDenyAssignmentIdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedDenyAssignmentIdId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedDenyAssignmentIdId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.DenyAssignmentId, ok = input.Parsed["denyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "denyAssignmentId", input)
	}

	return nil
}

// ValidateScopedDenyAssignmentIdID checks that 'input' can be parsed as a Scoped Deny Assignment Id ID
func ValidateScopedDenyAssignmentIdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedDenyAssignmentIdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Deny Assignment Id ID
func (id ScopedDenyAssignmentIdId) ID() string {
	fmtString := "/%s/providers/Microsoft.Authorization/denyAssignments/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), strings.TrimPrefix(id.DenyAssignmentId, "/"))
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Deny Assignment Id ID
func (id ScopedDenyAssignmentIdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticDenyAssignments", "denyAssignments", "denyAssignments"),
		resourceids.ScopeSegment("denyAssignmentId", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
	}
}

// String returns a human-readable description of this Scoped Deny Assignment Id ID
func (id ScopedDenyAssignmentIdId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Deny Assignment: %q", id.DenyAssignmentId),
	}
	return fmt.Sprintf("Scoped Deny Assignment Id (%s)", strings.Join(components, "\n"))
}
