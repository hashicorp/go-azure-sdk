package longtermretentionmanagedinstancebackups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LongTermRetentionManagedInstanceId{})
}

var _ resourceids.ResourceId = &LongTermRetentionManagedInstanceId{}

// LongTermRetentionManagedInstanceId is a struct representing the Resource ID for a Long Term Retention Managed Instance
type LongTermRetentionManagedInstanceId struct {
	SubscriptionId                       string
	LocationName                         string
	LongTermRetentionManagedInstanceName string
}

// NewLongTermRetentionManagedInstanceID returns a new LongTermRetentionManagedInstanceId struct
func NewLongTermRetentionManagedInstanceID(subscriptionId string, locationName string, longTermRetentionManagedInstanceName string) LongTermRetentionManagedInstanceId {
	return LongTermRetentionManagedInstanceId{
		SubscriptionId:                       subscriptionId,
		LocationName:                         locationName,
		LongTermRetentionManagedInstanceName: longTermRetentionManagedInstanceName,
	}
}

// ParseLongTermRetentionManagedInstanceID parses 'input' into a LongTermRetentionManagedInstanceId
func ParseLongTermRetentionManagedInstanceID(input string) (*LongTermRetentionManagedInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionManagedInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionManagedInstanceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLongTermRetentionManagedInstanceIDInsensitively parses 'input' case-insensitively into a LongTermRetentionManagedInstanceId
// note: this method should only be used for API response data and not user input
func ParseLongTermRetentionManagedInstanceIDInsensitively(input string) (*LongTermRetentionManagedInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionManagedInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionManagedInstanceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LongTermRetentionManagedInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.LocationName, ok = input.Parsed["locationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "locationName", input)
	}

	if id.LongTermRetentionManagedInstanceName, ok = input.Parsed["longTermRetentionManagedInstanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionManagedInstanceName", input)
	}

	return nil
}

// ValidateLongTermRetentionManagedInstanceID checks that 'input' can be parsed as a Long Term Retention Managed Instance ID
func ValidateLongTermRetentionManagedInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLongTermRetentionManagedInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Long Term Retention Managed Instance ID
func (id LongTermRetentionManagedInstanceId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Sql/locations/%s/longTermRetentionManagedInstances/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.LongTermRetentionManagedInstanceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Long Term Retention Managed Instance ID
func (id LongTermRetentionManagedInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticLongTermRetentionManagedInstances", "longTermRetentionManagedInstances", "longTermRetentionManagedInstances"),
		resourceids.UserSpecifiedSegment("longTermRetentionManagedInstanceName", "longTermRetentionManagedInstanceValue"),
	}
}

// String returns a human-readable description of this Long Term Retention Managed Instance ID
func (id LongTermRetentionManagedInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Long Term Retention Managed Instance Name: %q", id.LongTermRetentionManagedInstanceName),
	}
	return fmt.Sprintf("Long Term Retention Managed Instance (%s)", strings.Join(components, "\n"))
}
