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
	recaser.RegisterResourceId(&LongTermRetentionManagedInstanceBackupId{})
}

var _ resourceids.ResourceId = &LongTermRetentionManagedInstanceBackupId{}

// LongTermRetentionManagedInstanceBackupId is a struct representing the Resource ID for a Long Term Retention Managed Instance Backup
type LongTermRetentionManagedInstanceBackupId struct {
	SubscriptionId                             string
	LocationName                               string
	LongTermRetentionManagedInstanceName       string
	LongTermRetentionDatabaseName              string
	LongTermRetentionManagedInstanceBackupName string
}

// NewLongTermRetentionManagedInstanceBackupID returns a new LongTermRetentionManagedInstanceBackupId struct
func NewLongTermRetentionManagedInstanceBackupID(subscriptionId string, locationName string, longTermRetentionManagedInstanceName string, longTermRetentionDatabaseName string, longTermRetentionManagedInstanceBackupName string) LongTermRetentionManagedInstanceBackupId {
	return LongTermRetentionManagedInstanceBackupId{
		SubscriptionId:                             subscriptionId,
		LocationName:                               locationName,
		LongTermRetentionManagedInstanceName:       longTermRetentionManagedInstanceName,
		LongTermRetentionDatabaseName:              longTermRetentionDatabaseName,
		LongTermRetentionManagedInstanceBackupName: longTermRetentionManagedInstanceBackupName,
	}
}

// ParseLongTermRetentionManagedInstanceBackupID parses 'input' into a LongTermRetentionManagedInstanceBackupId
func ParseLongTermRetentionManagedInstanceBackupID(input string) (*LongTermRetentionManagedInstanceBackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionManagedInstanceBackupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionManagedInstanceBackupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseLongTermRetentionManagedInstanceBackupIDInsensitively parses 'input' case-insensitively into a LongTermRetentionManagedInstanceBackupId
// note: this method should only be used for API response data and not user input
func ParseLongTermRetentionManagedInstanceBackupIDInsensitively(input string) (*LongTermRetentionManagedInstanceBackupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&LongTermRetentionManagedInstanceBackupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := LongTermRetentionManagedInstanceBackupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *LongTermRetentionManagedInstanceBackupId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.LongTermRetentionManagedInstanceBackupName, ok = input.Parsed["longTermRetentionManagedInstanceBackupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "longTermRetentionManagedInstanceBackupName", input)
	}

	return nil
}

// ValidateLongTermRetentionManagedInstanceBackupID checks that 'input' can be parsed as a Long Term Retention Managed Instance Backup ID
func ValidateLongTermRetentionManagedInstanceBackupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseLongTermRetentionManagedInstanceBackupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Long Term Retention Managed Instance Backup ID
func (id LongTermRetentionManagedInstanceBackupId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Sql/locations/%s/longTermRetentionManagedInstances/%s/longTermRetentionDatabases/%s/longTermRetentionManagedInstanceBackups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.LongTermRetentionManagedInstanceName, id.LongTermRetentionDatabaseName, id.LongTermRetentionManagedInstanceBackupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Long Term Retention Managed Instance Backup ID
func (id LongTermRetentionManagedInstanceBackupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationName"),
		resourceids.StaticSegment("staticLongTermRetentionManagedInstances", "longTermRetentionManagedInstances", "longTermRetentionManagedInstances"),
		resourceids.UserSpecifiedSegment("longTermRetentionManagedInstanceName", "longTermRetentionManagedInstanceName"),
		resourceids.StaticSegment("staticLongTermRetentionDatabases", "longTermRetentionDatabases", "longTermRetentionDatabases"),
		resourceids.UserSpecifiedSegment("longTermRetentionDatabaseName", "longTermRetentionDatabaseName"),
		resourceids.StaticSegment("staticLongTermRetentionManagedInstanceBackups", "longTermRetentionManagedInstanceBackups", "longTermRetentionManagedInstanceBackups"),
		resourceids.UserSpecifiedSegment("longTermRetentionManagedInstanceBackupName", "longTermRetentionManagedInstanceBackupName"),
	}
}

// String returns a human-readable description of this Long Term Retention Managed Instance Backup ID
func (id LongTermRetentionManagedInstanceBackupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Long Term Retention Managed Instance Name: %q", id.LongTermRetentionManagedInstanceName),
		fmt.Sprintf("Long Term Retention Database Name: %q", id.LongTermRetentionDatabaseName),
		fmt.Sprintf("Long Term Retention Managed Instance Backup Name: %q", id.LongTermRetentionManagedInstanceBackupName),
	}
	return fmt.Sprintf("Long Term Retention Managed Instance Backup (%s)", strings.Join(components, "\n"))
}
