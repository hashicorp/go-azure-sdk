package rbacs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AccountId{}

// AccountId is a struct representing the Resource ID for a Account
type AccountId struct {
	SubscriptionId      string
	ResourceGroupName   string
	DatabaseAccountName string
	RoleAssignmentId    string
}

// NewAccountID returns a new AccountId struct
func NewAccountID(subscriptionId string, resourceGroupName string, databaseAccountName string, roleAssignmentId string) AccountId {
	return AccountId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		DatabaseAccountName: databaseAccountName,
		RoleAssignmentId:    roleAssignmentId,
	}
}

// ParseAccountID parses 'input' into a AccountId
func ParseAccountID(input string) (*AccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccountId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.DatabaseAccountName, ok = parsed.Parsed["databaseAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseAccountName", *parsed)
	}

	if id.RoleAssignmentId, ok = parsed.Parsed["roleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "roleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseAccountIDInsensitively parses 'input' case-insensitively into a AccountId
// note: this method should only be used for API response data and not user input
func ParseAccountIDInsensitively(input string) (*AccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccountId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.DatabaseAccountName, ok = parsed.Parsed["databaseAccountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "databaseAccountName", *parsed)
	}

	if id.RoleAssignmentId, ok = parsed.Parsed["roleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "roleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateAccountID checks that 'input' can be parsed as a Account ID
func ValidateAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Account ID
func (id AccountId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/sqlRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DatabaseAccountName, strings.TrimPrefix(id.RoleAssignmentId, "/"))
}

// Segments returns a slice of Resource ID Segments which comprise this Account ID
func (id AccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDocumentDB", "Microsoft.DocumentDB", "Microsoft.DocumentDB"),
		resourceids.StaticSegment("staticDatabaseAccounts", "databaseAccounts", "databaseAccounts"),
		resourceids.UserSpecifiedSegment("databaseAccountName", "databaseAccountValue"),
		resourceids.StaticSegment("staticSqlRoleAssignments", "sqlRoleAssignments", "sqlRoleAssignments"),
		resourceids.ScopeSegment("roleAssignmentId", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
	}
}

// String returns a human-readable description of this Account ID
func (id AccountId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Database Account Name: %q", id.DatabaseAccountName),
		fmt.Sprintf("Role Assignment: %q", id.RoleAssignmentId),
	}
	return fmt.Sprintf("Account (%s)", strings.Join(components, "\n"))
}
