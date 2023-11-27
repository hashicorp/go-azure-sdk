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

	id := ReplicationMigrationItemId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
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

	id := ReplicationMigrationItemId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ReplicationMigrationItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.VaultName, ok = input.Parsed["vaultName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "vaultName", input)
	}

	if id.ReplicationFabricName, ok = input.Parsed["replicationFabricName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationFabricName", input)
	}

	if id.ReplicationProtectionContainerName, ok = input.Parsed["replicationProtectionContainerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationProtectionContainerName", input)
	}

	if id.ReplicationMigrationItemName, ok = input.Parsed["replicationMigrationItemName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "replicationMigrationItemName", input)
	}

	return nil
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
