package subscriptionusages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&UsageId{})
}

var _ resourceids.ResourceId = &UsageId{}

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
	parser := resourceids.NewParserFromResourceIdType(&UsageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UsageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUsageIDInsensitively parses 'input' case-insensitively into a UsageId
// note: this method should only be used for API response data and not user input
func ParseUsageIDInsensitively(input string) (*UsageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UsageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UsageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UsageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.UsageName, ok = input.Parsed["usageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "usageName", input)
	}

	return nil
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
	fmtString := "/subscriptions/%s/providers/Microsoft.Sql/locations/%s/usages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.UsageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Usage ID
func (id UsageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
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
