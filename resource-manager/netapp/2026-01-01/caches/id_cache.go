package caches

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CacheId{})
}

var _ resourceids.ResourceId = &CacheId{}

// CacheId is a struct representing the Resource ID for a Cache
type CacheId struct {
	SubscriptionId    string
	ResourceGroupName string
	NetAppAccountName string
	CapacityPoolName  string
	CacheName         string
}

// NewCacheID returns a new CacheId struct
func NewCacheID(subscriptionId string, resourceGroupName string, netAppAccountName string, capacityPoolName string, cacheName string) CacheId {
	return CacheId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NetAppAccountName: netAppAccountName,
		CapacityPoolName:  capacityPoolName,
		CacheName:         cacheName,
	}
}

// ParseCacheID parses 'input' into a CacheId
func ParseCacheID(input string) (*CacheId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CacheId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CacheId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCacheIDInsensitively parses 'input' case-insensitively into a CacheId
// note: this method should only be used for API response data and not user input
func ParseCacheIDInsensitively(input string) (*CacheId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CacheId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CacheId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CacheId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NetAppAccountName, ok = input.Parsed["netAppAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", input)
	}

	if id.CapacityPoolName, ok = input.Parsed["capacityPoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "capacityPoolName", input)
	}

	if id.CacheName, ok = input.Parsed["cacheName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cacheName", input)
	}

	return nil
}

// ValidateCacheID checks that 'input' can be parsed as a Cache ID
func ValidateCacheID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCacheID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cache ID
func (id CacheId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/capacityPools/%s/caches/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetAppAccountName, id.CapacityPoolName, id.CacheName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cache ID
func (id CacheId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("netAppAccountName", "netAppAccountName"),
		resourceids.StaticSegment("staticCapacityPools", "capacityPools", "capacityPools"),
		resourceids.UserSpecifiedSegment("capacityPoolName", "capacityPoolName"),
		resourceids.StaticSegment("staticCaches", "caches", "caches"),
		resourceids.UserSpecifiedSegment("cacheName", "cacheName"),
	}
}

// String returns a human-readable description of this Cache ID
func (id CacheId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Net App Account Name: %q", id.NetAppAccountName),
		fmt.Sprintf("Capacity Pool Name: %q", id.CapacityPoolName),
		fmt.Sprintf("Cache Name: %q", id.CacheName),
	}
	return fmt.Sprintf("Cache (%s)", strings.Join(components, "\n"))
}
