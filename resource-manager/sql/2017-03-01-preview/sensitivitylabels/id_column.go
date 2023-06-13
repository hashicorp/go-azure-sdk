package sensitivitylabels

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ColumnId{}

// ColumnId is a struct representing the Resource ID for a Column
type ColumnId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	DatabaseName      string
	SchemaName        string
	TableName         string
	ColumnName        string
}

// NewColumnID returns a new ColumnId struct
func NewColumnID(subscriptionId string, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string) ColumnId {
	return ColumnId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		DatabaseName:      databaseName,
		SchemaName:        schemaName,
		TableName:         tableName,
		ColumnName:        columnName,
	}
}

// ParseColumnID parses 'input' into a ColumnId
func ParseColumnID(input string) (*ColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(ColumnId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ColumnId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseName", *parsed)
	}

	if id.SchemaName, ok = parsed.Parsed["schemaName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "schemaName", *parsed)
	}

	if id.TableName, ok = parsed.Parsed["tableName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tableName", *parsed)
	}

	if id.ColumnName, ok = parsed.Parsed["columnName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnName", *parsed)
	}

	return &id, nil
}

// ParseColumnIDInsensitively parses 'input' case-insensitively into a ColumnId
// note: this method should only be used for API response data and not user input
func ParseColumnIDInsensitively(input string) (*ColumnId, error) {
	parser := resourceids.NewParserFromResourceIdType(ColumnId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ColumnId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServerName, ok = parsed.Parsed["serverName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverName", *parsed)
	}

	if id.DatabaseName, ok = parsed.Parsed["databaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseName", *parsed)
	}

	if id.SchemaName, ok = parsed.Parsed["schemaName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "schemaName", *parsed)
	}

	if id.TableName, ok = parsed.Parsed["tableName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tableName", *parsed)
	}

	if id.ColumnName, ok = parsed.Parsed["columnName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnName", *parsed)
	}

	return &id, nil
}

// ValidateColumnID checks that 'input' can be parsed as a Column ID
func ValidateColumnID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseColumnID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Column ID
func (id ColumnId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/databases/%s/schemas/%s/tables/%s/columns/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseName, id.SchemaName, id.TableName, id.ColumnName)
}

// Segments returns a slice of Resource ID Segments which comprise this Column ID
func (id ColumnId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticServers", "servers", "servers"),
		resourceids.UserSpecifiedSegment("serverName", "serverValue"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
		resourceids.StaticSegment("staticSchemas", "schemas", "schemas"),
		resourceids.UserSpecifiedSegment("schemaName", "schemaValue"),
		resourceids.StaticSegment("staticTables", "tables", "tables"),
		resourceids.UserSpecifiedSegment("tableName", "tableValue"),
		resourceids.StaticSegment("staticColumns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnName", "columnValue"),
	}
}

// String returns a human-readable description of this Column ID
func (id ColumnId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Schema Name: %q", id.SchemaName),
		fmt.Sprintf("Table Name: %q", id.TableName),
		fmt.Sprintf("Column Name: %q", id.ColumnName),
	}
	return fmt.Sprintf("Column (%s)", strings.Join(components, "\n"))
}
