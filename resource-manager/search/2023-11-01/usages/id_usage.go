package usages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UsageId{}

// UsageId is a struct representing the Resource ID for a Usage
type UsageId struct {
	SubscriptionId string
	LocationName   string
	UsageName      string
}

// NewUsageID returns a new UsageId struct
func NewUsageID(subscriptionId string, locationName string, usageName string) UsageId {
	return UsageId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		UsageName:      usageName,
	}
}

// ParseUsageID parses 'input' into a UsageId
func ParseUsageID(input string) (*UsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UsageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UsageId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.UsageName, ok = parsed.Parsed["usageName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usageName", *parsed)
	}

	return &id, nil
}

// ParseUsageIDInsensitively parses 'input' case-insensitively into a UsageId
// note: this method should only be used for API response data and not user input
func ParseUsageIDInsensitively(input string) (*UsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UsageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UsageId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.UsageName, ok = parsed.Parsed["usageName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "usageName", *parsed)
	}

	return &id, nil
}

// ValidateUsageID checks that 'input' can be parsed as a Usage ID
func ValidateUsageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUsageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Usage ID
func (id UsageId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Search/locations/%s/usages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.UsageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Usage ID
func (id UsageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSearch", "Microsoft.Search", "Microsoft.Search"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticUsages", "usages", "usages"),
		resourceids.UserSpecifiedSegment("usageName", "usageValue"),
	}
}

// String returns a human-readable description of this Usage ID
func (id UsageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Usage Name: %q", id.UsageName),
	}
	return fmt.Sprintf("Usage (%s)", strings.Join(components, "\n"))
}
