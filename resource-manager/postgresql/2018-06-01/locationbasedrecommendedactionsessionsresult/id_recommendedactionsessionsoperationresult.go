package locationbasedrecommendedactionsessionsresult

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RecommendedActionSessionsOperationResultId{}

// RecommendedActionSessionsOperationResultId is a struct representing the Resource ID for a Recommended Action Sessions Operation Result
type RecommendedActionSessionsOperationResultId struct {
	SubscriptionId string
	LocationName   string
	OperationId    string
}

// NewRecommendedActionSessionsOperationResultID returns a new RecommendedActionSessionsOperationResultId struct
func NewRecommendedActionSessionsOperationResultID(subscriptionId string, locationName string, operationId string) RecommendedActionSessionsOperationResultId {
	return RecommendedActionSessionsOperationResultId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		OperationId:    operationId,
	}
}

// ParseRecommendedActionSessionsOperationResultID parses 'input' into a RecommendedActionSessionsOperationResultId
func ParseRecommendedActionSessionsOperationResultID(input string) (*RecommendedActionSessionsOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecommendedActionSessionsOperationResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecommendedActionSessionsOperationResultId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, fmt.Errorf("the segment 'operationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRecommendedActionSessionsOperationResultIDInsensitively parses 'input' case-insensitively into a RecommendedActionSessionsOperationResultId
// note: this method should only be used for API response data and not user input
func ParseRecommendedActionSessionsOperationResultIDInsensitively(input string) (*RecommendedActionSessionsOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecommendedActionSessionsOperationResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecommendedActionSessionsOperationResultId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, fmt.Errorf("the segment 'locationName' was not found in the resource id %q", input)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, fmt.Errorf("the segment 'operationId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRecommendedActionSessionsOperationResultID checks that 'input' can be parsed as a Recommended Action Sessions Operation Result ID
func ValidateRecommendedActionSessionsOperationResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRecommendedActionSessionsOperationResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Recommended Action Sessions Operation Result ID
func (id RecommendedActionSessionsOperationResultId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.DBforPostgreSQL/locations/%s/recommendedActionSessionsOperationResults/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Recommended Action Sessions Operation Result ID
func (id RecommendedActionSessionsOperationResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDBforPostgreSQL", "Microsoft.DBforPostgreSQL", "Microsoft.DBforPostgreSQL"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticRecommendedActionSessionsOperationResults", "recommendedActionSessionsOperationResults", "recommendedActionSessionsOperationResults"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Recommended Action Sessions Operation Result ID
func (id RecommendedActionSessionsOperationResultId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Recommended Action Sessions Operation Result (%s)", strings.Join(components, "\n"))
}
