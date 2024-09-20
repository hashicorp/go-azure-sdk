package longtermretentionbackups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&LongTermRetentionDatabaseLongTermRetentionBackupId{})
}

var _ resourceids.ResourceId = &LongTermRetentionDatabaseLongTermRetentionBackupId{}

// LongTermRetentionDatabaseLongTermRetentionBackupId is a struct representing the Resource ID for a Long Term Retention Database Long Term Retention Backup
type LongTermRetentionDatabaseLongTermRetentionBackupId struct {
	SubscriptionId                string
	ResourceGroupName             string
	LocationName                  string
	LongTermRetentionServerName   string
	LongTermRetentionDatabaseName string
	LongTermRetentionBackupName   string
}

// NewLongTermRetentionDatabaseLongTermRetentionBackupID returns a new LongTermRetentionDatabaseLongTermRetentionBackupId struct
func NewLongTermRetentionDatabaseLongTermRetentionBackupID(subscriptionId string, resourceGroupName string, locationName string, longTermRetentionServerName string, longTermRetentionDatabaseName string, longTermRetentionBackupName string) LongTermRetentionDatabaseLongTermRetentionBackupId {
	return LongTermRetentionDatabaseLongTermRetentionBackupId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		LocationName:                  locationName,
		LongTermRetentionServerName:   longTermRetentionServerName,
		LongTermRetentionDatabaseName: longTermRetentionDatabaseName,
		LongTermRetentionBackupName:   longTermRetentionBackupName,
	}
}

// ParseLongTermRetentionDatabaseLongTermRetentionBackupID parses 'input' into a LongTermRetentionDatabaseLongTermRetentionBackupId
func ParseLongTermRetentionDatabaseLongTermRetentionBackupID(input string) (*LongTermRetentionDatabaseLongTermRetentionBackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionDatabaseLongTermRetentionBackupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionDatabaseLongTermRetentionBackupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLongTermRetentionDatabaseLongTermRetentionBackupIDInsensitively parses 'input' case-insensitively into a LongTermRetentionDatabaseLongTermRetentionBackupId
// note: this method should only be used for API response data and not user input
func ParseLongTermRetentionDatabaseLongTermRetentionBackupIDInsensitively(input string) (*LongTermRetentionDatabaseLongTermRetentionBackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionDatabaseLongTermRetentionBackupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionDatabaseLongTermRetentionBackupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LongTermRetentionDatabaseLongTermRetentionBackupId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LongTermRetentionServerName, ok = input.Parsed["longTermRetentionServerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionServerName", input)
	}

	if id.LongTermRetentionDatabaseName, ok = input.Parsed["longTermRetentionDatabaseName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionDatabaseName", input)
	}

	if id.LongTermRetentionBackupName, ok = input.Parsed["longTermRetentionBackupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionBackupName", input)
	}

	return nil
}

// ValidateLongTermRetentionDatabaseLongTermRetentionBackupID checks that 'input' can be parsed as a Long Term Retention Database Long Term Retention Backup ID
func ValidateLongTermRetentionDatabaseLongTermRetentionBackupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLongTermRetentionDatabaseLongTermRetentionBackupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Long Term Retention Database Long Term Retention Backup ID
func (id LongTermRetentionDatabaseLongTermRetentionBackupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/locations/%s/longTermRetentionServers/%s/longTermRetentionDatabases/%s/longTermRetentionBackups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.LongTermRetentionServerName, id.LongTermRetentionDatabaseName, id.LongTermRetentionBackupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Long Term Retention Database Long Term Retention Backup ID
func (id LongTermRetentionDatabaseLongTermRetentionBackupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticLongTermRetentionServers", "longTermRetentionServers", "longTermRetentionServers"),
		resourceids.UserSpecifiedSegment("longTermRetentionServerName", "longTermRetentionServerName"),
		resourceids.StaticSegment("staticLongTermRetentionDatabases", "longTermRetentionDatabases", "longTermRetentionDatabases"),
		resourceids.UserSpecifiedSegment("longTermRetentionDatabaseName", "longTermRetentionDatabaseName"),
		resourceids.StaticSegment("staticLongTermRetentionBackups", "longTermRetentionBackups", "longTermRetentionBackups"),
		resourceids.UserSpecifiedSegment("longTermRetentionBackupName", "backupName"),
	}
}

// String returns a human-readable description of this Long Term Retention Database Long Term Retention Backup ID
func (id LongTermRetentionDatabaseLongTermRetentionBackupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Long Term Retention Server Name: %q", id.LongTermRetentionServerName),
		fmt.Sprintf("Long Term Retention Database Name: %q", id.LongTermRetentionDatabaseName),
		fmt.Sprintf("Long Term Retention Backup Name: %q", id.LongTermRetentionBackupName),
	}
	return fmt.Sprintf("Long Term Retention Database Long Term Retention Backup (%s)", strings.Join(components, "\n"))
}
