package recommendations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RecommendationId{}

// RecommendationId is a struct representing the Resource ID for a Recommendation
type RecommendationId struct {
	SubscriptionId     string
	RecommendationName string
}

// NewRecommendationID returns a new RecommendationId struct
func NewRecommendationID(subscriptionId string, recommendationName string) RecommendationId {
	return RecommendationId{
		SubscriptionId:     subscriptionId,
		RecommendationName: recommendationName,
	}
}

// ParseRecommendationID parses 'input' into a RecommendationId
func ParseRecommendationID(input string) (*RecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.RecommendationName, ok = parsed.Parsed["recommendationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationName", *parsed)
	}

	return &id, nil
}

// ParseRecommendationIDInsensitively parses 'input' case-insensitively into a RecommendationId
// note: this method should only be used for API response data and not user input
func ParseRecommendationIDInsensitively(input string) (*RecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.RecommendationName, ok = parsed.Parsed["recommendationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recommendationName", *parsed)
	}

	return &id, nil
}

// ValidateRecommendationID checks that 'input' can be parsed as a Recommendation ID
func ValidateRecommendationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRecommendationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Recommendation ID
func (id RecommendationId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Web/recommendations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.RecommendationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Recommendation ID
func (id RecommendationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticRecommendations", "recommendations", "recommendations"),
		resourceids.UserSpecifiedSegment("recommendationName", "recommendationValue"),
	}
}

// String returns a human-readable description of this Recommendation ID
func (id RecommendationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Recommendation Name: %q", id.RecommendationName),
	}
	return fmt.Sprintf("Recommendation (%s)", strings.Join(components, "\n"))
}
