package permissions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopedResourceId{}

// ScopedResourceId is a struct representing the Resource ID for a Scoped Resource
type ScopedResourceId struct {
	Scope        string
	ResourceName string
}

// NewScopedResourceID returns a new ScopedResourceId struct
func NewScopedResourceID(scope string, resourceName string) ScopedResourceId {
	return ScopedResourceId{
		Scope:        scope,
		ResourceName: resourceName,
	}
}

// ParseScopedResourceID parses 'input' into a ScopedResourceId
func ParseScopedResourceID(input string) (*ScopedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedResourceId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	return &id, nil
}

// ParseScopedResourceIDInsensitively parses 'input' case-insensitively into a ScopedResourceId
// note: this method should only be used for API response data and not user input
func ParseScopedResourceIDInsensitively(input string) (*ScopedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedResourceId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	return &id, nil
}

// ValidateScopedResourceID checks that 'input' can be parsed as a Scoped Resource ID
func ValidateScopedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Resource ID
func (id ScopedResourceId) ID() string {
	fmtString := "/%s/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.ResourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Resource ID
func (id ScopedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
	}
}

// String returns a human-readable description of this Scoped Resource ID
func (id ScopedResourceId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
	}
	return fmt.Sprintf("Scoped Resource (%s)", strings.Join(components, "\n"))
}
