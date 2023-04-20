package generaterecommendations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GenerateRecommendationId{}

// GenerateRecommendationId is a struct representing the Resource ID for a Generate Recommendation
type GenerateRecommendationId struct {
	SubscriptionId string
	OperationId    string
}

// NewGenerateRecommendationID returns a new GenerateRecommendationId struct
func NewGenerateRecommendationID(subscriptionId string, operationId string) GenerateRecommendationId {
	return GenerateRecommendationId{
		SubscriptionId: subscriptionId,
		OperationId:    operationId,
	}
}

// ParseGenerateRecommendationID parses 'input' into a GenerateRecommendationId
func ParseGenerateRecommendationID(input string) (*GenerateRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GenerateRecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GenerateRecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, fmt.Errorf("the segment 'operationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseGenerateRecommendationIDInsensitively parses 'input' case-insensitively into a GenerateRecommendationId
// note: this method should only be used for API response data and not user input
func ParseGenerateRecommendationIDInsensitively(input string) (*GenerateRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(GenerateRecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GenerateRecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, fmt.Errorf("the segment 'operationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateGenerateRecommendationID checks that 'input' can be parsed as a Generate Recommendation ID
func ValidateGenerateRecommendationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGenerateRecommendationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Generate Recommendation ID
func (id GenerateRecommendationId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Advisor/generateRecommendations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Generate Recommendation ID
func (id GenerateRecommendationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAdvisor", "Microsoft.Advisor", "Microsoft.Advisor"),
		resourceids.StaticSegment("staticGenerateRecommendations", "generateRecommendations", "generateRecommendations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Generate Recommendation ID
func (id GenerateRecommendationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Generate Recommendation (%s)", strings.Join(components, "\n"))
}
