package timezones

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TimeZoneId{})
}

var _ resourceids.ResourceId = &TimeZoneId{}

// TimeZoneId is a struct representing the Resource ID for a Time Zone
type TimeZoneId struct {
	SubscriptionId string
	LocationName   string
	TimeZoneId     string
}

// NewTimeZoneID returns a new TimeZoneId struct
func NewTimeZoneID(subscriptionId string, locationName string, timeZoneId string) TimeZoneId {
	return TimeZoneId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		TimeZoneId:     timeZoneId,
	}
}

// ParseTimeZoneID parses 'input' into a TimeZoneId
func ParseTimeZoneID(input string) (*TimeZoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TimeZoneId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TimeZoneId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTimeZoneIDInsensitively parses 'input' case-insensitively into a TimeZoneId
// note: this method should only be used for API response data and not user input
func ParseTimeZoneIDInsensitively(input string) (*TimeZoneId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TimeZoneId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TimeZoneId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TimeZoneId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.TimeZoneId, ok = input.Parsed["timeZoneId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "timeZoneId", input)
	}

	return nil
}

// ValidateTimeZoneID checks that 'input' can be parsed as a Time Zone ID
func ValidateTimeZoneID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTimeZoneID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Time Zone ID
func (id TimeZoneId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Sql/locations/%s/timeZones/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.TimeZoneId)
}

// Segments returns a slice of Resource ID Segments which comprise this Time Zone ID
func (id TimeZoneId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticTimeZones", "timeZones", "timeZones"),
		resourceids.UserSpecifiedSegment("timeZoneId", "timeZoneIdValue"),
	}
}

// String returns a human-readable description of this Time Zone ID
func (id TimeZoneId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Time Zone: %q", id.TimeZoneId),
	}
	return fmt.Sprintf("Time Zone (%s)", strings.Join(components, "\n"))
}
