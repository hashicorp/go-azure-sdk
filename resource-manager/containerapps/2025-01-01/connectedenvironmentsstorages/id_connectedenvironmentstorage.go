package connectedenvironmentsstorages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ConnectedEnvironmentStorageId{})
}

var _ resourceids.ResourceId = &ConnectedEnvironmentStorageId{}

// ConnectedEnvironmentStorageId is a struct representing the Resource ID for a Connected Environment Storage
type ConnectedEnvironmentStorageId struct {
	SubscriptionId           string
	ResourceGroupName        string
	ConnectedEnvironmentName string
	StorageName              string
}

// NewConnectedEnvironmentStorageID returns a new ConnectedEnvironmentStorageId struct
func NewConnectedEnvironmentStorageID(subscriptionId string, resourceGroupName string, connectedEnvironmentName string, storageName string) ConnectedEnvironmentStorageId {
	return ConnectedEnvironmentStorageId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		ConnectedEnvironmentName: connectedEnvironmentName,
		StorageName:              storageName,
	}
}

// ParseConnectedEnvironmentStorageID parses 'input' into a ConnectedEnvironmentStorageId
func ParseConnectedEnvironmentStorageID(input string) (*ConnectedEnvironmentStorageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConnectedEnvironmentStorageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConnectedEnvironmentStorageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseConnectedEnvironmentStorageIDInsensitively parses 'input' case-insensitively into a ConnectedEnvironmentStorageId
// note: this method should only be used for API response data and not user input
func ParseConnectedEnvironmentStorageIDInsensitively(input string) (*ConnectedEnvironmentStorageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConnectedEnvironmentStorageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConnectedEnvironmentStorageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ConnectedEnvironmentStorageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ConnectedEnvironmentName, ok = input.Parsed["connectedEnvironmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "connectedEnvironmentName", input)
	}

	if id.StorageName, ok = input.Parsed["storageName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "storageName", input)
	}

	return nil
}

// ValidateConnectedEnvironmentStorageID checks that 'input' can be parsed as a Connected Environment Storage ID
func ValidateConnectedEnvironmentStorageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConnectedEnvironmentStorageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Connected Environment Storage ID
func (id ConnectedEnvironmentStorageId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.App/connectedEnvironments/%s/storages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ConnectedEnvironmentName, id.StorageName)
}

// Segments returns a slice of Resource ID Segments which comprise this Connected Environment Storage ID
func (id ConnectedEnvironmentStorageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApp", "Microsoft.App", "Microsoft.App"),
		resourceids.StaticSegment("staticConnectedEnvironments", "connectedEnvironments", "connectedEnvironments"),
		resourceids.UserSpecifiedSegment("connectedEnvironmentName", "connectedEnvironmentName"),
		resourceids.StaticSegment("staticStorages", "storages", "storages"),
		resourceids.UserSpecifiedSegment("storageName", "storageName"),
	}
}

// String returns a human-readable description of this Connected Environment Storage ID
func (id ConnectedEnvironmentStorageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Connected Environment Name: %q", id.ConnectedEnvironmentName),
		fmt.Sprintf("Storage Name: %q", id.StorageName),
	}
	return fmt.Sprintf("Connected Environment Storage (%s)", strings.Join(components, "\n"))
}
