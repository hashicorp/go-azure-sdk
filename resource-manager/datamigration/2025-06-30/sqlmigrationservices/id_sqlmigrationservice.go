package sqlmigrationservices

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SqlMigrationServiceId{})
}

var _ resourceids.ResourceId = &SqlMigrationServiceId{}

// SqlMigrationServiceId is a struct representing the Resource ID for a Sql Migration Service
type SqlMigrationServiceId struct {
	SubscriptionId          string
	ResourceGroupName       string
	SqlMigrationServiceName string
}

// NewSqlMigrationServiceID returns a new SqlMigrationServiceId struct
func NewSqlMigrationServiceID(subscriptionId string, resourceGroupName string, sqlMigrationServiceName string) SqlMigrationServiceId {
	return SqlMigrationServiceId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		SqlMigrationServiceName: sqlMigrationServiceName,
	}
}

// ParseSqlMigrationServiceID parses 'input' into a SqlMigrationServiceId
func ParseSqlMigrationServiceID(input string) (*SqlMigrationServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SqlMigrationServiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SqlMigrationServiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSqlMigrationServiceIDInsensitively parses 'input' case-insensitively into a SqlMigrationServiceId
// note: this method should only be used for API response data and not user input
func ParseSqlMigrationServiceIDInsensitively(input string) (*SqlMigrationServiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SqlMigrationServiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SqlMigrationServiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SqlMigrationServiceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SqlMigrationServiceName, ok = input.Parsed["sqlMigrationServiceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sqlMigrationServiceName", input)
	}

	return nil
}

// ValidateSqlMigrationServiceID checks that 'input' can be parsed as a Sql Migration Service ID
func ValidateSqlMigrationServiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSqlMigrationServiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sql Migration Service ID
func (id SqlMigrationServiceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataMigration/sqlMigrationServices/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SqlMigrationServiceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sql Migration Service ID
func (id SqlMigrationServiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataMigration", "Microsoft.DataMigration", "Microsoft.DataMigration"),
		resourceids.StaticSegment("staticSqlMigrationServices", "sqlMigrationServices", "sqlMigrationServices"),
		resourceids.UserSpecifiedSegment("sqlMigrationServiceName", "sqlMigrationServiceName"),
	}
}

// String returns a human-readable description of this Sql Migration Service ID
func (id SqlMigrationServiceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Sql Migration Service Name: %q", id.SqlMigrationServiceName),
	}
	return fmt.Sprintf("Sql Migration Service (%s)", strings.Join(components, "\n"))
}
