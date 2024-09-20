package recoverablemanageddatabases

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ManagedInstanceRecoverableDatabaseId{})
}

var _ resourceids.ResourceId = &ManagedInstanceRecoverableDatabaseId{}

// ManagedInstanceRecoverableDatabaseId is a struct representing the Resource ID for a Managed Instance Recoverable Database
type ManagedInstanceRecoverableDatabaseId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ManagedInstanceName     string
	RecoverableDatabaseName string
}

// NewManagedInstanceRecoverableDatabaseID returns a new ManagedInstanceRecoverableDatabaseId struct
func NewManagedInstanceRecoverableDatabaseID(subscriptionId string, resourceGroupName string, managedInstanceName string, recoverableDatabaseName string) ManagedInstanceRecoverableDatabaseId {
	return ManagedInstanceRecoverableDatabaseId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ManagedInstanceName:     managedInstanceName,
		RecoverableDatabaseName: recoverableDatabaseName,
	}
}

// ParseManagedInstanceRecoverableDatabaseID parses 'input' into a ManagedInstanceRecoverableDatabaseId
func ParseManagedInstanceRecoverableDatabaseID(input string) (*ManagedInstanceRecoverableDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedInstanceRecoverableDatabaseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedInstanceRecoverableDatabaseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedInstanceRecoverableDatabaseIDInsensitively parses 'input' case-insensitively into a ManagedInstanceRecoverableDatabaseId
// note: this method should only be used for API response data and not user input
func ParseManagedInstanceRecoverableDatabaseIDInsensitively(input string) (*ManagedInstanceRecoverableDatabaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedInstanceRecoverableDatabaseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedInstanceRecoverableDatabaseId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedInstanceRecoverableDatabaseId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.RecoverableDatabaseName, ok = input.Parsed["recoverableDatabaseName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "recoverableDatabaseName", input)
	}

	return nil
}

// ValidateManagedInstanceRecoverableDatabaseID checks that 'input' can be parsed as a Managed Instance Recoverable Database ID
func ValidateManagedInstanceRecoverableDatabaseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedInstanceRecoverableDatabaseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Instance Recoverable Database ID
func (id ManagedInstanceRecoverableDatabaseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/recoverableDatabases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.RecoverableDatabaseName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Instance Recoverable Database ID
func (id ManagedInstanceRecoverableDatabaseId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceName"),
		resourceids.StaticSegment("staticRecoverableDatabases", "recoverableDatabases", "recoverableDatabases"),
		resourceids.UserSpecifiedSegment("recoverableDatabaseName", "recoverableDatabaseName"),
	}
}

// String returns a human-readable description of this Managed Instance Recoverable Database ID
func (id ManagedInstanceRecoverableDatabaseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Recoverable Database Name: %q", id.RecoverableDatabaseName),
	}
	return fmt.Sprintf("Managed Instance Recoverable Database (%s)", strings.Join(components, "\n"))
}
