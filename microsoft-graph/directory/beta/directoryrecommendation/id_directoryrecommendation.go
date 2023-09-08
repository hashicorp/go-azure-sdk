package directoryrecommendation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DirectoryRecommendationId{}

// DirectoryRecommendationId is a struct representing the Resource ID for a Directory Recommendation
type DirectoryRecommendationId struct {
	RecommendationId string
}

// NewDirectoryRecommendationID returns a new DirectoryRecommendationId struct
func NewDirectoryRecommendationID(recommendationId string) DirectoryRecommendationId {
	return DirectoryRecommendationId{
		RecommendationId: recommendationId,
	}
}

// ParseDirectoryRecommendationID parses 'input' into a DirectoryRecommendationId
func ParseDirectoryRecommendationID(input string) (*DirectoryRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryRecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryRecommendationId{}

	if id.RecommendationId, ok = parsed.Parsed["recommendationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationId", *parsed)
	}

	return &id, nil
}

// ParseDirectoryRecommendationIDInsensitively parses 'input' case-insensitively into a DirectoryRecommendationId
// note: this method should only be used for API response data and not user input
func ParseDirectoryRecommendationIDInsensitively(input string) (*DirectoryRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(DirectoryRecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DirectoryRecommendationId{}

	if id.RecommendationId, ok = parsed.Parsed["recommendationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationId", *parsed)
	}

	return &id, nil
}

// ValidateDirectoryRecommendationID checks that 'input' can be parsed as a Directory Recommendation ID
func ValidateDirectoryRecommendationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryRecommendationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Recommendation ID
func (id DirectoryRecommendationId) ID() string {
	fmtString := "/directory/recommendations/%s"
	return fmt.Sprintf(fmtString, id.RecommendationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Recommendation ID
func (id DirectoryRecommendationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticRecommendations", "recommendations", "recommendations"),
		resourceids.UserSpecifiedSegment("recommendationId", "recommendationIdValue"),
	}
}

// String returns a human-readable description of this Directory Recommendation ID
func (id DirectoryRecommendationId) String() string {
	components := []string{
		fmt.Sprintf("Recommendation: %q", id.RecommendationId),
	}
	return fmt.Sprintf("Directory Recommendation (%s)", strings.Join(components, "\n"))
}
