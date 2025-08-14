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
	recaser.RegisterResourceId(&DatabaseMigrationId{})
}

var _ resourceids.ResourceId = &DatabaseMigrationId{}

// DatabaseMigrationId is a struct representing the Resource ID for a Database Migration
type DatabaseMigrationId struct {
	SubscriptionId        string
	ResourceGroupName     string
	ServerName            string
	DatabaseMigrationName string
}

// NewDatabaseMigrationID returns a new DatabaseMigrationId struct
func NewDatabaseMigrationID(subscriptionId string, resourceGroupName string, serverName string, databaseMigrationName string) DatabaseMigrationId {
	return DatabaseMigrationId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		ServerName:            serverName,
		DatabaseMigrationName: databaseMigrationName,
	}
}

// ParseDatabaseMigrationID parses 'input' into a DatabaseMigrationId
func ParseDatabaseMigrationID(input string) (*DatabaseMigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DatabaseMigrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DatabaseMigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDatabaseMigrationIDInsensitively parses 'input' case-insensitively into a DatabaseMigrationId
// note: this method should only be used for API response data and not user input
func ParseDatabaseMigrationIDInsensitively(input string) (*DatabaseMigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DatabaseMigrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DatabaseMigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DatabaseMigrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ServerName, ok = input.Parsed["serverName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serverName", input)
	}

	if id.DatabaseMigrationName, ok = input.Parsed["databaseMigrationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseMigrationName", input)
	}

	return nil
}

// ValidateDatabaseMigrationID checks that 'input' can be parsed as a Database Migration ID
func ValidateDatabaseMigrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDatabaseMigrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Database Migration ID
func (id DatabaseMigrationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/providers/Microsoft.DataMigration/databaseMigrations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseMigrationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Database Migration ID
func (id DatabaseMigrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataMigration", "Microsoft.DataMigration", "Microsoft.DataMigration"),
		resourceids.StaticSegment("staticDatabaseMigrations", "databaseMigrations", "databaseMigrations"),
		resourceids.UserSpecifiedSegment("databaseMigrationName", "databaseMigrationName"),
	}
}

// String returns a human-readable description of this Database Migration ID
func (id DatabaseMigrationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Migration Name: %q", id.DatabaseMigrationName),
	}
	return fmt.Sprintf("Database Migration (%s)", strings.Join(components, "\n"))
}
