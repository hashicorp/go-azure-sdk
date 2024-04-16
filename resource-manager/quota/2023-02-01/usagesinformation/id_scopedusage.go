package usagesinformation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedUsageId{})
}

var _ resourceids.ResourceId = &ScopedUsageId{}

// ScopedUsageId is a struct representing the Resource ID for a Scoped Usage
type ScopedUsageId struct {
	Scope     string
	UsageName string
}

// NewScopedUsageID returns a new ScopedUsageId struct
func NewScopedUsageID(scope string, usageName string) ScopedUsageId {
	return ScopedUsageId{
		Scope:     scope,
		UsageName: usageName,
	}
}

// ParseScopedUsageID parses 'input' into a ScopedUsageId
func ParseScopedUsageID(input string) (*ScopedUsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedUsageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedUsageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedUsageIDInsensitively parses 'input' case-insensitively into a ScopedUsageId
// note: this method should only be used for API response data and not user input
func ParseScopedUsageIDInsensitively(input string) (*ScopedUsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedUsageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedUsageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedUsageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.UsageName, ok = input.Parsed["usageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "usageName", input)
	}

	return nil
}

// ValidateScopedUsageID checks that 'input' can be parsed as a Scoped Usage ID
func ValidateScopedUsageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedUsageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Usage ID
func (id ScopedUsageId) ID() string {
	fmtString := "/%s/providers/Microsoft.Quota/usages/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.UsageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Usage ID
func (id ScopedUsageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftQuota", "Microsoft.Quota", "Microsoft.Quota"),
		resourceids.StaticSegment("staticUsages", "usages", "usages"),
		resourceids.UserSpecifiedSegment("usageName", "usageValue"),
	}
}

// String returns a human-readable description of this Scoped Usage ID
func (id ScopedUsageId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Usage Name: %q", id.UsageName),
	}
	return fmt.Sprintf("Scoped Usage (%s)", strings.Join(components, "\n"))
}
