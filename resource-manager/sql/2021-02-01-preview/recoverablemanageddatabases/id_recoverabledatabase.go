package recoverablemanageddatabases

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RecoverableDatabaseId{}

// RecoverableDatabaseId is a struct representing the Resource ID for a Recoverable Database
type RecoverableDatabaseId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ManagedInstanceName     string
	RecoverableDatabaseName string
}

// NewRecoverableDatabaseID returns a new RecoverableDatabaseId struct
func NewRecoverableDatabaseID(subscriptionId string, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string) RecoverableDatabaseId {
	return RecoverableDatabaseId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ManagedInstanceName:     managedInstanceName,
		RecoverableDatabaseName: recoverableDatabaseName,
	}
}

// ParseRecoverableDatabaseID parses 'input' into a RecoverableDatabaseId
func ParseRecoverableDatabaseID(input string) (*RecoverableDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecoverableDatabaseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecoverableDatabaseId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.RecoverableDatabaseName, ok = parsed.Parsed["recoverableDatabaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recoverableDatabaseName", *parsed)
	}

	return &id, nil
}

// ParseRecoverableDatabaseIDInsensitively parses 'input' case-insensitively into a RecoverableDatabaseId
// note: this method should only be used for API response data and not user input
func ParseRecoverableDatabaseIDInsensitively(input string) (*RecoverableDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(RecoverableDatabaseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RecoverableDatabaseId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.RecoverableDatabaseName, ok = parsed.Parsed["recoverableDatabaseName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recoverableDatabaseName", *parsed)
	}

	return &id, nil
}

// ValidateRecoverableDatabaseID checks that 'input' can be parsed as a Recoverable Database ID
func ValidateRecoverableDatabaseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRecoverableDatabaseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Recoverable Database ID
func (id RecoverableDatabaseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/recoverableDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.RecoverableDatabaseName)
}

// Segments returns a slice of Resource ID Segments which comprise this Recoverable Database ID
func (id RecoverableDatabaseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticRecoverableDatabases", "recoverableDatabases", "recoverableDatabases"),
		resourceids.UserSpecifiedSegment("recoverableDatabaseName", "recoverableDatabaseValue"),
	}
}

// String returns a human-readable description of this Recoverable Database ID
func (id RecoverableDatabaseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Recoverable Database Name: %q", id.RecoverableDatabaseName),
	}
	return fmt.Sprintf("Recoverable Database (%s)", strings.Join(components, "\n"))
}
