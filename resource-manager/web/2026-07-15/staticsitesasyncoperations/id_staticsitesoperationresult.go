package staticsitesasyncoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&StaticSitesOperationResultId{})
}

var _ resourceids.ResourceId = &StaticSitesOperationResultId{}

// StaticSitesOperationResultId is a struct representing the Resource ID for a Static Sites Operation Result
type StaticSitesOperationResultId struct {
	SubscriptionId string
	LocationName   string
	OperationId    string
}

// NewStaticSitesOperationResultID returns a new StaticSitesOperationResultId struct
func NewStaticSitesOperationResultID(subscriptionId string, locationName string, operationId string) StaticSitesOperationResultId {
	return StaticSitesOperationResultId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		OperationId:    operationId,
	}
}

// ParseStaticSitesOperationResultID parses 'input' into a StaticSitesOperationResultId
func ParseStaticSitesOperationResultID(input string) (*StaticSitesOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StaticSitesOperationResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StaticSitesOperationResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseStaticSitesOperationResultIDInsensitively parses 'input' case-insensitively into a StaticSitesOperationResultId
// note: this method should only be used for API response data and not user input
func ParseStaticSitesOperationResultIDInsensitively(input string) (*StaticSitesOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StaticSitesOperationResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StaticSitesOperationResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *StaticSitesOperationResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
}

// ValidateStaticSitesOperationResultID checks that 'input' can be parsed as a Static Sites Operation Result ID
func ValidateStaticSitesOperationResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStaticSitesOperationResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Static Sites Operation Result ID
func (id StaticSitesOperationResultId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Web/locations/%s/staticSitesOperationResults/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Static Sites Operation Result ID
func (id StaticSitesOperationResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticStaticSitesOperationResults", "staticSitesOperationResults", "staticSitesOperationResults"),
		resourceids.UserSpecifiedSegment("operationId", "operationId"),
	}
}

// String returns a human-readable description of this Static Sites Operation Result ID
func (id StaticSitesOperationResultId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Static Sites Operation Result (%s)", strings.Join(components, "\n"))
}
