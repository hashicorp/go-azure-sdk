package manageddatabasesensitivitylabels

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SensitivityLabelSensitivityLabelSourceId{}

// SensitivityLabelSensitivityLabelSourceId is a struct representing the Resource ID for a Sensitivity Label Sensitivity Label Source
type SensitivityLabelSensitivityLabelSourceId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ManagedInstanceName    string
	DatabaseName           string
	SchemaName             string
	TableName              string
	ColumnName             string
	SensitivityLabelSource SensitivityLabelSource
}

// NewSensitivityLabelSensitivityLabelSourceID returns a new SensitivityLabelSensitivityLabelSourceId struct
func NewSensitivityLabelSensitivityLabelSourceID(subscriptionId string, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource) SensitivityLabelSensitivityLabelSourceId {
	return SensitivityLabelSensitivityLabelSourceId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ManagedInstanceName:    managedInstanceName,
		DatabaseName:           databaseName,
		SchemaName:             schemaName,
		TableName:              tableName,
		ColumnName:             columnName,
		SensitivityLabelSource: sensitivityLabelSource,
	}
}

// ParseSensitivityLabelSensitivityLabelSourceID parses 'input' into a SensitivityLabelSensitivityLabelSourceId
func ParseSensitivityLabelSensitivityLabelSourceID(input string) (*SensitivityLabelSensitivityLabelSourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(SensitivityLabelSensitivityLabelSourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SensitivityLabelSensitivityLabelSourceId{}

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

	if id.ColumnName, ok = parsed.Parsed["columnName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnName", *parsed)
	}

	if v, ok := parsed.Parsed["sensitivityLabelSource"]; true {
		if !ok {
			return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelSource", *parsed)
		}

		sensitivityLabelSource, err := parseSensitivityLabelSource(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.SensitivityLabelSource = *sensitivityLabelSource
	}

	return &id, nil
}

// ParseSensitivityLabelSensitivityLabelSourceIDInsensitively parses 'input' case-insensitively into a SensitivityLabelSensitivityLabelSourceId
// note: this method should only be used for API response data and not user input
func ParseSensitivityLabelSensitivityLabelSourceIDInsensitively(input string) (*SensitivityLabelSensitivityLabelSourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(SensitivityLabelSensitivityLabelSourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SensitivityLabelSensitivityLabelSourceId{}

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

	if id.ColumnName, ok = parsed.Parsed["columnName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "columnName", *parsed)
	}

	if v, ok := parsed.Parsed["sensitivityLabelSource"]; true {
		if !ok {
			return nil, resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelSource", *parsed)
		}

		sensitivityLabelSource, err := parseSensitivityLabelSource(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.SensitivityLabelSource = *sensitivityLabelSource
	}

	return &id, nil
}

// ValidateSensitivityLabelSensitivityLabelSourceID checks that 'input' can be parsed as a Sensitivity Label Sensitivity Label Source ID
func ValidateSensitivityLabelSensitivityLabelSourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSensitivityLabelSensitivityLabelSourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sensitivity Label Sensitivity Label Source ID
func (id SensitivityLabelSensitivityLabelSourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/databases/%s/schemas/%s/tables/%s/columns/%s/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.DatabaseName, id.SchemaName, id.TableName, id.ColumnName, string(id.SensitivityLabelSource))
}

// Segments returns a slice of Resource ID Segments which comprise this Sensitivity Label Sensitivity Label Source ID
func (id SensitivityLabelSensitivityLabelSourceId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticColumns", "columns", "columns"),
		resourceids.UserSpecifiedSegment("columnName", "columnValue"),
		resourceids.StaticSegment("staticSensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.ConstantSegment("sensitivityLabelSource", PossibleValuesForSensitivityLabelSource(), "current"),
	}
}

// String returns a human-readable description of this Sensitivity Label Sensitivity Label Source ID
func (id SensitivityLabelSensitivityLabelSourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Schema Name: %q", id.SchemaName),
		fmt.Sprintf("Table Name: %q", id.TableName),
		fmt.Sprintf("Column Name: %q", id.ColumnName),
		fmt.Sprintf("Sensitivity Label Source: %q", string(id.SensitivityLabelSource)),
	}
	return fmt.Sprintf("Sensitivity Label Sensitivity Label Source (%s)", strings.Join(components, "\n"))
}
