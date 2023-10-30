package asyncoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = OperationsStatusId{}

// OperationsStatusId is a struct representing the Resource ID for a Operations Status
type OperationsStatusId struct {
	SubscriptionId   string
	LocationName     string
	AsyncOperationId string
}

// NewOperationsStatusID returns a new OperationsStatusId struct
func NewOperationsStatusID(subscriptionId string, locationName string, asyncOperationId string) OperationsStatusId {
	return OperationsStatusId{
		SubscriptionId:   subscriptionId,
		LocationName:     locationName,
		AsyncOperationId: asyncOperationId,
	}
}

// ParseOperationsStatusID parses 'input' into a OperationsStatusId
func ParseOperationsStatusID(input string) (*OperationsStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationsStatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationsStatusId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.AsyncOperationId, ok = parsed.Parsed["asyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "asyncOperationId", *parsed)
	}

	return &id, nil
}

// ParseOperationsStatusIDInsensitively parses 'input' case-insensitively into a OperationsStatusId
// note: this method should only be used for API response data and not user input
func ParseOperationsStatusIDInsensitively(input string) (*OperationsStatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(OperationsStatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := OperationsStatusId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.AsyncOperationId, ok = parsed.Parsed["asyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "asyncOperationId", *parsed)
	}

	return &id, nil
}

// ValidateOperationsStatusID checks that 'input' can be parsed as a Operations Status ID
func ValidateOperationsStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOperationsStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Operations Status ID
func (id OperationsStatusId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Chaos/locations/%s/operationsStatuses/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.AsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Operations Status ID
func (id OperationsStatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticOperationsStatuses", "operationsStatuses", "operationsStatuses"),
		resourceids.UserSpecifiedSegment("asyncOperationId", "asyncOperationIdValue"),
	}
}

// String returns a human-readable description of this Operations Status ID
func (id OperationsStatusId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Async Operation: %q", id.AsyncOperationId),
	}
	return fmt.Sprintf("Operations Status (%s)", strings.Join(components, "\n"))
}
