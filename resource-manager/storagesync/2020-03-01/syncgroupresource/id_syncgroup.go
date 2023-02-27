// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package syncgroupresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = SyncGroupId{}

// SyncGroupId is a struct representing the Resource ID for a Sync Group
type SyncGroupId struct {
	SubscriptionId         string
	ResourceGroupName      string
	StorageSyncServiceName string
	SyncGroupName          string
}

// NewSyncGroupID returns a new SyncGroupId struct
func NewSyncGroupID(subscriptionId string, resourceGroupName string, storageSyncServiceName string, syncGroupName string) SyncGroupId {
	return SyncGroupId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		StorageSyncServiceName: storageSyncServiceName,
		SyncGroupName:          syncGroupName,
	}
}

// ParseSyncGroupID parses 'input' into a SyncGroupId
func ParseSyncGroupID(input string) (*SyncGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(SyncGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SyncGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.StorageSyncServiceName, ok = parsed.Parsed["storageSyncServiceName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageSyncServiceName' was not found in the resource id %q", input)
	}

	if id.SyncGroupName, ok = parsed.Parsed["syncGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'syncGroupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseSyncGroupIDInsensitively parses 'input' case-insensitively into a SyncGroupId
// note: this method should only be used for API response data and not user input
func ParseSyncGroupIDInsensitively(input string) (*SyncGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(SyncGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SyncGroupId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.StorageSyncServiceName, ok = parsed.Parsed["storageSyncServiceName"]; !ok {
		return nil, fmt.Errorf("the segment 'storageSyncServiceName' was not found in the resource id %q", input)
	}

	if id.SyncGroupName, ok = parsed.Parsed["syncGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'syncGroupName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateSyncGroupID checks that 'input' can be parsed as a Sync Group ID
func ValidateSyncGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSyncGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Sync Group ID
func (id SyncGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.StorageSync/storageSyncServices/%s/syncGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.StorageSyncServiceName, id.SyncGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Sync Group ID
func (id SyncGroupId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Sync Group ID
func (id SyncGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Storage Sync Service Name: %q", id.StorageSyncServiceName),
		fmt.Sprintf("Sync Group Name: %q", id.SyncGroupName),
	}
	return fmt.Sprintf("Sync Group (%s)", strings.Join(components, "\n"))
}
