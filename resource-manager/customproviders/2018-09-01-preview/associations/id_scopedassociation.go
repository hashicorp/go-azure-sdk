package associations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopedAssociationId{}

// ScopedAssociationId is a struct representing the Resource ID for a Scoped Association
type ScopedAssociationId struct {
	Scope           string
	AssociationName string
}

// NewScopedAssociationID returns a new ScopedAssociationId struct
func NewScopedAssociationID(scope string, associationName string) ScopedAssociationId {
	return ScopedAssociationId{
		Scope:           scope,
		AssociationName: associationName,
	}
}

// ParseScopedAssociationID parses 'input' into a ScopedAssociationId
func ParseScopedAssociationID(input string) (*ScopedAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedAssociationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedAssociationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedAssociationIDInsensitively parses 'input' case-insensitively into a ScopedAssociationId
// note: this method should only be used for API response data and not user input
func ParseScopedAssociationIDInsensitively(input string) (*ScopedAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedAssociationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedAssociationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedAssociationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.AssociationName, ok = input.Parsed["associationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "associationName", input)
	}

	return nil
}

// ValidateScopedAssociationID checks that 'input' can be parsed as a Scoped Association ID
func ValidateScopedAssociationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedAssociationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Association ID
func (id ScopedAssociationId) ID() string {
	fmtString := "/%s/providers/Microsoft.CustomProviders/associations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.AssociationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Association ID
func (id ScopedAssociationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCustomProviders", "Microsoft.CustomProviders", "Microsoft.CustomProviders"),
		resourceids.StaticSegment("staticAssociations", "associations", "associations"),
		resourceids.UserSpecifiedSegment("associationName", "associationValue"),
	}
}

// String returns a human-readable description of this Scoped Association ID
func (id ScopedAssociationId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Association Name: %q", id.AssociationName),
	}
	return fmt.Sprintf("Scoped Association (%s)", strings.Join(components, "\n"))
}
