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
	recaser.RegisterResourceId(&LongTermRetentionDatabaseId{})
}

var _ resourceids.ResourceId = &LongTermRetentionDatabaseId{}

// LongTermRetentionDatabaseId is a struct representing the Resource ID for a Long Term Retention Database
type LongTermRetentionDatabaseId struct {
	SubscriptionId                       string
	LocationName                         string
	LongTermRetentionManagedInstanceName string
	LongTermRetentionDatabaseName        string
}

// NewLongTermRetentionDatabaseID returns a new LongTermRetentionDatabaseId struct
func NewLongTermRetentionDatabaseID(subscriptionId string, locationName string, longTermRetentionManagedInstanceName string, longTermRetentionDatabaseName string) LongTermRetentionDatabaseId {
	return LongTermRetentionDatabaseId{
		SubscriptionId:                       subscriptionId,
		LocationName:                         locationName,
		LongTermRetentionManagedInstanceName: longTermRetentionManagedInstanceName,
		LongTermRetentionDatabaseName:        longTermRetentionDatabaseName,
	}
}

// ParseLongTermRetentionDatabaseID parses 'input' into a LongTermRetentionDatabaseId
func ParseLongTermRetentionDatabaseID(input string) (*LongTermRetentionDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionDatabaseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionDatabaseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLongTermRetentionDatabaseIDInsensitively parses 'input' case-insensitively into a LongTermRetentionDatabaseId
// note: this method should only be used for API response data and not user input
func ParseLongTermRetentionDatabaseIDInsensitively(input string) (*LongTermRetentionDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionDatabaseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionDatabaseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LongTermRetentionDatabaseId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LongTermRetentionDatabaseName, ok = input.Parsed["longTermRetentionDatabaseName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionDatabaseName", input)
	}

	return nil
}

// ValidateLongTermRetentionDatabaseID checks that 'input' can be parsed as a Long Term Retention Database ID
func ValidateLongTermRetentionDatabaseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLongTermRetentionDatabaseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Long Term Retention Database ID
func (id LongTermRetentionDatabaseId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Sql/locations/%s/longTermRetentionManagedInstances/%s/longTermRetentionDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.LongTermRetentionManagedInstanceName, id.LongTermRetentionDatabaseName)
}

// Segments returns a slice of Resource ID Segments which comprise this Long Term Retention Database ID
func (id LongTermRetentionDatabaseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticLongTermRetentionManagedInstances", "longTermRetentionManagedInstances", "longTermRetentionManagedInstances"),
		resourceids.UserSpecifiedSegment("longTermRetentionManagedInstanceName", "managedInstanceName"),
		resourceids.StaticSegment("staticLongTermRetentionDatabases", "longTermRetentionDatabases", "longTermRetentionDatabases"),
		resourceids.UserSpecifiedSegment("longTermRetentionDatabaseName", "databaseName"),
	}
}

// String returns a human-readable description of this Long Term Retention Database ID
func (id LongTermRetentionDatabaseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Long Term Retention Managed Instance Name: %q", id.LongTermRetentionManagedInstanceName),
		fmt.Sprintf("Long Term Retention Database Name: %q", id.LongTermRetentionDatabaseName),
	}
	return fmt.Sprintf("Long Term Retention Database (%s)", strings.Join(components, "\n"))
}
