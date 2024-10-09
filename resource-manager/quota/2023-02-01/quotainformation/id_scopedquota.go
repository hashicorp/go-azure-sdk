package quotainformation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedQuotaId{})
}

var _ resourceids.ResourceId = &ScopedQuotaId{}

// ScopedQuotaId is a struct representing the Resource ID for a Scoped Quota
type ScopedQuotaId struct {
	Scope     string
	QuotaName string
}

// NewScopedQuotaID returns a new ScopedQuotaId struct
func NewScopedQuotaID(scope string, quotaName string) ScopedQuotaId {
	return ScopedQuotaId{
		Scope:     scope,
		QuotaName: quotaName,
	}
}

// ParseScopedQuotaID parses 'input' into a ScopedQuotaId
func ParseScopedQuotaID(input string) (*ScopedQuotaId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedQuotaId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedQuotaId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedQuotaIDInsensitively parses 'input' case-insensitively into a ScopedQuotaId
// note: this method should only be used for API response data and not user input
func ParseScopedQuotaIDInsensitively(input string) (*ScopedQuotaId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedQuotaId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedQuotaId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedQuotaId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.QuotaName, ok = input.Parsed["quotaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "quotaName", input)
	}

	return nil
}

// ValidateScopedQuotaID checks that 'input' can be parsed as a Scoped Quota ID
func ValidateScopedQuotaID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedQuotaID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Quota ID
func (id ScopedQuotaId) ID() string {
	fmtString := "/%s/providers/Microsoft.Quota/quotas/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.QuotaName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Quota ID
func (id ScopedQuotaId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftQuota", "Microsoft.Quota", "Microsoft.Quota"),
		resourceids.StaticSegment("staticQuotas", "quotas", "quotas"),
		resourceids.UserSpecifiedSegment("quotaName", "quotaName"),
	}
}

// String returns a human-readable description of this Scoped Quota ID
func (id ScopedQuotaId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Quota Name: %q", id.QuotaName),
	}
	return fmt.Sprintf("Scoped Quota (%s)", strings.Join(components, "\n"))
}
