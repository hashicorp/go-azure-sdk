package managementlocks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ResourceLockScopedId{}

// ResourceLockScopedId is a struct representing the Resource ID for a Resource Lock Scoped
type ResourceLockScopedId struct {
	Scope        string
	ResourceName string
	LockName     string
}

// NewResourceLockScopedID returns a new ResourceLockScopedId struct
func NewResourceLockScopedID(scope string, resourceName string, lockName string) ResourceLockScopedId {
	return ResourceLockScopedId{
		Scope:        scope,
		ResourceName: resourceName,
		LockName:     lockName,
	}
}

// ParseResourceLockScopedID parses 'input' into a ResourceLockScopedId
func ParseResourceLockScopedID(input string) (*ResourceLockScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResourceLockScopedId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ResourceLockScopedId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	if id.LockName, ok = parsed.Parsed["lockName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "lockName", *parsed)
	}

	return &id, nil
}

// ParseResourceLockScopedIDInsensitively parses 'input' case-insensitively into a ResourceLockScopedId
// note: this method should only be used for API response data and not user input
func ParseResourceLockScopedIDInsensitively(input string) (*ResourceLockScopedId, error) {
	parser := resourceids.NewParserFromResourceIdType(ResourceLockScopedId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ResourceLockScopedId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	if id.LockName, ok = parsed.Parsed["lockName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "lockName", *parsed)
	}

	return &id, nil
}

// ValidateResourceLockScopedID checks that 'input' can be parsed as a Resource Lock Scoped ID
func ValidateResourceLockScopedID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResourceLockScopedID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resource Lock Scoped ID
func (id ResourceLockScopedId) ID() string {
	fmtString := "/%s/%s/providers/Microsoft.Authorization/locks/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.ResourceName, id.LockName)
}

// Segments returns a slice of Resource ID Segments which comprise this Resource Lock Scoped ID
func (id ResourceLockScopedId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticLocks", "locks", "locks"),
		resourceids.UserSpecifiedSegment("lockName", "lockValue"),
	}
}

// String returns a human-readable description of this Resource Lock Scoped ID
func (id ResourceLockScopedId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Lock Name: %q", id.LockName),
	}
	return fmt.Sprintf("Resource Lock Scoped (%s)", strings.Join(components, "\n"))
}
