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
	recaser.RegisterResourceId(&StaticSitesOperationStatusId{})
}

var _ resourceids.ResourceId = &StaticSitesOperationStatusId{}

// StaticSitesOperationStatusId is a struct representing the Resource ID for a Static Sites Operation Status
type StaticSitesOperationStatusId struct {
	SubscriptionId string
	LocationName   string
	OperationId    string
}

// NewStaticSitesOperationStatusID returns a new StaticSitesOperationStatusId struct
func NewStaticSitesOperationStatusID(subscriptionId string, locationName string, operationId string) StaticSitesOperationStatusId {
	return StaticSitesOperationStatusId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		OperationId:    operationId,
	}
}

// ParseStaticSitesOperationStatusID parses 'input' into a StaticSitesOperationStatusId
func ParseStaticSitesOperationStatusID(input string) (*StaticSitesOperationStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StaticSitesOperationStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StaticSitesOperationStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseStaticSitesOperationStatusIDInsensitively parses 'input' case-insensitively into a StaticSitesOperationStatusId
// note: this method should only be used for API response data and not user input
func ParseStaticSitesOperationStatusIDInsensitively(input string) (*StaticSitesOperationStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StaticSitesOperationStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StaticSitesOperationStatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *StaticSitesOperationStatusId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateStaticSitesOperationStatusID checks that 'input' can be parsed as a Static Sites Operation Status ID
func ValidateStaticSitesOperationStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStaticSitesOperationStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Static Sites Operation Status ID
func (id StaticSitesOperationStatusId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Web/locations/%s/staticSitesOperationStatuses/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Static Sites Operation Status ID
func (id StaticSitesOperationStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticStaticSitesOperationStatuses", "staticSitesOperationStatuses", "staticSitesOperationStatuses"),
		resourceids.UserSpecifiedSegment("operationId", "operationId"),
	}
}

// String returns a human-readable description of this Static Sites Operation Status ID
func (id StaticSitesOperationStatusId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Static Sites Operation Status (%s)", strings.Join(components, "\n"))
}
