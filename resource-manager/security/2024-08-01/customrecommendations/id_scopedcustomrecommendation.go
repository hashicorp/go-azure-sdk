package customrecommendations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedCustomRecommendationId{})
}

var _ resourceids.ResourceId = &ScopedCustomRecommendationId{}

// ScopedCustomRecommendationId is a struct representing the Resource ID for a Scoped Custom Recommendation
type ScopedCustomRecommendationId struct {
	Scope                    string
	CustomRecommendationName string
}

// NewScopedCustomRecommendationID returns a new ScopedCustomRecommendationId struct
func NewScopedCustomRecommendationID(scope string, customRecommendationName string) ScopedCustomRecommendationId {
	return ScopedCustomRecommendationId{
		Scope:                    scope,
		CustomRecommendationName: customRecommendationName,
	}
}

// ParseScopedCustomRecommendationID parses 'input' into a ScopedCustomRecommendationId
func ParseScopedCustomRecommendationID(input string) (*ScopedCustomRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedCustomRecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedCustomRecommendationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedCustomRecommendationIDInsensitively parses 'input' case-insensitively into a ScopedCustomRecommendationId
// note: this method should only be used for API response data and not user input
func ParseScopedCustomRecommendationIDInsensitively(input string) (*ScopedCustomRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedCustomRecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedCustomRecommendationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedCustomRecommendationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.CustomRecommendationName, ok = input.Parsed["customRecommendationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customRecommendationName", input)
	}

	return nil
}

// ValidateScopedCustomRecommendationID checks that 'input' can be parsed as a Scoped Custom Recommendation ID
func ValidateScopedCustomRecommendationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedCustomRecommendationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Custom Recommendation ID
func (id ScopedCustomRecommendationId) ID() string {
	fmtString := "/%s/providers/Microsoft.Security/customRecommendations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.CustomRecommendationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Custom Recommendation ID
func (id ScopedCustomRecommendationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticCustomRecommendations", "customRecommendations", "customRecommendations"),
		resourceids.UserSpecifiedSegment("customRecommendationName", "customRecommendationName"),
	}
}

// String returns a human-readable description of this Scoped Custom Recommendation ID
func (id ScopedCustomRecommendationId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Custom Recommendation Name: %q", id.CustomRecommendationName),
	}
	return fmt.Sprintf("Scoped Custom Recommendation (%s)", strings.Join(components, "\n"))
}
