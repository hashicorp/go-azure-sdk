package manageddatabases

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ManagedDatabaseRestoreAzureAsyncOperationId{}

// ManagedDatabaseRestoreAzureAsyncOperationId is a struct representing the Resource ID for a Managed Database Restore Azure Async Operation
type ManagedDatabaseRestoreAzureAsyncOperationId struct {
	SubscriptionId string
	LocationName   string
	OperationId    string
}

// NewManagedDatabaseRestoreAzureAsyncOperationID returns a new ManagedDatabaseRestoreAzureAsyncOperationId struct
func NewManagedDatabaseRestoreAzureAsyncOperationID(subscriptionId string, locationName string, operationId string) ManagedDatabaseRestoreAzureAsyncOperationId {
	return ManagedDatabaseRestoreAzureAsyncOperationId{
		SubscriptionId: subscriptionId,
		LocationName:   locationName,
		OperationId:    operationId,
	}
}

// ParseManagedDatabaseRestoreAzureAsyncOperationID parses 'input' into a ManagedDatabaseRestoreAzureAsyncOperationId
func ParseManagedDatabaseRestoreAzureAsyncOperationID(input string) (*ManagedDatabaseRestoreAzureAsyncOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedDatabaseRestoreAzureAsyncOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedDatabaseRestoreAzureAsyncOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseManagedDatabaseRestoreAzureAsyncOperationIDInsensitively parses 'input' case-insensitively into a ManagedDatabaseRestoreAzureAsyncOperationId
// note: this method should only be used for API response data and not user input
func ParseManagedDatabaseRestoreAzureAsyncOperationIDInsensitively(input string) (*ManagedDatabaseRestoreAzureAsyncOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedDatabaseRestoreAzureAsyncOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedDatabaseRestoreAzureAsyncOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateManagedDatabaseRestoreAzureAsyncOperationID checks that 'input' can be parsed as a Managed Database Restore Azure Async Operation ID
func ValidateManagedDatabaseRestoreAzureAsyncOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedDatabaseRestoreAzureAsyncOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Database Restore Azure Async Operation ID
func (id ManagedDatabaseRestoreAzureAsyncOperationId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Sql/locations/%s/managedDatabaseRestoreAzureAsyncOperation/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.LocationName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Database Restore Azure Async Operation ID
func (id ManagedDatabaseRestoreAzureAsyncOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticManagedDatabaseRestoreAzureAsyncOperation", "managedDatabaseRestoreAzureAsyncOperation", "managedDatabaseRestoreAzureAsyncOperation"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Managed Database Restore Azure Async Operation ID
func (id ManagedDatabaseRestoreAzureAsyncOperationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Managed Database Restore Azure Async Operation (%s)", strings.Join(components, "\n"))
}
