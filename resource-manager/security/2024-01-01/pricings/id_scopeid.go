package pricings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopeIdId{}

// ScopeIdId is a struct representing the Resource ID for a Scope Id
type ScopeIdId struct {
	ScopeId string
}

// NewScopeIdID returns a new ScopeIdId struct
func NewScopeIdID(scopeId string) ScopeIdId {
	return ScopeIdId{
		ScopeId: scopeId,
	}
}

// ParseScopeIdID parses 'input' into a ScopeIdId
func ParseScopeIdID(input string) (*ScopeIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopeIdId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopeIdId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopeIdIDInsensitively parses 'input' case-insensitively into a ScopeIdId
// note: this method should only be used for API response data and not user input
func ParseScopeIdIDInsensitively(input string) (*ScopeIdId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopeIdId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopeIdId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopeIdId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ScopeId, ok = input.Parsed["scopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scopeId", input)
	}

	return nil
}

// ValidateScopeIdID checks that 'input' can be parsed as a Scope Id ID
func ValidateScopeIdID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopeIdID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scope Id ID
func (id ScopeIdId) ID() string {
	fmtString := "/%s"
	return fmt.Sprintf(fmtString, id.ScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Scope Id ID
func (id ScopeIdId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.UserSpecifiedSegment("scopeId", "scopeIdValue"),
	}
}

// String returns a human-readable description of this Scope Id ID
func (id ScopeIdId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.ScopeId),
	}
	return fmt.Sprintf("Scope Id (%s)", strings.Join(components, "\n"))
}
