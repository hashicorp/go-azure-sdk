package longtermretentionmanagedinstancebackups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId{}

// LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId is a struct representing the Resource ID for a Long Term Retention Database Long Term Retention Managed Instance Backup
type LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId struct {
	SubscriptionId                             string
	ResourceGroupName                          string
	LocationName                               string
	LongTermRetentionManagedInstanceName       string
	LongTermRetentionDatabaseName              string
	LongTermRetentionManagedInstanceBackupName string
}

// NewLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID returns a new LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId struct
func NewLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID(subscriptionId string, resourceGroupName string, locationName string, longTermRetentionManagedInstanceName string, longTermRetentionDatabaseName string, longTermRetentionManagedInstanceBackupName string) LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId {
	return LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId{
		SubscriptionId:                             subscriptionId,
		ResourceGroupName:                          resourceGroupName,
		LocationName:                               locationName,
		LongTermRetentionManagedInstanceName:       longTermRetentionManagedInstanceName,
		LongTermRetentionDatabaseName:              longTermRetentionDatabaseName,
		LongTermRetentionManagedInstanceBackupName: longTermRetentionManagedInstanceBackupName,
	}
}

// ParseLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID parses 'input' into a LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId
func ParseLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID(input string) (*LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.LongTermRetentionManagedInstanceName, ok = parsed.Parsed["longTermRetentionManagedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionManagedInstanceName", *parsed)
	}

	if id.LongTermRetentionDatabaseName, ok = parsed.Parsed["longTermRetentionDatabaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionDatabaseName", *parsed)
	}

	if id.LongTermRetentionManagedInstanceBackupName, ok = parsed.Parsed["longTermRetentionManagedInstanceBackupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionManagedInstanceBackupName", *parsed)
	}

	return &id, nil
}

// ParseLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupIDInsensitively parses 'input' case-insensitively into a LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId
// note: this method should only be used for API response data and not user input
func ParseLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupIDInsensitively(input string) (*LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.LongTermRetentionManagedInstanceName, ok = parsed.Parsed["longTermRetentionManagedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionManagedInstanceName", *parsed)
	}

	if id.LongTermRetentionDatabaseName, ok = parsed.Parsed["longTermRetentionDatabaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionDatabaseName", *parsed)
	}

	if id.LongTermRetentionManagedInstanceBackupName, ok = parsed.Parsed["longTermRetentionManagedInstanceBackupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionManagedInstanceBackupName", *parsed)
	}

	return &id, nil
}

// ValidateLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID checks that 'input' can be parsed as a Long Term Retention Database Long Term Retention Managed Instance Backup ID
func ValidateLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Long Term Retention Database Long Term Retention Managed Instance Backup ID
func (id LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/locations/%s/longTermRetentionManagedInstances/%s/longTermRetentionDatabases/%s/longTermRetentionManagedInstanceBackups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.LongTermRetentionManagedInstanceName, id.LongTermRetentionDatabaseName, id.LongTermRetentionManagedInstanceBackupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Long Term Retention Database Long Term Retention Managed Instance Backup ID
func (id LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticLongTermRetentionManagedInstances", "longTermRetentionManagedInstances", "longTermRetentionManagedInstances"),
		resourceids.UserSpecifiedSegment("longTermRetentionManagedInstanceName", "longTermRetentionManagedInstanceValue"),
		resourceids.StaticSegment("staticLongTermRetentionDatabases", "longTermRetentionDatabases", "longTermRetentionDatabases"),
		resourceids.UserSpecifiedSegment("longTermRetentionDatabaseName", "longTermRetentionDatabaseValue"),
		resourceids.StaticSegment("staticLongTermRetentionManagedInstanceBackups", "longTermRetentionManagedInstanceBackups", "longTermRetentionManagedInstanceBackups"),
		resourceids.UserSpecifiedSegment("longTermRetentionManagedInstanceBackupName", "longTermRetentionManagedInstanceBackupValue"),
	}
}

// String returns a human-readable description of this Long Term Retention Database Long Term Retention Managed Instance Backup ID
func (id LongTermRetentionDatabaseLongTermRetentionManagedInstanceBackupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Long Term Retention Managed Instance Name: %q", id.LongTermRetentionManagedInstanceName),
		fmt.Sprintf("Long Term Retention Database Name: %q", id.LongTermRetentionDatabaseName),
		fmt.Sprintf("Long Term Retention Managed Instance Backup Name: %q", id.LongTermRetentionManagedInstanceBackupName),
	}
	return fmt.Sprintf("Long Term Retention Database Long Term Retention Managed Instance Backup (%s)", strings.Join(components, "\n"))
}
