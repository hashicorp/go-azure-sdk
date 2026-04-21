package manageddatabasesensitivitylabels

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TableColumnId{})
}

var _ resourceids.ResourceId = &TableColumnId{}

// TableColumnId is a struct representing the Resource ID for a Table Column
type TableColumnId struct {
	SubscriptionId      string
	ResourceGroupName   string
	ManagedInstanceName string
	DatabaseName        string
	SchemaName          string
	TableName           string
	ColumnName          string
}

// NewTableColumnID returns a new TableColumnId struct
func NewTableColumnID(subscriptionId string, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string) TableColumnId {
	return TableColumnId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		ManagedInstanceName: managedInstanceName,
		DatabaseName:        databaseName,
		SchemaName:          schemaName,
		TableName:           tableName,
		ColumnName:          columnName,
	}
}

// ParseTableColumnID parses 'input' into a TableColumnId
func ParseTableColumnID(input string) (*TableColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TableColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TableColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTableColumnIDInsensitively parses 'input' case-insensitively into a TableColumnId
// note: this method should only be used for API response data and not user input
func ParseTableColumnIDInsensitively(input string) (*TableColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TableColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TableColumnId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TableColumnId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedInstanceName, ok = input.Parsed["managedInstanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", input)
	}

	if id.DatabaseName, ok = input.Parsed["databaseName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "databaseName", input)
	}

	if id.SchemaName, ok = input.Parsed["schemaName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "schemaName", input)
	}

	if id.TableName, ok = input.Parsed["tableName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "tableName", input)
	}

	if id.ColumnName, ok = input.Parsed["columnName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "columnName", input)
	}

	return nil
}

// ValidateTableColumnID checks that 'input' can be parsed as a Table Column ID
func ValidateTableColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTableColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Table Column ID
func (id TableColumnId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/databases/%s/schemas/%s/tables/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.DatabaseName, id.SchemaName, id.TableName, id.ColumnName)
}

// Segments returns a slice of Resource ID Segments which comprise this Table Column ID
func (id TableColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceName"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseName"),
		resourceids.StaticSegment("staticSchemas", "schemas", "schemas"),
		resourceids.UserSpecifiedSegment("schemaName", "schemaName"),
		resourceids.StaticSegment("staticTables", "tables", "tables"),
		resourceids.UserSpecifiedSegment("tableName", "tableName"),
		resourceids.StaticSegment("staticColumns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnName", "columnName"),
	}
}

// String returns a human-readable description of this Table Column ID
func (id TableColumnId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Schema Name: %q", id.SchemaName),
		fmt.Sprintf("Table Name: %q", id.TableName),
		fmt.Sprintf("Column Name: %q", id.ColumnName),
	}
	return fmt.Sprintf("Table Column (%s)", strings.Join(components, "\n"))
}
