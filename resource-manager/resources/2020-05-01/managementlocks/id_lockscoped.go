package managementlocks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = LockScopedId{}

// LockScopedId is a struct representing the Resource ID for a Lock Scoped
type LockScopedId struct {
	Scope    string
	LockName string
}

// NewLockScopedID returns a new LockScopedId struct
func NewLockScopedID(scope string, lockName string) LockScopedId {
	return LockScopedId{
		Scope:    scope,
		LockName: lockName,
	}
}

// ParseLockScopedID parses 'input' into a LockScopedId
func ParseLockScopedID(input string) (*LockScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(LockScopedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LockScopedId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.LockName, ok = parsed.Parsed["lockName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "lockName", *parsed)
	}

	return &id, nil
}

// ParseLockScopedIDInsensitively parses 'input' case-insensitively into a LockScopedId
// note: this method should only be used for API response data and not user input
func ParseLockScopedIDInsensitively(input string) (*LockScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(LockScopedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LockScopedId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.LockName, ok = parsed.Parsed["lockName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "lockName", *parsed)
	}

	return &id, nil
}

// ValidateLockScopedID checks that 'input' can be parsed as a Lock Scoped ID
func ValidateLockScopedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLockScopedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Lock Scoped ID
func (id LockScopedId) ID() string {
	fmtString := "/%s/providers/Microsoft.Authorization/locks/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.LockName)
}

// Segments returns a slice of Resource ID Segments which comprise this Lock Scoped ID
func (id LockScopedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticLocks", "locks", "locks"),
		resourceids.UserSpecifiedSegment("lockName", "lockValue"),
	}
}

// String returns a human-readable description of this Lock Scoped ID
func (id LockScopedId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Lock Name: %q", id.LockName),
	}
	return fmt.Sprintf("Lock Scoped (%s)", strings.Join(components, "\n"))
}
