package actions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ServerEndpointId{}

// ServerEndpointId is a struct representing the Resource ID for a Server Endpoint
type ServerEndpointId struct {
	SubscriptionId         string
	ResourceGroupName      string
	StorageSyncServiceName string
	SyncGroupName          string
	ServerEndpointName     string
}

// NewServerEndpointID returns a new ServerEndpointId struct
func NewServerEndpointID(subscriptionId string, resourceGroupName string, storageSyncServiceName string, syncGroupName string, serverEndpointName string) ServerEndpointId {
	return ServerEndpointId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		StorageSyncServiceName: storageSyncServiceName,
		SyncGroupName:          syncGroupName,
		ServerEndpointName:     serverEndpointName,
	}
}

// ParseServerEndpointID parses 'input' into a ServerEndpointId
func ParseServerEndpointID(input string) (*ServerEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerEndpointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.StorageSyncServiceName, ok = parsed.Parsed["storageSyncServiceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storageSyncServiceName", *parsed)
	}

	if id.SyncGroupName, ok = parsed.Parsed["syncGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "syncGroupName", *parsed)
	}

	if id.ServerEndpointName, ok = parsed.Parsed["serverEndpointName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverEndpointName", *parsed)
	}

	return &id, nil
}

// ParseServerEndpointIDInsensitively parses 'input' case-insensitively into a ServerEndpointId
// note: this method should only be used for API response data and not user input
func ParseServerEndpointIDInsensitively(input string) (*ServerEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(ServerEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ServerEndpointId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.StorageSyncServiceName, ok = parsed.Parsed["storageSyncServiceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "storageSyncServiceName", *parsed)
	}

	if id.SyncGroupName, ok = parsed.Parsed["syncGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "syncGroupName", *parsed)
	}

	if id.ServerEndpointName, ok = parsed.Parsed["serverEndpointName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serverEndpointName", *parsed)
	}

	return &id, nil
}

// ValidateServerEndpointID checks that 'input' can be parsed as a Server Endpoint ID
func ValidateServerEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServerEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Server Endpoint ID
func (id ServerEndpointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StorageSync/storageSyncServices/%s/syncGroups/%s/serverEndpoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StorageSyncServiceName, id.SyncGroupName, id.ServerEndpointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Server Endpoint ID
func (id ServerEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftStorageSync", "Microsoft.StorageSync", "Microsoft.StorageSync"),
		resourceids.StaticSegment("staticStorageSyncServices", "storageSyncServices", "storageSyncServices"),
		resourceids.UserSpecifiedSegment("storageSyncServiceName", "storageSyncServiceValue"),
		resourceids.StaticSegment("staticSyncGroups", "syncGroups", "syncGroups"),
		resourceids.UserSpecifiedSegment("syncGroupName", "syncGroupValue"),
		resourceids.StaticSegment("staticServerEndpoints", "serverEndpoints", "serverEndpoints"),
		resourceids.UserSpecifiedSegment("serverEndpointName", "serverEndpointValue"),
	}
}

// String returns a human-readable description of this Server Endpoint ID
func (id ServerEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Storage Sync Service Name: %q", id.StorageSyncServiceName),
		fmt.Sprintf("Sync Group Name: %q", id.SyncGroupName),
		fmt.Sprintf("Server Endpoint Name: %q", id.ServerEndpointName),
	}
	return fmt.Sprintf("Server Endpoint (%s)", strings.Join(components, "\n"))
}
