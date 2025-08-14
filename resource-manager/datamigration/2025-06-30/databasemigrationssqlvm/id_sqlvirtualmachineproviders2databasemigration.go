package databasemigrationssqlvm

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SqlVirtualMachineProviders2DatabaseMigrationId{})
}

var _ resourceids.ResourceId = &SqlVirtualMachineProviders2DatabaseMigrationId{}

// SqlVirtualMachineProviders2DatabaseMigrationId is a struct representing the Resource ID for a Sql Virtual Machine Providers 2 Database Migration
type SqlVirtualMachineProviders2DatabaseMigrationId struct {
	SubscriptionId        string
	ResourceGroupName     string
	SqlVirtualMachineName string
	DatabaseMigrationName string
}

// NewSqlVirtualMachineProviders2DatabaseMigrationID returns a new SqlVirtualMachineProviders2DatabaseMigrationId struct
func NewSqlVirtualMachineProviders2DatabaseMigrationID(subscriptionId string, resourceGroupName string, sqlVirtualMachineName string, databaseMigrationName string) SqlVirtualMachineProviders2DatabaseMigrationId {
	return SqlVirtualMachineProviders2DatabaseMigrationId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		SqlVirtualMachineName: sqlVirtualMachineName,
		DatabaseMigrationName: databaseMigrationName,
	}
}

// ParseSqlVirtualMachineProviders2DatabaseMigrationID parses 'input' into a SqlVirtualMachineProviders2DatabaseMigrationId
func ParseSqlVirtualMachineProviders2DatabaseMigrationID(input string) (*SqlVirtualMachineProviders2DatabaseMigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SqlVirtualMachineProviders2DatabaseMigrationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SqlVirtualMachineProviders2DatabaseMigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSqlVirtualMachineProviders2DatabaseMigrationIDInsensitively parses 'input' case-insensitively into a SqlVirtualMachineProviders2DatabaseMigrationId
// note: this method should only be used for API response data and not user input
func ParseSqlVirtualMachineProviders2DatabaseMigrationIDInsensitively(input string) (*SqlVirtualMachineProviders2DatabaseMigrationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SqlVirtualMachineProviders2DatabaseMigrationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SqlVirtualMachineProviders2DatabaseMigrationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SqlVirtualMachineProviders2DatabaseMigrationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SqlVirtualMachineName, ok = input.Parsed["sqlVirtualMachineName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sqlVirtualMachineName", input)
	}

	if id.DatabaseMigrationName, ok = input.Parsed["databaseMigrationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseMigrationName", input)
	}

	return nil
}

// ValidateSqlVirtualMachineProviders2DatabaseMigrationID checks that 'input' can be parsed as a Sql Virtual Machine Providers 2 Database Migration ID
func ValidateSqlVirtualMachineProviders2DatabaseMigrationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSqlVirtualMachineProviders2DatabaseMigrationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sql Virtual Machine Providers 2 Database Migration ID
func (id SqlVirtualMachineProviders2DatabaseMigrationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/%s/providers/Microsoft.DataMigration/databaseMigrations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SqlVirtualMachineName, id.DatabaseMigrationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sql Virtual Machine Providers 2 Database Migration ID
func (id SqlVirtualMachineProviders2DatabaseMigrationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSqlVirtualMachine", "Microsoft.SqlVirtualMachine", "Microsoft.SqlVirtualMachine"),
		resourceids.StaticSegment("staticSqlVirtualMachines", "sqlVirtualMachines", "sqlVirtualMachines"),
		resourceids.UserSpecifiedSegment("sqlVirtualMachineName", "sqlVirtualMachineName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataMigration", "Microsoft.DataMigration", "Microsoft.DataMigration"),
		resourceids.StaticSegment("staticDatabaseMigrations", "databaseMigrations", "databaseMigrations"),
		resourceids.UserSpecifiedSegment("databaseMigrationName", "databaseMigrationName"),
	}
}

// String returns a human-readable description of this Sql Virtual Machine Providers 2 Database Migration ID
func (id SqlVirtualMachineProviders2DatabaseMigrationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Sql Virtual Machine Name: %q", id.SqlVirtualMachineName),
		fmt.Sprintf("Database Migration Name: %q", id.DatabaseMigrationName),
	}
	return fmt.Sprintf("Sql Virtual Machine Providers 2 Database Migration (%s)", strings.Join(components, "\n"))
}
