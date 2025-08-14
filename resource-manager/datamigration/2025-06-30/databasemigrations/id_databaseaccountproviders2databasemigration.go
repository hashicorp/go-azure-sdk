package databasemigrations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DatabaseAccountProviders2DatabaseMigrationId{})
}

var _ resourceids.ResourceId = &DatabaseAccountProviders2DatabaseMigrationId{}

// DatabaseAccountProviders2DatabaseMigrationId is a struct representing the Resource ID for a Database Account Providers 2 Database Migration
type DatabaseAccountProviders2DatabaseMigrationId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DatabaseAccountName   string
	DatabaseMigrationName string
}

// NewDatabaseAccountProviders2DatabaseMigrationID returns a new DatabaseAccountProviders2DatabaseMigrationId struct
func NewDatabaseAccountProviders2DatabaseMigrationID(subscriptionId string, resourceGroupName string, databaseAccountName string, databaseMigrationName string) DatabaseAccountProviders2DatabaseMigrationId {
	return DatabaseAccountProviders2DatabaseMigrationId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DatabaseAccountName:   databaseAccountName,
		DatabaseMigrationName: databaseMigrationName,
	}
}

// ParseDatabaseAccountProviders2DatabaseMigrationID parses 'input' into a DatabaseAccountProviders2DatabaseMigrationId
func ParseDatabaseAccountProviders2DatabaseMigrationID(input string) (*DatabaseAccountProviders2DatabaseMigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DatabaseAccountProviders2DatabaseMigrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DatabaseAccountProviders2DatabaseMigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDatabaseAccountProviders2DatabaseMigrationIDInsensitively parses 'input' case-insensitively into a DatabaseAccountProviders2DatabaseMigrationId
// note: this method should only be used for API response data and not user input
func ParseDatabaseAccountProviders2DatabaseMigrationIDInsensitively(input string) (*DatabaseAccountProviders2DatabaseMigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DatabaseAccountProviders2DatabaseMigrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DatabaseAccountProviders2DatabaseMigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DatabaseAccountProviders2DatabaseMigrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DatabaseAccountName, ok = input.Parsed["databaseAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseAccountName", input)
	}

	if id.DatabaseMigrationName, ok = input.Parsed["databaseMigrationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseMigrationName", input)
	}

	return nil
}

// ValidateDatabaseAccountProviders2DatabaseMigrationID checks that 'input' can be parsed as a Database Account Providers 2 Database Migration ID
func ValidateDatabaseAccountProviders2DatabaseMigrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDatabaseAccountProviders2DatabaseMigrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Database Account Providers 2 Database Migration ID
func (id DatabaseAccountProviders2DatabaseMigrationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/providers/Microsoft.DataMigration/databaseMigrations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DatabaseAccountName, id.DatabaseMigrationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Database Account Providers 2 Database Migration ID
func (id DatabaseAccountProviders2DatabaseMigrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDocumentDB", "Microsoft.DocumentDB", "Microsoft.DocumentDB"),
		resourceids.StaticSegment("staticDatabaseAccounts", "databaseAccounts", "databaseAccounts"),
		resourceids.UserSpecifiedSegment("databaseAccountName", "databaseAccountName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataMigration", "Microsoft.DataMigration", "Microsoft.DataMigration"),
		resourceids.StaticSegment("staticDatabaseMigrations", "databaseMigrations", "databaseMigrations"),
		resourceids.UserSpecifiedSegment("databaseMigrationName", "databaseMigrationName"),
	}
}

// String returns a human-readable description of this Database Account Providers 2 Database Migration ID
func (id DatabaseAccountProviders2DatabaseMigrationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Database Account Name: %q", id.DatabaseAccountName),
		fmt.Sprintf("Database Migration Name: %q", id.DatabaseMigrationName),
	}
	return fmt.Sprintf("Database Account Providers 2 Database Migration (%s)", strings.Join(components, "\n"))
}
