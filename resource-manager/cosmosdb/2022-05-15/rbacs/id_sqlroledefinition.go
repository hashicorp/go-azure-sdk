package rbacs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SqlRoleDefinitionId{}

// SqlRoleDefinitionId is a struct representing the Resource ID for a Sql Role Definition
type SqlRoleDefinitionId struct {
	SubscriptionId      string
	ResourceGroupName   string
	DatabaseAccountName string
	RoleDefinitionId    string
}

// NewSqlRoleDefinitionID returns a new SqlRoleDefinitionId struct
func NewSqlRoleDefinitionID(subscriptionId string, resourceGroupName string, databaseAccountName string, roleDefinitionId string) SqlRoleDefinitionId {
	return SqlRoleDefinitionId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		DatabaseAccountName: databaseAccountName,
		RoleDefinitionId:    roleDefinitionId,
	}
}

// ParseSqlRoleDefinitionID parses 'input' into a SqlRoleDefinitionId
func ParseSqlRoleDefinitionID(input string) (*SqlRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(SqlRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SqlRoleDefinitionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.DatabaseAccountName, ok = parsed.Parsed["databaseAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseAccountName", *parsed)
	}

	if id.RoleDefinitionId, ok = parsed.Parsed["roleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "roleDefinitionId", *parsed)
	}

	return &id, nil
}

// ParseSqlRoleDefinitionIDInsensitively parses 'input' case-insensitively into a SqlRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseSqlRoleDefinitionIDInsensitively(input string) (*SqlRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(SqlRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SqlRoleDefinitionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.DatabaseAccountName, ok = parsed.Parsed["databaseAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseAccountName", *parsed)
	}

	if id.RoleDefinitionId, ok = parsed.Parsed["roleDefinitionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "roleDefinitionId", *parsed)
	}

	return &id, nil
}

// ValidateSqlRoleDefinitionID checks that 'input' can be parsed as a Sql Role Definition ID
func ValidateSqlRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSqlRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sql Role Definition ID
func (id SqlRoleDefinitionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/sqlRoleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DatabaseAccountName, id.RoleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Sql Role Definition ID
func (id SqlRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDocumentDB", "Microsoft.DocumentDB", "Microsoft.DocumentDB"),
		resourceids.StaticSegment("staticDatabaseAccounts", "databaseAccounts", "databaseAccounts"),
		resourceids.UserSpecifiedSegment("databaseAccountName", "databaseAccountValue"),
		resourceids.StaticSegment("staticSqlRoleDefinitions", "sqlRoleDefinitions", "sqlRoleDefinitions"),
		resourceids.UserSpecifiedSegment("roleDefinitionId", "roleDefinitionIdValue"),
	}
}

// String returns a human-readable description of this Sql Role Definition ID
func (id SqlRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Database Account Name: %q", id.DatabaseAccountName),
		fmt.Sprintf("Role Definition: %q", id.RoleDefinitionId),
	}
	return fmt.Sprintf("Sql Role Definition (%s)", strings.Join(components, "\n"))
}
