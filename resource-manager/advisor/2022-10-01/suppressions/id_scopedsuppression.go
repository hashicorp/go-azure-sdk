// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package suppressions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScopedSuppressionId{}

// ScopedSuppressionId is a struct representing the Resource ID for a Scoped Suppression
type ScopedSuppressionId struct {
	ResourceUri      string
	RecommendationId string
	SuppressionName  string
}

// NewScopedSuppressionID returns a new ScopedSuppressionId struct
func NewScopedSuppressionID(resourceUri string, recommendationId string, suppressionName string) ScopedSuppressionId {
	return ScopedSuppressionId{
		ResourceUri:      resourceUri,
		RecommendationId: recommendationId,
		SuppressionName:  suppressionName,
	}
}

// ParseScopedSuppressionID parses 'input' into a ScopedSuppressionId
func ParseScopedSuppressionID(input string) (*ScopedSuppressionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedSuppressionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedSuppressionId{}

	if id.ResourceUri, ok = parsed.Parsed["resourceUri"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceUri' was not found in the resource id %q", input)
	}

	if id.RecommendationId, ok = parsed.Parsed["recommendationId"]; !ok {
		return nil, fmt.Errorf("the segment 'recommendationId' was not found in the resource id %q", input)
	}

	if id.SuppressionName, ok = parsed.Parsed["suppressionName"]; !ok {
		return nil, fmt.Errorf("the segment 'suppressionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseScopedSuppressionIDInsensitively parses 'input' case-insensitively into a ScopedSuppressionId
// note: this method should only be used for API response data and not user input
func ParseScopedSuppressionIDInsensitively(input string) (*ScopedSuppressionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedSuppressionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedSuppressionId{}

	if id.ResourceUri, ok = parsed.Parsed["resourceUri"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceUri' was not found in the resource id %q", input)
	}

	if id.RecommendationId, ok = parsed.Parsed["recommendationId"]; !ok {
		return nil, fmt.Errorf("the segment 'recommendationId' was not found in the resource id %q", input)
	}

	if id.SuppressionName, ok = parsed.Parsed["suppressionName"]; !ok {
		return nil, fmt.Errorf("the segment 'suppressionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateScopedSuppressionID checks that 'input' can be parsed as a Scoped Suppression ID
func ValidateScopedSuppressionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedSuppressionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Suppression ID
func (id ScopedSuppressionId) ID() string {
	fmtString := "/%s/providers/Microsoft.Advisor/recommendations/%s/suppressions/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceUri, "/"), id.RecommendationId, id.SuppressionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Suppression ID
func (id ScopedSuppressionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceUri", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAdvisor", "Microsoft.Advisor", "Microsoft.Advisor"),
		resourceids.StaticSegment("staticRecommendations", "recommendations", "recommendations"),
		resourceids.UserSpecifiedSegment("recommendationId", "recommendationIdValue"),
		resourceids.StaticSegment("staticSuppressions", "suppressions", "suppressions"),
		resourceids.UserSpecifiedSegment("suppressionName", "suppressionValue"),
	}
}

// String returns a human-readable description of this Scoped Suppression ID
func (id ScopedSuppressionId) String() string {
	components := []string{
		fmt.Sprintf("Resource Uri: %q", id.ResourceUri),
		fmt.Sprintf("Recommendation: %q", id.RecommendationId),
		fmt.Sprintf("Suppression Name: %q", id.SuppressionName),
	}
	return fmt.Sprintf("Scoped Suppression (%s)", strings.Join(components, "\n"))
}
