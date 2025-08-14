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
	recaser.RegisterResourceId(&MongoClusterProviders2DatabaseMigrationId{})
}

var _ resourceids.ResourceId = &MongoClusterProviders2DatabaseMigrationId{}

// MongoClusterProviders2DatabaseMigrationId is a struct representing the Resource ID for a Mongo Cluster Providers 2 Database Migration
type MongoClusterProviders2DatabaseMigrationId struct {
	SubscriptionId        string
	ResourceGroupName     string
	MongoClusterName      string
	DatabaseMigrationName string
}

// NewMongoClusterProviders2DatabaseMigrationID returns a new MongoClusterProviders2DatabaseMigrationId struct
func NewMongoClusterProviders2DatabaseMigrationID(subscriptionId string, resourceGroupName string, mongoClusterName string, databaseMigrationName string) MongoClusterProviders2DatabaseMigrationId {
	return MongoClusterProviders2DatabaseMigrationId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		MongoClusterName:      mongoClusterName,
		DatabaseMigrationName: databaseMigrationName,
	}
}

// ParseMongoClusterProviders2DatabaseMigrationID parses 'input' into a MongoClusterProviders2DatabaseMigrationId
func ParseMongoClusterProviders2DatabaseMigrationID(input string) (*MongoClusterProviders2DatabaseMigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MongoClusterProviders2DatabaseMigrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MongoClusterProviders2DatabaseMigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMongoClusterProviders2DatabaseMigrationIDInsensitively parses 'input' case-insensitively into a MongoClusterProviders2DatabaseMigrationId
// note: this method should only be used for API response data and not user input
func ParseMongoClusterProviders2DatabaseMigrationIDInsensitively(input string) (*MongoClusterProviders2DatabaseMigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MongoClusterProviders2DatabaseMigrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MongoClusterProviders2DatabaseMigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MongoClusterProviders2DatabaseMigrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.MongoClusterName, ok = input.Parsed["mongoClusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "mongoClusterName", input)
	}

	if id.DatabaseMigrationName, ok = input.Parsed["databaseMigrationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseMigrationName", input)
	}

	return nil
}

// ValidateMongoClusterProviders2DatabaseMigrationID checks that 'input' can be parsed as a Mongo Cluster Providers 2 Database Migration ID
func ValidateMongoClusterProviders2DatabaseMigrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMongoClusterProviders2DatabaseMigrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Mongo Cluster Providers 2 Database Migration ID
func (id MongoClusterProviders2DatabaseMigrationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/mongoClusters/%s/providers/Microsoft.DataMigration/databaseMigrations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MongoClusterName, id.DatabaseMigrationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Mongo Cluster Providers 2 Database Migration ID
func (id MongoClusterProviders2DatabaseMigrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDocumentDB", "Microsoft.DocumentDB", "Microsoft.DocumentDB"),
		resourceids.StaticSegment("staticMongoClusters", "mongoClusters", "mongoClusters"),
		resourceids.UserSpecifiedSegment("mongoClusterName", "mongoClusterName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataMigration", "Microsoft.DataMigration", "Microsoft.DataMigration"),
		resourceids.StaticSegment("staticDatabaseMigrations", "databaseMigrations", "databaseMigrations"),
		resourceids.UserSpecifiedSegment("databaseMigrationName", "databaseMigrationName"),
	}
}

// String returns a human-readable description of this Mongo Cluster Providers 2 Database Migration ID
func (id MongoClusterProviders2DatabaseMigrationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Mongo Cluster Name: %q", id.MongoClusterName),
		fmt.Sprintf("Database Migration Name: %q", id.DatabaseMigrationName),
	}
	return fmt.Sprintf("Mongo Cluster Providers 2 Database Migration (%s)", strings.Join(components, "\n"))
}
