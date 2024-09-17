package sqlpoolssensitivitylabels

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SensitivityLabelSourceId{})
}

var _ resourceids.ResourceId = &SensitivityLabelSourceId{}

// SensitivityLabelSourceId is a struct representing the Resource ID for a Sensitivity Label Source
type SensitivityLabelSourceId struct {
	SubscriptionId         string
	ResourceGroupName      string
	WorkspaceName          string
	SqlPoolName            string
	SchemaName             string
	TableName              string
	ColumnName             string
	SensitivityLabelSource SensitivityLabelSource
}

// NewSensitivityLabelSourceID returns a new SensitivityLabelSourceId struct
func NewSensitivityLabelSourceID(subscriptionId string, resourceGroupName string, workspaceName string, sqlPoolName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource) SensitivityLabelSourceId {
	return SensitivityLabelSourceId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		WorkspaceName:          workspaceName,
		SqlPoolName:            sqlPoolName,
		SchemaName:             schemaName,
		TableName:              tableName,
		ColumnName:             columnName,
		SensitivityLabelSource: sensitivityLabelSource,
	}
}

// ParseSensitivityLabelSourceID parses 'input' into a SensitivityLabelSourceId
func ParseSensitivityLabelSourceID(input string) (*SensitivityLabelSourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SensitivityLabelSourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SensitivityLabelSourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSensitivityLabelSourceIDInsensitively parses 'input' case-insensitively into a SensitivityLabelSourceId
// note: this method should only be used for API response data and not user input
func ParseSensitivityLabelSourceIDInsensitively(input string) (*SensitivityLabelSourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SensitivityLabelSourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SensitivityLabelSourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SensitivityLabelSourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.SqlPoolName, ok = input.Parsed["sqlPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sqlPoolName", input)
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

	if v, ok := input.Parsed["sensitivityLabelSource"]; true {
		if !ok {
			return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelSource", input)
		}

		sensitivityLabelSource, err := parseSensitivityLabelSource(v)
		if err != nil {
			return fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.SensitivityLabelSource = *sensitivityLabelSource
	}

	return nil
}

// ValidateSensitivityLabelSourceID checks that 'input' can be parsed as a Sensitivity Label Source ID
func ValidateSensitivityLabelSourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSensitivityLabelSourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sensitivity Label Source ID
func (id SensitivityLabelSourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Synapse/workspaces/%s/sqlPools/%s/schemas/%s/tables/%s/columns/%s/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.SqlPoolName, id.SchemaName, id.TableName, id.ColumnName, string(id.SensitivityLabelSource))
}

// Segments returns a slice of Resource ID Segments which comprise this Sensitivity Label Source ID
func (id SensitivityLabelSourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSynapse", "Microsoft.Synapse", "Microsoft.Synapse"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticSqlPools", "sqlPools", "sqlPools"),
		resourceids.UserSpecifiedSegment("sqlPoolName", "sqlPoolValue"),
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

// String returns a human-readable description of this Sensitivity Label Source ID
func (id SensitivityLabelSourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Sql Pool Name: %q", id.SqlPoolName),
		fmt.Sprintf("Schema Name: %q", id.SchemaName),
		fmt.Sprintf("Table Name: %q", id.TableName),
		fmt.Sprintf("Column Name: %q", id.ColumnName),
		fmt.Sprintf("Sensitivity Label Source: %q", string(id.SensitivityLabelSource)),
	}
	return fmt.Sprintf("Sensitivity Label Source (%s)", strings.Join(components, "\n"))
}
