package directoryrecommendationimpactedresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryRecommendationImpactedResourceId{}

// DirectoryRecommendationImpactedResourceId is a struct representing the Resource ID for a Directory Recommendation Impacted Resource
type DirectoryRecommendationImpactedResourceId struct {
	RecommendationId   string
	ImpactedResourceId string
}

// NewDirectoryRecommendationImpactedResourceID returns a new DirectoryRecommendationImpactedResourceId struct
func NewDirectoryRecommendationImpactedResourceID(recommendationId string, impactedResourceId string) DirectoryRecommendationImpactedResourceId {
	return DirectoryRecommendationImpactedResourceId{
		RecommendationId:   recommendationId,
		ImpactedResourceId: impactedResourceId,
	}
}

// ParseDirectoryRecommendationImpactedResourceID parses 'input' into a DirectoryRecommendationImpactedResourceId
func ParseDirectoryRecommendationImpactedResourceID(input string) (*DirectoryRecommendationImpactedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryRecommendationImpactedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryRecommendationImpactedResourceId{}

	if id.RecommendationId, ok = parsed.Parsed["recommendationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationId", *parsed)
	}

	if id.ImpactedResourceId, ok = parsed.Parsed["impactedResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "impactedResourceId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryRecommendationImpactedResourceIDInsensitively parses 'input' case-insensitively into a DirectoryRecommendationImpactedResourceId
// note: this method should only be used for API response data and not user input
func ParseDirectoryRecommendationImpactedResourceIDInsensitively(input string) (*DirectoryRecommendationImpactedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryRecommendationImpactedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryRecommendationImpactedResourceId{}

	if id.RecommendationId, ok = parsed.Parsed["recommendationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationId", *parsed)
	}

	if id.ImpactedResourceId, ok = parsed.Parsed["impactedResourceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "impactedResourceId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryRecommendationImpactedResourceID checks that 'input' can be parsed as a Directory Recommendation Impacted Resource ID
func ValidateDirectoryRecommendationImpactedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryRecommendationImpactedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Recommendation Impacted Resource ID
func (id DirectoryRecommendationImpactedResourceId) ID() string {
	fmtString := "/directory/recommendations/%s/impactedResources/%s"
	return fmt.Sprintf(fmtString, id.RecommendationId, id.ImpactedResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Recommendation Impacted Resource ID
func (id DirectoryRecommendationImpactedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticRecommendations", "recommendations", "recommendations"),
		resourceids.UserSpecifiedSegment("recommendationId", "recommendationIdValue"),
		resourceids.StaticSegment("staticImpactedResources", "impactedResources", "impactedResources"),
		resourceids.UserSpecifiedSegment("impactedResourceId", "impactedResourceIdValue"),
	}
}

// String returns a human-readable description of this Directory Recommendation Impacted Resource ID
func (id DirectoryRecommendationImpactedResourceId) String() string {
	components := []string{
		fmt.Sprintf("Recommendation: %q", id.RecommendationId),
		fmt.Sprintf("Impacted Resource: %q", id.ImpactedResourceId),
	}
	return fmt.Sprintf("Directory Recommendation Impacted Resource (%s)", strings.Join(components, "\n"))
}
