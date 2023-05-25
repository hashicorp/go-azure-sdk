package manageddatabasemoveoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ManagedDatabaseMoveOperationResultId{}

// ManagedDatabaseMoveOperationResultId is a struct representing the Resource ID for a Managed Database Move Operation Result
type ManagedDatabaseMoveOperationResultId struct {
	SubscriptionId    string
	ResourceGroupName string
	LocationName      string
	OperationId       string
}

// NewManagedDatabaseMoveOperationResultID returns a new ManagedDatabaseMoveOperationResultId struct
func NewManagedDatabaseMoveOperationResultID(subscriptionId string, resourceGroupName string, locationName string, operationId string) ManagedDatabaseMoveOperationResultId {
	return ManagedDatabaseMoveOperationResultId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LocationName:      locationName,
		OperationId:       operationId,
	}
}

// ParseManagedDatabaseMoveOperationResultID parses 'input' into a ManagedDatabaseMoveOperationResultId
func ParseManagedDatabaseMoveOperationResultID(input string) (*ManagedDatabaseMoveOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedDatabaseMoveOperationResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedDatabaseMoveOperationResultId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseManagedDatabaseMoveOperationResultIDInsensitively parses 'input' case-insensitively into a ManagedDatabaseMoveOperationResultId
// note: this method should only be used for API response data and not user input
func ParseManagedDatabaseMoveOperationResultIDInsensitively(input string) (*ManagedDatabaseMoveOperationResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedDatabaseMoveOperationResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedDatabaseMoveOperationResultId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateManagedDatabaseMoveOperationResultID checks that 'input' can be parsed as a Managed Database Move Operation Result ID
func ValidateManagedDatabaseMoveOperationResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedDatabaseMoveOperationResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Database Move Operation Result ID
func (id ManagedDatabaseMoveOperationResultId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/locations/%s/managedDatabaseMoveOperationResults/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Database Move Operation Result ID
func (id ManagedDatabaseMoveOperationResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticManagedDatabaseMoveOperationResults", "managedDatabaseMoveOperationResults", "managedDatabaseMoveOperationResults"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Managed Database Move Operation Result ID
func (id ManagedDatabaseMoveOperationResultId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Managed Database Move Operation Result (%s)", strings.Join(components, "\n"))
}
