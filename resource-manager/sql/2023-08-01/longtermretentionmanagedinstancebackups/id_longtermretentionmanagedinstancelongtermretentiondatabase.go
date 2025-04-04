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
	recaser.RegisterResourceId(&LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{})
}

var _ resourceids.ResourceId = &LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{}

// LongTermRetentionManagedInstanceLongTermRetentionDatabaseId is a struct representing the Resource ID for a Long Term Retention Managed Instance Long Term Retention Database
type LongTermRetentionManagedInstanceLongTermRetentionDatabaseId struct {
	SubscriptionId                       string
	ResourceGroupName                    string
	LocationName                         string
	LongTermRetentionManagedInstanceName string
	LongTermRetentionDatabaseName        string
}

// NewLongTermRetentionManagedInstanceLongTermRetentionDatabaseID returns a new LongTermRetentionManagedInstanceLongTermRetentionDatabaseId struct
func NewLongTermRetentionManagedInstanceLongTermRetentionDatabaseID(subscriptionId string, resourceGroupName string, locationName string, longTermRetentionManagedInstanceName string, longTermRetentionDatabaseName string) LongTermRetentionManagedInstanceLongTermRetentionDatabaseId {
	return LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{
		SubscriptionId:                       subscriptionId,
		ResourceGroupName:                    resourceGroupName,
		LocationName:                         locationName,
		LongTermRetentionManagedInstanceName: longTermRetentionManagedInstanceName,
		LongTermRetentionDatabaseName:        longTermRetentionDatabaseName,
	}
}

// ParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseID parses 'input' into a LongTermRetentionManagedInstanceLongTermRetentionDatabaseId
func ParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseID(input string) (*LongTermRetentionManagedInstanceLongTermRetentionDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseIDInsensitively parses 'input' case-insensitively into a LongTermRetentionManagedInstanceLongTermRetentionDatabaseId
// note: this method should only be used for API response data and not user input
func ParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseIDInsensitively(input string) (*LongTermRetentionManagedInstanceLongTermRetentionDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionManagedInstanceLongTermRetentionDatabaseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LongTermRetentionManagedInstanceLongTermRetentionDatabaseId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
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

// ValidateLongTermRetentionManagedInstanceLongTermRetentionDatabaseID checks that 'input' can be parsed as a Long Term Retention Managed Instance Long Term Retention Database ID
func ValidateLongTermRetentionManagedInstanceLongTermRetentionDatabaseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLongTermRetentionManagedInstanceLongTermRetentionDatabaseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Long Term Retention Managed Instance Long Term Retention Database ID
func (id LongTermRetentionManagedInstanceLongTermRetentionDatabaseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/locations/%s/longTermRetentionManagedInstances/%s/longTermRetentionDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.LongTermRetentionManagedInstanceName, id.LongTermRetentionDatabaseName)
}

// Segments returns a slice of Resource ID Segments which comprise this Long Term Retention Managed Instance Long Term Retention Database ID
func (id LongTermRetentionManagedInstanceLongTermRetentionDatabaseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticLongTermRetentionManagedInstances", "longTermRetentionManagedInstances", "longTermRetentionManagedInstances"),
		resourceids.UserSpecifiedSegment("longTermRetentionManagedInstanceName", "longTermRetentionManagedInstanceName"),
		resourceids.StaticSegment("staticLongTermRetentionDatabases", "longTermRetentionDatabases", "longTermRetentionDatabases"),
		resourceids.UserSpecifiedSegment("longTermRetentionDatabaseName", "longTermRetentionDatabaseName"),
	}
}

// String returns a human-readable description of this Long Term Retention Managed Instance Long Term Retention Database ID
func (id LongTermRetentionManagedInstanceLongTermRetentionDatabaseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Long Term Retention Managed Instance Name: %q", id.LongTermRetentionManagedInstanceName),
		fmt.Sprintf("Long Term Retention Database Name: %q", id.LongTermRetentionDatabaseName),
	}
	return fmt.Sprintf("Long Term Retention Managed Instance Long Term Retention Database (%s)", strings.Join(components, "\n"))
}
