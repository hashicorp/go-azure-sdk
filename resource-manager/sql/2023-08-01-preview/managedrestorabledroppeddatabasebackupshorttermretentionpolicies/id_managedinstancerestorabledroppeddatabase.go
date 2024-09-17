package managedrestorabledroppeddatabasebackupshorttermretentionpolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ManagedInstanceRestorableDroppedDatabaseId{})
}

var _ resourceids.ResourceId = &ManagedInstanceRestorableDroppedDatabaseId{}

// ManagedInstanceRestorableDroppedDatabaseId is a struct representing the Resource ID for a Managed Instance Restorable Dropped Database
type ManagedInstanceRestorableDroppedDatabaseId struct {
	SubscriptionId              string
	ResourceGroupName           string
	ManagedInstanceName         string
	RestorableDroppedDatabaseId string
}

// NewManagedInstanceRestorableDroppedDatabaseID returns a new ManagedInstanceRestorableDroppedDatabaseId struct
func NewManagedInstanceRestorableDroppedDatabaseID(subscriptionId string, resourceGroupName string, managedInstanceName string, restorableDroppedDatabaseId string) ManagedInstanceRestorableDroppedDatabaseId {
	return ManagedInstanceRestorableDroppedDatabaseId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		ManagedInstanceName:         managedInstanceName,
		RestorableDroppedDatabaseId: restorableDroppedDatabaseId,
	}
}

// ParseManagedInstanceRestorableDroppedDatabaseID parses 'input' into a ManagedInstanceRestorableDroppedDatabaseId
func ParseManagedInstanceRestorableDroppedDatabaseID(input string) (*ManagedInstanceRestorableDroppedDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedInstanceRestorableDroppedDatabaseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedInstanceRestorableDroppedDatabaseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedInstanceRestorableDroppedDatabaseIDInsensitively parses 'input' case-insensitively into a ManagedInstanceRestorableDroppedDatabaseId
// note: this method should only be used for API response data and not user input
func ParseManagedInstanceRestorableDroppedDatabaseIDInsensitively(input string) (*ManagedInstanceRestorableDroppedDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedInstanceRestorableDroppedDatabaseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedInstanceRestorableDroppedDatabaseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedInstanceRestorableDroppedDatabaseId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedInstanceName, ok = input.Parsed["managedInstanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", input)
	}

	if id.RestorableDroppedDatabaseId, ok = input.Parsed["restorableDroppedDatabaseId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "restorableDroppedDatabaseId", input)
	}

	return nil
}

// ValidateManagedInstanceRestorableDroppedDatabaseID checks that 'input' can be parsed as a Managed Instance Restorable Dropped Database ID
func ValidateManagedInstanceRestorableDroppedDatabaseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedInstanceRestorableDroppedDatabaseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Instance Restorable Dropped Database ID
func (id ManagedInstanceRestorableDroppedDatabaseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/restorableDroppedDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.RestorableDroppedDatabaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Instance Restorable Dropped Database ID
func (id ManagedInstanceRestorableDroppedDatabaseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticRestorableDroppedDatabases", "restorableDroppedDatabases", "restorableDroppedDatabases"),
		resourceids.UserSpecifiedSegment("restorableDroppedDatabaseId", "restorableDroppedDatabaseIdValue"),
	}
}

// String returns a human-readable description of this Managed Instance Restorable Dropped Database ID
func (id ManagedInstanceRestorableDroppedDatabaseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Restorable Dropped Database: %q", id.RestorableDroppedDatabaseId),
	}
	return fmt.Sprintf("Managed Instance Restorable Dropped Database (%s)", strings.Join(components, "\n"))
}
