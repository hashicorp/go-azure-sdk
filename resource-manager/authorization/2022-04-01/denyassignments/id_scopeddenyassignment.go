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
	recaser.RegisterResourceId(&ScopedDenyAssignmentId{})
}

var _ resourceids.ResourceId = &ScopedDenyAssignmentId{}

// ScopedDenyAssignmentId is a struct representing the Resource ID for a Scoped Deny Assignment
type ScopedDenyAssignmentId struct {
	Scope            string
	DenyAssignmentId string
}

// NewScopedDenyAssignmentID returns a new ScopedDenyAssignmentId struct
func NewScopedDenyAssignmentID(scope string, denyAssignmentId string) ScopedDenyAssignmentId {
	return ScopedDenyAssignmentId{
		Scope:            scope,
		DenyAssignmentId: denyAssignmentId,
	}
}

// ParseScopedDenyAssignmentID parses 'input' into a ScopedDenyAssignmentId
func ParseScopedDenyAssignmentID(input string) (*ScopedDenyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedDenyAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedDenyAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedDenyAssignmentIDInsensitively parses 'input' case-insensitively into a ScopedDenyAssignmentId
// note: this method should only be used for API response data and not user input
func ParseScopedDenyAssignmentIDInsensitively(input string) (*ScopedDenyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedDenyAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedDenyAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedDenyAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.DenyAssignmentId, ok = input.Parsed["denyAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "denyAssignmentId", input)
	}

	return nil
}

// ValidateScopedDenyAssignmentID checks that 'input' can be parsed as a Scoped Deny Assignment ID
func ValidateScopedDenyAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedDenyAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Deny Assignment ID
func (id ScopedDenyAssignmentId) ID() string {
	fmtString := "/%s/providers/Microsoft.Authorization/denyAssignments/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.DenyAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Deny Assignment ID
func (id ScopedDenyAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticDenyAssignments", "denyAssignments", "denyAssignments"),
		resourceids.UserSpecifiedSegment("denyAssignmentId", "denyAssignmentId"),
	}
}

// String returns a human-readable description of this Scoped Deny Assignment ID
func (id ScopedDenyAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Deny Assignment: %q", id.DenyAssignmentId),
	}
	return fmt.Sprintf("Scoped Deny Assignment (%s)", strings.Join(components, "\n"))
}
