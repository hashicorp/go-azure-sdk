package manageddatabasetables

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SchemaTableId{}

// SchemaTableId is a struct representing the Resource ID for a Schema Table
type SchemaTableId struct {
	SubscriptionId      string
	ResourceGroupName   string
	ManagedInstanceName string
	DatabaseName        string
	SchemaName          string
	TableName           string
}

// NewSchemaTableID returns a new SchemaTableId struct
func NewSchemaTableID(subscriptionId string, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string) SchemaTableId {
	return SchemaTableId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		ManagedInstanceName: managedInstanceName,
		DatabaseName:        databaseName,
		SchemaName:          schemaName,
		TableName:           tableName,
	}
}

// ParseSchemaTableID parses 'input' into a SchemaTableId
func ParseSchemaTableID(input string) (*SchemaTableId, error) {
	parser := resourceids.NewParserFromResourceIdType(SchemaTableId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SchemaTableId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
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

	return &id, nil
}

// ParseSchemaTableIDInsensitively parses 'input' case-insensitively into a SchemaTableId
// note: this method should only be used for API response data and not user input
func ParseSchemaTableIDInsensitively(input string) (*SchemaTableId, error) {
	parser := resourceids.NewParserFromResourceIdType(SchemaTableId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SchemaTableId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
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

	return &id, nil
}

// ValidateSchemaTableID checks that 'input' can be parsed as a Schema Table ID
func ValidateSchemaTableID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSchemaTableID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Schema Table ID
func (id SchemaTableId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/databases/%s/schemas/%s/tables/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.DatabaseName, id.SchemaName, id.TableName)
}

// Segments returns a slice of Resource ID Segments which comprise this Schema Table ID
func (id SchemaTableId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticDatabases", "databases", "databases"),
		resourceids.UserSpecifiedSegment("databaseName", "databaseValue"),
		resourceids.StaticSegment("staticSchemas", "schemas", "schemas"),
		resourceids.UserSpecifiedSegment("schemaName", "schemaValue"),
		resourceids.StaticSegment("staticTables", "tables", "tables"),
		resourceids.UserSpecifiedSegment("tableName", "tableValue"),
	}
}

// String returns a human-readable description of this Schema Table ID
func (id SchemaTableId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Schema Name: %q", id.SchemaName),
		fmt.Sprintf("Table Name: %q", id.TableName),
	}
	return fmt.Sprintf("Schema Table (%s)", strings.Join(components, "\n"))
}
