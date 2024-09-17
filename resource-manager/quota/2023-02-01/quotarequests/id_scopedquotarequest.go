package quotarequests

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedQuotaRequestId{})
}

var _ resourceids.ResourceId = &ScopedQuotaRequestId{}

// ScopedQuotaRequestId is a struct representing the Resource ID for a Scoped Quota Request
type ScopedQuotaRequestId struct {
	Scope            string
	QuotaRequestName string
}

// NewScopedQuotaRequestID returns a new ScopedQuotaRequestId struct
func NewScopedQuotaRequestID(scope string, quotaRequestName string) ScopedQuotaRequestId {
	return ScopedQuotaRequestId{
		Scope:            scope,
		QuotaRequestName: quotaRequestName,
	}
}

// ParseScopedQuotaRequestID parses 'input' into a ScopedQuotaRequestId
func ParseScopedQuotaRequestID(input string) (*ScopedQuotaRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedQuotaRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedQuotaRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedQuotaRequestIDInsensitively parses 'input' case-insensitively into a ScopedQuotaRequestId
// note: this method should only be used for API response data and not user input
func ParseScopedQuotaRequestIDInsensitively(input string) (*ScopedQuotaRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedQuotaRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedQuotaRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedQuotaRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.QuotaRequestName, ok = input.Parsed["quotaRequestName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "quotaRequestName", input)
	}

	return nil
}

// ValidateScopedQuotaRequestID checks that 'input' can be parsed as a Scoped Quota Request ID
func ValidateScopedQuotaRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedQuotaRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Quota Request ID
func (id ScopedQuotaRequestId) ID() string {
	fmtString := "/%s/providers/Microsoft.Quota/quotaRequests/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.QuotaRequestName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Quota Request ID
func (id ScopedQuotaRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftQuota", "Microsoft.Quota", "Microsoft.Quota"),
		resourceids.StaticSegment("staticQuotaRequests", "quotaRequests", "quotaRequests"),
		resourceids.UserSpecifiedSegment("quotaRequestName", "quotaRequestValue"),
	}
}

// String returns a human-readable description of this Scoped Quota Request ID
func (id ScopedQuotaRequestId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Quota Request Name: %q", id.QuotaRequestName),
	}
	return fmt.Sprintf("Scoped Quota Request (%s)", strings.Join(components, "\n"))
}
