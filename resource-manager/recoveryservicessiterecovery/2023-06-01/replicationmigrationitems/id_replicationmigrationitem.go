package replicationmigrationitems

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ReplicationMigrationItemId{}

// ReplicationMigrationItemId is a struct representing the Resource ID for a Replication Migration Item
type ReplicationMigrationItemId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	VaultName                          string
	ReplicationFabricName              string
	ReplicationProtectionContainerName string
	ReplicationMigrationItemName       string
}

// NewReplicationMigrationItemID returns a new ReplicationMigrationItemId struct
func NewReplicationMigrationItemID(subscriptionId string, resourceGroupName string, vaultName string, replicationFabricName string, replicationProtectionContainerName string, replicationMigrationItemName string) ReplicationMigrationItemId {
	return ReplicationMigrationItemId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		VaultName:                          vaultName,
		ReplicationFabricName:              replicationFabricName,
		ReplicationProtectionContainerName: replicationProtectionContainerName,
		ReplicationMigrationItemName:       replicationMigrationItemName,
	}
}

// ParseReplicationMigrationItemID parses 'input' into a ReplicationMigrationItemId
func ParseReplicationMigrationItemID(input string) (*ReplicationMigrationItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationMigrationItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationMigrationItemId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "vaultName", *parsed)
	}

	if id.ReplicationFabricName, ok = parsed.Parsed["replicationFabricName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "replicationFabricName", *parsed)
	}

	if id.ReplicationProtectionContainerName, ok = parsed.Parsed["replicationProtectionContainerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "replicationProtectionContainerName", *parsed)
	}

	if id.ReplicationMigrationItemName, ok = parsed.Parsed["replicationMigrationItemName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "replicationMigrationItemName", *parsed)
	}

	return &id, nil
}

// ParseReplicationMigrationItemIDInsensitively parses 'input' case-insensitively into a ReplicationMigrationItemId
// note: this method should only be used for API response data and not user input
func ParseReplicationMigrationItemIDInsensitively(input string) (*ReplicationMigrationItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(ReplicationMigrationItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ReplicationMigrationItemId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "vaultName", *parsed)
	}

	if id.ReplicationFabricName, ok = parsed.Parsed["replicationFabricName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "replicationFabricName", *parsed)
	}

	if id.ReplicationProtectionContainerName, ok = parsed.Parsed["replicationProtectionContainerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "replicationProtectionContainerName", *parsed)
	}

	if id.ReplicationMigrationItemName, ok = parsed.Parsed["replicationMigrationItemName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "replicationMigrationItemName", *parsed)
	}

	return &id, nil
}

// ValidateReplicationMigrationItemID checks that 'input' can be parsed as a Replication Migration Item ID
func ValidateReplicationMigrationItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseReplicationMigrationItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Replication Migration Item ID
func (id ReplicationMigrationItemId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/replicationFabrics/%s/replicationProtectionContainers/%s/replicationMigrationItems/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.ReplicationFabricName, id.ReplicationProtectionContainerName, id.ReplicationMigrationItemName)
}

// Segments returns a slice of Resource ID Segments which comprise this Replication Migration Item ID
func (id ReplicationMigrationItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticReplicationFabrics", "replicationFabrics", "replicationFabrics"),
		resourceids.UserSpecifiedSegment("replicationFabricName", "replicationFabricValue"),
		resourceids.StaticSegment("staticReplicationProtectionContainers", "replicationProtectionContainers", "replicationProtectionContainers"),
		resourceids.UserSpecifiedSegment("replicationProtectionContainerName", "replicationProtectionContainerValue"),
		resourceids.StaticSegment("staticReplicationMigrationItems", "replicationMigrationItems", "replicationMigrationItems"),
		resourceids.UserSpecifiedSegment("replicationMigrationItemName", "replicationMigrationItemValue"),
	}
}

// String returns a human-readable description of this Replication Migration Item ID
func (id ReplicationMigrationItemId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Replication Fabric Name: %q", id.ReplicationFabricName),
		fmt.Sprintf("Replication Protection Container Name: %q", id.ReplicationProtectionContainerName),
		fmt.Sprintf("Replication Migration Item Name: %q", id.ReplicationMigrationItemName),
	}
	return fmt.Sprintf("Replication Migration Item (%s)", strings.Join(components, "\n"))
}
