package applyupdates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedApplyUpdateId{}

// ScopedApplyUpdateId is a struct representing the Resource ID for a Scoped Apply Update
type ScopedApplyUpdateId struct {
	Scope           string
	ApplyUpdateName string
}

// NewScopedApplyUpdateID returns a new ScopedApplyUpdateId struct
func NewScopedApplyUpdateID(scope string, applyUpdateName string) ScopedApplyUpdateId {
	return ScopedApplyUpdateId{
		Scope:           scope,
		ApplyUpdateName: applyUpdateName,
	}
}

// ParseScopedApplyUpdateID parses 'input' into a ScopedApplyUpdateId
func ParseScopedApplyUpdateID(input string) (*ScopedApplyUpdateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedApplyUpdateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedApplyUpdateId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedApplyUpdateIDInsensitively parses 'input' case-insensitively into a ScopedApplyUpdateId
// note: this method should only be used for API response data and not user input
func ParseScopedApplyUpdateIDInsensitively(input string) (*ScopedApplyUpdateId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedApplyUpdateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedApplyUpdateId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedApplyUpdateId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.ApplyUpdateName, ok = input.Parsed["applyUpdateName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "applyUpdateName", input)
	}

	return nil
}

// ValidateScopedApplyUpdateID checks that 'input' can be parsed as a Scoped Apply Update ID
func ValidateScopedApplyUpdateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedApplyUpdateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Apply Update ID
func (id ScopedApplyUpdateId) ID() string {
	fmtString := "/%s/providers/Microsoft.Maintenance/applyUpdates/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.ApplyUpdateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Apply Update ID
func (id ScopedApplyUpdateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMaintenance", "Microsoft.Maintenance", "Microsoft.Maintenance"),
		resourceids.StaticSegment("staticApplyUpdates", "applyUpdates", "applyUpdates"),
		resourceids.UserSpecifiedSegment("applyUpdateName", "applyUpdateValue"),
	}
}

// String returns a human-readable description of this Scoped Apply Update ID
func (id ScopedApplyUpdateId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Apply Update Name: %q", id.ApplyUpdateName),
	}
	return fmt.Sprintf("Scoped Apply Update (%s)", strings.Join(components, "\n"))
}
