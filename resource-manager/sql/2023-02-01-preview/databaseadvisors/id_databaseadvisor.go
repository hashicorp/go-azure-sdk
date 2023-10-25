package databaseadvisors

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DatabaseAdvisorId{}

// DatabaseAdvisorId is a struct representing the Resource ID for a Database Advisor
type DatabaseAdvisorId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServerName        string
	DatabaseName      string
	AdvisorName       string
}

// NewDatabaseAdvisorID returns a new DatabaseAdvisorId struct
func NewDatabaseAdvisorID(subscriptionId string, resourceGroupName string, serverName string, databaseName string, advisorName string) DatabaseAdvisorId {
	return DatabaseAdvisorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServerName:        serverName,
		DatabaseName:      databaseName,
		AdvisorName:       advisorName,
	}
}

// ParseDatabaseAdvisorID parses 'input' into a DatabaseAdvisorId
func ParseDatabaseAdvisorID(input string) (*DatabaseAdvisorId, error) {
	parser := resourceids.NewParserFromResourceIdType(DatabaseAdvisorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DatabaseAdvisorId{}

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

	if id.AdvisorName, ok = parsed.Parsed["advisorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "advisorName", *parsed)
	}

	return &id, nil
}

// ParseDatabaseAdvisorIDInsensitively parses 'input' case-insensitively into a DatabaseAdvisorId
// note: this method should only be used for API response data and not user input
func ParseDatabaseAdvisorIDInsensitively(input string) (*DatabaseAdvisorId, error) {
	parser := resourceids.NewParserFromResourceIdType(DatabaseAdvisorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DatabaseAdvisorId{}

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

	if id.AdvisorName, ok = parsed.Parsed["advisorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "advisorName", *parsed)
	}

	return &id, nil
}

// ValidateDatabaseAdvisorID checks that 'input' can be parsed as a Database Advisor ID
func ValidateDatabaseAdvisorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDatabaseAdvisorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Database Advisor ID
func (id DatabaseAdvisorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/servers/%s/databases/%s/advisors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServerName, id.DatabaseName, id.AdvisorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Database Advisor ID
func (id DatabaseAdvisorId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticAdvisors", "advisors", "advisors"),
		resourceids.UserSpecifiedSegment("advisorName", "advisorValue"),
	}
}

// String returns a human-readable description of this Database Advisor ID
func (id DatabaseAdvisorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Server Name: %q", id.ServerName),
		fmt.Sprintf("Database Name: %q", id.DatabaseName),
		fmt.Sprintf("Advisor Name: %q", id.AdvisorName),
	}
	return fmt.Sprintf("Database Advisor (%s)", strings.Join(components, "\n"))
}
